package sql

import (
	"database/sql"

	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryGetClassification = `
SELECT class.id, class.name
FROM class
JOIN classification ON (classification.class_id = class.id)
WHERE classification.series_id = ? AND classification.class_id = ?`
	qryInsertClassification = `
INSERT INTO classification (id, series_id, class_id, create_ts, update_ts)
VALUES (null, ?, ?, NOW(), NOW())`
	qryDeleteClassification = `DELETE FROM classification WHERE series_id = ? and class_id = ?`
)

//GetClassification returns a class from a series
func (m *Model) GetClassification(seriesID, classID int) (*model.Class, error) {
	var v = model.Class{}
	err := m.db.QueryRow(qryGetClassification, seriesID, classID).Scan(&v.ID, &v.Name)
	switch {
	case err == sql.ErrNoRows:
		return nil, model.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return &v, nil
	}
}

//InsertClassification insert a classification
func (m *Model) InsertClassification(seriesID, classID int) (*model.Class, error) {
	_, err := m.db.Exec(qryInsertClassification, seriesID, classID)
	if err != nil {
		return nil, filterError(err)
	}
	return &model.Class{ID: int64(classID)}, nil

}

//DeleteClassification deletes a classification
func (m *Model) DeleteClassification(seriesID, classID int) error {
	return m.exec(qryDeleteClassification, seriesID, classID)
}
