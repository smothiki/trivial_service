# ApacheBench results
$ ab -r -n 150 -c 5 http://10.0.2.3:6000/ // Proxy server POD IP and PORT
This is ApacheBench, Version 2.3 <$Revision: 655654 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 10.0.2.3 (be patient).....done


Server Software:        
Server Hostname:        10.0.2.3
Server Port:            6000

Document Path:          /
Document Length:        32 bytes

Concurrency Level:      5
Time taken for tests:   39.098 seconds
Complete requests:      150
Failed requests:        7
   (Connect: 0, Receive: 0, Length: 7, Exceptions: 0)
Write errors:           0
Non-2xx responses:      7
Total transferred:      25252 bytes
HTML transferred:       4576 bytes
Requests per second:    3.84 [#/sec] (mean)
Time per request:       1303.272 [ms] (mean)
Time per request:       260.654 [ms] (mean, across all concurrent requests)
Transfer rate:          0.63 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.2      1       1
Processing:   101 1276 885.2   1996    2059
Waiting:      101 1276 885.2   1996    2059
Total:        102 1277 885.1   1996    2060

Percentage of the requests served within a certain time (ms)
  50%   1996
  66%   2000
  75%   2000
  80%   2001
  90%   2003
  95%   2009
  98%   2032
  99%   2037
 100%   2060 (longest request)

# PODS after getting auto scaled

$ kubectl get pods -w
NAME             READY     STATUS    RESTARTS   AGE
backend-9m5m6    1/1       Running   0          11m
frontend-0aebw   1/1       Running   0          52s
frontend-28o6t   1/1       Running   1          44s
frontend-3224k   1/1       Running   0          11m
frontend-7bbsx   1/1       Running   0          28s
frontend-7cb3n   1/1       Running   0          32s
frontend-7wfck   1/1       Running   1          11m
frontend-d6bek   1/1       Running   0          1m
frontend-edpx6   1/1       Running   1          56s
frontend-i5k3v   1/1       Running   1          11m
frontend-imt56   1/1       Running   1          11m
frontend-lp8f3   1/1       Running   0          1m
frontend-mo9kg   1/1       Running   1          11m
frontend-nd9zd   1/1       Running   0          40s
frontend-tk02y   1/1       Running   0          36s
frontend-z5lx3   1/1       Running   0          48s
proxy-jvpe0      1/1       Running   0          2m

# Logs from Proxy server
$  kubectl logs -f proxy-jvpe0
server will run on : 6000
redirecting to :http://10.3.248.230:80
1
2 // Counter increment status
3
4
......
scaling application 13
131
scaling application 13
132
scaling application 13
133
scaling application 13
134
scaling application 13
135
scaling application 13
136
scaling application 13
2016/09/15 22:58:14 http: proxy error: EOF
137
scaling application 13
2016/09/15 22:58:15 http: proxy error: EOF
138
scaling application 13
139
scaling application 13
140
scaling application 14
141
scaling application 14
142
scaling application 14
143
scaling application 14
144
scaling application 14
145
scaling application 14
146
scaling application 14
147
scaling application 14
148
scaling application 14
149
scaling application 14
150
scaling application 15
