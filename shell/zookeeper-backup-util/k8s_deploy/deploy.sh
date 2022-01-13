#!/bin/bash

_namespace_exist=$(kubectl get namespaces | grep -c "utility")

[[ $_namespace_exist -eq 0 ]] && kubectl apply -f zkbackup-namespace.yaml

# Create the base resources
kubectl apply -f zkbackup-service-account.yaml

# Create the config map
kubectl apply -f zkbackup-configmap.yaml

# Create the cronjob
kubectl apply -f zkbackup-cronjob.yaml
