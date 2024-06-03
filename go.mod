module golaravel

go 1.22

// local copy, instead of pulling from github.com
// go get github.com/jamesyang124/celeritas to instruct builder get from local repo
// th
replace github.com/jamesyang124/celeritas => ../celeritas

require github.com/jamesyang124/celeritas v0.0.0-00010101000000-000000000000 // indirect
