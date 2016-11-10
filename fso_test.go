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
