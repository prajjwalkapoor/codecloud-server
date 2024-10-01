package model

import "encoding/json"

type WSEvent struct {
	EventType string          `json:"event_type"`
	Data      json.RawMessage `json:"data"`
}
