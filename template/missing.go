package template


import (

)

func Missing() string {
    const indexPage = `
        <div class="alert">
            <div></div>
            <div class="alert-items">
                <div class="alert-icon left-border dark-blue-border">
                    <div><svg class="svg-inline--fa fa-info fa-w-8 fa-fw" aria-hidden="true" focusable="false" data-prefix="far" data-icon="info" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 512" data-fa-i2svg=""><path fill="white" d="M224 352.589V224c0-16.475-6.258-31.517-16.521-42.872C225.905 161.14 236 135.346 236 108 236 48.313 187.697 0 128 0 68.313 0 20 48.303 20 108c0 20.882 5.886 40.859 16.874 58.037C15.107 176.264 0 198.401 0 224v39.314c0 23.641 12.884 44.329 32 55.411v33.864C12.884 363.671 0 384.359 0 408v40c0 35.29 28.71 64 64 64h128c35.29 0 64-28.71 64-64v-40c0-23.641-12.884-44.329-32-55.411zM128 48c33.137 0 60 26.863 60 60s-26.863 60-60 60-60-26.863-60-60 26.863-60 60-60zm80 400c0 8.836-7.164 16-16 16H64c-8.836 0-16-7.164-16-16v-40c0-8.836 7.164-16 16-16h16V279.314H64c-8.836 0-16-7.164-16-16V224c0-8.836 7.164-16 16-16h96c8.836 0 16 7.164 16 16v168h16c8.836 0 16 7.164 16 16v40z"></path></svg><!-- <i class="far fa-info  fa-fw "></i> Font Awesome fontawesome.com --></div>
                </div>
                <div class="alert-notice">
                    <b>Note:</b>
                    No Results Found
                </div>
            </div>
            <div></div>
        </div>
        `

    return indexPage
}
