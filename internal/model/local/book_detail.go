package local

import (
    	"github.com/NBR41/go-testgoa/internal/model"
        "github.com/NBR41/go-testgoa/internal/api"
)

//GetBookDetail return a Book detail by ISBN
func (db *Local) GetBookDetail(isbn string) (*model.BookDetail, error) {
    // TODO
    return nil,nil
}

//InsertBookDetail insert all values in book detail
func (db *Local) InsertBookDetail(isbn string, book *api.BookDetail) (*model.BookDetail, error) {
    // TODO
    return nil,nil
}
