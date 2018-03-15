FROM golang:latest

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/ozankasikci/apollo-telegram-tracker/

COPY . $SRC_DIR

RUN cd $SRC_DIR; go get ./...

RUN cd $SRC_DIR; go build -o tracker ./cmd/tracker/tracker.go; cp tracker /app/;
RUN mkdir -p /app/firebase/
RUN cd $SRC_DIR; cp ./firebase/service-account.json /app/firebase/service-account.json;
RUN ["chmod", "+x", "/app/firebase/service-account.json"]
ENTRYPOINT ["./tracker"]