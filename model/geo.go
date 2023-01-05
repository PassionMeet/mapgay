package model

type UploadGeoRequest struct {
	Openid             string  `json:"openid"`
	UploadTs           int64   `json:"upload_ts"`
	Latitude           float64 `json:"latitude"`  //纬度
	Longitude          float64 `json:"longitude"` //经度
	Speed              int32   `json:"speed"`     //速度m/s
	Accuracy           int32   `json:"accuracy"`  //位置精确度
	Altitude           int32   `json:"altitude"`  //高度m
	ErrMsg             string  `json:"err_msg"`
	VerticalAccuracy   int32   `json:"verticalAccuracy"`   //垂直精度m
	HorizontalAccuracy int32   `json:"horizontalAccuracy"` //水平精度
}
