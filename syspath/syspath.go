package syspath

import (
	"os"
	"path/filepath"

	"github.com/amy911/amy911/onfail"
)

type SysPath struct {
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

func New(vendor, application string, onFail ...onfail.OnFail) *SysPath {
	return new(SysPath).Init(vendor, application, onFail...)
}

func (sp *SysPath) Application() string {
	return sp.application
}

func (sp *SysPath) Cache() string {
	return sp.sCache
}

func (sp *SysPath) Config() string {
	return sp.sConfig
}

func (sp *SysPath) Data() string {
	return sp.sData
}

func (sp *SysPath) Desktop() string {
	return sp.sDesktop
}

func (sp *SysPath) Documents() string {
	return sp.sDocuments
}

func (sp *SysPath) Downloads() string {
	return sp.sDownloads
}

func (sp *SysPath) ExeDir() string {
	return sp.exedir
}

func (sp *SysPath) Home() string {
	return sp.sHome
}

func (sp *SysPath) Init(vendor, application string, onFail ...onfail.OnFail) *Dirs {
	sp.vendor, sp.application = vendor, application
	exefile, err := os.Executable()
	if err == nil {
		exefile, err = filepath.EvalSymlinks(exefile)
	}
	if err == nil {
		sp.exedir = filepath.Dir(exefile)
	} else {
		onfail.Fail(err, exefile, onfail.Panic, onFail)
	}
	initDirs(sp, vendor, application)
	return sp
}

func (sp *SysPath) Pictures() string {
	return sp.sPictures
}

func (sp *SysPath) Screenshots() string {
	return sp.sScreenshots
}

func (sp *SysPath) Vendor() string {
	return sp.vendor
}
