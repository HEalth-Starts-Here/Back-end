package filesaver

import (
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/randomizer"
	"mime/multipart"
	"path/filepath"

	"fmt"
	"io"
	"os"
)

func GetExtension(file *multipart.FileHeader) (string, bool) {
	k := len(file.Filename) - 1
	extension := ""
	for k != 0 {
		if k == 0 {
			return "", false
		}
		if (file.Filename)[k] == '.' {
			extension = (file.Filename)[k+1:]
		}
		k = k - 1
	}
	return extension, true
}

func ExtractNames(filePaths []string) (fileName []string) {
	imageNames := []string{}
	for i := range filePaths {
		imageNames = append(imageNames, ExtractName(filePaths[i]))
	}
	return imageNames
}

func ExtractName(filePath string) (fileName string) {
	i := len(filePath) - 1
	for i >= 0 {
		if filePath[i] == '/' || filePath[i] == '\\' {
			fileName = filePath[i+1:]
		}
		if i == 0 {
			fileName = filePath[i:]

		}
		i--
	}
	return fileName
}

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

func deleteFile(root, dir, name string) (error) {
	_, err := os.ReadDir(root + dir)
	if err != nil {
		err = os.MkdirAll(root+dir, 0777)
		if err != nil {
			return err
		}
	}
	err = os.Remove(root + dir + name)
	return err
}

func SaveMultipartDataFiles(fileNames []string, fileHeaders []*multipart.FileHeader) (error) {
	// TODO: add mapping from error to http code
	for i, _ := range fileNames {
		file, err := fileHeaders[i].Open()
		defer file.Close()
		if err != nil {
			return domain.Err.ErrObj.InternalServer
		}
		extension := filepath.Ext(fileNames[i])
		nameWithouExtension := fileNames[i][:len(fileNames[i])-len(extension)]
		_, err = UploadFile(file, "", config.DevConfigStore.LoadedFilesPath, nameWithouExtension, filepath.Ext(fileNames[i]))
		if err != nil {
			return domain.Err.ErrObj.InternalServer
		}
	}
	return nil
}

func DeleteFiles(root, path string, names []string) (error) {

	for i := range names {
		log.Info("Delete file with name " + names[i])
		err := deleteFile(root, path, names[i])
		if err != nil {
			return fmt.Errorf("file deleting error: %s", err)
		}

	}
	return nil
}

