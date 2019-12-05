package main

import (
	"github.com/ellsclytn/docker-credential-helpers/credentials"
	"github.com/ellsclytn/docker-credential-helpers/wincred"
)

func main() {
	credentials.Serve(wincred.Wincred{})
}
