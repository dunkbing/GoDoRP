package api

type HttpError struct {
	Id            string `json:"-"`
	Message       string `json:"message"`                 // Message to be display to the end user without debugging information
	DetailedError string `json:"detailedError,omitempty"` // Internal error string to help the developer
	RequestId     string `json:"-"`                       // The RequestId that's also set in the header
	StatusCode    int    `json:"statusCode,omitempty"`    // The http status code
	Where         string `json:"-"`                       // The function where it happened in the form of Struct.Func
	IsOAuth       bool   `json:"-"`                       // Whether the error is OAuth specific
	// params        map[string]interface{}
}

func (er *HttpError) Error() string {
	return er.Where + ": " + er.Message + ", " + er.DetailedError
}
