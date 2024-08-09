package template

import (
    "github.com/eagledb14/shodan-clone/utils"
)

func Host(scan utils.Scan , db *utils.ConcurrentMap) string {

    data := struct {
        Scan utils.Scan
    } {
        Scan: scan,
    }

    const hostPage = `
        <div class="host-grid">
            <div class="one-item-left">
                <h1 class="ip-title">
                    {{.Scan.Ip}}
                </h1>
                <div></div>
            </div>
            <div class="two-items">
                <div>
                    <div class="content-box top-border yellow-border">
                        <div class="content-title">
                            <svg class="svg-inline--fa fa-globe fa-w-16 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="globe" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 496 512" data-fa-i2svg=""><path fill="white" d="M248 8C111 8 0 119 0 256s111 248 248 248 248-111 248-248S385 8 248 8zm179.3 160h-67.2c-6.7-36.5-17.5-68.8-31.2-94.7 42.9 19 77.7 52.7 98.4 94.7zM248 56c18.6 0 48.6 41.2 63.2 112H184.8C199.4 97.2 229.4 56 248 56zM48 256c0-13.7 1.4-27.1 4-40h77.7c-1 13.1-1.7 26.3-1.7 40s.7 26.9 1.7 40H52c-2.6-12.9-4-26.3-4-40zm20.7 88h67.2c6.7 36.5 17.5 68.8 31.2 94.7-42.9-19-77.7-52.7-98.4-94.7zm67.2-176H68.7c20.7-42 55.5-75.7 98.4-94.7-13.7 25.9-24.5 58.2-31.2 94.7zM248 456c-18.6 0-48.6-41.2-63.2-112h126.5c-14.7 70.8-44.7 112-63.3 112zm70.1-160H177.9c-1.1-12.8-1.9-26-1.9-40s.8-27.2 1.9-40h140.3c1.1 12.8 1.9 26 1.9 40s-.9 27.2-2 40zm10.8 142.7c13.7-25.9 24.4-58.2 31.2-94.7h67.2c-20.7 42-55.5 75.7-98.4 94.7zM366.3 296c1-13.1 1.7-26.3 1.7-40s-.7-26.9-1.7-40H444c2.6 12.9 4 26.3 4 40s-1.4 27.1-4 40h-77.7z"></path></svg>
                            <b>General</b>
                            Information
                        </div>
                        <table class="table">
                            <tbody>
                                <tr>
                                    <th>Hostnames</td>
                                    <td>{{.Scan.Hostname}}</td>
                                </tr>
                                <tr>
                                    <th>Last Seen</th>
                                    <td>{{.Scan.Timestamp}}
                                </tr>
                            </tbody>
                        </table>
                    </div>
                </div>
                <div>
                    <div class="content-box top-border blue-border">
                        <div class="content-title">
                            <svg class="svg-inline--fa fa-sitemap fa-w-20 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="sitemap" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512" data-fa-i2svg=""><path fill="white" d="M104 272h192v48h48v-48h192v48h48v-57.59c0-21.17-17.22-38.41-38.41-38.41H344v-64h40c17.67 0 32-14.33 32-32V32c0-17.67-14.33-32-32-32H256c-17.67 0-32 14.33-32 32v96c0 8.84 3.58 16.84 9.37 22.63S247.16 160 256 160h40v64H94.41C73.22 224 56 241.23 56 262.41V320h48v-48zm168-160V48h96v64h-96zm336 240h-96c-17.67 0-32 14.33-32 32v96c0 17.67 14.33 32 32 32h96c17.67 0 32-14.33 32-32v-96c0-17.67-14.33-32-32-32zm-16 112h-64v-64h64v64zM368 352h-96c-17.67 0-32 14.33-32 32v96c0 17.67 14.33 32 32 32h96c17.67 0 32-14.33 32-32v-96c0-17.67-14.33-32-32-32zm-16 112h-64v-64h64v64zM128 352H32c-17.67 0-32 14.33-32 32v96c0 17.67 14.33 32 32 32h96c17.67 0 32-14.33 32-32v-96c0-17.67-14.33-32-32-32zm-16 112H48v-64h64v64z"></path></svg>
                            Open <b>Ports</b>
                        </div>
                        <div class=port-grid>
                            {{range .Scan.Ports}}
                                <a class="port-box" href="#{{.ID}}">
                                    {{.ID}}
                                </a>
                            {{end}}
                        </div>
                    </div>
                    <div>
                    {{range .Scan.Ports}}
                        <br>
                        <div class="highlight">
                            // {{.ID}} / {{.Protocol}}
                        <a href="http://{{$.Scan.Ip}}:{{.ID}}" target="_blank">
                            <svg class="svg-inline--fa fa-external-link fa-w-16 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="external-link" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="currentColor" d="M497.6,0,334.4.17A14.4,14.4,0,0,0,320,14.57V47.88a14.4,14.4,0,0,0,14.69,14.4l73.63-2.72,2.06,2.06L131.52,340.49a12,12,0,0,0,0,17l23,23a12,12,0,0,0,17,0L450.38,101.62l2.06,2.06-2.72,73.63A14.4,14.4,0,0,0,464.12,192h33.31a14.4,14.4,0,0,0,14.4-14.4L512,14.4A14.4,14.4,0,0,0,497.6,0ZM432,288H416a16,16,0,0,0-16,16V458a6,6,0,0,1-6,6H54a6,6,0,0,1-6-6V118a6,6,0,0,1,6-6H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V304A16,16,0,0,0,432,288Z"></path></svg>
                        </a>
                        </div>
                        {{if .Service.Name}}
                            <div class="content-box" id="{{.ID}}">
                                {{.State}}
                                {{.Service.Name}}
                            </div>
                        {{end}}
                    {{end}}
                </div>
            </div>

        </div>
    `

    return Execute("host", hostPage, data)
}
