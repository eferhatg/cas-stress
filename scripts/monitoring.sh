


# kubectl port-forward $(kubectl get  pods --selector=app=grafana -n  cas --output=jsonpath="{.items..metadata.name}") -n cas  3000
# kubectl port-forward -n cas prometheus-caspromet-prometheus-opera-prometheus-0 9090