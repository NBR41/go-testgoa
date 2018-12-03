package google

import (
	"errors"
	"testing"

	"github.com/NBR41/go-testgoa/internal/api"
	"github.com/golang/mock/gomock"

	"google.golang.org/api/books/v1"
)

func TestGetBookName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mock := NewMockCaller(ctrl)
	gomock.InOrder(
		mock.EXPECT().Get("foo").Return(nil, errors.New("get error")),
		mock.EXPECT().Get("foo").Return(&books.Volumes{}, nil),
		mock.EXPECT().Get("foo").Return(&books.Volumes{TotalItems: 1, Items: []*books.Volume{&books.Volume{VolumeInfo: &books.VolumeVolumeInfo{Title: "bar", SeriesInfo: &books.Volumeseriesinfo{BookDisplayNumber: "2"}}}}}, nil),
	)

	tests := []struct {
		desc string
		exp  string
		err  error
	}{
		{"get error", "", errors.New("get error")},
		{"not found", "", api.ErrNoResult},
		{"valid", "bar 2", nil},
	}
	g := New(mock)
	for i := range tests {
		v, err := g.GetBookName("foo")
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
		if v != tests[i].exp {
			t.Errorf("unexpected value, exp [%s] got [%s]", tests[i].exp, v)
		}
	}
}
