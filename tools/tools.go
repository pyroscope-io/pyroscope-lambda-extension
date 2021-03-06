//go:build tools
// +build tools

// Package tools is used to describe various tools we use.
// Think of it as "dev-dependencies" in ruby or node projects
// See: https://marcofranssen.nl/manage-go-tools-via-go-modules/
// See Makefile for an example of how it's used
package tools

import (
	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/mgechev/revive"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
