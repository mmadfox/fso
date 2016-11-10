# fso
Provides create special directories for storage large number of files 

##Install
$ go get github.com/mmadfox/fso

###Example

```Go
package fso

import (
	"math/rand"
	"os"
	"path"
	"testing"
	"time"
)

func getFilename() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	strlen := 10
	result := make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result) + ".go"
}

func isDir(path string) bool {
	i, err := os.Stat(path)
	if err != nil {
		return false
	}
	return i.IsDir()
}

func TestMkDir(t *testing.T) {
	depth := 5
	root := "/tmp"
	for i := 0; i <= 100; i++ {
		filename := getFilename()
		info := MkDir(root, filename, depth)
		if info.Error() != nil {
			t.Fatal(info.Error())
		}
		if ok := isDir(info.Path); !ok {
			t.Errorf("MkDir(%s, %s, %d) = %v; want Info{}", root, filename, depth, info)
		}
		tmp := path.Join(info.Map[0:2]...)
		if err := os.RemoveAll(tmp); err != nil {
			t.Fatal(err)
		}
		t.Log(info)
	}
}
```

### Output
```
$ fso_test.go:46: fd41c6e82b737d8ab8d27a618dfda71f fullpath: /tmp/fd/d4/41/1c/c6/fd41c6e82b737d8ab8d27a618dfda71f.go, path: /tmp/fd/d4/41/1c/c6, filename: fd41c6e82b737d8ab8d27a618dfda71f.go
$ fso_test.go:46: 6aba4d6874b22ce673a621b16bf2361c fullpath: /tmp/6a/ab/ba/a4/4d/6aba4d6874b22ce673a621b16bf2361c.go, path: /tmp/6a/ab/ba/a4/4d, filename: 6aba4d6874b22ce673a621b16bf2361c.go
$ fso_test.go:46: 6e7b3d867a88e3294d5026d63811745b fullpath: /tmp/6e/e7/7b/b3/3d/6e7b3d867a88e3294d5026d63811745b.go, path: /tmp/6e/e7/7b/b3/3d, filename: 6e7b3d867a88e3294d5026d63811745b.go
$ fso_test.go:46: 92787c44f62e65c109c2663cb2ede626 fullpath: /tmp/92/27/78/87/7c/92787c44f62e65c109c2663cb2ede626.go, path: /tmp/92/27/78/87/7c, filename: 92787c44f62e65c109c2663cb2ede626.go
$ fso_test.go:46: e7a5b864831612b6a5df1eec17130bfc fullpath: /tmp/e7/7a/a5/5b/b8/e7a5b864831612b6a5df1eec17130bfc.go, path: /tmp/e7/7a/a5/5b/b8, filename: e7a5b864831612b6a5df1eec17130bfc.go
$ fso_test.go:46: 14412254670c38978a666cc0c947d66c fullpath: /tmp/14/44/41/12/22/14412254670c38978a666cc0c947d66c.go, path: /tmp/14/44/41/12/22, filename: 14412254670c38978a666cc0c947d66c.go
$ fso_test.go:46: 5420b979c8a6812c03845209c4aea95c fullpath: /tmp/54/42/20/0b/b9/5420b979c8a6812c03845209c4aea95c.go, path: /tmp/54/42/20/0b/b9, filename: 5420b979c8a6812c03845209c4aea95c.go
$ fso_test.go:46: 7cb93fd3238f5b1fda3f8cddcbe30dc0 fullpath: /tmp/7c/cb/b9/93/3f/7cb93fd3238f5b1fda3f8cddcbe30dc0.go, path: /tmp/7c/cb/b9/93/3f, filename: 7cb93fd3238f5b1fda3f8cddcbe30dc0.go
$ fso_test.go:46: 32ef61f6aece1ce7b2f6c57b9c9cc175 fullpath: /tmp/32/2e/ef/f6/61/32ef61f6aece1ce7b2f6c57b9c9cc175.go, path: /tmp/32/2e/ef/f6/61, filename: 32ef61f6aece1ce7b2f6c57b9c9cc175.go

```
