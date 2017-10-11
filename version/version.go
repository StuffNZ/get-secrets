package version

const (
	application = "get-secrets"
	release     = "0.4.1.4"
)

// Application is the "friendly" name for this code
func Application() string {
	return application
}

// Release is the current version of "get-secrets"
func Release() string {
	return release
}
