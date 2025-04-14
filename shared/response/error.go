package response

// ErrorStruct model
type ErrorStruct struct {
	HTTPCode           int         `json:"-"`
	ResponseCode       string      `json:"rc"`
	Description        string      `json:"desc,omitempty"`
	Message            string      `json:"msg"`
	MessageDescription string      `json:"msg_desc"`
	Data               interface{} `json:"data,omitempty"`
}

// Implementasi interface error
func (e ErrorStruct) Error() string {
	return e.Message
}
