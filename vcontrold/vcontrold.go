package vcontrold

import (
    "net"
    "bufio"
    "fmt"
    log "github.com/Sirupsen/logrus"
    "strconv"
    "strings"
    "os"
)

type Vcontrold interface {
    Close()
    ReadPrompt() error
    GetTemperature(cmd string) (float64, error)
    GetFlag(cmd string) (bool, error)
}

type vcontrold struct {
    connection net.Conn
}

func NewVcontroldClient() (Vcontrold, error) {

    host := os.Getenv("VCONTROLD_HOST")

    connection, err := net.Dial("tcp", host)
    if err != nil {
        log.WithField("err", err).
            WithField("host", host).
            Error("Error connection to vcontrold")
        return nil, err
    }

    obj := vcontrold {
        connection: connection,
    }

    return obj, nil
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
        log.WithField("message", message).Error("Received wrong prompt from vcontrold")
        return fmt.Errorf("Received wrong prompt from vcontrold")
    }

    return nil
}

func (v vcontrold) GetTemperature(cmd string) (float64, error) {
    fmt.Fprintf(v.connection, "%s\n", cmd)
    message, err := bufio.NewReader(v.connection).ReadString('\n')
    if err != nil {
        return 0, err
    }
    splitted := strings.Split(message, " ")
    return strconv.ParseFloat(splitted[0], 64)
}

func (v vcontrold) GetFlag(cmd string) (bool, error) {
    fmt.Fprintf(v.connection, "%s\n", cmd)
    message, err := bufio.NewReader(v.connection).ReadString('\n')
    if err != nil {
        return false, err
    }
    splitted := strings.Split(message, "\n")
    intVal, err := strconv.ParseInt(splitted[0], 10, 8)
    if err != nil {
        return false, err
    }

    return intVal == 1, nil
}