package main

import (
	"strings"
	"fmt"
	"regexp"
	"errors"
)

// There are 3 queries, domain, port, and ips
// domain:url.com,example.com
// port:22,80,443
// ips:8.8.8.8,8.8.4.4/24 or 8.8.8.8
// multiple queries can be on the same line, separated by a space

var cidrRe *regexp.Regexp = regexp.MustCompile(`^([1-9]{1,3}\.){3}[0-9]{1,3}(\/([0-9]|[1-2][0-9]|3[0-2]))?`)
var netRe *regexp.Regexp = regexp.MustCompile(`^net:`)
var portRe *regexp.Regexp = regexp.MustCompile(`^port:`)
var domainRe *regexp.Regexp = regexp.MustCompile(`^domain:`)
var serviceRe *regexp.Regexp = regexp.MustCompile(`^service:`)

func query(params string, db *ConcurrentMap) ([]Scan, error) {

	queries := strings.Split(params, " ")

	for	_, query := range queries {
		fmt.Println(query)
	}

	if len(queries) == 0 {
		return nil, errors.New("No queries provided") 
	}

	// if cidrRe.MatchString(queries[0]) {
	// 	return parseCidr(queries[0]), nil
	// }
	

	// if regexp.MatchString(p)

	// return "hi"
	return nil, nil
}

func parseCidr(params string) string {
	return fmt.Sprintf("%v", params)
}

func parseNet(params []string) string {
	return fmt.Sprintf("%v", params)
}

func parsePort(params []string) string {
	return fmt.Sprintf("%v", params)
}

func parseDomain(params []string) string {
	return fmt.Sprintf("%v", params)
}


