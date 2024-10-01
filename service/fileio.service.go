package service

import (
	"codecloud/model"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type FileIOService struct{}

func NewFileIOService() *FileIOService {
	return &FileIOService{}
}

func (fileIOService *FileIOService) ParseAndGetFolders(basePath, path string) (model.FileInfo, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return model.FileInfo{}, fmt.Errorf("failed to read directory: %w", err)
	}

	var entries []model.FileInfo
	for _, file := range files {
		if file.IsDir() {
			subFolder, err := fileIOService.ParseAndGetFolders(basePath, filepath.Join(path, file.Name()))
			if err != nil {
				return model.FileInfo{}, err
			}
			entries = append(entries, subFolder)
		} else {
			relPath, err := filepath.Rel(basePath, filepath.Join(path, file.Name()))
			if err != nil {
				return model.FileInfo{}, fmt.Errorf("failed to compute relative path: %w", err)
			}
			entries = append(entries, model.FileInfo{
				Path:  relPath,
				Name:  file.Name(),
				Files: nil,
			})
		}
	}

	relPath, err := filepath.Rel(basePath, path)
	if err != nil {
		return model.FileInfo{}, fmt.Errorf("failed to compute relative path: %w", err)
	}

	return model.FileInfo{
		Path:  relPath,
		Name:  filepath.Base(path),
		Files: entries,
	}, nil
}

func (fileIOService *FileIOService) GetFileData(basePath, relativePath string) ([]byte, error) {
	fsys := os.DirFS(basePath)
	log.Println("Reading file from basePath:", basePath, "relativePath:", relativePath)

	file, err := fs.ReadFile(fsys, relativePath)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return file, nil
}
