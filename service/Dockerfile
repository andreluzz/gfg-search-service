# docker build andreluzz/go-gfg-search-service -t .
# docker run -it --network="host" andreluzz/go-gfg-search-service
# OBS.: Assuming elasticsearch server is running on the host on port 9200

FROM golang:1-alpine

RUN apk add --no-cache --quiet gcc g++ ca-certificates make curl git jq

WORKDIR /go-gfg-search-service

COPY go.mod .
RUN go mod download

ENV TERM xterm-256color
ENV ES_HOST=http://127.0.0.1:9200

COPY . ./

CMD go run main.go
