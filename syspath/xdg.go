// +build dragonfly freebsd linux netbsd openbsd solaris

package syspath

import (
	"os"
	"path/filepath"
)

func initSysPath(sp *SysPath, vendor, application string) {
	home := os.Getenv("HOME")
	if len(home) < 1 {
		panic("The HOME environment variable was unset!")
	}
	xdg_cache_home := os.Getenv("XDG_CACHE_HOME")
	if len(xdg_cache_home) < 1 {
		xdg_cache_home = filepath.Join(home, ".cache")
	}
	xdg_config_home := os.Getenv("XDG_CONFIG_HOME")
	if len(xdg_config_home) < 1 {
		xdg_config_home = filepath.Join(home, ".config")
	}
	xdg_data_home := os.Getenv("XDG_DATA_HOME")
	if len(xdg_data_home) < 1 {
		xdg_data_home = filepath.Join(home, ".local/share")
	}

	sp.sCache = filepath.Join(xdg_cache_home, vendor, application)
	sp.sConfig = filepath.Join(xdg_config_home, vendor, application)
	sp.sData = filepath.Join(xdg_data_home, vendor, application)
	sp.sDesktop = filepath.Join(home, "Desktop")
	sp.sDocuments = filepath.Join(home, "Documents")
	sp.sDownloads = filepath.Join(home, "Downloads")
	sp.sHome = home
	sp.sPictures = filepath.Join(home, "Pictures")
	sp.sScreenshots = filepath.Join(home, "Pictures", "Screenshots")
}
