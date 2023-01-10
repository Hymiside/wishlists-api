package service

import (
	"encoding/base64"
	"net/http"
	"os"

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
