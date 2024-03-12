package tilde

func homeFromEnv() (string, error) {
	if v := os.Getenv("HOME"); v != "" {
		return v, nil
	}
	return os.UserHomeDir()
}
