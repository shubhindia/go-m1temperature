# Temperature Sensor Monitor for Apple Silicon M1

- `main.m`: modified on the Objective-C [code](https://github.com/freedomtan/sensors/blob/master/sensors/sensors.m) for iOS sensor by [freedomtan](https://github.com/freedomtan)

## Usage

run using `go run main.go` or build using `./darwin-build-arm64.sh` and then execute `./go-m1temperature`

## Output 

![screen shot](https://raw.githubusercontent.com/shubhindia/go-m1temperature/master/resources/output/go-m1temperature.png)

## References

For **better names** (e.g. what is `PMU TP3w` ?) for the sensors, please refer to

https://github.com/exelban/stats/blob/master/Modules/Sensors/values.swift

https://github.com/acidanthera/VirtualSMC/blob/master/Docs/SMCSensorKeys.txt

Here is a similar code in swift for getting sensor values using IOKit (for intel Mac)

https://github.com/exelban/stats/blob/master/Modules/Sensors/values.swift

For intel Mac, [gopsutil](https://github.com/shirou/gopsutil) can be uesed to get temperature information. 





