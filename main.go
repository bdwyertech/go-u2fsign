package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	u2f "github.com/marshallbrekka/go-u2fhost"
	log "github.com/sirupsen/logrus"
)

func init() {
	if _, verbose := os.LookupEnv("U2F_VERBOSE"); verbose {
		log.SetLevel(log.DebugLevel)
	}
	if _, trace := os.LookupEnv("U2F_TRACE"); trace {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	}
}

func main() {
	var authR u2f.AuthenticateRequest

	err := json.NewDecoder(os.Stdin).Decode(&authR)
	if err != nil {
		log.Fatal(err)
	}

	response, err := u2fAuth(&authR, u2f.Devices())
	if err != nil {
		log.Fatal(err)
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseJson))

}

func u2fAuth(req *u2f.AuthenticateRequest, devices []*u2f.HidDevice) (response *u2f.AuthenticateResponse, err error) {
	log.Debugf("Authenticating with request %+v", req)
	openDevices := []u2f.Device{}
	for i, device := range devices {
		err = device.Open()
		if err == nil {
			openDevices = append(openDevices, u2f.Device(devices[i]))
			defer func(i int) {
				devices[i].Close()
			}(i)
			version, err := device.Version()
			if err != nil {
				log.Debugf("Device version error: %s", err.Error())
			} else {
				log.Debugf("Device version: %s", version)
			}
		}
	}
	if len(openDevices) == 0 {
		log.Fatalf("Failed to find any devices")
	}
	prompted := false
	timeout := time.After(time.Second * 25)
	interval := time.NewTicker(time.Millisecond * 250)
	defer interval.Stop()
	for {
		select {
		case <-timeout:
			err = errors.New("failed to get authentication response after 25 seconds")
			return
		case <-interval.C:
			for _, device := range openDevices {
				response, err = device.Authenticate(req)
				if err == nil {
					return
				} else if err.Error() == "Device is requesting test of use presence to fulfill the request." && !prompted {
					log.Infoln("Touch the flashing U2F device to authenticate...")
					prompted = true
				} else {
					log.Debugf("Got status response %s", err)
				}
			}
		}
	}
}
