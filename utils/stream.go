package utils

import (
	"context"
	"errors"
	"sync"
)

// StreamWaitFn is a method used to wait for a stream to end, catching any error in the process.
type StreamWaitFn func() error

type StreamerCallback[Data any] func(ctx context.Context, inC chan<- Data) error

var ErrStreamerClosed = errors.New("stream closed")

func NewStreamer[Data any](ctx context.Context, callback StreamerCallback[Data]) (<-chan Data, StreamWaitFn) {
	var err error

	outC := make(chan Data)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer func() { close(outC) }()
		defer wg.Done()

		for {
			if err = callback(ctx, outC); err != nil {
				return
			}
		}
	}()

	return outC, func() error {
		wg.Wait()

		if errors.Is(err, ErrStreamerClosed) {
			return nil
		}

		return err
	}
}
