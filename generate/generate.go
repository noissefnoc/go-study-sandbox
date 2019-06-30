package generate

//go:generate stringer -type=Status
type Status int

const (
	Active Status = iota
	Inactive
)
