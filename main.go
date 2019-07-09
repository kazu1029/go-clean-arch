package main

import (
	"github.com/kazu1029/go-clean-arch/external"
	"github.com/kazu1029/go-clean-arch/external/mysql"
)

func main() {
	defer mysql.CloseConn()

	external.Router.Run()
}
