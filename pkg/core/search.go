package core

import (
	"github.com/alecthomas/units"
	"github.com/ip-rw/ransack/pkg/factory"
	"github.com/ip-rw/ransack/pkg/hooks"
	"github.com/koron/jvgrep/mmap"
	"github.com/paulbellamy/ratecounter"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

func maybeBinary(b []byte) bool {
	l := len(b)
	if l > 10000000 {
		l = 1024
	} else if l > 1024 {
		l /= 2
	}
	for i := 0; i < l; i++ {
		if 0 < b[i] && b[i] < 0x9 {
			return true
		}
	}
	return false
}

var binary = []string{"3dm", "3ds", "3g2", "3gp", "7z", "a", "aac", "adp", "ai", "aif", "aiff", "alz", "ape", "apk", "appimage", "ar", "arj", "asf", "au", "avi", "bak", "baml", "bh", "bin", "bk", "bmp", "btif", "bz2", "bzip2", "cab", "caf", "cgm", "class", "cmx", "cpio", "cr2", "cur", "dat", "dcm", "deb", "dex", "djvu", "dll", "dmg", "dng", "doc", "docm", "docx", "dot", "dotm", "dra", "DS_Store", "dsk", "dts", "dtshd", "dvb", "dwg", "dxf", "ecelp4800", "ecelp7470", "ecelp9600", "egg", "eol", "eot", "epub", "exe", "f4v", "fbs", "fh", "fla", "flac", "flatpak", "fli", "flv", "fpx", "fst", "fvt", "g3", "gh", "gif", "graffle", "gz", "gzip", "h261", "h263", "h264", "icns", "ico", "ief", "img", "ipa", "iso", "jar", "jpeg", "jpg", "jpgv", "jpm", "jxr", "key", "ktx", "lha", "lib", "lvp", "lz", "lzh", "lzma", "lzo", "m3u", "m4a", "m4v", "mar", "mdi", "mht", "mid", "midi", "mj2", "mka", "mkv", "mmr", "mng", "mobi", "mov", "movie", "mp3", "mp4", "mp4a", "mpeg", "mpg", "mpga", "mxu", "nef", "npx", "numbers", "nupkg", "o", "odp", "ods", "odt", "oga", "ogg", "ogv", "otf", "ott", "pages", "pbm", "pcx", "pdb", "pdf", "pea", "pgm", "pic", "png", "pnm", "pot", "potm", "potx", "ppa", "ppam", "ppm", "pps", "ppsm", "ppsx", "ppt", "pptm", "pptx", "psd", "pya", "pyc", "pyo", "pyv", "qt", "rar", "ras", "raw", "resources", "rgb", "rip", "rlc", "rmf", "rmvb", "rpm", "rtf", "rz", "s3m", "s7z", "scpt", "sgi", "shar", "snap", "sil", "sketch", "slk", "smv", "snk", "so", "stl", "suo", "sub", "swf", "tar", "tbz", "tbz2", "tga", "tgz", "thmx", "tif", "tiff", "tlz", "ttc", "ttf", "txz", "udf", "uvh", "uvi", "uvm", "uvp", "uvs", "uvu", "viv", "vob", "war", "wav", "wax", "wbmp", "wdp", "weba", "webm", "webp", "whl", "wim", "wm", "wma", "wmv", "wmx", "woff", "woff2", "wrm", "wvx", "xbm", "xif", "xla", "xlam", "xls", "xlsb", "xlsm", "xlsx", "xlt", "xltm", "xltx", "xm", "xmind", "xpi", "xpm", "xwd", "xz", "z", "zip", "zipx"}

func SearchFile(searchPath string, info os.FileInfo) {
	defer func() { cps.Incr(1) }()

	if info == nil || !info.Mode().IsRegular() || info.Size() > int64(5*units.Megabyte) {
		return
	}
	lpath := strings.ToLower(searchPath)
	for i := range binary {
		ext := path.Ext(lpath)
		if len(ext) > 0 && binary[i] == ext[1:] {
			//logrus.WithField("file", searchPath).Debugln("likely binary")
			return
		}
	}

	f, err := mmap.Open(searchPath)
	if err != nil {
		return
	}
	defer f.Close()
	max := int64(1024)
	if f.Size() < max {
		max = f.Size()
	}
	if maybeBinary(f.Data()[:max]) {
		//logrus.WithField("file", searchPath).Debugln("binary")
		return
	}

	for _, plugin := range []hooks.Plugin{&hooks.FindKey{}, &hooks.KnownHosts{}, &hooks.TargetSearch{}} {
		if c, _ := plugin.Search(factory.GetMachineClient("127.0.0.1:50051"), hooks.FileInfo{FileInfo: info, Path: searchPath}, f.Data()); !c {
			continue
		}
	}
	return
}

//
//func FindHosts() []*structs.Host {
//	hosts := []*structs.Host{{
//		Ip:   GetLocalIP(),
//		Port: 22,
//	}}
//
//	sniffed := search.SniffIps()
//	for i := range sniffed {
//		hosts = append(hosts, sniffed[i])
//	}
//	return hosts
//}

var cps = ratecounter.NewRateCounter(60 * time.Second)
var lastFile string

func init() {
	go func() {
		start := time.Now()
		for {
			logrus.WithFields(logrus.Fields{"f/s": cps.Rate() / 60, "elapsed": time.Since(start)}).Warn()
			time.Sleep(5 * time.Second)
		}
	}()
}

//func Scan(dirs []string) *structs.ScanResults {
//	//dirs = []string{"/root/", "/home/", "/storage/nuts"}
//	result := FindFiles(dirs)
//	hosts := FindHosts()
//	logrus.WithField("found", len(hosts)).Info("finished host enumeration")
//	for _, host := range hosts {
//		result.AddHost(host)
//	}
//
//	if len(result.HashedKnownHosts) > 0 {
//		logrus.WithField("hashed_known_hosts", len(result.HashedKnownHosts)).Info("cracking known hosts")
//		knownhosts.CrackKnownHosts(result)
//	}
//	logrus.WithFields(logrus.Fields{"hosts": len(result.Hosts), "known_hosts": len(result.KnownHosts), "hashed_known_hosts": len(result.HashedKnownHosts), "keys": len(result.PrivateKeys)}).Info("finished collecting")
//	return result
//}
//
//func GetLocalIP() string {
//	addrs, err := net.InterfaceAddrs()
//	if err != nil {
//		return ""
//	}
//	for _, address := range addrs {
//		// check the address type and if it is not a loopback the display it
//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
//			if ipnet.IP.To4() != nil {
//				return ipnet.IP.String()
//			}
//		}
//	}
//	return ""
//}
