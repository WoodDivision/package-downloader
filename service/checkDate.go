package service

import (
	"time"
)

func CheckDate(pac time.Time) bool {
	tYear, tMonth, tDay := time.Now().Date()
	if pac.Day() <= tDay && pac.Month() >= (tMonth-2) && pac.Year() == tYear {
		return false
	}
	return true
}
