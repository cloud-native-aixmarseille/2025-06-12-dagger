package main

import (
	"context"
	"dagger/test/internal/dagger"
)

type Test struct{}

func (m *Test) Hello(ctx context.Context) (string, error) {
	str := "Meetup Cloud Native Aix-Marseille"
	return dag.Coucou().DitCoucou(ctx, dagger.CoucouDitCoucouOpts{
		Name: str,
	})
}
