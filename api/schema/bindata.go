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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\xd1\x8e\xdb\x2c\x10\x85\xef\x79\x8a\x89\xfe\x9b\xbf\x37\xfb\x00\xb9\xdb\x26\xaa\x14\x69\xa3\x56\x9b\xf8\x01\xa6\xf6\xac\x83\x8a\x81\xc2\xb0\xaa\x55\xe5\xdd\x2b\x13\x82\xc1\x4e\x2e\xf9\x32\xe7\x30\x9c\x13\xff\x07\x43\x60\x64\x69\xb4\x87\x16\x35\x0c\xa6\x93\x1f\x23\xf0\x85\xa0\xfb\xf9\x02\xf4\x87\xda\xc0\xd4\x81\xa7\xdf\x81\x34\x4b\x54\x6a\x7c\x11\x42\x6a\x1b\x18\x1a\xab\x0c\x76\xdf\xa4\xa2\x23\x31\xc2\x5f\x01\xf0\x21\x15\x9d\x47\x4b\x5b\x38\xb1\x93\xba\xdf\x08\x80\xd6\x68\x26\xcd\x6f\xa4\x7b\xbe\x6c\xe1\xa0\x79\x23\xae\x42\xf0\x68\xa9\xb0\x78\x27\x6f\xa3\x45\x70\xaa\x54\xfb\xd0\xb6\xe4\xfd\xd9\xfc\x22\x3d\xf3\xeb\x7a\x87\xd3\x6d\x30\x7a\x3c\x15\xc5\x4b\x8f\xe9\xc9\x71\x14\xbb\x41\xea\x37\xd3\x4b\xfd\x7f\x74\xdc\xc2\x6b\x26\x9b\x2f\x5b\x78\x0d\x7c\x89\x3e\x02\x60\x40\x8d\x3d\xb9\x6a\xfa\x58\xb0\xe5\x7c\x47\x8a\x7a\x64\xaa\x04\xfb\x12\x2e\x15\xad\x23\x64\x4a\x9e\x77\xc5\xae\x84\x87\x89\x4d\xb2\x74\x9e\x12\xb3\xdd\x5a\xd4\x94\xf0\x81\x28\xbd\xe5\x87\x33\x53\x67\xb7\x18\xdf\xa7\x96\x3d\xcf\x16\x65\xbf\x93\xbc\xae\xeb\x89\x4b\x2a\x62\xed\x92\x7e\xa8\xf7\x98\x32\x5a\x2d\xbf\x2f\x61\x5e\xfe\xab\x31\x8a\x50\x6f\x72\x4e\xb1\xa9\x3a\xa5\x88\xb2\x22\x9e\x72\x42\xd5\x78\x33\xa3\xd5\xf8\x6d\xa7\x6a\x7c\x3f\xa3\xe7\xfb\xdc\xab\xad\x57\xba\xd3\xac\xab\xf1\x94\xa4\xd1\x9e\xb2\xcb\xce\x0c\x16\xf5\x58\x9b\x24\x38\x7b\xdc\xce\xf9\x6d\x0b\x51\x53\xc2\x07\xa2\xea\xa6\x7b\xed\x02\x60\xfa\x5a\x23\x7b\x78\x6f\x1c\x48\x95\x3f\xfc\x5b\xc6\x01\x47\x2d\x5a\x6e\x2f\x58\x7e\xc6\x75\x5c\x68\xad\x33\x9f\x79\xe9\x10\x64\xb7\x85\xa6\x39\xec\xeb\x2d\x3d\x7e\xd2\x77\xad\xa4\xa6\x9d\x09\xce\xe7\x5c\x4f\x0b\x9e\x1f\x58\xc2\xa4\xdf\x29\xf4\xde\x19\x33\xac\x2d\x16\x3f\xcd\x31\xd5\x3c\xc7\x75\xc6\xbe\x2e\xe5\x8c\x7d\x16\x9d\xb1\x9f\x73\x45\xa6\xde\xb8\x65\x85\x89\xce\xf7\x24\x20\xae\xe2\x5f\x00\x00\x00\xff\xff\xd8\x91\xe8\xc6\x88\x05\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\x4b\x6e\xdc\x3a\x10\xdc\xf3\x14\x65\x78\x63\x03\xc6\x3b\x00\x77\xfe\xc0\x80\x17\x2f\x71\x10\x7b\x1d\xf4\x48\x2d\x89\x31\x3f\x0a\x49\x25\x1e\x04\xbe\x7b\x40\xea\x47\x8d\xc7\x03\x24\x3b\x51\x5d\xd5\x55\x4d\x56\x9f\xe3\xc7\xc0\x5e\x71\x40\xe8\xdc\xa0\x6b\x58\x17\x61\x5c\xad\x9a\x3d\x62\xc7\xa8\x77\xff\x81\x5f\xb9\x1a\x22\xd7\x50\x16\x3d\x79\x4f\x5a\xb3\x16\x71\xdf\x33\x1e\xa9\xe5\x07\xdb\x38\xfc\x16\x40\x74\x91\xb4\xc4\x83\x8d\x67\x38\xc7\xa7\xc1\xec\xd8\xc3\x35\xe8\xa9\xe5\x00\x6a\x22\x7b\xc4\x4e\x05\x38\xcb\x02\x70\x4d\x13\x38\x2e\xf8\xa7\x8e\xa7\x5f\x89\x93\x71\x89\x88\xc6\x3b\x93\xad\x84\x48\x3e\x0a\x40\x2b\xa3\xb6\x34\x43\xaf\xc9\xfe\x2f\x8b\x9e\x7d\x66\x09\xa0\x55\x3f\xd9\x1e\x71\xa3\x22\x9b\x30\x56\xc5\x9b\x10\xca\xf6\x43\xcc\x73\xe4\x19\x0a\x53\x1b\xa9\x15\xfa\x3f\x59\x6a\xd9\xdf\x2b\x9d\xe6\x49\x1c\x4b\x86\x25\xbe\x46\xaf\x6c\x2b\x00\x36\xa4\x74\x71\xfe\xee\x76\x4f\x2a\xea\x12\x32\x0c\xaa\x96\x78\x7e\x7e\xb8\x4b\xd7\xc6\x9a\xfb\xce\xd9\x15\xb0\x68\xdd\xb1\xe6\x96\x22\x17\x62\x5b\x6a\xac\xbe\xa5\xe3\xd2\xf8\xef\xad\x9c\x50\xbf\x75\x83\x0f\xf9\x79\x3f\x1e\x96\xaa\x8a\x43\x78\xda\xf7\x2c\x71\xbd\x7c\x0b\x60\x47\xd5\x4b\xeb\xdd\x60\xeb\xdb\x8e\xab\x17\x89\x1b\xe7\x34\x93\x15\x40\xef\x55\xc5\x12\xf7\xda\x51\x71\xaf\x9f\xad\x56\x96\x47\xcd\x42\xaf\x5a\x4c\xc8\x77\x86\x0a\xab\x9a\x42\xf0\xce\x99\x7f\xe5\x3b\xd3\x93\xdd\x7f\x74\xcf\x9b\xa1\x33\x29\xb2\x6f\xa8\xca\x0b\xa0\x2c\x45\xae\xbf\x0c\xec\xf7\x99\xd9\x4f\x3b\x21\x97\xed\x28\x86\xf4\x35\xfb\x9b\x11\x47\xa1\x62\x5b\x2b\xdb\x96\x77\xd3\x28\xd6\xcb\x83\x9e\x25\x62\xde\xb3\xb5\xbb\xca\x9d\xe7\x3a\x40\xb5\x51\xf6\x62\x75\x7b\x76\x29\x71\x9d\xfe\xcd\xb5\x70\x91\x0c\x8d\x66\xe6\x5a\xfa\x4e\x64\x33\x66\xb9\xa0\x5f\xca\x39\xe0\x6b\xb9\xec\x70\x85\x26\xdf\x91\xdc\xee\xc1\x15\xdc\x38\x99\x9c\x47\x5c\x1b\x3d\x8e\xfb\x58\x4f\x59\x3e\xf0\x3a\x47\xbc\x40\x1c\xd7\xdb\xee\xc2\x51\xc1\x19\x32\x29\x56\xe3\xa3\x1e\x08\x4e\x4f\xbd\xd4\xd5\x07\x7a\x9b\x48\x1c\x95\x9b\x10\x93\x9a\x2b\xf2\x1b\x2e\x04\x30\x26\x61\x6c\x9b\x8f\x73\xe7\xf7\x49\xcf\xe5\x43\x05\x01\x5c\x6e\xc1\xf3\x5c\xdb\xb0\x9f\x16\x3b\xba\x19\x27\xf4\x0e\xf0\xb9\xdf\x9b\xf8\x13\x00\x00\xff\xff\x1a\x8c\x7b\xb0\x27\x06\x00\x00")

func query_graphql() ([]byte, error) {
	return bindata_read(
		_query_graphql,
		"query.graphql",
	)
}

var _schema_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\x4e\xcc\x49\x2c\x52\x08\x0d\xf5\x74\xe1\x82\xb2\x43\x32\x73\x53\xb9\xb8\x8a\x93\x33\x52\x73\x13\x15\xaa\xb9\x14\x14\x0a\x4b\x53\x8b\x2a\xad\x14\x02\x41\x14\x97\x82\x42\x6e\x69\x49\x62\x49\x66\x7e\x9e\x95\x82\x2f\x94\xc5\x55\xcb\x05\x08\x00\x00\xff\xff\xc2\x82\x45\xdf\x48\x00\x00\x00")

func schema_graphql() ([]byte, error) {
	return bindata_read(
		_schema_graphql,
		"schema.graphql",
	)
}

var _type_admin_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xcd\x8a\x83\x30\x14\x85\xf7\x79\x8a\x3b\xaf\x91\xdd\x30\x6e\x84\x61\x70\x28\xae\x4a\x17\x17\x72\x0d\x17\x92\x18\xf2\x43\x11\xf1\xdd\x8b\xb1\x60\xd5\x52\xda\x5d\x38\x9c\x7c\xdf\x49\xd2\xe0\x09\xbe\x95\x65\x07\xa3\x00\xc8\x99\x95\x84\xb6\xad\xab\x2f\x01\x40\x16\xd9\x48\x38\xa5\xc0\x4e\xcf\x41\xc7\x21\xa6\x3f\xb4\xf4\x18\x1a\xdc\x67\x93\x10\x2b\xb7\x41\x4d\xc0\xd6\x1b\xb2\xe4\x52\x84\x06\x35\x3b\x4c\xa4\xfe\x33\x85\xa1\x58\x49\x69\x8a\x12\xce\xa5\x7f\x11\x00\x1e\x35\xd5\xae\xeb\xe5\xdc\x2e\xa7\x99\xc9\xce\xe7\xb4\x40\x7f\x7b\x7d\x5f\x7c\xd8\xe8\x31\xc6\x6b\x1f\xd4\x66\xce\x72\xf5\x27\x10\x26\x2a\x80\xba\x04\xe3\xfb\x6f\xfa\xcc\xd4\x7a\x75\x34\x6d\x3f\xf7\xe0\x7d\xa2\xdd\x5b\x57\x41\x45\x86\x5e\x0a\x26\x71\x0b\x00\x00\xff\xff\x68\xfa\x3e\xad\xdb\x01\x00\x00")

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

var _type_company_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\x4f\x6b\x33\x21\x18\xc4\xef\x7e\x8a\xc9\xed\x7d\x21\x97\xf6\xe8\x2d\x49\x29\x04\xfa\x97\xb2\xa7\xd2\x83\xc4\x27\x22\xec\x3e\x8a\xba\x85\x25\xe4\xbb\x17\x75\x53\x96\xec\xe6\xda\xdb\x38\x8e\xfa\x73\x78\xd2\xe0\x09\x3b\xd7\x79\xc5\x03\x4e\x02\x50\xde\x07\xf7\x4d\x5a\x62\xeb\x5c\x4b\x8a\x05\xd0\xf7\x56\x4b\x34\xcd\xfe\x61\x25\x80\x43\x20\x95\x48\x6f\x92\xc4\x47\x0a\x96\x8d\x00\x58\x75\x74\x59\xe6\x4c\xa7\x58\x19\x0a\xf1\x9f\x57\x86\x24\xde\x94\xa1\x35\x8e\xb6\x4d\x14\x24\x9e\xeb\xe6\x63\x59\xae\xe1\x82\xa6\xb0\x1d\x24\x5e\xab\xf8\xff\x9b\xc8\xc7\xf2\x6d\x4a\xeb\x40\x31\x4a\x6c\xaa\x58\x89\xb3\x10\x85\x7c\x34\x2a\x79\xd5\x4f\x96\xe9\x6e\xca\x32\xf1\xef\xa7\xfe\xc1\xf5\x9c\x86\xa9\xe3\x5d\x4c\x3b\xa7\x69\x96\x0a\x93\xd8\xe5\xe9\xb1\xb4\xcc\x08\xdb\xf9\x96\x3a\xe2\x14\xf3\x57\x2d\xe7\x7e\xde\x7b\x0a\xb5\x52\xd2\x86\xa2\xc4\xe7\x78\xe2\x2b\xbf\xa4\x0c\xed\xf9\xe8\x6a\x35\x59\xe5\x7b\x2d\xfb\x3e\x61\x57\xfa\x1d\xc3\xfb\x62\x9d\x0a\x48\x31\x5e\xae\x9a\xfe\xbb\x5f\x57\xba\xc6\xeb\x25\xba\xab\x11\x99\xb3\xde\x18\xad\x25\xfc\x1b\xf4\x33\xf8\x05\xf6\x39\xba\x38\x8b\x9f\x00\x00\x00\xff\xff\xc7\xdd\x28\x06\xe4\x02\x00\x00")

func type_company_graphql() ([]byte, error) {
	return bindata_read(
		_type_company_graphql,
		"type/company.graphql",
	)
}

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x55\x4d\x6b\x1b\x3d\x10\xbe\xef\xaf\x18\xdb\x97\xf7\x85\xf8\x90\xd8\x0e\xed\xde\x52\xc7\x05\x43\x43\xdc\x78\x0d\x85\xe0\x83\xa2\x1d\x6f\x44\x76\xa5\x45\x9a\x2d\x31\x21\xff\xbd\x48\xf2\x5a\xf2\xa6\x71\x9b\xd0\x4b\x29\x32\x1e\xad\x34\xcf\xe8\x99\x0f\x8d\x50\x36\x15\x5c\x70\x8e\xc6\x64\xdb\x1a\xe1\x29\x01\xd0\x68\x48\x0b\x4e\x98\x27\x00\xaa\x46\x99\x3c\x27\x89\x53\x5c\x92\x6e\x38\x35\x1a\x67\x25\x56\x28\xc9\xa9\x57\x2a\x6f\x4a\x4c\x00\x08\x0d\x25\x00\x25\x1a\xa3\x02\xe6\xca\x6d\xc7\x80\x17\x7a\x64\x4f\x9e\xaa\x46\x1b\x9c\xcb\x8d\x72\x4a\x92\x55\x98\xda\x03\x85\x2c\x12\x80\x3b\xc6\x1f\x0a\xad\x1a\x99\x4f\xef\x91\x3f\xa4\xf0\x49\xa9\x12\x99\x4c\x00\x6a\x2d\x38\xa6\xf0\xb9\x54\xcc\x9a\xe5\xaa\x54\x3a\x42\xe2\x23\x47\x5d\x53\xb4\x22\x24\x69\x95\x37\x9c\x84\x92\xd1\xb2\xa9\x91\x8b\x8d\xe0\x19\xea\xca\x44\xeb\x9c\x11\x16\x4a\x6f\x53\x98\xee\x66\x7b\xce\xd7\xb2\x14\x12\x3d\x73\xc7\xba\x69\x44\x9e\xc2\x6a\x35\xbf\xec\xb9\x83\x36\x2a\x8d\x1c\xeb\xfd\x14\xb8\x60\x05\x82\xa8\x6a\x1f\x21\x03\x0b\x56\x08\xc9\x08\xf3\xaf\x0d\xea\xad\x33\x8b\x79\x81\x26\x85\xdb\x18\xb6\xb6\xae\xb3\xc2\x19\x4e\x2d\xc8\xcd\x42\x38\x4b\x66\x8c\x56\xaa\x7a\x07\xbb\x0e\xf6\x0d\x04\x3b\xc8\xd7\x39\x0a\x59\x37\x04\x4b\xf6\x1d\x3b\x90\xb9\xdb\x38\x64\xfb\xb2\x1c\xda\x9c\xd8\xdd\xbd\xce\x6f\x67\xfa\x48\x31\xb1\xfd\x55\x48\xa3\x6b\xf1\xeb\x2a\x23\x56\x58\xff\x5d\x6c\xd7\x47\x8a\xe9\x8e\x49\x89\x7a\x5e\xb1\x02\x97\x8d\xb3\x1f\x1b\x69\x48\xe9\xe0\x12\x0c\xbc\x50\x1b\xa0\x7b\x84\xcc\xee\xc2\x7f\xbc\x8d\x17\x70\x9f\x59\x25\xcb\xed\xff\xf6\x1e\xb2\xc7\x05\xd3\x24\xb8\xa8\x99\x24\x93\xc2\x5c\x5a\xa6\x86\x98\xa6\x4b\x46\x98\x42\x26\x2a\xeb\x0a\xca\xfc\xe0\xbb\x54\x9c\x1d\x44\xe8\x20\x3f\x71\xcd\xfd\xcb\xc9\x31\x6d\xef\xb3\x65\xee\xa3\x41\x58\xf5\xd6\x30\xb0\xb7\x79\x0b\x1b\xa5\x6d\x2a\x84\xc4\x5d\x62\x4c\x88\x63\xd0\xf7\x0d\xd0\x79\xd0\x6d\xa6\xbd\x97\x37\x94\xd0\xf2\xbb\xf5\x2d\xd4\x1f\xd7\x0b\x56\xc3\x72\x64\xf5\xa0\xdd\x76\x4d\x06\x42\x1a\x19\x61\xc6\x8a\x90\xd1\x38\x89\xbd\x6e\x10\x43\x6f\xc8\x58\x71\xbc\x00\x8e\x60\xdb\xfe\xf9\x1e\x03\x31\xef\xd6\xce\x5b\xc8\x0f\xe0\x16\x66\xdf\x2e\xae\x16\x5f\x66\x30\xbd\x5e\xdd\x2c\x67\xb0\xcc\x6e\x56\xd3\x6c\x75\x33\x4b\x06\x00\xf0\xe4\xe3\xd7\xf7\xcf\x59\xff\x64\x47\xb0\x3f\x1a\x8f\xc6\xa3\x61\xfc\x3f\x76\xa2\xff\x7c\xe2\x71\xd0\x01\xba\xd5\xd6\xbf\xfe\xe9\xd9\xc8\xfe\x86\x3b\x79\xb6\x9b\x38\x31\xda\x2b\xb7\x99\xde\x7d\x06\x36\xfe\x91\x0c\x6c\x4e\xcf\xdc\x18\x7a\x61\xad\xed\xbf\x5a\x42\x31\xdc\x3e\xb5\x01\xfc\xe1\xa3\x1f\xc3\xc9\xb9\x1b\xc3\xc9\xf9\x7e\x32\x39\x9f\x04\x03\x6b\x27\xff\x88\x83\x30\x00\xb1\x01\x61\x80\xb0\xaa\x4b\x46\x08\x15\x7b\x40\x60\xc0\x55\xbd\xfd\x1b\xbc\x7f\x85\xca\xc8\x15\xc1\xd8\x97\x84\x9d\x4c\x26\x63\x37\x86\xee\xbf\xdf\x81\xbf\x83\xca\x3a\xf9\x11\x00\x00\xff\xff\xb4\x63\xd0\x24\x9d\x09\x00\x00")

func type_course_graphql() ([]byte, error) {
	return bindata_read(
		_type_course_graphql,
		"type/course.graphql",
	)
}

var _type_delegate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x4d\x4e\xf3\x30\x14\xdc\xfb\x14\xaf\xfb\xef\x04\xde\x7d\xb4\x9b\x48\x08\x15\x68\x57\x08\xa1\x47\x3d\x35\x46\xfe\x93\xed\x08\x55\x55\xef\x8e\xec\x12\x92\xb4\xdd\xb0\xb3\x46\x33\x93\x99\x37\x29\x87\x08\x5a\xc1\x42\x73\x01\xd1\x51\x10\xf5\xbd\x51\x92\xb6\xdb\x6e\xb5\x10\x44\xbb\x04\x2e\x50\xff\x8b\xa4\xe7\x92\x8c\xd7\x82\x68\xb3\x59\xbe\x75\xab\x01\xa8\x2c\x38\x36\x76\xc2\xd8\x9b\x94\xcb\x03\x3b\x4c\x49\x96\xaf\xb1\x02\x8b\xf8\x11\x3c\x26\xe2\xca\xbb\x0f\xda\xf8\x29\xf1\x33\xbc\x6f\x4c\xb1\x33\xf1\x2e\xb8\xc8\xfe\x20\x69\x79\x7e\x54\x2c\xa6\xb0\x37\x16\x9d\x63\x8d\x6d\x1a\x33\x9d\x84\x30\x3e\xf6\x85\x96\xad\xd0\x50\xb9\x6b\xd8\x71\x34\xab\xbd\xcf\xed\xff\x52\xe3\xb2\xff\xad\xb4\xb7\xaa\x6a\x78\x24\x2e\x58\x73\xce\x5f\x21\x29\x49\x77\x21\x58\xb0\xaf\x79\xdb\x36\xf3\xb8\x4f\xc8\x31\xf8\x8c\x96\x58\xfd\x80\xf2\x77\xc0\xc5\xbf\x89\xa7\x1a\x4d\xc7\x1b\xcc\xf6\x5e\xb3\x06\x19\x17\x2d\x1c\x7c\xc9\xb4\x66\x6d\x7c\x55\x3e\xf6\x48\x87\xf6\x0d\x28\x8d\x2c\xe9\x65\x90\xbc\xd6\x13\xb3\x46\xe7\xf7\x41\x56\x41\x7b\x8d\xd7\x1d\x78\x6d\xbf\xe6\x70\xfd\xb3\xc4\x8b\x5c\x0b\x71\x12\xdf\x01\x00\x00\xff\xff\x96\x86\x13\xaf\x88\x02\x00\x00")

func type_delegate_graphql() ([]byte, error) {
	return bindata_read(
		_type_delegate_graphql,
		"type/delegate.graphql",
	)
}

var _type_manager_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\x5d\x4b\x03\x31\x10\x7c\xcf\xaf\x58\xf1\x5f\xdc\x9b\xb4\x08\x07\x2a\x15\xbd\x27\xf1\x61\x6d\xf6\xd2\x95\x64\x13\x92\x1c\xa5\x94\xfe\x77\xc9\x5d\x3f\xae\x36\x08\xbe\x25\xc3\xec\x30\x33\xbb\x79\x17\x08\x9e\x51\xd0\x50\x04\xd8\x2b\x80\x61\x60\xdd\x40\xd7\xb5\xcb\x3b\x05\xb0\x8e\x84\x99\xf4\x43\x6e\xe0\x2d\x47\x16\xa3\x00\xc8\x21\xdb\xd3\xbf\x90\x7a\x8e\x29\xbf\xa0\xa3\x39\x68\xf1\x16\xcb\x64\x29\x6c\xbc\xdc\x10\x9f\xbc\x61\x99\x83\xdf\xfe\xeb\x9d\xb3\xbd\x22\xae\xbd\x0b\x28\xbb\x06\x16\xd3\xa3\x60\x21\xfa\x9e\x2d\xb5\x0e\x0d\x75\xf1\xec\x4a\x1d\x94\x9a\x47\x5b\xa1\x21\x60\x17\x2c\x39\x92\x9c\x60\x85\x86\xa5\x04\x7b\x1d\x28\xee\xc6\xdc\xa4\x0d\xa5\x06\x3e\x8e\x13\x9f\x45\x1c\x0d\xb5\xd2\xfb\xa6\xf0\xc7\x57\xd1\x65\x09\x43\x3e\x09\x8f\xc6\xa7\xf9\xdf\xad\x04\x4c\x69\xeb\xa3\xbe\x60\xe7\xe1\xc5\xd8\xea\x51\xa2\x1d\xa1\xfd\x25\x60\xe9\x7e\xda\x00\xdc\xc3\xa3\x8f\x80\xda\xb1\x24\xd8\x6e\x48\xa6\x85\xb0\x18\x70\xd3\x74\x02\xdf\x83\xe6\xbe\xa7\x48\x92\x8f\x12\x4c\xe9\x3f\x6b\xb9\xb1\x5e\x6b\xbf\xba\xbb\x3f\x32\x2e\xc9\x52\x25\xe3\xfc\xbc\xce\xdc\x2e\xe8\x5a\x1f\xd7\xa7\x78\xe5\xb2\x16\xaf\x92\xae\xe6\xbb\x92\x4f\x1d\xd4\x4f\x00\x00\x00\xff\xff\xd4\x4d\x42\xec\x09\x03\x00\x00")

func type_manager_graphql() ([]byte, error) {
	return bindata_read(
		_type_manager_graphql,
		"type/manager.graphql",
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
	"type/course.graphql": type_course_graphql,
	"type/delegate.graphql": type_delegate_graphql,
	"type/manager.graphql": type_manager_graphql,
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
		"course.graphql": &_bintree_t{type_course_graphql, map[string]*_bintree_t{
		}},
		"delegate.graphql": &_bintree_t{type_delegate_graphql, map[string]*_bintree_t{
		}},
		"manager.graphql": &_bintree_t{type_manager_graphql, map[string]*_bintree_t{
		}},
	}},
}}
