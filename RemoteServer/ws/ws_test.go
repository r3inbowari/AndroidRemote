package ws

import (
	"RemoteServer/event"
	"encoding/json"
	"testing"
)

func TestUM(t *testing.T) {
	ab := event.Event{}
	a := "{\"id\":\"a8f5f167f44f4964e6c998dee827110c\",\"contact\":0,\"x\":124,\"y\":391,\"ts\":1618574242}"
	k :=  json.Unmarshal([]byte(a), &ab)
	println(k)
}
