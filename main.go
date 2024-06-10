package main

import (
	"github.com/theoriz0/flome-go/internal/config"
	"github.com/theoriz0/flome-go/internal/database"
	"github.com/theoriz0/flome-go/internal/flag"
	"github.com/theoriz0/flome-go/internal/log"
	"github.com/theoriz0/flome-go/internal/server"
	"github.com/theoriz0/flome-go/internal/verflag"
)

func main() {
	flag.ParseAndHandleHelpFlag()
	verflag.PrintAndExitIfRequested()
	config.ParseConfig()
	log.Init(&config.GlobalConfig.Log)
	defer log.Sync()

	log.Info("config:", log.String("file", *config.ConfigFile))
	log.Info("version:", log.String("build_date", verflag.BuildDate),
		log.String("git_version", verflag.GitVersion), log.String("git_commit", verflag.GitCommit),
		log.String("git_tree_state", verflag.GitTreeState),
	)

	database.Init(&config.GlobalConfig.Database)

	server.Start()
}
