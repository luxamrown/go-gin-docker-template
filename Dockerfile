# Builder
FROM golang-alpine3.11 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /

COPY . .

RUN make engine

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 8080

COPY --from=builder /app/engine /app

CMD /