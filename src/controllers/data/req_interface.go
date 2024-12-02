package data

type ReqInterface interface {
	Validate() (bool, error)
}
