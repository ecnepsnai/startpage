package startpage

// This file is was generated automatically by Codegen v1.6.0
// Do not make changes to this file as they will be lost

import (
	"path"

	"github.com/ecnepsnai/ds"
)

type bookmarkStoreObject struct {
	Table *ds.Table
}

// BookmarkStore the global bookmark store
var BookmarkStore = bookmarkStoreObject{}

func cbgenDataStoreRegisterBookmarkStore() {
	table, err := ds.Register(Bookmark{}, path.Join(Directories.Data, "bookmark.db"), &ds.Options{

		DisableSorting: true,
	})
	if err != nil {
		log.Fatal("Error registering bookmark store: %s", err.Error())
	}
	BookmarkStore.Table = table
}

// dataStoreSetup set up the data store
func dataStoreSetup() {

	cbgenDataStoreRegisterBookmarkStore()

}

// dataStoreTeardown tear down the data store
func dataStoreTeardown() {

	if BookmarkStore.Table != nil {
		BookmarkStore.Table.Close()
	}

}
