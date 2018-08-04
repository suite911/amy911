package dirs

// +build darwin

/*
type Dirs struct {
	...
	sExecutableDirectory string
	sSystemCache         string
	sSystemConfig        string
	sSystemData          string
	sUserCache           string
	sUserConfig          string
	sUserData            string
	sUserDesktop         string
	sUserDocuments       string
	sUserDownloads       string
	sUserHome            string
	sUserPictures        string
	sUserScreenshots     string
}
*/

func init() {
	initDirs = func(d *Dirs, vendor, application string) {
	}
}
