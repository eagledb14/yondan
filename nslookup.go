package main

import (
	"os/exec"
	"strings"
	"errors"
)


func Lookup(target string) (string, error) {
	output, err := exec.Command("nslookup", target).Output()
	if err != nil {
		return "", err
	}

	outSplit := strings.Split(string(output), "name = ")
	if len(outSplit) < 2 {
		return "", errors.New("Host Not Found")
	}

	hostname := strings.Split(outSplit[1], "\n")
	if len(outSplit) == 0 {
		return "", errors.New("Host Not Found")
	}
	return hostname[0], nil
}
