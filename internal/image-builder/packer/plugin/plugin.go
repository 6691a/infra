package plugin

type Plugin struct {
	Name    string
	Version string
	Source  string
}

func New(name string, version string, source string) *Plugin {
	return &Plugin{
		Name:    name,
		Version: version,
		Source:  source,
	}
}
