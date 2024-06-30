package celeritas

type initPaths struct {
	rootPath   string
	folderName []string
}

type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}
