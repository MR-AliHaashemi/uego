package uego

import (
	"github.com/MR-AliHaashemi/uego/ue4/objects/core/misc"
)

type File = string

type FileProvider struct {
	folder   File
	versions *VersionContainer

	localFiles   map[string]File
	files        map[string]GameFile
	unloadedPaks []AbstractAesVfsReader
	requiredKeys []misc.FGuid
	keys         map[misc.FGuid][]byte
	mountedPaks  []AbstractAesVfsReader

	mappingsProvider TypeMappingsProvider
}

func NewFileProvider(folder string, versions *VersionContainer, mappingsProvider *TypeMappingsProvider) *FileProvider {
	// scanFiles()
	return &FileProvider{
		folder:   folder,
		versions: versions,
	}
}

func (p *FileProvider) scanFiles(folder string) error {}

func (p *FileProvider) registerFile(file string) error {}

func (p *FileProvider) SaveGameFile(filePath string) ([]byte, error) {}

func (p *FileProvider) Close(filePath string) error {}
