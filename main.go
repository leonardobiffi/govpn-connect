package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/alecthomas/kingpin/v2"
)

var (
	config = kingpin.Flag("config", "OpenVPN Config File").Required().String()
)

func main() {
	kingpin.Parse()

	if *config == "" {
		kingpin.Usage()
		return
	}

	// Load username and base password from environment variables
	username := os.Getenv("OPENVPN_USERNAME")
	basePassword := os.Getenv("OPENVPN_PASSWORD")
	if username == "" || basePassword == "" {
		fmt.Println("Error: OPENVPN_USERNAME or OPENVPN_PASSWORD is not set")
		os.Exit(1)
	}

	// Prompt user for OTP
	fmt.Print("Enter OTP Code: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	otpCode := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading OTP:", err)
		os.Exit(1)
	}

	// Combine base password with OTP
	finalPassword := basePassword + otpCode

	// Run openvpn3 command
	cmd := exec.Command("openvpn3", "session-start", "--config", *config)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error creating stdin pipe:", err)
		os.Exit(1)
	}

	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, username)
		fmt.Fprintln(stdin, finalPassword)
	}()

	cmd.Stdout = nil
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running OpenVPN:", err)
		os.Exit(1)
	}

	fmt.Println("Connected!")
}
