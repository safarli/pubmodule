package pubmodule

import (
	"fmt"
)

func Hello() string {
	return "Hello, You!"
}

func Welcome(name string) string {
	return fmt.Sprintf("%s %s", "Welcome", name)
}
