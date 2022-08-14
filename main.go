package day

import (
	"time"
)

type Day struct {
	Time   time.Time
	Format string
}

func (d Day) New(t time.Time) Day {
	d = Day{
		Time:   t,
		Format: "2006-01-02 15:04:05",
	}
	return d
}

//set time time.Time
func (d Day) SetTime(t time.Time) Day {
	d.Time = t
	return d
}

//set time format
func (d Day) SetFormat(format string) Day {
	d.Format = format
	return d
}

//set string time to time.Time
func (d Day) SetStringTime(s string) Day {
	d.Time = d.StringToTime(s)
	return d
}

//time.Time to string time
func (d Day) TimeToString() string {
	return d.Time.Format(d.Format)
}

//time.Time to int time
func (d Day) TimeToInt() int {
	return int(d.Time.Unix())
}

//string time to time.Time
func (d Day) StringToTime(s string) time.Time {
	t, _ := time.Parse(d.Format, s)
	return t
}

//string time to int time
func (d Day) StringToInt(s string) int {
	t, _ := time.Parse(d.Format, s)
	return int(t.Unix())
}

//int time to time.Time
func (d Day) IntToTime(i int) time.Time {
	return time.Unix(int64(i), 0)
}

//int time to string time
func (d Day) IntToString(i int) string {
	return d.IntToTime(i).Format(d.Format)
}
