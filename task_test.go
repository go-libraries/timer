package timer

import (
    "testing"
    "time"
    "fmt"
)

func TestTime(t *testing.T) {
    startTime := "2019-08-11 21:30:01"
    endTime := "2019-08-11 21:45:00"
    cal := "2019-08-12 21:30:07"

    t1, err := time.ParseInLocation("2006-01-02 15:04:05", cal, time.Local)
    s, err1 := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
    e, err2 := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)

    sTime := time.Now().Format("2006-01-02") + " " + s.Format("15:04:05")
    eTime := time.Now().Format("2006-01-02") + " " + e.Format("15:04:05")
    cal = time.Now().Format("2006-01-02") + " " + t1.Format("15:04:05")

    fmt.Println(t1, err, err1, err2, sTime, eTime)
    if cal >= sTime && cal <= eTime {
        fmt.Println(cal )
    }

}