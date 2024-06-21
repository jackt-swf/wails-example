package main

import (
	"context"
	"net"
	"os"
	"os/user"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type ComputerInformation struct {
	Hostname  string `json:"hostname"`
	MACAdress string `json:"macAddress"`
	User      string `json:"user"`
}

func getMacAdresses() []string {
	ifas, _ := net.Interfaces()
	var macs []string
	for _, ifa := range ifas {
		mac := ifa.HardwareAddr.String()
		if mac != "" {
			macs = append(macs, strings.ToUpper(mac))
		}
	}
	return macs
}

// Greet returns a greeting for the given name
func (a *App) GetComputerInformation() (ComputerInformation, error) {
	hostname, _ := os.Hostname()
	mac := getMacAdresses()[0]
	user, _ := user.Current()

	return ComputerInformation{
		Hostname:  hostname,
		MACAdress: mac,
		User:      user.Username,
	}, nil
}
