# Problem statement

Write a Go Program to capture any changes in a configMap deloyed within a namespace by watching over it and enforcing the change on the service consuming the configMap without having to restart the service. 

## Importance
This is a very helpful design pattern or implementation best practices in cloud native applications (or microservices deployed on K8s) where they consume the configuration metadata like DB host, server host-port, loglevels etc from the ConfigMap. Any changes to the ConfigMap should be automatically reflected in the service at almost real time without having to restart the service. 

## Example implementation
In this example implementation, a http server is started consuming its configuration metadata from a K8s ConfigMap. Any changes to the configuration shall result in almost instant reflection within the server without having to restart the server.

## Steps to run 

1. Clone this repository and navigate the folder

    ```bash
    cd example-06
    ```

2. Ensure the K8s cluster is running & then create the configmap

    ```bash
    k apply -f manifests/configmap.yaml
    ```

3. Start the service

    ```bash
    go run cmd/main.go
    ```

4. The output will be something similar to the below

    ```bash
    go run cmd/main.go 
    2024/03/24 17:01:38 starting the service
    2024/03/24 17:01:38 starting the server at localhost:8080
    got / request
    got / request

    # update the server-config configmap to use different port
    2024/03/24 17:01:59 server config changes observed. Restarting the server with new address localhost:8082
    2024/03/24 17:01:59 starting the server at localhost:8082
    got / request
    got / request
    ```