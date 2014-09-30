package tempmonitor

import (
	"testing"
	"fmt"
)

func TestGetTemperatureSensors(t *testing.T) {
	data, err := GetTemperatureSensors()
	if err != nil {
		t.Fail()
	} else {
		for key, _ := range data {
			line := fmt.Sprintf("%s: %s", key, data[key])
			fmt.Println(line)
		}
	}
}