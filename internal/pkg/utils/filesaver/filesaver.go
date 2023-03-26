package filesaver

import (
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/randomizer"
	"path/filepath"

	"fmt"
	"io"
	"os"
)

func createFile(root, dir, name string) (*os.File, error) {
	_, err := os.ReadDir(root + dir)
	if err != nil {
		err = os.MkdirAll(root+dir, 0777)
		if err != nil {
			return nil, err
		}
	}
	file, err := os.Create(root + dir + name)
	return file, err
}

func UploadFile(reader io.Reader, root, path, ext string, alreadyUsed map[string]struct{}) (string, error) {
	randString := "" 
	for {
		randStringTemporary, err := randomizer.GenerateRandomString(6)
		if err != nil {
			return "", err
		}
		randString = randStringTemporary
		log.Info("randString" + randString)
		_, alreadyUsed := alreadyUsed[randString]
		if !alreadyUsed {
			break
		}
	}
	log.Info("randString" + randString)

	filename := randString + ext
	log.Info("Created file with name " + filename)
	file, err := createFile(root, path, filename)
	if err != nil {
		return "", fmt.Errorf("file creating error: %s", err)
	}
	defer file.Close()

	filename = path + filename
	_, err = io.Copy(file, reader)
	if err != nil {
		return "", fmt.Errorf("copy error: %s", err)
	}
	return filepath.Base(filename), nil
}
