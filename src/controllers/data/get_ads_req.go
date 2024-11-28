package data

type GetAdsReq struct {
	Offset   *int   `form:"offset" validate:"omitempty,gte=0"`
	Limit    *int   `form:"limit" validate:"omitempty,gte=1,lte=100"`
	Age      *int   `form:"age" validate:"omitempty,gte=1,lte=100"`
	Gender   string `form:"gender" validate:"omitempty,oneof=F M"`
	Country  string `form:"country" validate:"omitempty,country_code"`
	Platform string `form:"platform" validate:"omitempty,oneof=ios android web"`
}

func (req GetAdsReq) Validate() (bool, error) {
	return true, nil
}
