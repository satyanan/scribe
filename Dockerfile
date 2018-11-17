# build stage
FROM golang:1.11.1 AS build-env
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/scribe/main.go

# final stage
FROM scratch
COPY --from=build-env /app/main /main
ENTRYPOINT ["/main"]
