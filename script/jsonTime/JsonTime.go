package jsonTime
import "time"

type JsonTime struct {
	time.Time
}

func (jsonTime JsonTime) format(formatStr string) string {
	return jsonTime.Time.Format(formatStr)
}

func (jsonTime JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + jsonTime.format("2006-01-02") + `"`), nil
}
