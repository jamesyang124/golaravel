package celeritas

import "net/http"

func (c *Celeritas) SessionLoad(next http.Handler) http.Handler {
	c.InfoLog.Println("SessionLoad called")
	// load and save session to each handler
	return c.Session.LoadAndSave(next)
}
