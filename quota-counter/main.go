package main

import (
	"flag"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	listen := flag.String("listen", ":8080", "Server Listen Address")

	flag.Parse()

	var svc QuotaService
	svc = quotaService{}

	quotaServiceHandler := httptransport.NewServer(
		makeQuotaEndpoint(svc),
		decodeQuotaRequest,
		encodeResponse,
	)

	http.Handle("/quota", quotaServiceHandler)

	http.ListenAndServe(*listen, nil)
}
