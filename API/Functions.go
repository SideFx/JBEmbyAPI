/////////////////////////////////////////////////////////////////////////////
// Name:        Functions.go
// Purpose:     Exported functions
// Author:      Jan Buchholz
// Created:     2026-04-15
/////////////////////////////////////////////////////////////////////////////

package API

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

func UserLoginToServer(secure bool, hostname string, port string, username string, password string) EmbyLogonResultExp {
	result := EmbyLogonResultExp{}
	url := createBasicURL(secure, hostname, port)
	result.Session.BaseUrl = url.BaseUrl
	result.Result = url.Result
	if url.Result.Code != NoErrorConst {
		return result
	}
	id := findUserIdByName(url.BaseUrl, username)
	result.Session.UserId = id.Id
	result.Result = id.Result
	if id.Result.Code != NoErrorConst {
		return result
	}
	tk := authenticateUserByCredentials(url.BaseUrl, id.Id, username, password)
	result.Session.AccessToken = tk.Token
	result.Result = tk.Result
	return result
}

func UserGetViews(baseurl string, userid string, accesstoken string) UserViewsExp {
	var views UserViewsExp
	p := rESTParams[QueryResultBaseItemDto]{}
	p.url = baseurl + getUserViews
	p.url = strings.Replace(p.url, "&1", userid, 1)
	p.url += "?" + apiKey + accesstoken
	p.genericHttpGet()
	if p.error.Code != NoErrorConst {
		views.Result = p.error
		return views
	}
	for _, item := range p.data.Items {
		for _, collectionType := range SupportedCollectionTypes {
			if item.CollectionType == collectionType {
				var v = UserView{
					Name:           item.Name,
					CollectionType: item.CollectionType,
					Id:             item.Id,
				}
				views.UserViews = append(views.UserViews, v)
			}
		}
	}
	views.Result = NoError
	return views
}

func UserGetMovies(baseurl string, collectionid string, userid string, accesstoken string) MoviesDataExp {
	MovieTable.Data = MovieData{}
	err := checkCollectionType(baseurl, userid, accesstoken, collectionid, MovieTable.CollectionType)
	if err.Code != NoErrorConst {
		return MoviesDataExp{MovieData{}, err}
	}
	items := userGetItems(baseurl, collectionid, MovieTable.CollectionType, userid, accesstoken)
	if items.Result.Code != NoErrorConst {
		return MoviesDataExp{MovieData{}, items.Result}
	}
	for _, item := range items.Items {
		var movie MovieDataInc
		movie.Name = item.Name
		movie.MovieId = item.Id
		movie.OriginalTitle = item.OriginalTitle
		movie.ProductionYear = strconv.Itoa(int(item.ProductionYear))
		movie.Studios = evalNameLongIdPairs(item.Studios)
		movie.Actors = evalPersons(item.People, ActorPersonType, GuestStarPersonType)
		movie.Directors = evalPersons(item.People, DirectorPersonType)
		movie.Genres = evalNameLongIdPairs(item.GenreItems)
		movie.Container = item.Container
		movie.AudioCodec, movie.VideoCodec = evalCodecs(item.MediaSources)
		movie.Resolution = evalResolution(item.Width, item.Height)
		movie.Bitrate = evalBitrate(item.Bitrate)
		movie.Runtime = evalRuntime(item.RunTimeTicks)
		movie.AddedAt = evalTime(item.DateCreated)
		movie.PrimaryImageId = item.PrimaryImageItemId
		if movie.PrimaryImageId == "" {
			movie.PrimaryImageId = item.Id
		}
		movie.PrimaryImageTag = item.PrimaryImageTag
		if movie.PrimaryImageTag == "" {
			movie.PrimaryImageTag = item.ImageTags[PrimaryImage]
		}
		movie.FileSize = evalFileSize(item.Size)
		movie.FileName = item.FileName
		movie.Overview = item.Overview
		movie.ImDBId = item.ProviderIds[ImDb]
		movie.TheMovieDBId = item.ProviderIds[TheMovieDb]
		MovieTable.Data.TMovieData = append(MovieTable.Data.TMovieData, movie)
	}
	return MoviesDataExp{MovieTable.Data, NoError}
}

func UserGetSeries(baseurl string, collectionid string, userid string, accesstoken string) SeriesDataExp {
	SeriesTable.Data = SeriesData{}
	err := checkCollectionType(baseurl, userid, accesstoken, collectionid, SeriesTable.CollectionType)
	if err.Code != NoErrorConst {
		return SeriesDataExp{SeriesData{}, err}
	}
	items := userGetItems(baseurl, collectionid, SeriesTable.CollectionType, userid, accesstoken)
	if items.Result.Code != NoErrorConst {
		return SeriesDataExp{SeriesData{}, items.Result}
	}
	for _, item := range items.Items {
		switch item.Type {
		case SeriesType:
			series := SeriesDataInc{}
			series.Name = item.Name
			series.OriginalTitle = item.OriginalTitle
			series.ProductionYear = strconv.Itoa(int(item.ProductionYear))
			series.Actors = evalPersons(item.People, ActorPersonType, GuestStarPersonType)
			series.Directors = evalPersons(item.People, DirectorPersonType)
			series.Genres = evalNameLongIdPairs(item.GenreItems)
			series.Studios = evalNameLongIdPairs(item.Studios)
			series.Overview = item.Overview
			series.AddedAt = evalTime(item.DateCreated)
			series.PrimaryImageId = item.PrimaryImageItemId
			if series.PrimaryImageId == "" {
				series.PrimaryImageId = item.Id
			}
			series.PrimaryImageTag = item.PrimaryImageTag
			if series.PrimaryImageTag == "" {
				series.PrimaryImageTag = item.ImageTags[PrimaryImage]
			}
			series.ImDBId = item.ProviderIds[ImDb]
			series.TheMovieDBId = item.ProviderIds[TheMovieDb]
			series.SeriesId = item.Id
			series.Type = item.Type
			SeriesTable.Data.TSeriesData = append(SeriesTable.Data.TSeriesData, series)
			break
		case SeasonType:
			season := SeasonDataInc{}
			season.Name = item.Name
			season.SeriesId = item.SeriesId
			season.SeasonId = item.Id
			season.ProductionYear = strconv.Itoa(int(item.ProductionYear))
			season.AddedAt = evalTime(item.DateCreated)
			season.PrimaryImageId = item.PrimaryImageItemId
			if season.PrimaryImageId == "" {
				season.PrimaryImageId = item.Id
			}
			season.PrimaryImageTag = item.PrimaryImageTag
			if season.PrimaryImageTag == "" {
				season.PrimaryImageTag = item.ImageTags[PrimaryImage]
			}
			season.SortIndex = item.IndexNumber
			season.Type = item.Type
			SeriesTable.Data.TSeasonData = append(SeriesTable.Data.TSeasonData, season)
			break
		case EpisodeType:
			episode := EpisodeDataInc{}
			episode.Name = item.Name
			episode.OriginalTitle = item.OriginalTitle
			episode.EpisodeId = item.Id
			episode.ProductionYear = strconv.Itoa(int(item.ProductionYear))
			episode.Actors = evalPersons(item.People, ActorPersonType, GuestStarPersonType)
			episode.Directors = evalPersons(item.People, DirectorPersonType)
			episode.Runtime = evalRuntime(item.RunTimeTicks)
			episode.Container = item.Container
			episode.AudioCodec, episode.VideoCodec = evalCodecs(item.MediaSources)
			episode.Resolution = evalResolution(item.Width, item.Height)
			episode.Bitrate = evalBitrate(item.Bitrate)
			episode.SortIndex = item.IndexNumber
			episode.FileSize = evalFileSize(item.Size)
			episode.FileName = item.FileName
			episode.Overview = item.Overview
			episode.AddedAt = evalTime(item.DateCreated)
			episode.PrimaryImageId = item.PrimaryImageItemId
			if episode.PrimaryImageId == "" {
				episode.PrimaryImageId = item.Id
			}
			episode.PrimaryImageTag = item.PrimaryImageTag
			if episode.PrimaryImageTag == "" {
				episode.PrimaryImageTag = item.ImageTags[PrimaryImage]
			}
			episode.ImDBId = item.ProviderIds[ImDb]
			episode.TheMovieDBId = item.ProviderIds[TheMovieDb]
			episode.SeriesId = item.SeriesId
			episode.SeasonId = item.SeasonId
			episode.Type = item.Type
			SeriesTable.Data.TEpisodeData = append(SeriesTable.Data.TEpisodeData, episode)
			break
		}
	}
	return SeriesDataExp{SeriesTable.Data, NoError}
}

func UserGetHomeVideos(baseurl string, collectionid string, userid string, accesstoken string) HomeVideosDataExp {
	HomeVideoTable.Data = HomeVideoData{}
	err := checkCollectionType(baseurl, userid, accesstoken, collectionid, HomeVideoTable.CollectionType)
	if err.Code != NoErrorConst {
		return HomeVideosDataExp{HomeVideoData{}, err}
	}
	items := userGetItems(baseurl, collectionid, HomeVideoTable.CollectionType, userid, accesstoken)
	if items.Result.Code != NoErrorConst {
		return HomeVideosDataExp{HomeVideoData{}, items.Result}
	}
	for _, item := range items.Items {
		switch item.Type {
		case VideoType:
			video := HomeVideoDataInc{}
			video.Name = item.Name
			video.ProductionYear = strconv.Itoa(int(item.ProductionYear))
			video.Genres = evalNameLongIdPairs(item.GenreItems)
			video.Overview = item.Overview
			video.Container = item.Container
			video.Resolution = evalResolution(item.Width, item.Height)
			video.AudioCodec, video.VideoCodec = evalCodecs(item.MediaSources)
			video.Runtime = evalRuntime(item.RunTimeTicks)
			video.Bitrate = evalBitrate(item.Bitrate)
			video.FileSize = evalFileSize(item.Size)
			video.FileName = item.FileName
			video.PrimaryImageId = item.PrimaryImageItemId
			if video.PrimaryImageId == "" {
				video.PrimaryImageId = item.Id
			}
			video.PrimaryImageTag = item.PrimaryImageTag
			if video.PrimaryImageTag == "" {
				video.PrimaryImageTag = item.ImageTags[PrimaryImage]
			}
			video.AddedAt = evalTime(item.DateCreated)
			video.FolderId = item.ParentId
			HomeVideoTable.Data.THomeVideoData = append(HomeVideoTable.Data.THomeVideoData, video)
			break
		case FolderType:
			folder := FolderDataInc{}
			folder.Name = item.Name
			folder.FolderId = item.Id
			HomeVideoTable.Data.TFolderData = append(HomeVideoTable.Data.TFolderData, folder)
			break
		}
	}
	return HomeVideosDataExp{HomeVideoTable.Data, NoError}
}

func UserGetMusicVideos(baseurl string, collectionid string, userid string, accesstoken string) MusicVideosDataExp {
	MusicVideoTable.Data = MusicVideoData{}
	err := checkCollectionType(baseurl, userid, accesstoken, collectionid, MusicVideoTable.CollectionType)
	if err.Code != NoErrorConst {
		return MusicVideosDataExp{MusicVideoData{}, err}
	}
	items := userGetItems(baseurl, collectionid, MusicVideoTable.CollectionType, userid, accesstoken)
	if items.Result.Code != NoErrorConst {
		return MusicVideosDataExp{MusicVideoData{}, items.Result}
	}
	for _, item := range items.Items {
		switch item.Type {
		case MusicVideoType:
			video := MusicVideoDataInc{}
			video.Name = item.Name
			video.ProductionYear = strconv.Itoa(int(item.ProductionYear))
			video.Genres = evalNameLongIdPairs(item.GenreItems)
			video.Container = item.Container
			video.Resolution = evalResolution(item.Width, item.Height)
			video.AudioCodec, video.VideoCodec = evalCodecs(item.MediaSources)
			video.Runtime = evalRuntime(item.RunTimeTicks)
			video.Bitrate = evalBitrate(item.Bitrate)
			video.AddedAt = evalTime(item.DateCreated)
			video.FileSize = evalFileSize(item.Size)
			video.FileName = item.FileName
			video.PrimaryImageId = item.PrimaryImageItemId
			if video.PrimaryImageId == "" {
				video.PrimaryImageId = item.Id
			}
			video.PrimaryImageTag = item.PrimaryImageTag
			if video.PrimaryImageTag == "" {
				video.PrimaryImageTag = item.ImageTags[PrimaryImage]
			}
			video.ImDBId = item.ProviderIds[ImDb]
			video.TheMovieDBId = item.ProviderIds[TheMovieDb]
			video.FolderId = item.ParentId
			MusicVideoTable.Data.TMusicVideoData = append(MusicVideoTable.Data.TMusicVideoData, video)
			break
		case FolderType:
			folder := FolderDataInc{}
			folder.Name = item.Name
			folder.FolderId = item.Id
			MusicVideoTable.Data.TFolderData = append(MusicVideoTable.Data.TFolderData, folder)
			break
		}
	}
	return MusicVideosDataExp{MusicVideoTable.Data, NoError}
}

func UserGetMusic(baseurl string, collectionid string, userid string, accesstoken string) MusicDataExp {
	MusicTable.Data = MusicData{}
	err := checkCollectionType(baseurl, userid, accesstoken, collectionid, MusicTable.CollectionType)
	if err.Code != NoErrorConst {
		return MusicDataExp{MusicData{}, err}
	}
	items := userGetItems(baseurl, collectionid, MusicTable.CollectionType, userid, accesstoken)
	if items.Result.Code != NoErrorConst {
		return MusicDataExp{MusicData{}, items.Result}
	}
	for _, item := range items.Items {
		switch item.Type {
		case AudioType:
			audio := AudioDataInc{}
			audio.Name = item.Name
			audio.ProductionYear = strconv.Itoa(int(item.ProductionYear))
			audio.Artists = evalNameIdPairs(item.ArtistItems)
			audio.Genres = evalNameLongIdPairs(item.GenreItems)
			audio.Container = item.Container
			audio.Album = item.Album
			audio.AlbumArtist, audio.AlbumArtistId = evalAlbumArtists(item.AlbumArtists)
			if audio.AlbumArtistId == "" {
				audio.AlbumArtist = item.AlbumArtist
				audio.AlbumArtistId = ""
			}
			audio.FileSize = evalFileSize(item.Size)
			audio.FileName = item.FileName
			audio.Bitrate = evalBitrate(item.Bitrate)
			audio.AudioCodec, _ = evalCodecs(item.MediaSources)
			audio.TrackNumber = strconv.Itoa(int(item.IndexNumber))
			audio.Runtime = evalRuntime(item.RunTimeTicks)
			audio.Type = item.Type
			audio.MediaType = item.MediaType
			audio.AddedAt = evalTime(item.DateCreated)
			audio.PrimaryImageId = item.PrimaryImageItemId
			if audio.PrimaryImageId == "" {
				audio.PrimaryImageId = item.Id
			}
			audio.PrimaryImageTag = item.PrimaryImageTag
			if audio.PrimaryImageTag == "" {
				audio.PrimaryImageTag = item.ImageTags[PrimaryImage]
			}
			audio.AudioId = item.Id
			audio.AlbumId = item.AlbumId
			MusicTable.Data.TAudioData = append(MusicTable.Data.TAudioData, audio)
			break
		case MusicAlbumType:
			album := AlbumDataInc{}
			album.Name = item.Name
			album.ProductionYear = strconv.Itoa(int(item.ProductionYear))
			album.Artists = evalNameIdPairs(item.ArtistItems)
			album.Genres = evalNameLongIdPairs(item.GenreItems)
			album.AlbumId = item.Id
			album.AlbumArtist, album.AlbumArtistId = evalAlbumArtists(item.AlbumArtists)
			if album.AlbumArtistId == "" {
				album.AlbumArtist = item.AlbumArtist
				album.AlbumArtistId = ""
			}
			album.Genres = evalNameLongIdPairs(item.GenreItems)
			album.Runtime = evalRuntime(item.RunTimeTicks)
			album.AddedAt = evalTime(item.DateCreated)
			album.PrimaryImageId = item.PrimaryImageItemId
			if album.PrimaryImageId == "" {
				album.PrimaryImageId = item.Id
			}
			album.PrimaryImageTag = item.PrimaryImageTag
			if album.PrimaryImageTag == "" {
				album.PrimaryImageTag = item.ImageTags[PrimaryImage]
			}
			album.MusicBrainzId = item.ProviderIds[MusicBrainzAlbum]
			album.Type = item.Type
			MusicTable.Data.TAlbumData = append(MusicTable.Data.TAlbumData, album)
			break
		}
	}
	return MusicDataExp{MusicTable.Data, NoError}
}

func GetPrimaryImageForItem(baseurl string, itemid string, format string, imagetag string,
	maxwidth int, maxheight int, accesstoken string) ItemImageExp {
	img := ItemImageExp{}
	url := baseurl + getPrimaryImage
	url = strings.Replace(url, "&1", itemid, 1)
	url += "?" + apiKey + accesstoken
	if format == ImageFormatBmp || format == ImageFormatGif || format == ImageFormatJpp || format == ImageFormatPng {
		url += "&" + paraFormat + string(format)
	}
	if maxwidth > 0 {
		url += "&" + paraMaxWidth + strconv.Itoa(maxwidth)
	}
	if maxheight > 0 {
		url += "&" + paraMaxHeight + strconv.Itoa(maxheight)
	}
	if imagetag != "" {
		url += "&" + paraImageTag + imagetag
	}
	response, err := http.Get(url)
	defer func() { _ = response.Body.Close() }()
	if err != nil {
		return ItemImageExp{itemid, nil, HttpGetFailed}
	}
	if response.StatusCode != http.StatusOK {
		return ItemImageExp{itemid, nil, HttpStatusError}
	}
	img.ImageData, err = io.ReadAll(response.Body)
	if err != nil {
		return ItemImageExp{itemid, nil, IoError}
	}
	img.Result = NoError
	return img
}

func Init() {
	sendNetworkBroadcast()
}
