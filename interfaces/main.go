package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) (string, int) {
	s, ok := msg.(birthdayMessage)
	if ok {
		return s.getMessage(), 1111
	}
	n := msg.getMessage()
	cost := len(n) / 3
	return n, cost
}

type message interface {
	getMessage() string
}

type report interface {
	getReport() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func (sr sendingReport) getReport() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func setReport(rp report) string {
	return rp.getReport()
}

func main() {
	birthdayMsg := birthdayMessage{time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC), "John Doe"}
	// reportMsg := sendingReport{"test", 12}
	s, v := sendMessage(birthdayMsg)
	fmt.Println(v, s)
}
