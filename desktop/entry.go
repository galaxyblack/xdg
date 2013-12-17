package desktop

import (
	"errors"
	"github.com/rkoesters/xdg/ini"
	"io"
)

var (
	ErrMissingType = errors.New("missing entry type")
	ErrMissingName = errors.New("missing entry name")
	ErrMissingURL  = errors.New("missing entry url")
)

const dent = "Desktop Entry"

type Key string

const (
	Version         Key = "Version"
	Name                = "Name"
	GenericName         = "GenericName"
	NoDisplay           = "NoDisplay"
	Comment             = "Comment"
	Icon                = "Icon"
	Hidden              = "Hidden"
	OnlyShowin          = "OnlyShowIn"
	NotShowIn           = "NotShowIn"
	DBusActivatable     = "DBusActivatable"
	TryExec             = "TryExec"
	Exec                = "Exec"
	Path                = "Path"
	Terminal            = "Terminal"
	MimeType            = "MimeType"
	Categories          = "Categories"
	Keywords            = "Keywords"
	StartupNotify       = "StartupNotify"
	StartupWMClass      = "StartupWMClass"
	URL                 = "URL"
)

// Entry represents a desktop entry file.
type Entry struct {
	m ini.Map
}

func New(r io.Reader) (*Entry, error) {
	dfile, err := ini.New(r)
	if err != nil {
		return nil, err
	}

	e := &Entry{dfile}

	// Check that the desktop file is valid.
	_, ok := e.m[dent]["Type"]
	if !ok {
		return nil, ErrMissingType
	}
	switch e.Type() {
	case Link:
		_, ok = e.m[dent]["URL"]
		if !ok {
			return nil, ErrMissingURL
		}
		fallthrough
	case Application, Directory:
		_, ok = e.m[dent]["Name"]
		if !ok {
			return nil, ErrMissingName
		}
	}
	return e, nil
}

func (e *Entry) Type() Type {
	return ParseType(e.m.Get(dent, "Type"))
}

func (e *Entry) String(k Key) string {
	return e.m.Get(dent, string(k))
}

func (e *Entry) Bool(k Key) bool {
	return e.m.Bool(dent, string(k))
}

func (e *Entry) List(k Key) []string {
	return e.m.List(dent, string(k))
}
