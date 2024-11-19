package celeritas

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/jamesyang124/celeritas/render"
	"github.com/jamesyang124/celeritas/session"
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
	Render   *render.Render
	JetViews *jet.Set
	Session  *scs.SessionManager
	DB       Database
	config   config // config does not expose public fields, Celeritas should follow its accessibility
}

// for internal modulized purpose
type config struct {
	port        string
	renderer    string
	cookie      cookieConfig
	database    databaseConfig
	sessionType string
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

	// connect to db
	if os.Getenv("DATABASE_TYPE") != "" {
		db, err := c.OpenDB(os.Getenv("DATABASE_TYPE"), c.BuildDSN())
		if err != nil {
			c.ErrorLog.Println(err)
			os.Exit(1)
		}
		c.DB = Database{
			DataType: os.Getenv("DATABASE_TYPE"),
			Pool:     db,
		}
	}

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
		cookie: cookieConfig{
			name:     os.Getenv("COOKIE_NAME"),
			lifetime: os.Getenv("COOKIE_LIFETIME"),
			persist:  os.Getenv("COOKIE_PRESIST"),
			secure:   os.Getenv("COOKIE_SECURE"),
			domain:   os.Getenv("COOKIE_DOMAIN"),
		},
		sessionType: os.Getenv("SESSION_TYPE"),
	}

	ses := session.Session{
		CookieLifetime: c.config.cookie.lifetime,
		CookiePersist:  c.config.cookie.persist,
		CookieName:     c.config.cookie.name,
		SessionType:    c.config.sessionType,
		CookieDomain:   c.config.cookie.domain,
	}

	c.Session = ses.InitSession()

	jetViews := jet.NewSet(
		jet.NewOSFileSystemLoader(fmt.Sprintf("%s/views", rootPath)),
		jet.InDevelopmentMode(),
	)

	c.JetViews = jetViews
	c.createRender()

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

// put reciever type as function argument may benefit its test stubbing, or some readability reason
// updated: remove above impl. finally, unless for above reason..
func (c *Celeritas) createRender() {
	myRender := render.Render{
		Renderer: c.config.renderer,
		RootPath: c.RootPath,
		Port:     c.config.port,
		JetViews: c.JetViews,
	}

	c.Render = &myRender
}

func (c *Celeritas) BuildDSN() string {
	var dsn string

	switch os.Getenv("DATABASE_TYPE") {
	case "postgres", "postgresql":
		dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s timezone=UTC connect_timeout=5",
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_NAME"),
			os.Getenv("DATABASE_SSL_MODE"),
		)

		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, os.Getenv("DATABASE_PASS"))
		}
	default:

	}

	return dsn
}
