// Code generated by go-bindata.
// sources:
// mutation.graphql
// query.graphql
// schema.graphql
// type/admin.graphql
// type/auth.graphql
// type/company.graphql
// type/course.graphql
// type/users.graphql
// DO NOT EDIT!

package schema

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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\x41\x6e\xdb\x30\x10\xbc\xf3\x15\x6b\xf4\xd2\x5e\xf2\x00\xdd\x52\x1b\x05\x0c\xc4\x68\x11\xc7\x0f\xd8\x4a\x1b\x9b\x28\xb5\x64\xc9\x65\x50\xa1\xf0\xdf\x03\xd1\x34\x2d\x4a\xf6\x71\x87\x33\xc3\xd9\x21\xbf\x40\x1f\x05\x45\x5b\x0e\xd0\x22\x43\x6f\x3b\xfd\x3e\x80\x9c\x08\xba\xdf\x4f\x40\xff\xa8\x8d\x42\x1d\x04\xfa\x1b\x89\x45\xa3\x31\xc3\x93\x52\x9a\x5d\x14\x38\x38\x63\xb1\xfb\xa1\x0d\xed\x48\x10\xfe\x2b\x80\x77\x6d\xe8\x6d\x70\xd4\xc0\x5e\xbc\xe6\xe3\x4a\x01\xb4\x96\x85\x58\x5e\x88\x8f\x72\x6a\x60\xcb\xb2\x52\x67\xa5\x64\x70\x34\xb1\x78\xa5\xe0\x92\x45\xf4\x66\xaa\x0e\xb1\x6d\x29\x84\x37\xfb\x87\xf8\x86\x9f\x97\x19\xf6\x17\x62\xf2\x78\x28\x4a\x97\xee\xf2\xca\x89\x8a\x5d\xaf\xf9\xc5\x1e\x35\x7f\x4d\x8e\x0d\x3c\x17\x64\xf5\xad\x81\xe7\x28\xa7\xe4\xa3\x00\x7a\x64\x3c\x92\xaf\xd8\xbb\x09\x36\xe7\xb7\x9e\x50\x28\x33\xae\x82\xf5\x14\xdc\x8e\xd8\x28\xcb\xf3\xed\x92\x5f\xde\x8e\x65\x5e\xf6\x7b\x1d\xeb\x0f\x72\xb5\xa8\x8b\x1f\xe5\x75\x8f\x0f\x5c\x72\x43\x4b\x97\x7c\x50\xe7\xe8\xc8\xd0\x22\xfc\x66\x0a\x96\xf0\xdf\xad\x35\x84\xbc\x2a\x2b\xa7\x0a\xeb\x85\x13\x54\x14\x69\x1a\x1f\xdb\x75\x73\xfa\xe1\x06\x2d\xe8\x97\x4c\x15\x7d\x73\x83\x1e\xe7\x59\xdb\xde\x21\x0f\x75\xa2\x0c\x16\x55\x9e\xe7\xa2\x6b\xf7\x0a\x60\xfc\xcb\x09\xbb\x6b\x91\x08\xb9\xf7\xbb\xcf\x9c\x08\x9e\x5a\x74\xd2\x9e\x70\xfa\xc9\xeb\xcc\xe8\x9c\xb7\x1f\x25\x74\x8c\xba\x6b\xe0\x70\xd8\x6e\xea\x94\x01\x3f\xe8\x27\x1b\xcd\xb4\xb6\xd1\x07\xba\x6e\xb7\x9f\xe1\x65\xc1\x29\x98\xf5\x6b\x83\x21\x78\x6b\xfb\xa5\xc5\xec\xa8\xb8\x5c\x52\xab\xb3\xfa\x0c\x00\x00\xff\xff\xda\xa8\x68\x70\x3c\x04\x00\x00")

func mutationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_mutationGraphql,
		"mutation.graphql",
	)
}

func mutationGraphql() (*asset, error) {
	bytes, err := mutationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mutation.graphql", size: 1084, mode: os.FileMode(420), modTime: time.Unix(1587076210, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\xc1\x6e\xdb\x3a\x10\xbc\xf3\x2b\xc6\xc8\x25\x0f\x08\xde\x07\xe8\xd6\xb4\x28\xe0\x43\xdb\x14\x4d\x3e\x80\x36\x97\xd2\x16\xe4\x52\xa5\x56\x6d\x8c\x22\xff\x5e\x90\x92\x2c\x29\xf0\xcd\xe6\xce\xec\xcc\xce\xe8\x0e\xbf\x46\xca\x4c\x03\x86\x2e\x8d\xc1\x41\x92\x22\x26\xc7\xfe\x02\xed\x08\xee\xf4\x3f\xe8\x95\xce\xa3\x92\x03\x0b\x7a\x9b\xb3\x0d\x81\x82\xd1\x4b\x4f\x78\xb2\x2d\x1d\xc5\x27\xfc\x35\x80\x26\xb5\xa1\xc1\x51\xf4\x80\x3b\x7c\x1d\xe3\x89\x32\x92\x47\x6f\x5b\x1a\x60\xbd\x52\x86\x76\x3c\x20\x09\x19\x20\x79\x3f\x90\x5e\xf1\xcf\x1d\xcd\x4f\x85\x53\x71\x85\x08\x9f\x53\xac\x56\x06\xb5\x59\x0d\x10\x38\xf2\x9e\x16\xed\x6b\xb1\xff\x47\xd0\x53\xae\x2c\x03\xb4\xfc\x9b\xe4\x86\x1b\x56\x8a\xc3\x34\x35\x6f\xc6\xb0\xf4\xa3\xd6\x3b\xea\x0d\x1b\x53\x3b\xa9\x15\xfa\xc5\x8a\x6d\x29\x7f\xe6\x50\xee\x29\x1c\xb1\x91\x1a\xfc\xd0\xcc\xd2\x1a\x80\xa2\xe5\xb0\xf9\xff\x33\x9d\x9e\x59\xc3\x16\x32\x8e\xec\x1a\xbc\xbc\x1c\x3f\x95\xd8\x28\x50\xdf\x25\x59\x01\x57\xad\x8f\x29\xf6\x56\x2e\x1b\xad\x1d\x73\x27\x5c\x49\x4a\xd9\xdb\x73\xed\x85\xc5\x2a\xb9\xef\x23\xe5\x4b\x65\xf6\x73\x55\xcd\xb5\xb4\x55\xe7\x5b\x76\x94\x1f\x27\x9c\x1d\xce\x24\x8e\xa5\x6d\xf0\x98\x52\x20\x2b\x06\xf0\x4c\xc1\x2d\x52\x87\x42\xac\xf5\xaf\xdb\xb9\x6e\x5e\xe6\x80\x75\x91\xe5\x7e\x75\x7b\xf8\xaf\xc1\x87\xf2\xb6\xcc\x86\xfb\x62\x68\x32\xb3\xcc\xca\xef\x42\x8e\x53\xc4\xef\xe8\x73\xf0\xeb\x7c\xbb\xe2\x01\xbe\x86\xd4\xec\xfb\x79\x40\x9a\x4e\x6b\x96\x1b\xd7\x45\x4f\xd3\x77\x72\x9e\x32\x7e\x27\x36\x27\x7f\x9d\x33\xdd\x56\xdb\x35\x74\x53\x6d\x46\x54\xb5\x37\xf3\x2f\x00\x00\xff\xff\xca\xac\x30\x89\x70\x03\x00\x00")

func queryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_queryGraphql,
		"query.graphql",
	)
}

func queryGraphql() (*asset, error) {
	bytes, err := queryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.graphql", size: 880, mode: os.FileMode(420), modTime: time.Unix(1587076223, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\x4e\xcc\x49\x2c\x52\x08\x0d\xf5\x74\xe1\xe2\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\x0a\x0a\xb9\xa5\x25\x89\x25\x99\xf9\x79\x56\x0a\xbe\x50\x16\x57\x2d\x17\x20\x00\x00\xff\xff\x7d\xb7\x88\x41\x3c\x00\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 60, mode: os.FileMode(420), modTime: time.Unix(1587045855, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeAdminGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xcd\x8a\x83\x30\x14\x85\xf7\x79\x8a\x3b\xaf\x91\xdd\x30\x6e\x84\x61\x70\x28\xae\x4a\x17\x17\x72\x0d\x17\x92\x18\xf2\x43\x11\xf1\xdd\x8b\xb1\x60\xd5\x52\xda\x5d\x38\x9c\x7c\xdf\x49\xd2\xe0\x09\xbe\x95\x65\x07\xa3\x00\xc8\x99\x95\x84\xb6\xad\xab\x2f\x01\x40\x16\xd9\x48\x38\xa5\xc0\x4e\xcf\x41\xc7\x21\xa6\x3f\xb4\xf4\x18\x1a\xdc\x67\x93\x10\x2b\xb7\x41\x4d\xc0\xd6\x1b\xb2\xe4\x52\x84\x06\x35\x3b\x4c\xa4\xfe\x33\x85\xa1\x58\x49\x69\x8a\x12\xce\xa5\x7f\x11\x00\x1e\x35\xd5\xae\xeb\xe5\xdc\x2e\xa7\x99\xc9\xce\xe7\xb4\x40\x7f\x7b\x7d\x5f\x7c\xd8\xe8\x31\xc6\x6b\x1f\xd4\x66\xce\x72\xf5\x27\x10\x26\x2a\x80\xba\x04\xe3\xfb\x6f\xfa\xcc\xd4\x7a\x75\x34\x6d\x3f\xf7\xe0\x7d\xa2\xdd\x5b\x57\x41\x45\x86\x5e\x0a\x26\x71\x0b\x00\x00\xff\xff\x68\xfa\x3e\xad\xdb\x01\x00\x00")

func typeAdminGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeAdminGraphql,
		"type/admin.graphql",
	)
}

func typeAdminGraphql() (*asset, error) {
	bytes, err := typeAdminGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/admin.graphql", size: 475, mode: os.FileMode(420), modTime: time.Unix(1587069388, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeAuthGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x2c\x2d\xc9\x08\xc9\xcf\x4e\xcd\x53\xa8\xe6\x52\x50\x28\x01\xb1\xac\x14\x82\x4b\x8a\x32\xf3\xd2\x15\xb9\x6a\xb9\x00\x01\x00\x00\xff\xff\xa3\x33\x16\xdc\x24\x00\x00\x00")

func typeAuthGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeAuthGraphql,
		"type/auth.graphql",
	)
}

func typeAuthGraphql() (*asset, error) {
	bytes, err := typeAuthGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/auth.graphql", size: 36, mode: os.FileMode(420), modTime: time.Unix(1586436668, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeCompanyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x50\x3d\x6b\x03\x31\x0c\xdd\xfd\x2b\x5e\xb6\x16\xb2\xb4\xa3\xb7\xe4\x4a\xe1\xa0\x9f\x94\x4c\xa5\x83\x88\x15\x63\xb8\x93\x8d\xed\x2b\x1c\xa1\xff\xbd\xd8\xbe\x94\xa3\xdd\xbb\x3d\x3d\x3d\x49\x4f\x2f\xcf\x81\xd1\xf9\x31\x90\xcc\x38\x2b\x80\x42\x88\xfe\x93\x8d\xc6\xde\xfb\x81\x49\x14\x30\x4d\xce\x68\x1c\x0e\xfd\xdd\x46\x01\xc7\xc8\x94\xd9\xec\xb2\xc6\x5b\x8e\x4e\xac\x02\x84\x46\xbe\x94\x45\x33\x92\x90\xe5\x98\xae\x02\x59\xd6\x78\x21\xcb\x5b\x9c\xdc\x90\x39\x6a\x3c\xb6\xe6\x7d\x2d\xb7\xf0\xd1\x70\xdc\xcf\x1a\xcf\x0d\x5c\xff\x28\xca\x58\xd9\x46\xc6\x44\x4e\x49\x63\xd7\xc0\x46\x7d\x29\x55\x9d\x2f\x44\x73\xde\xf0\x83\x13\xbe\x59\x7b\x59\xf1\xb7\x6b\xfe\xe8\x27\xc9\xf3\x9a\x09\x3e\xe5\xce\x1b\xfe\xa3\x8a\x2b\xd9\xe5\xf4\x12\x5a\xf1\x08\x37\x86\x81\x47\x96\x9c\xca\xab\x4e\x4a\x3e\xaf\x13\xc7\x16\x29\x1b\xcb\x49\xe3\x7d\x99\xf8\x28\x97\xc8\x72\x2f\x27\xdf\xa2\x29\xa8\xec\x75\x12\xa6\x8c\xae\xe6\xbb\x88\xfb\x4a\x9d\xab\x91\x4a\x3c\xfd\x4a\xfa\xdf\xbe\xfe\x0e\x00\x00\xff\xff\x10\xb8\xf9\xd1\x2a\x02\x00\x00")

func typeCompanyGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeCompanyGraphql,
		"type/company.graphql",
	)
}

func typeCompanyGraphql() (*asset, error) {
	bytes, err := typeCompanyGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/company.graphql", size: 554, mode: os.FileMode(420), modTime: time.Unix(1587076231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeCourseGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x54\x4d\x6b\x1b\x31\x10\xbd\xeb\x57\x3c\xdb\x97\x16\xe2\x43\x62\x3b\xb4\x7b\x4b\x5d\x17\x0c\x0d\x09\xb1\x0d\x05\xe3\x83\x2a\xcf\x3a\x22\xbb\xd2\xa2\x8f\x12\x13\xf2\xdf\x8b\xa4\xf5\xee\xda\x4e\x2f\xa1\x97\x16\x2d\x3b\xfa\x78\xef\xcd\x68\xa4\x11\x29\x5f\xe2\x46\x08\xb2\x76\xb9\xaf\x08\x2f\x0c\x30\x64\x9d\x91\xc2\xd1\x96\x01\xba\x22\xc5\x5e\x19\x8b\xc0\x85\x33\x5e\x38\x6f\x68\x56\x50\x49\xca\x45\x78\xa9\xb7\xbe\x20\x06\x38\xb2\x8e\x01\x05\x59\xab\x5b\xce\x6d\x5c\xee\x12\xce\x70\x2e\x78\x9e\x6a\x6f\x2c\xcd\x55\xae\x23\x48\xf1\x92\xb2\xe0\x50\xaa\x5d\x83\xb9\x53\x85\x54\x94\x90\x11\xe5\xbd\xdc\x66\x58\xad\xe6\x5f\x7b\x0c\x90\x2a\xd7\x59\x47\xa8\xd7\x8a\x17\xdc\x5a\xa3\x75\xf9\x26\xf7\x0d\x6a\x60\x4a\x55\x79\x87\x05\xff\x45\x27\xf4\x79\x5c\x38\xd3\x38\x0a\x19\x10\xdc\xd1\x4e\x9b\x7d\x58\x6d\x30\xf4\x2c\xc8\x54\xae\x03\x93\xca\x19\xbd\xf5\xc2\x49\xad\x3a\xd3\x3f\xb9\x78\xda\x19\xed\xd5\x76\xfa\x48\xe2\x29\xc3\x17\xad\x0b\xe2\x8a\x01\xbc\x39\xae\xac\x73\x74\x0c\xa8\x8c\x14\x94\xe1\x5b\xa1\x79\x48\xb0\xd0\x85\x36\x1d\x49\xc7\x77\x36\xc3\x3a\x8d\x37\x0c\xb0\x15\x09\x99\x4b\xb1\x24\x53\xda\x23\xdf\x4a\x91\x99\x97\x7c\x47\x0b\x1f\x1d\x74\x55\xbc\xd3\xa6\xdd\x13\x06\xc9\xe8\x1c\xee\x91\xb0\x0c\xab\xf8\x20\x0e\x09\x83\x48\x09\xd7\xaa\xd8\x7f\x0c\x97\x85\x3f\xdf\x73\xe3\xa4\x90\x15\x57\xce\x66\x98\x2b\x77\x9c\xeb\xee\x19\xff\x0f\x89\xee\xbd\x3b\xd3\xf6\x50\x6d\x19\xd6\x75\x3e\x1c\x95\xbd\x0d\x06\xa1\x10\xf6\xc8\xb5\x09\x79\x95\x8a\xea\x2c\xdb\x36\x93\x2d\x3e\x95\x5c\xdc\xc3\x69\xf9\xf6\xce\x2b\xc8\x51\x88\x6f\x9d\x8a\x36\xb9\xeb\xb5\xaa\xed\x74\x47\xf5\xa8\xc0\x4f\x25\x5f\x19\x1b\x60\x8d\xd9\x8f\x9b\xdb\xfb\xef\x33\x4c\xef\x56\x0f\x8b\x19\x16\xcb\x87\xd5\x74\xb9\x7a\x98\xb1\x01\x80\x97\xa4\xd3\x4f\x0f\x49\xff\xa2\x16\xe8\x8f\xc6\xa3\xf1\x68\xd8\xfd\x8f\xa3\xe9\xbf\x5e\x24\x1e\x4e\x88\x71\xf6\xe0\xbf\x7f\x79\x35\x0a\xdf\xb0\xb6\x57\x75\x27\x9a\x51\x03\x3e\xec\xb8\x1e\xb6\xd1\xa4\xe7\xa9\x8d\xe6\xf2\x2a\xb6\x61\x32\x41\xad\x19\x1d\x02\xea\xd2\xc3\x23\xd7\x92\x3f\x7d\x4e\x6d\x38\xb9\x8e\x6d\x38\xb9\x6e\x3a\x93\xeb\x49\x2b\xb0\x89\xf6\xaf\x6c\x10\x03\xc8\x1c\xd2\xc2\x51\x59\x15\xdc\x11\x4a\xfe\x44\xe0\x10\xba\xda\xff\x0b\xbb\xff\x43\x28\xa3\x78\x09\xc6\xe9\x4a\x84\xce\x64\x32\x8e\x6d\x18\xff\xfd\x13\xfa\x3b\x42\xd9\xb0\xdf\x01\x00\x00\xff\xff\x70\xfa\x76\x15\x17\x07\x00\x00")

func typeCourseGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeCourseGraphql,
		"type/course.graphql",
	)
}

func typeCourseGraphql() (*asset, error) {
	bytes, err := typeCourseGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/course.graphql", size: 1815, mode: os.FileMode(420), modTime: time.Unix(1587062092, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeUsersGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x92\xc1\x8e\xd4\x30\x0c\x86\xef\x7d\x0a\xa3\xb9\x2c\x12\x4f\xd0\x1b\x9a\xd1\x4a\x95\x00\x2d\x82\x9e\x10\x07\xd3\x3a\x59\xa3\xc4\xa9\x12\x57\xab\x11\x9a\x77\x47\x49\xdb\xa1\x4c\x3b\x20\x6e\xdc\xd2\xbf\xfe\x6d\xff\x5f\xc2\xa2\x14\x0d\x76\x04\x6d\xa2\x08\x3f\x2a\x80\x71\xe4\xbe\x86\xb6\x6d\x4e\xaf\x2a\x80\x2e\x12\x2a\xf5\x6f\xb5\x86\x4f\x1a\x59\x6c\x05\x40\x1e\xd9\x2d\xdf\xb9\xc8\x70\x4c\xfa\x01\x3d\xad\x45\x87\x5b\x4d\xc9\xd1\xf0\x1c\x64\x53\xf8\x2e\x58\x96\xb5\xf8\x3d\x7c\xfb\xcc\xea\x56\x85\x97\xaa\x62\x19\x46\x85\xf7\x28\x68\x29\x16\x4b\xd9\x78\xb3\xcf\x80\x29\xbd\x84\xd8\xef\x98\x8f\x25\xcf\xdc\xa2\x29\x52\x6e\xd1\x05\x3f\xa0\x9c\x73\xea\x29\x3b\x1c\xe0\x31\x44\xc0\xde\xb3\x24\x78\x79\x26\x99\x50\xb0\x58\xf0\x93\x3b\x41\x30\xd0\xb3\x31\x14\x49\x74\x6e\xc1\x94\xfe\x05\xc8\x66\xf5\x6d\xee\x3b\xd4\xfe\x90\xf1\x44\x8e\x76\x32\xae\x2f\xf6\x52\x55\x7a\x1e\x68\x61\x09\xec\x07\x47\x9e\x44\xd3\xff\xff\x10\xae\xb7\x55\xc3\x71\x3a\x14\x1e\x31\x18\x76\xd4\x78\xb4\xd4\xc6\xeb\x56\xb7\x49\x9f\xd0\xd2\x3a\xed\x13\x5a\x96\x1c\xec\xe3\x48\xf1\x3c\x3d\xa7\xde\x52\xaa\xe1\xcb\xec\xf8\x5a\x60\x5b\x6a\xc4\x84\x3a\xd7\x97\x53\xee\x7b\x80\xd2\x39\xe3\xb6\xa8\xb4\x03\xf1\x00\xbf\x63\xcc\xdf\x37\xd0\xb2\xb4\x83\x2d\xcb\x5b\x70\x59\xdd\x41\xb7\x14\xcf\xf0\x4e\xa8\x54\xb4\x2d\xbb\xac\x76\x61\x8c\x89\xd2\x43\x30\x26\x91\xd6\xd0\x88\xbe\x01\xc7\x9e\x97\xb3\x61\xa7\x14\xeb\x6b\xb0\x63\x31\x3c\x16\xf5\x75\x86\x5e\xfc\x4b\x2f\xd1\x59\x78\xf8\x8b\xaf\x11\x9d\x3c\x14\x95\x0d\x77\xa8\xf7\x97\xc8\x63\x56\x65\xab\x59\x6b\xf3\x66\xe0\xaf\x9f\x37\x53\x2f\xd5\xcf\x00\x00\x00\xff\xff\x7e\xaf\xba\x37\xea\x04\x00\x00")

func typeUsersGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeUsersGraphql,
		"type/users.graphql",
	)
}

func typeUsersGraphql() (*asset, error) {
	bytes, err := typeUsersGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/users.graphql", size: 1258, mode: os.FileMode(420), modTime: time.Unix(1587071844, 0)}
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
	"mutation.graphql": mutationGraphql,
	"query.graphql": queryGraphql,
	"schema.graphql": schemaGraphql,
	"type/admin.graphql": typeAdminGraphql,
	"type/auth.graphql": typeAuthGraphql,
	"type/company.graphql": typeCompanyGraphql,
	"type/course.graphql": typeCourseGraphql,
	"type/users.graphql": typeUsersGraphql,
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
	"mutation.graphql": &bintree{mutationGraphql, map[string]*bintree{}},
	"query.graphql": &bintree{queryGraphql, map[string]*bintree{}},
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
	"type": &bintree{nil, map[string]*bintree{
		"admin.graphql": &bintree{typeAdminGraphql, map[string]*bintree{}},
		"auth.graphql": &bintree{typeAuthGraphql, map[string]*bintree{}},
		"company.graphql": &bintree{typeCompanyGraphql, map[string]*bintree{}},
		"course.graphql": &bintree{typeCourseGraphql, map[string]*bintree{}},
		"users.graphql": &bintree{typeUsersGraphql, map[string]*bintree{}},
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
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

