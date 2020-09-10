// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package main generated by go-bindata.// sources:
// ../template_amber_0/views/index.amber
// ../template_amber_0/views/layouts/main.amber
package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _indexAmber = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8d\x31\x8a\xc3\x30\x14\x44\x7b\x83\xef\x30\xe8\x00\xbb\xec\xf6\xae\xb6\xda\x3e\x9d\x71\xf1\x6d\x4d\xb0\xb1\xbe\x24\xa4\x6f\x12\xdf\x3e\x10\x07\x4c\xd2\xcd\x63\x86\x79\xbc\x1b\xa3\xaf\x08\xb2\xa7\xcd\xea\xb7\xca\x12\xbf\x44\x47\x96\xb6\x69\x9b\x31\xa4\x69\x85\xe4\xcc\xe8\xa1\x34\x69\x1b\x00\xcf\xd4\x47\x51\x76\x6e\xe5\x7e\x4b\xc5\x57\x37\xf4\x53\x8a\xc6\x68\x9d\xbb\xcc\xac\x84\x14\x42\xbc\xa7\xc7\xb8\xc3\x66\x62\x9a\x97\xe0\x61\xd4\x1c\xc4\xe8\x86\x53\xa0\x8c\xdb\xf1\x1c\x16\xfc\x1b\x15\x3f\xef\xf8\x7b\x6e\x5f\x96\xa3\xcf\xf8\x3b\x10\xd7\x92\xf4\xc3\xf0\x08\x00\x00\xff\xff\xaa\xd2\x87\xd5\xdb\x00\x00\x00")

func indexAmberBytes() ([]byte, error) {
	return bindataRead(
		_indexAmber,
		"index.amber",
	)
}

func indexAmber() (*asset, error) {
	bytes, err := indexAmberBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.amber", size: 219, mode: os.FileMode(438), modTime: time.Unix(1599264617, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _layoutsMainAmber = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x4e\xc1\xae\x83\x30\x0c\xbb\x23\xf1\x0f\xa1\x7c\x07\xf7\xf7\x01\xdc\x10\x87\x40\x23\x11\xbd\x34\x45\x34\x20\x4d\xd3\xfe\x7d\xea\xd8\x04\xf5\xa5\xae\xdd\xda\x6e\x9a\x06\x6c\x43\x4d\x6c\x1c\x15\xa5\xae\x16\x0b\x52\x57\x00\x00\x0b\xa1\x3f\x59\x86\xb1\x09\x41\xfb\xec\xf3\xf9\xaa\xab\xcb\x99\x24\xce\xff\x10\xc8\xf0\xd2\x32\xb2\x32\x28\x06\xea\x9c\xa7\x34\x6f\xbc\xe6\x0e\x37\x0e\x07\xca\x4e\x9d\xeb\x17\x4e\xc0\x09\x10\x12\x86\x55\xc8\x8d\xbf\xd4\x29\xfa\xc7\x95\x95\x77\xd0\xd6\x06\x64\xfd\xfb\xd0\xb2\x66\x97\xf2\x7e\x9f\xa4\xfb\x7d\xa8\xe7\xa3\x9d\xa3\x1a\xa9\x95\x5f\xce\xe7\x5f\xeb\x1d\x00\x00\xff\xff\x56\x99\x59\x53\x13\x01\x00\x00")

func layoutsMainAmberBytes() ([]byte, error) {
	return bindataRead(
		_layoutsMainAmber,
		"layouts/main.amber",
	)
}

func layoutsMainAmber() (*asset, error) {
	bytes, err := layoutsMainAmberBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "layouts/main.amber", size: 275, mode: os.FileMode(438), modTime: time.Unix(1599263735, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"index.amber":        indexAmber,
	"layouts/main.amber": layoutsMainAmber,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"index.amber": {indexAmber, map[string]*bintree{}},
	"layouts": {nil, map[string]*bintree{
		"main.amber": {layoutsMainAmber, map[string]*bintree{}},
	}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}