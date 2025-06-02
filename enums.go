package main

import "fmt"

type ServerState int

const (
	serverIdle ServerState = iota
	serverConnected
	serverReconnecting
	serverError
)

var StateName = map[ServerState]string{
	serverIdle:         "idle",
	serverConnected:    "connected",
	serverReconnecting: "reconnecting",
	serverError:        "server error",
}

func (s ServerState) String() string {
	return StateName[s]
}

func main() {
	s := transition(serverIdle)
	fmt.Println(s)

	n := transition(serverConnected)
	fmt.Println(n)
}

func transition(s ServerState) ServerState {
	switch s {
	case serverIdle:
		return serverConnected
	case serverConnected, serverReconnecting:
		return serverIdle
	case serverError:
		return serverError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
