#!/bin/bash

APISERVER=$(kubectl config view --minify -o jsonpath='{.clusters[0].cluster.server}')
TOKEN=$(kubectl get secret -n job-operator-system $(kubectl get serviceaccount job-operator-controller-manager -n job-operator-system -o jsonpath='{.secrets[0].name}') -o jsonpath='{.data.token}' | base64 --decode )

curl -X POST $APISERVER/apis/job.vr.fmwk.com/v1alpha1/namespaces/job-operator-system/vrtestjobs --header "Content-Type: application/json" --header "Authorization: Bearer $TOKEN" --insecure --data "@kubernetes-crd-request.json"

#curl -s $APISERVER/openapi/v2  --header "Authorization: Bearer $TOKEN" --cacert /home/valkyrie/.minikube/ca.crt | jq . > openapispecs.json