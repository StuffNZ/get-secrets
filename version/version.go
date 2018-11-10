package version

//nolint:gochecknoglobals
var (
	application = "get-secrets"
	release     = "---"
	buildDate   = ""
)

// Application is the "friendly" name for this code
func Application() string {
	return application
}

// Release is the current version of "get-secrets"
func Release() string {
	return release
}

// BuildDate is the current build-date of "get-secrets"
func BuildDate() string {
	return buildDate
}
