package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"package/gqlnet/internal/domain/models"
	"path/filepath"
	"regexp"
	"strings"
)

func createDefault(path string) error {
	defaultCfg := models.Config{
		SolutionName: "webapi-graphql",
		ProjectName:  "Api",
	}
	defaultCfg.DB = struct {
		Host     string `json:"host"`
		Database string `json:"database"`
		User     string `json:"user"`
		Password string `json:"password"`
		Timeout  int    `json:"timeout"`
		Encrypt  bool   `json:"encrypt"`
	}{
		Host:     "localhost",
		Database: "eStatements",
		User:     "username",
		Password: "password",
		Timeout:  20,
		Encrypt:  false,
	}

	data, err := json.MarshalIndent(defaultCfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func Load(path string) (*models.Config, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {

		_ = createDefault(path)
		return nil, fmt.Errorf("config file does not exist, a default one was created")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg models.Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}

func BuildConnectionString(cfg *models.Config) string {
	return fmt.Sprintf(
		"data source=%s;initial catalog=%s;uid=%s;pwd=%s;Connection Timeout=%d;Encrypt=%t;",
		cfg.DB.Host, cfg.DB.Database, cfg.DB.User, cfg.DB.Password, cfg.DB.Timeout, cfg.DB.Encrypt,
	)
}

func GetModelToDbSetMap() (map[string]string, error) {
	modelDir := "Domain/Models"
	contextPath := "Domain/Context/EStatementsContext.cs"

	files, err := os.ReadDir(modelDir)
	if err != nil {
		return nil, err
	}
	modelNames := make(map[string]bool)
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".cs" {
			modelNames[strings.TrimSuffix(file.Name(), ".cs")] = true
		}
	}

	content, err := os.ReadFile(contextPath)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`DbSet<(\w+)> (\w+)`)
	matches := re.FindAllStringSubmatch(string(content), -1)

	result := make(map[string]string)
	for _, match := range matches {
		if modelNames[match[1]] {
			result[match[1]] = match[2]
		}
	}
	return result, nil
}
