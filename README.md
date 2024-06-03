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

