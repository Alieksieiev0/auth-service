# Project auth-service

Auth-Service is a microservice, containing logic for authenticating of the user.
It provides functionality to Register, Login and Read Claims from given JWT token. 
It does not contain logic to save a user, but it uses a separate microservice for such purposes. 
Auth-Service uses HTTP API for Register and Login actions. 
For Read Claims action gRPC service is used.

## Getting Started

To run this service, just clone it, and start it 
using either [Make Run](#run) or [Make Docker Run](#run-in-docker). 
However, to run properly, it requires a separate microservice with
gRPC connection, to save and fetch users. In scope of the [Feed Project](https://github.com/Alieksieiev0/feed-templ)
microservice called [feed-service](https://github.com/Alieksieiev0/feed-service) was used.

## MakeFile

### Build
```bash
make build
```

### Run
```bash
make run
```

### Run in docker
```bash
make docker-run
```

### Run and rebuild in docker
```bash
make docker-build-n-run
```

### Shutdown docker
```bash
make docker-down
```

### Test
```bash
make test
```

### Clean
```bash
make clean
```

### Proto
```bash
make proto
```

### Live Reload
```bash
make watch
```

## Flags
This application supports startup flags, 
that can be passed to change servers and clients urls. 
However, be careful changing auth-service servers urls 
if you are running it using docker-compose, because by default
only ports 3001 and 4001 are exposed 

### REST server
- Long Name: rest-server
- Default: 3001

### gRPC server
- Name: grpc-server
- Default: 4001

### gRPC client
- Name: grpc-client
- Default: 4000
