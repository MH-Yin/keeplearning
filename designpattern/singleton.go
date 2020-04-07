package designpattern

import (
	"sync"
)

type instance struct {
}

var singleton *instance
var once sync.Once

func getInstance() *instance {
	once.Do(func() {
		singleton = &instance{}
	})
	return singleton
}
