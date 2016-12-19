package vcontrold

import (
    "os"
    "net"
    "bufio"
    "fmt"
    log "github.com/Sirupsen/logrus"
    "strconv"
    "strings"
)

type Vcontrold interface {
    Close()
    ReadPrompt() error
    GetTemp(cmd string) (float64, error)
}

type vcontrold struct {
    connection net.Conn
}

func NewVcontroldClient() Vcontrold {

    host := os.Getenv("VCONTROLD_HOST")

    connection, err := net.Dial("tcp", host)
    if err != nil {
        log.Fatal(err)
    }

    obj := vcontrold {
        connection: connection,
    }

    return obj
}

func (v vcontrold) Close() {
    v.connection.Close()
}

func (v vcontrold) ReadPrompt() error {
    message, err := bufio.NewReader(v.connection).ReadString('>')
    if err != nil {
        return err
    }

    if message != "vctrld>" {
        log.WithField("message", message).Fatal("Received wrong prompt")
        return fmt.Errorf("Received wrong prompt")
    }

    return nil
}

func (v vcontrold) GetTemp(cmd string) (float64, error) {
    fmt.Fprintf(v.connection, "%s\n", cmd)
    message, err := bufio.NewReader(v.connection).ReadString('\n')
    if err != nil {
        return 0, err
    }
    splitted := strings.Split(message, " ")
    return strconv.ParseFloat(splitted[0], 64)
}