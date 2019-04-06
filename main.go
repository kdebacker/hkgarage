package main

import (
	"log"

	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
)

func main() {
	info := accessory.Info{
		Name: "Garage Door",
	}
	acc := NewGarageDoorOpener(info)

	cfg := hc.Config{
		Pin: "00102003",
		Port: "12345",
		StoragePath: "./db",
	}

	t, err := hc.NewIPTransport(cfg, acc.Accessory)
	if err != nil {
		log.Fatalf("Failed to initialize HomeControl: %s", err)
	}

	hc.OnTermination(func() {
		<-t.Stop()
	})
	t.Start()
}