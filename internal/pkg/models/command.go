package models

type Command struct {
	Name  string
	Usage string
	Flags []Flag
}
