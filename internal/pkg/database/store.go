package database

type Store interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	GetSlice(key string) ([]any, error)
	Values() ([]any, error)
}
