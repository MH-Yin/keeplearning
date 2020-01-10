package golang

import (
	"fmt"
	"testing"
)

type Actor interface {
	baseActor
}
type context map[string]chan string

type baseActor interface {
	Act()
	GetName() string
	Receive(s string)
}

type abstractActor struct {
	name string
	mp   context
}

func (a *abstractActor) Act() {
	fmt.Println("I am abstractActor")
}

func (a *abstractActor) GetName() string {
	return a.name
}

func (a *abstractActor) Receive(s string) {
	a.mp[s] <- s
}

func NewActor(name string) Actor {
	return &abstractActor{name: name, mp: make(map[string]chan string)}
}

type heartManager struct {
	Actor
}

func TestNewActor(t *testing.T) {
	a := heartManager{NewActor("heartManager")}
	a.Act()
	fmt.Println(a.GetName())
}

func (r *heartManager) Act() {
	fmt.Println("I am heartManager")
}
