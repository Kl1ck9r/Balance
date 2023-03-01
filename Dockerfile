FROM golang:alpine 

WORKDIR /dir/balance/api

COPY . /dir/balance/api

COPY go.mod /dir/balance/api 
COPY go.mod /dir/balance/api

RUN go mod tidy 

RUN go build -o /main ./cmd/balance-api/

EXPOSE 8080

ENTRYPOINT ["/main"]