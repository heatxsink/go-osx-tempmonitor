go-osx-tempmonitor
==================

A golang package to access TemperatureMonitor.app temperature sensor data from your MacBook Pro/Air running OSX.

Setup
-----
- Download the dmg to TemperatureMonitor.app located over at http://www.bresink.com/osx/0TemperatureMonitor/download.php
- Install the downloaded dmg.
- go get github.com/heatxsink/go-osx-tempmonitor

Example
-------
```go
import(
	"github.com/heatxsink/go-osx-tempmonitor"
	"fmt"
)

func main() {
	data, err := tempmonitor.GetTemperatureSensors()
	if err != nil {
		fmt.Println(err)
	}
	for key, _ := range data {
		line := fmt.Sprintf("%s: %s", key, data[key])
		fmt.Println(line)
	}
}
```