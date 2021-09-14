package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	u2f "github.com/marshallbrekka/go-u2fhost"
	log "github.com/sirupsen/logrus"
)

func main() {
	var authR *u2f.AuthenticateRequest

	err := json.NewDecoder(os.Stdin).Decode(&authR)
	if err != nil {
		log.Fatal(err)
	}

	response := authenticateHelper(authR, u2f.Devices())
	responseJson, _ := json.Marshal(response)
	fmt.Println(string(responseJson))

}

func authenticateHelper(req *u2f.AuthenticateRequest, devices []*u2f.HidDevice) *u2f.AuthenticateResponse {
	log.Debugf("Authenticating with request %+v", req)
	openDevices := []u2f.Device{}
	for i, device := range devices {
		err := device.Open()
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
			fmt.Println("Failed to get authentication response after 25 seconds")
			return nil
		case <-interval.C:
			for _, device := range openDevices {
				response, err := device.Authenticate(req)
				if err == nil {
					return response
				} else if _, ok := err.(u2f.TestOfUserPresenceRequiredError); ok && !prompted {
					fmt.Println("\nTouch the flashing U2F device to authenticate...")
					prompted = true
				} else {
					log.Debugf("Got status response %s", err)
				}
			}
		}
	}
}
