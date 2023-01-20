package service

import (
	"encoding/base64"
	"fmt"
	"github.com/Hymiside/wishlists-api/pkg/models"
	"github.com/google/uuid"
	"net/http"
	"os"
	"strings"

	"github.com/Hymiside/wishlists-api/pkg/repository"
)

type ProfileService struct {
	repo repository.PersonalCabinet
}

func NewProfile(repo repository.PersonalCabinet) *ProfileService {
	return &ProfileService{repo: repo}
}

func (p *ProfileService) GetProfile(userId string) (map[string]string, error) {
	profile, err := p.repo.GetProfile(userId)
	if err != nil {
		return nil, err
	}

	if profile["image_base64"] != "none" {
		var (
			bytes       []byte
			base64Image string
		)

		bytes, err = os.ReadFile(profile["image_base64"])
		if err != nil {
			return profile, nil
		}

		mimeType := http.DetectContentType(bytes)
		switch mimeType {
		case "image/jpeg":
			base64Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(bytes)
		case "image/png":
			base64Image = "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes)
		}
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
		if wish.ImageURL != "none" {
			var (
				bytes       []byte
				base64Image string
			)

			bytes, err = os.ReadFile(wish.ImageURL)
			if err == nil {
				mimeType := http.DetectContentType(bytes)
				switch mimeType {
				case "image/jpeg":
					base64Image = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(bytes)
				case "image/png":
					base64Image = "data:image/png;base64," + base64.StdEncoding.EncodeToString(bytes)
				}
				wish.ImageURL = base64Image
			} else {
				// TODO сделать лог
				wish.ImageURL = ""
			}
		}
		wishes[i] = wish
	}
	return wishes, nil
}

func (p *ProfileService) CreateWish(wish models.Wish) (string, error) {
	wish.Id = uuid.New().String()

	if wish.ImageURL == "" {
		// TODO
	}

	f1 := strings.Index(wish.ImageURL, ":")
	f2 := strings.Index(wish.ImageURL, ";")
	f3 := strings.Index(wish.ImageURL, ",")
	mimeType := wish.ImageURL[f1+1 : f2]

	switch mimeType {
	case "image/jpeg":
		if err := writeImage(wish.ImageURL[f3+1:], "jpg", wish.Id); err != nil {
			wish.ImageURL = ""
			// TODO сделать лог
		}
		wish.ImageURL = fmt.Sprintf("pkg/images/wishes/%s.jpg", wish.Id)
	case "image/png":
		if err := writeImage(wish.ImageURL[f3+1:], "png", wish.Id); err != nil {
			wish.ImageURL = ""
			// TODO сделать лог
		}
		wish.ImageURL = fmt.Sprintf("pkg/images/wishes/%s.png", wish.Id)
	}

	// TODO
	return "", nil
}

func writeImage(imgBase64, typeImg, wishId string) error {
	file, err := os.Create(fmt.Sprintf("pkg/images/wishes/%s.%s", wishId, typeImg))
	if err != nil {
		return ErrCreateImage
	}
	defer file.Close()

	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return ErrDecodeImage
	}
	if _, err = file.Write(dec); err != nil {
		return ErrWriteImage
	}
	file.Seek(0, 0)
	return nil
}
