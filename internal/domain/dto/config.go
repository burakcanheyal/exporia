package dto

type Config struct {
	DBURL        string `mapstructure:"DB_URL"`
	DBTestUrl    string `mapstructure:"DB_TEST_URL"`
	Secret       string `mapstructure:"SECRET"`
	Secret2      string `mapstructure:"SECRET2"`
	MailAddress  string `mapstructure:"MAIL_ADDRESS"`
	MailPassword string `mapstructure:"MAIL_PASS"`
}
