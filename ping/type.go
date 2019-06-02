package ping

import (
	"time"
)

// Ping data struct
type Ping struct {
	target *Target
	done   chan struct{}
	result *Result
}

// Target data struct
type Target struct {
	Host string
	Port uint16

	Counter  int
	Interval time.Duration
	Timeout  time.Duration
}

// Result data struct
type Result struct {
	Counter        int
	SuccessCounter int
	Target         *Target

	MinDuration   time.Duration
	MaxDuration   time.Duration
	TotalDuration time.Duration
}
