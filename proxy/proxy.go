/* A simple proxy program taken from the following URL
http://www.darul.io/post/2015-07-22_go-lang-simple-reverse-proxy*/

package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"k8s.io/kubernetes/pkg/api"
	kcl "k8s.io/kubernetes/pkg/client/unversioned"
)

var count int32

// Prox is a proxy struct
type Prox struct {
	// target url of reverse proxy
	target *url.URL
	// instance of Go ReverseProxy thatwill do the job for us
	proxy *httputil.ReverseProxy
	// count number of call incoming calls
	counter chan int32
	// clinet to configure replica sets
	rsClient kcl.ReplicaSetInterface
	sync.Mutex
}

// New a small factory method
func New(target string, rscli kcl.ReplicaSetInterface) *Prox {
	url, _ := url.Parse(target)
	// you should handle error on parsing
	c := make(chan int32)
	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url), counter: c, rsClient: rscli}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")
	// call to magic method from ReverseProxy object
	p.Lock()
	count++
	p.counter <- count
	p.Unlock()
	p.proxy.ServeHTTP(w, r)
}

func (p *Prox) counterReset() {
	tickChan := time.NewTicker(time.Millisecond * 1000).C
	for {
		select {
		case m := <-p.counter:
			var scalefactor int32
			if m > 50 {
				scalefactor = (m / 10)
			} else if m <= 50 {
				scalefactor = 5
			}
			fmt.Println(scalefactor)
			rsfrontend, err := p.rsClient.Get("frontend")
			if err != nil {
				fmt.Println("replicaset not found")
			}
			rsfrontend.Spec.Replicas = scalefactor
			rsfrontend, err = p.rsClient.Update(rsfrontend)
			if err != nil {
				fmt.Println("replicaset not updated")
			}
		case <-tickChan:
			p.Lock()
			count = 0
			p.Unlock()
		}
	}
}

func main() {
	// come constants and usage helper
	const (
		defaultPort        = "6000"
		portUsage          = "Port to listen on"
		defaultTargetUsage = "default redirect url, 'http://127.0.0.1:8080'"
		defaultTarget      = "http://127.0.0.1:8080"
	)
	var (
		port string
		url  string
	)
	// flags
	flag.StringVar(&url, "url", defaultTarget, defaultTargetUsage)
	flag.StringVar(&port, "port", defaultPort, portUsage)

	flag.Parse()

	fmt.Printf("server will run on : %s\n", port)
	fmt.Printf("redirecting to :%s\n", url)
	//frontend := os.Getenv("FRONTEND_SERVICE_HOST")
	kubeClient, err := kcl.NewInCluster()
	if err != nil {
		fmt.Println("check for api auth volume mount")
	}
	rsCli := kubeClient.Extensions().ReplicaSets(api.NamespaceDefault)
	// proxy
	proxy := &Prox{}
	proxy = New(url, rsCli)
	// server
	go proxy.counterReset()
	http.HandleFunc("/", proxy.handle)
	http.ListenAndServe(":"+port, nil)
}
