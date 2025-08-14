// Copyright (c) Kyle Huggins and contributors
// SPDX-License-Identifier: BSD-3-Clause

package cal

import (
	"strings"
	"testing"
	"time"
)

func TestRenderMonthOk(t *testing.T) {
	tests := []struct {
		name   string
		config Config
		want   []string
	}{
		{
			name: "January 2023 without week numbers",
			config: Config{
				Year:         2023,
				Month:        time.January,
				WeekNumbers:  false,
				StartWeekday: time.Monday,
			},
			want: []string{
				"    January 2023",
				"Mo Tu We Th Fr Sa Su",
				"                   1",
				" 2  3  4  5  6  7  8",
				" 9 10 11 12 13 14 15",
				"16 17 18 19 20 21 22",
				"23 24 25 26 27 28 29",
				"30 31               ",
			},
		},
		{
			name: "February 2024 (leap year) without week numbers",
			config: Config{
				Year:         2024,
				Month:        time.February,
				WeekNumbers:  false,
				StartWeekday: time.Monday,
			},
			want: []string{
				"    February 2024",
				"Mo Tu We Th Fr Sa Su",
				"          1  2  3  4",
				" 5  6  7  8  9 10 11",
				"12 13 14 15 16 17 18",
				"19 20 21 22 23 24 25",
				"26 27 28 29         ",
			},
		},
		{
			name: "January 2023 starting Sunday",
			config: Config{
				Year:         2023,
				Month:        time.January,
				WeekNumbers:  false,
				StartWeekday: time.Sunday,
			},
			want: []string{
				"    January 2023",
				"Su Mo Tu We Th Fr Sa",
				" 1  2  3  4  5  6  7",
				" 8  9 10 11 12 13 14",
				"15 16 17 18 19 20 21",
				"22 23 24 25 26 27 28",
				"29 30 31            ",
			},
		},
		{
			name: "January 2023 starting Wednesday",
			config: Config{
				Year:         2023,
				Month:        time.January,
				WeekNumbers:  false,
				StartWeekday: time.Wednesday,
			},
			want: []string{
				"    January 2023",
				"We Th Fr Sa Su Mo Tu",
				"             1  2  3",
				" 4  5  6  7  8  9 10",
				"11 12 13 14 15 16 17",
				"18 19 20 21 22 23 24",
				"25 26 27 28 29 30 31",
			},
		},
		{
			name: "January 2023 starting Sunday with week numbers",
			config: Config{
				Year:         2023,
				Month:        time.January,
				WeekNumbers:  true,
				StartWeekday: time.Sunday,
			},
			want: []string{
				"      January 2023",
				"Wk Su Mo Tu We Th Fr Sa",
				"52  1  2  3  4  5  6  7",
				" 1  8  9 10 11 12 13 14",
				" 2 15 16 17 18 19 20 21",
				" 3 22 23 24 25 26 27 28",
				" 4 29 30 31            ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RenderMonth(tt.config)
			gotLines := strings.Split(strings.TrimRight(got, "\n"), "\n")

			if len(gotLines) != len(tt.want) {
				t.Errorf("Expected %d lines, got %d lines", len(tt.want), len(gotLines))
				t.Errorf("Got:\n%s", got)
				return
			}

			for i, line := range gotLines {
				if line != tt.want[i] {
					t.Errorf("Line %d: expected %q, got %q", i, tt.want[i], line)
				}
			}
		})
	}
}
