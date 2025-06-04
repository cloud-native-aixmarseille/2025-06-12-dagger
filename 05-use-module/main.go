package main

import (
	"context"
)

type Test struct{}

// will run all Go tests
func (m *Test) Hello(ctx context.Context) (string, error) {
	m = Coucou{}
	m.Name = "Meetup Cloud Native Aix Marseille"
	return Meetup.DitCoucou()
}
