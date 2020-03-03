FROM golang:1.14-alpine AS builder

ENV GO111MODULES=on
WORKDIR /src
COPY . .
RUN apk update && apk add git \
      && go build .

FROM alpine:3

WORKDIR /app
COPY --from=builder /src/just-redirect /app/just-redirect

RUN chmod 0755 /app/just-redirect

CMD /app/just-redirect
