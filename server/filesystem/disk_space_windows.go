package filesystem

import (
	//"os/exec"
	//"encoding/binary"
	
	"fmt"
	"os"

	"github.com/apex/log"
)

// Determines the directory size of a given location by running parallel tasks to iterate
// through all of the folders. Returns the size in bytes. This can be a fairly taxing operation
// on locations with tons of files, so it is recommended that you cache the output.
func (fs *Filesystem) DirectorySize(path string) (int64, error) {
	d, err := fs.SafePath(path)
	// if err != nil {
	// 	return 0, err
	// }



	log.WithField("d", d).Info("checking DirectorySize for")

	var size int64

	dir, err := os.Open(d)
	if err != nil {
		return size, err
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			var s, _ = fs.DirectorySize(fmt.Sprintf("%s/%s", d, file.Name()))
			size += s
		} else {
			size += file.Size()
		}
	}

	return size, err

	// if err != nil {
	// 	return 0, err
	// }
	
	// size := binary.BigEndian.Uint64(co)
	// return int64(size), err
}