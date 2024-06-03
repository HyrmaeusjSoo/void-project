package service

import (
	"os"
	"void-project/global"
	"void-project/internal/model"

	"github.com/gin-gonic/gin"
)

type StorageService struct{}

func NewStorageService() *StorageService {
	return &StorageService{}
}

func (s *StorageService) List(path string) ([]model.DirEntry, error) {
	dirs := []model.DirEntry{}
	entries, err := os.ReadDir(global.Config.System.StorageLocation + path)
	if err != nil {
		return dirs, err
	}
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			return dirs, err
		}
		dirs = append(dirs, model.DirEntry{info.Name(), info.IsDir(), info.Size(), info.ModTime()})
	}
	return dirs, err
}

func (s *StorageService) Upload(c *gin.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	path := c.PostForm("path") + "/" + file.Filename
	return c.SaveUploadedFile(file, global.Config.System.StorageLocation+path)
}
