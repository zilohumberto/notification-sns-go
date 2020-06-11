package sending

// Notification body to load a SNS message
type Notification struct {
	Target string `json:"target"`
	Topic string `json:"topic"`
	Message string `json:"message"`
	Subject string `json:"subject"`
}
