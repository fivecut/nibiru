package types

import (
	"fmt"
	"strings"
)

const (
	// WeekEpochID defines the identifier for weekly epochs
	MonthEpochID = "month"
	// WeekEpochID defines the identifier for weekly epochs
	WeekEpochID = "week"
	// DayEpochID defines the identifier for daily epochs
	DayEpochID = "day"
	// HourEpochID defines the identifier for hourly epochs
	HourEpochID = "hour"
	// ThirtyMinuteEpochID defines the identifier for 30 minute epochs
	ThirtyMinuteEpochID = "30 min"
	// FifteenMinuteEpochID defines the identifier for 15 minute epochs
	FifteenMinuteEpochID = "15 min"
)

// ValidateEpochIdentifierInterface performs a stateless
// validation of the epoch ID interface.
func ValidateEpochIdentifierInterface(i any) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return ValidateEpochIdentifierString(v)
}

// ValidateEpochIdentifierString performs a stateless
// validation of the epoch ID.
func ValidateEpochIdentifierString(s string) error {
	s = strings.TrimSpace(s)
	if s == "" {
		return fmt.Errorf("blank epoch identifier: %s", s)
	}
	return nil
}
