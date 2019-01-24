package api

// Config api to request postgres
type Config struct {
	PostgresHost     string `split_words:"true" required:"true"`
	PostgresPort     string `split_words:"true" required:"true"`
	PostgresUser     string `split_words:"true" required:"true"`
	PostgresPassword string `split_words:"true" required:"true"`
	PostgresDB       string `split_words:"true" required:"true"`
}

func (c Config) connSTR() string {
	var connSTR string

	connSTR = "host=" + c.PostgresHost + " "
	connSTR += "port=" + c.PostgresPort + " "
	connSTR += "user=" + c.PostgresUser + " "
	connSTR += "password= " + c.PostgresPassword + " "
	connSTR += "dbname=" + c.PostgresDB + " "
	connSTR += "sslmode=disable"
	return connSTR
}
