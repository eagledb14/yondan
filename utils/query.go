package utils

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
var queryRe *regexp.Regexp = regexp.MustCompile(`^[a-zA-z]+`)

func Query(params string, db *ConcurrentMap) ([]Scan, error) {

	queries := strings.Split(params, " ")

	if len(queries) == 0 {
		return nil, errors.New("No queries provided")
	}

	queryScans := [][]Scan{}

	re := regexp.MustCompile(`[,:]`)
	for _, q := range queries {
		res := []Scan{}

		if cidrRe.MatchString(q){
			res = parseCidr(q, db)
		} else if netRe.MatchString(q) {
			res = parseNet(re.Split(q, -1), db)
		} else if queryRe.MatchString(q) {
			res = parseQuery(re.Split(q, -1), db)
		} else {
			res = parseString(re.Split(q, -1), db)
		}
		queryScans = append(queryScans, res)
	}

	return filter(queryScans, db), nil
}

func filter(scans [][]Scan, db *ConcurrentMap) []Scan {
	ipCount := map[string]int{}
	scanCount := len(scans)

	for _, scan := range scans {
		for _, s := range scan {
			ipCount[s.Ip]++
		}
	}

	filteredScans := []Scan{}
	for key, value := range ipCount {
		if value == scanCount {
			scan, err := db.Read(key)
			if err != nil {
				continue
			}
			filteredScans = append(filteredScans, scan...)
		}
	}

	return filteredScans
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

func parseString(params []string, db *ConcurrentMap) []Scan {
	scans := []Scan{}
	index := db.ReadAll()
	for _, p := range params[1:] {
		for key, value := range index {
			if strings.Contains(key, p) {
				scans = append(scans, value...)
			}
		}
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
