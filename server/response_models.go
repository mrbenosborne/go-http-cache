package server

// Acknowledged a simple response for acknowledging
// various operations.
type Acknowledged struct {
	Acknowledged bool `json:"acknowledged"`
}
