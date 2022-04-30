package main

import "log"

func main() {
	log.Println(`Get Global instance with sync.Mutex`)
	_ = GetInstanceWithMutex()
	_ = GetInstanceWithMutex()
	log.Println(`Get Global instance with sync.Once`)
	_ = GetInstanceWithOnce()
	_ = GetInstanceWithOnce()
}
