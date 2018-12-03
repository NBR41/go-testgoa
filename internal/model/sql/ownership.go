package sql

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

// GetOwnershipList returns book list by user ID
func (m *Model) GetOwnershipList(userID int) ([]*model.Ownership, error) {
	rows, err := m.db.Query(
		`SELECT b.book_id, b.isbn, b.name FROM ownerships u JOIN books b USING(book_id) where user_id = ?`,
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
SELECT b.book_id, b.isbn, b.name
FROM ownerships u
JOIN books b USING(book_id)
where u.user_id = ? and b.book_id = ?`,
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
INSERT INTO ownerships (user_id, book_id, create_ts, update_ts)
VALUES (?, ?, NOW(), NOW())
ON DUPLICATE KEY UPDATE update_ts = VALUES(update_ts)`,
		userID, bookID,
	)
	if err != nil {
		return nil, err
	}
	return &model.Ownership{UserID: int64(userID), BookID: int64(bookID)}, nil
}

//UpdateOwnership update ownership
func (m *Model) UpdateOwnership(userID, bookID int) error {
	return m.exec(
		`UPDATE ownerships set update_ts = NOW() where user_id = ? and book_id = ?`,
		userID, bookID,
	)
}

// DeleteOwnership deletes user book association
func (m *Model) DeleteOwnership(userID, bookID int) error {
	return m.exec(`DELETE FROM ownerships where user_id = ? and book_id = ?`, userID, bookID)
}
