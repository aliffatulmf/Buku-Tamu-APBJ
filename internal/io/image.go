package io

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrWriteImage        = errors.New("can't write image into storage")
	ErrImageDecode       = errors.New("an error occurred during the decoding process")
	ErrUnsupportedFormat = errors.New("unsupported image format")
	ErrInvalidFormat     = errors.New("invalid image data format")
)

type ImageStorage struct {
	dir string
}

func NewImageStorage(dir string) *ImageStorage {
	return &ImageStorage{dir: dir}
}

func (s *ImageStorage) Save(base64Data string) (string, error) {
	parts := strings.Split(base64Data, ",")
	if len(parts) != 2 {
		return "", ErrInvalidFormat
	}

	var ext string
	switch parts[0] {
	case "data:image/png;base64":
		ext = "png"
	case "data:image/jpeg;base64":
		ext = "jpeg"
	default:
		return "", ErrUnsupportedFormat
	}

	dec, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", ErrImageDecode
	}

	if err := os.MkdirAll(s.dir, os.ModeDir); err != nil {
		return "", err
	}

	name := uuid.NewString() + "." + ext
	filePath := filepath.Join(s.dir, name)
	mediaPath := strings.ReplaceAll(filePath, "\\", "/")

	if err := os.WriteFile(filePath, dec, 0644); err != nil {
		return "", ErrWriteImage
	}

	return mediaPath, nil
}
