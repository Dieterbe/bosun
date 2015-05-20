package btime

import (
	"strconv"
	"time"

	"bosun.org/opentsdb"
)

// getTime takes in a time specification, which can be a unix timestamp or
// a duration offset which will be applied against now.
// the latter case can be expressed as in time.ParseDuration
// or using units specified at http://opentsdb.net/docs/build/html/user_guide/query/dates.html
func Get(now time.Time, spec string) (time.Time, error) {
	n, err := strconv.ParseUint(spec, 10, 0)
	if err == nil {
		return time.Unix(int64(n), 0), nil
	}
	duration, err := opentsdb.ParseDuration(spec)
	if err != nil {
		return time.Time{}, err
	}
	return now.Add(-time.Duration(duration)), nil
}
