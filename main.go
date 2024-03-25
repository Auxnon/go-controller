package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	fmt.Println("Gontroller Active")

	http.HandleFunc("/r", restart)
	http.HandleFunc("/s", shutdown)
	http.HandleFunc("/d", display_restart)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}

func restart(w http.ResponseWriter, r *http.Request) {
	exec_respond(w, r, "shutdown -r -t 3")
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	exec_respond(w, r, "shutdown -s -t 3")
}

func display_restart(w http.ResponseWriter, r *http.Request) {
	exec_respond(w, r, "systemctl restart display-manager")
}

//func getHello(w http.ResponseWriter, r *http.Request) {
//	fmt.Printf("got /hello request\n")
//	io.WriteString(w, "Hello, HTTP!\n")
//	// popup()
//}

func exec_respond(w http.ResponseWriter, r *http.Request, s string) {
	out := execute(s)
	fmt.Println(out)
	io.WriteString(w, out)
}

func wexecute(s string) {
	command := s
	err := exec.Command("powershell", "-NoProfile", command).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func execute(s string) string {
	out, err := exec.Command(s).Output()
	if err != nil {
		return "failed"
	}
	output := string(out)
	return output
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
