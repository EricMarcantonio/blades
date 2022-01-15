## Blades

TLDR; An inventory system build with Go, GraphQL and MariaDB. Dockerized!

You can view a deployed version at http://blades.ericmarcantonio.com/, or you can build from source:

You will need:
1. docker
2. docker-compose

To build and run from source:

1. `git clone https://github.com/EricMarcantonio/blades.git`
2. `cd blades/`
3. `docker-compose up --build -d`
4. It's running on port 80 (if you need to change this, edit docker-compose)
