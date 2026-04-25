/////////////////////////////////////////////////////////////////////////////
// Name:        parser.hpp
// Purpose:     Json parser and data definitions
// Author:      Jan Buchholz (Header was mainly created by Copilot)
// Created:     2026-04-24
/////////////////////////////////////////////////////////////////////////////

#pragma once

#include "json.hpp"
using json = nlohmann::json;

// Supported collections (user views)
constexpr std::string_view CollectionMovies = "movies";
constexpr std::string_view CollectionSeries = "tvshows";
constexpr std::string_view CollectionHomeVideos = "homevideos";
constexpr std::string_view CollectionMusic = "music";
constexpr std::string_view CollectionMusicVideos = "musicvideos";

struct EmbyLogonResult {
    std::string baseUrl;
    std::string userId;
    std::string accessToken;
    int code = -1;
    std::string message;
};

inline EmbyLogonResult parseLogonResult(const std::string& raw) {
    json j = json::parse(raw);
    EmbyLogonResult r;
    r.code = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code == 0) {
        r.baseUrl     = j["Session"]["BaseUrl"].get<std::string>();
        r.userId      = j["Session"]["UserId"].get<std::string>();
        r.accessToken = j["Session"]["AccessToken"].get<std::string>();
    }
    return r;
}

struct UserView {
    std::string name;
    std::string collectionType;
    std::string id;
};

struct UserViewsResult {
    std::vector<UserView> views;
    int code = -1;
    std::string message;
};

inline UserViewsResult parseUserViews(const std::string& raw) {
    json j = json::parse(raw);
    UserViewsResult r;
    r.code = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code != 0) {
        return r;
    }
    for (auto& v : j["UserViews"]) {
        UserView uv;
        uv.name           = v["Name"].get<std::string>();
        uv.collectionType = v["CollectionType"].get<std::string>();
        uv.id             = v["Id"].get<std::string>();
        r.views.push_back(std::move(uv));
    }
    return r;
}

struct MovieDataInc {
    std::string name;
    std::string originalTitle;
    std::string productionYear;
    std::string runtime;
    std::vector<std::string> actors;
    std::vector<std::string> directors;
    std::vector<std::string> studios;
    std::vector<std::string> genres;
    std::string overview;
    std::string container;
    std::string audioCodec;
    std::string videoCodec;
    std::string resolution;
    std::string bitrate;
    std::string fileSize;
    std::string fileName;
    std::string addedAt;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string imDbId;
    std::string theMovieDbId;
    std::string movieId;
};

struct MovieData {
    std::vector<MovieDataInc> tMovieData;
};

struct MoviesDataImp {
    MovieData movies;
    int code = -1;
    std::string message;
};

inline MoviesDataImp parseMovies(const std::string& raw) {
    json j = json::parse(raw);
    MoviesDataImp r;
    r.code    = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code != 0) {
        return r;
    }
    // --- Movies ---
    for (auto& m : j["Movies"]["MovieData"]) {
        MovieDataInc md;
        md.name            = m["Name"].get<std::string>();
        md.originalTitle   = m["OriginalTitle"].get<std::string>();
        md.productionYear  = m["ProductionYear"].get<std::string>();
        md.runtime         = m["Runtime"].get<std::string>();
        md.actors          = m["Actors"].get<std::vector<std::string>>();
        md.directors       = m["Directors"].get<std::vector<std::string>>();
        md.studios         = m["Studios"].get<std::vector<std::string>>();
        md.genres          = m["Genres"].get<std::vector<std::string>>();
        md.overview        = m["Overview"].get<std::string>();
        md.container       = m["Container"].get<std::string>();
        md.audioCodec      = m["AudioCodec"].get<std::string>();
        md.videoCodec      = m["VideoCodec"].get<std::string>();
        md.resolution      = m["Resolution"].get<std::string>();
        md.bitrate         = m["Bitrate"].get<std::string>();
        md.fileSize        = m["FileSize"].get<std::string>();
        md.fileName        = m["FileName"].get<std::string>();
        md.addedAt         = m["AddedAt"].get<std::string>();
        md.primaryImageId  = m["PrimaryImageId"].get<std::string>();
        md.primaryImageTag = m["PrimaryImageTag"].get<std::string>();
        md.imDbId          = m["ImdbId"].get<std::string>();
        md.theMovieDbId    = m["TheMovieDbId"].get<std::string>();
        md.movieId         = m["MovieId"].get<std::string>();
        r.movies.tMovieData.push_back(std::move(md));
    }
    return r;
}

struct SeriesDataInc {
    std::string name;
    std::string originalTitle;
    std::string productionYear;
    std::vector<std::string> actors;
    std::vector<std::string> directors;
    std::vector<std::string> studios;
    std::vector<std::string> genres;
    std::string overview;
    std::string addedAt;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string imDbId;
    std::string theMovieDbId;
    std::string seriesId;
    std::string type;
};

struct SeasonDataInc {
    std::string name;
    std::string productionYear;
    std::string runtime;
    std::string addedAt;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string seriesId;
    std::string seasonId;
    std::string type;
    int sortIndex{};
};

struct EpisodeDataInc {
    std::string name;
    std::string originalTitle;
    std::string productionYear;
    std::string runtime;
    std::vector<std::string> actors;
    std::vector<std::string> directors;
    std::string overview;
    std::string container;
    std::string audioCodec;
    std::string videoCodec;
    std::string resolution;
    std::string bitrate;
    std::string fileSize;
    std::string fileName;
    std::string addedAt;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string imDbId;
    std::string theMovieDbId;
    std::string seriesId;
    std::string seasonId;
    std::string episodeId;
    std::string type;
    int sortIndex{};
};

struct SeriesData {
    std::vector<SeriesDataInc> tSeriesData;
    std::vector<SeasonDataInc> tSeasonData;
    std::vector<EpisodeDataInc> tEpisodeData;
};

struct SeriesDataImp {
    SeriesData series;
    int code = -1;
    std::string message;
};

inline SeriesDataImp parseSeries(const std::string& raw) {
    json j = json::parse(raw);
    SeriesDataImp r;
    r.code    = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code != 0) {
        return r;
    }
    // --- Series ---
    for (auto& s : j["Series"]["SeriesData"]) {
        SeriesDataInc sd;
        sd.name            = s["Name"].get<std::string>();
        sd.originalTitle   = s["OriginalTitle"].get<std::string>();
        sd.productionYear  = s["ProductionYear"].get<std::string>();
        sd.actors          = s["Actors"].get<std::vector<std::string>>();
        sd.directors       = s["Directors"].get<std::vector<std::string>>();
        sd.studios         = s["Studios"].get<std::vector<std::string>>();
        sd.genres          = s["Genres"].get<std::vector<std::string>>();
        sd.overview        = s["Overview"].get<std::string>();
        sd.addedAt         = s["AddedAt"].get<std::string>();
        sd.primaryImageId  = s["PrimaryImageId"].get<std::string>();
        sd.primaryImageTag = s["PrimaryImageTag"].get<std::string>();
        sd.imDbId          = s["ImdbId"].get<std::string>();
        sd.theMovieDbId    = s["TheMovieDbId"].get<std::string>();
        sd.seriesId        = s["SeriesId"].get<std::string>();
        sd.type            = s["Type"].get<std::string>();
        r.series.tSeriesData.push_back(std::move(sd));
    }
    // --- Seasons ---
    for (auto& s : j["Series"]["SeasonData"]) {
        SeasonDataInc sd;
        sd.name            = s["Name"].get<std::string>();
        sd.productionYear  = s["ProductionYear"].get<std::string>();
        sd.runtime         = s["Runtime"].get<std::string>();
        sd.addedAt         = s["AddedAt"].get<std::string>();
        sd.primaryImageId  = s["PrimaryImageId"].get<std::string>();
        sd.primaryImageTag = s["PrimaryImageTag"].get<std::string>();
        sd.seriesId        = s["SeriesId"].get<std::string>();
        sd.seasonId        = s["SeasonId"].get<std::string>();
        sd.type            = s["Type"].get<std::string>();
        sd.sortIndex       = s["SortIndex"].get<int>();
        r.series.tSeasonData.push_back(std::move(sd));
    }
    // --- Episodes ---
    for (auto& e : j["Series"]["EpisodeData"]) {
        EpisodeDataInc ed;
        ed.name            = e["Name"].get<std::string>();
        ed.originalTitle   = e["OriginalTitle"].get<std::string>();
        ed.productionYear  = e["ProductionYear"].get<std::string>();
        ed.runtime         = e["Runtime"].get<std::string>();
        ed.actors          = e["Actors"].get<std::vector<std::string>>();
        ed.directors       = e["Directors"].get<std::vector<std::string>>();
        ed.overview        = e["Overview"].get<std::string>();
        ed.container       = e["Container"].get<std::string>();
        ed.audioCodec      = e["AudioCodec"].get<std::string>();
        ed.videoCodec      = e["VideoCodec"].get<std::string>();
        ed.resolution      = e["Resolution"].get<std::string>();
        ed.bitrate         = e["Bitrate"].get<std::string>();
        ed.fileSize        = e["FileSize"].get<std::string>();
        ed.fileName        = e["FileName"].get<std::string>();
        ed.addedAt         = e["AddedAt"].get<std::string>();
        ed.primaryImageId  = e["PrimaryImageId"].get<std::string>();
        ed.primaryImageTag = e["PrimaryImageTag"].get<std::string>();
        ed.imDbId          = e["ImdbId"].get<std::string>();
        ed.theMovieDbId    = e["TheMovieDbId"].get<std::string>();
        ed.seriesId        = e["SeriesId"].get<std::string>();
        ed.seasonId        = e["SeasonId"].get<std::string>();
        ed.episodeId       = e["EpisodeId"].get<std::string>();
        ed.type            = e["Type"].get<std::string>();
        ed.sortIndex       = e["SortIndex"].get<int>();
        r.series.tEpisodeData.push_back(std::move(ed));
    }
    return r;
}

struct HomeVideoDataInc {
    std::string name;
    std::string productionYear;
    std::vector<std::string> genres;
    std::string runtime;
    std::string overview;
    std::string container;
    std::string audioCodec;
    std::string videoCodec;
    std::string resolution;
    std::string bitrate;
    std::string fileSize;
    std::string fileName;
    std::string addedAt;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string folderId;
    std::string videoId;
};

struct FolderDataInc {
    std::string name;
    std::string folderId;
};

struct HomeVideoData {
    std::vector<HomeVideoDataInc> tHomeVideoData;
    std::vector<FolderDataInc> tFolderData;
};

struct HomeVideosDataImp {
    HomeVideoData homeVideos;
    int code = -1;
    std::string message;
};

inline HomeVideosDataImp parseHomeVideos(const std::string& raw) {
    json j = json::parse(raw);
    HomeVideosDataImp r;
    r.code    = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code != 0) {
        return r;
    }
    // --- HomeVideoData ---
    for (auto& hv : j["HomeVideos"]["HomeVideoData"]) {
        HomeVideoDataInc h;
        h.name            = hv["Name"].get<std::string>();
        h.productionYear  = hv["ProductionYear"].get<std::string>();
        h.genres          = hv["Genres"].get<std::vector<std::string>>();
        h.runtime         = hv["Runtime"].get<std::string>();
        h.overview        = hv["Overview"].get<std::string>();
        h.container       = hv["Container"].get<std::string>();
        h.audioCodec      = hv["AudioCodec"].get<std::string>();
        h.videoCodec      = hv["VideoCodec"].get<std::string>();
        h.resolution      = hv["Resolution"].get<std::string>();
        h.bitrate         = hv["Bitrate"].get<std::string>();
        h.fileSize        = hv["FileSize"].get<std::string>();
        h.fileName        = hv["FileName"].get<std::string>();
        h.addedAt         = hv["AddedAt"].get<std::string>();
        h.primaryImageId  = hv["PrimaryImageId"].get<std::string>();
        h.primaryImageTag = hv["PrimaryImageTag"].get<std::string>();
        h.folderId        = hv["FolderId"].get<std::string>();
        h.videoId         = hv["VideoId"].get<std::string>();
        r.homeVideos.tHomeVideoData.push_back(std::move(h));
    }
    // --- FolderData ---
    for (auto& f : j["HomeVideos"]["FolderData"]) {
        FolderDataInc fd;
        fd.name     = f["Name"].get<std::string>();
        fd.folderId = f["FolderId"].get<std::string>();
        r.homeVideos.tFolderData.push_back(std::move(fd));
    }
    return r;
}

struct MusicVideoDataInc {
    std::string name;
    std::string productionYear;
    std::string runtime;
    std::vector<std::string> genres;
    std::string overview;
    std::string container;
    std::string audioCodec;
    std::string videoCodec;
    std::string resolution;
    std::string bitrate;
    std::string fileSize;
    std::string fileName;
    std::string addedAt;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string imDbId;
    std::string theMovieDbId;
    std::string movieId;
    std::string folderId;
};

struct MusicVideoData {
    std::vector<MusicVideoDataInc> tMusicVideoData;
    std::vector<FolderDataInc> tFolderData;
};

struct MusicVideosDataImp {
    MusicVideoData musicVideos;
    int code = -1;
    std::string message;
};

inline MusicVideosDataImp parseMusicVideos(const std::string& raw) {
    json j = json::parse(raw);
    MusicVideosDataImp r;
    r.code    = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code != 0) {
        return r;
    }
    // --- MusicVideoData ---
    for (auto& mv : j["MusicVideos"]["MusicVideoData"]) {
        MusicVideoDataInc m;
        m.name            = mv["Name"].get<std::string>();
        m.productionYear  = mv["ProductionYear"].get<std::string>();
        m.runtime         = mv["Runtime"].get<std::string>();
        m.genres          = mv["Genres"].get<std::vector<std::string>>();
        m.overview        = mv["Overview"].get<std::string>();
        m.container       = mv["Container"].get<std::string>();
        m.audioCodec      = mv["AudioCodec"].get<std::string>();
        m.videoCodec      = mv["VideoCodec"].get<std::string>();
        m.resolution      = mv["Resolution"].get<std::string>();
        m.bitrate         = mv["Bitrate"].get<std::string>();
        m.fileSize        = mv["FileSize"].get<std::string>();
        m.fileName        = mv["FileName"].get<std::string>();
        m.addedAt         = mv["AddedAt"].get<std::string>();
        m.primaryImageId  = mv["PrimaryImageId"].get<std::string>();
        m.primaryImageTag = mv["PrimaryImageTag"].get<std::string>();
        m.imDbId          = mv["ImdbId"].get<std::string>();
        m.theMovieDbId    = mv["TheMovieDbId"].get<std::string>();
        m.movieId         = mv["MovieId"].get<std::string>();
        m.folderId        = mv["FolderId"].get<std::string>();
        r.musicVideos.tMusicVideoData.push_back(std::move(m));
    }
    // --- FolderData ---
    for (auto& f : j["MusicVideos"]["FolderData"]) {
        FolderDataInc fd;
        fd.name     = f["Name"].get<std::string>();
        fd.folderId = f["FolderId"].get<std::string>();

        r.musicVideos.tFolderData.push_back(std::move(fd));
    }
    return r;
}

struct AlbumDataInc {
    std::string name;
    std::string productionYear;
    std::string albumArtist;
    std::string runtime;
    std::vector<std::string> artists;
    std::vector<std::string> genres;
    std::string addedAt;
    std::string albumId;
    std::string albumArtistId;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string musicBrainzId;
    std::string type;
};

struct AudioDataInc {
    std::string name;
    std::string productionYear;
    std::string trackNumber;
    std::string album;
    std::string albumArtist;
    std::string runtime;
    std::vector<std::string> artists;
    std::vector<std::string> genres;
    std::string container;
    std::string audioCodec;
    std::string bitrate;
    std::string addedAt;
    std::string fileSize;
    std::string fileName;
    std::string primaryImageId;
    std::string primaryImageTag;
    std::string audioId;
    std::string albumId;
    std::string albumArtistId;
    std::string mediaType;
    std::string type;
};

struct MusicData {
    std::vector<AlbumDataInc> tAlbumData;
    std::vector<AudioDataInc> tAudioData;
};

struct MusicDataImp {
    MusicData music;
    int code = -1;
    std::string message;
};

inline MusicDataImp parseMusic(const std::string& raw) {
    json j = json::parse(raw);
    MusicDataImp r;
    r.code    = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code != 0) {
        return r;
    }
    // --- AlbumData ---
    for (auto& a : j["Music"]["AlbumData"]) {
        AlbumDataInc ad;
        ad.name            = a["Name"].get<std::string>();
        ad.productionYear  = a["ProductionYear"].get<std::string>();
        ad.albumArtist     = a["AlbumArtist"].get<std::string>();
        ad.runtime         = a["Runtime"].get<std::string>();
        ad.artists         = a["Artists"].get<std::vector<std::string>>();
        ad.genres          = a["Genres"].get<std::vector<std::string>>();
        ad.addedAt         = a["AddedAt"].get<std::string>();
        ad.albumId         = a["AlbumId"].get<std::string>();
        ad.albumArtistId   = a["ArtistId"].get<std::string>();
        ad.primaryImageId  = a["PrimaryImageId"].get<std::string>();
        ad.primaryImageTag = a["PrimaryImageTag"].get<std::string>();
        ad.musicBrainzId   = a["MusicBrainzId"].get<std::string>();
        ad.type            = a["Type"].get<std::string>();
        r.music.tAlbumData.push_back(std::move(ad));
    }
    // --- AudioData ---
    for (auto& au : j["Music"]["AudioData"]) {
        AudioDataInc ad;
        ad.name            = au["Name"].get<std::string>();
        ad.productionYear  = au["ProductionYear"].get<std::string>();
        ad.trackNumber     = au["TrackNumber"].get<std::string>();
        ad.album           = au["Album"].get<std::string>();
        ad.albumArtist     = au["AlbumArtist"].get<std::string>();
        ad.runtime         = au["Runtime"].get<std::string>();
        ad.artists         = au["Artists"].get<std::vector<std::string>>();
        ad.genres          = au["Genres"].get<std::vector<std::string>>();
        ad.container       = au["Container"].get<std::string>();
        ad.audioCodec      = au["AudioCodec"].get<std::string>();
        ad.bitrate         = au["Bitrate"].get<std::string>();
        ad.addedAt         = au["AddedAt"].get<std::string>();
        ad.fileSize        = au["FileSize"].get<std::string>();
        ad.fileName        = au["FileName"].get<std::string>();
        ad.primaryImageId  = au["PrimaryImageId"].get<std::string>();
        ad.primaryImageTag = au["PrimaryImageTag"].get<std::string>();
        ad.audioId         = au["AudioId"].get<std::string>();
        ad.albumId         = au["AlbumId"].get<std::string>();
        ad.albumArtistId   = au["ArtistId"].get<std::string>();
        ad.mediaType       = au["MediaType"].get<std::string>();
        ad.type            = au["Type"].get<std::string>();
        r.music.tAudioData.push_back(std::move(ad));
    }
    return r;
}

struct ItemImageImp {
    std::string itemId;
    std::string imageData;  // Base64-String aus Go
    int code = -1;
    std::string message;
};

inline ItemImageImp parseItemImage(const std::string& raw) {
    json j = json::parse(raw);
    ItemImageImp r;
    r.code    = j["Result"]["Code"].get<int>();
    r.message = j["Result"]["Message"].get<std::string>();
    if (r.code != 0) {
        return r;
    }
    // --- ItemId for local buffering ---
    r.itemId = j["ItemId"].get<std::string>();
    r.imageData = j["ImageData"].get<std::string>(); // Base64
    return r;
}


