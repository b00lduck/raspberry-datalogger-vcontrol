package reading
import (
    "github.com/b00lduck/raspberry-datalogger-vcontrol/vcontrold"
    log "github.com/Sirupsen/logrus"
    "github.com/b00lduck/raspberry-datalogger-dataservice-client"
)

type flag struct {
    reading
    oldValue bool
}

func NewFlagReading(code string, command string) Reading {
    return &flag{
        reading: reading{
            code: code,
            command: command},
        oldValue: false}
}

func (t *flag) Process() error {

    vc, err := vcontrold.NewVcontroldClient()
    if err != nil {
        return err
    }
    defer vc.Close()

    err = vc.ReadPrompt()
    if err != nil {
        return err
    }

    state, err := vc.GetFlag(t.command)
    if err != nil {
        return err
    }

    return t.setNewReading(state)
}

func (t *flag) setNewReading(state bool) error {

    log.WithField("code", t.code).
        WithField("state", state).
        Debug("Set new reading")

    if state != t.oldValue {
        log.WithField("code", t.code).
            WithField("oldValue", t.oldValue).
            WithField("state", state).
            Info("Value has changed, updating now")

        if err := client.SendFlagState(t.code, state); err != nil {
    	    log.WithField("err", err).Error("Error sending thermometer reading")
            return err
        } else {
            t.oldValue = state
        }

    } else {
        log.WithField("code", t.code).
            WithField("oldValue", t.oldValue).
            WithField("state", state).
            Info("Value has not changed, no update necessary")
    }

    return nil
}