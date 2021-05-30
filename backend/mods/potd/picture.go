package potd

import (
	"io"
	"os"
	"time"
)

// Picture describes a picture object
type Picture struct {
	URL       string
	Path      string
	Size      uint64
	StartDate time.Time
	EndDate   time.Time
}

// Open return a reader for the picture
func (p *Picture) Open() (io.ReadCloser, error) {
	f, err := os.OpenFile(p.Path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening picture file", map[string]interface{}{
			"file_path": p.Path,
			"error":     err.Error(),
		})
		return nil, err
	}
	return f, nil
}
