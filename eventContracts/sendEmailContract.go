package eventcontracts

// SendEmailEvent Event type
type SendEmailEvent struct {
	From    string `json:"from"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
