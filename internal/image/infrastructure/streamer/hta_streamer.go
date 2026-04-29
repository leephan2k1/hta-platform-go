package streamer

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

type htaStreamer struct {
}

func NewHTAStreamer() ImageStreamer {
	return &htaStreamer{}
}

func (s *htaStreamer) Stream(ctx context.Context, url string) (io.ReadCloser, error) {
	// Placeholder for now
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed to fetch image: %s", resp.Status)
	}
	return resp.Body, nil
}
