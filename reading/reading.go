package reading

import "github.com/b00lduck/raspberry-datalogger-vcontrol/vcontrold"

type Reading interface {
    Process() error
}

type reading struct {
    vcontrold vcontrold.Vcontrold

    oldValue float64
    code string
    command string
}