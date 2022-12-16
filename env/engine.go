package env

import (
	"os"
)

// Used for tests
func Reset() {
	engine = envEngine{}
}

// Support setting/getting process variables without heavy calls
var engine = envEngine{}

func (eng envEngine) get(key, _default string) string {
	value, ok := eng[key]
	if ok == false {
		tmp := os.Getenv(key)
		if len(tmp) != 0 {
			value = tmp
		} else {
			value = _default
		}

		eng.set(key, value)
	}

	return value
}

func (eng envEngine) set(key, value string) {
	eng[key] = value
}
