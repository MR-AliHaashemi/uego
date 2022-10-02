package vfs

import (
	"strings"

	"github.com/agrison/go-commons-lang/stringUtils"
)

type AbstractVfsReader struct {
	name  string
	files []pak.GameFile
}

func NewAbstractVfsReader(path string, versions any) *AbstractVfsReader {
	return &AbstractVfsReader{
		name: stringUtils.SubstringAfterLast(strings.ReplaceAll(path, "\\", "/"), "/"),
	}
}

func (a *AbstractVfsReader) FileCount() int {
	return len(a.files)
}
