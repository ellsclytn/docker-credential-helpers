package main

import (
	"github.com/ellsclytn/docker-credential-helpers/credentials"
	"github.com/ellsclytn/docker-credential-helpers/pass"
)

func main() {
	credentials.Serve(pass.Pass{})
}
