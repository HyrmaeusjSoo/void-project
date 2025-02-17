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

// 存储
type Storage struct {
	service *service.StorageService
}

func NewStorage() *Storage {
	return &Storage{service.NewStorageService()}
}

// 获取文件列表
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

// 创建目录
func (s *Storage) Mkdir(c *gin.Context) {
	var param struct {
		Path string
		Name string
	}
	if err := c.ShouldBind(&param); err != nil {
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}

	path := global.Configs.System.StorageLocation + param.Path + "/" + param.Name
	if err := os.Mkdir(path, os.ModePerm); err != nil {
		response.FailError(c, apierr.CreateFailed, err)
		return
	}
	response.SuccessOk(c)
}

// 上传文件
func (s *Storage) Upload(c *gin.Context) {
	err := s.service.Upload(c)
	if err != nil {
		response.FailError(c, apierr.FileUploadFailed, err)
		return
	}
	response.SuccessOk(c)
}

// 下载文件
func (s *Storage) Download(c *gin.Context) {
	path := global.Configs.System.StorageLocation + c.Query("path")
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

// 重命名文件
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

	param.Oldpath = global.Configs.System.StorageLocation + param.Oldpath
	if _, err := os.Stat(param.Oldpath); err != nil {
		if os.IsNotExist(err) {
			response.FailError(c, apierr.FileNotExist)
			return
		} else {
			response.FailError(c, apierr.UpdateFailed, err)
			return
		}
	}
	param.Newpath = global.Configs.System.StorageLocation + param.Newpath
	if err := os.Rename(param.Oldpath, param.Newpath); err != nil {
		response.FailError(c, apierr.UpdateFailed, err)
		return
	}
	response.SuccessOk(c)
}

// 删除文件
func (s *Storage) Delete(c *gin.Context) {
	var paths struct {
		Path []string
	}
	if err := c.ShouldBind(&paths); err != nil {
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}

	for _, v := range paths.Path {
		path := global.Configs.System.StorageLocation + v
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
	}
	response.SuccessOk(c)
}

// 复制文件
func (s *Storage) Copy(c *gin.Context) {
	var param struct {
		Origin string
		Target string
	}
	if err := c.ShouldBind(&param); err != nil {
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}

	//复制

	response.SuccessOk(c)
}

// 移动文件
func (s *Storage) Move(c *gin.Context) {
	var param struct {
		Oldpath string
		Newpath string
	}
	if err := c.ShouldBind(&param); err != nil {
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}

	// 移动

	response.SuccessOk(c)
}
