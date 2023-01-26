package filehandling

import (
	"os"
)

type FileManager interface {
	ReadFile(name string) ([]byte, error)
	WriteFile(name string, data []byte, perm os.FileMode) error
}
