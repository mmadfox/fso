//Package fso provides create special directories
//with depth 1-7 level like /tmp/f1/15/5c/c4/4b/f15c4b179ccbbcfbb4267eb0f0b61d39.go
package fso

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Info struct {
	Id       string
	Path     string
	Filename string
	FullPath string
	Map      []string
	err      error
}

func (i Info) String() string {
	return fmt.Sprintf("%s fullpath: %s, path: %s, filename: %s",
		i.Id,
		i.FullPath,
		i.Path,
		i.Filename)
}

func (i Info) Error() error {
	return i.err
}

func makeHash(f string) string {
	filename := strings.Trim(f, "")
	hasher := md5.New()
	hasher.Write([]byte(filename))
	return hex.EncodeToString(hasher.Sum(nil))
}

func checkDepth(d int) int {
	if d <= 0 {
		return 1
	}
	if d >= 7 {
		return 7
	}
	return d
}

//MkDir creates a new directory with the specified dir, name and depth
func MkDir(dir string, filename string, depth int) Info {
	depth = checkDepth(depth)
	sum := makeHash(filename)
	ext := filepath.Ext(filename)
	nf := sum + ext
	p := make([]string, 0)
	p = append(p, dir)
	for i := 0; i < depth; i++ {
		p = append(p, sum[i:i+2])
	}
	full := filepath.Join(p...)
	if err := os.MkdirAll(full, 0755); err != nil {
		return Info{err: err}
	}
	return Info{
		Id:       sum,
		err:      nil,
		Map:      p,
		Filename: nf,
		Path:     full,
		FullPath: filepath.Join(full, nf),
	}
}
