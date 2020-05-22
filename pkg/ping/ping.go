package ping

import (
	"fmt"
	"net"
	"time"

	"github.com/TorchPing/go-torch/pkg/utils"
)

// NewPing Create Ping Instane
func NewPing() *Ping {
	ping := Ping{
		done: make(chan struct{}),
	}
	return &ping
}

// SetTarget ..
func (ping *Ping) SetTarget(target *Target) {
	ping.target = target
	if ping.result == nil {
		ping.result = &Result{Target: target}
	}
}

// Start the process
func (ping Ping) Start() <-chan struct{} {
	go func() {
		t := time.NewTicker(time.Nanosecond)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				if ping.result.Counter >= ping.target.Counter && ping.target.Counter != 0 {
					ping.done <- struct{}{}
					return
				}

				duration, err := ping.ping()
				ping.result.Counter++
				if err != nil {
					fmt.Printf("Ping %s - failed: %s\n", ping.target.Host, err)
				} else {
					if ping.result.MinDuration == 0 {
						ping.result.MinDuration = duration
					}
					if ping.result.MaxDuration == 0 {
						ping.result.MaxDuration = duration
					}
					ping.result.SuccessCounter++
					if duration > ping.result.MaxDuration {
						ping.result.MaxDuration = duration
					} else if duration < ping.result.MinDuration {
						ping.result.MinDuration = duration
					}
					ping.result.TotalDuration += duration
				}
			case <-ping.done:
				return
			}
		}
	}()
	return ping.done
}

// Result return the result
func (ping Ping) Result() *Result {
	return ping.result
}

func (ping Ping) ping() (time.Duration, error) {
	duration, errIfce := utils.TimeIt(func() interface{} {
		host := fmt.Sprintf("%s:%d", ping.target.Host, ping.target.Port)

		conn, err := net.DialTimeout("tcp", host, ping.target.Timeout)
		if err != nil {
			return err
		}
		conn.Close()
		return nil
	})

	if errIfce != nil {
		err := errIfce.(error)
		return 0, err
	}
	return time.Duration(duration), nil
}
