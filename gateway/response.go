package gateway

type ErrData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message int    `json:"message"`
}

//func ResponseError(ctx context.Context, c echo.Context, httpCode int, err error) error {
//	return c.JSON(httpCode, &ErrData{Code: httpCode, Status: "Failed", Message: errors.w})
//}
