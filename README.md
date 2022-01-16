## Blades

TLDR; An POC inventory system build with Go, GraphQL and MariaDB. Dockerized!

You can view a deployed version at http://blades.ericmarcantonio.com/, or you can build from source:

You will need:
1. docker
2. docker-compose

To build and run from source:

1. `git clone https://github.com/EricMarcantonio/blades.git`
2. `cd blades/`
3. `docker-compose -f docker-compose-local.yml up --build -d`
4. It's running on port 80 (if you need to change this, edit docker-compose-local.yml)


## Features

Blades is a simple POC inventory system, with CRUD and the ability to export as a CSV.
It has a simple React frontend, designed with Polaris from Shopify, and a dockerized backend running GraphQL with connections to MariaDB.

Repo Features
- Lean and concurrent queries, only what you request in GraphQL is queried from a table
- All complicated code is commented for easy reading
- A collection of postman requests are included for Web testing
- `go test` friendly testing is available for database requests


## Future changes
- Add TLS and authentication
- Set up kubernetes for scaling and managing a container fleet
- Add logging servers that can digest logs (ELK Software)

