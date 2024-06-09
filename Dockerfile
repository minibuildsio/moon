FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/
COPY . .

RUN go build -o /bin/moon

FROM scratch

WORKDIR /app
COPY --from=builder /bin/moon /bin/moon
CMD ["/bin/moon"]