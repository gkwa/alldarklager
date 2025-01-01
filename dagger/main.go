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

func (m *Alldarklager) CloneRepo(container *dagger.Container) *dagger.Container {
	return container.WithExec([]string{
		"git",
		"clone",
		"https://github.com/paulovcmedeiros/toml-formatter",
	})
}

func (m *Alldarklager) InstallDependencies(container *dagger.Container) *dagger.Container {
	return container.
		WithWorkdir("toml-formatter").
		WithExec([]string{"poetry", "install"})
}

func (m *Alldarklager) ProvisionEnvironment(ctx context.Context) *dagger.Container {
	base := m.CreateBaseContainer()
	withPoetry := m.InstallPoetry(base)
	withRepo := m.CloneRepo(withPoetry)
	return m.InstallDependencies(withRepo)
}
