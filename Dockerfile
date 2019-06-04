FROM balenalib/raspberry-pi-golang AS builder
COPY . /src
WORKDIR /src
RUN CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o /app .

FROM scratch
COPY --from=builder /app /
ENTRYPOINT [ "/app" ]

