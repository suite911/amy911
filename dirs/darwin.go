// +build darwin

package dirs

/*
type Dirs struct {
	...
	sCache           string
	sConfig          string
	sData            string
	sDesktop         string
	sDocuments       string
	sDownloads       string
	sHome            string
	sPictures        string
	sScreenshots     string
}
*/

func init() {
	initDirs = func(d *Dirs, vendor, application string) {
	}
}
