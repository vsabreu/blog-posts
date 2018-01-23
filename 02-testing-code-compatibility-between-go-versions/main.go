package main

import (
	"fmt"
	"sync"
)

var sm sync.Map

func init() {
	sm.Store("meow", "mmmmeow!")
}

func main() {
	fmt.Println(getValue("meow"))
}

func getValue(v string) string {
	r, ok := sm.Load(v)
	if ok {
		return r.(string)
	}
	return "sorry, nothing here"
}
