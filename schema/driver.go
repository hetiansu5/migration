package schema

type Driver interface {
	Run(statement string) error
}
