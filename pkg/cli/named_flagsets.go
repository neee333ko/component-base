package cli

import "github.com/spf13/pflag"

type NamedFlagSets struct {
	Order    []string
	FlagSets map[string]*pflag.FlagSet
}

func (nfs *NamedFlagSets) AddFlagSet(name string, fs *pflag.FlagSet) {
	if _, ok := nfs.FlagSets[name]; !ok {
		nfs.Order = append(nfs.Order, name)
		nfs.FlagSets[name] = fs
		return
	}

	nfs.FlagSets[name].AddFlagSet(fs)
}
