package safemap

import "sync"

var myPool Pool

func Set(id string) error {
	return myPool.Set(id)
}

func IsExist(id string) bool {
	return myPool.IsExist(id)
}

func Delete(id string) error {
	return myPool.Delete(id)
}

type value struct{}

type Pool interface {
	Set(id string) error
	IsExist(id string) bool
	Delete(id string) error
}

type pool struct {
	rw          sync.RWMutex
	connections map[string]struct{}
}

func InitDefaultPool() {
	myPool = &pool{
		connections: make(map[string]struct{}),
	}
}

func (c *pool) Set(id string) error {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.connections[id] = value{}
	return nil
}

func (c *pool) IsExist(id string) bool {
	c.rw.RLock()
	defer c.rw.RUnlock()
	_, ok := c.connections[id]
	return ok
}

func (c *pool) Delete(id string) error {
	c.rw.Lock()
	defer c.rw.Unlock()
	delete(c.connections, id)
	return nil
}
