package twiliserver

type Config struct {
    BindAddr          string `toml:"bind_addr"`
    DBSenumURL        string `toml:"db_senum_url"`
    DBAsterURL        string `toml:"db_aster_url"`

    TwilioAccountSID  string `toml:"twilio_account_sid"`
    TwilioAuthToken   string `toml:"twilio_auth_token"`
    TwilioCallerIDURL string `toml:"twilio_callerid_url"`
    TwilioCallbackURL string `toml:"twilio_callback_url"`
}

func NewConfig() *Config {
    return &Config{
        BindAddr: ":8080",
    }
}
