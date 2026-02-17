package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC,
	)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Nanoseconds())
	p(diff.Seconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	// epoch

	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())
	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, now.UnixNano()))

	// time parsing

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%2dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

	// math random number
	p("============ math random number ===============")
	p(rand.IntN(100), ",")
	p(rand.IntN(100))
	p(rand.Float64())
	p((rand.Float64()*5)+5, ",")
	p((rand.Float64() * 5) + 5)

	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)
	p(r2.IntN(100), ",")
	p(r2.IntN(100))
	p()
	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	p(r3.IntN(100), ",")
	p(r3.IntN(100))
	p()

}
