# ESI-Homework1

Project Overview: https://docs.google.com/document/d/13FwCWFViaeDJ6JYXD6Pzfjf5vReP47ZDiZakf-0EL5M/edit

# Usage
Make sure that docker registry is running in your local machine \
**docker-compose -f docker-compose-registry.yml up -d** \
After first time compose it restarts automatically after system restart. So it is needed to do that only once

Other docker compose files use that registry get their images

Use script **build.sh** to build images and push them to registry. It is great start to CI/CD system because CI can run that script and new images are available in the registry


### These commands use previously mentioned registry to get their image :
**docker-compose -f docker-compose-app.yml up** \
**docker-compose -f docker-compose-proxy.yml up**
