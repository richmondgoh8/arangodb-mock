package redirects

var globalPath string

// Setter
func SetConfigPath(basic string) {
	globalPath = basic
}

// Getter
func GetConfigPath() string {
	return globalPath
}

func UnMount() {
	globalPath = ""
}