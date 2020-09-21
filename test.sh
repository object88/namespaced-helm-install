#!/usr/bin/env bash

docker build . --tag object88/namespaced-helm-install:latest
docker push object88/namespaced-helm-install:latest

kubens default

helm delete nhi --namespace nhi
helm install nhi charts/namespaced-helm-install --namespace nhi --wait

kubectl get pods --namespace nhi