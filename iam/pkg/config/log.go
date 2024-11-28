package config

type Log struct {
	Level   string `description:"Log level set to logs." json:"level,omitempty" toml:"level,omitempty" yaml:"level,omitempty" export:"true"`
	Format  string `description:"Log format: json | common" json:"format,omitempty" toml:"format,omitempty" yaml:"format,omitempty" export:"true"`
	NoColor bool   `description:"When using the 'common' format, disables the colorized output." json:"noColor,omitempty" toml:"noColor,omitempty" yaml:"noColor,omitempty" export:"true"`

	FilePath   string `description:"Log file path. Stdout is used when omitted or empty." json:"filePath,omitempty" toml:"filePath,omitempty" yaml:"filePath,omitempty"`
	MaxSize    int    `description:"Maximum size in megabytes of the log file before it gets rotated." json:"maxSize,omitempty" toml:"maxSize,omitempty" yaml:"maxSize,omitempty" export:"true"`
	MaxAge     int    `description:"Maximum number of days to retain old log files based on the timestamp encoded in their filename." json:"maxAge,omitempty" toml:"maxAge,omitempty" yaml:"maxAge,omitempty" export:"true"`
	MaxBackups int    `description:"Maximum number of old log files to retain." json:"maxBackups,omitempty" toml:"maxBackups,omitempty" yaml:"maxBackups,omitempty" export:"true"`
	Compress   bool   `description:"Determines if the rotated log files should be compressed using gzip." json:"compress,omitempty" toml:"compress,omitempty" yaml:"compress,omitempty" export:"true"`
}
