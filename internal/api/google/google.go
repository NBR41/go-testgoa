package google

import (
	"net/http"

	"github.com/NBR41/go-testgoa/internal/api"
	"google.golang.org/api/books/v1"
)

type Caller interface {
	Get(isbn string) (*books.Volumes, error)
}

type HttpCaller struct{}

func (HttpCaller) Get(isbn string) (*books.Volumes, error) {
	svc, err := books.New(http.DefaultClient)
	if err != nil {
		return nil, err
	}
	return books.NewVolumesService(svc).List("isbn:" + isbn).Do()
}

type Google struct {
	c Caller
}

func New(c Caller) *Google {
	return &Google{c: c}
}

// GetBookName returns book name by calling Google Book API
func (g Google) GetBookName(isbn string) (string, error) {
	vols, err := g.c.Get(isbn)
	if err != nil {
		return "", err
	}

	if vols.TotalItems == 0 {
		return "", api.ErrNoResult
	}

	var title = vols.Items[0].VolumeInfo.Title

	if vols.Items[0].VolumeInfo.SeriesInfo != nil {
		title += " " + vols.Items[0].VolumeInfo.SeriesInfo.BookDisplayNumber
	}
	return title, nil
}
