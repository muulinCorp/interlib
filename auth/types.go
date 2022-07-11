package auth

import "github.com/94peter/sterna/auth"

const (
	PermOwner  = auth.UserPerm("owner")
	PermBuyer  = auth.UserPerm("buyer")
	PermClassA = auth.UserPerm("classA")
	PermClassB = auth.UserPerm("classB")
)
