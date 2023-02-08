
// Self-contained webserver running Snort or another Nostr web client

package main

import (
	"net/http"
	"os/exec"
	"runtime"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	go http.ListenAndServe(":8080", nil)

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}

	args = append(args, "http://localhost:8080")

	err := exec.Command(cmd, args...).Start()
	if err != nil {
		panic(err)
	}
}