package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	val, ok := stateName[ss]
	if !ok {
		return "unknown-state"
	}
	return val
}

func main() {
	fmt.Println("===========")
	// StateNgasal := ServerState(98)

	// ns := transition(StateNgasal)
	ns := transition(StateRetrying)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
	fmt.Println("===========")
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("Unknown State: %s", s))
	}
}
