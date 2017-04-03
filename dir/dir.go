package dir

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"fmt"
	"runtime"
	"sync"
)

func init() {
	// Allocate one logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)
}

func WalkDirs (dirNames string) int64{

	dirs := strings.Split(dirNames, ",")
	fmt.Println("Cantida de Gorutinas:", len(dirs))
	var wg sync.WaitGroup
	wg.Add(len(dirs))

	size := make(chan int64)

	for _, dirName :=range dirs {
		fmt.Println(dirName)
		go func() {
			fmt.Println("Lanzo Gorutina para dir: ", dirName)
			walk(dirName, size)
			wg.Done()
		}()
	}
	var total int64
	for s := range size {
		total += s
	}

	wg.Wait()
	close(size)
	return total
}


func walk(dir string, size chan int64) {
	fmt.Println(dir);
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walk(subdir, size)
		} else {
			size <- entry.Size()
		}
	}

}

func dirents(dir string) []os.FileInfo {
	dirFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil
	}
	return dirFiles

}
