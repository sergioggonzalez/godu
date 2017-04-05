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

func WalkDirs (dirNames string) {

	dirs := strings.Split(dirNames, ",")
	fmt.Println("Cantida de Gorutinas:", len(dirs))
	var wg sync.WaitGroup
	wg.Add(len(dirs))

	size := make(chan int64)

	for _, dirName :=range dirs {
		go func(dirName string) {
			fmt.Println("Lanzo Gorutina para dir: ", dirName)
			walk(dirName, size)
			wg.Done()
		}(dirName)
	}



	var total int64
	go func(){
		fmt.Println("Lanzo Gorutina final ")
		for s := range size {
			total += s
		}
	}()

	// Wait for the goroutines to finish.
	fmt.Println("Waiting To Finish")
	wg.Wait()

	close(size)
	fmt.Println("Finaliza ejecucion de gorutinas!")
	fmt.Printf("%.2f GB\n", float32(total)/1e9)
}


func walk(dir string, size chan int64) {
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


