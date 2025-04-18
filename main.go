package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

const (
	PORT = "7070"
)

type Response struct {
	Message string `json:"message,omitempty"`
	Success bool   `json:"success,omitempty"`
}

func main() {
	fmt.Println("Gontroller Active on port " + PORT)

	http.HandleFunc("/r", restart)
	http.HandleFunc("/s", shutdown)
	http.HandleFunc("/d", display_restart)
	http.HandleFunc("/h", get_hostname)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}

func restart(w http.ResponseWriter, r *http.Request) {
	execute(w, "/bin/sh", "-c", "sleep 5; shutdown -r now")
	//exec_respond(w, r, "shutdown --help")
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	execute(w, "/bin/sh", "-c", `"sleep 5; shutdown now"`)
}

func display_restart(w http.ResponseWriter, r *http.Request) {
	execute(w, "systemctl", "restart", "display-manager")
}

func get_hostname(w http.ResponseWriter, r *http.Request) {
	execute(w, "hostname")
}

func wexecute(s string) {
	command := s
	err := exec.Command("powershell", "-NoProfile", command).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func execute(w http.ResponseWriter, s string, args ...string) {
	cmd := exec.Command(s, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	w.Header().Set("Content-Type", "application/json")
	response := Response{}
	response.Success = err == nil
	if !response.Success {
		w.WriteHeader(http.StatusInternalServerError)
		response.Message =
			strings.TrimSpace(fmt.Sprint(err) + ": " + stderr.String())
	} else {
		response.Message = strings.TrimSpace(out.String())
	}
	fmt.Println(response.Message)
	json.NewEncoder(w).Encode(response)

}

/* func popup() {
	command := `
	Add-Type -AssemblyName PresentationCore,PresentationFramework
	$ButtonType = [System.Windows.MessageBoxButton]::YesNoCancel
	$MessageIcon = [System.Windows.MessageBoxImage]::Error
	$MessageBody = "Are you sure you want to delete the log file?"
	$MessageTitle = "Confirm Deletion"
	$Result = [System.Windows.MessageBox]::Show($MessageBody,$MessageTitle,$ButtonType,$MessageIcon)
	Write-Host "Your choice is $Result"
	`
	err := exec.Command("powershell", "-NoProfile", command).Run()
	if err != nil {
		log.Fatal(err)
	}
}*/
