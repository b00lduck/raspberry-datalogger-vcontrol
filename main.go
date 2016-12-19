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

    t1 := reading.NewTemperatureReading(vcd, "temperature_air", "getTempA", 0.2)

    for {

        err := t1.Process()
        if err != nil {
            log.Error(err)
        }

        time.Sleep(time.Second * 5)
    }

}
