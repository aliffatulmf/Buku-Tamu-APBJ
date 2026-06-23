package lib

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func SaveImageFile(enc string) (string, error) {
	sp := strings.Split(enc, ",")
	if len(sp) != 2 {
		return "", ErrInvalidImageFormat
	}

	var ext string
	switch sp[0] {
	case "data:image/png;base64":
		ext = "png"
	case "data:image/jpeg;base64":
		ext = "jpeg"
	default:
		return "", ErrUnsupportedFormat
	}

	dec, err := base64.StdEncoding.DecodeString(sp[1])
	if err != nil {
		return "", ErrImageDecode
	}

	name := uuid.NewString()
	loc := filepath.Join("media/img", name+"."+ext)
	med := strings.ReplaceAll(loc, "\\", "/")

	f, err := os.Create(loc)
	if err != nil {
		return "", ErrIOWriteImage
	}
	defer f.Close()

	if _, err = f.Write(dec); err != nil {
		return "", ErrIOWriteImage
	}

	return med, nil
}
