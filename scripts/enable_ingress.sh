kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/static/mandatory.yaml 

# You have to change it for other platforms as GCE, Azure, AWS
minikube addons enable ingress