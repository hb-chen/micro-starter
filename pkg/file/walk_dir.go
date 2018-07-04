package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"path"
)

//sema is a counting semaphore for limiting concurrency in dirents
var sema = make(chan struct{}, 20)

//读取目录dir下的文件信息
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

//获取目录dir下的文件
func walkDir(dir, suffix string, wg *sync.WaitGroup, fileInfos chan<- FileInfo) {
	defer wg.Done()
	for _, entry := range dirents(dir) {
		subDir := filepath.Join(dir, entry.Name())
		if entry.IsDir() { //目录
			wg.Add(1)
			go walkDir(subDir, suffix, wg, fileInfos)
		} else {
			if suffix == "" || suffix == path.Ext(entry.Name()) {
				fi := FileInfo{
					FileInfo: entry,
					Path:     subDir,
				}
				fileInfos <- fi
			}
		}
	}
}

func WalkDirs(roots []string, suffix string, files *FileSlice, nbytes *int64) {
	fileInfos := make(chan FileInfo)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, suffix, &wg, fileInfos)
	}
	go func() {
		wg.Wait() //等待goroutine结束
		close(fileInfos)
	}()

loop:
	for {
		select {
		case fi, ok := <-fileInfos:
			if !ok {
				break loop
			}

			*files = append(*files, fi)
			*nbytes += fi.Size()
		}
	}

	return
}
