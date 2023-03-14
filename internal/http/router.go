package webserver

import(
    "log"
    "net/http"
)

func getDo(w http.ResponseWriter, r *http,Request) {
    fmt.Fprint(w, "Return Do")
}
