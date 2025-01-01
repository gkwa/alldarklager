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

func (m *Alldarklager) FormatToml(ctx context.Context, source *dagger.Directory, filename string) (*dagger.Directory, error) {
	container := m.InstallProject(m.InstallPoetry(m.CreateBaseContainer()))

	containerWithSource := container.WithDirectory("/src", source)

	output := containerWithSource.
		WithWorkdir("/src").
		WithExec([]string{"toml-formatter", "check", "--fix-inplace", filename}).
		Directory("/src")

	return output, nil
}

func (m *Alldarklager) ProvisionEnvironment(ctx context.Context) *dagger.Container {
	base := m.CreateBaseContainer()
	withPoetry := m.InstallPoetry(base)
	return m.InstallProject(withPoetry)
}
