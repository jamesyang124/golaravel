module golaravel

go 1.22

// local copy, instead of pulling from github.com
// go get github.com/jamesyang124/celeritas to instruct builder get from local repo
replace github.com/jamesyang124/celeritas => ../celeritas

require (
	github.com/go-chi/chi/v5 v5.0.12
	github.com/jamesyang124/celeritas v0.0.0-00010101000000-000000000000
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.2.0 // indirect
	github.com/alexedwards/scs/v2 v2.8.0 // indirect
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
