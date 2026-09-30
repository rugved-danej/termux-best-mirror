package mirror

import "time"

type Result struct {
	Path     string
	Name     string
	Duration time.Duration
	Success  bool
}
