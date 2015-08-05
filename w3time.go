package passbook

import "time"

type W3Time time.Time

func (t *W3Time) UnmarshalJSON(data []byte) error {
	format := "\"2006-01-02T15:04-07:00\""
	if ti, err := time.Parse(format, string(data)); err != nil {
		return err
	} else {
		t = (*W3Time)(&ti)
	}
	return nil
}

func (t W3Time) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).UTC().Format("2006-01-02T15:04-07:00")), nil
}
