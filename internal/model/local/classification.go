package local

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

func (db *Local) getClassification(seriesID, classID int) (*model.Class, error) {
	for _, v := range db.classifications {
		if v.SeriesID == seriesID && v.ClassID == classID {
			return v.Class, nil
		}
	}
	return nil, model.ErrNotFound
}

//GetClassification get a classification
func (db *Local) GetClassification(seriesID, classID int) (*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	return db.getClassification(seriesID, classID)
}

//InsertClassification inserts a classification
func (db *Local) InsertClassification(seriesID, classID int) (*model.Class, error) {
	db.Lock()
	defer db.Unlock()
	for _, v := range db.classifications {
		if v.SeriesID == seriesID && v.ClassID == classID {
			return nil, model.ErrDuplicateKey
		}
	}

	if _, ok := db.series[seriesID]; !ok {
		return nil, model.ErrInvalidID
	}
	c, err := db.getClassByID(classID)
	if err != nil {
		return nil, model.ErrInvalidID
	}

	idx := len(db.classifications) + 1
	v := &classification{ID: idx, SeriesID: seriesID, ClassID: classID, Class: c}
	db.classifications[idx] = v
	return c, nil
}

//DeleteClassification deletes a classification
func (db *Local) DeleteClassification(seriesID, classID int) error {
	db.Lock()
	defer db.Unlock()
	for id, v := range db.classifications {
		if v.SeriesID == seriesID && v.ClassID == classID {
			delete(db.classifications, id)
			return nil
		}
	}
	return model.ErrNotFound
}
