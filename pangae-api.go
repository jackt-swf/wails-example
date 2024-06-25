package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type StatusResult struct {
	Status string
}

func (app *App) CheckProcessingStatus() (string, error) {
	mac := strings.ReplaceAll(getMacAddress(), ":", "")
	url := app.config.TenantURL + "api/agent/status?mac=" + mac
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	status := &StatusResult{}
	if res.StatusCode == 200 {
		readJson(res.Body, status)
	} else {
		str := readString(res.Body)
		return "", errors.New(str)
	}

	fmt.Println(status)

	return status.Status, nil
}

type AuthenticateRequest struct {
	Email      string     `json:"email"`
	Option     string     `json:"option"`
	DeviceInfo DeviceInfo `json:"deviceInfo"`
	AgentId    string     `json:"agentId"`
}

type DeviceInfo struct {
	Nics []Nic `json:"nics"`
}

type Nic struct {
	MacAddress string `json:"macAddress"`
}

func (app *App) Authenticate() int {
	mac := getMacAddress()
	time := time.Now().String()

	h := sha1.New()
	h.Write([]byte(mac + time))

	agentId := hex.EncodeToString(h.Sum(nil))

	requestBody := AuthenticateRequest{
		Email:  app.config.UserEmail,
		Option: "biothenticate",
		DeviceInfo: DeviceInfo{
			Nics: []Nic{{mac}},
		},
		AgentId: agentId,
	}

	json, _ := json.Marshal(requestBody)

	url := app.config.TenantURL + "api/agent/v2/authenticate"
	res, _ := http.Post(url, "application/json", bytes.NewBuffer(json))

	fmt.Println(res)
	return res.StatusCode
}

func readString(body io.ReadCloser) string {
	defer body.Close()
	buf, _ := io.ReadAll(body)
	return string(buf)
}

func readJson(body io.ReadCloser, target interface{}) {
	defer body.Close()
	decoder := json.NewDecoder(body)
	decoder.Decode(target)
}
