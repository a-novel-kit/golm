package models

import "github.com/samber/lo"

type StreamOptions struct {
	// If set, an additional chunk will be streamed before the data: [DONE] message. The usage field on this chunk
	// shows the token usage statistics for the entire request, and the choices field will always be an empty array.
	// All other chunks will also include a usage field, but with a null value.
	IncludeUsage bool `json:"include_usage"`
}

// Stream allows partial message deltas to be sent. Tokens will be sent as data-only server-sent events as they
// become available, with the stream terminated by a data: [DONE] message.
type Stream bool

func NewStream(s bool) *Stream {
	return lo.ToPtr(Stream(s))
}
