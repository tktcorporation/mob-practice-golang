package repository

import (
	"errors"

	"github.com/tktcorporation/mob-practice-golang/web-service-gin/album"
)

var albums = []album.Album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type error interface {
    Error() string
}

func GetAll() []album.Album {
    return albums
}

func Create(newAlbum album.Album) {
	albums = append(albums, newAlbum)
}

func GetById(id string) (*album.Album, error) {
	for _, a := range albums {
        if a.ID == id {
            return &a, nil
        }
    }
	return nil, errors.New("not found")
}