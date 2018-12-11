package local

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) GetCollectionByID(id int) (*model.Collection, error) {
	return nil, nil
}

func (db *Local) GetCollectionByName(name string) (*model.Collection, error) {
	return nil, nil
}

func (db *Local) InsertCollection(name string, editorID int) (*model.Collection, error) {
	return nil, nil
}

func (db *Local) UpdateCollection(name *string, editorID *int) error {
	return nil
}

func (db *Local) DeleteCollection(id int) error {
	return nil
}

func (db *Local) ListCollections() ([]*model.Collection, error) {
	return nil, nil
}

func (db *Local) ListCollectionsByEditorID(id int) ([]*model.Collection, error) {
	return nil, nil
}
