package sending

// Notification body to load a SNS message
type Notification struct {
	Topic string `json:"topic"`
	Message string `json:"message"`
}
