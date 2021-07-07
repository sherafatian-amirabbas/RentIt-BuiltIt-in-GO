#!/bin/sh
docker image build -f Backend/Dockerfile_backend -t localhost:5000/backend Backend
docker push localhost:5000/backend

docker image build -f Backend/Dockerfile_test -t localhost:5000/backend_test Backend
docker push localhost:5000/backend_test

docker image build -f Frontend/Dockerfile_frontend -t localhost:5000/frontend Frontend
docker push localhost:5000/frontend

docker image build -f Frontend/Dockerfile_test -t localhost:5000/frontend_test Frontend
docker push localhost:5000/frontend_test

docker image build -f docker-compose-proxy-assets/Dockerfile_nginx -t localhost:5000/proxy docker-compose-proxy-assets
docker push localhost:5000/proxy
