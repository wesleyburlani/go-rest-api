package albums_service

import (
	"fmt"
	"test/web-service/database"
	"test/web-service/models"
)

type AlbumsService struct {
	database *database.Database
}

func NewAlbumsService(
	database *database.Database,
) *AlbumsService {
	return &AlbumsService{
		database,
	}
}

func (instance *AlbumsService) CreateAlbum(props models.AlbumProps) models.Album {
	return instance.database.CreateAlbum(props)
}

func (instance *AlbumsService) GetAlbums(page int, limit int) []models.Album {
	return instance.database.GetAlbums(page, limit)
}

func (instance *AlbumsService) GetAlbum(id string) (models.Album, error) {
	if album, exists := instance.database.GetAlbum(id); exists {
		return album, nil
	}
	return models.Album{}, fmt.Errorf("album with id %s not found", id)
}

func (instance *AlbumsService) UpdateAlbum(id string, props models.AlbumProps) (models.Album, error) {
	if album, updated := instance.database.UpdateAlbum(id, props); updated {
		return album, nil
	}
	return models.Album{}, fmt.Errorf("album with id %s not found", id)
}