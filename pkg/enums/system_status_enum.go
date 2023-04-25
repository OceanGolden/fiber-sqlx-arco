package enums

import "strings"

type StatusType int

const (
	Disable StatusType = iota
	Enable
	Unknown
)

func (b StatusType) String() string {
	switch b {
	case Disable:
		return "Disable"
	case Enable:
		return "Enable"
	default:
		return "Unknown"
	}
}
func (b StatusType) MarshalText() ([]byte, error) {
	return []byte(b.String()), nil
}

func (b *StatusType) UnmarshalText(text []byte) error {
	*b = StatusFromText(string(text))
	return nil
}

func StatusFromText(text string) StatusType {
	switch strings.ToLower(text) {
	default:
		return Unknown
	case "Disable":
		return Disable
	case "Enable":
		return Enable
	}
}
