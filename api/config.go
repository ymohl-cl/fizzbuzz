package api

// Config api to request postgres
type Config struct {
	PostgresUser string `split_words:"true" required:"true"`
	PostgresDB   string `split_words:"true" required:"true"`
}
