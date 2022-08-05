package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
)

//go:embed resources/bin/m1templib
var f []byte

func main() {
	_ = os.WriteFile("/tmp/m1templib", f, 0755)
	sensors, _ := exec.Command("/tmp/m1templib", "n").Output()
	temperatures, _ := exec.Command("/tmp/m1templib", "v").Output()

	fmt.Printf("Sensors: %s\n", sensors)
	fmt.Printf("Temperatures: %s\n", temperatures)

}
