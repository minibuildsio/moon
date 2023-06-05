package main

import (
	"fmt"
	"math"
	"time"
)

type LunarPhase string

const (
	NEW             LunarPhase = "New 🌑"
	WAXING_CRESCENT LunarPhase = "Waxing Crescent 🌒"
	FIRST_QUARTER   LunarPhase = "First Quarter 🌓"
	WAXING_GIBBOUS  LunarPhase = "Waxing Gibbous 🌔"
	FULL            LunarPhase = "Full 🌕"
	WANING_GIBBOUS  LunarPhase = "Waning Gibbous 🌖"
	LAST_QUARTER    LunarPhase = "Last Quarter 🌗"
	WANING_CRESCENT LunarPhase = "Waning Crescent 🌘"
)

type Interval struct {
	start float32
	end   float32
}

func (interval Interval) Contains(v float32) bool {
	return v >= interval.start && v < interval.end
}

type LunarPhaseInterval struct {
	interval Interval
	phase    LunarPhase
}

var KnownNewMoon = time.Date(2023, 5, 19, 15, 53, 0, 0, time.UTC)

const LunarMonth = 29.531

var LunarPhaseLookup = [9]LunarPhaseInterval{
	{Interval{0, 1}, NEW},
	{Interval{1, 6.383}, WAXING_CRESCENT},
	{Interval{6.383, 8.382}, FIRST_QUARTER},
	{Interval{8.382, 13.765}, WAXING_GIBBOUS},
	{Interval{13.765, 15.765}, FULL},
	{Interval{15.765, 21.148}, WANING_GIBBOUS},
	{Interval{21.148, 23.148}, LAST_QUARTER},
	{Interval{23.148, 28.531}, WANING_CRESCENT},
	{Interval{28.531, 29.531}, NEW},
}

func LunarPhaseOn(date time.Time) LunarPhase {
	daysSinceNewMoon := math.Mod(date.Sub(KnownNewMoon).Hours()/24, LunarMonth)
	if daysSinceNewMoon < 0 {
		daysSinceNewMoon += LunarMonth
	}

	for _, entry := range LunarPhaseLookup {
		if entry.interval.Contains(float32(daysSinceNewMoon)) {
			return entry.phase
		}
	}

	return NEW
}

func main() {
	fmt.Printf("Current lunar phase: %s", LunarPhaseOn(time.Now()))
}
