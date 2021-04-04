FROM golang:1.16.3-buster

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN apt-get -y update && apt-get -y install postgresql postgresql-client

EXPOSE 3000

CMD ["runproj"]