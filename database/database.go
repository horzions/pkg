package database

import (
	"fmt"

	"github.com/horzions/pkg/app"
)

type Database struct {
}

func (c *Database) Load(a *app.App) {
	fmt.Println("load Database")
}
