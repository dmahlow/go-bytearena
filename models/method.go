package models

const (
	HandshakeType = "Handshake"
	MutationType  = "Actions"
)

// TODO: Find a better name for that
type Method struct {
	AgentID string      `json:"agentid"`
	Method  string      `json:"method"`
	Payload interface{} `json:"payload"`
}
