package service

import (
	"time"
)

func CheckDate(pac time.Time) bool {
	today := time.Now()
	if today.Before(pac) {
		return false
	}
	return true
}
