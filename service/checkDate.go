package service

import (
	"time"
)

func CheckDate(pac time.Time) bool {
	tYear, tMonth, tDay := time.Now().Date()
	if pac.Year() == tYear {
		if pac.Month() >= tMonth-2 {
			if pac.Day() <= tDay {
				return false
			}
		}
	}
	return true
}
