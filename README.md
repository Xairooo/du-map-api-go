# Go API Server for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This server was generated by the [openapi-generator]
(https://openapi-generator.tech) project.
By using the [OpenAPI-Spec](https://github.com/OAI/OpenAPI-Specification) from a remote server, you can easily generate a server stub.

To see how to make this your own, look here:

[README](https://openapi-generator.tech)

- API version: 0.0.10
- Build date: 2024-09-02T13:07:41.871312900+02:00[Europe/Berlin]
- Generator version: 7.8.0


### Running the server
To run the server, follow these simple steps:

```
go run main.go
```

The server will be available on `http://localhost:8080`.

To run the server in a docker container
```
docker build --network=host -t openapi .
```

Once image is built use
```
docker run --rm -it openapi
```
