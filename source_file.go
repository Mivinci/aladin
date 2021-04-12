package aladin

import (
	"io"
	"os"

	"github.com/fsnotify/fsnotify"
)

var (
	_ Source  = (*file)(nil)
	_ Watcher = (*fileWatcher)(nil)
)

const DefaultPath = "config.yml"

type file struct {
	path string
}

func NewFileSource(path string) Source {
	if len(path) == 0 {
		path = DefaultPath
	}
	return &file{path: path}
}

func (f *file) Read() (*Snapshot, error) {
	fp, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	info, err := fp.Stat()
	if err != nil {
		return nil, err
	}

	snap := Snapshot{
		Data:      data,
		Path:      f.path,
		Source:    f.String(),
		Timestamp: info.ModTime(),
	}
	snap.Checksum = snap.Sum()

	return &snap, nil
}

func (f *file) Write(*Snapshot) error {
	return nil
}

func (f *file) Watch() (Watcher, error) {
	return newFileWatcher(f)
}

func (f *file) Close() error {
	return nil
}

func (f *file) String() string {
	return "file"
}

type fileWatcher struct {
	f *file
	w *fsnotify.Watcher
}

func newFileWatcher(f *file) (Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	w.Add(f.path)
	return &fileWatcher{f: f, w: w}, nil
}

func (w *fileWatcher) Next() (*Snapshot, error) {
	select {
	case event := <-w.w.Events:
		if event.Op == fsnotify.Rename {
			_, err := os.Stat(event.Name)
			if err == nil || os.IsExist(err) {
				w.w.Add(event.Name)
			}
		}
		return w.f.Read()
	case err := <-w.w.Errors:
		return nil, err
	}
}

func (w *fileWatcher) Stop() error {
	return w.w.Close()
}
