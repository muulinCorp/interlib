package sensor

import (
	"time"
)

type UpsertData struct {
	Mac 	string
	DvID	string
	Data	map[uint16]struct{
		Value float64
		Time  time.Time
		DP  uint8
	} 
}