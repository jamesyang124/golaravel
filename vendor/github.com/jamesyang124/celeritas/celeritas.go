package celeritas

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Celeritas struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Routes   *chi.Mux
	RootPath string
	config   config // config does not expose public fields, Celeritas should follow its accessibility
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

	if routes, ok := c.routes().(*chi.Mux); ok {
		c.Routes = routes
	} else {
		c.ErrorLog.Fatal("Celeritas.routes() cannot type assertion to *chi.Mux")
	}

	c.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	return nil
}

// inti to create application folders for each component
func (c *Celeritas) Init(initPaths initPaths) error {
	rp := initPaths.rootPath

	for _, path := range initPaths.folderName {
		if err := c.CreateDirIfNotExisted(rp + "/" + path); err != nil {
			return err
		}
	}

	return nil
}

// create server listening, and attach handler
func (c *Celeritas) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", c.config.port),
		ErrorLog:     c.ErrorLog,
		Handler:      c.Routes,
		IdleTimeout:  30 * time.Second, // wait time out for keep-alive headered requests
		ReadTimeout:  30 * time.Second, // reuqest in
		WriteTimeout: 30 * time.Second, // response out
	}

	c.InfoLog.Printf("Listening on port %s", c.config.port)
	err := srv.ListenAndServe()
	c.ErrorLog.Fatal(err)
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
