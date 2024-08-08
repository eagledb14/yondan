package main

import (
	"errors"
	"fmt"
	"net"
	"regexp"
	"strings"
)

// There are 3 queries, domain, port, and ips
// domain:url.com,example.com
// port:22,80,443
// ips:8.8.8.8,8.8.4.4/24 or 8.8.8.8
// multiple queries can be on the same line, separated by a space

var cidrRe *regexp.Regexp = regexp.MustCompile(`^([0-9]{1,3}\.){3}[0-9]{1,3}(\/([0-9]|[1-2][0-9]|3[0-2]))?`)
var netRe *regexp.Regexp = regexp.MustCompile(`^net:`)
var portRe *regexp.Regexp = regexp.MustCompile(`^port:`)
var domainRe *regexp.Regexp = regexp.MustCompile(`^domain:`)
var serviceRe *regexp.Regexp = regexp.MustCompile(`^service:`)

func query(params string, db *ConcurrentMap) ([]Scan, error) {

	queries := strings.Split(params, " ")

	if len(queries) == 0 {
		return nil, errors.New("No queries provided")
	}

	if cidrRe.MatchString(queries[0]) {
		return parseCidr(queries[0], db), nil
	}

	// if regexp.MatchString(p)

	// return "hi"
	return nil, nil
}

func filter([][]Scan) []Scan {

	return nil
}

func parseCidr(params string, db *ConcurrentMap) []Scan {

	if !strings.Contains(params, "/") {
		scan, err := db.Read(params)
		if err != nil {
			return nil
		}
		return scan
	}

	_, ipList, err := net.ParseCIDR(params)
	if err != nil {
		return nil
	}

	index := db.ReadAll()
	scanList := []Scan{}
	for key, v := range index {
		if ipList.Contains(net.ParseIP(key)) {
			fmt.Println(key)
			scanList = append(scanList, v...)
		}
	}

	return scanList
}

func parseNet(params []string) []Scan {
	// return fmt.Sprintf("%v", parems)
	return nil
}

func parsePort(params []string) []Scan {
	// return fmt.Sprintf("%v", params)
	return nil
}

func parseDomain(params []string) []Scan {
	// return fmt.Sprintf("%v", params)
	return nil
}
