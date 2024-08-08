package main

import (
	"errors"
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
var domainRe *regexp.Regexp = regexp.MustCompile(`^domain:`)
var queryRe *regexp.Regexp = regexp.MustCompile(`^[a-zA-z]+`)

func query(params string, db *ConcurrentMap) ([]Scan, error) {

	queries := strings.Split(params, " ")

	if len(queries) == 0 {
		return nil, errors.New("No queries provided")
	}

	if cidrRe.MatchString(queries[0]) {
		return parseCidr(queries[0], db), nil
	}

	queryScans := [][]Scan{}

	re := regexp.MustCompile(`[,:]`)
	for _, q := range queries {
		res := []Scan{}
		if netRe.MatchString(q) {
			res = parseNet(re.Split(q, -1), db)
		} else if queryRe.MatchString(q) {
			res = parseQuery(re.Split(q, -1), db)
		}
		queryScans = append(queryScans, res)
	}

	return filter(queryScans), nil
}

func filter(scans [][]Scan) []Scan {

	return scans [0]
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
			scanList = append(scanList, v...)
		}
	}

	return scanList
}

func parseNet(params []string, db *ConcurrentMap) []Scan {
	scans := []Scan{}
	for _, p := range params[1:] {
		scans = append(scans, parseCidr(p, db)...)
	}
	return scans
}

func parseQuery(params []string, db *ConcurrentMap) []Scan {
	scans := []Scan{}
	for _, p := range params[1:] {
		run, err:= db.Read(p)
		if err != nil {
			continue
		}
		scans = append(scans, run...)
	}

	return scans
}

// func parseDomain(params []string) []Scan {
// 	// return fmt.Sprintf("%v", params)
// 	return nil
// }
