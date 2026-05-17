package main

import (
	"embed"

	"github.com/zkep/my-geektime/cmd"
)

//go:embed i18n/*
//go:embed web/*
//go:embed config.yml
var Assets embed.FS

func main() { cmd.Execute(Assets) }
