package twiliserver

type Config struct {
    BindAddr          string `toml:"bind_addr"`

    TwilioAccountSID  string `toml:"twilio_account_sid"`
    TwilioAuthToken   string `toml:"twilio_auth_token"`
    TwilioCallerIDURL string `toml:"twilio_callerid_url"`
}

func NewConfig() *Config {
    return &Config{
        BindAddr: ":8080",
    }
}
