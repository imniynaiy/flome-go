package main

import (
	"example.com/app-structure/internal/config"
	"example.com/app-structure/internal/database"
	"example.com/app-structure/internal/flag"
	"example.com/app-structure/internal/log"
	"example.com/app-structure/internal/server"
	"example.com/app-structure/internal/verflag"
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
