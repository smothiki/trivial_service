# trivial_service

This is a simple web services example that is used for teaching.

There are two web servers in this project.

www is a simple frontend web server that only responds to requests to '/'.

backend is a simple backend web server that usually returns the current time.  Usually because it has an
approximately 2% error rate.

www for it's part has some issues too. For one it generally never exceeds 10 requests per second.  Secondly
it crashes if the backend server returns an error.

Both web servers accept various command-line parameters and this should be pretty obvious from the code.


Note: Project assumes that you have go version >= 1.6 and knowledge of kubernetes.

#Instructions to test the project

* Once the project is deployed. Run ```kubectl get pod -o yaml <proxy-pod-name>``` and look for PODIP. log in to one of your hosts and use apache bench or just curl the podIP and port.
* check ```kubectl logs <proxy-pod-name>```

#GKE from GCS

* I used GKE google container engine from Google cloud services. Which provides a kubernetes cluster. But this example will work on any kubernetes setup. GKE is pretty easy.
  * once you create the cluster in your project get the cluster credentials <command>
  * Check if you can connect to the cluster just execute ```kubectl get nodes``` should give you nodes in the cluster.
  * Follow steps instructions to deploy then login to a node with ```gcloud compute ssh <node-name>```.
  * Follow instructions to test the project.
