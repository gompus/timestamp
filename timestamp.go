package timestamp

import (
	"encoding/json"
	"time"
)

const layout = time.RFC3339

// Timestamp is an ISO8601 timestamp.
type Timestamp time.Time

// Parse constructs a Timestamp from s.
func Parse(s string) (Timestamp, error) {
	parsed, err := time.Parse(layout, s)
	return Timestamp(parsed), err
}

// MustParse is like Parse, but panics if it encounters an error.
func MustParse(s string) Timestamp {
	t, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return t
}

// MarshalJSON returns the JSON representation of t.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	s := time.Time(t).Format(layout)
	return json.Marshal(s)
}

// UnmarshalJSON unmarshals data into t.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	dec, err := Parse(raw)
	if err == nil {
		*t = dec
	}
	return err
}
