package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed resources/bin/m1templib
var f []byte

func main() {
	_ = os.WriteFile("/tmp/m1templib", f, 0755)
	sensors, _ := exec.Command("/tmp/m1templib", "n").Output()
	temperatures, _ := exec.Command("/tmp/m1templib", "v").Output()

	sensorslice := strings.Split(string(sensors), ",")
	templice := strings.Split(string(temperatures), ",")

	for i := 0; i < (len(sensorslice) - 1); i++ {
		fmt.Printf("%s : %s\n", sensorslice[i], templice[i])
	}

}
