package base

type ReqInterface interface {
	Validate() (bool, error)
}
