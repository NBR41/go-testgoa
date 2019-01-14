package sql

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

// ListOwnershipsByUserID returns book list by user ID
func (m *Model) ListOwnershipsByUserID(userID int) ([]*model.Ownership, error) {
	_, err := m.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	rows, err := m.db.Query(
		`
SELECT b.id, b.isbn, b.name
FROM ownership u
JOIN books b ON (u.book_id = b.id) where user_id = ?`,
		userID,
	)
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
	b, err := m.getBook(
		`
SELECT b.id, b.isbn, b.name, b.series_id
FROM ownership u
JOIN books b ON (u.book_id = b.id)
where u.user_id = ? and b.id = ?`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	return &model.Ownership{UserID: int64(userID), BookID: int64(bookID), Book: b}, nil
}

// InsertOwnership inserts user book association
func (m *Model) InsertOwnership(userID, bookID int) (*model.Ownership, error) {
	_, err := m.db.Exec(
		`
INSERT INTO ownership (user_id, book_id, create_ts, update_ts)
VALUES (?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		userID, bookID,
	)
	if err != nil {
		return nil, filterError(err)
	}
	return &model.Ownership{UserID: int64(userID), BookID: int64(bookID)}, nil
}

//UpdateOwnership update ownership
func (m *Model) UpdateOwnership(userID, bookID int) error {
	return m.exec(
		`UPDATE ownership set update_ts = NOW() where user_id = ? and book_id = ?`,
		userID, bookID,
	)
}

// DeleteOwnership deletes user book association
func (m *Model) DeleteOwnership(userID, bookID int) error {
	return m.exec(`DELETE FROM ownership where user_id = ? and book_id = ?`, userID, bookID)
}
