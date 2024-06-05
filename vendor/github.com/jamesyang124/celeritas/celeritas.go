package celeritas

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
}

type config struct {
	port     string
	renderer string
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

	// create log
	c.InfoLog, c.ErrorLog = c.startLoggers()

	// don't care error if debug env not found, treat it as default mode instead
	c.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	c.Version = version
	c.RootPath = rootPath

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

func (c *Celeritas) startLoggers() (*log.Logger, *log.Logger) {
	// set to file for PROD in later development
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
