
package filesystem

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"emperror.dev/errors"
	"github.com/gabriel-vasile/mimetype"
	"github.com/karrick/godirwalk"
	ignore "github.com/sabhiram/go-gitignore"

	"github.com/pterodactyl/wings/config"
	"github.com/pterodactyl/wings/system"
)

// Recursively iterates over a file or directory and sets the permissions on all of the
// underlying files. Iterate over all of the files and directories. If it is a file just
// go ahead and perform the chown operation. Otherwise dig deeper into the directory until
// we've run out of directories to dig into.
func (fs *Filesystem) Chown(path string) error {
	cleaned, err := fs.SafePath(path)
	if err != nil {
		return err
	}

	if fs.isTest {
		return nil
	}

	uid := config.Get().System.User.Uid
	gid := config.Get().System.User.Gid

	// Start by just chowning the initial path that we received.
	if err := os.Chown(cleaned, uid, gid); err != nil {
		return errors.Wrap(err, "server/filesystem: chown: failed to chown path")
	}

	// If this is not a directory we can now return from the function, there is nothing
	// left that we need to do.
	if st, err := os.Stat(cleaned); err != nil || !st.IsDir() {
		return nil
	}

	// If this was a directory, begin walking over its contents recursively and ensure that all
	// of the subfiles and directories get their permissions updated as well.
	err = godirwalk.Walk(cleaned, &godirwalk.Options{
		Unsorted: true,
		Callback: func(p string, e *godirwalk.Dirent) error {
			// Do not attempt to chown a symlink. Go's os.Chown function will affect the symlink
			// so if it points to a location outside the data directory the user would be able to
			// (un)intentionally modify that files permissions.
			if e.IsSymlink() {
				if e.IsDir() {
					return godirwalk.SkipThis
				}

				return nil
			}

			return os.Chown(p, uid, gid)
		},
	})

	return errors.Wrap(err, "server/filesystem: chown: failed to chown during walk function")
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