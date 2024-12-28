package responses

import "echo-base/internal/contract"

type Response struct {
	Code     int                    `json:"code"`
	Message  string                 `json:"message"`
	Data     interface{}            `json:"data"`
	Metadata *contract.ListResponse `json:"metadata"`
	Error    error                  `json:"error"`
	Success  bool                   `json:"success"`
}

func ResponseData(data interface{}, metadata *contract.ListResponse, err error, code int) Response {
	if err != nil {
		return Response{
			Code:     code,
			Message:  err.Error(),
			Data:     data,
			Metadata: metadata,
			Error:    err,
			Success:  false,
		}
	}
	return Response{
		Code:     code,
		Message:  "successfully",
		Data:     data,
		Metadata: metadata,
		Error:    err,
		Success:  true,
	}
}
