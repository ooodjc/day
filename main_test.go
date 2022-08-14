package day

import (
	"fmt"
	"testing"
	"time"
)

func TestDay(t *testing.T) {
	var d Day
	d.SetTime(time.Now()).SetFormat("2006-01-02 15:04:05")
	fmt.Printf("TimeToString:%s\n", d.TimeToString())
}
