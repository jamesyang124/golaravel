package celeritas

import "database/sql"

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

type databaseConfig struct {
	// data source name for db connection
	dsn      string
	database string
}

type Database struct {
	DataType string
	Pool     *sql.DB
}
