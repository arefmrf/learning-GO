package domain

import "time"

type Timestamp int64

func Now() Timestamp {
	return Timestamp(time.Now().Unix())
}

func NormalizeDate(ts Timestamp) Timestamp {
	t := time.Unix(int64(ts), 0).
		Truncate(24 * time.Hour).
		Add(12 * time.Hour)
	return Timestamp(t.Unix())
}
