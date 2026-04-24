package API

type DataDescription[T any] struct {
	CollectionType   string
	IncludeItemTypes []string
	APIFields        []string
	Data             T
}

type MovieDataInc struct {
	Name            string   `json:"Name"`
	OriginalTitle   string   `json:"OriginalTitle"`
	ProductionYear  string   `json:"ProductionYear"`
	Runtime         string   `json:"Runtime"`
	Actors          []string `json:"Actors"`
	Directors       []string `json:"Directors"`
	Studios         []string `json:"Studios"`
	Genres          []string `json:"Genres"`
	Overview        string   `json:"Overview"`
	Container       string   `json:"Container"`
	AudioCodec      string   `json:"AudioCodec"`
	VideoCodec      string   `json:"VideoCodec"`
	Resolution      string   `json:"Resolution"`
	Bitrate         string   `json:"Bitrate"`
	FileSize        string   `json:"FileSize"`
	FileName        string   `json:"FileName"`
	AddedAt         string   `json:"AddedAt"`
	PrimaryImageId  string   `json:"PrimaryImageId"`
	PrimaryImageTag string   `json:"PrimaryImageTag"`
	ImDBId          string   `json:"ImdbId"`
	TheMovieDBId    string   `json:"TheMovieDbId"`
	MovieId         string   `json:"MovieId"`
}
type MovieData struct {
	TMovieData []MovieDataInc `json:"MovieData"`
}

var MovieTable = DataDescription[MovieData]{
	CollectionMovies,
	[]string{MovieType},
	[]string{"Name", "OriginalTitle", "MediaSources", "FileName", "Genres", "ProductionYear",
		"People", "Studios", "Width", "Height", "Container", "DateCreated", "Overview", "RunTimeTicks",
		"Id", "Type", "ProviderIds"},
	MovieData{},
}

type SeriesDataInc struct {
	Name            string   `json:"Name"`
	OriginalTitle   string   `json:"OriginalTitle"`
	ProductionYear  string   `json:"ProductionYear"`
	Actors          []string `json:"Actors"`
	Directors       []string `json:"Directors"`
	Studios         []string `json:"Studios"`
	Genres          []string `json:"Genres"`
	Overview        string   `json:"Overview"`
	AddedAt         string   `json:"AddedAt"`
	PrimaryImageId  string   `json:"PrimaryImageId"`
	PrimaryImageTag string   `json:"PrimaryImageTag"`
	ImDBId          string   `json:"ImdbId"`
	TheMovieDBId    string   `json:"TheMovieDbId"`
	SeriesId        string   `json:"SeriesId"`
	Type            string   `json:"Type"`
}

type SeasonDataInc struct {
	Name            string `json:"Name"`
	ProductionYear  string `json:"ProductionYear"`
	Runtime         string `json:"Runtime"`
	AddedAt         string `json:"AddedAt"`
	PrimaryImageId  string `json:"PrimaryImageId"`
	PrimaryImageTag string `json:"PrimaryImageTag"`
	SeriesId        string `json:"SeriesId"`
	SeasonId        string `json:"SeasonId"`
	Type            string `json:"Type"`
	SortIndex       int32  `json:"SortIndex"`
}

type EpisodeDataInc struct {
	Name            string   `json:"Name"`
	OriginalTitle   string   `json:"OriginalTitle"`
	ProductionYear  string   `json:"ProductionYear"`
	Runtime         string   `json:"Runtime"`
	Actors          []string `json:"Actors"`
	Directors       []string `json:"Directors"`
	Overview        string   `json:"Overview"`
	Container       string   `json:"Container"`
	AudioCodec      string   `json:"AudioCodec"`
	VideoCodec      string   `json:"VideoCodec"`
	Resolution      string   `json:"Resolution"`
	Bitrate         string   `json:"Bitrate"`
	FileSize        string   `json:"FileSize"`
	FileName        string   `json:"FileName"`
	AddedAt         string   `json:"AddedAt"`
	PrimaryImageId  string   `json:"PrimaryImageId"`
	PrimaryImageTag string   `json:"PrimaryImageTag"`
	ImDBId          string   `json:"ImdbId"`
	TheMovieDBId    string   `json:"TheMovieDbId"`
	SeriesId        string   `json:"SeriesId"`
	SeasonId        string   `json:"SeasonId"`
	EpisodeId       string   `json:"EpisodeId"`
	Type            string   `json:"Type"`
	SortIndex       int32    `json:"SortIndex"`
}

type SeriesData struct {
	TSeriesData  []SeriesDataInc  `json:"SeriesData"`
	TSeasonData  []SeasonDataInc  `json:"SeasonData"`
	TEpisodeData []EpisodeDataInc `json:"EpisodeData"`
}

var SeriesTable = DataDescription[SeriesData]{
	CollectionSeries,
	[]string{SeriesType, SeasonType, EpisodeType},
	[]string{"Name", "MediaSources", "FileName", "Genres", "ProductionYear", "People",
		"Studios", "Width", "Height", "Container", "RunTimeTicks", "Overview", "DateCreated",
		"SeriesId", "SeasonId", "Id", "ParentId", "IndexNumber", "Type", "ProviderIds"},
	SeriesData{},
}

type HomeVideoDataInc struct {
	Name            string   `json:"Name"`
	ProductionYear  string   `json:"ProductionYear"`
	Genres          []string `json:"Genres"`
	Runtime         string   `json:"Runtime"`
	Overview        string   `json:"Overview"`
	Container       string   `json:"Container"`
	AudioCodec      string   `json:"AudioCodec"`
	VideoCodec      string   `json:"VideoCodec"`
	Resolution      string   `json:"Resolution"`
	Bitrate         string   `json:"Bitrate"`
	FileSize        string   `json:"FileSize"`
	FileName        string   `json:"FileName"`
	AddedAt         string   `json:"AddedAt"`
	PrimaryImageId  string   `json:"PrimaryImageId"`
	PrimaryImageTag string   `json:"PrimaryImageTag"`
	FolderId        string   `json:"FolderId"`
	VideoId         string   `json:"VideoId"`
}

type FolderDataInc struct {
	Name     string `json:"Name"`
	FolderId string `json:"FolderId"`
}

type HomeVideoData struct {
	THomeVideoData []HomeVideoDataInc `json:"HomeVideoData"`
	TFolderData    []FolderDataInc    `json:"FolderData"`
}

var HomeVideoTable = DataDescription[HomeVideoData]{
	CollectionHomeVideos,
	[]string{VideoType, FolderType},
	[]string{"Name", "MediaSources", "FileName", "Genres", "Width", "Height", "Container",
		"Genres", "ProductionYear", "RunTimeTicks", "DateCreated", "Id", "ParentId", "Type"},
	HomeVideoData{},
}

type MusicVideoDataInc struct {
	Name            string   `json:"Name"`
	ProductionYear  string   `json:"ProductionYear"`
	Runtime         string   `json:"Runtime"`
	Genres          []string `json:"Genres"`
	Overview        string   `json:"Overview"`
	Container       string   `json:"Container"`
	AudioCodec      string   `json:"AudioCodec"`
	VideoCodec      string   `json:"VideoCodec"`
	Resolution      string   `json:"Resolution"`
	Bitrate         string   `json:"Bitrate"`
	FileSize        string   `json:"FileSize"`
	FileName        string   `json:"FileName"`
	AddedAt         string   `json:"AddedAt"`
	PrimaryImageId  string   `json:"PrimaryImageId"`
	PrimaryImageTag string   `json:"PrimaryImageTag"`
	ImDBId          string   `json:"ImdbId"`
	TheMovieDBId    string   `json:"TheMovieDbId"`
	MovieId         string   `json:"MovieId"`
	FolderId        string   `json:"FolderId"`
}

type MusicVideoData struct {
	TMusicVideoData []MusicVideoDataInc `json:"MusicVideoData"`
	TFolderData     []FolderDataInc     `json:"FolderData"`
}

var MusicVideoTable = DataDescription[MusicVideoData]{
	CollectionMusicVideos,
	[]string{MusicVideoType, FolderType},
	[]string{"Name", "MediaSources", "FileName", "Genres", "ProductionYear", "DateCreated",
		"Width", "Height", "Container", "Overview", "RunTimeTicks", "Id", "ParentId", "Type", "ProviderIds"},
	MusicVideoData{},
}

type AlbumDataInc struct {
	Name            string   `json:"Name"`
	ProductionYear  string   `json:"ProductionYear"`
	AlbumArtist     string   `json:"AlbumArtist"`
	Runtime         string   `json:"Runtime"`
	Artists         []string `json:"Artists"`
	Genres          []string `json:"Genres"`
	AddedAt         string   `json:"AddedAt"`
	AlbumId         string   `json:"AlbumId"`
	AlbumArtistId   string   `json:"ArtistId"`
	PrimaryImageId  string   `json:"PrimaryImageId"`
	PrimaryImageTag string   `json:"PrimaryImageTag"`
	MusicBrainzId   string   `json:"MusicBrainzId"`
	Type            string   `json:"Type"`
}

type AudioDataInc struct {
	Name            string   `json:"Name"`
	ProductionYear  string   `json:"ProductionYear"`
	TrackNumber     string   `json:"TrackNumber"`
	Album           string   `json:"Album"`
	AlbumArtist     string   `json:"AlbumArtist"`
	Runtime         string   `json:"Runtime"`
	Artists         []string `json:"Artists"`
	Genres          []string `json:"Genres"`
	Container       string   `json:"Container"`
	AudioCodec      string   `json:"AudioCodec"`
	Bitrate         string   `json:"Bitrate"`
	AddedAt         string   `json:"AddedAt"`
	FileSize        string   `json:"FileSize"`
	FileName        string   `json:"FileName"`
	PrimaryImageId  string   `json:"PrimaryImageId"`
	PrimaryImageTag string   `json:"PrimaryImageTag"`
	AudioId         string   `json:"AudioId"`
	AlbumId         string   `json:"AlbumId"`
	AlbumArtistId   string   `json:"ArtistId"`
	MediaType       string   `json:"MediaType"`
	Type            string   `json:"Type"`
}

type MusicData struct {
	TAlbumData []AlbumDataInc `json:"AlbumData"`
	TAudioData []AudioDataInc `json:"AudioData"`
}

var MusicTable = DataDescription[MusicData]{
	CollectionMusic,
	[]string{MusicAlbumType, AudioType},
	[]string{"Name", "MediaSources", "FileName", "Genres", "ProductionYear", "People", "Studios",
		"Container", "Overview", "RunTimeTicks", "Id", "ParentId", "Type", "DateCreated", "ProviderIds"},
	MusicData{},
}

type MoviesDataExp struct {
	Movies MovieData   `json:"Movies"`
	Result ErrorStruct `json:"Result"`
}

type SeriesDataExp struct {
	Series SeriesData  `json:"Series"`
	Result ErrorStruct `json:"Result"`
}

type HomeVideosDataExp struct {
	HomeVideos HomeVideoData `json:"HomeVideos"`
	Result     ErrorStruct   `json:"Result"`
}

type MusicVideosDataExp struct {
	MusicVideos MusicVideoData `json:"MusicVideos"`
	Result      ErrorStruct    `json:"Result"`
}

type MusicDataExp struct {
	Music  MusicData   `json:"Music"`
	Result ErrorStruct `json:"Result"`
}
