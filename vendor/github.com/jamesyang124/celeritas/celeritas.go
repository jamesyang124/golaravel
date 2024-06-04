package celeritas

import (
	"fmt"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName string
	Debug   bool
	Version string
}

func (c *Celeritas) New(rootPath string) error {
	pathConfigs := initPaths{
		rootPath:   rootPath,
		folderName: []string{"handler", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	if err := c.Init(pathConfigs); err != nil {
		return err
	}

	if err := c.checkDotEnv(rootPath); err != nil {
		return err
	}

	// read dotenv by:
	// https://github.com/joho/godotenv
	if err := godotenv.Load(rootPath + "/.env"); err != nil {
		return err
	}

	return nil
}

func (c *Celeritas) Init(initPaths initPaths) error {
	rp := initPaths.rootPath

	for _, path := range initPaths.folderName {
		if err := c.CreateDirIfNotExisted(rp + "/" + path); err != nil {
			return err
		}
	}

	return nil
}

func (c *Celeritas) checkDotEnv(rootPath string) error {
	err := c.CreateFileIfNotExisted(fmt.Sprintf("%s/.env", rootPath))

	if err != nil {
		return err
	}

	return nil
}
