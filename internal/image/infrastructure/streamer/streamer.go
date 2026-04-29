package streamer

import (
	"context"
	"io"
)

type ImageStreamer interface {
	Stream(ctx context.Context, url string) (io.ReadCloser, error)
}
