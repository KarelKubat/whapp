package streamreplaced

import "go.mau.fi/whatsmeow/types/events"

type streamreplaced struct{}

func New(s *events.StreamReplaced) *streamreplaced {
	return &streamreplaced{}
}

func (s *streamreplaced) String() string {
	return "StreamReplaced"
}
