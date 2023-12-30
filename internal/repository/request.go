package repository

type Request interface {
	Add(data string) error
	All() ([]string, error)
	Truncate() error
	Sync(data []string) error
}
