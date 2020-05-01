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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x53\xcb\x6e\xdb\x30\x10\xbc\xf3\x2b\xd6\xe8\xa5\xbd\xe4\x03\x74\x4b\x6d\x14\x30\x10\xa3\x45\x2c\x7d\xc0\x56\xda\xc8\x44\x29\x92\xe5\xc3\xa8\x50\xf8\xdf\x03\xd1\x34\x45\x4a\xf2\x91\xa3\x99\xe1\xec\xac\xf8\x05\x06\xef\xd0\x71\x25\x2d\xb4\x28\x61\x50\x1d\xff\x18\xc1\x5d\x08\xba\xdf\x2f\x40\xff\xa8\xf5\x8e\x3a\xb0\xf4\xd7\x93\x74\x1c\x85\x18\x5f\x18\xe3\x52\x7b\x07\x8d\x16\x0a\xbb\x1f\x5c\xd0\x89\x1c\xc2\x7f\x06\xf0\xc1\x05\xd5\xa3\xa6\x0a\xce\xce\x70\xd9\xef\x18\x40\xab\xa4\x23\xe9\xde\x48\xf6\xee\x52\xc1\x51\xba\x1d\xbb\x31\xe6\x46\x4d\x99\xc5\x3b\x59\x1d\x2c\xbc\x11\xb9\xda\xfa\xb6\x25\x6b\x6b\xf5\x87\xe4\x8c\xdf\xd6\x19\xce\x77\x62\xf0\x78\x2a\x0a\x97\x9e\xe2\xc8\x81\x8a\xdd\xc0\xe5\x9b\xea\xb9\xfc\x1a\x1c\x2b\x78\x4d\xc8\xee\x5b\x05\xaf\xde\x5d\x82\x0f\x03\x18\x50\x62\x4f\xa6\x60\x9f\x32\x6c\xc9\x6f\x0d\xa1\xa3\xc8\x78\x08\xf6\x39\x78\x9c\xb0\x49\x16\xcf\xd3\xfc\xba\x5b\x8b\x9a\x1c\xdc\x10\xc5\x64\xbf\x8c\x9a\x36\x70\x2f\xe5\x7d\xda\x99\x75\xb3\x45\xbe\xad\x49\x5e\x96\xff\xc4\x25\xd6\xba\x76\x89\x1f\xca\x1c\x1d\x09\x5a\x85\x3f\xe4\x60\x0a\xff\x5d\x29\x41\x28\x77\xa9\xa7\xd0\x7b\xd9\x52\x80\x92\x22\x9c\x52\x43\x05\xbd\x99\xa1\x15\xfd\x9e\xa9\xa0\x1f\x66\xe8\x79\x9e\xbd\x1a\x34\xca\xb1\x4c\x14\xc1\xa4\x8a\xe7\x94\x6a\x21\x6a\x72\x70\x43\x54\xdc\xf4\x58\x18\x03\x98\x5e\x4d\xc0\x36\xef\x0d\x84\xb8\xac\xcd\x1f\x2a\x10\x0c\xb5\xa8\x5d\x7b\xc1\xfc\x39\x95\x83\xa2\xd6\x46\x5d\x53\x68\xef\x79\x57\x41\xd3\x1c\x0f\x65\x4a\x8b\x57\xfa\x29\x05\x97\xb4\x57\xde\x58\x7a\x4c\x77\x5e\xe0\x69\xc0\x1c\x8c\xfa\xbd\x40\x6b\x8d\x52\xc3\xda\x62\xf1\x69\xae\xa9\xc4\x53\x5d\x35\xf6\xe5\x52\x6a\xec\x93\xa8\xc6\x7e\xee\x15\x1d\xf5\xca\x2c\x57\x18\xd1\xf9\x9e\x08\xb0\x1b\xfb\x0c\x00\x00\xff\xff\xac\x42\x99\x5a\x10\x05\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 1296, mode: os.FileMode(420), modTime: time.Unix(1588339256, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\xcd\x6e\xdb\x3c\x10\xbc\xf3\x29\xc6\xc8\xc5\x06\x82\xef\x01\x78\x4b\xf2\x21\x80\x0f\x6d\x53\xd4\x79\x00\x5a\x5a\x49\xdb\xf0\x47\x25\xa9\x36\x46\x91\x77\x2f\x48\x59\x12\xe5\xb8\x39\xf4\x66\x69\x67\x76\x66\x56\xe3\x1b\xfc\x18\xc8\x33\x05\x84\xce\x0d\xba\x86\x75\x11\xc6\xd5\xdc\x9c\x10\x3b\x42\x7d\xfc\x0f\xf4\x4a\xd5\x10\xa9\x06\x5b\xf4\xca\x7b\xa5\x35\x69\x11\x4f\x3d\xe1\x49\xb5\xb4\xb7\x8d\xc3\x6f\x01\x44\x17\x95\x96\xd8\xdb\xb8\xc1\x0d\x3e\x0f\xe6\x48\x1e\xae\x41\xaf\x5a\x0a\x50\x4d\x24\x8f\xd8\x71\x80\xb3\x24\x00\xd7\x34\x81\xe2\x8c\x3f\x74\x74\x7e\x95\x38\x19\x97\x88\x68\xbc\x33\xd9\x4a\x88\xca\x47\x01\x68\x36\xbc\xa6\x19\xf5\x9a\xec\xff\xb2\xe8\xc9\x67\x96\x00\x5a\xfe\x49\xf6\x8a\x1b\x8e\x64\xc2\x38\x15\x6f\x42\xb0\xed\x87\x98\x73\xe4\x0c\x85\xa9\x95\xd4\x02\xfd\xa4\xac\x6a\xc9\x3f\xb2\x4e\x79\x12\xc7\x2a\x43\x12\xdf\xa2\x67\xdb\x0a\x80\x8c\x62\x5d\x3c\x7f\x77\xc7\x03\x47\x5d\x42\x86\x81\x6b\x89\xe7\xe7\xfd\xff\xe9\x6c\xa4\xa9\xef\x9c\x5d\x00\xb3\xd6\x83\x1b\x7c\xc8\x07\xfe\xbb\x9c\xaa\x2a\x0a\xe1\x70\xea\x49\xe2\x6e\xfe\x2d\x80\xa3\xaa\x5e\x5a\xef\x06\x5b\x3f\x74\x54\xbd\x48\xdc\x3b\xa7\x49\x59\x01\xf4\x9e\x2b\x92\x78\xd4\x4e\x15\xc9\xbe\x58\xcd\x96\x46\xcd\x42\xaf\x9a\x4d\xc8\x77\x86\x0a\xab\x5a\x85\xe0\x9d\x33\xff\xca\x77\xa6\x57\xf6\x54\xf0\x56\x47\x5a\x85\xce\xa4\x48\xbe\x51\x55\xae\x20\x5b\x15\xa9\xfe\x3a\x90\x3f\x65\x66\x7f\x6e\xa5\x9c\xfb\x59\x84\xf4\x35\xf9\xfb\x11\xa7\x42\x45\xb6\x66\xdb\x96\xb7\x69\x98\x74\x3d\x49\x6d\x12\x31\x37\x7d\xd9\xce\x79\xf3\x34\x07\x54\x6d\xd8\x6e\x17\xb7\x9b\x9d\xc4\x5d\x7a\x37\xcd\xc2\x36\x19\x1a\xcd\x4c\xb3\xf4\x3b\x91\xcd\xd8\xa6\x82\xbe\x93\x53\xc5\x96\x71\xb9\xe1\x16\x4d\xbe\x91\x5c\x37\xf1\x16\x6e\x4c\x26\xa7\x88\xcb\xa2\xa7\xf1\x1f\x51\x8d\x27\xbe\xb0\x7a\x3e\xfc\x3c\x67\xba\xae\xb6\xfa\x40\x57\xd5\xce\x88\xb3\x9a\x2b\xda\x14\xb6\x02\x18\xbf\xcb\xb8\x36\x3f\x4e\x9b\xdf\xf7\x2e\x8f\x2f\x15\x04\xb0\x5b\x83\xa7\x5c\xeb\xea\x7d\x2c\x76\xb5\xa7\x1f\xe8\x5d\xe0\xf3\xbe\x37\xf1\x27\x00\x00\xff\xff\x34\x13\x0e\xff\x37\x05\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 1335, mode: os.FileMode(420), modTime: time.Unix(1588339244, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\x4e\xcc\x49\x2c\x52\x08\x0d\xf5\x74\xe1\x82\xb2\x43\x32\x73\x53\xb9\xb8\x8a\x93\x33\x52\x73\x13\x15\xaa\xb9\x14\x14\x0a\x4b\x53\x8b\x2a\xad\x14\x02\x41\x14\x97\x82\x42\x6e\x69\x49\x62\x49\x66\x7e\x9e\x95\x82\x2f\x94\xc5\x55\xcb\x05\x08\x00\x00\xff\xff\xc2\x82\x45\xdf\x48\x00\x00\x00")

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

	info := bindataFileInfo{name: "schema.graphql", size: 72, mode: os.FileMode(420), modTime: time.Unix(1588339244, 0)}
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

	info := bindataFileInfo{name: "type/admin.graphql", size: 475, mode: os.FileMode(420), modTime: time.Unix(1588338400, 0)}
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

	info := bindataFileInfo{name: "type/auth.graphql", size: 36, mode: os.FileMode(420), modTime: time.Unix(1588338400, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeCompanyGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\x4f\x6b\x33\x21\x18\xc4\xef\x7e\x8a\xc9\xed\x7d\x21\x97\xf6\xe8\x2d\x49\x29\x04\xfa\x97\xb2\xa7\xd2\x83\xc4\x27\x22\xec\x3e\x8a\xba\x85\x25\xe4\xbb\x17\x75\x53\x96\xec\xe6\xda\xdb\x38\x8e\xfa\x73\x78\xd2\xe0\x09\x3b\xd7\x79\xc5\x03\x4e\x02\x50\xde\x07\xf7\x4d\x5a\x62\xeb\x5c\x4b\x8a\x05\xd0\xf7\x56\x4b\x34\xcd\xfe\x61\x25\x80\x43\x20\x95\x48\x6f\x92\xc4\x47\x0a\x96\x8d\x00\x58\x75\x74\x59\xe6\x4c\xa7\x58\x19\x0a\xf1\x9f\x57\x86\x24\xde\x94\xa1\x35\x8e\xb6\x4d\x14\x24\x9e\xeb\xe6\x63\x59\xae\xe1\x82\xa6\xb0\x1d\x24\x5e\xab\xf8\xff\x9b\xc8\xc7\xf2\x6d\x4a\xeb\x40\x31\x4a\x6c\xaa\x58\x89\xb3\x10\x85\x7c\x34\x2a\x79\xd5\x4f\x96\xe9\x6e\xca\x32\xf1\xef\xa7\xfe\xc1\xf5\x9c\x86\xa9\xe3\x5d\x4c\x3b\xa7\x69\x96\x0a\x93\xd8\xe5\xe9\xb1\xb4\xcc\x08\xdb\xf9\x96\x3a\xe2\x14\xf3\x57\x2d\xe7\x7e\xde\x7b\x0a\xb5\x52\xd2\x86\xa2\xc4\xe7\x78\xe2\x2b\xbf\xa4\x0c\xed\xf9\xe8\x6a\x35\x59\xe5\x7b\x2d\xfb\x3e\x61\x57\xfa\x1d\xc3\xfb\x62\x9d\x0a\x48\x31\x5e\xae\x9a\xfe\xbb\x5f\x57\xba\xc6\xeb\x25\xba\xab\x11\x99\xb3\xde\x18\xad\x25\xfc\x1b\xf4\x33\xf8\x05\xf6\x39\xba\x38\x8b\x9f\x00\x00\x00\xff\xff\xc7\xdd\x28\x06\xe4\x02\x00\x00")

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

	info := bindataFileInfo{name: "type/company.graphql", size: 740, mode: os.FileMode(420), modTime: time.Unix(1588339256, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeCourseGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x55\x4d\x6b\x1b\x3d\x10\xbe\xef\xaf\x18\xdb\x97\xf7\x85\xf8\x90\xd8\x0e\xed\xde\x52\xc7\x05\x43\x43\xdc\x78\x0d\x85\xe0\x83\xa2\x1d\x6f\x44\x76\xa5\x45\x9a\x2d\x31\x21\xff\xbd\x48\xf2\x5a\xf2\xa6\x71\x9b\xd0\x4b\x29\x32\x1e\xad\x34\xcf\xe8\x99\x0f\x8d\x50\x36\x15\x5c\x70\x8e\xc6\x64\xdb\x1a\xe1\x29\x01\xd0\x68\x48\x0b\x4e\x98\x27\x00\xaa\x46\x99\x3c\x27\x89\x53\x5c\x92\x6e\x38\x35\x1a\x67\x25\x56\x28\xc9\xa9\x57\x2a\x6f\x4a\x4c\x00\x08\x0d\x25\x00\x25\x1a\xa3\x02\xe6\xca\x6d\xc7\x80\x17\x7a\x64\x4f\x9e\xaa\x46\x1b\x9c\xcb\x8d\x72\x4a\x92\x55\x98\xda\x03\x85\x2c\x12\x80\x3b\xc6\x1f\x0a\xad\x1a\x99\x4f\xef\x91\x3f\xa4\xf0\x49\xa9\x12\x99\x4c\x00\x6a\x2d\x38\xa6\xf0\xb9\x54\xcc\x9a\xe5\xaa\x54\x3a\x42\xe2\x23\x47\x5d\x53\xb4\x22\x24\x69\x95\x37\x9c\x84\x92\xd1\xb2\xa9\x91\x8b\x8d\xe0\x19\xea\xca\x44\xeb\x9c\x11\x16\x4a\x6f\x53\x98\xee\x66\x7b\xce\xd7\xb2\x14\x12\x3d\x73\xc7\xba\x69\x44\x9e\xc2\x6a\x35\xbf\xec\xb9\x83\x36\x2a\x8d\x1c\xeb\xfd\x14\xb8\x60\x05\x82\xa8\x6a\x1f\x21\x03\x0b\x56\x08\xc9\x08\xf3\xaf\x0d\xea\xad\x33\x8b\x79\x81\x26\x85\xdb\x18\xb6\xb6\xae\xb3\xc2\x19\x4e\x2d\xc8\xcd\x42\x38\x4b\x66\x8c\x56\xaa\x7a\x07\xbb\x0e\xf6\x0d\x04\x3b\xc8\xd7\x39\x0a\x59\x37\x04\x4b\xf6\x1d\x3b\x90\xb9\xdb\x38\x64\xfb\xb2\x1c\xda\x9c\xd8\xdd\xbd\xce\x6f\x67\xfa\x48\x31\xb1\xfd\x55\x48\xa3\x6b\xf1\xeb\x2a\x23\x56\x58\xff\x5d\x6c\xd7\x47\x8a\xe9\x8e\x49\x89\x7a\x5e\xb1\x02\x97\x8d\xb3\x1f\x1b\x69\x48\xe9\xe0\x12\x0c\xbc\x50\x1b\xa0\x7b\x84\xcc\xee\xc2\x7f\xbc\x8d\x17\x70\x9f\x59\x25\xcb\xed\xff\xf6\x1e\xb2\xc7\x05\xd3\x24\xb8\xa8\x99\x24\x93\xc2\x5c\x5a\xa6\x86\x98\xa6\x4b\x46\x98\x42\x26\x2a\xeb\x0a\xca\xfc\xe0\xbb\x54\x9c\x1d\x44\xe8\x20\x3f\x71\xcd\xfd\xcb\xc9\x31\x6d\xef\xb3\x65\xee\xa3\x41\x58\xf5\xd6\x30\xb0\xb7\x79\x0b\x1b\xa5\x6d\x2a\x84\xc4\x5d\x62\x4c\x88\x63\xd0\xf7\x0d\xd0\x79\xd0\x6d\xa6\xbd\x97\x37\x94\xd0\xf2\xbb\xf5\x2d\xd4\x1f\xd7\x0b\x56\xc3\x72\x64\xf5\xa0\xdd\x76\x4d\x06\x42\x1a\x19\x61\xc6\x8a\x90\xd1\x38\x89\xbd\x6e\x10\x43\x6f\xc8\x58\x71\xbc\x00\x8e\x60\xdb\xfe\xf9\x1e\x03\x31\xef\xd6\xce\x5b\xc8\x0f\xe0\x16\x66\xdf\x2e\xae\x16\x5f\x66\x30\xbd\x5e\xdd\x2c\x67\xb0\xcc\x6e\x56\xd3\x6c\x75\x33\x4b\x06\x00\xf0\xe4\xe3\xd7\xf7\xcf\x59\xff\x64\x47\xb0\x3f\x1a\x8f\xc6\xa3\x61\xfc\x3f\x76\xa2\xff\x7c\xe2\x71\xd0\x01\xba\xd5\xd6\xbf\xfe\xe9\xd9\xc8\xfe\x86\x3b\x79\xb6\x9b\x38\x31\xda\x2b\xb7\x99\xde\x7d\x06\x36\xfe\x91\x0c\x6c\x4e\xcf\xdc\x18\x7a\x61\xad\xed\xbf\x5a\x42\x31\xdc\x3e\xb5\x01\xfc\xe1\xa3\x1f\xc3\xc9\xb9\x1b\xc3\xc9\xf9\x7e\x32\x39\x9f\x04\x03\x6b\x27\xff\x88\x83\x30\x00\xb1\x01\x61\x80\xb0\xaa\x4b\x46\x08\x15\x7b\x40\x60\xc0\x55\xbd\xfd\x1b\xbc\x7f\x85\xca\xc8\x15\xc1\xd8\x97\x84\x9d\x4c\x26\x63\x37\x86\xee\xbf\xdf\x81\xbf\x83\xca\x3a\xf9\x11\x00\x00\xff\xff\xb4\x63\xd0\x24\x9d\x09\x00\x00")

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

	info := bindataFileInfo{name: "type/course.graphql", size: 2461, mode: os.FileMode(420), modTime: time.Unix(1588339244, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeUsersGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x52\xc1\xae\xd3\x30\x10\xbc\xe7\x2b\x16\xf5\xf2\x90\xf8\x82\xdc\x50\xab\x27\x45\x02\xf4\x10\xe4\x84\x38\x98\x64\xed\xb7\xc8\x5e\x47\xf6\x46\x55\x85\xfa\xef\xc8\x4e\xd2\xa6\x8d\x0b\xe2\xc6\x6d\x33\xd9\xd9\xdd\x99\x31\xb1\x60\xd0\xaa\x43\x68\x23\x06\xf8\x55\x01\x8c\x23\xf5\x35\xb4\x6d\x73\x78\x53\x01\x74\x01\x95\x60\xff\x5e\x6a\xf8\x22\x81\xd8\x54\x00\xe8\x14\xd9\xe5\x3b\x35\x69\x0a\x51\x3e\x29\x87\x6b\xd0\xaa\x2d\x26\x68\x71\x78\xf5\xbc\x69\xfc\xe0\x0d\xf1\x1a\xfc\xe9\x7f\x7c\x25\xb1\xab\xc6\x73\x55\x11\x0f\xa3\xc0\x47\xc5\xca\x60\xc8\x94\x7c\xf1\xe6\x9e\x41\xc5\x78\xf4\xa1\x2f\x90\xf7\x59\xcf\x3c\xa2\xc9\x50\x1a\xd1\x79\x37\x28\x3e\x25\xd5\x93\x76\xd8\xc1\xb3\x0f\xa0\x7a\x47\x1c\xe1\xf8\x8a\x3c\x59\x41\x6c\xc0\x4d\xec\x08\x5e\x43\x4f\x5a\x63\x40\x96\x79\x04\x61\xfc\x17\x43\x36\xa7\x6f\x75\x3f\x70\xed\x0f\x1a\x0f\x68\xb1\xa0\x71\x1d\xec\xb9\xaa\xe4\x34\xe0\xe2\x25\x90\x1b\x2c\x3a\x64\x89\xff\xff\x43\xb8\xa4\x55\xc3\x7e\x2a\xb2\x1f\xc1\x6b\xb2\xd8\x38\x65\xb0\x0d\x97\xab\xae\xae\xb4\x43\x5f\x4a\xfe\x56\xe5\x8d\xa2\x92\xa0\x82\x9e\x92\x9c\xc2\xe1\xf7\x9e\xbf\x28\x83\x6b\xdf\x5f\x94\x21\x4e\x16\x7f\x1e\x31\x9c\xa6\x87\xdd\x1b\x8c\x35\x7c\x9b\x19\xdf\x73\xec\x06\x1b\xd6\xbe\x4e\xfd\xb9\x4a\x73\x77\x90\x27\xa7\xe0\x8d\x12\x2c\xc4\xb9\x83\x5b\xa9\xe9\xfb\x2e\xbe\x04\x15\x02\x4c\xf0\x36\xc2\x84\x16\x42\x5c\x9a\xe7\x18\x0f\x4a\x30\x63\xdb\x14\x13\xda\xf9\x31\x44\x8c\x4f\x5e\xeb\x88\x52\x43\xc3\xf2\x0e\x2c\x39\x5a\x6a\x4d\x56\x30\xd4\x17\x61\xfb\x4c\x78\xce\xe8\xdb\x14\x7f\xe6\x2f\xb3\x58\x66\xe0\xe9\x2f\xbc\x86\x65\xe2\x60\x10\xd2\xd4\x29\x79\x7c\x44\x5a\xb3\x6a\x5b\xed\x5a\x93\x37\x0b\xaf\x3f\xef\xb6\x9e\xab\xdf\x01\x00\x00\xff\xff\x93\xe1\xac\xe7\x74\x05\x00\x00")

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

	info := bindataFileInfo{name: "type/users.graphql", size: 1396, mode: os.FileMode(420), modTime: time.Unix(1588338400, 0)}
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

