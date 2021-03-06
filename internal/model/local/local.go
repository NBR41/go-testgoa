package local

import (
	"sync"

	"github.com/NBR41/go-testgoa/internal/model"
)

// Local emulates a database driver using in-memory data structures.
type Local struct {
	pass model.Passworder
	sync.Mutex
	users           map[int]*model.User
	books           map[int]*model.Book
	ownerships      map[int][]*model.Book
	authors         map[int]*model.Author
	categories      map[int]*model.Category
	collections     map[int]*model.Collection
	prints          map[int]*model.Print
	editors         map[int]*model.Editor
	classes         map[int]*model.Class
	roles           map[int]*model.Role
	series          map[int]*model.Series
	authorships     map[int]*model.Authorship
	editions        map[int]*model.Edition
	classifications map[int]*classification
}

type classification struct {
	ID       int
	SeriesID int
	ClassID  int
	Class    *model.Class
}

// New returns new instance of Local storage
func New(pass model.Passworder) *Local {
	book := &model.Book{ID: 1, ISBN: "isbn-123", Name: "test1", SeriesID: 1}
	book2 := &model.Book{ID: 2, ISBN: "isbn-456", Name: "test2"}
	book3 := &model.Book{ID: 3, ISBN: "isbn-789", Name: "test3"}
	book4 := &model.Book{ID: 4, ISBN: "isbn-135", Name: "test4"}
	author := &model.Author{ID: 1, Name: "author1"}
	category := &model.Category{ID: 1, Name: "category1"}
	print := &model.Print{ID: 1, Name: "print1"}
	editor := &model.Editor{ID: 1, Name: "editor1"}
	class := &model.Class{ID: 1, Name: "class1"}
	role := &model.Role{ID: 1, Name: "role1"}
	collection := &model.Collection{ID: 1, Name: "collection1", EditorID: 1, Editor: editor}
	series := &model.Series{ID: 1, Name: "series1", CategoryID: 1, Category: category}
	authorship := &model.Authorship{ID: 1, AuthorID: 1, RoleID: 1, BookID: 1}
	edition := &model.Edition{ID: 1, CollectionID: 1, BookID: 1, PrintID: 1}
	classific := &classification{ID: 1, SeriesID: 1, ClassID: 1, Class: class}

	return &Local{
		pass: pass,
		users: map[int]*model.User{
			3: &model.User{ID: 3, Email: `user@myinventory.com`, Nickname: `user`, IsValidated: true, IsAdmin: false},
			2: &model.User{ID: 2, Email: `new@myinventory.com`, Nickname: `new`, IsValidated: false, IsAdmin: false},
			1: &model.User{
				ID:          1,
				Email:       `admin@myinventory.com`,
				Nickname:    `admin`,
				IsValidated: true,
				IsAdmin:     true,
				Salt:        []byte("\xd6\xe8\u007f Yg\xbc\xe7@\x8b\xe4E\x9b\xb8\xc3\xeepZ\xe0\x90Z\xe4C\xd5%\xe7RP9a(\xfb"),
				Password:    []byte("'\xeb\xbe\x1f\xbaaG\xe1&>\x9f \u007f\xc94^\xdf\xca*\xdb\xf6<\x05\x05A8q\x94\xd0k\xc23\xf9\xd5\xdb-\x8f\x1c\f\xa5\xa1\xcf\xcf\xe1\t\xde\xf4\x89\x81B\x06\x16\x0ecQ\x94*\xa0D\x82\x1dUeJ"),
			},
			4: &model.User{ID: 4, Email: `nobooks@myinventory.com`, Nickname: `nobooks`, IsValidated: true, IsAdmin: false},
		},
		books: map[int]*model.Book{
			1: book,
			2: book2,
			3: book3,
			4: book4,
		},
		ownerships: map[int][]*model.Book{
			1: []*model.Book{book, book4},
			2: []*model.Book{},
			3: []*model.Book{book2, book3},
		},
		authors:         map[int]*model.Author{1: author},
		categories:      map[int]*model.Category{1: category},
		prints:          map[int]*model.Print{1: print},
		editors:         map[int]*model.Editor{1: editor},
		classes:         map[int]*model.Class{1: class},
		roles:           map[int]*model.Role{1: role},
		collections:     map[int]*model.Collection{1: collection},
		series:          map[int]*model.Series{1: series},
		authorships:     map[int]*model.Authorship{1: authorship},
		editions:        map[int]*model.Edition{1: edition},
		classifications: map[int]*classification{1: classific},
	}
}

// Close close the connextion
func (db *Local) Close() error {
	return nil
}
