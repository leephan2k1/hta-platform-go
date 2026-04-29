package streamer

import (
	"context"
	"fmt"
	"io"

	"github.com/go-resty/resty/v2"
)

type mmStreamer struct {
	client  *resty.Client
	referer string
}

func NewMMStreamer(referer string) ImageStreamer {
	return &mmStreamer{
		client:  resty.New(),
		referer: referer,
	}
}

func (s *mmStreamer) Stream(ctx context.Context, url string) (io.ReadCloser, error) {
	resp, err := s.client.R().
		SetContext(ctx).
		SetHeaders(map[string]string{
			"sec-ch-ua-platform": "\"macOS\"",
			"Referer":            s.referer,
			"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/145.0.0.0 Safari/537.36",
			"sec-ch-ua":          "\"Chromium\";v=\"145\", \"Not:A-Brand\";v=\"99\"",
			"sec-ch-ua-mobile":   "?0",
		}).
		SetDoNotParseResponse(true).
		Get(url)

	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf("failed to fetch image: %s", resp.Status())
	}

	return resp.RawBody(), nil
}
