FROM balenalib/raspberry-pi-golang AS builder
COPY . /src
WORKDIR /src
RUN go build -o /app .

FROM scratch
COPY --from=builder /app /
ENTRYPOINT [ "/app" ]

