package tempmonitor

import (
	"os/exec"
	"strings"
	"bytes"
)

type TemperatureSensors map[string] string

func GetTemperatureSensors() (TemperatureSensors, error) {
	command := "/Applications/TemperatureMonitor.app/Contents/MacOS/tempmonitor"
	cmd := exec.Command(command, "-c", "-l", "-a")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	temp_data := TemperatureSensors {}
	lines := strings.Split(out.String(), "\n")
	for _, line := range lines {
		datum := strings.Split(line, ": ")
		if len(datum) > 1 {
			key := datum[0]
			value := strings.TrimRight(datum[1], " C")
			temp_data[key] = value
		}
	}
	return temp_data, nil
}