package types

import "time"

type Log struct {
	Timestamp time.Time
	Username  string
	Operation string
	Size      int
}
