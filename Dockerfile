FROM golang:1.12.5-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git
# RUN apk --no-cache and make ca-certificates
#vRUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/inven-store/
COPY . /go/src/inven-store/

RUN cd /go/src/inven-store/ && CGO_ENABLED=0 go build

WORKDIR /go/src/inven-store/
CMD /go/src/inven-store/inven-store
EXPOSE 8080