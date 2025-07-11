package user

type Role int

const (
	Owner Role = iota
	Commander
	Planner
)
