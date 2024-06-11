package domain

import "github.com/dguhr/arch-go-action-example/domain/baz"

func Execute() string {
	return "Hello World" + baz.Bar()
}
