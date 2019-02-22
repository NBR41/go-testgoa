package google

import (
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/internal/api"
	"github.com/golang/mock/gomock"
	"github.com/kylelemons/godebug/pretty"

	"google.golang.org/api/books/v1"
)

func TestGetBookName(t *testing.T) {
	var (
		author, editor, subtitle, description = "fooAuthor", "fooEditor", "fooSubtitle", "fooDesc"
		volume                                = 2
	)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockCaller(ctrl)
	gomock.InOrder(
		mock.EXPECT().Get("foo").Return(nil, errors.New("get error")),
		mock.EXPECT().Get("foo").Return(&books.Volumes{}, nil),
		mock.EXPECT().Get("foo").Return(
			&books.Volumes{
				TotalItems: 1,
				Items: []*books.Volume{
					&books.Volume{
						VolumeInfo: &books.VolumeVolumeInfo{
							Title: "bar",
							SeriesInfo: &books.Volumeseriesinfo{
								BookDisplayNumber: "a",
							},
						},
					},
				},
			},
			nil,
		),
		mock.EXPECT().Get("foo").Return(
			&books.Volumes{
				TotalItems: 1,
				Items: []*books.Volume{
					&books.Volume{
						VolumeInfo: &books.VolumeVolumeInfo{
							Title:       "bar",
							Subtitle:    subtitle,
							Authors:     []string{author},
							Publisher:   editor,
							Description: description,
							SeriesInfo: &books.Volumeseriesinfo{
								BookDisplayNumber: "2",
							},
						},
					},
				},
			},
			nil,
		),
	)

	tests := []struct {
		desc string
		exp  *api.BookDetail
		err  error
	}{
		{"get error", nil, errors.New("get error")},
		{"not found", nil, api.ErrNoResult},
		{"strconv error", nil, errors.New(`strconv.Atoi: parsing "a": invalid syntax`)},
		{
			"valid",
			&api.BookDetail{
				Title:       "bar",
				Subtitle:    &subtitle,
				Authors:     []*api.Author{&api.Author{Name: author}},
				Editor:      &editor,
				Description: &description,
				Volume:      &volume,
			},
			nil,
		},
	}
	g := New(mock)
	for i := range tests {
		v, err := g.GetBookDetail("foo")
		if err != nil {
			if tests[i].err == nil {
				t.Errorf("unexpected error for [%s], [%v]", tests[i].desc, err)
				continue
			}
			if tests[i].err.Error() != err.Error() {
				t.Errorf("unexpected error for [%s], exp [%v] got [%v]", tests[i].desc, tests[i].err, err)
				continue
			}
			continue
		}
		if tests[i].err != nil {
			t.Errorf("expecting error for [%s]", tests[i].desc)
		}
		if diff := pretty.Compare(v, tests[i].exp); diff != "" {
			t.Errorf("unexpected value for [%s]\n%s", tests[i].desc, diff)
		}
	}
}
