package main

import (
	"JBEmbyAPI/API"
	"fmt"
)

func main() {
	log := API.UserLoginToServer(false,
		"<hostname>",
		"<port>",
		"<username>",
		"<password>")
	if log.Result.Code != API.NoErrorConst {
		panic(log.Result.Message)
	}
	views := API.UserGetViews(log.Session.BaseUrl, log.Session.UserId, log.Session.AccessToken)
	for _, view := range views.UserViews {
		/*
			if view.CollectionType == API.CollectionMovies {
				movies := API.UserGetMovies(log.Session.BaseUrl, view.Id, log.Session.UserId, log.Session.AccessToken)
				if movies.Result.Code == API.NoErrorConst {
					for _, movie := range movies.Movies.TMovieData {
						fmt.Println(movie.Name, movie.ProductionYear, movie.ImDBId, movie.TheMovieDBId,
							movie.AudioCodec, movie.VideoCodec, movie.Resolution, movie.AddedAt, movie.FileName)
					}
				} else {
					fmt.Println(movies.Result.Message)
				}
			}
		*/
		/*
			if view.CollectionType == API.CollectionSeries {
				series := API.UserGetSeries(log.Session.BaseUrl, view.Id, log.Session.UserId, log.Session.AccessToken,
					maxactors, maxdirectors, maxgenres, maxstudios)
				if series.Result.Code == API.NoErrorConst {
					for _, episode := range series.Series.TEpisodeData {
						fmt.Println(episode.Name, episode.FileSize, episode.Bitrate, episode.Container, episode.AudioCodec,
							episode.VideoCodec)
					}
				} else {
					fmt.Println(series.Result.Message)
				}
			}
		*/
		/*
			if view.CollectionType == API.CollectionHomeVideos {
				videos := API.UserGetHomeVideos(log.Session.BaseUrl, view.Id, log.Session.UserId, log.Session.AccessToken)
				if videos.Result.Code == API.NoErrorConst {
					for _, video := range videos.HomeVideos.THomeVideoData {
						fmt.Println(video.Name, video.FileSize, video.Bitrate, video.Container, video.AudioCodec,
							video.VideoCodec, video.Resolution)
					}
					for _, folder := range videos.HomeVideos.TFolderData {
						fmt.Println(folder.Name, folder.FolderId)
					}
				} else {
					fmt.Println(videos.Result.Message)
				}
			}
		*/
		/*
			if view.CollectionType == API.CollectionMusicVideos {
				videos := API.UserGetMusicVideos(log.Session.BaseUrl, view.Id, log.Session.UserId, log.Session.AccessToken)
				if videos.Result.Code == API.NoErrorConst {
					for _, video := range videos.MusicVideos.TMusicVideoData {
						fmt.Println(video.Name, video.FileSize, video.Bitrate, video.Container, video.AudioCodec,
							video.VideoCodec, video.Resolution)
					}
					for _, folder := range videos.MusicVideos.TFolderData {
						fmt.Println(folder.Name, folder.FolderId)
					}
				} else {
					fmt.Println(videos.Result.Message)
				}
			}
		*/
		if view.CollectionType == API.CollectionMusic {
			music := API.UserGetMusic(log.Session.BaseUrl, view.Id, log.Session.UserId, log.Session.AccessToken)
			if music.Result.Code == API.NoErrorConst {
				for _, m := range music.Music.TAlbumData {
					fmt.Println(m.Name, m.AlbumArtist, m.PrimaryImageTag)
					if m.AlbumArtist != "" && m.PrimaryImageTag != "" {
						image := API.GetPrimaryImageForItem(log.Session.BaseUrl, m.PrimaryImageId, "jpg",
							m.PrimaryImageTag, 300, 300, log.Session.AccessToken)
						if image.Result.Code == API.NoErrorConst {
							fmt.Println("Image retrieved successfully")
						} else {
							fmt.Println(image.Result.Message)
						}
					}
				}
			} else {
				fmt.Println(music.Result.Message)
			}
		}
	}
}
