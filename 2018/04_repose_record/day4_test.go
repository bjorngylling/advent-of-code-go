package main

import (
	"log"
	"reflect"
	"testing"
	"time"
)

func parseTime(s string) (t time.Time) {
	t, err := time.Parse("2006-01-02 15:04", s)
	if err != nil {
		log.Fatal(err)
	}
	return
}

var data = []string{
	"[1518-11-01 00:25] wakes up",
	"[1518-11-01 00:00] Guard #10 begins shift",
	"[1518-11-02 00:00] Guard #12 begins shift",
	"[1518-11-01 00:05] falls asleep",
	"[1518-11-02 00:30] wakes up",
	"[1518-11-02 00:05] falls asleep",
}

func TestParseLog(t *testing.T) {
	expected := EventLog{
		10: {
			{parseTime("1518-11-01 00:05"), SLEEP},
			{parseTime("1518-11-01 00:25"), WAKEUP},
		},
		12: {
			{parseTime("1518-11-02 00:05"), SLEEP},
			{parseTime("1518-11-02 00:30"), WAKEUP},
		},
	}
	result := parseLog(data)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}

func TestSleepiestGuard(t *testing.T) {
	expected := 12
	result := sleepiestGuard(parseLog(data))
	if result != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}

func TestMostCommonSleepingMinute(t *testing.T) {
	events := []Event{
		{parseTime("1518-11-02 00:05"), SLEEP},
		{parseTime("1518-11-02 00:30"), WAKEUP},
		{parseTime("1518-11-03 00:25"), SLEEP},
		{parseTime("1518-11-03 00:40"), WAKEUP},
		{parseTime("1518-11-04 00:34"), SLEEP},
		{parseTime("1518-11-04 00:46"), WAKEUP},
		{parseTime("1518-11-05 00:19"), SLEEP},
		{parseTime("1518-11-05 00:26"), WAKEUP},
	}
	expected := 25
	result := mostCommonSleepingMinute(events)
	if result.min != expected {
		t.Errorf("Expected %+v but was %+v", expected, result)
	}
}
