package infra

type Executor interface {
	Exec(command string) error
}

type DbScaffolder interface {
	ScaffoldDatabase(connStr string) error
}
