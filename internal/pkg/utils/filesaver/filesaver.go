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
func GetUniqueFileNames (quantity int, alreadyUsed map[string]struct{}) ([]string){
	fileNames := make([]string, 0)
	for i := 0; i < quantity; i++{
		randStringTemporary := ""
		for {
			randStringTemporary, err := randomizer.GenerateRandomString(6)
			if err != nil {
				return nil
			}
			fileNames = append(fileNames, randStringTemporary)

			_, alreadyUsed := alreadyUsed[randStringTemporary]
			if !alreadyUsed {
				break
			}
		}
		alreadyUsed[randStringTemporary] = struct{}{}
	}
	return fileNames
}

func UploadFile(reader io.Reader, root, path, name, ext string) (string, error) {
	log.Info("randString" + name)

	filename := name + ext
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
