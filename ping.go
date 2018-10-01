package main

import (
	"fmt"
	"os/exec"
	"bytes"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
)

func execute(program string, args ...string) string {
	cmd := exec.Command(program, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		// for strange reasons, err was 1 but the stderr.String() was empty. So i commented the following line:
		// fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}

	return out.String()
}

func getAvailableDevices(ip string) []string {
	fmt.Println("Pinging hosts...")
	hosts := execute("fping","-a", "-q", "-g", ip)
	splittedHosts := strings.Split(hosts, "\n")
	splittedHosts = splittedHosts[:len(splittedHosts)-1]
	return iotDevices(splittedHosts)
}


func iotDevices(hosts []string) []string {
	var iotHosts []string
	pwd, _ := os.Getwd()
	dat, _ := ioutil.ReadFile(pwd + "/assets/pages/relayhtml1.txt")
	dat2, _ := ioutil.ReadFile(pwd + "/assets/pages/relayhtml2.txt")

	for _, host := range hosts {
		response, err := http.Get("http://" + host)
		if err != nil {
			
		} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
			}
			if(string(contents) == string(dat) || string(contents) == string(dat2)) {
				iotHosts = append(iotHosts, host)
			}
		}
	}
	return iotHosts
}
