package v1

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os/exec"
	"strings"
)

func New(done chan struct{}) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", ChooseConnect)
	mux.HandleFunc("/connect", TryToConnect(done))

	return mux
}

func ChooseConnect(w http.ResponseWriter, r *http.Request) {
	slog.Info("get connector request")

	ssidsList := getWifiList(r.Context())

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

	slog.Info("return html with wifi list")

	fmt.Fprintf(w, res)
}

func TryToConnect(done chan struct{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("post connector request")

		ssid := r.FormValue("ssid")
		password := r.FormValue("password")

		app := "/bin/bash"

		if ifWasConnected(r.Context(), ssid) {
			dropConnection := exec.CommandContext(
				r.Context(),
				app,
				"-c",
				"nmcli connection delete "+ssid,
			)

			_, err := dropConnection.Output()
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
		}

		cmd := exec.CommandContext(
			r.Context(),
			app,
			"-c",
			"sudo nmcli dev wifi connect "+ssid,
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

		slog.Info("successfully connected to wifi")

		fmt.Fprintf(w, res)

		done <- struct{}{}
		close(done)
	}
}

func ifWasConnected(ctx context.Context, ssid string) bool {
	app := "/bin/bash"
	cmd := exec.CommandContext(ctx, app, "-c", "nmcli connection | grep "+ssid)
	out, err := cmd.Output()
	if err != nil {
		return false
	}

	if len(out) <= 0 {
		return false
	}

	return true
}

func getWifiList(ctx context.Context) []string {
	app := "/bin/bash"
	cmd := exec.CommandContext(
		ctx,
		app,
		"-c",
		"nmcli --colors no -m multiline --get-value SSID dev wifi list",
	)

	out, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	ssids := strings.Replace(string(out), "SSID:", "", -1)
	ssidsList := strings.Split(ssids, "\n")

	return ssidsList
}
