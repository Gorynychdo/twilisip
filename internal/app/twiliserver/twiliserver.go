package twiliserver

func Start(config *Config) error {
    srv := newServer(config)

    return srv.serve()
}
