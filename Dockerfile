FROM rem/rpi-golang-1.7:latest

WORKDIR /gopath/src/github.com/b00lduck/raspberry-datalogger-vcontrol
ENTRYPOINT ["raspberry-datalogger-vcontrol"]

ADD . /gopath/src/github.com/b00lduck/raspberry-datalogger-vcontrol
RUN go get
RUN go build
USER root
