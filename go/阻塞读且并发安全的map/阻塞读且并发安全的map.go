package main

//https://github.com/lifei6671/interview-go/blob/master/question/q010.md

import (
	"fmt"
	"sync"
	"time"
)

type Map struct {
	c   map[string]*entry
	rmx *sync.RWMutex
}
type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Out(key string, val interface{}) {
	m.rmx.Lock()
	defer m.rmx.Unlock()
	if e, ok := m.c[key]; ok {
		e.value = val
		e.isExist = true
		close(e.ch)
	} else {
		e = &entry{ch: make(chan struct{}), isExist: true, value: val}
		m.c[key] = e
		close(e.ch)
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.Lock()
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.Unlock()
		fmt.Println(e.value)
		return e.value
	} else if !ok {
		e = &entry{ch: make(chan struct{}), isExist: false}
		m.c[key] = e
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			fmt.Println(e.value)
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			return nil
		}
	} else {
		m.rmx.Unlock()
		fmt.Println("协程阻塞 -> ", key)
		select {
		case <-e.ch:
			fmt.Println(e.value)
			return e.value
		case <-time.After(timeout):
			fmt.Println("协程超时 -> ", key)
			return nil
		}
	}
}

func main() {
	obj := &Map{make(map[string]*entry), &sync.RWMutex{}}
	obj.Out("qiubing", "enheng")
	tim := time.Duration(500000000)
	obj.Rd("qiubing", tim)
	go obj.Rd("qiubin", tim)
	go obj.Out("qiubin", "enhen")
	go obj.Rd("qiubi", tim)
	//go obj.Out("qiubi", "enhen")
	time.Sleep(tim * 2)
}
