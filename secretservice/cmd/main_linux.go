package main

import (
	"github.com/ellsclytn/docker-credential-helpers/credentials"
	"github.com/ellsclytn/docker-credential-helpers/secretservice"
)

func main() {
	credentials.Serve(secretservice.Secretservice{})
}
