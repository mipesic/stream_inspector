package domain

import (
	"time"
)

type Channel struct {
	ID string
}

type WatchPoint struct {
	Channel Channel
	Timestamp time.Time
	Streams []Stream
}
