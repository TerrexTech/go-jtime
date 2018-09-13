package jtime

import (
	"time"
)

type JTime struct {
	time.Time
}

func (j *JTime) UnmarshalJSON(jsonString []byte) error {
	s := string(jsonString)

	s = s[1 : len(s)-1]
	changedTime, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return err
	}

	j.Time = changedTime
	return nil
}
