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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8c\x3d\x0a\xc2\x40\x10\x46\xfb\x85\xbd\xc3\x27\x36\xda\xe4\x00\xe9\x52\x0a\xa6\xf3\x02\x63\x76\x4c\x06\xb3\xb3\xd1\xcc\x82\x41\xbc\xbb\x04\xe2\x1f\x96\xf3\xe6\xbd\x6f\x8d\x98\x8d\x4c\x92\x8e\x68\x48\x11\x53\x90\xd3\x04\xeb\x18\xe1\x58\x80\x6f\xdc\x64\xe3\x80\x91\x2f\x99\xd5\x84\xfa\x7e\x2a\xbc\xb3\x69\x60\xd4\x4b\x89\xbb\x77\x00\x85\x28\xba\x4f\xad\xe8\x46\x74\xc8\x56\xa2\x7a\x93\xd5\xb6\x44\x95\xad\x3b\xa4\x33\xeb\x2c\x47\x52\x6a\xf9\xfa\xa3\xd7\x5f\xec\x2f\xa0\x10\x96\xff\x67\xfd\x45\x76\x33\x98\x8b\xe5\xf6\xee\xe1\xdd\x33\x00\x00\xff\xff\x9c\x12\x46\xe4\xd9\x00\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xc1\x4e\xc3\x30\x10\x44\xef\x91\xf2\x0f\x5b\xf5\x52\x24\xc4\x07\xf4\xd6\x63\x0f\x20\x10\xfc\x80\x5b\xaf\x93\x95\xe2\x75\x58\xaf\xa1\x11\xe2\xdf\xd1\x3a\x25\xa4\xe2\x3a\x33\xcf\xde\x99\x2d\xbc\x17\x14\xc2\x0c\xb9\x4f\x65\xf0\xc0\x49\x21\x26\x4f\x61\x02\xed\x11\xfc\xe9\x01\xf0\x82\xe7\xa2\xe8\x81\x18\x46\x27\xe2\x86\x01\x87\xb6\xd1\x69\x44\x78\x76\x1d\x1e\x39\x24\xf8\x6a\x1b\x80\xd1\x75\x98\x0f\x41\x51\xf6\x70\x64\xdd\xc0\x16\x9e\x4a\x3c\xa1\x40\x0a\xb3\x09\xce\x5c\xd0\x9e\x32\x24\x46\x83\x52\x08\x19\x75\x01\xde\x7a\xbc\x4a\x06\xd5\xa0\x91\x10\x24\xc5\x7a\x52\x56\x27\x6a\xe0\x40\x91\x6e\xb9\xe8\x2e\xd6\xe3\x93\x61\x44\xa9\x98\xe5\x3a\xfa\x40\xbe\xc9\xf1\x72\x14\x29\xc6\x0c\x82\x5a\x84\xe7\x86\xf6\x85\x6d\x32\xb5\xcd\x77\xdb\xb4\x0d\xb1\xa2\x04\x77\xae\x5d\x89\x9d\xa2\x7f\x31\xfb\xaf\xb1\xf5\xdf\x2f\x4b\x5c\xb1\xba\xce\x2a\x48\x35\xf4\xaa\x42\xdc\x6d\x4c\x70\x3e\x12\xef\x4a\x21\xbf\xc8\x77\x7b\x38\x98\xba\xd8\x79\xb7\x1a\xe7\x7e\x55\xf8\x37\x69\x9f\xd6\xd7\xa2\x63\xd7\xa1\xfc\x7b\xef\x71\xd6\xeb\x51\x3f\x01\x00\x00\xff\xff\xb5\x81\xdc\x28\xec\x01\x00\x00")

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

var _type_admin_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x4d\x0a\xc2\x30\x10\x85\xf7\x39\xc5\xf3\x1a\xdd\xb9\x92\x82\x48\xc5\xa5\xb8\x18\xc8\x34\x0c\x34\x3f\x24\x13\xa4\x88\x77\x97\x06\xd1\x6a\x77\xc9\xc7\x37\x6f\xe6\xe9\x9c\x18\x7b\xeb\x25\xe0\x61\x80\x5a\xc5\x76\xb8\x68\x96\xe0\x76\x06\x60\x4f\x32\xad\xc1\x28\xb9\xe8\x89\x3c\xaf\xe1\x44\xff\xec\x69\xcc\x37\x79\x20\xc7\x10\x9f\x26\xf6\x1c\xb4\x60\x20\x27\x81\x94\xed\xb9\x72\x9e\xdb\x5e\xb6\x8e\x4b\x87\x6b\xf3\x6f\x06\x48\xe4\xb8\x0f\x63\xec\x16\xbb\xbd\x96\x4c\x09\xa9\x2a\x0e\xac\xcd\xeb\xdb\x6f\x7b\xf6\x47\x6c\xd6\x31\xba\x77\xb9\x4d\x99\x44\xa5\xdc\x63\xfe\x19\x7d\x05\x00\x00\xff\xff\x16\x90\x7b\x4e\x11\x01\x00\x00")

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

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x91\xcf\x4a\xc3\x40\x10\xc6\xef\x81\xbc\xc3\x94\x5e\x2a\xf8\x04\xb9\x89\x45\x08\xa8\x17\xcd\x03\xac\xc9\x6c\x1c\xd9\xcc\x86\xdd\x09\x22\xa5\xef\x2e\x9b\x3f\x35\x69\x36\x88\x37\x6f\xdd\x5f\x67\xe6\xfb\x7d\x84\x58\xd0\x69\x55\x22\x14\x1e\x1d\x9c\xd2\x04\xa0\xeb\xa8\xca\xe0\x45\x1c\x71\xbd\x0b\x00\x1b\x45\x66\x41\x34\x39\x2f\xcf\xaa\xc1\x05\x35\x2a\x02\x05\x0d\xb6\xef\x96\xd7\xa3\x8f\xb6\x26\x5e\xd0\x0f\xfb\xf6\x4a\x62\xe6\xa3\xe7\x34\x49\x13\xe2\xb6\x13\x78\x52\xac\x6a\x74\xfd\xda\x60\xba\x16\x6b\x95\xf7\x9f\xd6\x55\xf1\x0b\x77\x55\x35\x1e\xc9\xfb\xf7\xe9\x8f\x5d\xd6\x79\x11\xe3\xad\xca\x5b\x6a\xf2\xd5\xe2\xd4\x0d\xa8\x69\x0d\x36\xc8\xe2\xff\xf5\x07\xd9\x43\x6f\x7d\x44\x83\xb5\x12\x8c\x69\xef\x61\x12\x2f\x8a\xfc\xb8\x1b\xc0\xb5\x78\x60\x31\xf5\xc0\x23\xf2\x01\xc7\xf4\xa7\xf1\xb1\xc0\x51\x09\x0e\x30\xe2\x1f\x70\x69\x3b\xe7\xd1\x1f\xac\xd6\x1e\x25\x83\x9c\xe5\x16\x0c\x35\x34\xfd\xd6\x64\x04\x5d\x76\x29\x78\xdf\x2f\x3c\xf4\xf4\x26\x83\xe1\xe9\x2f\xc7\x58\x46\x72\xf8\x65\x31\x67\x19\x97\xd0\x09\x69\x2a\x95\x6c\x6b\x84\xa0\xd9\xd8\x3c\x6d\xbe\xbd\x8a\xfc\xf9\xf3\x3a\xf7\x9c\x26\xdf\x01\x00\x00\xff\xff\x1c\xc9\xcc\x3f\xf1\x03\x00\x00")

func type_users_graphql() ([]byte, error) {
	return bindata_read(
		_type_users_graphql,
		"type/users.graphql",
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
	"mutation.graphql": mutation_graphql,
	"query.graphql": query_graphql,
	"schema.graphql": schema_graphql,
	"type/admin.graphql": type_admin_graphql,
	"type/auth.graphql": type_auth_graphql,
	"type/users.graphql": type_users_graphql,
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
		"users.graphql": &_bintree_t{type_users_graphql, map[string]*_bintree_t{
		}},
	}},
}}
