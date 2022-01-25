package timekit

import (
	"time"
)

// ParseJavaScriptTime will convert the number of milliseconds since the Unix Epoch parameter into Golang `time` format. As a result, the output of the JavaScript `getTime()` function can be used as the parameter in this function.
func ParseJavaScriptTime(i int64) time.Time {
	// Special thanks JavaScript timestamp to golang time.Time
	// https://gist.github.com/alextanhongpin/3b6b2ee47665ac9c1c32c805b86380a6
	return time.Unix(i/1000, (i%1000)*1000*1000)
}
