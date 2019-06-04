package main

import (
    "time"
    log "github.com/sirupsen/logrus"
    "github.com/b00lduck/raspberry-datalogger-vcontrol/reading"
)

func main() {

    readings := make([]reading.Reading, 14)

    readings[0] = reading.NewTemperatureReading("temperature_air", "getTempA", 0.1, -40, 60)
    readings[1] = reading.NewTemperatureReading("temperature_hot_water_tank", "getTempWW", 0.2, 10, 100)
    readings[2] = reading.NewTemperatureReading("temperature_boiler", "getTempKist", 0.2, 10, 100)
    readings[3] = reading.NewTemperatureReading("temperature_exhaust", "getTempAbg", 0.2, 10, 100)
    readings[4] = reading.NewTemperatureReading("temperature_forward_flow", "getTempVorlaufHk1", 0.2, 10, 100)
    readings[5] = reading.NewTemperatureReading("temperature_return_flow", "getTempRuecklaufHk1", 0.2, 10, 100)
    readings[6] = reading.NewTemperatureReading("temperature_boiler_target", "getTempKsoll", 0.2, 10, 100)


    readings[7] = reading.NewFlagReading("flag_heating_circulation_pump_active", "getPumpeStatusHk1")
    readings[8] = reading.NewFlagReading("flag_hot_water_loading_pump_active", "getPumpeStatusSp")
    readings[9] = reading.NewFlagReading("flag_boiler_internal_pump", "getPumpeStatusBr")

    readings[10] = reading.NewPercentReading("percent_boiler_throttle", "getDrosselklappenPosition")
    readings[11] = reading.NewPercentReading("percent_boiler_mixer", "getMischerPosition")
    readings[12] = reading.NewPercentReading("percent_boiler_power", "getBrennerLeistung")
    readings[13] = reading.NewPercentReading("percent_boiler_throttle", "getDrosselklappenPosition")

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
