package database

type Drivers struct {
	InMemoryStore Store
}

func NewDrivers() Drivers {
	return Drivers{
		InMemoryStore: NewInMemory(),
	}
}
