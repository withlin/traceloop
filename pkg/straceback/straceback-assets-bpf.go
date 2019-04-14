// Code generated by go-bindata.
// sources:
// ../dist/straceback-main-bpf.o
// ../dist/straceback-tailcall-bpf.o
// DO NOT EDIT!

package straceback

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _stracebackMainBpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x55\xbb\x6e\x13\x41\x14\x3d\xe3\xb5\x71\x88\x53\xa4\x01\x59\x02\xa4\xd0\xa5\x5a\x23\xbe\x20\x0a\xe2\x51\xb8\x88\xa0\xa1\x1b\x0d\xd6\x24\x58\x5a\xdb\xab\x99\x01\x12\x88\x04\x4d\x28\xe9\xa0\xa3\xe1\x0b\xe8\x28\xcd\x67\x50\xf2\x09\x14\x48\xb8\x62\xd1\xcc\xde\x61\xc3\xdd\x5d\x85\xf4\xb9\xc5\x5e\x9f\x33\xf7\x31\x73\xe6\xe1\xd7\x77\xc7\xf7\x3a\x42\x20\x9a\xc0\x2f\x54\xa8\xb2\xcf\xdd\xea\xf7\x0e\x7d\x37\x20\xb0\xbc\x5a\x72\x27\x00\xf6\x00\xbc\x5a\x5f\x15\x1e\x2f\x3f\x95\x7c\xbf\x03\xac\x8a\xa2\x18\xb2\xa2\x27\xa1\x17\x70\x05\xbd\x80\x55\x52\xf2\x4b\x55\xfa\x61\xa7\x1e\xbf\x01\xe0\x0b\xe1\xf7\x7f\xe7\x0b\xac\x01\xe8\x92\xbf\xb0\xd2\x12\xd2\xe4\x42\x97\x7f\xed\xfe\xde\x18\xbf\x8b\xa2\xd8\x24\x2c\x5e\x3e\xc4\xda\xf1\x40\x5c\x03\xe0\xb9\x21\xf1\x6f\xce\xa8\xf3\x21\x7c\x13\xe4\x0d\x63\x09\x92\x1a\xb7\x08\xf5\xfb\x35\x7e\x3f\xf0\xf5\x1d\x7a\x10\xf8\x5e\x8d\x7f\x1c\xf8\x4b\x35\xbe\x1f\xf8\x7a\xdf\x6d\xf2\xfe\x9e\xf4\x28\x3f\x62\x5f\xe5\x3a\xe1\x75\x00\xe1\xca\xa5\x4e\x1f\x3a\x38\xa3\x26\x3a\x5f\x4c\xe7\x4e\x4a\x7b\x64\xa5\x9e\x3b\x6d\x90\x1a\x9d\x55\x23\x23\xa3\x5e\xf8\xc1\x89\xca\x32\x3b\xaa\xa2\x66\x2a\xb7\xa3\xc9\x81\x59\x3c\xcb\xe5\x4c\xe5\x25\x76\x6a\x9a\x49\x1f\x19\x28\xf9\x5c\x1b\x3b\x5d\xcc\x21\xb3\xe9\x44\xcf\xad\x0e\xa5\x53\xfd\x54\xee\x1b\x35\xd3\x48\xad\x33\x4e\x3d\x41\x6a\x8f\x66\xde\x8f\x77\x77\x6f\xc9\xdb\x0d\x52\x9f\xdb\xde\xd2\xdd\xe0\xb6\x49\xef\xdb\x47\xc6\xf3\xb7\x50\x9c\xd2\xee\xb4\xed\xb4\xf4\xeb\x32\x7c\xf3\x8c\x7c\x7e\xee\xf8\xc9\xb8\x01\xe0\x72\x43\x9f\xef\xb4\xa8\x2d\xc2\x03\x5a\x67\xcc\x8f\xe7\xfd\x0e\xf5\xe7\x1a\x7c\x25\xcf\xdf\x69\x3e\xff\x47\x2d\xf9\xdf\xc4\xff\xe5\xe7\x2d\xf9\x3f\x3a\xcd\xf1\x5c\xff\x83\x96\xfc\x9f\x2d\xf9\x1c\x1f\x53\x3e\xfb\x7b\xc1\x8a\x88\x6d\xc6\x73\xfd\x0f\x5b\xf4\x8f\x8d\xa2\xce\x03\x8a\xe3\xfa\xbf\x6b\xe8\xed\x6d\x8b\x16\x14\xf7\x41\xb0\xfd\x8b\xef\xd2\x9f\x00\x00\x00\xff\xff\x9c\x4c\xdf\x84\xb0\x07\x00\x00")

func stracebackMainBpfOBytes() ([]byte, error) {
	return bindataRead(
		_stracebackMainBpfO,
		"straceback-main-bpf.o",
	)
}

func stracebackMainBpfO() (*asset, error) {
	bytes, err := stracebackMainBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "straceback-main-bpf.o", size: 1968, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _stracebackTailcallBpfO = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xe4\x94\xbf\x6f\xd3\x40\x14\xc7\xbf\x6e\x93\xc6\x4d\x2b\x54\x09\x09\x05\xc3\x10\x55\xaa\x54\x06\x9c\x96\x81\x01\x09\xa9\x0b\xb0\x54\x50\x75\x2a\x93\x75\x24\x07\x44\x38\x3f\x64\xbb\x0d\xa9\x91\x3a\x20\x16\x16\x58\x90\x18\x83\xc4\xc0\xc8\x96\x6e\xe6\x4f\xe8\xc8\xd8\xb1\x63\x37\x3a\x61\xf4\xce\xcf\x71\x39\x3b\xe4\x0f\xe0\x49\xc9\xbb\xf7\xb1\xdf\xbd\x5f\x77\x3e\x7a\xb0\xfd\x70\xce\x30\x90\x8a\x81\x5f\xc8\xac\x4c\x3e\x97\xb2\xf5\x16\xff\x2f\xc3\x40\x74\x2d\x61\xef\x00\x98\x00\x22\x33\xb3\xaf\x90\x5d\xc9\xec\x32\x80\xb0\x7a\x12\x93\xfd\xc2\x04\xea\x00\x06\xac\xc3\xf7\x3f\x15\x0f\x0f\x4f\x95\x1e\x8a\x64\xa3\xd0\x3a\x53\xf6\x98\x93\x0a\xad\x8b\xe4\x3d\xeb\x5c\xe9\xe8\x4b\xc2\x2b\x06\x70\x1e\xc7\xf1\x78\x0e\x58\xe1\x78\xa4\xa3\xaf\xfc\xbc\x04\x9c\xc4\x71\x1c\x89\xc4\xae\xcd\xfd\x5d\x5f\xf4\x36\xd1\xe3\x32\xb0\xc1\xfe\xd7\x39\x6e\x15\xc0\x2b\xeb\x87\x8a\x57\x33\x80\xfd\xba\xf4\x68\xed\xdd\x5f\x6b\x85\xd6\xf1\x84\xbb\xfb\xf5\xc7\xb4\xde\xad\xaf\xb9\xa1\xf5\x7d\xc2\xef\xd5\xfb\x6d\x5a\xb7\x14\xff\x36\xe1\x81\x68\xbb\xb4\x6e\x0a\xd7\x0d\xad\x91\xe2\x83\x4a\xd2\x8f\x61\x2b\xa9\xff\x72\x7d\x23\xae\x6f\x95\xb8\xcf\x79\x97\xb3\xfe\x2e\x50\xbe\x5c\xcf\x27\xd6\xa5\x4b\x3f\x94\x0a\x06\xfb\x9f\x0b\xcd\x80\xfa\x4f\x33\xa2\xf9\xd0\x0c\x69\x7e\x34\x63\x9a\x6f\x15\x8f\x76\xb6\x01\xfc\x8e\x63\x35\x9f\x15\xf6\x33\x0e\x77\x61\xbe\x59\x32\x6e\x30\xab\x31\x1f\x15\x5d\x1e\x4d\xe6\x31\x9f\x63\x7b\x6a\x9f\x4a\x8e\x3f\x51\xdc\xcc\xf1\x65\xc5\xcb\x39\x7e\x55\xf1\xfc\xfe\xc7\x69\xde\x7c\x0c\x6e\xb2\x5d\x65\x06\x3b\x90\xaf\x03\x74\x44\xdf\x6f\xc8\x03\xd9\x0d\x7c\x04\x9e\x68\xca\x7e\xaf\xdd\x0d\x1c\xc7\x1f\xfa\x8e\xec\x06\xd2\x83\xed\x49\x37\x7b\xd2\xf0\xc4\x80\x1e\x52\x0b\xfd\x46\xf6\x96\x73\x20\x3d\xbf\xdd\xeb\xc2\x71\xdb\x4d\xd9\xf5\xa5\xf2\xb3\xe5\x4b\xe7\xb9\x27\x3a\x12\xb6\x1f\x78\x81\x78\x06\xdb\x1f\x76\x94\xf6\x7a\x2d\x11\x08\xc2\x9b\xf6\xe6\xdd\xd9\x3d\x9c\x25\x3d\xd5\xe7\xbc\x98\x7c\x07\x3e\x68\x5c\x1f\x9b\xc1\xbf\x05\x8d\x6f\x4d\x89\xa7\x5f\xad\xdb\x33\xfc\xf5\x73\xa2\x4f\xf8\x16\x80\xc5\x82\x38\x67\x5c\x54\x7a\x0e\x97\xb8\xce\xd4\x3f\xe5\x15\x8e\xaf\xf7\xe0\x94\xe3\xd6\x8c\x7f\xe7\x7f\xc4\xfe\x77\xf4\x04\x78\xc3\x55\x0d\x4f\xee\x05\xeb\xa7\x53\xe2\xaf\xce\x17\xc7\xd3\xfb\xbf\x33\xc5\x7f\x7d\x8a\xbf\x6e\x4b\xf6\xd7\x3e\xf3\xd8\x60\xff\x75\x8d\xeb\xfd\x17\x53\xfa\x7f\x51\xd0\xff\xc5\x82\xfe\x0f\x0a\x62\x93\xec\xb1\xff\x47\xb6\xd3\x1c\x53\xff\xf4\x3b\xf2\x27\x00\x00\xff\xff\x40\xfb\xe1\x74\x98\x07\x00\x00")

func stracebackTailcallBpfOBytes() ([]byte, error) {
	return bindataRead(
		_stracebackTailcallBpfO,
		"straceback-tailcall-bpf.o",
	)
}

func stracebackTailcallBpfO() (*asset, error) {
	bytes, err := stracebackTailcallBpfOBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "straceback-tailcall-bpf.o", size: 1944, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"straceback-main-bpf.o": stracebackMainBpfO,
	"straceback-tailcall-bpf.o": stracebackTailcallBpfO,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"straceback-main-bpf.o": &bintree{stracebackMainBpfO, map[string]*bintree{}},
	"straceback-tailcall-bpf.o": &bintree{stracebackTailcallBpfO, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

