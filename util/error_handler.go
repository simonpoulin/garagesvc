package util

// IsEmptyListError ...
func IsEmptyListError(err error) bool {
	if err.Error() == "mongo: no documents in result" || err.Error() == "out of pages" {
		return true
	}
	return false
}
