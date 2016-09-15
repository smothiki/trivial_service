# Instructions to Deploy the project

Note: This sections assumes that you have kubernetes environment setup and kubectl binary is configured.
If you want to build your own container. Please read the next section. This project comes with a default image smothiki/quantum:latest

* If you want to use the default image proceed to next step or continue here.Once you build the project and image is pushed to your registry.In manifests folder edit ```*_rs.yaml``` files image section inside container section to point to your image in the registry from above section.
* Once you edit the image names makefile has target ```deploy``` which will deploy the service manifests first and replica set manifests next.

# Instructions to build the Project

Note: This section assumes that you have the Docker environment set up on your host machine to build docker containers.

* This project comes with a Makefile and has target ```docker-build```. which builds all the necessary binaries which are statically linked for Alpine Docker image.
* For customization in rootfs/usr/bin we provided start bash script which you can customize hwo to start www or backend or proxy servers.
* Once build is done it generates quantum image. Tag the image to your registry using ```docker tag quantum <registry/imagename>``` and push the image using ```docker push <registry/imagename>```
