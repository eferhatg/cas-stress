
kubectl delete deployments  -n cas --all
kubectl delete pods -n cas --all
kubectl delete services -n cas --all

kubectl apply -f ./k8s/namespace.yml
kubectl apply -f ./k8s --namespace cas