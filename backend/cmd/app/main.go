package main

import (
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("GET /", ChooseConnect)
	http.HandleFunc("POST /connect", TryToConnect)

	http.ListenAndServe(":18080", nil)
}

func ChooseConnect(w http.ResponseWriter, r *http.Request) {
	app := "/bin/bash"
	cmd := exec.CommandContext(
		r.Context(),
		app,
		"-c",
		"nmcli --colors no -m multiline --get-value SSID dev wifi list",
	)

	out, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	res := `
  <!DOCTYPE html>
        <html>
        <head>
            <title>Wifi Control</title>
        </head>
        <body>
            <h1>Wifi Control</h1>
            <form action="/connect" method="post">
                <label for="ssid">Choose a WiFi network:</label>
                <select name="ssid" id="ssid">
  `

	ssids := strings.Replace(string(out), "SSID:", "", -1)
	ssidsList := strings.Split(ssids, "\n")

	for _, ssid := range ssidsList {
		if ssid != "" {
			res += fmt.Sprintf("<option value=\"%s\">%s</option>", ssid, ssid)
		}
	}

	res += `
                </select>
                <p/>
                <label for="password">Password: <input type="password" name="password"/></label>
                <p/>
                <input type="submit" value="Connect">
            </form>
        </body>
        </html>
  `

	fmt.Fprintf(w, res)
}

func TryToConnect(w http.ResponseWriter, r *http.Request) {
	ssid := r.FormValue("ssid")
	password := r.FormValue("password")

	app := "/bin/bash"
	cmd := exec.CommandContext(
		r.Context(),
		app,
		"-c",
		"nmcli dev wifi connect "+ssid,
	)

	if len(password) > 0 {
		cmd.Args[2] += " password " + password
	}

	_, err := cmd.Output()
	if err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			fmt.Println(string(exitErr.Stderr))
			fmt.Fprintf(w, "Error: %s", string(exitErr.Stderr))
			return
		}

		fmt.Println(err.Error())
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}

	res := `
        <!DOCTYPE html>
        <html>
        <head>
            <title>Wifi Control</title>
        </head>
        <body>
  `
	res += fmt.Sprintf("Successfully connected to %s", ssid)
	res += `
        </body>
        </html>
  `

	fmt.Fprintf(w, res)
}
