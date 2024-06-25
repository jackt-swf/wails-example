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
	ctx    context.Context
	config Config
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (app *App) startup(ctx context.Context) {
	app.ctx = ctx
}

// Config

type Config struct {
	TenantURL string `json:"tenantUrl"`
	UserEmail string `json:"userEmail"`
}

func (app *App) GetConfig() Config {
	return app.config
}

func (app *App) SaveConfig(config Config) {
	app.config = config
}

// ComputerInformation

type ComputerInformation struct {
	Hostname  string `json:"hostname"`
	MACAdress string `json:"macAddress"`
	User      string `json:"user"`
}

func getMacAddress() string {
	ifas, _ := net.Interfaces()
	for _, ifa := range ifas {
		mac := ifa.HardwareAddr.String()
		if hasFlag(ifa.Flags, net.FlagUp) && mac != "" {
			return strings.ToUpper(mac)
		}
	}
	return ""
}

func hasFlag(flags, flag net.Flags) bool {
	return (flags & flag) == flag
}

// Greet returns a greeting for the given name
func (a *App) GetComputerInformation() (ComputerInformation, error) {
	hostname, _ := os.Hostname()
	mac := getMacAddress()
	user, _ := user.Current()

	return ComputerInformation{
		Hostname:  hostname,
		MACAdress: mac,
		User:      user.Username,
	}, nil
}
