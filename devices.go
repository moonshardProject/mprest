package main

import(
	"io"
	"net/http"
	"fmt"
	"encoding/json"
)

type Devices struct {
	jdevices string
}

type DevicesJson struct {
	NumberOfDevices int
	Devices []string
}

func newDevices(ip string) *Devices {
	devices := getAvailableDevices(ip)
	fmt.Println("Found devices...")
	for _, device := range devices {
		fmt.Println(device)
	}
	dj := DevicesJson{len(devices), devices}
	jdevices, _ := json.Marshal(dj)
	return &Devices{string(jdevices)}
}

func (devices *Devices) devicesHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, devices.jdevices)
}

