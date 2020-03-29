package schema

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _input_admin_login_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xca\xcc\x2b\x28\x2d\x51\x70\x4c\xc9\xcd\xcc\xf3\xc9\x4f\xcf\xcc\x53\xa8\xe6\x52\x50\x48\xcd\x4d\xcc\xcc\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\x28\x48\x2c\x2e\x2e\xcf\x2f\x4a\x41\x88\xd5\x72\x01\x02\x00\x00\xff\xff\x61\xed\xb3\xbe\x3a\x00\x00\x00")

func input_admin_login_graphql() ([]byte, error) {
	return bindata_read(
		_input_admin_login_graphql,
		"input/admin_login.graphql",
	)
}

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\xcb\x31\x0e\x82\x40\x10\x05\xd0\x7e\x4f\xf1\x89\x8d\x36\x1c\x80\x8e\x5e\x3b\x2f\xb0\xb2\xa3\x4c\x64\x67\x30\xfc\x49\xdc\x18\xef\x6e\x63\x68\x5f\xf2\x0e\xa8\xc1\x4c\x75\xdb\x30\x65\x43\xf5\xa2\xf7\x06\xce\x82\x72\xeb\x21\x6f\x99\x82\x52\xb0\xc9\x2b\xc4\xa8\x79\x59\x5a\x9f\xd8\x56\xc1\xe5\x1f\xf1\x49\x40\x2e\x55\xed\xec\x0f\xb5\xa3\xda\x1a\x1c\x30\xee\xd2\x9d\x06\x8c\xc1\xf9\xea\x4f\xb1\x2e\x7d\xd3\x2f\x00\x00\xff\xff\x9a\xd9\x42\xa0\x75\x00\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xc1\x0d\xc2\x30\x0c\x05\xd0\x7b\xa6\xf8\x88\x7b\x07\x60\x0b\xc4\x04\x40\x7e\xa9\xa5\xd4\x86\xd8\x96\xb0\x10\xbb\xf7\x9d\xf1\x49\x4e\xa1\xc3\x37\xcb\xd1\xa1\x16\xd8\xad\xcb\x5a\x88\x8d\xe8\x8f\x05\xfc\xf2\x99\xc1\x0e\x97\x3d\x47\xdc\x95\x96\x3e\x6a\x69\x51\x6f\xe2\x9a\x9c\x85\x5f\x03\x44\x57\xbb\xe0\x16\x53\xf4\x75\x6a\xff\x76\x04\x00\x00\xff\xff\x4d\x33\x0f\x06\x5c\x00\x00\x00")

func query_graphql() ([]byte, error) {
	return bindata_read(
		_query_graphql,
		"query.graphql",
	)
}

var _schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\x0a\x0a\xb9\xa5\x25\x89\x25\x99\xf9\x79\x56\x0a\xbe\x50\x16\x57\x2d\x17\x20\x00\x00\xff\xff\x1f\x5d\x4a\xc1\x2f\x00\x00\x00")

func schema_graphql() ([]byte, error) {
	return bindata_read(
		_schema_graphql,
		"schema.graphql",
	)
}

var _type_admin_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x4c\xc9\xcd\xcc\x0b\x2d\x4e\x2d\x52\xa8\xe6\x52\x50\x48\xcd\x4d\xcc\xcc\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x57\xe4\x52\x50\x48\xcb\x2c\x2a\x2e\xf1\x4b\xcc\x4d\x45\x16\xcc\x49\x44\x17\xab\xe5\x02\x04\x00\x00\xff\xff\x16\xe5\x7a\x07\x4d\x00\x00\x00")

func type_admin_graphql() ([]byte, error) {
	return bindata_read(
		_type_admin_graphql,
		"type/admin.graphql",
	)
}

var _type_auth_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x2c\x2d\xc9\x08\xc9\xcf\x4e\xcd\x53\xa8\xe6\x52\x50\x28\x01\xb1\xac\x14\x82\x4b\x8a\x32\xf3\xd2\x15\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\xa3\x33\x16\xdc\x24\x00\x00\x00")

func type_auth_graphql() ([]byte, error) {
	return bindata_read(
		_type_auth_graphql,
		"type/auth.graphql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"input/admin_login.graphql": input_admin_login_graphql,
	"mutation.graphql": mutation_graphql,
	"query.graphql": query_graphql,
	"schema.graphql": schema_graphql,
	"type/admin.graphql": type_admin_graphql,
	"type/auth.graphql": type_auth_graphql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"input": &_bintree_t{nil, map[string]*_bintree_t{
		"admin_login.graphql": &_bintree_t{input_admin_login_graphql, map[string]*_bintree_t{
		}},
	}},
	"mutation.graphql": &_bintree_t{mutation_graphql, map[string]*_bintree_t{
	}},
	"query.graphql": &_bintree_t{query_graphql, map[string]*_bintree_t{
	}},
	"schema.graphql": &_bintree_t{schema_graphql, map[string]*_bintree_t{
	}},
	"type": &_bintree_t{nil, map[string]*_bintree_t{
		"admin.graphql": &_bintree_t{type_admin_graphql, map[string]*_bintree_t{
		}},
		"auth.graphql": &_bintree_t{type_auth_graphql, map[string]*_bintree_t{
		}},
	}},
}}
