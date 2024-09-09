package utils

import (
	"errors"
	"net"
	"fmt"
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
var queryRe *regexp.Regexp = regexp.MustCompile(`^[a-zA-z]+:.*`)
var filterRe *regexp.Regexp = regexp.MustCompile(`^-`)

func Query(params string, db *ConcurrentMap) ([]*Scan, error) {

	queries := strings.Split(params, " ")

	if len(queries) == 0 {
		return nil, errors.New("No queries provided")
	}

	queryScans := [][]*Scan{}
	queryCount := 0

	re := regexp.MustCompile(`[,:]`)
	filterRes := []string{}
	for _, q := range queries {
		res := []*Scan{}

		if cidrRe.MatchString(q){
			res = parseCidr(q, db)
			queryCount++
		} else if netRe.MatchString(q) {
			params := re.Split(q, -1)
			queryCount += len(params) - 1
			res = parseNet(params, db)
		} else if queryRe.MatchString(q) {
			params := re.Split(q, -1)
			queryCount += len(params) - 1
			res = parseQuery(params, db)
		} else if filterRe.MatchString(q) { // This handles picking ips that should be filtered
			fmt.Println(q)
			filterQ, _ := Query(q[1:], db)
			for _, query := range filterQ {
				filterRes = append(filterRes, query.Ip)
			}
		} else {
			queryCount ++
			res = parseString(q, db)
		}
		queryScans = append(queryScans, res)
	}

	foundQueries := filterQueriesCount(queryScans, queryCount, db)
	return removeFilteredQueries(foundQueries, filterRes), nil
}

// removes ips that are not in every query
// so if there are 5 queries and the ip only shows up 3 times, then it removes the ip from the search
func filterQueriesCount(scans [][]*Scan, queryCount int, db *ConcurrentMap) []*Scan {
	ipCount := map[string]int{}

	for _, scan := range scans {
		for _, s := range scan {
			ipCount[s.Ip]++
		}
	}

	filteredScans := []*Scan{}
	for key, value := range ipCount {
		if value == queryCount {
			scan, err := db.Read(key)
			if err != nil {
				continue
			}
			filteredScans = append(filteredScans, scan...)
		}
	}

	return filteredScans
}

// removes ips that are to be removed from the query
func removeFilteredQueries(scans []*Scan, filterScans []string) []*Scan{
	filteredScans := []*Scan{}

	for _, scan := range scans {
		included := false
		for _, filterHost := range filterScans {
			if filterHost == scan.Ip {
				included = true 
				break
			}
		}

		if included == false {
			filteredScans = append(filteredScans, scan)
		}
	}

	return filteredScans
}

func parseCidr(params string, db *ConcurrentMap) []*Scan {

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
	scanList := []*Scan{}
	for key, v := range index {
		if ipList.Contains(net.ParseIP(key)) {
			scanList = append(scanList, v...)
		}
	}

	return scanList
}

func parseNet(params []string, db *ConcurrentMap) []*Scan {
	scans := []*Scan{}
	for _, p := range params[1:] {
		scans = append(scans, parseCidr(p, db)...)
	}
	return scans
}

func parseString(params string, db *ConcurrentMap) []*Scan {
	scans := []*Scan{}
	index := db.ReadAll()

	for key, value := range index {
		if strings.Contains(key, params) {
			scans = append(scans, value...)
		}
	}

	return scans
}

func parseQuery(params []string, db *ConcurrentMap) []*Scan {
	scans := []*Scan{}
	for _, p := range params[1:] {
		run, err:= db.Read(p)
		if err != nil {
			continue
		}
		scans = append(scans, run...)
	}

	return scans
}
