package config

type Config struct {
	Services     []string
	PollInterval int // seconds
	SMTPHost     string
	SMTPPort     string
	SMTPUser     string
	SMTPPass     string
	EmailTo      string
}
