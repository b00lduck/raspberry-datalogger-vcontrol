package reading
import (
    "math"
    "github.com/b00lduck/raspberry-datalogger-dataservice-client"
    "github.com/b00lduck/raspberry-datalogger-vcontrol/vcontrold"
    log "github.com/Sirupsen/logrus"
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

    log.WithField("lpv", limitedPrecisionValue).
	WithField("reading", reading).
	WithField("precision", t.precision).
	WithField("oldValue", t.oldValue).
	Info("Set new reading")

    if math.Abs(float64(limitedPrecisionValue - t.oldValue)) > t.precision {
        if err := client.SendThermometerReading(t.code, limitedPrecisionValue); err != nil {
	    log.Error(err)
            return err
        } else {
	    log.Info("oldvalue")
            t.reading.oldValue = limitedPrecisionValue
        }
    }

    return nil
}

func round(x float64) float64 {
    return math.Floor(x + 0.5)
}