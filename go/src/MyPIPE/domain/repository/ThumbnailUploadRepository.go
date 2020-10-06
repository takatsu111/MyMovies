package repository

import (
	"MyPIPE/domain/model"
	"mime/multipart"
)

type ThumbnailUploadRepository interface {
	Upload(file multipart.File,movieFileHeader multipart.FileHeader,movieID model.MovieID) error
}