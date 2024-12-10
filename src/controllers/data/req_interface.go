package data

type IRequest interface {
	Validate() (bool, error)
}
