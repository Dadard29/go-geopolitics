FROM golang:1.14-alpine

ARG ARG_VERSION
ARG ARG_HOST_SUB
ARG ARG_DB_PASSWORD

ENV VERSION=$ARG_VERSION
ENV HOST_SUB=$ARG_HOST_SUB
ENV DB_PASSWORD=$ARG_DB_PASSWORD

ENV CORS_ORIGIN=https://dadard.fr

RUN apk add --update git gcc libc-dev

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]