package nostrhttp

import "net/http"

var DefaultClient = &http.Client{Transport: http.DefaultTransport}

// TODO
