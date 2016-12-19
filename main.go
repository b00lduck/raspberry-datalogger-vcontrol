package main

import (
    "time"
    log "github.com/Sirupsen/logrus"
    "github.com/b00lduck/raspberry-datalogger-vcontrol/vcontrold"
    "github.com/b00lduck/raspberry-datalogger-vcontrol/reading"
)

func main() {

    vcd := vcontrold.NewVcontroldClient()
    defer vcd.Close()

    readings := make([]reading.Reading, 0)

    readings = append(readings, reading.NewTemperatureReading(vcd, "temperature_air", "getTempA", 0.2))

    for {
        for _, v := range readings {
            err := v.Process()
            if err != nil {
                log.Error(err)
            }
        }
        time.Sleep(time.Second * 5)
    }

}
