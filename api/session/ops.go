package session

import "sync"

type SimpleSession struct {
	Username string
	TTL int64
}
var session *sync.Map
