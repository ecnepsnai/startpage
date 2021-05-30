package startpage

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/ecnepsnai/web"
)

func (a *apiHandle) BookmarkGetAll(request web.Request) (interface{}, *web.Error) {
	return BookmarkStore.AllBookmarks(), nil
}

func (a *apiHandle) BookmarkNew(request web.Request) (interface{}, *web.Error) {
	params := newBookmarkParameters{}
	if err := request.Decode(&params); err != nil {
		return nil, err
	}

	bookmark, err := BookmarkStore.New(params)
	if err != nil {
		if err.Server {
			return nil, web.CommonErrors.ServerError
		}
		return nil, web.ValidationError(err.Message)
	}

	return bookmark, nil
}

func (a *apiHandle) BookmarkGet(request web.Request) (interface{}, *web.Error) {
	id := request.Params.ByName("id")
	bookmark := BookmarkStore.Get(id)
	if bookmark == nil {
		return nil, web.ValidationError("No bookmark with ID %s", id)
	}
	return bookmark, nil
}

func (v *viewHandle) BookmarkGetPicture(request web.Request, writer web.Writer) web.Response {
	id := request.Params.ByName("id")
	bookmark := BookmarkStore.Get(id)
	if bookmark == nil {
		return web.Response{
			Status: 404,
		}
	}

	r, err := bookmark.Picture()
	if err != nil {
		return web.Response{
			Status: 500,
		}
	}

	return web.Response{
		Headers: map[string]string{
			"Content-Length": fmt.Sprintf("%d", bookmark.ImageSize),
			"Content-Type":   bookmark.ImageMime,
		},
		Reader: r,
	}
}

func (a *apiHandle) BookmarkEdit(request web.Request) (interface{}, *web.Error) {
	id := request.Params.ByName("id")
	params := updateBookmarkParameters{}
	if err := request.Decode(&params); err != nil {
		return nil, err
	}
	bookmark, err := BookmarkStore.Update(id, params)
	if err != nil {
		if err.Server {
			return nil, web.CommonErrors.ServerError
		}
		return nil, web.ValidationError(err.Message)
	}

	return bookmark, nil
}

func (a *apiHandle) BookmarkEditPicture(request web.Request) (interface{}, *web.Error) {
	id := request.Params.ByName("id")
	bookmark := BookmarkStore.Get(id)
	if bookmark == nil {
		return nil, web.ValidationError("No bookmark with ID %s", id)
	}

	file, info, erro := request.HTTP.FormFile("picture")
	if erro != nil {
		return nil, web.ValidationError("No picture field")
	}
	size := uint64(info.Size)
	mime := info.Header.Get("Content-Type")

	picturePath := path.Join(Directories.Bookmarks, id)
	f, erro := os.OpenFile(picturePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if erro != nil {
		log.PError("Error opening bookmark picture", map[string]interface{}{
			"id":    id,
			"path":  picturePath,
			"error": erro.Error(),
		})
		return nil, web.CommonErrors.ServerError
	}
	if _, err := io.Copy(f, file); err != nil {
		log.PError("Error writing bookmark picture", map[string]interface{}{
			"id":    id,
			"path":  picturePath,
			"error": err.Error(),
		})
		return nil, web.CommonErrors.ServerError
	}

	bookmark, err := BookmarkStore.UpdateImage(id, size, mime)
	if err != nil {
		if err.Server {
			return nil, web.CommonErrors.ServerError
		}
		return nil, web.ValidationError(err.Message)
	}

	return bookmark, nil
}

func (a *apiHandle) BookmarkDelete(request web.Request) (interface{}, *web.Error) {
	id := request.Params.ByName("id")
	err := BookmarkStore.Delete(id)
	if err != nil {
		if err.Server {
			return nil, web.CommonErrors.ServerError
		}
		return nil, web.ValidationError(err.Message)
	}

	return true, nil
}
