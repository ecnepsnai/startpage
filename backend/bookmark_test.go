package startpage

import (
	"testing"

	"github.com/ecnepsnai/web"
)

func TestBookmarkAddGetDelete(t *testing.T) {
	bookmark, err := BookmarkStore.New(newBookmarkParameters{
		URL:  randomString(6),
		Name: randomString(6),
	})
	if err != nil {
		t.Fatalf("Error adding bookmark: %s", err.Message)
	}

	if BookmarkStore.Get(bookmark.ID) == nil {
		t.Fatalf("No bookmark returned when one expected")
	}

	if err := BookmarkStore.Delete(bookmark.ID); err != nil {
		t.Fatalf("Error deleting bookmark: %s", err.Message)
	}
}

func TestBookmarkReorder(t *testing.T) {
	BookmarkStore.Table.DeleteAll()

	i := 0
	for i < 5 {
		_, err := BookmarkStore.New(newBookmarkParameters{
			URL:  randomString(6),
			Name: randomString(6),
		})
		if err != nil {
			t.Fatalf("Error adding bookmark: %s", err.Message)
		}
		i++
	}

	bookmarks := BookmarkStore.AllBookmarks()
	order := map[string]int{}
	for _, bookmark := range bookmarks {
		i--
		order[bookmark.ID] = i
	}

	if err := BookmarkStore.Reorder(order); err != nil {
		t.Fatalf("Error reordering bookmarks: %s", err.Message)
	}
}

func TestHandleBookmarkGetAll(t *testing.T) {
	t.Parallel()

	i := 0
	for i < 5 {
		_, err := BookmarkStore.New(newBookmarkParameters{
			URL:  randomString(6),
			Name: randomString(6),
		})
		if err != nil {
			t.Fatalf("Error adding bookmark: %s", err.Message)
		}
		i++
	}

	api := apiHandle{}
	obj, err := api.BookmarkGetAll(web.MockRequest(nil, nil, nil))
	if err != nil {
		t.Fatalf("Error getting all bookmarks: %s", err.Message)
	}
	results, ok := obj.([]Bookmark)
	if !ok {
		t.Fatalf("Unexpected object from API handle")
	}
	if len(results) < 5 {
		t.Fatalf("Unexpected result from API handle")
	}
}

func TestHandleBookmarkNewGet(t *testing.T) {
	t.Parallel()

	name := randomString(6)
	url := randomString(6)

	api := apiHandle{}
	obj, err := api.BookmarkNew(web.MockRequest(nil, nil, newBookmarkParameters{
		Name: name,
		URL:  url,
	}))
	if err != nil {
		t.Fatalf("Error making new bookmark: %s", err.Message)
	}
	result, ok := obj.(*Bookmark)
	if !ok {
		t.Fatalf("Unexpected object from API handle")
	}

	if result.Name != name {
		t.Fatalf("Unexpected name for created bookmark")
	}
	if result.URL != url {
		t.Fatalf("Unexpected url for created bookmark")
	}

	obj, err = api.BookmarkGet(web.MockRequest(nil, map[string]string{"id": result.ID}, nil))
	if err != nil {
		t.Fatalf("Error getting bookmark: %s", err.Message)
	}
	result, ok = obj.(*Bookmark)
	if !ok {
		t.Fatalf("Unexpected object from API handle")
	}

	if result.Name != name {
		t.Fatalf("Unexpected name for created bookmark")
	}
	if result.URL != url {
		t.Fatalf("Unexpected url for created bookmark")
	}
}

func TestHandleBookmarkEdit(t *testing.T) {
	t.Parallel()

	name := randomString(6)
	url := randomString(6)

	api := apiHandle{}
	obj, err := api.BookmarkNew(web.MockRequest(nil, nil, newBookmarkParameters{
		Name: name,
		URL:  url,
	}))
	if err != nil {
		t.Fatalf("Error making new bookmark: %s", err.Message)
	}
	result, ok := obj.(*Bookmark)
	if !ok {
		t.Fatalf("Unexpected object from API handle")
	}

	if result.Name != name {
		t.Fatalf("Unexpected name for created bookmark")
	}
	if result.URL != url {
		t.Fatalf("Unexpected url for created bookmark")
	}

	obj, err = api.BookmarkEdit(web.MockRequest(nil, map[string]string{
		"id": result.ID,
	}, updateBookmarkParameters{
		Name: randomString(6),
		URL:  randomString(6),
	}))
	if err != nil {
		t.Fatalf("Error updating bookmark: %s", err.Message)
	}
	result, ok = obj.(*Bookmark)
	if !ok {
		t.Fatalf("Unexpected object from API handle")
	}

	if result.Name == name {
		t.Fatalf("Unexpected name for created bookmark")
	}
	if result.URL == url {
		t.Fatalf("Unexpected url for created bookmark")
	}
}

func TestHandleBookmarkDelete(t *testing.T) {
	t.Parallel()

	api := apiHandle{}
	obj, err := api.BookmarkNew(web.MockRequest(nil, nil, newBookmarkParameters{
		Name: randomString(6),
		URL:  randomString(6),
	}))
	if err != nil {
		t.Fatalf("Error making new bookmark: %s", err.Message)
	}
	result, ok := obj.(*Bookmark)
	if !ok {
		t.Fatalf("Unexpected object from API handle")
	}

	obj, err = api.BookmarkDelete(web.MockRequest(nil, map[string]string{
		"id": result.ID,
	}, nil))
	if err != nil {
		t.Fatalf("Error deleting bookmark: %s", err.Message)
	}
	didDelete, ok := obj.(bool)
	if !ok {
		t.Fatalf("Unexpected object from API handle")
	}

	if !didDelete {
		t.Fatalf("Unexpected result from API handle")
	}
}
