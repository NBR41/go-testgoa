package sql

import (
	//	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

func (m *Model) GetCollectionByID(id int) (*model.Collection, error) {
	return nil, nil
}

func (m *Model) GetCollectionByName(name string) (*model.Collection, error) {
	return nil, nil
}

func (m *Model) InsertCollection(name string, editorID int) (*model.Collection, error) {
	return nil, nil
}

func (m *Model) UpdateCollection(name *string, editorID *int) error {
	return nil
}

func (m *Model) DeleteCollection(id int) error {
	return nil
}

func (m *Model) ListCollections() ([]*model.Collection, error) {
	return nil, nil
}

func (m *Model) ListCollectionsByEditorID(id int) ([]*model.Collection, error) {
	return nil, nil
}
