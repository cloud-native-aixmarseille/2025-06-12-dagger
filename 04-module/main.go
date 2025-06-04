package main

import "context"

type Coucou struct{}

func (m *Coucou) DitCoucou(
	ctx context.Context,
	// +default="le monde"
	name string,
) (string, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", "Coucou " + name + " ! 👋👋👋"}).Stdout(ctx)
}
