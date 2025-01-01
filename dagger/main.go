package main

import (
	"context"

	"dagger/alldarklager/internal/dagger"
)

type Alldarklager struct{}

func (m *Alldarklager) CreateBaseContainer() *dagger.Container {
	return dag.Container().From("python:3.12")
}

func (m *Alldarklager) InstallPoetry(base *dagger.Container) *dagger.Container {
	return base.WithExec([]string{"pip", "install", "poetry"})
}

func (m *Alldarklager) InstallProject(container *dagger.Container) *dagger.Container {
	return container.
		WithExec([]string{"poetry", "new", "formatter"}).
		WithWorkdir("formatter").
		WithExec([]string{"poetry", "config", "virtualenvs.create", "false"}).
		WithExec([]string{
			"poetry",
			"add",
			"--python=>=3.12,<3.13",
			"git+https://github.com/paulovcmedeiros/toml-formatter.git",
		})
}

func (m *Alldarklager) RunTomlFormatter(ctx context.Context, source *dagger.Directory, args []string) (*dagger.Directory, error) {
	container := m.InstallProject(m.InstallPoetry(m.CreateBaseContainer()))

	containerWithSource := container.WithDirectory("/src", source)

	cmd := append([]string{"toml-formatter"}, args...)

	output := containerWithSource.
		WithWorkdir("/src").
		WithExec(cmd).
		Directory("/src")

	return output, nil
}

func (m *Alldarklager) Debug(ctx context.Context, source *dagger.Directory) *dagger.Container {
	container := m.InstallProject(m.InstallPoetry(m.CreateBaseContainer()))
	return container.
		WithDirectory("/src", source).
		WithWorkdir("/src").
		Terminal()
}
