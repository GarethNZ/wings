
package filesystem

import (
	"os"
)

// Not applicable for windows
func (fs *Filesystem) Chown(path string) error {
	return nil
}

func (fs *Filesystem) Chmod(path string, mode os.FileMode) error {
	cleaned, err := fs.SafePath(path)
	if err != nil {
		return err
	}

	if fs.isTest {
		return nil
	}

	if err := os.Chmod(cleaned, mode); err != nil {
		return err
	}

	return nil
}
