## Build the docker image
```bash
docker build . --tag=us-states-webservice
```

## Run the docker image
```bash
docker run -dp 8001:8001 us-states-webservice
```

### Ez :)