package repositories

type RepositoryInterface[T any] interface {
	FindByID(id uint) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}
