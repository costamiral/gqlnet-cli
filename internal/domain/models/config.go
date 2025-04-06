package models

type Config struct {
	SolutionName string `json:"solutionName"`
	ProjectName  string `json:"projectName"`
	DB           struct {
		Host     string `json:"host"`
		Database string `json:"database"`
		User     string `json:"user"`
		Password string `json:"password"`
		Timeout  int    `json:"timeout"`
		Encrypt  bool   `json:"encrypt"`
	} `json:"db"`
}
