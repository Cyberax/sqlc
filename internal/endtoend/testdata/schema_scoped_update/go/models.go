// Code generated by sqlc. DO NOT EDIT.

package querytest

import ()

type FooBar struct {
	ID   int32
	Name string
}

func (t *FooBar) GetID() int32 {
	return t.ID
}
func (t *FooBar) GetName() string {
	return t.Name
}
