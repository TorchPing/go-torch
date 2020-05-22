package resolve

import (
	"time"
)

// Ping data struct
type Resolve struct {
	target *Target
	done   chan struct{}
	result *Result
}

// Target data struct
type Target struct {
	Host string

	Counter  int
	Interval time.Duration
	Timeout  time.Duration
}

// Result data struct
type Result struct {
	Counter        int
	SuccessCounter int
	Target         *Target
	Addrs          []string

	MinDuration   time.Duration
	MaxDuration   time.Duration
	TotalDuration time.Duration
}
