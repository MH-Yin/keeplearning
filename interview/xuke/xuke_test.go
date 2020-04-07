package xuke

import (
	"fmt"
	"testing"
)

func TestArrayToString(t *testing.T) {
	fmt.Println(ArrayToString([]string{"alice", "bob", "jack", "ross"}))
	fmt.Println(ArrayToString([]string{"alice"}))
	fmt.Println(ArrayToString([]string{"alice", "bob"}))
}

func TestArrayToStringWithLimit(t *testing.T) {
	fmt.Println(ArrayToStringWithLimit([]string{"alice", "bob", "jack", "ross"},0))
}