package verflag

import (
	"fmt"
	"os"
	"strconv"

	ver "github.com/neee333ko/component-base/pkg/version"
	"github.com/neee333ko/errors"
	"github.com/neee333ko/log"
	"github.com/spf13/pflag"
)

type version int

const (
	versionRaw   = -1
	versionFalse = 0
	versionTrue  = 1
)

var versionRawMessage string = "raw"

func (v *version) Type() string {
	return "version"
}

func (v *version) String() string {
	var s string
	switch *v {
	case versionRaw:
		s = versionRawMessage
	case versionFalse:
		s = "false"
	case versionTrue:
		s = "true"
	default:
		s = "unknown"
	}

	return s
}

func (v *version) Set(s string) error {
	if s == versionRawMessage {
		*v = versionRaw
	}

	value, err := strconv.ParseBool(s)
	if err != nil {
		log.Warnf("invalid value for version flag")

		return errors.New("invalid value for version flag")
	}

	if value {
		*v = versionTrue
	} else {
		*v = versionFalse
	}

	return nil
}

var versionFlagName = "version"

func Version(fs *pflag.FlagSet, basename string) {
	p := new(version)
	fs.VarP(p, versionFlagName, "v", fmt.Sprintf("version flag for %s", basename))
	fs.Lookup(versionFlagName).NoOptDefVal = "true"
}

func AddFlag(fs *pflag.FlagSet, basename string) {
	Version(fs, basename)
}

func PrintAndExit(fs *pflag.FlagSet) {
	f := fs.Lookup(versionFlagName)
	if f.Value.String() == versionRawMessage {
		log.Infof("%#v\n", ver.Get())
		os.Exit(0)
	} else if f.Value.String() == "true" {
		log.Infof("%s\n", ver.Get())
		os.Exit(0)
	}
}
