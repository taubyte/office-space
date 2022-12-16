package go_work

import "fmt"

func (w *goWork) Clean() error {
	w.Remove()

	err := w.Init()
	if err != nil {
		return fmt.Errorf("GoWork init failed with: %s", err)
	}

	return nil
}
