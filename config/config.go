package config

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the composition of yml settings.
type Config struct {
	Database struct {
		Dialect   string `default:"postgresql"`
		Host      string `default:"127.0.0.1"`
		Port      string `default:"5432"`
		Dbname    string `default:"healthy_web"`
		Username  string `default:"username"`
		Password  string `default:"password"`
		Migration bool   `default:"false"`
	}
}

const (
	// DEV represents development environment
	DEV = "develop"
	// PRD represents production environment
	PRD = "production"
)

// Load reads the settings written to the yml file
func Load(yamlFile embed.FS) (*Config, string) {
	var env *string
	if value := os.Getenv("HEALTHY_WEB_APP"); value != "" {
		env = &value
	} else {
		env = flag.String("env", "develop", "To switch configurations.")
		flag.Parse()
	}

	file, err := yamlFile.ReadFile("application." + *env + ".yml")
	if err != nil {
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(2)
	}

	config := &Config{}
	if err := yaml.Unmarshal(file, config); err != nil {
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(2)
	}

	return config, *env
}
