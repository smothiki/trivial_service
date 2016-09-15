# Project Desciption

Apart from trivial_service original files a lot other files and folder have been added in the project.
* Docs Folder contains Docs about Project descirption. Architecture , Installations and other scripts about
output and logs.
* manifests folder contains kubernetes manifest files like replica sets and services in yaml format.
* Proxy folder contains Proxy go lang server.
* rootfs is the root file system that we add in to the Alpine Image and also has Dockerfile we use to build the project.
* Vendor Folder contains Golang dependencies to build the files.
* glide.lock is the file used by glide container to fetch the dependencies.
* glide.yaml is the file used by glide container to set the dependencies.
* Makefile target has scripts to build and deploy the project.
