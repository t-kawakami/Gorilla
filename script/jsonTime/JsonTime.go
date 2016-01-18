package jsonTime
import "time"

type JsonTime struct {
	time.Time
}

func (jsonTime JsonTime) format() string {
	return jsonTime.Time.Format("2006-01-02")
}

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + jsonTime.format() + `"`), nil
}