# trivial_service

This is a simple web services example that is used for teaching.

There are two web servers in this project.

www is a simple frontend web server that only responds to requests to '/'.

backend is a simple backend web server that usually returns the current time.  Usually because it has an
approximately 2% error rate.

www for it's part has some issues too. For one it generally never exceeds 10 requests per second.  Secondly
it crashes if the backend server returns an error.

Both web servers accept various command-line parameters and this should be pretty obvious from the code.

#Instructions to Deploy the project

Note: This sections assumes that you have kubernetes environment setup and kubectl binary is configured.
If you want to build your own container. Please read the next section. This project comes with a default image smothiki/quantum:latest

* If you want to use the default image proceed to next step or continue here.Once you build the project and image is pushed to your registry.In manifests folder edit ```*_rs.yaml``` files image section inside container section to point to your image in the registry from above section. 
* Once you edit the image names makefile has target deploy which will deploy the service manifests first and replica set manifests next.

#Instructions to build the Project

Note: This section assumes that you have the Docker environment set up on your host machine to build docker containers.

* This project comes with a Makefile and has target docker-build. which builds all the necessary binaries which are statically linked for Alpine Docker image.
* For customization in rootfs/usr/bin we provided start bash script which you can customize hwo to start www or backend or proxy servers.
* Once build is done it generates quantum image. Tag the image to your registry using ```docker tag quantum <registry/imagename>``` and push the image using ```docker push <registry/imagename>```

#Instructions to test the project

* Once the project is deployed. Run ```kubectl get pod -o yaml <proxy-pod-name>``` and look for PODIP. log in to one of your hosts and use apache bench or just curl the podIP and port.
* check ```kubectl logs <proxy-pod-name>```

#Design
* Apart from frontend www and backend I have additonally added a Proxy service which proxies requests to Frontend and in the Proxy service for each connections I'm incrementing the count and as count exceeds 50 requests scaling the frontend ReplicaSet to actually sustain the load. Its a pretty basic and simple design. But doing a distributed proxy service requires storing the count and getting the count value from a distributed key value store and maintaining a lock on the key.

#GKE from GCS

* I used GKE google container engine from Google cloud services. Which provides a kubernetes cluster. But this example will work on any kubernetes setup. GKE is pretty easy.
  * once you create the cluster in your project get the cluster credentials <command>
  * Check if you can connect to the cluster just execute ```kubectl get nodes``` should give you nodes in the cluster.
  * Follow steps instructions to deploy then login to a node with ```gcloud compute ssh <node-anme>```.
  * Follow instructions to test the project.
