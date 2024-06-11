package domain

import "github.com/dguhr/arch-go-action-example/domain/bar"

func Execute() string {
	return "Hello World" + bar.Bar()
}
