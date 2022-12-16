package common

// Displayer should not return errors, it should always display something and execution should not depend on display succeeding
type Displayer interface {
	SetVerbose(bool)

	WroteDirectoriesInFile(wrote []string, file string, dir string)
}
