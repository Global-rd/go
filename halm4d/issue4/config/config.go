package config

import "flag"

type Config struct {
	verbose     bool
	help        bool
	zip         bool
	sourceType  string
	sourceFile  string
	output      string
	ollamaUrl   string
	ollamaModel string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Parse() *Config {
	flag.BoolVar(&c.help, "help", false, "show help")
	flag.BoolVar(&c.help, "h", false, "show help")

	flag.BoolVar(&c.zip, "zip", true, "zip the output")
	flag.BoolVar(&c.zip, "z", true, "zip the output")

	flag.StringVar(&c.sourceType, "source-type", "ollama", "source type can be 'ollama' or 'file'")
	flag.StringVar(&c.sourceType, "st", "ollama", "source type can be 'ollama' or 'file'")

	flag.StringVar(&c.sourceFile, "source", "source.json", "source file")
	flag.StringVar(&c.sourceFile, "s", "source.json", "source file")

	flag.StringVar(&c.output, "output", "output", "output file")
	flag.StringVar(&c.output, "o", "output", "output file")

	flag.BoolVar(&c.verbose, "verbose", false, "verbose mode")
	flag.BoolVar(&c.verbose, "v", false, "verbose mode")

	flag.StringVar(&c.ollamaUrl, "ollama-url", "http://localhost:11434", "ollama URL to use for generating responses")
	flag.StringVar(&c.ollamaModel, "ollama-model", "llama3.2", "ollama model to use for generating responses")

	flag.Parse()
	return c
}

func (c *Config) Verbose() bool {
	return c.verbose
}

func (c *Config) Help() bool {
	return c.help
}

func (c *Config) Zip() bool {
	return c.zip
}

func (c *Config) SourceType() string {
	return c.sourceType
}

func (c *Config) SourceFile() string {
	return c.sourceFile
}

func (c *Config) Output() string {
	return c.output
}

func (c *Config) OllamaUrl() string {
	return c.ollamaUrl
}

func (c *Config) OllamaModel() string {
	return c.ollamaModel
}

func (c *Config) PrintHelp() {
	flag.PrintDefaults()
}
