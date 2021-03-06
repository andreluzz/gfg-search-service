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

## API Requests

Available query string parameters:
- __q__: query products.
- __filter__: filter products. Should be used passing the field_name:value.
- __sort__: define products order. Should be used passing filed_name:order_type. Available order types are: __desc__ and __asc__.
- __page__: defines the page to select the results. Should be an integer.
- __limit__: defines the amount of items per page. Should be an integer.

Example:
``` 
http://localhost:8080/products?q=shoes&filter=brand:adidas&sort=title:desc&page=1&limit=20
```

### Version and Authentication
Every request should have the headers to define API Version and Authentication.

For defining the API version use the header: "X-Service-Version" in the request. For test purpose there are two versions available: "v1" and "v2".
``` 
"X-Service-Version": "v2"
```

Authentication use the header: "Authorization" passing a valid token. For test use the token: 
``` 
"Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.wDWyyGem9YgXDDbH3Un7YYcTB8IcN_BE4BMmS1tvlnE"
```

In this implementation was used the JSON Web Token (JWT) to generate the token. If is necessary create other tokens use https://jwt.io/ with the key "top-secret-signin-value-key".

## Solution Architecture

![Solution Architecture](architecture_diagram.png)
