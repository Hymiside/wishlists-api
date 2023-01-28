package service

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/Hymiside/wishlists-api/pkg/repository"
	"github.com/google/uuid"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfile(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (p *ProfileService) GetProfile(userId string) (map[string]string, error) {
	profile, err := p.repo.GetProfile(userId)
	if err != nil {
		return nil, err
	}

	if profile["image_base64"] != "" {
		base64Image := getBase64Image(profile["image_base64"])
		profile["image_base64"] = base64Image
	}
	return profile, nil
}

func (p *ProfileService) GetWishes(userId string) ([]models.Wish, error) {
	wishes, err := p.repo.GetWishes(userId)
	if err != nil {
		return nil, err
	}

	if wishes == nil {
		// TODO сделать лог
		return wishes, nil
	}

	for i, wish := range wishes {
		if wish.ImageURL != "" {
			base64Image := getBase64Image(wish.ImageURL)
			wish.ImageURL = base64Image
		}
		wishes[i] = wish
	}
	return wishes, nil
}

func (p *ProfileService) CreateWish(wish models.Wish) (string, error) {
	wish.Id = uuid.New().String()

	if wish.ImageURL == "" {
		wishId, err := p.repo.CreateWish(wish)
		if err != nil {
			// TODO сделаг лог
			return "", err
		}
		return wishId, nil
	}

	f1 := strings.Index(wish.ImageURL, ":")
	f2 := strings.Index(wish.ImageURL, ";")
	f3 := strings.Index(wish.ImageURL, ",")
	mimeType := wish.ImageURL[f1+1 : f2]

	typeImg, err := writeImage(wish.ImageURL[f3+1:], wish.Id, mimeType)
	if err != nil {
		wish.ImageURL = ""
		// TODO сделать лог
	} else {
		wish.ImageURL = fmt.Sprintf("pkg/images/wishes/%s.%s", wish.Id, typeImg)
	}

	var wishId string
	wishId, err = p.repo.CreateWish(wish)
	if err != nil {
		// TODO сделаг лог
		return "", err
	}
	return wishId, nil
}

func (p *ProfileService) GetFavorites(userId string) ([]map[string]string, error) {
	//TODO implement me
	panic("implement me")
}

func writeImage(imgBase64, wishId, mimeType string) (string, error) {
	var typeImg string

	switch mimeType {
	case "image/jpeg":
		typeImg = "jpg"
	case "image/png":
		typeImg = "png"
	}

	file, err := os.Create(fmt.Sprintf("pkg/images/wishes/%s.%s", wishId, typeImg))
	if err != nil {
		return "", ErrCreateImage
	}
	defer file.Close()

	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", ErrDecodeImage
	}
	if _, err = file.Write(dec); err != nil {
		return "", ErrWriteImage
	}
	return typeImg, nil
}

func getBase64Image(imageURL string) string {
	var base64Image string

	bytes, err := os.ReadFile(imageURL)
	if err != nil {
		// TODO сделать лог
		return ""
	}

	mimeType := http.DetectContentType(bytes)
	switch mimeType {
	case "image/jpeg":
		base64Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(bytes)
	case "image/png":
		base64Image = "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes)
	}
	return base64Image
}
