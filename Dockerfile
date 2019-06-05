FROM golang:alpine AS builder
RUN apk add --no-cache git gcc musl-dev
ADD . /src
WORKDIR /src
RUN go get ./... \
 && go vet ./... \
 && go test ./...\
 && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app

FROM scratch
USER 10000:10000
COPY --from=builder /app /app
ENTRYPOINT ["/app"]

