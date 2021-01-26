# Description

A purpose of this repository is to keep a track of steps performed during doing a homework from Equal Experts.

I decided to go with *Mini Kube* project.

Any `make ...` command used in below description should be launched from project main directory.

### Running
In order to start this project you will need minikube and kubectl installed so CLI environment has `minikube` and `kubectl` binaries available.

1. Run `make minikube-start`
2. Run `make apply-k8s-resources`

Now you can test with curl using:
`curl -H 'Host: ee-homework-app.local' http://$(minikube ip)/`

To make test app available in the browser you need to modify your host `/etc/hosts` file and add entry for `ee-homework-app.local` pointing to whatever `minikube ip` returns.

### Development

If you want to introduce changes to image used for deployment, you should set your env up and running and then:
1. Change imagePullPolicy for app deployment to Never: `kubectl -nee-homework-dev patch deployment app -p '{"spec":{"template":{"spec":{"containers":[{"name":"ee-homework","imagePullPolicy":"Never"}]}}}}`
2. import minikube docker env variables: `eval $(minikube docker-env)`
3. build image with minikube docker daemon context: `make image-build`
4. restart pods with: `make restart-pods`

At this point you should have latest local changes baked in into image built inside minikube docker space, available for minikube cluster and already applied.

To revert above, run:
1. revert docker env variables `eval $(minikube docker-env -u)` changes
2. Change imagePullPolicy for app deployment to Always: `kubectl -nee-homework-dev patch deployment app -p '{"spec":{"template":{"spec":{"containers":[{"name":"ee-homework","imagePullPolicy":"Always"}]}}}}`

To push out changes more permanently, push out latest image to registry with `make image-push`

### ToDo

* Introduce secrets management
* Introduce support for private docker registry
* Extract base build image into separate creation process so we limit calls like "go get"
* Make sure to introduce PodAntiAffinity for non-DEV scenarios to prevent our app multiple PODS from being provisioned on same NODE.
* When needed for any reason, put extra effort into enabling multiple nodes minikube setup with application spread among many nodes - to more accurately simulate production environment. Out of the box, just increasing nodes count does not work.
