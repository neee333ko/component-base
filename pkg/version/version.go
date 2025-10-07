package version

import (
	"fmt"
	"runtime"

	"github.com/gosuri/uitable"
	"github.com/neee333ko/component-base/pkg/json"
)

var (
	// GitVersion is semantic version.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ').
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit sha1 from git, output of $(git rev-parse HEAD).
	GitCommit = "$Format:%H$"
	// GitTreeState state of git tree, either "clean" or "dirty".
	GitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func (info *Info) String() string {
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true
	table.RightAlign(0)

	table.AddRow("GitVersion:", info.GitVersion)
	table.AddRow("GitCommit:", info.GitCommit)
	table.AddRow("GitTreeState:", info.GitTreeState)
	table.AddRow("BuildDate:", info.BuildDate)
	table.AddRow("GoVersion:", info.GoVersion)
	table.AddRow("Compiler:", info.Compiler)
	table.AddRow("Platform:", info.Platform)

	return table.String()
}

func (info *Info) ToJson() ([]byte, error) {
	return json.Marshal(info)
}

func Get() Info {
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s-%s", runtime.GOOS, runtime.GOARCH),
	}
}
