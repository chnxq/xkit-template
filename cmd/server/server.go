package main

import (
	"context"
	"fmt"

	"xkit-template-v01/internal/bootstrap"
)

var ConfigPath string

func runServer() error {
	app, cleanup, err := bootstrap.Initialize(context.Background(), bootstrap.Options{
		Name:       Name,
		Version:    Version,
		BuildTime:  BuildTime,
		GitCommit:  GitCommit,
		ConfigPath: ConfigPath,
	})
	if err != nil {
		return err
	}
	defer cleanup()

	fmt.Printf("starting %s %s\n", Name, Version)
	return app.Run()
}
