##
## Build
##

FROM golang:alpine AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /main

##
## Deploy
##

FROM alpine

WORKDIR /

COPY --from=build /main /main
CMD ["/main"]
