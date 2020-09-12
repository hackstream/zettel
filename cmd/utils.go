package cmd

import (
	"fmt"
	"os"
)

// createFile takes a default config template file and writes to the current directory
func createFile(cfgFile []byte, configName string) error {
	f, err := os.Create(configName)
	if err != nil {
		return fmt.Errorf("error while creating default config: %v", err)
	}
	_, err = f.Write(cfgFile)
	if err != nil {
		return fmt.Errorf("error while copying default config: %v", err)
	}
	return nil
}
