/////////////////////////////////////////////////////////////////////////////
// Name:        Constants.go
// Purpose:     Constants for JBEmbyAPI (REST API)
// Author:      Jan Buchholz
// Created:     2026-04-13
/////////////////////////////////////////////////////////////////////////////

package API

// Protocols
const (
	ProtocolHttp  = "http"
	ProtocolHttps = "https"
)

// REST Header Constants
const (
	contentType     = "Content-Type"
	contentTypeJSON = "application/json"
)

// URL parameters
const (
	paraParentId         = "ParentId="
	paraRecursive        = "Recursive="
	paraIncludeItemTypes = "IncludeItemTypes="
	paraFields           = "Fields="
	paraFormat           = "format="
	paraMaxWidth         = "MaxWidth="
	paraMaxHeight        = "MaxHeight="
	paraImageTag         = "tag="
	apiKey               = "api_key="
)

// Parameters for the Authentication Request Header
const (
	authType        = "Emby"
	authHeader      = "Authorization"
	authKeyUserId   = "UserId"
	authKeyClient   = "Client"
	authKeyDevice   = "Device"
	authKeyDeviceId = "DeviceId"
	authKeyVersion  = "Version"
	client          = "EmbyExplorer"
)

// REST Endpoints
const (
	getUsersPublic       = "/Users/Public"
	postAuthenticateUser = "/Users/AuthenticateByName"
	getUserViews         = "/Users/" + "&1" + "/Views"
	getUserItems         = "/Users/" + "&1" + "/Items"
	getPrimaryImage      = "/Items/" + "&1" + "/Images/Primary"
)

// Collection Types
const (
	CollectionMovies      = "movies"
	CollectionSeries      = "tvshows"
	CollectionHomeVideos  = "homevideos"
	CollectionMusic       = "music"
	CollectionMusicVideos = "musicvideos"
)

var SupportedCollectionTypes = []string{
	CollectionMovies,
	CollectionSeries,
	CollectionHomeVideos,
	CollectionMusicVideos,
	CollectionMusic,
}

const (
	VideoType       = "Video"
	SeriesType      = "Series"
	SeasonType      = "Season"
	EpisodeType     = "Episode"
	MovieType       = "Movie"
	FolderType      = "Folder"
	AudioType       = "Audio"
	MusicVideoType  = "MusicVideo"
	MusicAlbumType  = "MusicAlbum"
	MusicArtistType = "MusicArtist"
)

const (
	TheMovieDb       = "Tmdb"
	ImDb             = "IMDB"
	MusicBrainzAlbum = "MusicBrainzAlbum"
)

const PrimaryImage = "Primary"

const (
	ImageFormatBmp string = "bmp"
	ImageFormatGif string = "gif"
	ImageFormatJpp string = "jpp"
	ImageFormatPng string = "png"
)
