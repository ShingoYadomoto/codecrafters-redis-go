package store

import (
	"sync"
)

var store = &sync.Map{}

func GetStore() *sync.Map {
	return store
}
