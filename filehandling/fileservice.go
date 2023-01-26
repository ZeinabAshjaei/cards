package filehandling

import (
	"os"
)

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

func (fs *FileService) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
func (fs *FileService) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}
