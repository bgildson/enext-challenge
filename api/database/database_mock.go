package database

type mockDatabase struct {
	get     func() ([]map[string]interface{}, error)
	getByID func(id string) (map[string]interface{}, error)
}

// NewMockDatabase generate a new Database instance for mock
func NewMockDatabase(
	get func() ([]map[string]interface{}, error),
	getByID func(id string) (map[string]interface{}, error),
) Database {
	return &mockDatabase{
		get:     get,
		getByID: getByID,
	}
}

func (d *mockDatabase) Get() ([]map[string]interface{}, error) {
	return d.get()
}

func (d *mockDatabase) GetByID(id string) (map[string]interface{}, error) {
	return d.getByID(id)
}
