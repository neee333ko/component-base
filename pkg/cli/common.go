package cli

import (
	"flag"
	"strings"

	"github.com/neee333ko/log"
	"github.com/spf13/pflag"
)

func WordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	r := strings.NewReplacer("_", "-", ".", "-")
	if strings.ContainsAny(name, "_.") {
		return pflag.NormalizedName(r.Replace(name))
	}

	return pflag.NormalizedName(name)
}

func WordSepNormalizeWarnFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
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
