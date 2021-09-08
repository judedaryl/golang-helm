# Getting Started

This is a quick-start on deploying an API written in golang to a **kubernetes** cluster using **helm**. If you only want to learn about the deployment process you can fully ignore the golang code and focus on the deployment side of things by using an existing image published to docker hub [judedaryl/myapi](https://hub.docker.com/repository/docker/judedaryl/myapi).


## Requirements
* Docker Desktop

## Minikube

### Install
Check for the installation instructions here https://minikube.sigs.k8s.io/docs/start/

> Note that installing minikube also installs kubectl so you don't need to install that separately

#### macOS
```sh
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64
sudo install minikube-darwin-amd64 /usr/local/bin/minikube
```

### Setup

You can start the cluster by running:

```sh
minikube start
```

Add an **nginx ingress**

```
minikube addons enable ingress
```

###

From this point, **minikube** should be operational and you should have a live kubernetes cluster in your local machine. You can interact with your cluster using kubectl or you can run a gen-purpose kubernetes dashboard.

```sh
# interact with the cluster using kubectl
kubectl get po -A

# run the gen-purpose dashboard
minikube dashboard
```

## Helm

We'll use helm to deploy our app to the cluster.
> Helm helps you manage Kubernetes applications — Helm Charts help you define, install, and upgrade even the most complex Kubernetes application.

### Install

Helm has an installer script that will automatically grab the latest version of Helm and install it locally.


```sh
curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
```

## Deploy

### Use the existing image

There is an existing docker image that you can use available in docker hub [judedaryl/myapi](https://hub.docker.com/repository/docker/judedaryl/myapi). By default this is already referenced in ``chart/values.yml`` in ``line 8``

1. When installing for the first time run:

    ```sh
    helm install myapi chart/ --values chart/values.yaml
    ```
2. When making changes to the chart template or values

    ```sh
    helm upgrade myapi chart/ --values chart/values.yaml
    ```
3. Check the status of the pod

    ```sh
    kubectl get pods

    ### or grep it
    kubectl get pods | grep myapi
    ```

4. When the pod has a status of **Running**, create a tunnel to your cluster by running 

    ```sh
    minikube service myapi --url
    ```
    This will printout something like this on your terminal.

    ```
    🏃  Starting tunnel for service myapi.
    |-----------|-------|-------------|------------------------|
    | NAMESPACE | NAME  | TARGET PORT |          URL           |
    |-----------|-------|-------------|------------------------|
    | default   | myapi |             | http://127.0.0.1:{PORT} |
    |-----------|-------|-------------|------------------------|
    http://127.0.0.1:{PORT}
    ```

    Take note of the value of **{PORT}**.

5. In this repository go to the file ``test.rest`` and update ``line 1`` with the port value from step 4.

### Build from source code

You don't need to install any golang dependencies for this part as everything is already configured in the container specified in the ``Dockerfile``. You will however need to have a [docker hub](https://hub.docker.com) account.

1. Create an image from the Dockerfile
    ```sh
    docker build --tag {docker_username}/myapi .
    ```
2. Push the image to docker hub
    ```sh
    docker push {docker_username}/myapi
    ```
3. Update the chart references to the image name in ``chart/values.yml`` ``line 8`` to

    ```
      repository: {docker_username}/myapi
    ```
4. Follow the steps in **Using an existing image**