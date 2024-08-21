package types

// Fig Represents a plotly figure
// use the Info method to get details about the plotly version
type Fig interface {
	Info() Version
}

type Version struct {
	Name      string `yaml:"Name"`      // name of the version
	Tag       string `yaml:"Tag"`       // git tag of the plotly version
	URL       string `yaml:"URL"`       // url under which the plotly schema json file can be downloaded directly
	Path      string `yaml:"Path"`      // path under which the schema file will be saved locally for future use
	Generated string `yaml:"Generated"` // path for the generated package
	Cdn       string `yaml:"CDN"`       // url for the cdn which should be included in the head of the generated html
}
