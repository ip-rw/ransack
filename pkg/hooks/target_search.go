package hooks

import (
	"github.com/ip-rw/sshpider/pkg/search"
	"github.com/ip-rw/sshpider/pkg/structs"
	"github.com/sirupsen/logrus"
	"os"
)

type TargetSearch struct{}

func (p *TargetSearch) Log() *logrus.Entry {
	return logrus.WithField("plugin", p.Name())
}
func IsExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}
func (p *TargetSearch) Name() string {
	return "target search"
}
func (p *TargetSearch) Search(results *structs.ScanResults, file FileInfo, data []byte) (bool, error) {
	targets := search.ExtractTargets(data)
	if targets != nil {
		for _, target := range targets {
			if results.AddHost(target) {
				p.Log().WithFields(logrus.Fields{"path": file.Path, "host": target.String()}).Info("found host")
			}
		}
	}
	return true, nil
}
