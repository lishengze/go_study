package main

import (
	"package_test/auth"
	"package_test/lucky"
	"package_test/route"
)

func main() {
	auth.Login("lsz")
	route.Site("lsz")
	lucky.TryTry("lsz")
}
