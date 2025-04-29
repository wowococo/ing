package main

import (
	"github.com/google/uuid"
	"github.com/rs/xid"
)

func genXid() {
	guid := xid.New()

	println(guid.String())
	// Output: 9m4e2mr0ui3e8a215n4g
}

func genGoogleUuid() {
	u1 := uuid.New()
	u2, _ := uuid.NewRandom()
	println(u1.Version().String())
	println(u2.Version().String())

	println(u1.String())
	println(u2.String())
}

func main() {
	genXid()
	genGoogleUuid()
}
