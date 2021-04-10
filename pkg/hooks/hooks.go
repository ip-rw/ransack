package hooks

import (
	"github.com/ip-rw/ransack/pkg/models"
	"os"
	"strings"
)

type FileInfo struct {
	os.FileInfo
	Path string
}

func (p *FileInfo) Lpath() string {
	return strings.ToLower(p.Path)
}

type Plugin interface {
	Search(models.MachineServiceClient, FileInfo, []byte) (bool, error)
	Name() string
}
