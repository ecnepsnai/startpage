package startpage

import (
	nanoid "github.com/matoous/go-nanoid"
)

func newID() string {
	id, err := nanoid.Generate("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890-_.", 12)
	if err != nil {
		log.Fatal("Error generating nanoid: %s", err.Error())
	}
	return id
}
