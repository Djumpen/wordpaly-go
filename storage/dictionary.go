package storage

type Dictionary struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	CreatedAt   string `db:"created_at" json:",omitempty"`
}

func (s *Storage) GetDictionaries(uid int) ([]*Dictionary, error) {
	return []*Dictionary{
		&Dictionary{
			ID:          100,
			Title:       "MyDict1",
			Description: "Funny words",
		},
		&Dictionary{
			ID:          101,
			Title:       "MyDict2",
			Description: "Another word",
		},
	}, nil
}

func (s *Storage) GetDictionary(uid, did int) (*Dictionary, error) {
	return &Dictionary{
		ID:          100,
		Title:       "MyDict1",
		Description: "Funny words",
	}, nil
}
