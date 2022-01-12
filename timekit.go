package timekit

import (
	"time"
)

// FirstDayOfLastYear returns first date (with 0:00 hour) from last year.
func FirstDayOfLastYear(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year()-1, 1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfThisYear returns the date (with 0:00 hour) from the first date of this year.
func FirstDayOfThisYear(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), 1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfNextYear returns date (12AM) of the first date of next year.
func FirstDayOfNextYear(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year()+1, 1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfLastMonth returns the date (with 0:00 hour) of the first day from last month.
func FirstDayOfLastMonth(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month()-1, 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfThisMonth returns the first date (with 0:00 hour) from this month.
func FirstDayOfThisMonth(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, dt.Location())
}

// FirstDayOfNextMonth returns next months first day (in 12 AM hours).
func FirstDayOfNextMonth(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month()+1, 1, 0, 0, 0, 0, dt.Location())
}

// MidnightYesterday return 12 AM date of yesterday.
func MidnightYesterday(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day()-1, 0, 0, 0, 0, dt.Location())
}

// Midnight return today's date at 12 o’clock (or 0:00) during the night.
func Midnight(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, dt.Location())
}

// MidnightTomorrow will return tomorrows date at 12 o’clock (or 0:00) during the night.
func MidnightTomorrow(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, dt.Location())
}

// Noon will return today's date at 12 o'clock (or 12:00) during the day.
func Noon(now func() time.Time) time.Time {
	dt := now()
	return time.Date(dt.Year(), dt.Month(), dt.Day(), 12, 0, 0, 0, dt.Location())
}

// FirstDayOfLastISOWeek returns the previous week's monday date.
func FirstDayOfLastISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate back to Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}
	dt = dt.AddDate(0, 0, -1) // Skip the current monday we are on!

	// iterate ONCE AGAIN back to Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}

	return dt
}

// FirstDayOfThisISOWeek return monday's date of this week. Please note monday is considered the first day of the week according to ISO 8601 and not sunday (which is what is used in Canada and USA).
func FirstDayOfThisISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate back to Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, -1)
	}

	return dt
}

// LastDayOfThisISOWeek return sunday's date of this week. Please note sunday is considered the last day of the week according to ISO 8601.
func LastDayOfThisISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate forward to Sunday
	for dt.Weekday() != time.Sunday {
		dt = dt.AddDate(0, 0, 1)
	}

	return dt
}

// FirstDayOfNextISOWeek return date of the upcoming monday.
func FirstDayOfNextISOWeek(now func() time.Time) time.Time {
	dt := now()

	// iterate forward to next Monday
	for dt.Weekday() != time.Monday {
		dt = dt.AddDate(0, 0, 1)
	}

	return dt
}

// Range function returns an array of datetime values from the starting date to the finish date according to the step pattern specified in the parameter.
func Range(start time.Time, end time.Time, yearStep int, monthStep int, dayStep int, hourStep int, minuteStep int, secondStep int) []time.Time {
	var results []time.Time

	dur := time.Hour*time.Duration(hourStep) + time.Minute*time.Duration(minuteStep) + time.Second*time.Duration(secondStep)
	for d := start; d.After(end) == false; d = d.AddDate(yearStep, monthStep, dayStep).Add(dur) {
		results = append(results, d)
	}
	return results
}

// TimeStepper is a structure to hold keep track of the position we are in the datetime range which we are stepping through.
type TimeStepper struct {
	// Details hidden to keep library simple.

	curr       time.Time
	start      time.Time
	end        time.Time
	yearStep   int
	monthStep  int
	dayStep    int
	hourStep   int
	minuteStep int
	secondStep int
}

// NewTimeStepper is a constructor of the `TimeStepper` struct.
func NewTimeStepper(start time.Time, end time.Time, yearStep int, monthStep int, dayStep int, hourStep int, minuteStep int, secondStep int) *TimeStepper {
	return &TimeStepper{
		curr:       start,
		start:      start,
		end:        end,
		yearStep:   yearStep,
		monthStep:  monthStep,
		dayStep:    dayStep,
		hourStep:   hourStep,
		minuteStep: minuteStep,
		secondStep: secondStep,
	}
}

// Step makes one time step over and returns true or false depending if the stepper has stepped over the end datetime.
func (ts *TimeStepper) Step() bool {
	dur := time.Hour*time.Duration(ts.hourStep) + time.Minute*time.Duration(ts.minuteStep) + time.Second*time.Duration(ts.secondStep)
	ts.curr = ts.curr.AddDate(ts.yearStep, ts.monthStep, ts.dayStep).Add(dur)
	return ts.curr.After(ts.end) == false
}

// Done checks to see if the stepper has stepped over the end datetime and will return true or false according.
func (ts *TimeStepper) Done() bool {
	return ts.curr.After(ts.end)
}

// Value will return the value that that the stepper is currently on.
func (ts *TimeStepper) Value() time.Time {
	return ts.curr
}
