package must

import "time"

func ParseDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err.Error())
	}
	return t
}
