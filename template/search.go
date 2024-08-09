package template

import (
	"sort"

	"github.com/eagledb14/shodan-clone/utils"
)

func Search(scans []utils.Scan, query string) string {
    totalScan := len(scans)
    ports := sortPorts(scans)
    if len(ports) > 10 {
        ports = ports[:10]
    }

    services := sortServices(scans)
    if len(services) > 10 {
        services = services[:10]
    }

    data := struct {
        Scans []utils.Scan
        TotalScan int
        SortedPorts []sortedPorts
        SortedServices []sortedServices
        Query string
    } {
        Scans: scans,
        TotalScan: totalScan,
        SortedPorts: ports,
        SortedServices: services,
        Query: query,
    }

    const searchPage = `
        <div class="index-grid">
            <div class="two-items">
                <div>
                    <div class="subheader">TOTAL RESULTS:</div><br>
                    <div class="large-data">{{ .TotalScan }}</div>
                    <br>
                    <div class="subheader">TOP PORTS</div>

                    {{range .SortedPorts}}
                        <a href="/search?query={{$.Query}}%20port%3A{{.Port}}" class="display-port-grid">
                            {{.Port}}: {{.Count}}<br>
                        </a>
                    {{end}}
                    <br>

                    <div class="subheader">TOP SERVICES</div>

                    {{range .SortedServices}}
                        <a href="/search?query={{$.Query}}%20service%3A{{.Service}}" class="display-port-grid">
                            {{.Service}}: {{.Count}}<br>
                        </a>
                    {{end}}
                </div>
                <div>
                    {{range .Scans}}
                        <div class="scan-display">
                            <a href="/host/{{.Ip}}">{{.Ip}}</a>
                            <br>
                            {{if ne .Hostname ""}}
                                <p class="item">Hostname: {{.Hostname}}</p>
                            {{end}}
                            {{if ne (len .Ports) 0}}
                                <p class="item">Ports: {{range .Ports}} {{.ID}},  {{end}}</p>
                            {{end}}
                            <p class="item">TimeStamp: {{.Timestamp}}</p>
                        </div>
                    {{end}}
                </div>
            </div>
        </div>
    `


    return Execute("search", searchPage, data)
}

type sortedPorts struct {
    Port uint16
    Count int
}

func sortPorts(scans []utils.Scan) []sortedPorts {
    portCount := make(map[uint16]int)

    for _, scan := range scans {
        for _, port := range scan.Ports {
            portCount[port.ID]++
        }
    }

    ports := []sortedPorts{}

    for port, count :=  range portCount {
        ports = append(ports, sortedPorts{port, count})
    }

    sort.Slice(ports, func(i, j int) bool {
        return ports[i].Count > ports[j].Count
    })

    return ports
}

type sortedServices struct {
    Service string
    Count int
}

func sortServices(scans []utils.Scan) []sortedServices {
    serviceCount := make(map[string]int)

    for _, scan := range scans {
        for _, port := range scan.Ports {
            serviceCount[port.Service.Name]++
        }
    }

    services := []sortedServices{}
    for service, count := range serviceCount {
        services = append(services, sortedServices{service, count})
    }

    sort.Slice(services, func(i, j int) bool {
        return services[i].Count > services[j].Count
    })

    return services
}
