FROM golang:1.21

ARG PORT

WORKDIR /api

COPY . /api
RUN go get 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

EXPOSE ${PORT} 

ENTRYPOINT ["/api/server"]
