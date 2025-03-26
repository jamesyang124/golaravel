module golaravel

go 1.23.0

toolchain go1.23.6

// local copy, instead of pulling from github.com
// go get github.com/jamesyang124/celeritas to instruct builder get from local repo
replace github.com/jamesyang124/celeritas => ./celeritas

require (
	github.com/CloudyKit/jet/v6 v6.2.0
	github.com/go-chi/chi/v5 v5.0.12
	github.com/jamesyang124/celeritas v0.0.0-00010101000000-000000000000
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/alexedwards/scs/v2 v2.8.0 // indirect
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/go-sql-driver/mysql v1.9.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.14.3 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.3 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgtype v1.14.4 // indirect
	github.com/jackc/pgx/v4 v4.18.3 // indirect
	github.com/jackc/pgx/v5 v5.7.2 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/segmentio/fasthash v1.0.3 // indirect
	github.com/upper/db/v4 v4.10.0 // indirect
	golang.org/x/crypto v0.35.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)
