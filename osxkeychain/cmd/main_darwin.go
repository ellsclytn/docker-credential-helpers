package main

import (
	"github.com/ellsclytn/docker-credential-helpers/credentials"
	"github.com/ellsclytn/docker-credential-helpers/osxkeychain"
)

func main() {
	credentials.Serve(osxkeychain.Osxkeychain{})
}
