package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	fmt.Println("Hello, world.")

	http.HandleFunc("/hello", getHello)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
	popup()
}

func popup() {
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
}
