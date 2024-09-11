package main

import (
	"dagger/geit-ravi/internal/dagger"
)

type GeitRavi struct{}

// Returns a scratch container containing a go binary built from git project
func (m *GeitRavi) BuildInScratch() *dagger.Container {
	repo := dag.Git("https://github.com/jpadams/hello").Branch("main").Tree()

	binary := dag.Container().
		From("golang:alpine").
		WithDirectory("/src", repo).
		WithWorkdir("/src").
		WithExec([]string{"go", "build", "-o", "/go/bin/hello"}).
		File("/go/bin/hello")

	return dag.Container().
		WithFile("/go/bin/hello", binary).
		WithDefaultArgs([]string{"/go/bin/hello"}).
		WithExec(nil)
}
