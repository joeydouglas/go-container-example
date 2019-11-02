#### Go Container Example
This is a simple Go http server demonstrating the use of containers for multi-environment deployments and multi-stage builds. 

## Steps:
1. ```docker build --no-cache -t go-container-example:v1.0 .```
2. ```docker run --rm --name go-container-example -e APP_ENV="dev" -p 8887:80 go-container-example:v1.0```