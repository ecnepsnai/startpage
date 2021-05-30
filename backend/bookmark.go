package startpage

import (
	"io"
	"os"
	"path"
	"reflect"

	"github.com/ecnepsnai/limits"
)

// Bookmark describes a bookmark
type Bookmark struct {
	ID        string `ds:"primary"`
	URL       string `ds:"unique" min:"1"`
	Name      string `min:"1" max:"128"`
	Index     int
	ImageSize uint64
	ImageMime string
}

// AllBookmarks will return an ordered list of all bookmarks
func (s *bookmarkStoreObject) AllBookmarks() []Bookmark {
	objects, err := s.Table.GetAll(nil)
	if err != nil {
		log.PError("Error getting all bookmarks", map[string]interface{}{
			"error": err.Error(),
		})
		return []Bookmark{}
	}
	count := len(objects)
	if count == 0 {
		return []Bookmark{}
	}

	bookmarks := make([]Bookmark, count)
	for i, obj := range objects {
		bookmark, ok := obj.(Bookmark)
		if !ok {
			log.PPanic("Invalid object in bookmark store", map[string]interface{}{
				"index": i,
				"type":  reflect.TypeOf(obj),
			})
		}
		bookmarks[i] = bookmark
	}

	return bookmarks
}

// Get will return the bookmark with the given ID
func (s *bookmarkStoreObject) Get(id string) *Bookmark {
	object, err := s.Table.Get(id)
	if err != nil {
		log.PError("Error getting bookmark with ID", map[string]interface{}{
			"id":    id,
			"error": err.Error(),
		})
		return nil
	}
	if object == nil {
		return nil
	}
	bookmark, ok := object.(Bookmark)
	if !ok {
		log.PPanic("Invalid object in bookmark store", map[string]interface{}{
			"id":   id,
			"type": reflect.TypeOf(object),
		})
	}
	return &bookmark
}

type newBookmarkParameters struct {
	URL  string
	Name string
}

// New will save a new bookmark
func (s *bookmarkStoreObject) New(params newBookmarkParameters) (*Bookmark, *Error) {
	bookmark := Bookmark{
		ID:    newID(),
		URL:   params.URL,
		Name:  params.Name,
		Index: len(s.AllBookmarks()),
	}

	if err := limits.Check(bookmark); err != nil {
		return nil, ErrorUser(err.Error())
	}

	if err := s.Table.Add(bookmark); err != nil {
		log.PError("Error adding bookmark", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, ErrorFrom(err)
	}

	log.PInfo("Added new bookmark", map[string]interface{}{
		"id":    bookmark.ID,
		"url":   bookmark.URL,
		"name":  bookmark.Name,
		"index": bookmark.Index,
	})
	return &bookmark, nil
}

type updateBookmarkParameters struct {
	URL  string
	Name string
}

// Update will modify an existing bookmark
func (s *bookmarkStoreObject) Update(id string, params updateBookmarkParameters) (*Bookmark, *Error) {
	bookmark := s.Get(id)
	if bookmark == nil {
		return nil, ErrorUser("No bookmark with ID %s", id)
	}

	bookmark.URL = params.URL
	bookmark.Name = params.Name
	if err := limits.Check(bookmark); err != nil {
		return nil, ErrorUser(err.Error())
	}

	if err := s.Table.Update(*bookmark); err != nil {
		log.PError("Error updating bookmark", map[string]interface{}{
			"id":    id,
			"error": err.Error(),
		})
		return nil, ErrorFrom(err)
	}

	log.PInfo("Updated bookmark", map[string]interface{}{
		"id":    bookmark.ID,
		"url":   bookmark.URL,
		"name":  bookmark.Name,
		"index": bookmark.Index,
	})
	return bookmark, nil
}

// UpdateImage will update the image size for this bookmark
func (s *bookmarkStoreObject) UpdateImage(id string, size uint64, mimetype string) (*Bookmark, *Error) {
	bookmark := s.Get(id)
	if bookmark == nil {
		return nil, ErrorUser("No bookmark with ID %s", id)
	}

	bookmark.ImageSize = size
	bookmark.ImageMime = mimetype
	if err := s.Table.Update(*bookmark); err != nil {
		log.PError("Error updating bookmark", map[string]interface{}{
			"id":    id,
			"error": err.Error(),
		})
		return nil, ErrorFrom(err)
	}

	log.PInfo("Updated bookmark image", map[string]interface{}{
		"id":        bookmark.ID,
		"size":      size,
		"mime_type": mimetype,
	})
	return bookmark, nil
}

// Reorder will set the index of all bookmarks to reflect the map of ID->index
func (s *bookmarkStoreObject) Reorder(order map[string]int) *Error {
	if len(order) != len(s.AllBookmarks()) {
		return ErrorUser("All bookmarks must be specified to reorder")
	}

	for id := range order {
		bookmark := s.Get(id)
		if bookmark == nil {
			return ErrorUser("No bookmark with ID %s", id)
		}
	}

	for id, index := range order {
		bookmark := s.Get(id)
		bookmark.Index = index
		if err := s.Table.Update(*bookmark); err != nil {
			log.PError("Error updating bookmark", map[string]interface{}{
				"id":    id,
				"error": err.Error(),
			})
			return ErrorFrom(err)
		}

		log.PInfo("Updated bookmark image", map[string]interface{}{
			"id":    bookmark.ID,
			"index": index,
		})
	}

	return nil
}

// Delete will delete the bookmark with the given ID
func (s *bookmarkStoreObject) Delete(id string) *Error {
	bookmark := s.Get(id)
	if bookmark == nil {
		return ErrorUser("No bookmark with ID %s", id)
	}

	if err := s.Table.Delete(*bookmark); err != nil {
		log.PError("Error deleting bookmark", map[string]interface{}{
			"id":    id,
			"error": err.Error(),
		})
		return ErrorFrom(err)
	}

	log.PInfo("Deleted bookmark", map[string]interface{}{
		"id":    bookmark.ID,
		"url":   bookmark.URL,
		"name":  bookmark.Name,
		"index": bookmark.Index,
	})
	return nil
}

func (b *Bookmark) Picture() (io.ReadCloser, error) {
	picturePath := path.Join(Directories.Bookmarks, b.ID)
	f, err := os.OpenFile(picturePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening bookmark picture", map[string]interface{}{
			"id":    b.ID,
			"path":  picturePath,
			"error": err.Error(),
		})
		return nil, err
	}
	return f, nil
}
