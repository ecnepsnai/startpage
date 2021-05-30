package startpage

import (
	"os"
	"path"

	"github.com/ecnepsnai/web"
)

func (v viewHandle) Index(request web.Request, writer web.Writer) web.Response {
	file, err := os.OpenFile(path.Join(StaticPath, "index.html"), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return web.Response{
		Reader: file,
	}
}
