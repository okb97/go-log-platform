package model

import "time"

type LogEntry struct {
	TimeStamp time.Time
	Level     string
	Message   string
}
