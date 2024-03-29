package initialize

import (
	"fmt"

	"github.com/facebookgo/inject"
)

var dependencies = make(map[string]interface{})

func RegisterIOCs(name string, value interface{}) {
	dependencies[name] = value
}

func InitIOCs() error {
	var g inject.Graph
	var objects []*inject.Object

	for name, value := range dependencies {
		objects = append(objects, &inject.Object{
			Name:  name,
			Value: value,
		})
	}

	if err := g.Provide(objects...); err != nil {
		return fmt.Errorf("provide : %v", err)
	}

	if err := g.Populate(); err != nil {
		return fmt.Errorf("populate : %v", err)
	}
	return nil
}
