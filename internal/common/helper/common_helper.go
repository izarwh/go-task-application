package helper

import (
	"time"

	"github.com/google/uuid"
)

func StringTimeToTime(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	t, _ := time.Parse("2006-01-02", dateStr)
	return &t
}

func StringUUIDToUUID(uuidStr string) *uuid.UUID {
	if uuidStr == "" {
		return nil
	}
	u, _ := uuid.Parse(uuidStr)
	return &u
}
