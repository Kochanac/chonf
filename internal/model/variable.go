package model

type Type uint64

const (
	EnumType Type = iota
	StringType
	IntType
	BoolType
	ListType
	ChoiceListType
)

type RegistryEntry interface {
	Path() string
	Type() Type
}
