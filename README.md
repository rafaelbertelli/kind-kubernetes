# Kind

> <https://kind.sigs.k8s.io/docs/user/configuration/>

> <https://kubernetes.io/docs/tasks/configure-pod-container/>

## Kind e Kubectl

```bash
https://kubernetes.io/docs/reference/kubectl/cheatsheet/

# cria cluster a partir de um arquivo de config
kind create cluster --config=k8s/kind.yaml --name=fullcycle

# muda de contexto 
kubectl cluster-info --context kind-fullcycle

kind get nodes
kind get configs
kind get clusters

cat ~/.kube/config
kubectl get nodes
kubectl cluster-info –context kind-kind

# seta um determinado contexto
kubectl config get-clusters
kubectl config use-context kind-fullcycle
kubectl get nodes

# aplica o arquivo de configuração criando meu pod
kubectl apply -f k8s/pod.yaml

# comandos usuais
kubectl port-forward pod/goserver 8080:8080
kubectl delete pod goserver

kubectl apply -f k8s/replicaset.yaml
kubectl get pods
kubectl get replicaset
kubectl delete pods 'nome do pod'
kubectl get pods


```

## Go

```bash
go run server.go
```

## Docker

```bash
docker build -t rafaelbertelli89/hello-go .
docker run --rm -p 8080:8080 rafaelbertelli89/hello-go
docker push rafaelbertelli89/hello-go
```
