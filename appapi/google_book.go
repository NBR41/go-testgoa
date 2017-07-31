package appapi

import (
	"net/http"

	"google.golang.org/api/books/v1"
)

// GetBookName returns book name by calling Google Book API
func GetBookName(isbn string) (string, error) {
	svc, err := books.New(http.DefaultClient)
	if err != nil {
		return "", err
	}
	vols, err := books.NewVolumesService(svc).List("isbn:" + isbn).Do()
	if err != nil {
		return "", err
	}

	if vols.TotalItems == 0 {
		return "", err
	}

	var title = vols.Items[0].VolumeInfo.Title

	if vols.Items[0].VolumeInfo.SeriesInfo != nil {
		title += " " + vols.Items[0].VolumeInfo.SeriesInfo.BookDisplayNumber
	}
	return title, nil
}
