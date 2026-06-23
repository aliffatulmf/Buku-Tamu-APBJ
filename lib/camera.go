package lib

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"

	liberr "github.com/aliffatulmf/buku-tamu-apbj/lib/errors"

	"github.com/google/uuid"
)

func SaveImageFile(enc string) (string, error) {
	var name string
	var ext string

	sp := strings.Split(enc, ",")

	name = uuid.NewString()
	switch sp[0] {
	case "data:image/png;base64":
		ext = "png"
	case "data:image/jpeg;base64":
		ext = "jpeg"
	}

	dec, err := base64.StdEncoding.DecodeString(sp[1])
	if err != nil {
		return "", liberr.ErrImageDecode
	}

	loc := filepath.Join("media/img/", strings.Join([]string{name, ext}, "."))
	med := strings.Replace(loc, "\\", "/", 2)
	f, err := os.Create(loc)
	if err != nil {
		return "", liberr.ErrIOWriteImage
	}

	_, err = f.Write(dec)
	if err != nil {
		return "", liberr.ErrIOWriteImage
	}

	return med, nil
}
