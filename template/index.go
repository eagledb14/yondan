package template


import (
)

func Index(searchIP string) string {

    data := struct {
        Banner string
    } {
        Banner: Banner(searchIP),
    }

    const indexPage = `
        {{.Banner}}
        <div class="index-grid">
            <div class="one-item-left">
                <h1 class="heading-box">DashBoard</h1>
            </div>

            <div class="three-items">
                <div class="content-box top-border blue-border">
                    <h2>
                        <svg class="svg-inline--fa fa-book fa-w-14 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="book" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" data-fa-i2svg=""><path fill="white" d="M128 152v-32c0-4.4 3.6-8 8-8h208c4.4 0 8 3.6 8 8v32c0 4.4-3.6 8-8 8H136c-4.4 0-8-3.6-8-8zm8 88h208c4.4 0 8-3.6 8-8v-32c0-4.4-3.6-8-8-8H136c-4.4 0-8 3.6-8 8v32c0 4.4 3.6 8 8 8zm299.1 159.7c-4.2 13-4.2 51.6 0 64.6 7.3 1.4 12.9 7.9 12.9 15.7v16c0 8.8-7.2 16-16 16H80c-44.2 0-80-35.8-80-80V80C0 35.8 35.8 0 80 0h352c8.8 0 16 7.2 16 16v368c0 7.8-5.5 14.2-12.9 15.7zm-41.1.3H80c-17.6 0-32 14.4-32 32 0 17.7 14.3 32 32 32h314c-2.7-17.3-2.7-46.7 0-64zm6-352H80c-17.7 0-32 14.3-32 32v278.7c9.8-4.3 20.6-6.7 32-6.7h320V48z"></path></svg>
                        Getting Started
                    </h2>
                    <div class="fake-link">
                        <div>What is Shodan?</div>
                        <div>Search Query Fundamentals</div>
                        <div>Working with Shodan Data Files</div>
                    </div>
                </div>
                <div class="content-box top-border red-border">
                    <h2>
                        <svg class="svg-inline--fa fa-terminal fa-w-20 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="terminal" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512" data-fa-i2svg=""><path fill="white" d="M41.678 38.101l209.414 209.414c4.686 4.686 4.686 12.284 0 16.971L41.678 473.899c-4.686 4.686-12.284 4.686-16.971 0L4.908 454.101c-4.686-4.686-4.686-12.284 0-16.971L185.607 256 4.908 74.87c-4.686-4.686-4.686-12.284 0-16.971L24.707 38.1c4.686-4.686 12.284-4.686 16.971.001zM640 468v-28c0-6.627-5.373-12-12-12H300c-6.627 0-12 5.373-12 12v28c0 6.627 5.373 12 12 12h328c6.627 0 12-5.373 12-12z"></path></svg>
                        ASCII Videos 
                    </h2>
                    <div class="fake-link">
                        <div>Setting up Real-Time Network Monitoring</div>
                        <div>Measuring Public SMB Exposure</div>
                        <div>Analyzing the Vulnerabilities for a Network</div>
                    </div>
                </div>
                <div class="content-box top-border yellow-border">
                    <h2>
                        <svg class="svg-inline--fa fa-code fa-w-18 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="code" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512" data-fa-i2svg=""><path fill="white" d="M234.8 511.7L196 500.4c-4.2-1.2-6.7-5.7-5.5-9.9L331.3 5.8c1.2-4.2 5.7-6.7 9.9-5.5L380 11.6c4.2 1.2 6.7 5.7 5.5 9.9L244.7 506.2c-1.2 4.3-5.6 6.7-9.9 5.5zm-83.2-121.1l27.2-29c3.1-3.3 2.8-8.5-.5-11.5L72.2 256l106.1-94.1c3.4-3 3.6-8.2.5-11.5l-27.2-29c-3-3.2-8.1-3.4-11.3-.4L2.5 250.2c-3.4 3.2-3.4 8.5 0 11.7L140.3 391c3.2 3 8.2 2.8 11.3-.4zm284.1.4l137.7-129.1c3.4-3.2 3.4-8.5 0-11.7L435.7 121c-3.2-3-8.3-2.9-11.3.4l-27.2 29c-3.1 3.3-2.8 8.5.5 11.5L503.8 256l-106.1 94.1c-3.4 3-3.6 8.2-.5 11.5l27.2 29c3.1 3.2 8.1 3.4 11.3.4z"></path></svg>
                        Developer Access
                    </h2>
                    <div class="fake-link">
                        <div>How to Download Data with the API</div>
                        <div>Looking up IP Information</div>
                        <div>Working with Shodan Data Files</div>
                    </div>
                </div>
            </div>

            <div class="two-items">
                <div class="content-box top-border green-border">
                    <h2>
                        <svg class="svg-inline--fa fa-download fa-w-18 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="download" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512" data-fa-i2svg=""><path fill="white" d="M528 288h-92.1l46.1-46.1c30.1-30.1 8.8-81.9-33.9-81.9h-64V48c0-26.5-21.5-48-48-48h-96c-26.5 0-48 21.5-48 48v112h-64c-42.6 0-64.2 51.7-33.9 81.9l46.1 46.1H48c-26.5 0-48 21.5-48 48v128c0 26.5 21.5 48 48 48h480c26.5 0 48-21.5 48-48V336c0-26.5-21.5-48-48-48zm-400-80h112V48h96v160h112L288 368 128 208zm400 256H48V336h140.1l65.9 65.9c18.8 18.8 49.1 18.7 67.9 0l65.9-65.9H528v128zm-88-64c0-13.3 10.7-24 24-24s24 10.7 24 24-10.7 24-24 24-24-10.7-24-24z"></path></svg>
                        Enterprice Access
                    </h2>
                    <p class="description">Need bulk data access? Check out our enterprise offering which includes full, unlimited access to the entire Shodan platform:</p>
                </div>

                <div class="content-box">
                    <h2>Filters Cheat Sheet</h2>
                    <p class="description">Yondan currently crawls nearly 1,500 ports across the Internet. Here are a few of the most commonly-used search filters to get started.</p>

                    <table class="table">
                        <tbody>
                            <tr>
                                <th>Filter Name</th>
                                <th>Description</th>
                                <th>Example</th>
                            </tr>
                            <tr>
                                <th>net</th>
                                <td>Network range or IP in CIDR notation</td>
                                <td><a href="/query/net:8.8.0.0/24">Services in the range of 8.8.0.0 to 8.8.255.255</a></td>
                            </tr>
                            <tr>
                                <th>port</th>
                                <td>Port number for the service that is running</td>
                                <td><a href="/query/port:22">SSH servers</a></td>
                            </tr>
                            <tr>
                                <th>domain</th>
                                <td>Domain name pulled from DNS records</td>
                                <td><a href="/query/domain:google.com">Put Something Here</a></td>
                            </tr>
        
                        </tbody>
                    </table>
                </div>
            </div>

        </div>
        `

    return ExecuteText("index", indexPage, data)
}

