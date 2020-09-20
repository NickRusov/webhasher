FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build webhasher.go 
ENTRYPOINT ["/app/webhasher"]
EXPOSE 8080
