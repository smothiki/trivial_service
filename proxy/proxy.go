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

//A global variable to keep track of number of requests
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
	//Mutex for locking
	sync.Mutex
}

// New a small factory method
func New(target string, rscli kcl.ReplicaSetInterface) *Prox {
	url, _ := url.Parse(target)
	// should handle error on parsing
	c := make(chan int32)
	return &Prox{target: url, proxy: httputil.NewSingleHostReverseProxy(url), counter: c, rsClient: rscli}
}

func (p *Prox) handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-GoProxy", "GoProxy")
	p.Lock()
	count++
	p.counter <- count
	p.Unlock()
	// call to magic method from ReverseProxy object
	p.proxy.ServeHTTP(w, r)
}

//counterReset resets the count global variable for every 500 milliseconds
//gets the count from handle func through counter chan and calculates the scale factor
func (p *Prox) counterReset() {
	//tickerchan sends a signal for every 500 milliseconds
	// tickChanReset := time.NewTicker(time.Millisecond * 50).C
	tickChanScale := time.NewTicker(time.Second * 60).C
	var scalefactor int32
	for {
		select {
		//if we receive an updated count this case gets executed and calculates scale factor
		case m := <-p.counter:
			fmt.Println(m)
			if m > 50 {
				scalefactor = (m / 10)
			} else if m <= 50 {
				scalefactor = 5
			}
			fmt.Printf("scaling applicaiton %d\n", scalefactor)
			rsfrontend, err := p.rsClient.Get("frontend")
			if err != nil {
				fmt.Println("replicaset not found")
			}
			rsfrontend.Spec.Replicas = scalefactor
			rsfrontend, err = p.rsClient.Update(rsfrontend)
			if err != nil {
				fmt.Println("replicaset not updated")
			}
		case <-tickChanScale:
			p.Lock()
			count = 0
			scalefactor = 5
			p.counter <- count
			p.Unlock()
		}
	}
}

func main() {
	// constants and usage helper
	const (
		defaultPort        = "6000"
		portUsage          = "Port to listen on"
		defaultTargetUsage = "default redirect url, 'http://127.0.0.1:8080'"
		defaultTarget      = "http://127.0.0.1:8080"
	)
	//port and url store flag values
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
	//kubeclinet which reads default credentials mounted on each pod and talks to API
	kubeClient, err := kcl.NewInCluster()
	if err != nil {
		fmt.Println("check for api auth volume mount")
	}
	//rsCli is replicaset clinet that gives acces to replicaset in the given Namespace
	rsCli := kubeClient.Extensions().ReplicaSets(api.NamespaceDefault)
	// proxy
	proxy := &Prox{}
	proxy = New(url, rsCli)
	//calling counterReset in a seprate go routine.
	go proxy.counterReset()
	// server
	http.HandleFunc("/", proxy.handle)
	http.ListenAndServe(":"+port, nil)
}
