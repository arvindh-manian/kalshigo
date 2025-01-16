package structs

import (
	"encoding/json"
	"time"
)

type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return e.Body
}

type Timestamp struct {
	time.Time
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)

	if err != nil {
		return err
	}

	t.Time = time.Unix(timestamp, 0)

	return nil
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

type OptionalTime struct {
	time.Time
}

func (t *OptionalTime) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)

	if err != nil {
		return err
	}

	if s == "" {
		return nil
	}

	v, err := time.Parse(time.RFC3339, s)

	if err != nil {
		return err
	}

	t.Time = v

	return nil
}

func (t OptionalTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}

	return json.Marshal(t.Format(time.RFC3339))
}
