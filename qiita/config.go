package qiita

type Config struct {
	Endpoint string
}

// APIEndpoint constants
const (
	APIEndpointBase = "http://qiita.com/api/v2"
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
