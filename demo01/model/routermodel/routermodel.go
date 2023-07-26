package routermodel

type returnData struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    map[string]any `json:"data"`
}

func NewReturnData() returnData {
	returnData := returnData{}
	returnData.Status = 2000
	returnData.Message = ""
	returnData.Data = make(map[string]any)
	return returnData
}
