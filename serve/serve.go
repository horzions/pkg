package serve

type Serve struct {
	Port   string `yaml:"port"`
	JwtKey string `yaml:"jwtKey"`
	Salt   string `yaml:"salt"`
}
