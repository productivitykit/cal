// Copyright (c) Kyle Huggins and contributors
// SPDX-License-Identifier: BSD-3-Clause

package cal

import (
	"bytes"
	"fmt"
	"time"
)

// Config defines options for rendering.
type Config struct {
	Year         int          // e.g. 2023
	Month        time.Month   // e.g. time.January
	WeekNumbers  bool         // include ISO week numbers column
	StartWeekday time.Weekday // first day of week e.g. time.Sunday, time.Monday
}

// RenderMonth returns the month calendar as a string.
func RenderMonth(cfg Config) string {
	var buf bytes.Buffer

	bodyWidth := 21
	if cfg.WeekNumbers {
		bodyWidth += 3
	}

	monthName := cfg.Month.String()
	header := fmt.Sprintf("%s %d", monthName, cfg.Year)
	headerPadding := max((bodyWidth-len(header))/2, 0)

	buf.WriteString(fmt.Sprintf("%*s%s\n", headerPadding, "", header))

	if cfg.WeekNumbers {
		buf.WriteString("Wk ")
	}

	// Write day headers based on start weekday
	weekdayNames := []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
	for i := range 7 {
		dayIndex := (int(cfg.StartWeekday) + i) % 7
		buf.WriteString(weekdayNames[dayIndex])
		if i < 6 {
			buf.WriteString(" ")
		}
	}
	buf.WriteString("\n")

	firstOfMonth := time.Date(cfg.Year, cfg.Month, 1, 0, 0, 0, 0, time.UTC)

	// Calculate offset from the configured start weekday
	firstDayOfMonth := int(firstOfMonth.Weekday())
	startWeekdayInt := int(cfg.StartWeekday)
	offset := (firstDayOfMonth - startWeekdayInt + 7) % 7

	day := 1
	daysInMonth := daysIn(cfg.Year, cfg.Month)

	// Print week rows
	for row := 0; day <= daysInMonth; row++ {
		// Optional week number
		if cfg.WeekNumbers {
			_, week := firstOfMonth.AddDate(0, 0, day-1).ISOWeek()
			buf.WriteString(fmt.Sprintf("%2d ", week))
		}

		for wd := range 7 {
			if row == 0 && wd < offset {
				buf.WriteString("  ")
			} else if day > daysInMonth {
				buf.WriteString("  ")
			} else {
				buf.WriteString(fmt.Sprintf("%2d", day))
				day++
			}
			if wd < 6 {
				buf.WriteString(" ")
			}
		}

		buf.WriteString("\n")
	}

	return buf.String()
}

func daysIn(year int, month time.Month) int {
	firstOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	firstOfNextMonth := firstOfMonth.AddDate(0, 1, 0)
	return int(firstOfNextMonth.Sub(firstOfMonth).Hours() / 24)
}
