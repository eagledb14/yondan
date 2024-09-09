package utils

import (
	"sync"
	"errors"
	"strings"
)

// Writes probably won't scale well when getting to larger ip spaces, but it should be good enough for a /24
type ConcurrentMap struct {
	data map[string][]*Scan
	mutex sync.RWMutex
}

func NewConcurrentMap() *ConcurrentMap {
	newConcurrentMap := ConcurrentMap {
		data: make(map[string][]*Scan),
		mutex: sync.RWMutex{},
	}

	return &newConcurrentMap
}

func (c *ConcurrentMap) Read(target string) ([]*Scan, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	run, ok := c.data[target]
	if !ok {
		return nil, errors.New("Target Not Found: ")
	}
	return run, nil
}

func (c *ConcurrentMap) ReadAll() map[string][]*Scan {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	newMap := make(map[string][]*Scan)
	for k, v := range c.data {
		newMap[k] = v
	}
	// return c.data
	return newMap
}

func (c *ConcurrentMap) append(target string, scan *Scan) {
	if _, ok := c.data[target]; ok {
		for i, v := range c.data[target] {
			if target == "" {
				continue
			}
			if v.Ip == scan.Ip {
				c.data[target] = append(c.data[target][:i], c.data[target][i+1:]...)
				c.data[target] = append(c.data[target], scan)
				return
			}
		}
		c.data[target] = append(c.data[target], scan)
	} else {
		c.data[target] = []*Scan{scan}
	}
}

func (c *ConcurrentMap) Write(target string, run *Scan) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.append(target, run)
}

//returns true if the ip was already found from a scan, as to not rewrite it
func (c *ConcurrentMap) DummyWrite(target string, run *Scan) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.data[target]; ok {
		return true
	}

	c.append(target, run)
	return false
}

func (c *ConcurrentMap) MassWrite(newMap *map[string][]*Scan) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for target, value := range *newMap {
		for _, v := range value {
			c.append(target, v)
		}
	}
}

func (c *ConcurrentMap) Delete(target string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.data, target)
}

func (c *ConcurrentMap) Drop() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data = map[string][]*Scan{}
}

func (c *ConcurrentMap) String() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	out := strings.Builder{}
	for key, value := range c.data {
		out.WriteString(key + ":\n")
		for _, v := range value {
			out.WriteString("\t" + v.Ip + "\n")
		}
		out.WriteString("\n")
	}

	return out.String()
}

func (c *ConcurrentMap) Len() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.data)
}
