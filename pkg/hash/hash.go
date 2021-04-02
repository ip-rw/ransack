package hash

import (
	"encoding/gob"
	"fmt"
	"hash"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"leb.io/aeshash"
)

type Kfe struct {
	root  int
	level int
	Path  string
	Size  int64
	Hash  uint64
}

var (
	zeroHash   = aeshash.HashStr("", 0) // the null Hash (no bytes passed)
	verbose    = false
	ignoreList = []string{".DS_Store", ".Spotlight-V100", ".fseventsd", ".git", ".idea", "Thumbs.db"}
)

func fullName(path string, fi os.FileInfo) string {
	p := ""
	if path == "" {
		p = fi.Name()
	} else {
		p = path + "/" + fi.Name()
	}
	return p
}

type FsHash struct {
	lock     *sync.RWMutex
	hf       hash.Hash64 //= aeshash.NewAES(0)
	Hmap     map[uint64][]Kfe
	Pmap     map[string]uint64
	Roots    []string
	Scanning bool
}

func NewFsHash(roots []string) *FsHash {
	return &FsHash{
		lock:  &sync.RWMutex{},
		Hmap:  make(map[uint64][]Kfe, 100),
		Pmap:  make(map[string]uint64),
		hf:    aeshash.NewAES(0),
		Roots: roots,
	}
}

// calculate Hash using name+Size only!
func (fs *FsHash) SaveGob(fp string) error {
	f, err := os.Create(fp)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	if err := gob.NewEncoder(f).Encode(fs); err != nil {
		return fmt.Errorf("%v\n", err)
	}
	return nil
}

func Load(fp string) (*FsHash, error) {
	var fs FsHash
	f, err := os.Open(fp)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	enc := gob.NewDecoder(f)
	err = enc.Decode(&fs)
	fs.lock = &sync.RWMutex{}
	fs.hf = aeshash.NewAES(0)
	if err != nil {
		return nil, err
	}
	return &fs, nil
}

func (f *FsHash) fileHash(fi os.FileInfo) (r uint64) {
	//p := fullName(Path, fi)
	//fmt.Printf("fileHash: Path=%q fi.Name=%q, p=%q Size=%d\n", Path, fi.Name(), p, fi.Size())
	if fi.Size() == 0 {
		return zeroHash
	}
	if f.hf == nil {
		f.hf = aeshash.NewAES(0)
	}
	f.hf.Reset()
	f.hf.Write([]byte(fi.Name() + fmt.Sprintf("%d", fi.Size())))
	r = f.hf.Sum64()
	//fmt.Printf("fileHash: file=%q, Hash=0x%016x\n", p, r)
	return r
}

// add a Kfe to the Hash map. check for inline
func (f *FsHash) add(hash uint64, k *Kfe) {
	f.lock.Lock()
	defer f.lock.Unlock()
	//fmt.Printf("add: Kfe=%v\n", k)
	_, ok := f.Hmap[hash]
	if !ok {
		f.Hmap[hash] = []Kfe{*k}
	} else {
		f.Hmap[hash] = append(f.Hmap[hash], *k)
	}
	f.Pmap[k.Path] = hash
}

// addDir a file entry to the Hash map.
func (f *FsHash) addFile(root, level int, path string, fi os.FileInfo, hash uint64, size int64) {
	p := fullName(path, fi)
	if verbose {
		fmt.Printf("addFile: Hash=%016x, p=%q\n", hash, path)
	}
	k1 := Kfe{root, level, p, fi.Size(), 0}
	f.add(hash, &k1)
}

// addDir a directory entry to the Hash map.
func (f *FsHash) addDir(root, level int, path string, fi os.FileInfo, hash uint64, size int64) {
	p := fullName(path, fi)
	//fmt.Printf("addDir: Path=%q, fi.Name()=%q, p=%q, Size=%d, level=%d, Hash=0x%016x\n", Path, fi.Name(), p, Size, level, Hash)
	if verbose {
		fmt.Printf("addDir : Hash=%016x, p=%q\n", hash, p)
	}
	k1 := Kfe{root, level, p, size, hash}
	f.add(hash, &k1)
}

// descent recursively descends the directory hierarchy.
func (f *FsHash) descend(root int, path string, fis []os.FileInfo, callback func(path string, hash uint64, level int)) (uint64, int64) {

	var level = -1
	var des func(root int, path string, fis []os.FileInfo) (uint64, int64)
	des = func(root int, path string, fis []os.FileInfo) (uint64, int64) {
		var hash uint64
		var sizes int64
		var gh = aeshash.NewAES(0)
		level++
	outer:
		for _, fi := range fis {
			//fmt.Printf("descend: enter fi.Name=%q\n", fi.Name())
			switch {
			case fi.Mode()&os.ModeDir == os.ModeDir:
				//fmt.Printf("descend: dir: Path=%q, fi.Name()=%q\n", Path, fi.Name())
				for _, name := range ignoreList {
					if fi.Name() == name {
						//fmt.Printf("descend: skip dir=%q\n", fi.Name())
						continue outer
					}
				}
				p := fullName(path, fi)
				//fmt.Printf("descend: dir=%q\n", p)
				d, err := os.Open(p)
				if err != nil {
					continue
				}
				fis, err := d.Readdir(-1)
				if err != nil || fis == nil {
					//fmt.Printf("descend: can't read %q\n", fullName(Path, fi))
					continue
				}
				d.Close()
				h, size := des(root, p, fis)
				hash = h
				gh.Write64(hash)
				sizes += size
				f.addDir(root, level, path, fi, hash, size)
			case fi.Mode()&os.ModeType == 0:
				for _, name := range ignoreList {
					if fi.Name() == name {
						//fmt.Printf("descend: skip file=%q\n", fi.Name())
						continue outer
					}
				}
				if fi.Size() >= 0 {
					hash = f.fileHash(fi)
					if hash == 0 {
						continue
					}
					gh.Write64(hash)
					sizes += fi.Size()
					//fmt.Printf("descend: file: Path=%q, fi.Name()=%q, Hash=%016x, Size=%d, sizes=%d\n", Path, fi.Name(), Hash, Size, sizes)
				}
			default:
				continue
			}
		}
		hashes := gh.Sum64()
		level--
		if callback != nil {
			defer callback(path, hashes, level)
		}
		return hashes, sizes
	}
	return des(root, path, fis)
}

// scan the Roots (dirs) and files passed on the command line and records their hashes in a map.
func (f *FsHash) LookupHash(hash uint64) []Kfe {
	f.lock.RLock()
	defer f.lock.RUnlock()

	if v, ok := f.Hmap[hash]; ok {
		return v
	}
	return nil
}
func (f *FsHash) LookupPath(path string) uint64 {
	f.lock.RLock()
	defer f.lock.RUnlock()

	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	if v, ok := f.Pmap[abs]; ok {
		return v
	}
	return 0
}

func (f *FsHash) Scan(callback func(path string, hash uint64, level int)) {
	f.Scanning = true
	for k, path := range f.Roots {
		fi, err := os.Stat(path)
		if err != nil || fi == nil {
			fmt.Printf("fi=%#v, err=%v\n", fi, err)
			panic("bad")
		}
		prefix := ""
		idx := strings.LastIndex(path, "/")
		if idx != -1 {
			prefix = path[0:idx]
		}
		fis := []os.FileInfo{fi}
		f.descend(k, prefix, fis, callback)
	}
	f.Scanning = false
}
