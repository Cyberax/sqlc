// Code generated by sqlc. DO NOT EDIT.

package querytest

import ()

type Bar struct {
	ID   string
	Info []string
}

func (t *Bar) GetID() string {
	return t.ID
}
func (t *Bar) GetInfo() []string {
	return t.Info
}

type Foo struct {
	ID  string
	Bar string
}

func (t *Foo) GetID() string {
	return t.ID
}
func (t *Foo) GetBar() string {
	return t.Bar
}
