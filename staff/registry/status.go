package registry

import (
	"encoding/json"
	"gitlab.com/andile/go/popcorn/log"
)

type Status int

const (
	Invalid Status = iota
	Awol
	Sick
	HalfDay
	FullDay
)

func (a Status) String() string {
	switch a {
	default:
		return "Invalid"
	case Awol:
		return "Awol"
	case Sick:
		return "Sick"
	case HalfDay:
		return "HalfDay"
	case FullDay:
		return "FullDay"
	}
}

func (a *Status) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	default:
		log.Error("Invalid status: " + s)
		*a = Invalid
	case Awol.String():
		*a = Awol
	case Sick.String():
		*a = Sick
	case HalfDay.String():
		*a = HalfDay
	case FullDay.String():
		*a = FullDay
	}

	return nil
}

func (a Status) MarshalJSON() ([]byte, error) {
	var s string
	switch a {
	default:
		s = Invalid.String()
	case Awol:
		s = Awol.String()
	case Sick:
		s = Sick.String()
	case HalfDay:
		s = HalfDay.String()
	case FullDay:
		s = FullDay.String()
	}

	return json.Marshal(s)
}
