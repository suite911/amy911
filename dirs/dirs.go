package dirs

import (
	"os"
	"path/filepath"

	"github.com/amyadzuki/amygolib/onfail"
)

type Dirs struct {
	vendor, application, exedir string

	sCache       string
	sConfig      string
	sData        string
	sDesktop     string
	sDocuments   string
	sDownloads   string
	sHome        string
	sPictures    string
	sScreenshots string
}

func New(vendor, application string, onFail ...onfail.OnFail) *Dirs {
	return new(Dirs).Init(vendor, application, onFail...)
}

func (d *Dirs) Application() string {
	return d.application
}

func (d *Dirs) Cache() string {
	return d.sCache
}

func (d *Dirs) Config() string {
	return d.sConfig
}

func (d *Dirs) Data() string {
	return d.sData
}

func (d *Dirs) Desktop() string {
	return d.sDesktop
}

func (d *Dirs) Documents() string {
	return d.sDocuments
}

func (d *Dirs) Downloads() string {
	return d.sDownloads
}

func (d *Dirs) ExeDir() string {
	return d.exedir
}

func (d *Dirs) Home() string {
	return d.sHome
}

func (d *Dirs) Init(vendor, application string, onFail ...onfail.OnFail) *Dirs {
	d.vendor, d.application = vendor, application
	exedir, err := os.Executable()
	if err == nil {
		exedir, err = filepath.EvalSymlinks(exedir)
	}
	if err == nil {
		d.exedir = exedir
	} else {
		onfail.Fail(err, nil, onfail.Panic, onFail...)
	}
	initDirs(d, vendor, application)
	return d
}

func (d *Dirs) Pictures() string {
	return d.sPictures
}

func (d *Dirs) Screenshots() string {
	return d.sScreenshots
}

func (d *Dirs) Vendor() string {
	return d.vendor
}
