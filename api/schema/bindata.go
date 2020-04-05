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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\x4f\x4e\x87\x30\x10\x85\xf7\x3d\xc5\x23\x6e\x74\xc3\x01\xd8\x61\xd8\x98\xc8\xce\x0b\x8c\x74\x84\xc6\x76\x8a\x32\x4d\x24\xc6\xbb\x1b\x10\xcb\x9f\xdf\xf2\x7d\x7d\xdf\x9b\xde\x21\x24\x25\x75\x51\x26\x74\x24\x08\xd1\xba\xb7\x19\x3a\x30\xec\x6b\x09\xfe\xe2\x2e\x29\x5b\x4c\xfc\x91\x58\xd4\x91\xf7\x73\x69\x74\x1e\x19\xed\x26\xe2\xdb\x00\x64\x83\x93\xe7\xd8\x3b\xb9\x77\x32\x26\xad\x50\x67\x52\x3c\x54\xa8\x93\x0e\x2f\xf1\x9d\xc5\x00\x81\x84\x7a\xfe\x3c\xb5\xdb\x03\xbb\xf6\xc9\xda\xed\x79\xdf\xfe\x27\x4f\x0b\x58\x84\x2d\x1b\xc0\xb2\x67\xe5\x8b\xd1\x1c\x61\x96\x1e\x63\xf4\x4c\x52\xfc\x1d\x59\x7f\x7c\x38\xb1\xe6\xdc\x5d\x53\x9e\x3f\x75\x9b\x1d\xdd\x4e\xff\x98\xdf\x00\x00\x00\xff\xff\x2d\x57\xe3\x8a\x64\x01\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xcb\x6e\xdc\x30\x0c\xbc\xeb\x2b\x66\x91\x4b\x0a\x04\xfd\x00\xdf\x9a\x02\x05\x72\x68\x9b\xa2\xf9\x01\xed\x9a\xb2\x59\x48\x94\x2b\xd1\x6d\x16\x45\xfe\xbd\x90\xec\xf5\x23\xbb\x37\x49\x9c\xe1\x0c\x39\xba\xc3\xef\x91\x12\x53\x46\xee\xe3\xe8\x5b\x48\x54\x84\xd8\xb2\x3b\x43\x7b\x42\x7b\xfc\x08\x7a\xa5\xd3\xa8\xd4\x82\x05\x83\x4d\xc9\x7a\x4f\xde\xe8\x79\x20\x3c\xdb\x8e\x9e\xc4\x45\xfc\x33\x80\x46\xb5\xbe\xc1\x93\xe8\x01\x77\xf8\x36\x86\x23\x25\x44\x87\xc1\x76\x94\x61\x9d\x52\x82\xf6\x9c\x11\x85\x0c\x10\x9d\xcb\xa4\x0b\xfe\xa5\xa7\xf9\xa9\x70\x2a\xae\x10\xe1\x52\x0c\xd5\x4a\x56\x9b\xd4\x00\x9e\x03\xef\x69\xc1\xbe\x16\xfb\x7f\x05\x03\xa5\xca\x32\x40\xc7\x7f\x48\x6e\xb8\x61\xa5\x90\xa7\xaa\x79\x33\x86\x65\x18\xb5\xce\x51\x67\xd8\x98\xda\x49\xad\xd0\xaf\x56\x6c\x47\xe9\x0b\xfb\x32\x4f\xe1\x88\x0d\xd4\xe0\xa7\x26\x96\xce\x00\x14\x2c\xfb\xcd\xfd\x57\x3c\xbe\xb0\xfa\x2d\x64\x1c\xb9\xdd\x5c\x95\x3c\x0d\x7d\x94\x15\xb2\xa8\x7d\x8e\x61\xb0\x72\xde\xa8\xbd\xe3\xee\xc4\x2b\x4d\x29\x39\x7b\xaa\xd9\xb0\x58\xa5\xf6\xc7\x48\xe9\x5c\xb9\xc3\x1c\x57\xb3\x04\xb7\x2a\x7d\x4f\x2d\xa5\xc7\x09\x67\xf3\x89\xa4\x65\xe9\x1a\x3c\xc6\xe8\xc9\x8a\x01\x1c\x93\x5f\x94\x0f\x85\x58\xbf\xc0\xda\x9d\x6b\xe7\x4b\x1d\xb0\x6d\x60\xb9\xdf\xfa\x3d\x7c\x68\xf0\xa9\xbc\x5e\xaa\xf9\xbe\x58\x9a\xec\x5c\x6a\xe5\x5c\xe8\x61\x5a\xf4\x55\x83\x39\x80\x15\xb1\x6d\xf2\x00\x57\x57\xd5\xec\x73\x7a\x40\x9c\xc6\x6b\x2e\x73\xae\x8d\x9e\xa7\xff\x72\x9a\x36\x7d\x25\x37\x27\xb0\x20\x98\x6e\xeb\xed\x92\xba\xa9\x37\x23\xaa\xde\x9b\xf9\x1f\x00\x00\xff\xff\x43\x83\x89\xca\x7a\x03\x00\x00")

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

var _type_admin_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\xc1\x6a\xc3\x30\x10\x44\xef\xfa\x8a\xed\x6f\xf8\x56\x28\x14\x43\x29\x2e\x3d\x96\x1e\x16\x34\x16\x0b\x92\x2c\xa4\x15\xc1\x84\xfc\x7b\xb0\x08\x89\xb1\x03\x49\x6e\xd2\xf0\x34\xa3\xa7\x73\x02\xbd\xdb\x20\x91\x8e\x86\xa8\x56\xb1\x1d\xfd\x6a\x96\xe8\xde\x0c\x11\x02\x8b\x5f\x07\xa3\xe4\xa2\xdf\x1c\xb0\x0e\x3d\x6f\xb3\x93\x31\xb7\xe6\x81\x1d\x48\x42\xf2\x08\x88\x5a\x68\x60\x27\x91\x15\xf6\xa7\x22\xcf\x6d\x17\xd6\xa1\x74\xf4\xd7\xf8\x7f\x43\x94\xd8\xa1\x8f\xe3\xd4\x2d\x74\x3b\x2d\x9d\x12\x53\x55\xfa\x84\x36\xae\x6f\xb7\xfd\xb7\xaf\x60\xa3\xbe\x26\x77\x91\xdb\xc9\x24\x2e\xe5\x30\xe5\xfb\x4f\xed\x66\xe3\x49\xf3\xd7\x66\x3e\xe0\xa1\x78\x60\x73\x0e\x00\x00\xff\xff\x7f\xeb\x7e\xa5\xa4\x01\x00\x00")

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

var _type_company_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\x41\xca\xc2\x30\x14\x84\xf7\xef\x14\xd3\xdd\xff\x43\x4f\x90\x9d\x0a\x82\x0b\x51\x71\x29\x2e\x82\x99\x86\x40\x93\x86\x34\x5d\x14\xe9\xdd\x25\x46\xa5\xbb\x79\xf3\x66\x86\x2f\xcf\x91\xd8\x0d\x3e\xea\x30\xe3\x29\xc0\x34\x39\xa3\x70\xcd\xc9\x05\xdb\x08\xf0\x48\xd4\x99\x66\x93\xbf\xa6\x00\x41\x7b\xae\x33\x5e\x07\x6d\x99\xc6\xbf\xa8\x2d\x15\xce\xda\xb2\x45\xe7\xfa\xcc\xa4\x70\xac\xcf\xfd\xfb\x6c\x31\x24\xc3\xb4\x9d\x15\x4e\x55\xfc\xff\x12\xa5\xd6\xc8\x22\xb2\x66\x2a\x26\x9c\x8f\x3d\x3d\x43\x1e\xcb\xb6\x0b\x05\xe8\x32\x31\x55\x62\x1a\xcb\x51\xe1\xf6\x69\xdc\x05\x28\x1c\x87\xd0\x0d\x95\xa5\x28\x59\xe4\x15\x00\x00\xff\xff\x00\x22\xec\x14\xeb\x00\x00\x00")

func type_company_graphql() ([]byte, error) {
	return bindata_read(
		_type_company_graphql,
		"type/company.graphql",
	)
}

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x92\xc1\x6a\xe3\x30\x10\x86\xef\x7a\x8a\x09\xb9\x64\x61\x9f\xc0\xb7\x90\xb0\x60\xd8\x5d\x52\xda\x9c\x4a\x0f\xaa\x3d\x52\xa7\xc8\x92\x91\xc6\x94\x50\xf2\xee\x45\x8a\x9d\xaa\x91\xd3\x42\x6f\xbd\x49\xbf\xe7\xd7\xfc\xf3\x8d\xc9\x32\x7a\x25\x1b\x84\x7d\x40\x0f\xaf\x02\x60\x18\xa8\xad\xe0\x96\x3d\x59\xbd\x10\x00\x8d\x47\xc9\xd8\xae\x79\x12\x05\x00\x76\x92\x4c\x5e\xa4\xc8\x07\xfe\x2f\x3b\xcc\x45\x23\x4b\x8d\xd1\x60\xff\xe4\x6c\x51\xf8\xd7\x69\xb2\xb9\xf8\xec\x1e\xef\x88\x4d\x56\x78\x14\x82\x6c\x3f\x30\xfc\x93\x56\x6a\xf4\xc9\x92\x32\x17\x79\x7a\x19\xc2\x8b\xf3\xed\x8c\x79\xdd\xb6\xa3\xbf\x4e\xf7\xe8\x6f\x5c\xd7\x4b\x7b\xd8\xef\xeb\xed\xb7\xa6\x2a\xfa\x97\xe1\xaf\x8c\xfe\x49\xd0\x2d\x1a\x64\x2c\xb2\x7e\xdc\xcf\x51\x08\x3e\xf4\x38\x21\x01\xea\x7a\x83\x1d\x5a\x0e\x3f\x61\xa3\x67\xf2\x15\x6c\x4e\x87\x62\xa2\x9d\xd4\x98\x4f\xb5\x93\x9a\x6c\x8c\x7f\x33\xa0\x3f\x9c\xb6\xdf\x6a\x0c\x15\xdc\x8f\x8e\x87\x84\x55\x63\x6d\x95\xab\x62\x7d\x3a\xc5\x77\x97\x90\x5e\x8e\x60\xb5\x64\x9c\x81\xb5\x84\x09\x57\xfc\x17\x16\xe9\x7e\x81\x26\x4a\x33\x70\xa2\x5c\xe2\x89\xea\x0c\xa0\xa9\x78\x44\xb4\x95\x8c\x49\x2b\x09\x45\xb5\x71\x83\x0f\x18\x56\x4e\xa9\x80\x5c\x41\x6d\xf9\x37\x18\xea\x68\x3a\x2b\x32\x8c\xbe\x3a\x0f\xb6\x49\x86\x3f\x49\xfd\x15\xd1\x26\xff\xf4\x96\xe5\x51\x58\x7d\xe1\xab\x2d\x9f\x3c\xe8\x99\x14\x35\x92\xaf\x87\x88\x6d\xb2\xb2\xac\x57\x6e\x2e\x1a\xbe\x7f\xbc\xe8\x7a\x14\x6f\x01\x00\x00\xff\xff\x2d\xec\x31\x46\x9b\x04\x00\x00")

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
	"type/company.graphql": type_company_graphql,
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
		"company.graphql": &_bintree_t{type_company_graphql, map[string]*_bintree_t{
		}},
		"users.graphql": &_bintree_t{type_users_graphql, map[string]*_bintree_t{
		}},
	}},
}}
