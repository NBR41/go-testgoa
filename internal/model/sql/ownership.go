package sql

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

const (
	qryListOwnershipsByUserID = `
SELECT books.id, books.isbn, books.name
FROM ownership
JOIN book ON (ownership.book_id = book.id)
WHERE ownership.user_id = ?`
	qryGetOwnership = `
SELECT book.id, book.isbn, book.name, book.series_id
FROM ownership
JOIN book ON (ownership.book_id = book.id)
WHERE ownership.user_id = ? AND book.id = ?`
	qryInsertOwnership = `
INSERT INTO ownership (user_id, book_id, create_ts, update_ts)
VALUES (?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`
	qryUpdateOwnership = `UPDATE ownership set update_ts = NOW() WHERE user_id = ? AND book_id = ?`
	qryDeleteOwnership = `DELETE FROM ownership WHERE user_id = ? AND book_id = ?`
)

// ListOwnershipsByUserID returns book list by user ID
func (m *Model) ListOwnershipsByUserID(userID int) ([]*model.Ownership, error) {
	_, err := m.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	rows, err := m.db.Query(qryListOwnershipsByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var l = []*model.Ownership{}
	for rows.Next() {
		b := model.Book{}
		if err := rows.Scan(&b.ID, &b.ISBN, &b.Name); err != nil {
			return nil, err
		}
		l = append(
			l,
			&model.Ownership{
				UserID: int64(userID),
				BookID: b.ID,
				Book:   &b,
			},
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return l, nil
}

// GetOwnership returns user book association
func (m *Model) GetOwnership(userID, bookID int) (*model.Ownership, error) {
	b, err := m.getBook(qryGetOwnership, userID, bookID)
	if err != nil {
		return nil, err
	}
	return &model.Ownership{UserID: int64(userID), BookID: int64(bookID), Book: b}, nil
}

// InsertOwnership inserts user book association
func (m *Model) InsertOwnership(userID, bookID int) (*model.Ownership, error) {
	_, err := m.db.Exec(qryInsertOwnership, userID, bookID)
	if err != nil {
		return nil, filterError(err)
	}
	return &model.Ownership{UserID: int64(userID), BookID: int64(bookID)}, nil
}

//UpdateOwnership update ownership
func (m *Model) UpdateOwnership(userID, bookID int) error {
	return m.exec(qryUpdateOwnership, userID, bookID)
}

// DeleteOwnership deletes user book association
func (m *Model) DeleteOwnership(userID, bookID int) error {
	return m.exec(qryDeleteOwnership, userID, bookID)
}
