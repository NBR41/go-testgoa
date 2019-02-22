package google

import (
	"net/http"
	"strconv"

	"github.com/NBR41/go-testgoa/internal/api"
	"google.golang.org/api/books/v1"
)

//Caller interface to get google book volumes
type Caller interface {
	Get(isbn string) (*books.Volumes, error)
}

//HTTPCaller struct for HTTP caller, implementing Caller
type HTTPCaller struct{}

//Get return books.Volumes for an isbn
func (HTTPCaller) Get(isbn string) (*books.Volumes, error) {
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

// GetBookDetail returns book detail by calling Google Book API
func (g Google) GetBookDetail(isbn string) (*api.BookDetail, error) {
	vols, err := g.c.Get(isbn)
	if err != nil {
		return nil, err
	}

	if vols.TotalItems == 0 {
		return nil, api.ErrNoResult
	}

	ret := &api.BookDetail{}
	ret.Title = vols.Items[0].VolumeInfo.Title

	if vols.Items[0].VolumeInfo.SeriesInfo != nil {
		vol, err := strconv.Atoi(vols.Items[0].VolumeInfo.SeriesInfo.BookDisplayNumber)
		if err != nil {
			return nil, err
		}
		ret.Volume = &vol
	}
	if vols.Items[0].VolumeInfo.Subtitle != "" {
		ret.Subtitle = &vols.Items[0].VolumeInfo.Subtitle
	}
	if len(vols.Items[0].VolumeInfo.Authors) > 0 {
		for i := range vols.Items[0].VolumeInfo.Authors {
			ret.Authors = append(ret.Authors, &api.Author{Name: vols.Items[0].VolumeInfo.Authors[i]})
		}
	}
	if vols.Items[0].VolumeInfo.Publisher != "" {
		ret.Editor = &vols.Items[0].VolumeInfo.Publisher
	}
	if vols.Items[0].VolumeInfo.Description != "" {
		ret.Description = &vols.Items[0].VolumeInfo.Description
	}
	return ret, nil
}
