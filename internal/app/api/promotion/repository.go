package promotion

// Repository is the interface
type Repository interface {
	Create(data *Promotion) error
	Find(id int) (Promotion, error)
	FindList() ([]Promotion, error)
}

// Encapsulated Implementation of Repository Interface
type repository struct {
	Ds DataSource
}

// NewRepository is the function
func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}

func (repo *repository) Create(data *Promotion) error {

	return repo.Ds.Create(data)
}

func (repo *repository) FindList() ([]Promotion, error) {

	return repo.Ds.FindList()
}

func (repo *repository) Find(id int) (Promotion, error) {

	return repo.Ds.Find(id)
}
