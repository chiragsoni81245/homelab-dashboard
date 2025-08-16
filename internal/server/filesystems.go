package server

import (
	"embed"
)

//go:embed static/*
var staticFS embed.FS

//go:embed templates/*
var templateFS embed.FS
