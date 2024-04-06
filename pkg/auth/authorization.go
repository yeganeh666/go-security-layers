package auth

var accessControlList = map[string][]string{
	"user1": {},
	"user2": {"/data"},
}

func IsAuthorized(username, path string) bool {
	allowedPaths, ok := accessControlList[username]
	if !ok {
		return false
	}

	for _, allowedPath := range allowedPaths {
		if allowedPath == path {
			return true
		}
	}
	return false
}
