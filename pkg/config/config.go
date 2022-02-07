package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/OpenPeeDeeP/xdg"

	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	Name     string `long:"name" env:"NAME" default:"lazygit"`
	Version  string `long:"version" env:"VERSION" default:"unversioned"`
	AppState *AppState
}

type AppState struct {
	LastUpdateCheck int64
	AllRepos        []string
}

// NewAppConfig makes a new app config
func NewAppConfig(name, version string) (*AppConfig, error) {

	appState, err := loadAppState()
	if err != nil {
		return nil, err
	}

	appConfig := &AppConfig{
		Name:     "ga",
		Version:  version,
		AppState: appState,
	}

	return appConfig, nil
}

// SaveAppState marshalls the AppState struct and writes it to the disk
func (c *AppConfig) SaveAppState() error {
	marshalledAppState, err := yaml.Marshal(c.AppState)
	if err != nil {
		return err
	}

	filepath, err := configFilePath("state.yml")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, marshalledAppState, 0644)
	if err != nil && os.IsPermission(err) {
		// apparently when people have read-only permissions they prefer us to fail silently
		return nil
	}

	return err
}

// loadAppState loads recorded AppState from file
func loadAppState() (*AppState, error) {
	filepath, err := configFilePath("state.yml")
	if err != nil {
		if os.IsPermission(err) {
			// apparently when people have read-only permissions they prefer us to fail silently
			return getDefaultAppState(), nil
		}
		return nil, err
	}

	appStateBytes, err := ioutil.ReadFile(filepath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	if len(appStateBytes) == 0 {
		return getDefaultAppState(), nil
	}

	appState := &AppState{}
	err = yaml.Unmarshal(appStateBytes, appState)
	if err != nil {
		return nil, err
	}

	return appState, nil
}

func getDefaultAppState() *AppState {
	return &AppState{
		LastUpdateCheck: 0,
		AllRepos:        []string{},
	}
}

func configFilePath(filename string) (string, error) {
	folder, err := findOrCreateConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(folder, filename), nil
}

func findOrCreateConfigDir() (string, error) {
	folder := ConfigDir()
	return folder, os.MkdirAll(folder, 0755)
}

func ConfigDir() string {
	configDirectory := configDirForVendor("")
	return configDirectory
}

func configDirForVendor(vendor string) string {
	envConfigDir := os.Getenv("CONFIG_DIR")
	if envConfigDir != "" {
		return envConfigDir
	}
	configDirs := xdg.New(vendor, "ga")
	return configDirs.ConfigHome()
}
