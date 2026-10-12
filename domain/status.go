package domain

import "fmt"

type Status int

const (
	StatusInactive Status = 0
	StatusActive   Status = 1 // começa em 1, evita que 0 signifique algo
)

func ParseStatus(s string) (Status, error) {
	switch s {
	case "active":
		return StatusActive, nil
	case "inactive":
		return StatusInactive, nil
	default:
		return 0, fmt.Errorf("status inválido: %q", s)
	}
}
