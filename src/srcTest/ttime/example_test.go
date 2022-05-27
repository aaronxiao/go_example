package ttime


import (
	"fmt"
	"testing"
	"time"
)

func expensiveCall() {}

func TestExampleDuration(t *testing.T) {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0) )		//可统计运行时间

	fmt.Printf("The call took %v to run.\n", time.Since(t0) )		//这样也可以
}


func TestExampleDuration_Round(t *testing.T) {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
		fmt.Println("1111: ", d.Round(r).Nanoseconds() )
	}
	// Output:
	// d.Round(   1ns) = 1h15m30.918273645s
	// d.Round(   1µs) = 1h15m30.918274s
	// d.Round(   1ms) = 1h15m30.918s
	// d.Round(    1s) = 1h15m31s
	// d.Round(    2s) = 1h15m30s
	// d.Round(  1m0s) = 1h16m0s
	// d.Round( 10m0s) = 1h20m0s
	// d.Round(1h0m0s) = 1h0m0s
}


func TestExampleDuration_String(t *testing.T) {
	fmt.Println(1*time.Hour + 2*time.Minute + 300*time.Millisecond)
	fmt.Println(300 * time.Millisecond)
	// Output:
	// 1h2m0.3s			//是这种格式是因为Duration实现了String方法
	// 300ms
}


func ExampleDuration_Truncate() {
	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	trunc := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, t := range trunc {
		fmt.Printf("d.Truncate(%6s) = %s\n", t, d.Truncate(t).String())		//注意这里不是Round
	}
	// Output:
	// d.Truncate(   1ns) = 1h15m30.918273645s
	// d.Truncate(   1µs) = 1h15m30.918273s
	// d.Truncate(   1ms) = 1h15m30.918s
	// d.Truncate(    1s) = 1h15m30s
	// d.Truncate(    2s) = 1h15m30s
	// d.Truncate(  1m0s) = 1h15m0s
	// d.Truncate( 10m0s) = 1h10m0s
	// d.Truncate(1h0m0s) = 1h0m0s
}


func ExampleParseDuration() {
	hours, _ := time.ParseDuration("10h")
	complex, _ := time.ParseDuration("1h10m10s")
	micro, _ := time.ParseDuration("1µs")
	// The package also accepts the incorrect but common prefix u for micro.
	micro2, _ := time.ParseDuration("1us")

	fmt.Println(hours)
	fmt.Println(complex)
	fmt.Printf("There are %.0f seconds in %v.\n", complex.Seconds(), complex)
	fmt.Printf("There are %d nanoseconds in %v.\n", micro.Nanoseconds(), micro)
	fmt.Printf("There are %6.2e seconds in %v.\n", micro2.Seconds(), micro)
	// Output:
	// 10h0m0s
	// 1h10m10s
	// There are 4210 seconds in 1h10m10s.
	// There are 1000 nanoseconds in 1µs.
	// There are 1.00e-06 seconds in 1µs.
}


var c chan int

func handle(t int) {
	fmt.Println("1111: ", t)
}

func TestExampleAfter(t *testing.T) {
	tc := time.After(10 * time.Second)
	select {
	case m := <-c:
		handle(m)
	case <-tc:
	//case <- time.After(10 * time.Second):		//都可以
		fmt.Println("timed out")

	}
}

func statusUpdate() string { return "" }

func TestExampleTick(t *testing.T) {
	c := time.Tick(5 * time.Second)			//每5秒触发一次
	for next := range c {
		fmt.Printf("%v %s\n", next, statusUpdate())
	}
}

func ExampleDate() {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	// Output: Go launched at 2009-11-10 15:00:00 -0800 PST
}


func TestExampleNewTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)		//每秒触发一次
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			ticker.Stop()
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
			ticker.Reset(2* time.Second)
		}
	}
}

func TestExampleTime_GoString(t1 *testing.T) {
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(t.GoString())
	t = t.Add(1 * time.Minute)
	fmt.Println(t.GoString())
	t = t.AddDate(0, 1, 0)
	fmt.Println(t.GoString())
	t, _ = time.Parse("Jan 2, 2006 at 3:04pm (MST)", "Feb 3, 2013 at 7:54pm (UTC)")
	fmt.Println(t.GoString())

	// Output:
	// time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	// time.Date(2009, time.November, 10, 23, 1, 0, 0, time.UTC)
	// time.Date(2009, time.December, 10, 23, 1, 0, 0, time.UTC)
	// time.Date(2013, time.February, 3, 19, 54, 0, 0, time.UTC)
}


func TestExampleParse(t1 *testing.T) {
	// See the example for Time.Format for a thorough description of how
	// to define the layout string to parse a time.Time value; Parse and
	// Format use the same model to describe their input and output.

	// longForm shows by example how the reference time would be represented in
	// the desired layout.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)

	// Some valid layouts are invalid time values, due to format specifiers
	// such as _ for space padding and Z for zone information.
	// For example the RFC3339 layout 2006-01-02T15:04:05Z07:00
	// contains both Z and a time zone offset in order to handle both valid options:
	// 2006-01-02T15:04:05Z
	// 2006-01-02T15:04:05+07:00
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	fmt.Println(t)
	t, _ = time.Parse(time.RFC3339, "2006-01-02T15:04:05+07:00")
	fmt.Println(t)
	_, err := time.Parse(time.RFC3339, time.RFC3339)
	fmt.Println("error", err) // Returns an error as the layout is not a valid time value

	// Output:
	// 2013-02-03 19:54:00 -0800 PST
	// 2013-02-03 00:00:00 +0000 UTC
	// 2006-01-02 15:04:05 +0000 UTC
	// 2006-01-02 15:04:05 +0700 +0700
	// error parsing time "2006-01-02T15:04:05Z07:00": extra text: "07:00"
}


func TestExampleParseInLocation(t1 *testing.T) {
	loc, _ := time.LoadLocation("Europe/Berlin")

	// This will look for the name CEST in the Europe/Berlin time zone.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	// Note: without explicit zone, returns time in given location.
	const shortForm = "2006-Jan-02"
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

	// Output:
	// 2012-07-09 05:02:00 +0200 CEST
	// 2012-07-09 00:00:00 +0200 CEST
}


func TestExampleUnix(t1 *testing.T) {
	unixTime := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println(unixTime.Unix())			//秒数
	t := time.Unix(unixTime.Unix(), 0).UTC()		//UCT
	fmt.Println(t)


	fmt.Println(unixTime.UnixMicro())
	t = time.UnixMicro(unixTime.UnixMicro()).UTC()
	fmt.Println(t)

	fmt.Println(unixTime.UnixMilli())
	t = time.UnixMilli(unixTime.UnixMilli()).UTC()
	fmt.Println(t)


	// Output:
	// 1257894000
	// 2009-11-10 23:00:00 +0000 UTC
}


func ExampleTime_Round() {
	t := time.Date(0, 0, 0, 12, 15, 30, 918273645, time.UTC)
	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, d := range round {
		fmt.Printf("t.Round(%6s) = %s\n", d, t.Round(d).Format("15:04:05.999999999"))
		fmt.Printf("t.Truncate(%5s) = %s\n", d, t.Truncate(d).Format("15:04:05.999999999"))
	}
	// Output:
	// t.Round(   1ns) = 12:15:30.918273645
	// t.Round(   1µs) = 12:15:30.918274
	// t.Round(   1ms) = 12:15:30.918
	// t.Round(    1s) = 12:15:31
	// t.Round(    2s) = 12:15:30
	// t.Round(  1m0s) = 12:16:00
	// t.Round( 10m0s) = 12:20:00
	// t.Round(1h0m0s) = 12:00:00
}

func ExampleTime_Unix() {
	// 1 billion seconds of Unix, three ways.
	fmt.Println(time.Unix(1e9, 0).UTC())     // 1e9 seconds
	fmt.Println(time.Unix(0, 1e18).UTC())    // 1e18 nanoseconds
	fmt.Println(time.Unix(2e9, -1e18).UTC()) // 2e9 seconds - 1e18 nanoseconds

	t := time.Date(2001, time.September, 9, 1, 46, 40, 0, time.UTC)
	fmt.Println(t.Unix())     // seconds since 1970
	fmt.Println(t.UnixNano()) // nanoseconds since 1970

	// Output:
	// 2001-09-09 01:46:40 +0000 UTC
	// 2001-09-09 01:46:40 +0000 UTC
	// 2001-09-09 01:46:40 +0000 UTC
	// 1000000000
	// 1000000000000000000
}

func TestExampleLocation(t *testing.T) {
	// China doesn't have daylight saving. It uses a fixed 8 hour offset from UTC.
	secondsEastOfUTC := int((8 * time.Hour).Seconds())		//8小时的秒数
	fmt.Println(secondsEastOfUTC)

	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)

	// If the system has a timezone database present, it's possible to load a location
	// from that, e.g.:
	//    newYork, err := time.LoadLocation("America/New_York")

	// Creating a time requires a location. Common locations are time.Local and time.UTC.
	timeInUTC := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	sameTimeInBeijing := time.Date(2009, 1, 1, 20, 0, 0, 0, beijing)

	fmt.Println(timeInUTC.Unix(), sameTimeInBeijing.Unix())

	// Although the UTC clock time is 1200 and the Beijing clock time is 2000, Beijing is
	// 8 hours ahead so the two dates actually represent the same instant.
	timesAreEqual := timeInUTC.Equal(sameTimeInBeijing)		//判断相等
	fmt.Println(timesAreEqual)

	// Output:
	// true
}


func ExampleTime_Add() {
	start := time.Date(2009, 1, 1, 12, 0, 0, 0, time.UTC)
	afterTenSeconds := start.Add(time.Second * 10)
	afterTenMinutes := start.Add(time.Minute * 10)
	afterTenHours := start.Add(time.Hour * 10)
	afterTenDays := start.Add(time.Hour * 24 * 10)

	fmt.Printf("start = %v\n", start)
	fmt.Printf("start.Add(time.Second * 10) = %v\n", afterTenSeconds)
	fmt.Printf("start.Add(time.Minute * 10) = %v\n", afterTenMinutes)
	fmt.Printf("start.Add(time.Hour * 10) = %v\n", afterTenHours)
	fmt.Printf("start.Add(time.Hour * 24 * 10) = %v\n", afterTenDays)

	// Output:
	// start = 2009-01-01 12:00:00 +0000 UTC
	// start.Add(time.Second * 10) = 2009-01-01 12:00:10 +0000 UTC
	// start.Add(time.Minute * 10) = 2009-01-01 12:10:00 +0000 UTC
	// start.Add(time.Hour * 10) = 2009-01-01 22:00:00 +0000 UTC
	// start.Add(time.Hour * 24 * 10) = 2009-01-11 12:00:00 +0000 UTC
}

func ExampleTime_AddDate() {
	start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	oneDayLater := start.AddDate(0, 0, 1)
	oneMonthLater := start.AddDate(0, 1, 0)
	oneYearLater := start.AddDate(1, 0, 0)

	fmt.Printf("oneDayLater: start.AddDate(0, 0, 1) = %v\n", oneDayLater)
	fmt.Printf("oneMonthLater: start.AddDate(0, 1, 0) = %v\n", oneMonthLater)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)

	// Output:
	// oneDayLater: start.AddDate(0, 0, 1) = 2009-01-02 00:00:00 +0000 UTC
	// oneMonthLater: start.AddDate(0, 1, 0) = 2009-02-01 00:00:00 +0000 UTC
	// oneYearLater: start.AddDate(1, 0, 0) = 2010-01-01 00:00:00 +0000 UTC
}

func ExampleTime_After() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear3000AfterYear2000 := year3000.After(year2000) // True			//判断year2000在year3000之后
	isYear2000AfterYear3000 := year2000.After(year3000) // False

	fmt.Printf("year3000.After(year2000) = %v\n", isYear3000AfterYear2000)
	fmt.Printf("year2000.After(year3000) = %v\n", isYear2000AfterYear3000)

	// Output:
	// year3000.After(year2000) = true
	// year2000.After(year3000) = false
}

func ExampleTime_Before() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	isYear2000BeforeYear3000 := year2000.Before(year3000) // True  判断year2000在year3000之前
	isYear3000BeforeYear2000 := year3000.Before(year2000) // False

	fmt.Printf("year2000.Before(year3000) = %v\n", isYear2000BeforeYear3000)
	fmt.Printf("year3000.Before(year2000) = %v\n", isYear3000BeforeYear2000)

	// Output:
	// year2000.Before(year3000) = true
	// year3000.Before(year2000) = false
}

func TestExampleTime_Date(t *testing.T) {
	d := time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC)
	year, month, day := d.Date()
	hour, min, sec := d.Clock()


	fmt.Printf("year = %v\n", year)
	fmt.Printf("month = %v\n", month)
	fmt.Printf("day = %v\n", day)

	fmt.Printf("hour = %v\n", hour)
	fmt.Printf("min = %v\n", min)
	fmt.Printf("sec = %v\n", sec)

	// Output:
	// year = 2000
	// month = February
	// day = 1
	// hour = 12
	// min = 30
	// sec = 0
}


func TestExampleTime_Sub(t *testing.T) {
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	difference := end.Sub(start)
	fmt.Printf("difference = %v\n", int64(difference) )

	// Output:
	// difference = 12h0m0s
}


func TestExampleTime_AppendFormat(t1 *testing.T) {
	t := time.Date(2017, time.November, 4, 11, 0, 0, 0, time.UTC)
	text := []byte("Time: ")

	text = t.AppendFormat(text, time.Kitchen)
	fmt.Println(string(text))

	// Output:
	// Time: 11:00AM
}