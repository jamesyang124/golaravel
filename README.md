# golaravel

learning purpose only.

- section 34 - pick `pgx` as postgres connection client lib
- section 28 - update session with cookie session type
- section 25 - add scs session, hook with middleware
- section 24 - module test data as input struct
- section 23 - test for render engine/template
- section 19 - add jet template engine
- section 17 - fix os.Mkdir file mode value err
- section 15 - add render for go default template engine
- section 14 - add chi router, init server config and mount handler for it
- section 13 - add config from env
- section 12 - add logger and parse env
- section 11 - add dot env dep.
- section 08 - add folder generator for code base

## Section 28-34

- docker composed for local run time
- include dbs
- pick pgx for postgres client

## Section 25-27

- use `https://github.com/alexedwards/scs` for session mgmt
- hook cookie config from env
- prepare middleware function for handling scs session load and save for each handler iteration

## Section 23-24

- test guide - https://ieftimov.com/posts/testing-in-go-go-test/#:~:text=If%20you%20remember%20anything%20from,it's%20available%20on%20your%20machine.
- review coverage report to see how to tackle test coverage for source code
- add `testify/assertion` for testing
- module test data as input struct, this abstraction may force its flexibility is sticked in same test unit. 

## Section 19

- add jet template engine
- go around with iris web framework's template engine testing tool & report:

```
https://github.com/kataras/iris/tree/main/_benchmarks/view
```

## Section 17

- fix os.Mkdir wrong file mode value err
- move actual HTTP routes definition to golaravel's `routes.go` file instead
- `http.StripPrefix` for strip reuqest url's path prefix to next handler
- expose `handlers` to `app`, and with handler function

## Section 15

1. check `template.text.execute` how `reflect.ValueOf` for pointers
2. review method reciever as function argument for reasons.

## Section 14

1. run `go mod tidy` to add back missing module if `go get` not import new pacakage.
2. `F12` to see document.
3. review `/celeritas/routes.go` for `chi` router handler impl.

## Section 13

1. pass by value or pointer?

https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values

2. pointers pro & cons

https://medium.com/@briankworld/in-go-a-pointer-is-a-variable-that-stores-the-memory-address-of-another-variable-7584ac788041

## Section 12

1. add logger, `int` flag for properties
2. add `strconv` for parsing env for dotenv 

## Section 11

1. defer function handling, note defer arguments vs defer function call
2. get mod by `go get mod_name` ex: `go get github.com/joho/godotenv`
3. remove mod by `go get mod_name@none`
4. update go version management tool by `https://github.com/moovweb/gvm`

```sh
gvm install go1.20 -B
gvm install go1.22.4 -B
```

## Section 8 

#### celeritas & go mod

assume have another local repo, then its dep. in `go.mod`:

```sh
# go.mod
replace github.com/jamesyang124/celeritas => ../celeritas
```

then in `main.go`

```go
package main

import (
  "github.com/jamesyang124/celeritas"
)
```

#### go mod vendor

```sh
# first run go mod vendor would require to go get module
go mod vendor
#golaravel imports
#        github.com/jamesyang124/celeritas: module github.com/jamesyang124/celeritas provides package github.com/jamesyang124/celeritas and is replaced but not required; to add it:
#        go get github.com/jamesyang124/celeritas

go get github.com/jamesyang124/celeritas
go mod vendor
# will create vendor folder + modulex.txt and copy celeritas project, 
# note original ../celeritas will no longer be updated if we directly made change in vendor's code base
# but once we call go mod vendor again, it will sync back to ../celeritas
# so make sure always change code on local repo instead of vendor folder
```

#### Makefile

in order to in-sync go mod vendor every time, crate make file for it.

https://hackmd.io/@sysprog/SySTMXPvl