package handler

import (
	"net/url"
	"os"
	"path/filepath"
	"void-project/global"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/service"

	"github.com/gin-gonic/gin"
)

type Storage struct {
	service *service.StorageService
}

func NewStorage() *Storage {
	return &Storage{service.NewStorageService()}
}

func (s *Storage) List(c *gin.Context) {
	dirs, err := s.service.List(c.Query("path"))
	if err != nil {
		if os.IsNotExist(err) {
			response.FailError(c, apierr.DirNotExist)
			return
		}
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.Success(c, dirs)
}

func (s *Storage) Mkdir(c *gin.Context) {
	var param struct {
		Path string
		Name string
	}
	if err := c.ShouldBind(&param); err != nil {
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}

	path := global.Config.System.StorageLocation + param.Path + "/" + param.Name
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		response.FailError(c, apierr.CreateFailed, err)
		return
	}
	response.SuccessOk(c)
}

func (s *Storage) Upload(c *gin.Context) {
	err := s.service.Upload(c)
	if err != nil {
		response.FailError(c, apierr.FileUploadFailed, err)
		return
	}
	response.SuccessOk(c)
}

func (s *Storage) Download(c *gin.Context) {
	path := global.Config.System.StorageLocation + c.Query("path")
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			response.FailError(c, apierr.FileNotExist)
			return
		} else {
			response.FailError(c, apierr.FetchFailed, err)
			return
		}
	}
	_, filename := filepath.Split(path)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment;filename*=UTF-8''"+url.QueryEscape(filename))
	c.Header("fileName", url.QueryEscape(filename))
	c.File(path)
}

func (s *Storage) Rename(c *gin.Context) {
	var param struct {
		Oldpath string
		Newpath string
	}
	if err := c.ShouldBind(&param); err != nil {
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}
	if param.Oldpath == "" || param.Newpath == "" {
		response.FailError(c, apierr.InvalidPath)
		return
	}

	param.Oldpath = global.Config.System.StorageLocation + param.Oldpath
	if _, err := os.Stat(param.Oldpath); err != nil {
		if os.IsNotExist(err) {
			response.FailError(c, apierr.FileNotExist)
			return
		} else {
			response.FailError(c, apierr.UpdateFailed, err)
			return
		}
	}
	param.Newpath = global.Config.System.StorageLocation + param.Newpath
	if err := os.Rename(param.Oldpath, param.Newpath); err != nil {
		response.FailError(c, apierr.UpdateFailed, err)
		return
	}
	response.SuccessOk(c)
}

func (s *Storage) Delete(c *gin.Context) {
	path := c.Query("path")
	if path == "" {
		response.FailError(c, apierr.InvalidPath)
		return
	}
	path = global.Config.System.StorageLocation + path
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			response.FailError(c, apierr.FileNotExist)
			return
		} else {
			response.FailError(c, apierr.DeleteFailed, err)
			return
		}
	}
	if err := os.RemoveAll(path); err != nil {
		response.FailError(c, apierr.DeleteFailed, err)
	}
	response.SuccessOk(c)
}

func (s *Storage) Copy(c *gin.Context) {}

func (s *Storage) Move(c *gin.Context) {}
