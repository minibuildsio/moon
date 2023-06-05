package main

import (
	"testing"
	"time"
)

type TestCase struct {
	input    time.Time
	expected LunarPhase
}

func TestLunarPhaseOn(t *testing.T) {
	testCases := [10]TestCase{
		{time.Date(2000, 1, 1, 20, 0, 0, 0, time.UTC), WANING_CRESCENT},
		{time.Date(2023, 6, 5, 20, 0, 0, 0, time.UTC), WANING_GIBBOUS},
		{time.Date(2023, 6, 10, 20, 0, 0, 0, time.UTC), LAST_QUARTER},
		{time.Date(2023, 6, 12, 20, 0, 0, 0, time.UTC), WANING_CRESCENT},
		{time.Date(2023, 6, 17, 20, 0, 0, 0, time.UTC), NEW},
		{time.Date(2023, 6, 19, 20, 0, 0, 0, time.UTC), WAXING_CRESCENT},
		{time.Date(2023, 6, 25, 20, 0, 0, 0, time.UTC), FIRST_QUARTER},
		{time.Date(2023, 6, 27, 20, 0, 0, 0, time.UTC), WAXING_GIBBOUS},
		{time.Date(2023, 7, 3, 18, 0, 0, 0, time.UTC), FULL},
		{time.Date(2050, 1, 1, 20, 0, 0, 0, time.UTC), WAXING_GIBBOUS},
	}

	for _, testCase := range testCases {
		phase := LunarPhaseOn(testCase.input)

		if phase != testCase.expected {
			t.Errorf("Lunar phase on %s was %s not %s", testCase.input, testCase.expected, phase)
		}
	}
}
