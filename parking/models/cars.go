package models

import "time"

type Cars struct {
	Type  string
	Regd  int64
	Entry time.Time
}
