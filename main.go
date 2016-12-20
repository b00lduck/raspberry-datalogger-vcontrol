package main

import (
    "time"
    log "github.com/Sirupsen/logrus"
    "github.com/b00lduck/raspberry-datalogger-vcontrol/reading"
)

func main() {

    readings := make([]reading.Reading, 4)

    readings[0] = reading.NewTemperatureReading("temperature_air", "getTempA", 0.2, -40, 60)
    readings[1] = reading.NewTemperatureReading("temperature_hot_water_tank", "getTempWW", 0.2, 10, 100)
    readings[2] = reading.NewTemperatureReading("temperature_boiler", "getTempKist", 0.2, 10, 100)
    readings[3] = reading.NewTemperatureReading("temperature_exhaust", "getTempAbg", 0.2, 10, 100)

    for {
        for _, v := range readings {
            err := v.Process()
            if err != nil {
                log.Error(err)
            }
            time.Sleep(time.Second * 1)
        }
        time.Sleep(time.Second * 15)
    }

}
