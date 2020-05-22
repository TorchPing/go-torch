package resolve

import (
	"fmt"
	"net"
	"time"

	"github.com/TorchPing/go-torch/pkg/utils"
)

// NewResolve Create Resolve Instane
func NewResolve() *Resolve {
	resolve := Resolve{
		done: make(chan struct{}),
	}
	return &resolve
}

// SetTarget ..
func (resolve *Resolve) SetTarget(target *Target) {
	resolve.target = target
	if resolve.result == nil {
		resolve.result = &Result{Target: target}
	}
}

// Start the process
func (resolve Resolve) Start() <-chan struct{} {
	go func() {
		t := time.NewTicker(time.Nanosecond)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				if resolve.result.Counter >= resolve.target.Counter && resolve.target.Counter != 0 {
					resolve.done <- struct{}{}
					return
				}

				duration, addrs := resolve.resolve()
				resolve.result.Counter++
				if duration == 0 {
					fmt.Printf("Resolve %s - failed: %s\n", resolve.target.Host, addrs)
				} else {
					if resolve.result.MinDuration == 0 {
						resolve.result.MinDuration = duration
					}
					if resolve.result.MaxDuration == 0 {
						resolve.result.MaxDuration = duration
					}
					resolve.result.SuccessCounter++
					if duration > resolve.result.MaxDuration {
						resolve.result.MaxDuration = duration
					} else if duration < resolve.result.MinDuration {
						resolve.result.MinDuration = duration
					}
					resolve.result.TotalDuration += duration
					resolve.result.Addrs = addrs.([]string)
				}
			case <-resolve.done:
				return
			}
		}
	}()
	return resolve.done
}

// Result return the result
func (resolve Resolve) Result() *Result {
	return resolve.result
}

func (resolve Resolve) resolve() (time.Duration, interface{}) {
	duration, res, errIfce := utils.TimeItWithResult(func() (interface{}, interface{}) {
		host := fmt.Sprintf("%s", resolve.target.Host)

		addrs, err := net.LookupHost(host)

		return addrs, err
	})

	if errIfce != nil {
		err := errIfce.(error)
		return 0, err
	}

	return time.Duration(duration), res
}
