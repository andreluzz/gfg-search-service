# Search Service Rest API

### Running the service with Docker
Access the service folder and execute the commands

1. Build the image
``` 
docker build -t andreluzz/go-gfg-search-service .
```
2. Run the image
``` 
docker run -it --network="host" andreluzz/go-gfg-search-service
```
> Using --network="host" assuming elasticsearch server is running on the host at port 9200

### Executing the service unit tests
Access the service folder and execute the commands

1. Run the image overriding main command
``` 
docker run -it andreluzz/go-gfg-search-service go test ./... -cover 
```

## Docker Compose
To execute the hole soluction with UI, Elasticsearch Server and Rest API service use docker-compose

``` 
docker-compose up 
```

Access http://localhost to use the UI

## Solution Architecture

![Solution Architecture](architecture_diagram.png)
