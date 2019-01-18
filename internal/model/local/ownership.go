package local

import (
	"github.com/NBR41/go-testgoa/internal/model"
)

// ListOwnershipsByUserID returns book list by user ID
func (db *Local) ListOwnershipsByUserID(userID int) ([]*model.Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, model.ErrNotFound
	}

	var list = []*model.Ownership{}
	for i := range l {
		list = append(
			list,
			&model.Ownership{
				UserID: int64(userID),
				BookID: l[i].ID,
				Book:   l[i],
			},
		)
	}
	return list, nil
}

// GetOwnership returns user book association
func (db *Local) GetOwnership(userID, bookID int) (*model.Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, model.ErrNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			return &model.Ownership{UserID: int64(userID), BookID: int64(bookID), Book: l[i]}, nil
		}
	}
	return nil, model.ErrNotFound
}

// InsertOwnership inserts user book association
func (db *Local) InsertOwnership(userID, bookID int) (*model.Ownership, error) {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return nil, model.ErrInvalidID
	}

	b, ok := db.books[bookID]
	if !ok {
		return nil, model.ErrInvalidID
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			return nil, model.ErrDuplicateKey
		}
	}

	db.ownerships[userID] = append(l, b)
	return &model.Ownership{UserID: int64(userID), BookID: int64(bookID), Book: b}, nil
}

//UpdateOwnership update the ownership
func (db *Local) UpdateOwnership(userID, bookID int) error {
	return nil
}

// DeleteOwnership deletes user book association
func (db *Local) DeleteOwnership(userID, bookID int) error {
	db.Lock()
	defer db.Unlock()
	l, ok := db.ownerships[userID]
	if !ok {
		return model.ErrNotFound
	}

	for i := range l {
		if l[i].ID == int64(bookID) {
			db.ownerships[userID] = append(l[:i], l[i+1:]...)
			return nil
		}
	}
	return model.ErrNotFound
}
