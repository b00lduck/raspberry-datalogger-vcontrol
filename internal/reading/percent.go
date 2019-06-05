package reading
import (
    "github.com/b00lduck/raspberry-datalogger-vcontrol/internal/vcontrold"
    log "github.com/sirupsen/logrus"
    "github.com/b00lduck/raspberry-datalogger-dataservice-client"
)

type percent struct {
    reading
    oldValue float64
}

func NewPercentReading(code string, command string) Reading {
    return &percent{
        reading: reading{
            code: code,
            command: command},
        oldValue: 0}
}

func (t *percent) Process() error {

    vc, err := vcontrold.NewVcontroldClient()
    if err != nil {
        return err
    }
    defer vc.Close()

    err = vc.ReadPrompt()
    if err != nil {
        return err
    }

    percent, err := vc.GetPercent(t.command)
    if err != nil {
        return err
    }

    return t.setNewReading(percent)
}

func (t *percent) setNewReading(percent float64) error {

    log.WithField("code", t.code).
        WithField("percent", percent).
        Debug("Set new reading")

    if percent != t.oldValue {
        log.WithField("code", t.code).
            WithField("oldValue", t.oldValue).
            WithField("percent", percent).
            Info("Value has changed, updating now")

        if err := client.SendPercentage(t.code, percent); err != nil {
            log.WithField("err", err).Error("Error sending percentage reading")
            return err
        } else {
            t.oldValue = percent
        }

    } else {
        log.WithField("code", t.code).
            WithField("oldValue", t.oldValue).
            WithField("percent", percent).
            Info("Value has not changed, no update necessary")
    }

    return nil
}

