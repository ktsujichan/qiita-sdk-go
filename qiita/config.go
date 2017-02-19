package qiita

type Config struct {
	Endpoint string
}

// APIEndpoint constants
const (
	APIEndpointBase = "http://qiita.com"
)

func NewConfig() *Config {
	return &Config{
		Endpoint: APIEndpointBase,
	}
}

func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}
