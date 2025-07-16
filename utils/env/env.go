package env

import "os"

// Comment
func Get(key string, default_ ...string) string {
	value := os.Getenv(key)

	if value == "" && len(default_) != 0 {
		return default_[0]
	}
	return value
}
