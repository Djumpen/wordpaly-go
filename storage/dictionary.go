package storage

import "github.com/djumpen/wordplay-go/entities"

func (s *Storage) GetDictionaries(uid int) ([]*entities.Dictionary, error) {
	return []*entities.Dictionary{
		&entities.Dictionary{
			ID:          100,
			Title:       "MyDict1",
			Description: "Funny words",
		},
		&entities.Dictionary{
			ID:          101,
			Title:       "MyDict2",
			Description: "Another word",
		},
	}, nil
}

func (s *Storage) GetDictionary(uid, did int) (*entities.Dictionary, error) {
	return &entities.Dictionary{
		ID:          100,
		Title:       "MyDict1",
		Description: "Funny words",
	}, nil
}
