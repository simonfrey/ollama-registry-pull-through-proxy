package dumptransport

import (
	"github.com/rs/zerolog/log"

	"moul.io/http2curl"
	"net/http"
	"net/http/httputil"
	"time"
)

type Transport struct {
}

func (p *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqDump, _ := httputil.DumpRequestOut(req, true)
	curlDump, _ := http2curl.GetCurlCommand(req)

	requestStart := time.Now()
	resp, err := http.DefaultTransport.RoundTrip(req)
	requestDuration := time.Since(requestStart)
	if err != nil {
		log.Error().Str("duration", requestDuration.String()).
			Str("request", string(reqDump)).
			Str("request_curl", curlDump.String()).
			Err(err).
			Msgf("%s %s Could not be loaded as of error", req.Method, req.URL)
		return nil, err
	}
	respDump, _ := httputil.DumpResponse(resp, true)
	log.Info().Str("duration", requestDuration.String()).
		Str("request", string(reqDump)).
		Str("request_curl", curlDump.String()).
		Str("response", string(respDump)).
		Msgf("%s %s", req.Method, req.URL)
	return resp, err
}
