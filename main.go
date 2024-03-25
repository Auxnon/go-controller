package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

const (
	PORT = "7070"
)

func main() {
	fmt.Println("Gontroller Active on port " + PORT)

	http.HandleFunc("/r", restart)
	http.HandleFunc("/s", shutdown)
	http.HandleFunc("/d", display_restart)
	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}

func restart(w http.ResponseWriter, r *http.Request) {
	exec_respond(w, r, "/bin/sh", "-c", "sleep 5; shutdown -r now")
	//exec_respond(w, r, "shutdown --help")
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	exec_respond(w, r, "/bin/sh", "-c", `"sleep 5; shutdown now"`)
}

func display_restart(w http.ResponseWriter, r *http.Request) {
	exec_respond(w, r, "systemctl", "restart", "display-manager")
}

//func getHello(w http.ResponseWriter, r *http.Request) {
//	fmt.Printf("got /hello request\n")
//	io.WriteString(w, "Hello, HTTP!\n")
//	// popup()
//}

func exec_respond(w http.ResponseWriter, r *http.Request, s string, args ...string) {
	out := execute(s, args...)
	fmt.Println(out)
	io.WriteString(w, "success: "+out)
}

func wexecute(s string) {
	command := s
	err := exec.Command("powershell", "-NoProfile", command).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func execute(s string, args ...string) string {
	cmd := exec.Command(s, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return "failed"
	}

	fmt.Println("Exec Result: " + out.String() + " (" + stderr.String() + ")")

	return out.String()

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
