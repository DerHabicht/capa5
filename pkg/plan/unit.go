package plan

import (
	"github.com/ag7if/go-files"
)

type Unit struct {
	seal    files.File
	patch   files.File
	name    string
	address string
}

func NewUnit(
	seal files.File,
	patch files.File,
	name string,
	address string,

) Unit {
	return Unit{
		seal:    seal,
		patch:   patch,
		name:    name,
		address: address,
	}
}

func (unit Unit) Seal() files.File {
	return unit.seal
}

func (unit Unit) Patch() files.File {
	return unit.patch
}

func (unit Unit) Name() string {
	return unit.name
}

func (unit Unit) Address() string {
	return unit.address
}
