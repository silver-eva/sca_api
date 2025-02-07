# Spy Cat Agency API
======================

## Overview

The Spy Cat Agency API is a RESTful API built with Golang and Postgres, containerized with Docker Compose. The API provides a set of endpoints for managing spy cats and missions.

## Features

* Containerized with Docker Compose for easy deployment
* Configurable with environment variables
* Swagger documentation available at `/docs/*` path
* Built with Golang and Postgres for high performance and reliability

## Endpoints

The API provides the following endpoints:

* `GET /cats`: List all spy cats
* `POST /cats`: Create a new spy cat
* `GET /cats/:id`: Retrieve a spy cat by ID
* `PUT /cats/:id`: Update a spy cat
* `DELETE /cats/:id`: Delete a spy cat
* `GET /missions`: List all missions
* `POST /missions`: Create a new mission
* `GET /missions/:id`: Retrieve a mission by ID
* `PUT /missions/:id`: Update a mission
* `DELETE /missions/:id`: Delete a mission

## Configuration

The API can be configured with file/environment variables. The following variables are supported:

* `DB_HOST`: The hostname or IP address of the Postgres database (default: localhost)
* `DB_PORT`: The port number of the Postgres database (default: 5432)
* `DB_USER`: The username to use for the Postgres database (default: postgres)
* `DB_PASS`: The password to use for the Postgres database (default: postgres)
* `DB_NAME`: The name of the Postgres database (default: postgres)

## Running the API

To run the API, use the following command:

```bash
docker-compose up
```

This will start the API and make it available at `http://localhost:8000`.

## Swagger Documentation

Swagger documentation is available at `/docs/*` path. You can access it by visiting `http://localhost:8000/docs` in your web browser.

## Start project

```bash
sudo docker compose --env-file .env -f docker-compose.yaml up --build -d
```

Command will build project and start it in detached mode.

## Contributing

Contributions are welcome! If you'd like to contribute to the API, please fork the repository and submit a pull request.

## License

The Spy Cat Agency API is licensed under the MIT License.