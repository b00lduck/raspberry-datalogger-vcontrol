package reading
import (
    "math"
    "github.com/b00lduck/raspberry-datalogger-dataservice-client"
    log "github.com/Sirupsen/logrus"
    "github.com/b00lduck/raspberry-datalogger-vcontrol/vcontrold"
)

type temperature struct {
    reading
    precision float64
}

func NewTemperatureReading(vcontrold vcontrold.Vcontrold, code string, command string, precision float64) Reading {
    return temperature{
        reading: reading{
            vcontrold: vcontrold,
            code: code,
            command: command,
            oldValue: 0},
        precision: precision,
    }
}

func (t temperature) Process() error {

    err := t.vcontrold.ReadPrompt()
    if err != nil {
        return err
    }

    temp, err := t.vcontrold.GetTemp(t.command)
    if err != nil {
        return err
    }

    log.Info(temp)
    return t.setNewReading(temp)
}

func (t temperature) setNewReading(reading float64) error {

    // precision reduction
    limitedPrecisionValue := round(reading / t.precision) * t.precision

    if math.Abs(float64(limitedPrecisionValue - t.oldValue)) > t.precision {
        if err := client.SendThermometerReading(t.code, limitedPrecisionValue); err != nil {
            return err
        } else {
            t.oldValue = limitedPrecisionValue
        }
    }
    return nil
}

func round(x float64) float64 {
    return math.Floor(x + 0.5)
}