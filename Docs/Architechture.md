# Design

# challenges

* This project assumes that you are familiar with kubernetes architecture.
* Given a Frontend server called www and backend server which gives time with 2% error rate.
* If there is an error from Backend Frontend dies and Frontend has a throughput of 10 requests per Minute.
* The architecture should be accessible through single IP or DNS. To cope up with this Developed a proxy server which is a single reverse proxy proxies traffic to Frontend.

# Components

* Frontend and Backend servers are launches as replica sets which is a kubernetes extension an advanced Replication controller type. Each replica set was given appropriate labels to get backed by kubernetes services.
* Frontend and Backend runs on ports 7000 and 8000. Proxy server runs on port 6000.

# Functionality

* Proxy service which proxies requests to Frontend service and Frontend service talks to backend service.
* Frontend service is a HA proxy endpoint which round robins request to pods in Frontend replicaset.
* A replica set makes sure that even though there is a POD failure it restarts the POD in the same or different node.
* In the Proxy service for each connections a counter is incremented and for every 100 milliseconds we are scaling the frontend according to the counter. If there is no Load the replica set of Fronteneds is scaled back to 5.
* Counter is reset for every 10 seconds.
 
