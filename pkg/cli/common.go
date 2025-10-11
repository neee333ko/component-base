package cli

import (
	"flag"
	"fmt"
	"strings"

	"github.com/neee333ko/log"
	"github.com/spf13/pflag"
)

func AddHelpFlag(fs *pflag.FlagSet, basename string) {
	fs.BoolP("help", "h", false, fmt.Sprintf("help flag for %s", basename))
}

func WordSepNormalizeFunc(fs *pflag.FlagSet, name string) pflag.NormalizedName {
	r := strings.NewReplacer("_", "-", ".", "-")
	if strings.ContainsAny(name, "_.") {
		return pflag.NormalizedName(r.Replace(name))
	}

	return pflag.NormalizedName(name)
}

func WordSepNormalizeWarnFunc(fs *pflag.FlagSet, name string) pflag.NormalizedName {
	r := strings.NewReplacer("_", "-", ".", "-")
	if strings.ContainsAny(name, "_.") {
		log.Warnf("flag name should use '-' as word separator but found flag: %s\n", name)

		return pflag.NormalizedName(r.Replace(name))
	}

	return pflag.NormalizedName(name)
}

func InitFS(fs *pflag.FlagSet) {
	fs.SetNormalizeFunc(WordSepNormalizeFunc)
	fs.AddGoFlagSet(flag.CommandLine)
}

func PrintFlagSet(fs *pflag.FlagSet) string {
	s := "\n"
	fs.VisitAll(func(f *pflag.Flag) {
		s += fmt.Sprintf("--%s: %s\n", f.Name, f.Value.String())
	})

	return s
}
