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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x92\x4d\x8e\xe2\x30\x10\x85\xf7\x3e\x45\x45\xb3\x99\xd9\x70\x80\xec\x18\xd0\x48\x48\x20\x8d\x80\x3e\x40\xb5\x5d\x24\x56\x3b\xe5\x74\x52\x91\x3a\x6a\x71\xf7\x56\x8c\x93\x38\xfc\x2c\xeb\xf9\xbd\x8f\xaa\x47\x7e\x41\xd5\x09\x8a\xf5\xdc\x82\x46\x86\xca\x1b\x7b\xe9\x41\x4a\x02\xf3\xbe\x02\xfa\x22\xdd\x09\x19\x68\xe9\xb3\x23\x16\x8b\xce\xf5\x2b\xa5\x2c\xd7\x9d\xc0\x5b\xed\x3c\x9a\x7f\xd6\xd1\x81\x04\xe1\x5b\x01\x5c\xac\xa3\x73\x5f\x53\x0e\x27\x69\x2c\x17\x99\x02\xd0\x9e\x85\x58\xf6\xc4\x85\x94\x39\xec\x58\x32\x75\x55\x4a\xfa\x9a\x12\xc4\x91\xda\x3a\x20\xba\xc6\xa5\xe9\xb6\xd3\x9a\xda\xf6\xec\x3f\x88\x67\xfd\xfa\xb8\xc3\xe9\x66\x0c\x8c\x97\xa1\xf0\xa3\x87\x78\x72\xb0\xa2\xa9\x2c\xef\x7d\x61\xf9\x77\x20\xe6\xb0\x9e\x94\xec\x4f\x0e\xeb\x4e\xca\xc0\x51\x00\x15\x32\x16\xd4\x2c\xdc\x87\x44\xbb\xf7\xa3\x31\xf1\x79\x66\x8f\xca\x6e\x10\x86\x40\x9c\x67\xfc\xff\xc6\x0f\x35\xde\x2e\x3b\x0e\xc5\xb7\x32\xe6\x97\x95\x0f\xf1\x65\x83\x2f\x28\xb1\x9b\x47\x4a\x7c\x58\xee\x61\xc8\x91\xd0\xdd\xe6\xdb\x54\x9c\x96\xff\xeb\xbd\x23\xe4\xec\x76\x6c\x68\x2e\x39\x35\xcc\x93\x37\x4c\x13\x7e\xe1\xdd\xce\xd2\x33\xb4\x6e\x08\x85\x36\xbe\xaa\x91\xfb\x31\xb3\x49\xc5\x29\x15\xe7\xfb\xd0\x58\xa3\x02\x18\x3e\xc8\xa0\x3d\x45\x04\x43\xac\xf0\xf1\xef\x0a\xaf\x0d\x69\xac\x45\x97\x98\x7e\xa6\xe9\xc2\x57\xf5\x13\x00\x00\xff\xff\x23\x2b\x4d\x03\x57\x03\x00\x00")

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

var _type_company_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x50\xbb\x6a\x03\x31\x10\xec\xf5\x15\xe3\x2e\x01\x37\x49\xa9\xce\x3e\x08\x18\xf2\x24\x65\x48\xb1\x44\x6b\x21\xb8\x5b\x09\x49\x17\x38\x4c\xfe\x3d\x48\x3a\x07\x25\xe9\xdd\x8d\x46\xb3\xbb\x33\x93\x97\xc0\x18\xfc\x14\x48\x16\x9c\x14\x40\x21\x44\xff\xc9\x46\x63\xef\xfd\xc8\x24\x0a\x98\x67\x67\x34\x5e\x73\x74\x62\x37\x0a\xf8\x88\x4c\x99\xcd\x2e\x9f\x49\x05\x08\x4d\xdc\x6b\x26\x12\xb2\x1c\xd3\x55\x20\xcb\x1a\xcf\x64\x79\x8b\xa3\x1b\x33\x47\x8d\x87\xf6\x79\x57\x9f\x5b\xf8\x68\x38\xee\x17\x8d\xa7\x06\xae\x7f\x14\x65\xac\x6c\x23\x63\x22\xa7\xa4\xb1\x6b\x60\xa3\xbe\x94\xaa\xde\x57\xa2\x79\x6f\xf8\xde\x09\xdf\xf4\x5e\x3a\xfe\xf6\x57\x0e\x3f\x4b\x5e\x7a\x26\xf8\x94\x07\x6f\xf8\x9f\x2a\x76\xb2\xf3\xe9\xb5\xb6\xe2\x11\x6e\x0a\x23\x4f\x2c\x39\x95\xa8\x4e\x4a\x3f\x2f\x33\xc7\x56\x2a\x1b\xcb\x49\xe3\x6d\x9d\x78\x2f\x97\xc8\xf2\x41\x8e\xbe\x55\x53\x50\xd9\xeb\x24\xcc\x19\x43\xed\x77\x15\x1f\x2a\x75\xaa\x46\x2a\xf1\xf8\xa7\xe9\x8b\xa5\xfe\x0e\x00\x00\xff\xff\x87\xf2\x57\xad\x2c\x02\x00\x00")

func type_company_graphql() ([]byte, error) {
	return bindata_read(
		_type_company_graphql,
		"type/company.graphql",
	)
}

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x92\x4f\x8f\xd3\x30\x10\xc5\xef\xf9\x14\xb3\xea\x65\x91\xf8\x04\xb9\xad\xb6\x5a\x29\x12\xa0\x45\xd0\x13\xe2\x60\x92\x67\xef\x20\xff\x89\xec\x89\x56\x15\xea\x77\x47\x76\x92\x12\xea\x16\xc4\x91\x9b\xf3\x32\x6f\x66\xde\xcf\x66\x2f\x88\x5a\xf5\xa0\x43\x42\xa4\x1f\x0d\xd1\x34\xf1\xd0\xd2\x27\x89\xec\xcd\x5d\x43\xd4\x47\x28\xc1\xf0\x20\xab\xd8\x10\xc1\x29\xb6\xdb\x22\xcd\x31\xc9\x07\xe5\xb0\x15\xad\xaa\x35\x81\xc5\xf8\x12\x7c\x55\xf8\x2e\x18\xf6\x5b\xf1\x7b\xf8\xf6\x99\xc5\x6e\x0a\x4f\x4d\xc3\x7e\x9c\x84\xde\x2b\xaf\x0c\x62\xb1\x94\x9d\xab\x7d\x46\x95\xd2\x6b\x88\xc3\x15\xf3\xc3\x30\x2c\xfe\xae\x7c\x67\x7f\x1f\xdc\xa8\xfc\xf1\x70\xe8\xf6\xab\x83\x76\xf4\x14\x22\xa9\xc1\xb1\x4f\xf4\xfa\x02\x3f\x93\xc8\xbf\xdc\xec\x4f\x14\x34\x0d\xac\x35\x22\xbc\x2c\x4d\x18\xe9\x5f\x78\x54\x9b\xd7\xb1\x6f\x40\xfb\x43\xc4\x3d\x2c\x04\x55\xca\xdf\x6f\xf6\xd4\x34\x72\x1c\xb1\xc2\x24\x76\xa3\x85\x83\x97\xf4\x3f\xbc\x85\xf3\x9d\xb5\xf4\x38\x1f\x0a\x93\x18\x34\x5b\x74\x4e\x19\x1c\xe2\x79\xab\xcb\xac\xcf\xca\x60\x9b\xf7\x59\x19\xf6\x39\xd8\xc7\x09\xf1\x38\xbf\xa8\xc1\x20\xb5\xf4\x65\x71\x7c\x2d\xc0\x0d\x3a\xaf\x43\x9b\xeb\xcb\x29\xf7\xdd\x51\xe9\x9c\x91\x1b\x25\xb8\x82\x71\x47\x2b\xc8\xfc\xbe\xee\xca\xf7\x05\xb4\x2c\x5d\xc1\x96\xe5\x1a\x5c\x56\xaf\xa0\x5b\x8b\x17\x78\x7b\x25\x28\x5a\xcd\x2e\xab\x7d\x98\x62\x42\xba\x0f\x5a\x27\x48\x4b\x9d\x97\xb7\x64\xd9\xf1\x7a\xd6\x6c\x05\xb1\x3d\x07\x7b\x2c\x86\xa7\xa2\xbe\xc9\xd0\x8b\x7f\xed\xe5\x65\x11\xee\xff\xe2\xeb\xbc\xcc\x1e\x44\x61\xcd\xbd\x92\xdb\x4b\xe4\x31\x9b\xb2\xcd\xac\xad\xb9\x1a\xf8\xeb\xe7\xc5\xd4\x53\xf3\x33\x00\x00\xff\xff\xdd\xd6\xb4\x69\xef\x04\x00\x00")

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
