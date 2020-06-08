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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\x4d\x8e\xdb\x30\x0c\x85\xf7\x3a\x05\x83\x6e\xda\xcd\x1c\xc0\xbb\x69\x82\x02\x01\x12\xb4\x98\xd8\x07\x60\x6d\x8e\x22\x54\x96\x54\xfd\x0c\x6a\x14\xb9\xfb\xc0\xb2\x23\x4b\x76\xb2\xd4\x67\xbe\x47\xea\xd1\xfa\x02\x7d\xf0\xe8\x85\x56\x0e\x5a\x54\xd0\xeb\x4e\xbc\x0f\xe0\xaf\x04\xdd\xef\x17\xa0\x7f\xd4\x06\x4f\x1d\x38\xfa\x1b\x48\x79\x81\x52\x0e\x2f\x8c\x09\x65\x82\x87\xc6\x48\x8d\xdd\x0f\x21\xe9\x4c\x1e\xe1\x3f\x03\x78\x17\x92\xea\xc1\x50\x05\x17\x6f\x85\xe2\x3b\x06\xd0\x6a\xe5\x49\xf9\x13\x29\xee\xaf\x15\x1c\x95\xdf\xb1\x1b\x63\x7e\x30\x94\x59\xbc\x91\x33\xd1\x22\x58\x99\xab\x5d\x68\x5b\x72\xae\xd6\x7f\x48\x2d\xfc\xb6\x9d\xe1\x32\x15\x46\x8f\xa7\xa2\xd8\xf4\x3c\x5f\x39\x96\x62\xd7\x0b\x75\xd2\x5c\xa8\xaf\xd1\xb1\x82\xd7\x44\x76\xdf\x2a\x78\x0d\xfe\x1a\x7d\x18\x40\x8f\x0a\x39\xd9\xa2\xfa\x9c\xb1\x75\x7d\x47\x92\x38\x7a\x2a\x04\x87\x1c\xae\x15\xad\x25\xf4\x34\x7b\xde\x15\xfb\x1c\x1e\x47\x36\xca\xe6\xf3\x98\x98\xe9\xb6\xa2\x26\x87\x0f\x44\xc6\xea\x71\x59\xc7\x1e\x39\x4d\x21\xbe\x8d\x3b\x76\x7e\x31\xc8\xb7\x3b\x8a\xcb\x65\xad\x1b\xff\xca\x0c\xb7\x1e\xf3\x76\xca\x19\xc6\x7c\x36\x83\x1f\x72\x98\x06\xff\xae\xb5\x24\x54\xbb\x94\x51\xdc\x52\x99\x50\x44\x49\x11\x4f\x69\xc8\xa2\xbc\x59\xd0\xa6\x7c\x9a\xa9\x28\x3f\x2c\xe8\xf9\x3c\xf7\xb5\x96\x23\xdd\x69\xd2\x95\x78\xcc\x51\x2b\x47\xc9\x65\xaf\x7b\x83\x6a\x28\x4d\x66\xb8\x78\x4c\xe7\x74\xb7\x95\xa8\xc9\xe1\x03\x51\xd1\xe9\xbe\x74\x06\x30\xbe\xd4\xc8\x1e\xf6\x8d\x05\xf3\x03\x78\xf8\x4b\xc6\x02\x4b\x2d\x1a\xdf\x5e\x31\x7f\xc2\x65\x5c\x68\x8c\xd5\x1f\x69\xe8\x10\x44\x57\x41\xd3\x1c\x0f\xe5\x94\x0e\x3f\xe8\xa7\x92\x42\xd1\x5e\x07\xeb\x52\xae\x97\x15\x4f\x17\xcc\xe1\xac\xdf\x4b\x74\xce\x6a\xdd\x6f\x2d\x56\x9f\x96\x98\x4a\x9e\xe2\xaa\x91\x97\x4b\xa9\x91\x27\x51\x8d\x7c\xc9\x15\x3d\x71\x6d\xd7\x2b\x9c\xe9\xd2\x67\x06\x49\x77\x22\xe7\xf4\xea\x87\x9e\x58\xd2\x4c\x47\x76\x63\x9f\x01\x00\x00\xff\xff\xac\x85\xdf\x8d\xb6\x05\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x53\xcb\x6e\xdb\x30\x10\xbc\xf3\x2b\xc6\xc8\xc5\x06\x82\x7e\x00\x6f\x79\x20\x80\x0f\x6d\x53\xd4\x3e\x17\xb4\xb4\x92\xd8\xf0\xa1\x92\x54\x1b\xa3\xc8\xbf\x17\xa4\x5e\x94\x63\x1b\x68\x6e\xa2\x76\x66\x67\x96\x9c\xbd\xc1\xaf\x8e\x9c\x24\x0f\xdf\xd8\x4e\x95\x30\x36\x40\xdb\x52\x56\x47\x84\x86\x50\x1e\x3e\x81\x5e\xa9\xe8\x02\x95\x90\x06\xad\x70\x4e\x28\x45\x8a\x85\x63\x4b\x78\x16\x35\x6d\x4d\x65\xf1\x97\x01\xc1\x06\xa1\x38\xb6\x26\xac\x70\x83\x2f\x9d\x3e\x90\x83\xad\xd0\x8a\x9a\x3c\x44\x15\xc8\x21\x34\xd2\xc3\x1a\x62\x80\xad\x2a\x4f\x61\xc2\xef\x1a\x1a\x7e\x45\x4e\xc2\x45\x22\x2a\x67\x75\xb2\xe2\x83\x70\x81\x01\x4a\x6a\xb9\xa4\x69\xf1\x1a\xed\xff\x31\x68\xc9\x25\x16\x03\x6a\xf9\x9b\xcc\x19\x37\x32\x90\xf6\x7d\x95\xbd\x31\x26\x4d\xdb\x85\x34\x47\x9a\x21\x33\xb5\x90\x9a\xa1\x9f\x85\x11\x35\xb9\x27\xa9\xe2\x3c\x91\x63\x84\x26\x8e\xef\xc1\x49\x53\x33\x80\xb4\x90\x2a\x3b\xff\xb4\x87\x9d\x0c\x2a\x87\x74\x9d\x2c\x39\xf6\xfb\xed\x63\xbc\x36\x52\xd4\x36\xd6\xcc\x80\x49\xeb\x91\x14\xd5\x22\x50\x26\xb6\xa4\x86\xe2\x47\x3c\x4e\x8d\xff\xdf\xca\x15\xf5\x07\xdb\x39\x9f\x9e\xf7\xf2\xb0\xa2\x28\xc8\xfb\xdd\xb1\x25\x8e\xbb\xe9\x9b\x01\x07\x51\xbc\xd4\xce\x76\xa6\x7c\x68\xa8\x78\xe1\xb8\xb7\x56\x91\x30\x0c\x68\x9d\x2c\x88\xe3\x49\x59\x91\xdd\xeb\x57\xa3\xa4\xa1\x5e\x33\xd3\x2b\x26\x13\xfc\x9d\xa1\xcc\xaa\x12\xde\x3b\x6b\xf5\x47\xf9\x56\xb7\xc2\x1c\x2f\xdd\xf3\x62\xe8\x44\x0a\xe4\x2a\x51\xa4\x05\x90\x46\x04\x2a\xbf\x75\xe4\x8e\x89\xd9\x0e\x3b\xc1\xa7\xed\xc8\x86\x74\x25\xb9\xfb\x1e\x27\x7c\x41\xa6\x94\xa6\xce\xef\xa6\x92\xa4\xa6\x07\x5d\x45\x62\xda\xb3\xb9\xbb\x4c\x9d\xc7\x3a\x20\x4a\x2d\xcd\x7a\x76\xbb\xda\x70\xdc\xc5\x7f\x63\xcd\xaf\xa3\xa1\xde\xcc\x58\x8b\xdf\x91\xac\xfb\x2c\x67\xf4\x0d\x1f\x03\x3e\x97\xf3\x0e\xb7\xa8\xd2\x1d\xf1\xe5\x1e\xdc\xc2\xf6\x93\xf1\x71\xc4\xb9\xd1\x73\xbf\x8f\xe5\x90\xe5\x13\xaf\x63\xc4\x33\xc4\x79\xbd\xe5\x2e\x9c\x15\x1c\x21\x83\x62\xd1\x3f\xea\x89\xe0\xf0\xd4\x53\x5d\x5e\xd0\x5b\x44\xe2\xac\xdc\x80\x18\xd4\x6c\x96\x5f\xbf\x66\x40\x9f\x84\xbe\x6d\x3a\x8e\x9d\xdf\x27\x3d\x95\x4f\x15\x18\xb0\x59\x82\xc7\xb9\x96\x61\xbf\x2e\x76\x76\x33\xae\xe8\x9d\xe0\x87\x7e\x9d\x8f\xbd\xf6\x9e\x1c\xd8\x1b\xfb\x17\x00\x00\xff\xff\x15\xe2\x85\x35\x35\x06\x00\x00")

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

var _type_delegate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x41\x4f\xc3\x30\x0c\x85\xef\xfd\x15\xde\x9d\x5f\xd0\x1b\x6c\x97\x4a\x08\x0d\xe8\x4e\x08\x21\xb3\xbc\x85\x40\x1a\x47\x49\x2a\x34\x4d\xfb\xef\x28\x19\xa5\x5d\xb7\x0b\xb7\xe8\xe9\x3d\xdb\x9f\x9d\xb4\xf7\xa0\x15\x2c\x34\x27\x10\x1d\x2a\xa2\xbe\x37\xaa\xa6\xcd\xa6\x59\x2d\x2a\xa2\x6d\x00\x27\xa8\xdb\x54\xd3\x73\x0a\xc6\xe9\x8a\xa8\x6d\x97\x6f\xcd\x6a\x10\xb2\x0b\x1d\x1b\x3b\x71\xec\x4c\x88\xe9\x81\x3b\x4c\x4d\x96\x2f\xb5\x04\x0b\xff\x21\x0e\x93\x70\xf6\xdd\x8b\x36\x6e\x6a\xfc\x94\xf7\xd6\x24\x7b\x16\xde\x4a\xe7\xd9\xed\x6b\x5a\x9e\x1e\x59\xf3\x41\x76\xc6\xa2\xe9\x58\x63\x13\xc6\x99\x8e\x55\x65\x9c\xef\x13\x2d\x0b\xd0\x80\xdc\x14\xed\x30\x16\xcb\xdc\x27\xfa\xff\x60\xcc\xf9\xaf\x4d\x7b\x0d\x55\xc3\x21\x70\xc2\x9a\x63\xfc\x96\xa0\x6a\xba\x13\xb1\x60\x37\x27\xf1\x56\x58\xb5\xf2\x05\x37\x25\x2a\xd7\x3b\x07\x7a\x42\xf4\xe2\x22\x0a\x93\xfa\x15\xeb\xbf\x13\x2f\x6e\x26\x5d\xd5\xd8\x76\x56\x73\xb0\xaf\x59\x83\x4c\xe7\x2d\x3a\xb8\x14\x69\xcd\xda\xb8\x9c\x7c\xec\x11\xf6\xa5\x07\x94\x46\xac\xe9\x65\x88\xbc\xe6\xd1\x59\xa3\x71\x3b\xa9\x73\xa0\xbc\xc6\xfd\x0f\xbe\x72\xe1\x52\xe1\xf2\x3b\xf9\xd9\x5c\x59\x73\x92\xd1\xc6\x05\x1d\xab\x9f\x00\x00\x00\xff\xff\x18\xe4\xd4\x61\xbc\x02\x00\x00")

func type_delegate_graphql() ([]byte, error) {
	return bindata_read(
		_type_delegate_graphql,
		"type/delegate.graphql",
	)
}

var _type_lesson_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xf0\x49\x2d\x2e\xce\xcf\x53\xa8\xe6\x52\x50\x28\x2d\xcd\x4c\xb1\x52\x08\x0d\xf5\x74\x51\xe4\x52\x50\x28\xc9\x2c\xc9\x49\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x07\x0b\x24\xa6\x17\x5b\x29\x44\x87\x24\xa6\x2b\xc6\x82\xf9\xa9\x15\x25\x08\xf9\x5a\x2e\xae\xcc\xbc\x82\xd2\x12\x05\xe7\xa2\xd4\xc4\x92\x54\x88\xb1\x9e\x60\x91\x6a\x3c\xa6\x81\x6d\x8b\xc5\x34\x0d\x10\x00\x00\xff\xff\x00\x95\x8f\xcf\x9c\x00\x00\x00")

func type_lesson_graphql() ([]byte, error) {
	return bindata_read(
		_type_lesson_graphql,
		"type/lesson.graphql",
	)
}

var _type_manager_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xcd\x6e\xe3\x30\x0c\x84\xef\x7a\x0a\x2e\xf6\x2d\x7c\xdb\x4d\xb0\x80\x81\x6d\x91\xfe\xf8\x54\xf4\xc0\x46\xb4\xc2\x42\xa2\x04\x49\x46\x10\x04\x79\xf7\x42\x76\x7e\x9c\x5a\x97\xde\xec\xc1\x70\xc0\x8f\xa3\x7c\x08\x04\x0f\x28\x68\x28\x02\x1c\x15\xc0\x30\xb0\x6e\xa0\xeb\xda\xf5\x2f\x05\xb0\x8d\x84\x99\xf4\x9f\xdc\xc0\x4b\x8e\x2c\x46\x01\x90\x43\xb6\x97\xff\x62\xea\x39\xa6\xfc\x88\x8e\xe6\xa2\xc5\xa5\x96\xc9\x52\xd8\x79\x59\x18\xff\x7b\xc3\x32\x17\x3f\xfd\xc7\x2b\x67\x7b\x67\xdc\x7a\x17\x50\x0e\x0d\xac\xa6\x8f\xa2\x85\xe8\x7b\xb6\xd4\x3a\x34\xd4\xc5\xeb\x56\xea\xa4\xd4\x1c\x6d\x83\x86\x80\x5d\xb0\xe4\x48\x72\x82\x0d\x1a\x96\x02\xf6\x34\x50\x3c\x8c\xdc\xa4\x0d\xa5\x06\xde\xce\x13\xef\x25\x1c\x0d\xb5\xd2\xfb\xa6\xf8\xc7\xaf\x92\xcb\x12\x86\x7c\x09\x1e\x17\x9f\xe6\xbf\x5f\x25\x60\x4a\x7b\x1f\xf5\x5c\x13\xff\x4c\x29\x34\xf0\xd7\x7b\x4b\x28\xb7\xb8\xd5\x78\xe7\x73\x68\x3b\x4a\xc7\x1b\x72\x69\x63\xea\x04\x7e\xc3\x3f\x1f\x01\xb5\x63\x49\xb0\xdf\x91\x4c\x15\xb1\x18\x70\xd3\x74\x02\xdf\x83\xe6\xbe\xa7\x48\x92\xcf\x11\x4c\xe9\x27\x45\x2d\x60\x6a\x7d\x54\xdb\x5c\x52\x5f\x19\xd7\x64\xa9\xc2\x38\x7f\x70\x57\x6f\x17\x74\xed\x1e\xf7\x8f\xf3\x6e\xcb\x1a\x5e\x85\xae\xb6\x77\x85\x4f\x9d\xd4\x57\x00\x00\x00\xff\xff\xf3\x0d\x49\x60\x1b\x03\x00\x00")

func type_manager_graphql() ([]byte, error) {
	return bindata_read(
		_type_manager_graphql,
		"type/manager.graphql",
	)
}

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xc1\x8a\x1c\x31\x0c\x44\xef\xfe\x8a\x0a\x7d\x0d\xfb\x01\x7d\x0b\x7b\x0a\x84\x24\xb0\xbb\x1f\xa0\x69\xab\x3d\x0a\xb6\xd4\xb1\xd5\x03\x4d\xd8\x7f\x0f\xf6\x30\x49\xc3\xde\xec\x42\xf5\xaa\x24\xd6\xbd\xe0\xad\x71\x7d\x3d\x36\xc6\x9f\x00\x14\x52\x4a\x5c\x03\x20\x1a\xe5\x26\x71\xa7\x1c\x80\xc8\x99\x13\x39\x87\xf7\x10\xa6\xe1\x80\x34\x10\xfc\x2a\x0a\xef\x66\xbf\x92\x63\xb1\x9c\xc9\xb9\x41\x74\xb5\x5a\xc8\xc5\x14\x74\xb1\xdd\xe1\x57\x7e\xc0\x3f\xff\xc7\x4d\x20\x8d\xa7\x28\xec\x1d\xdd\x81\xed\x09\x2f\x56\x18\xab\x70\x8e\x0d\x54\x19\x6a\x8e\x64\xa2\x09\x6e\xb8\x30\xe8\x46\x92\xe9\x92\x19\x91\x37\xd6\x28\x9a\xc2\x04\xd3\x91\x35\x4a\xd9\x7a\x07\x1e\xb6\x0f\x42\xe5\xdf\x3b\x37\xef\x88\x48\x4e\x58\xad\x3e\x85\x31\x39\x56\xea\x07\x98\xb0\x70\x75\x59\x65\xe9\x8b\x0c\xc1\x34\x8b\xf2\xb3\xed\xb5\x71\x9b\xf1\xe3\xf4\xfd\x49\x89\xc7\x0c\x2d\x2e\x37\xf1\x23\x60\x24\xcf\xff\xae\xfa\x29\x00\x5c\x48\xf2\x8c\x17\xaf\xbd\x22\xb0\x4a\x6d\xfe\x9d\x0a\x3f\xb4\x3e\x94\xe9\xa3\xe6\x9c\x79\xbb\x9a\xf2\xc9\xfc\xcb\x2e\xaf\xe2\xf9\x2c\x75\xeb\x37\x4b\xa2\x67\xef\x62\x65\x23\x3d\x66\x3c\xdf\x1f\x01\xd8\xaa\xad\x92\xf9\x6b\xa1\xc4\x6f\xf5\xdc\x68\x02\xc5\x58\xb9\xb5\x19\x5f\xee\x8f\xf0\x1e\xfe\x06\x00\x00\xff\xff\x2e\x35\xad\x4a\x1e\x02\x00\x00")

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
	"type/course.graphql": type_course_graphql,
	"type/delegate.graphql": type_delegate_graphql,
	"type/lesson.graphql": type_lesson_graphql,
	"type/manager.graphql": type_manager_graphql,
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
		"course.graphql": &_bintree_t{type_course_graphql, map[string]*_bintree_t{
		}},
		"delegate.graphql": &_bintree_t{type_delegate_graphql, map[string]*_bintree_t{
		}},
		"lesson.graphql": &_bintree_t{type_lesson_graphql, map[string]*_bintree_t{
		}},
		"manager.graphql": &_bintree_t{type_manager_graphql, map[string]*_bintree_t{
		}},
		"users.graphql": &_bintree_t{type_users_graphql, map[string]*_bintree_t{
		}},
	}},
}}
