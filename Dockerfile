# builder image
FROM golang:1.20.1-alpine as gobuilder

RUN apk update && apk add curl protobuf make \
    rm -rf /var/cache/apk/*

RUN mkdir /app
COPY . /app/
WORKDIR /app

# Download Go modules
RUN go mod download

# Build
RUN make build


# executable image
FROM scratch

COPY --from=gobuilder /app/build/linux_amd64/taraxa-indexer /taraxa-indexer

EXPOSE 8080
ENV HTTP_PORT=8080

# Run
CMD [ "/taraxa-indexer" ]