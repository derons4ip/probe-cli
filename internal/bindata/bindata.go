// Code generated by go-bindata. DO NOT EDIT.
// sources:
// data/default-config.json
// data/migrations/1_create_msmt_results.sql

package bindata


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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataDataDefaultconfigjson = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xcd\x6e\xe3\x38\x0c\xc7\xef\x7d\x0a\x41\xe7\x3a\xcd\x62\x6f\x39" +
	"\xee\x6d\x0f\xbb\x1d\x60\xe6\x56\x14\x82\x62\xd1\x36\x31\x32\xa9\x11\xe9\x64\x82\x41\xdf\x7d\x20\x35\x89\xed\x7e" +
	"\x4d\x8f\xe2\x9f\xa2\xc8\x1f\x29\xfe\xba\x31\xc6\x3a\xbb\x33\xf6\xdb\x80\x62\x50\xcc\x89\xa7\x6c\xee\xef\xff\xff" +
	"\xd7\x7c\xc9\xbc\x07\xd3\x32\x75\xd8\x9b\x0e\x23\x6c\xcc\x57\x00\x33\xa8\x26\xd9\xdd\xdd\x31\x13\x6e\x90\xef\x06" +
	"\x88\xa9\x1e\x52\xf1\x6f\xda\x88\xa6\xe3\x6c\x8a\xd9\xde\x96\xf0\x7e\x52\x76\x53\x0a\x5e\xc1\xee\x8c\xe6\x09\xaa" +
	"\x59\x06\x9f\x91\x7a\xbb\x33\x25\x09\x63\x2c\x52\x1b\xa7\x00\x0e\x93\xdd\x99\xce\x47\xa9\x7e\x0b\xc1\x0b\x2d\x02" +
	"\x2c\x84\x3e\xc9\x5a\x98\x52\x64\x1f\x5c\x06\x99\xa2\xbe\xd0\x04\x28\xb8\x36\x7b\x19\x5c\x86\xc4\xf9\xaa\xdf\x18" +
	"\xf3\x54\x33\x23\x56\xec\xb0\xf5\x8a\x4c\x32\xe7\x07\xe4\xf7\x11\xc2\x3a\x5a\xf5\x3d\x39\x26\xa7\x20\xea\x5a\x1e" +
	"\x53\x84\x72\xf1\x3d\x37\x82\xa3\x5c\xea\xbb\xbe\x58\x10\x8d\x5e\x21\xd4\x28\x2b\x2a\xf3\xab\x4b\x24\x67\x6b\x75" +
	"\x2f\xe1\x1e\xaa\xd9\x18\x7b\x84\x7d\xd3\x32\x11\xb4\x8a\x07\xd4\x93\xbd\xbd\x28\x9d\x6f\x61\xcf\xfc\xbd\x19\x41" +
	"\x04\xa8\x87\x3c\x6b\xc7\xc1\xab\xf8\x94\x66\x8b\x42\x84\x3e\xfb\x71\xb6\x04\x2f\xc3\x7c\xa2\xa0\xf3\xa1\x8c\x44" +
	"\x83\x74\xf0\x11\x43\x93\xe1\xc7\x04\xa2\x4d\x44\x82\x17\x2e\x03\xf8\x00\xb9\xe9\x10\x62\x68\x46\x4f\x98\xa6\x58" +
	"\x29\xdb\xea\xf6\x78\x2e\x6e\x64\xd2\x21\x9e\x9c\x8f\x91\x8f\x9e\xda\x32\x36\xf6\xef\xed\xf6\xbf\x7f\xec\x95\x58" +
	"\xa5\x2d\xa0\x05\xd6\xa2\x47\x47\xd8\x0b\x2a\xcc\x96\x05\xab\xd6\x2b\xf4\x9c\xb1\xaa\x0f\x8f\x55\x7e\xba\x4e\x92" +
	"\xa8\x27\x75\x85\x8d\xef\x97\x0d\xf8\x00\xf6\xc7\x50\xdf\xc2\xba\x04\x7b\x36\xad\xf3\x48\x90\x3b\xce\xe3\xb9\xe8" +
	"\xcf\x64\x50\x1a\x71\x09\xb5\xec\x8e\x13\xc8\x07\xc8\x05\x5d\x99\x2e\xfb\x86\xe6\xca\xf4\xbf\x76\x28\x8d\x7e\xf7" +
	"\xf6\x42\x5c\x5f\x5f\x95\x31\x62\x08\x11\xf6\xfc\xf3\x93\x45\xfc\x79\x80\x3e\x39\x42\x57\x9e\xf3\xd7\x0a\x87\x42" +
	"\x33\xbc\xde\x33\x2d\x4f\xa4\xf9\xf4\x62\x73\x08\xb8\xc0\xa3\x47\x72\x5d\x66\x3a\xff\xc5\xd5\x7a\x70\xcf\x2b\xd1" +
	"\x1d\x20\xcb\xf3\x47\xb7\xdb\xcd\x76\xf3\xd7\xf3\xb6\x73\x48\xa5\x83\x65\xde\x98\x04\x48\x2f\xd7\x9f\x6e\x7e\x07" +
	"\x00\x00\xff\xff\x0f\x7e\x15\xb3\x6d\x05\x00\x00")

func bindataDataDefaultconfigjsonBytes() ([]byte, error) {
	return bindataRead(
		_bindataDataDefaultconfigjson,
		"data/default-config.json",
	)
}



func bindataDataDefaultconfigjson() (*asset, error) {
	bytes, err := bindataDataDefaultconfigjsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "data/default-config.json",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataDataMigrations1createmsmtresultssql = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\x4f\x6f\xf2\x30\x0c\xc6\xef\xfd\x14\x3e\x82\xde\x97\xc3\x26\x71" +
	"\xe2\x14\x5a\x6f\xeb\x56\x52\x94\xa6\xd3\x38\xb5\xd1\x1a\x50\x34\x9a\x56\x69\x22\xb4\x6f\x3f\x85\x7f\x2a\xac\x70" +
	"\xde\xd5\xbf\xc7\x8e\xfd\xd8\x99\x4c\xe0\x5f\xad\x36\x46\x58\x09\x51\xb3\xd3\x41\x3f\x90\x59\x61\x65\x2d\xb5\x9d" +
	"\xcb\x8d\xd2\x41\x10\xb1\x74\x09\x9c\xcc\x13\x84\xd2\xc8\xce\x6d\x6d\x57\xce\x2e\xa2\xb5\x14\x9d\x33\xfb\x1c\x8f" +
	"\x86\xab\xa1\xae\x2e\x49\xde\xde\x7d\x36\x64\x48\x38\x5e\x3f\x0c\xa3\x00\x00\xa0\x54\x55\x09\x31\xe5\xf8\x8c\x0c" +
	"\x96\x2c\x5e\x10\xb6\x82\x37\x5c\x01\xc9\x79\x1a\xd3\x90\xe1\x02\x29\xff\x7f\xd0\x6a\x51\xcb\x12\xde\x09\x0b\x5f" +
	"\x08\x1b\x3d\x4e\xa7\xe3\x23\xe8\xac\x30\xb6\xb0\xca\xe3\x88\x70\xe4\xf1\x02\x8f\x48\xea\x6a\x18\x74\xae\xae\x85" +
	"\xf9\x2e\xe1\x35\x4b\xe9\x31\x56\x35\x5a\x96\xc0\x63\xba\x8a\x29\x1f\x3d\x9c\xca\x57\xc2\x8a\xc2\x75\x62\x23\x0b" +
	"\xd7\x9e\xdb\xfd\x0d\xab\x66\xa7\xcf\x38\x18\xcf\xae\x67\xbf\xb0\xf7\x4f\x1a\xa0\xda\xc1\xf2\xa2\xd3\xd7\x63\x7f" +
	"\x36\x4e\x5b\x9f\x7d\x96\x9f\xc4\x5a\xda\x5d\x63\xbe\x8a\x7b\xcd\x5a\x6f\x33\x7e\x9c\xe6\x5a\x0b\xb5\x75\x66\x58" +
	"\x6d\x64\xdb\x18\x5b\xac\xd5\xf6\x2e\xf7\x36\x0e\x50\xa5\x5b\x67\x07\x49\x6f\x19\xb7\x92\x0f\xb7\x5a\xf4\x37\xc4" +
	"\xf0\x09\x19\xd2\x10\xb3\xfe\x29\xfb\x25\x8e\x21\xa5\x10\x61\x82\x1c\x21\x43\x0e\x34\x4f\x12\x1f\xca\x97\xde\x77" +
	"\x08\x49\x16\x92\x08\xf7\x57\x71\xf3\x5b\xfd\x04\x00\x00\xff\xff\xc4\x16\x0a\x20\xcf\x03\x00\x00")

func bindataDataMigrations1createmsmtresultssqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataDataMigrations1createmsmtresultssql,
		"data/migrations/1_create_msmt_results.sql",
	)
}



func bindataDataMigrations1createmsmtresultssql() (*asset, error) {
	bytes, err := bindataDataMigrations1createmsmtresultssqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "data/migrations/1_create_msmt_results.sql",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"data/default-config.json":                  bindataDataDefaultconfigjson,
	"data/migrations/1_create_msmt_results.sql": bindataDataMigrations1createmsmtresultssql,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"data": {Func: nil, Children: map[string]*bintree{
		"default-config.json": {Func: bindataDataDefaultconfigjson, Children: map[string]*bintree{}},
		"migrations": {Func: nil, Children: map[string]*bintree{
			"1_create_msmt_results.sql": {Func: bindataDataMigrations1createmsmtresultssql, Children: map[string]*bintree{}},
		}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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