package template

import (
    "html/template"
    text "text/template"
    "bytes"
)

func Execute(name string, t string, data interface{}) string {
    tmpl, err := template.New(name).Parse(t)
    if err != nil {
        return err.Error()
    }
    var b bytes.Buffer
    err = tmpl.Execute(&b, data)
    if err != nil {
        return err.Error()
    }

    return b.String()
}

//The reason both functions are needed is because html/template sanitizes
//the html input, which is something we want, unless we already
//sanitized the input
func ExecuteText(name string, t string, data interface{}) string {
    tmpl, err := text.New(name).Parse(t)
    if err != nil {
        return err.Error()
    }
    var b bytes.Buffer
    err = tmpl.Execute(&b, data)
    if err != nil {
        return err.Error()
    }

    return b.String()
}

func Banner(searchIP string) string {
    data := struct {
        SearchIP string
    } {
        SearchIP: searchIP,
    }

    const banner = `
    <div class="banner">
        <div class="banner-box">
            <div class="banner-button" hx-get="/" hx-push-url="true" hx-target="body">
                <img src="/logo"></img>
                Yondan
            </div>
            <div class="banner-button banner-shrink">
                Explore
            </div>
            <input id="search" name="query" type="text" placeholder="Type / to search" value="{{.SearchIP}}"></input>
            <button id="submitSearch" class="banner-search-button">
                <svg class="svg-inline--fa fa-search fa-w-16 fa-fw" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="white" d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path></svg>
            </button>
            <script>
                document.getElementById('submitSearch').addEventListener('click', function(event) {
                    event.preventDefault()
                    
                    let input = document.getElementById('search').value
                    window.location.href = '/search?query=' + encodeURIComponent(input)
                })

                document.addEventListener('keydown', function(event) {
                    const searchBox = document.getElementById('search')
                    if (event.key ==='/' && document.activeElement !== searchBox) {
                        event.preventDefault()
                        searchBox.focus()
                    }
                    if (event.key === 'Enter' && document.activeElement === searchBox) {
                        event.preventDefault(); 
                        let input = document.getElementById('search').value
                        window.location.href = '/search?query=' + encodeURIComponent(input)
                    }
                })
            </script>
        </div>
    </div>
    `

    return Execute("banner", banner, data)
}

func header() string {
    return `
        <head>
            <title>Yondan Search Engine</title>
            <script src="/htmx"></script>
            <link rel="shortcut icon" type="image/png" href="/favicon.ico"/>
            <link rel="stylesheet" type="text/css" href="/styles.css">
        </head>
        `
}

func BuildPage(body string, searchQuery string) string {
    data := struct {
        Header string
        Body string
        Banner string
    } {
        Header: header(),
        Body: body,
        Banner: Banner(searchQuery),
    }
    const page = `
        <!DOCTYPE html>
        <html lang="en">
        {{.Header}}
        <body hx-boost="true">
            {{.Banner}}
            {{.Body}}
        </body>
        </html>
        `

    return ExecuteText("page", page, data)
}



