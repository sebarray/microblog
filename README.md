# microblog


k  apply -f k8s/command_service.yaml

k kubectl apply -f k8s/mongo_statefulset.yaml

k kubectl apply -f k8s/mongo-pvc.yml

k kubectl apply -f k8s/query_service.yaml