package config

var Config config

type matcher struct {
	Name             string  `yaml:"name"`
	Regex            string  `yaml:"regex"`
	Color            string  `yaml:"color"`
	Print            string  `yaml:"print"`
	MultilineLines   int     `yaml:"multiline_lines"`
	MultilineRegex   string  `yaml:"multiline_regex"`
	MultilinePrint   string  `yaml:"multiline_print"`
	MultilineEndline *string `yaml:"multiline_endline"`
}

type config struct {
	Matcher []matcher `yaml:"beautifier"`
}
