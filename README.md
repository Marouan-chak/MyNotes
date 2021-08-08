# MyNotes
## Environment setup
OS: linux (Ubuntu/centos/raspbian tested)  
Docker version: +19.03.13  
Docker compose version: 1.29.2  
## Getting and running MyNotes  
```
# Clone MyNotes repository
git clone https://github.com/Marouan-chak/MyNotes.git
cd MyNotes
# Run docker-compose
docker-compose up -d
```
## Scaling the backend:  
In order to run multiple instances of MyNotes Backend, "--scale" option must be added to the docker-compose command:  
```
docker-compose up -d --scale mynote-be=5
```
## Kubernetes setup:
A helm chart is provided in order to deploy MyNotes App on a Kubernetes cluster. In order to do so:

### Requierments:
- A working K8s cluster
- Helm v3
- A docker registry
### Steps
Build and push docker images to a registry
Customize the `values.yaml` with correct values
Install the Chart
```
helm install mynotes MyNotes
```
## Architecture  
The following picture describe the different components of MyNotes App:

![alt text](https://github.com/Marouan-chak/MyNotes/blob/master/MyNotes.png?raw=true)
