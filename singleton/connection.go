package singleton

import (
	"log"
	"sync"
)

type (
	Config any
)

// using mutex to lock the instance
// so instance will lock when it is creating
// and unlock when it is created
var lock = &sync.Mutex{}

// dummy connection struct
type connection struct {
	config *Config
}

var connectionInstance *connection

// get connection without singleton
// this will create new instance every time
func GetConnectionInstanceWithoutSingleton() connection {
	return connection{}
}

// get connection with singleton
// this will return same instance every time when instance is already created
func GetConnectionInstance() *connection {
	if connectionInstance == nil {
		lock.Lock()         // lock the instance when it is creating
		defer lock.Unlock() // unlock the instance
		log.Println("Connection Instance created")
		connectionInstance = &connection{}
	} else {
		log.Println("Connection Instance already created")
	}

	return connectionInstance
}
