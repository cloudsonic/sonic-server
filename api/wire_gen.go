// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package api

import (
	"github.com/cloudsonic/sonic-server/itunesbridge"
	"github.com/cloudsonic/sonic-server/persistence/db_ledis"
	"github.com/deluan/gomate"
	"github.com/deluan/gomate/ledis"
	"github.com/google/wire"
)

// Injectors from wire_injectors.go:

func initSystemController(router *Router) *SystemController {
	systemController := NewSystemController()
	return systemController
}

func initBrowsingController(router *Router) *BrowsingController {
	browser := router.Browser
	browsingController := NewBrowsingController(browser)
	return browsingController
}

func initAlbumListController(router *Router) *AlbumListController {
	listGenerator := router.ListGenerator
	albumListController := NewAlbumListController(listGenerator)
	return albumListController
}

func initMediaAnnotationController(router *Router) *MediaAnnotationController {
	scrobbler := router.Scrobbler
	ratings := router.Ratings
	mediaAnnotationController := NewMediaAnnotationController(scrobbler, ratings)
	return mediaAnnotationController
}

func initPlaylistsController(router *Router) *PlaylistsController {
	playlists := router.Playlists
	playlistsController := NewPlaylistsController(playlists)
	return playlistsController
}

func initSearchingController(router *Router) *SearchingController {
	search := router.Search
	searchingController := NewSearchingController(search)
	return searchingController
}

func initUsersController(router *Router) *UsersController {
	usersController := NewUsersController()
	return usersController
}

func initMediaRetrievalController(router *Router) *MediaRetrievalController {
	cover := router.Cover
	mediaRetrievalController := NewMediaRetrievalController(cover)
	return mediaRetrievalController
}

func initStreamController(router *Router) *StreamController {
	browser := router.Browser
	streamController := NewStreamController(browser)
	return streamController
}

// wire_injectors.go:

var allProviders = wire.NewSet(itunesbridge.NewItunesControl, NewSystemController,
	NewBrowsingController,
	NewAlbumListController,
	NewMediaAnnotationController,
	NewPlaylistsController,
	NewSearchingController,
	NewUsersController,
	NewMediaRetrievalController,
	NewStreamController,
	newDB, wire.FieldsOf(new(*Router), "Browser", "Cover", "ListGenerator", "Playlists", "Ratings", "Scrobbler", "Search"),
)

func newDB() gomate.DB {
	return ledis.NewEmbeddedDB(db_ledis.Db())
}
