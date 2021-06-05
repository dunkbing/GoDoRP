package utils

type AppError struct {
	Id            string `json:"id,omitempty"`
	Message       string `json:"message"`                  // Message to be display to the end user without debugging information
	DetailedError string `json:"detailed_error,omitempty"` // Internal error string to help the developer
	RequestId     string `json:"request_id,omitempty"`     // The RequestId that's also set in the header
	StatusCode    int    `json:"status_code,omitempty"`    // The http status code
	Where         string `json:"-"`                        // The function where it happened in the form of Struct.Func
	IsOAuth       bool   `json:"is_oauth,omitempty"`       // Whether the error is OAuth specific
	// params        map[string]interface{}
}

func (er *AppError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}
