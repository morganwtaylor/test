package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ReadHostnames reads hostnames from a file.
func ReadHostnames(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}

// SSHAndExecuteCommandWindows executes a command on a remote host via SSH on WINDOWS.
func SSHAndExecuteCommandWindows(username, password, hostname, command string) []string {
	cmd := exec.Command("plink.exe", "-ssh", "-l", username, "-pw", password, hostname, command)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error executing command: ", err)
		return nil
	}
	output := string(outputBytes)
	// Split the output string by the newline character and create a slice
	outputLines := strings.Split(output, "\n")
	return outputLines
}

// SSHAndExecuteCommandLinux executes a command on a remote host via SSH on LINUX.
func SSHAndExecuteCommandLinux(username, password, hostname, command string) []string {
	sshCommand := fmt.Sprintf(`expect -c '
	spawn ssh %s@%s %s
	expect "password:"
	send "%s\r"
	expect eof
	'`, username, hostname, command, password)
	cmd := exec.Command("bash", "-c", sshCommand)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("Error executing command: ", err)
		return nil
	}
	output := string(outputBytes)
	// Split the output string by the newline character and create a slice
	outputLines := strings.Split(output, "\n")
	return outputLines
}

// Swift Control 2.2 - Security Updates. Required to SSH and view os release to verify it is up to date.
func Ccontrol2_2SecurityUpdates() {
	username := "root"
	password := os.Getenv("SSH_PASSWORD")
	hostnames, err := ReadHostnames("hostnames.txt")
	if err != nil {
		log.Fatal(err)
	}

	results := make(map[string][]string)
	for _, hostname := range hostnames {
		if hostname != "" {
			command := "more /etc/os-release" // command to be executed
			var outputLines []string
			if runtime.GOOS == "windows" {
				outputLines = SSHAndExecuteCommandWindows(username, password, hostname, command)
			} else {
				outputLines = SSHAndExecuteCommandLinux(username, password, hostname, command)
			}
			results[hostname] = outputLines
		}
	}

	jsonData, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("2.2_output.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	Control2_2SecurityUpdates()
}
