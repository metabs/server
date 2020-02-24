# Server

Server for unprogettosenzanomecheforseinizieremo.

Packages:
 
- customer: contains aggregate root for the customer and it main behaviours
- workspace: contains aggregate root for the workspace, collection, tag (entity) and their main behaviours
- internal: glue together the external layers (HTTP, database and so on) with the domain models 
- cmd: contains the main files to build the binaries
- tests: contains BDD tests
- vendor: contains the dependencies

Other directories:

- data: contains all the volumes bind for the docker compose
- db: contains the database migrations
- deploy: contains the deployment files for K8S
- .github: contains the github action to empower CI/CD

# Development

## Build app

- make

## Lint

- make lint

## Run tests

- cp .env.dist .env
- make tests

## Run server

- docker-compose up -d
