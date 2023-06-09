package albums

import (
	"context"
	"fmt"

	"github.com/wesleyburlani/go-rest-api/models"

	linq "github.com/ahmetb/go-linq/v3"
)

type MockAlbumsService struct {
	Albums []models.Album
	ctx    context.Context
}

func NewMockAlbumsService() *MockAlbumsService {
	return &MockAlbumsService{
		Albums: []models.Album{
			{
				ID:     1,
				Title:  "Blue Train",
				Artist: "John Coltrane",
				Price:  56.99,
			},
			{
				ID:     2,
				Title:  "Jeru",
				Artist: "Gerry Mulligan",
				Price:  17.99,
			},
			{
				ID:     3,
				Title:  "Sarah Vaughan and Clifford Brown",
				Artist: "Sarah Vaughan",
				Price:  39.99,
			},
		},
	}
}

func (instance *MockAlbumsService) WithContext(ctx context.Context) IAlbumsService {
	instance.ctx = ctx
	return instance
}

func (instance *MockAlbumsService) CreateAlbum(props models.AlbumProps) models.Album {
	album := models.Album{
		ID:     uint(len(instance.Albums) + 2),
		Title:  props.Title,
		Artist: props.Artist,
		Price:  props.Price,
	}

	instance.Albums = append(instance.Albums, album)
	return album
}

func (instance *MockAlbumsService) GetAlbums(page int, limit int) []models.Album {
	skip := (page) * (limit)
	var results []models.Album
	linq.From(instance.Albums).Skip(skip).Take(limit).ToSlice(&results)
	return results
}

func (instance *MockAlbumsService) GetAlbum(id uint) (models.Album, error) {
	album := linq.From(instance.Albums).FirstWithT(func(a models.Album) bool {
		return a.ID == id
	})

	if album == nil {
		return models.Album{}, fmt.Errorf("album with id %d not found", id)
	}

	return album.(models.Album), nil
}

func (instance *MockAlbumsService) UpdateAlbum(id uint, props models.AlbumProps) (models.Album, error) {
	for index, album := range instance.Albums {
		if album.ID == id {
			instance.Albums[index] = models.Album{
				ID:     id,
				Title:  props.Title,
				Artist: props.Artist,
				Price:  props.Price,
			}
			return instance.Albums[index], nil
		}
	}

	return models.Album{}, fmt.Errorf("album with id %d not found", id)
}
