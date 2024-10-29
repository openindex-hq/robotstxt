package robotstxt

func IsValidUserAgentToObey(userAgent string) bool {
	return NewRobotsMatcher().isValidUserAgentToObey(userAgent)
}
