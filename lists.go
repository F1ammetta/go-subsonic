package subsonic

import (
	"errors"
	"fmt"
)

func validateListType(input string) bool {
	validTypes := map[string]bool{
		"random":               true,
		"newest":               true,
		"highest":              true,
		"frequent":             true,
		"recent":               true,
		"alphabeticalByName":   true,
		"alphabeticalByArtist": true,
		"starred":              true,
		"byYear":               true,
		"byGenre":              true,
	}
	_, ok := validTypes[input]
	return ok
}

// GetAlbumList returns a list of random, newest, highest rated etc. albums. Similar to the album lists on the home page of the Subsonic web interface.
// Optional Parameters:
// size          No              10      The number of albums to return. Max 500.
// offset        No              0       The list offset. Useful if you for example want to page through the list of newest albums.
// fromYear      Yes (if type is         The first year in the range. If fromYear > toYear a reverse chronological list is returned.
//               byYear)
// toYear        Yes (if type is         The last year in the range.
//               byYear)
// genre         Yes (if type is         The name of the genre, e.g., "Rock".
//               byGenre)
// musicFolderId No                      (Since 1.11.0) Only return albums in the music folder with the given ID. See getMusicFolders.
func (s *SubsonicClient) GetAlbumList(listType string, parameters map[string]string) ([]*Album, error) {
	if !validateListType(listType) {
		return nil, fmt.Errorf("List type %s is invalid, see http://www.subsonic.org/pages/api.jsp#getAlbumList", listType)
	}
	if listType == "byYear" {
		_, ok := parameters["fromYear"]
		if !ok {
			return nil, errors.New("Required argument fromYear was not found when using GetAlbumList byYear")
		}
		_, ok = parameters["toYear"]
		if !ok {
			return nil, errors.New("Required argument toYear was not found when using GetAlbumList byYear")
		}
	} else if listType == "byGenre" {
		_, ok := parameters["genre"]
		if !ok {
			return nil, errors.New("Required argument genre was not found when using GetAlbumList byGenre")
		}
	}
	params := make(map[string]string)
	params["type"] = listType
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getAlbumList", params)
	if err != nil {
		return nil, err
	}
	return resp.AlbumList.Albums, nil
}

// GetAlbumList2 returns a list of albums like GetAlbumList, but organized according to id3 tags.
// Optional Parameters:
// size          No              10      The number of albums to return. Max 500.
// offset        No              0       The list offset. Useful if you for example want to page through the list of newest albums.
// fromYear      Yes (if type is         The first year in the range. If fromYear > toYear a reverse chronological list is returned.
//               byYear)
// toYear        Yes (if type is         The last year in the range.
//               byYear)
// genre         Yes (if type is         The name of the genre, e.g., "Rock".
//               byGenre)
// musicFolderId No                      (Since 1.11.0) Only return albums in the music folder with the given ID. See getMusicFolders.
func (s *SubsonicClient) GetAlbumList2(listType string, parameters map[string]string) ([]*Album, error) {
	if !validateListType(listType) {
		return nil, fmt.Errorf("List type %s is invalid, see http://www.subsonic.org/pages/api.jsp#getAlbumList", listType)
	}
	if listType == "byYear" {
		_, ok := parameters["fromYear"]
		if !ok {
			return nil, errors.New("Required argument fromYear was not found when using GetAlbumList2 byYear")
		}
		_, ok = parameters["toYear"]
		if !ok {
			return nil, errors.New("Required argument toYear was not found when using GetAlbumList2 byYear")
		}
	} else if listType == "byGenre" {
		_, ok := parameters["genre"]
		if !ok {
			return nil, errors.New("Required argument genre was not found when using GetAlbumList2 byGenre")
		}
	}
	params := make(map[string]string)
	params["type"] = listType
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getAlbumList2", params)
	if err != nil {
		return nil, err
	}
	return resp.AlbumList2.Albums, nil
}

// GetRandomSongs returns a randomly selected set of songs limited by the optional parameters.
// Optional Parameters:
// * size:           The maximum number of songs to return. Max 500, default 10.
// * genre:          Only returns songs belonging to this genre.
// * fromYear:       Only return songs published after or in this year.
// * toYear:         Only return songs published before or in this year.
// * musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *SubsonicClient) GetRandomSongs(parameters map[string]string) ([]*Song, error) {
	resp, err := s.Get("getRandomSongs", parameters)
	if err != nil {
		return nil, err
	}
	return resp.RandomSongs.Songs, nil
}

// GetSongsByGenre returns songs in a given genre name.
// Optional Parameters:
// * count:          The maximum number of songs to return. Max 500, default 10.
// * offset:         The offset. Useful if you want to page through the songs in a genre.
// * musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *SubsonicClient) GetSongsByGenre(name string, parameters map[string]string) ([]*Song, error) {
	params := make(map[string]string)
	params["genre"] = name
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSongsByGenre", params)
	if err != nil {
		return nil, err
	}
	return resp.SongsByGenre.Songs, nil
}

func (s *SubsonicClient) GetNowPlaying() ([]*NowPlaying, error) {
	resp, err := s.Get("getNowPlaying", nil)
	if err != nil {
		return nil, err
	}
	return resp.NowPlaying.Entries, nil
}
