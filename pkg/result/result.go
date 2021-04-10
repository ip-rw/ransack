package result

import (
	"context"
	"fmt"
	"github.com/ip-rw/ransack/pkg/models"
	"github.com/ip-rw/ransack/pkg/proto"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"os/user"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

type ScanResults struct {
	MachineID string
	Client    proto.MachineServiceClient
}

func GetOwner(path string) string {
	info, err := os.Stat(path)
	if err != nil {
		return ""
	}
	var uid int
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		uid = int(stat.Uid)
	} else {
		uid = os.Getuid()
	}
	user, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		return ""
	}
	return user.Username
}

func (s *ScanResults) GetMachineRequest() *proto.MachineRequest {
	return &proto.MachineRequest{MachineID: s.MachineID}
}

func (s *ScanResults) AddUser(username string) *proto.AddResult {
	res, err := s.Client.AddUser(context.Background(), &proto.AddUserRequest{
		Request: s.GetMachineRequest(),
		User: &proto.User{
			Username: username,
		},
	})
	if err != nil {
		return &proto.AddResult{
			Added: false,
			Id:    -1,
		}
	}
	return res
}

func (s *ScanResults) AddPrivateKey(user string, path string, data []byte) models.AddResult {
	ures := s.AddUser(user)
	key, err := models.ParseKey(path, data)
	if err != nil {
		return proto.AddResult{
			Added: false,
			Id:    -1,
		}
	}
	s.Client.AddKey(context.Background(), &proto.AddKeyRequest{
		Request: s.GetMachineRequest(),
		Key: &proto.Key{
			Data:          data,
			Path:          path,
			UserID:        ures.Id,
			CryptPassword: "",
		},
	})
}

func (s *ScanResults) AddHost(host *Host) bool {
	if host != nil {
		if _, loaded := s.dupMap.LoadOrStore("Host"+host.String(), 1); !loaded {
			if host.User != "git" && host.Ip != "openssh.com" && host.Ip != "github.com" && host.Ip != "bitbucket.org" && !strings.Contains(host.String(), "example") {
				//logrus.WithFields(logrus.Fields{"host": host.String(), "user": host.User}).Info("host added")
				s.Hosts = append(s.Hosts, host)
				return true
			}
		}
	}
	return false
}

func (s *ScanResults) AddKnownHost(knownhost *KnownHost) bool {
	if knownhost != nil {
		if knownhost.Hostname != "" {
			h := ParseHost(fmt.Sprintf("%s:%d", knownhost.Hostname, knownhost.Port))
			if knownhost.Owner != "" {
				h.User = knownhost.Owner
			}
			if s.AddHost(h) {
				return true
			}
		}

		if knownhost.IP != "" {
			h := ParseHost(fmt.Sprintf("%s:%d", knownhost.IP, knownhost.Port))
			if knownhost.Owner != "" {
				h.User = knownhost.Owner
			}
			if s.AddHost(h) {
				return true
			}
		}

		if knownhost.Hash != nil && knownhost.Salt != nil {
			if _, loaded := s.dupMap.LoadOrStore("knownhost"+string(knownhost.Hash), 1); !loaded {
				logrus.WithFields(logrus.Fields{"hash": string(knownhost.Hash)}).Info("host hashed known host found")
				s.HashedKnownHosts = append(s.HashedKnownHosts, knownhost)
			}
			return true
		}
	}
	return false
}
