package medailydeal

import (
	"io"
	"os"
	"time"
)

// DailyDeal describes a daily deal
type DailyDeal struct {
	SKU          string
	Title        string
	URL          string
	RegularPrice string
	SalePrice    string
	ImageSize    uint64
	ImagePath    string
	Expires      time.Time
}

// Open return a reader for the picture
func (p *DailyDeal) Open() (io.ReadCloser, error) {
	f, err := os.OpenFile(p.ImagePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.PError("Error opening picture file", map[string]interface{}{
			"file_path": p.ImagePath,
			"error":     err.Error(),
		})
		return nil, err
	}
	return f, nil
}
