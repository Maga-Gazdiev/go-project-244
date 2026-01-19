package model

const (
	StatusAdded    = "added"
	StatusRemoved  = "remove"
	StatusUnchanged = "unchanged"
	StatusChanged  = "changed"
	StatusNested   = "nested"
)

type DiffNode struct {
	Key      string
	Status   string
	OldValue any
	NewValue any
	Children []DiffNode
}





