# Kind

> <https://kind.sigs.k8s.io/docs/user/configuration/>

## Kind e Kubectl

```bash
# cria cluster a partir de um arquivo de config
kind create cluster --config=kind.yaml --name=fullcycle

# muda de contexto 
kubectl cluster-info --context kind-fullcycle

kind get nodes
kind get configs
kind get clusters

cat ~/.kube/config
kubectl get nodes
kubectl cluster-info –context kind-kind

kubectl config get-clusters
kubectl config use-context kind-fullcycle
kubectl get nodes
```

## Go

```bash
# Roda a app Go
go run server.go
```

## Docker

```bash
# Cria o build da app
docker build -t rafaelbertelli89/hello-go .
docker run --rm -p 8080:8080 rafaelbertelli89/hello-go
docker push rafaelbertelli89/hello-go
```
