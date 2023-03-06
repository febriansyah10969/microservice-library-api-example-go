package helper

type Response struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    interface{}       `json:"meta"`
	Error   []interface{} `json:"error"`
}

func APIResponse(code int, status bool, message string, pagination, data, err interface{}) Response {
	apiResponse := Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    pagination,
	}

	if err == nil {
		apiResponse.Error = []interface{}{}
	} else {
		apiResponse.Error = []interface{}{err}
	}

	return apiResponse
}