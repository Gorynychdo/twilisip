package main

import (
    "flag"
    "log"

    "github.com/BurntSushi/toml"
    "github.com/Gorynychdo/twilisip/internal/app/twiliserver"
)

var (
    configPath string
)

func init() {
    flag.StringVar(&configPath, "config-path", "configs/twiliserver.toml", "path to config file")
}

func main() {
    flag.Parse()

    config := twiliserver.NewConfig()
    if _, err := toml.DecodeFile(configPath, config); err != nil {
        log.Fatal(err)
    }

    if err := twiliserver.Start(config); err != nil {
        log.Fatal(err)
    }
}
