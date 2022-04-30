package main

import (
	"log"
	"sync"
)

var (
	lock           = &sync.Mutex{}
	global *Global = nil
)

//GetInstanceWithMutex goroutine safely, return Global instance
func GetInstanceWithMutex() *Global {
	lock.Lock()
	defer lock.Unlock()
	if global != nil {
		log.Println("Global instance already created")
		return global
	}
	log.Println("Create new Global instance")
	global = new(Global)
	return global
}
