FROM golang:1.12.5-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/inven-store/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}inven-store
EXPOSE 8080