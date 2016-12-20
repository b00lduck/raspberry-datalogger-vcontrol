package reading
import (
    "math"
    "github.com/b00lduck/raspberry-datalogger-vcontrol/vcontrold"
    log "github.com/Sirupsen/logrus"
    "github.com/b00lduck/raspberry-datalogger-dataservice-client"
)

type temperature struct {
    reading
    precision float64
}

func NewTemperatureReading(vcontrold vcontrold.Vcontrold, code string, command string, precision float64) Reading {
    return &temperature{
        reading: reading{
            vcontrold: vcontrold,
            code: code,
            command: command,
            oldValue: 0},
        precision: precision,
    }
}

func (t *temperature) Process() error {

    err := t.vcontrold.ReadPrompt()
    if err != nil {
        return err
    }

    temp, err := t.vcontrold.GetTemp(t.command)
    if err != nil {
        return err
    }

    return t.setNewReading(temp)
}

func (t *temperature) setNewReading(reading float64) error {

    // precision reduction
    limitedPrecisionValue := round(reading / t.precision) * t.precision

    log.WithField("code", t.code).
        WithField("reading", reading).
	    Info("Set new reading")

    if math.Abs(float64(limitedPrecisionValue - t.oldValue)) >= t.precision {
        log.WithField("code", t.code).
            WithField("oldValue", t.oldValue).
            WithField("limitedPrecisionValue", limitedPrecisionValue).
            Info("Value has changed equal or more than the precision limit, updating now")

        if err := client.SendThermometerReading(t.code, limitedPrecisionValue); err != nil {
    	    log.WithField("err", err).Error("Error sending thermometer reading")
            return err
        } else {
            t.reading.oldValue = limitedPrecisionValue
        }

    } else {
        log.WithField("code", t.code).
            WithField("oldValue", t.oldValue).
            WithField("limitedPrecisionValue", limitedPrecisionValue).
            Info("Value has not changed equal or more than the precision limit, no update necessary")
    }

    return nil
}

func round(x float64) float64 {
    return math.Floor(x + 0.5)
}