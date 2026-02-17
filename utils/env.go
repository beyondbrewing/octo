package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// CheckFileorFolderPath checks whether the given path exists on disk.
// Returns true if the path exists (file or directory), false otherwise.
func CheckFileorFolderPath(filepath string) (bool, error) {
	if strings.TrimSpace(filepath) == "" {
		return false, fmt.Errorf("filepath must not be empty")
	}

	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, fmt.Errorf("unable to stat path %q: %w", filepath, err)
}

func ReturnAbsolutePath(fp string) (string, error) {
	if strings.TrimSpace(fp) == "" {
		return "", fmt.Errorf("filepath must not be empty")
	}

	absPath, err := filepath.Abs(fp)
	if err != nil {
		return "", fmt.Errorf("unable to resolve absolute path for %q: %w", fp, err)
	}

	return absPath, nil
}

// ReadConfigutableVariables loads configuration from the given files in order.
// Later files override values from earlier files (e.g., pass .env first, then
// config.yaml so that YAML takes priority over .env defaults).
func ReadConfigutableVariables(filepaths ...string) error {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	for _, fp := range filepaths {
		exists, err := CheckFileorFolderPath(fp)
		if err != nil {
			return fmt.Errorf("checking config file: %w", err)
		}
		if !exists {
			continue // skip missing optional files
		}

		viper.SetConfigFile(fp)
		if err := viper.MergeInConfig(); err != nil {
			return fmt.Errorf("reading config file %q: %w", fp, err)
		}
	}

	return nil
}
