package main

import (
	"log"
	"sync"
)

var (
	once               = &sync.Once{}
	globalOnce *Global = nil
)

//GetInstanceWithOnce goroutine safely, return Global instance
func GetInstanceWithOnce() *Global {
	if globalOnce != nil {
		log.Println("Global instance already created")
		return globalOnce
	}
	once.Do(func() {
		log.Println("Create new Global instance")
		globalOnce = new(Global)
	})
	return globalOnce
}
