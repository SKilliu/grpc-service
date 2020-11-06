### gRPC-service

This is simple gRPC service for saving coordinates from GPS device to the database.
In this repo are two services. The first one is gRPC client that communicates with GPS device with REST API and further sends data to the gRPC server. The last one saves data to the database.

**Pre-requirements**

- Go v1.11.4+
- Docker version 18.09.06+

#### **How to run**

*Start the database, download migrations to db and run the gRPC server*

`make compose-up`

`make migration-up`

`make run`

*After this you go to client folder, generate Swagger documentation and starting the gRPC client.*

`make docs-generat`

`make run`

#### Commands list

- make docs-generate -- *generate the swagger docs*

- make compose-up -- *up postgreSQL in docker container*

- make build -- *build the project*

- make run -- *run the project*

- migration-up -- *upload migrations to db*

- migration-down -- *rollback all migrations*

#### **Swagger documentation url**

{{baseUrl}}/swagger/index.html