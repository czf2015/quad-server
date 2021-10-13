package media

import (
	"time"
	"io"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Streamer interface {
	Stream() (io.ReaderCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Audio interface {
	Streamer
}

type Vedio interface {
	Streamer
	Resolution() (x, y int)
}