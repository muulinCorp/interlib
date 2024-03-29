package auth

import (
	apiAuth "github.com/94peter/api-toolkit/auth"
)

const (
	PermGuest  = apiAuth.ApiPerm("guest")
	PermAdmin  = apiAuth.ApiPerm("admin")
	PermOwner  = apiAuth.ApiPerm("owner")
	PermEditor = apiAuth.ApiPerm("editor")
)
