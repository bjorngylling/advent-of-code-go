package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type EventType string

const (
	SLEEP  = "falls asleep"
	WAKEUP = "wakes up"
)

type Event struct {
	T  time.Time
	Ev EventType
}
type EventLog map[int][]Event

var logLineRegexp = regexp.MustCompile(`\[(.+)] (.+)`)
var shiftStartRegexp = regexp.MustCompile(`Guard #(\d+) begins shift`)

func parseLog(lines []string) EventLog {
	sort.Strings(lines)
	var id int
	events := make(EventLog)
	for _, ln := range lines {
		matches := logLineRegexp.FindAllStringSubmatch(ln, -1)
		t, err := time.Parse("2006-01-02 15:04", matches[0][1])
		if err != nil {
			log.Fatal(err)
		}
		logMsg := strings.Trim(matches[0][2], " ")
		if idMatch := shiftStartRegexp.FindAllStringSubmatch(logMsg, -1); len(idMatch) == 1 {
			id, _ = strconv.Atoi(idMatch[0][1])
		} else {
			events[id] = append(events[id], Event{t, EventType(logMsg)})
		}

	}
	return events
}

type KeyValue struct {
	Key, Value int
}

func sleepiestGuard(eventLog EventLog) int {
	minsAsleep := make(map[int]int)
	for id, events := range eventLog {
		for idx := 0; idx < len(events); idx += 2 {
			sleep, wake := events[idx], events[idx+1]
			minsAsleep[id] += int(wake.T.Sub(sleep.T).Minutes())
		}
	}
	var keyValues []KeyValue
	for k, v := range minsAsleep {
		keyValues = append(keyValues, KeyValue{k, v})
	}
	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].Value > keyValues[j].Value
	})
	return keyValues[0].Key
}

type minuteSleepCount struct {
	min        int
	sleepCount int
}

func mostCommonSleepingMinute(events []Event) minuteSleepCount {
	sleepPerMinute := make([]int, 60)
	for idx := 0; idx < len(events); idx += 2 {
		sleep, wake := events[idx], events[idx+1]
		for i := sleep.T.Minute(); i < wake.T.Minute(); i++ {
			sleepPerMinute[i]++
		}
	}
	m := 0
	minIdx := 0
	for i, e := range sleepPerMinute {
		if e > m {
			m = e
			minIdx = i
		}
	}
	return minuteSleepCount{minIdx, sleepPerMinute[minIdx]}
}

func guardMostFrequentAsleepMinute(eventLog EventLog) (int, int) {
	guardSleepData := make(map[int]minuteSleepCount)
	for id, events := range eventLog {
		guardSleepData[id] = mostCommonSleepingMinute(events)
	}
	mId := -1
	for id, sleepData := range guardSleepData {
		if mId == -1 || sleepData.sleepCount > guardSleepData[mId].sleepCount {
			mId = id
		}
	}
	return mId, guardSleepData[mId].min
}

func main() {
	fileContent, err := ioutil.ReadFile("04_repose_record/day4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	events := parseLog(strings.Split(string(fileContent), "\n"))

	id := sleepiestGuard(events)
	min := mostCommonSleepingMinute(events[id]).min
	fmt.Printf("Day 4 part 1 result: id=%d minute=%d answer=%d\n", id, min, id*min)

	id, min = guardMostFrequentAsleepMinute(events)
	fmt.Printf("Day 4 part 2 result: id=%d minute=%d answer=%d\n", id, min, id*min)
}
