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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xcf\x6e\xe3\xb8\x0f\xc7\xef\x7e\x0a\x16\xbd\x74\x0e\xbf\xf9\xdd\x73\x9b\xc6\x18\x6c\x80\x16\xed\x4e\xec\x07\x60\x2d\xc6\x11\x46\x96\xbc\x92\xdc\xd9\x60\xd1\x77\x5f\x48\xb6\xf5\xcf\xd9\x5c\x3a\xb7\xe8\x63\x7e\x49\x8a\xa4\xa4\xdc\xc3\x30\x59\xb4\x5c\x49\x03\x1d\x4a\x18\x14\xe3\xa7\x0b\xd8\x33\x01\x7b\xfb\x0a\xf4\x37\x75\x93\x25\x06\x86\xfe\x9a\x48\x5a\x8e\x42\x5c\xbe\x56\x15\x97\xe3\x64\xa1\x1d\x85\x42\xf6\x9d\x0b\x7a\x26\x8b\xf0\x4f\x05\x70\xe2\x82\x9a\xcb\x48\x3b\x38\x5a\xcd\x65\x7f\x57\x01\x74\x4a\x5a\x92\xf6\x89\x64\x6f\xcf\x3b\x38\x48\x7b\x57\x7d\x54\x95\xbd\x8c\x94\xb8\xf8\x41\x66\xf4\x2e\x26\x2d\x52\xb5\x99\xba\x8e\x8c\x69\xd4\x4f\x92\x91\x7f\x6c\x73\x38\xce\x86\xde\xc7\x7f\x8a\x7c\xd0\xe7\x65\xcb\xde\x14\xd9\xc0\xe5\x93\xea\xb9\x7c\xf0\x1e\x77\xf0\x2d\x90\xbb\x2f\x3b\xf8\x36\xd9\xb3\xf7\x53\x01\x0c\x28\xb1\x27\x9d\x59\x3f\x27\xac\xb4\x67\x24\xa8\x47\x4b\x99\xa0\x4e\x61\xae\x70\xc5\xd2\x84\x96\x16\xa7\xab\x64\x9f\xc2\x83\x63\x4e\xb7\xac\x5d\xc9\x46\xb6\x15\xb5\x29\xbc\x22\x1a\xb5\x72\xdd\x3a\x0c\xd8\xd3\x5c\xc5\x1f\xae\xc9\xc6\x46\x07\x69\x7b\x9d\x38\xef\x56\x19\xf8\x35\x71\xb8\xf5\xb1\xb4\x27\xcf\xc1\x15\x68\x93\x78\x9d\xc2\x90\xf8\xa3\x52\x82\x50\xde\xc5\x22\xf9\x3e\xe5\x25\xf2\x28\x48\xfc\x2a\x64\x99\x99\xb7\x11\x6d\xcc\xe7\xa4\x32\xf3\x3a\xa2\x1b\x09\xad\x9d\xcd\x73\x5a\x69\x10\xe6\xd8\x55\x52\x49\x43\x21\xcf\xd2\x4b\x9b\xd1\xe0\x65\x05\x31\xfc\x5e\x0d\x23\xca\x4b\x1e\x7d\x81\x31\xf8\xbc\x0e\xd1\x0a\x51\x9b\xc2\x2b\xa2\xfb\xa5\x3a\x91\x64\xb1\xd7\x09\xaa\x00\xdc\xb9\xf7\xec\x6a\x26\xde\x60\x39\x4e\x57\xe7\xdb\x1b\x68\xea\x70\xb4\xdd\x19\xd3\x0b\x21\x2d\x3d\x00\x8e\xa3\x56\xef\x61\x1b\xd3\xc4\xd9\x0e\xda\xf6\x50\xa7\x79\xbb\x8b\x69\x12\x27\x2e\x5e\x49\x32\x2e\xfb\x17\xcd\x48\x3f\x74\x82\x93\xb4\x47\xea\x34\xd9\x10\x20\xf7\x6e\xf0\x9d\x5e\xa4\xe0\x92\xf6\x6a\xd2\x26\xf4\xe4\x58\xf0\xa4\x52\x6e\x09\xf7\xd0\xbc\xd4\x2f\x80\x42\xa8\x5f\x20\xb8\xfc\xb9\xf8\xda\x0b\x34\x46\x2b\x35\x6c\xdd\x15\x9f\xae\x78\x54\x4c\x01\x32\x06\x76\xb2\x4a\x83\x55\xd0\xf9\x4f\xfe\x86\x75\x3f\x1e\x51\x4a\xd2\x9f\x3a\xcf\xe3\xa4\xbb\x33\x9a\x65\x57\x66\x55\xbe\xe6\x38\xe4\x56\xf0\x64\x94\xd7\x29\x49\x77\x59\x27\xec\xc6\x29\x6a\xb0\xcf\x47\xb8\xc1\x3e\x98\x37\xd8\xc7\x99\x43\x4b\xbd\xd2\xe5\xc0\x2f\x34\x16\x70\x01\x31\xc2\x13\x19\xa3\x8a\x9b\x63\x66\x41\x34\x2f\xc3\x29\xc9\x15\x6d\xc2\xb6\x8a\x79\xeb\xb9\xa2\x4e\xd8\x8d\xad\x1f\x24\xe3\xef\x9c\x4d\x28\xf2\xe4\x22\x2f\x2e\x91\xf8\x21\xa9\xbd\x99\xde\x06\x6e\x9b\xa4\xf3\xc7\x40\x82\x3e\xa2\x57\xbc\xb8\x29\x80\xfb\x85\x19\xc0\xf0\x70\xfd\x9f\x87\x00\x06\x50\x9a\x5f\xa4\x8d\x9b\x3c\x04\x4b\xc6\xc6\x96\x25\xb1\xf6\x81\x14\xb9\x26\xb1\x42\x5d\x9b\x6c\x3c\x59\x29\x8c\x28\x0a\xe7\xf2\xa6\xc2\x3a\x90\xab\xa5\x9d\xd3\xfe\xd4\xb1\x98\x77\xf9\xa7\x53\xf1\x72\x70\x56\x5a\xec\x76\xc5\xe5\x8e\x4b\x27\x6d\x46\x8b\x9d\x6f\x9d\xcc\xbb\x2f\x9d\xd4\x19\xbd\x31\x60\xcf\x8a\x4d\xa2\x78\x9f\x66\x56\xa4\x3f\xc3\x32\xf9\x5c\xde\x26\xac\x48\xbc\x94\x2f\xef\x7c\x26\xaf\x13\x76\xeb\x3a\x70\xb7\x5d\x31\x5c\x0e\xc5\x2b\xc1\xad\xe2\x44\xa5\xe6\x6d\x44\x1b\x73\x7f\x8b\x1e\x79\x2f\xd1\x4e\xfa\x77\xfc\x0b\x6a\x32\x87\x57\x52\x88\xc1\xf2\x5c\xfc\xb3\x6a\xc8\xae\x2f\x7a\xf2\x12\x0c\xa3\x20\xff\x4f\xf5\x7f\xce\xc0\x00\x3f\x25\x87\x13\x46\x34\x86\x18\x28\x0d\x27\xe4\x82\x18\x20\x74\xab\x76\x79\x13\xe0\xe1\xbb\xd2\x73\x1c\xf3\xe5\xf7\x5d\xf0\xa1\x3b\x8f\x42\x15\xb7\xb5\x23\xb1\x9d\x42\xf5\xa1\x40\xa9\x6d\x1b\x48\x69\xfb\x26\x54\xff\x07\x21\xfb\xe4\x91\x8d\x21\x13\x67\xdb\xe8\xc9\xc7\x6b\x89\x3c\x2a\x76\xf9\x44\x1a\x1f\xd5\xbf\x01\x00\x00\xff\xff\xed\x85\x29\xf8\x60\x0d\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x56\xc1\x6e\xa4\x46\x10\xbd\xf3\x15\xcf\x9a\xcb\x58\xb2\xb2\x77\x6e\xb6\x57\x2b\x59\xca\x26\xbb\xf1\xf8\x14\x45\x51\x0f\x14\xd0\x71\xd3\x4d\xba\x8b\xd8\x24\xda\x7f\x8f\xaa\x69\x18\x98\x30\xa3\xec\x8d\xe1\xbd\xaa\x57\x55\xae\x57\x66\x87\x3f\x7b\xf2\x9a\x02\x42\xe3\x7a\x53\xc2\x3a\x46\xeb\x4a\x5d\x0d\xe0\x86\x50\x1e\x7f\x00\xbd\x53\xd1\x33\x95\xd0\x16\x9d\xf2\x5e\x19\x43\x26\xe3\xa1\x23\x7c\x51\x35\x3d\xd9\xca\xe1\x9f\x0c\x60\xc7\xca\xe4\x78\xb2\x7c\x83\x1d\x7e\xea\xdb\x23\x79\xb8\x0a\x9d\xaa\x29\x40\x55\x4c\x1e\xdc\xe8\x00\x67\x29\x03\x5c\x55\x05\xe2\x99\x7f\x68\x28\xbd\x92\x98\xc8\x93\x40\x54\xde\xb5\xb1\x94\xc0\xca\x73\x06\x18\xdd\xea\x75\x58\xab\xde\xa5\xfc\x37\x8b\x8e\x7c\x8c\xca\x80\x5a\xff\x45\x76\xa3\x1a\xcd\xd4\x86\x11\xcd\xbe\x65\x99\xb6\x5d\xcf\xb1\x8f\xd8\xc3\xa2\xa8\x95\xd4\x89\xfa\x59\x59\x55\x93\xff\xa4\x8d\xf4\x23\x31\x56\xb5\x94\xe3\x99\xbd\xb6\x75\x06\x50\xab\xb4\x59\xfc\xfe\xc3\x1d\x0f\x9a\xcd\x92\xd2\xf7\xba\xcc\xf1\xf2\xf2\xf4\x51\xc6\x46\x86\xba\xc6\xd9\x13\x61\xd6\xfa\x48\x86\x6a\xc5\xb4\x10\x5b\x87\x72\xf1\xbb\xfc\x9c\x13\x7f\x7f\x29\x57\xd4\x1f\x5d\xef\x03\x85\xcb\x9d\xaa\xa2\xa0\x10\x0e\x43\x47\x39\xee\xe7\xe7\x0c\x38\xaa\xe2\xb5\xf6\xae\xb7\xe5\x63\x43\xc5\x6b\x8e\x07\xe7\x0c\x29\x9b\x01\x9d\xd7\x05\xe5\xf8\x64\x9c\x92\x09\x2b\x63\xdc\x1b\x95\x07\xf7\xd0\x0f\x33\x0d\x3b\x3c\x37\xee\x0d\xce\x9a\x01\xc5\x58\x05\xb8\x51\x8c\xc1\xf5\x50\x9e\xa0\x7a\x6e\x9c\xd7\x7f\x53\x09\x76\x38\x3a\xf7\xba\x2c\xbb\xed\x94\x1d\x2e\xcd\x6c\xd5\xc3\x1c\xf4\x23\x85\xe0\xec\xff\x8a\x01\x58\xd5\x21\xc7\xaf\x02\xfe\x76\x4a\xf1\x3c\x18\xa3\x8e\xfd\x95\x79\xd1\x7b\x61\xfa\x92\x46\xad\xe5\x50\x12\xf0\xd9\x95\xbd\xfc\x71\xfe\x03\x1c\x28\xf0\xe9\x75\x54\x64\xf2\x95\x2a\xa2\x03\xb5\x55\x4c\xe5\xd7\x9e\xfc\x10\x55\xbb\x64\xca\x7c\xb6\xe7\xa9\xc8\x9f\x7d\x49\xfe\x61\xe4\xa9\x50\x90\x2d\xb5\xad\x97\x8a\x95\x26\x33\x6f\xd4\x8d\x04\x46\xa3\x9f\xb2\xeb\x98\x79\xc2\x01\x55\xb6\xda\xee\x4f\xe3\xba\xcd\x71\x2f\xaf\x26\x28\xec\xa5\x9e\xb1\x96\x09\x93\x67\x89\x6d\x47\x2f\xad\xa3\x93\xc1\x4e\xf0\x32\xc3\x1d\xaa\x38\xde\x7c\xed\xc3\x3b\xb8\xb1\xb1\x7c\xea\xf0\x94\xe8\xcb\x78\x0f\xca\xe4\xa5\x85\xd8\xcd\x6d\x3e\x5b\x6c\xc1\xd8\xd6\x5b\x7b\x71\x53\x70\xa2\x24\xc5\x62\x5c\xc4\x33\xc1\xb4\x9e\x33\xae\x2f\xe8\xad\xd6\x78\x53\x2e\x31\x92\x9a\x89\x5b\x75\x26\x36\xae\xda\x8c\x6e\x2b\x2d\x77\x7f\x53\x68\x24\x24\x9d\xa3\x71\xf5\x99\xca\x83\x71\x75\x42\xd6\x0a\x1b\xb9\x84\x3b\xcf\x27\x3a\xfb\x42\xf7\x8b\xdb\x73\xa1\x7b\x61\xac\x52\xed\xa5\x28\x39\xf8\x33\x2c\x46\x0e\x92\xef\x25\x90\xcf\x32\xc8\x7f\xbb\xc0\xfa\xd2\x24\xbe\x26\xf4\x8a\xec\x44\x49\xc2\x53\xbe\xb3\x89\x4c\xac\x78\x60\x03\x6f\xab\x89\xab\xaf\x28\x09\x9c\x54\x24\xc7\x99\x82\xa0\x62\x91\x78\x32\x2e\x38\x24\x62\xd7\x0c\x12\x09\x49\x63\xcc\x74\xa6\x32\x32\x64\x70\x81\x94\x2f\x9a\xe9\xc6\x6d\xea\xad\x0f\xe0\x6d\x8e\xe7\x55\xcc\x2f\x14\x7a\xc3\xd8\xe1\x5e\x6e\x7e\x48\x19\xb5\xad\xa7\x26\xee\xa6\x2d\x85\xb2\xe5\x38\xb7\x0c\xd8\xc5\x73\xbb\x67\x7a\xe7\xe9\xec\xdc\xe2\xc3\x87\x14\x2e\x1f\x16\xc6\x8c\x5f\x0c\x14\x99\x52\x6c\x41\x9e\x75\xa5\x0b\xc5\xf1\xfc\xed\xd9\xbd\xca\xb7\x40\x3a\x5a\xb2\x1d\x6b\x82\xdc\x22\x79\xae\x9d\x3f\xb7\xe3\x4a\x37\xc7\xe3\xc8\x9a\x4c\xb7\x43\xa0\xc2\xd9\x52\xf9\xe1\xf1\xfb\xe3\xbf\x65\xff\x06\x00\x00\xff\xff\xeb\xc5\x20\xba\x7f\x09\x00\x00")

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

var _type_blog_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xcb\x6e\xc3\x20\x10\xbc\xfb\x2b\xc8\x6f\x70\xcb\x43\x55\x2d\x55\x55\xda\xd4\xa7\x2a\x07\x6a\xd6\x84\x16\x83\x85\xf1\xc1\xaa\xf2\xef\xd5\xae\x13\x03\x8d\x15\xb5\x37\xef\xc0\x0c\x33\xbb\xeb\x30\x76\xc0\x36\xc6\xa9\xf5\x10\x4e\xce\xb3\xef\x82\xb1\x46\xfb\x3e\x3c\x8b\x16\x38\x3b\x04\xaf\xad\x5a\x15\x8c\x19\x71\x8b\x75\xde\x35\xda\xc0\x5e\xd7\xd5\xeb\xd3\xf5\xa0\x38\x17\xc5\xac\x4a\x7a\xb5\x07\x11\x40\xae\x43\xca\x1d\x3a\x99\x83\x88\x0d\x5a\x72\x56\x55\xe5\x0e\x6f\x04\x1d\x4c\xf6\xdc\x87\x93\x63\x5a\xd7\x22\x80\x72\x7e\xe4\x6c\x7b\xf9\x42\xf4\x04\x42\x82\x2f\x5b\xa1\x20\x71\x85\x27\x82\x22\xf2\x24\x2e\xa9\x1a\xa7\x36\x4e\x8e\xc4\xe8\x39\x7b\xc7\x63\x2a\x56\xc7\x2c\x0b\x61\x14\xe8\xb3\x77\xb6\xdc\x65\x69\xbc\x89\x65\x4a\xda\x23\x47\xb7\x9d\x81\x16\x6c\xe8\xd9\x5e\x28\x6d\x31\xf8\xcb\x00\x7e\x24\x35\x90\xf1\xdd\x23\x76\x55\x28\x28\x6d\xe3\x38\x5e\xa6\x2f\x54\xd4\xb6\x1b\x42\xf4\x51\x52\xb9\x6c\x26\xb8\x2f\xb0\x99\x9d\x89\xbc\xa5\x39\x90\xc4\xcc\xfe\x73\x93\x71\x2a\x71\x36\x49\x93\xdf\xd2\xd7\xe6\x2e\xc7\xeb\x17\xcd\xdb\xf6\x92\x89\xa9\xc7\x93\xbf\x8a\x56\xe2\xff\x11\x17\x15\x66\xf2\x9d\xa5\xfa\x15\x77\x31\xed\xfd\xb0\x59\xb2\x25\xff\xcb\xf9\x1e\xa3\x62\x34\x8a\xbf\xd2\x61\xa8\xa1\xef\x39\xab\x3a\xe3\x84\x7c\x98\x10\x84\xae\x9b\x9a\x8e\xe1\xfc\x13\x00\x00\xff\xff\x1d\xb2\xaa\x5e\xbd\x03\x00\x00")

func type_blog_graphql() ([]byte, error) {
	return bindata_read(
		_type_blog_graphql,
		"type/blog.graphql",
	)
}

var _type_certificate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\xcf\x41\x0e\x82\x30\x10\x05\xd0\x7d\x4f\x31\x5c\xc3\x1d\x62\x4c\x48\x08\x0b\xc1\x03\x4c\xea\x40\x26\x42\x4b\x86\x69\x42\x63\xbc\xbb\x71\xa1\x54\x61\xfb\x5f\xdb\xff\xab\x71\x22\x28\x48\x94\x3b\xb6\xa8\x54\xba\xce\xc3\xc3\x00\x58\x1f\x64\xa6\x96\x75\xa0\x03\x34\x2a\xec\xfa\xcc\x00\xd0\x32\xb1\xc4\x13\xea\x4f\x6a\xfd\x38\x0d\xa4\xec\xdd\x9e\xa0\x8b\x35\x8e\xdf\xd8\x00\x28\xde\x49\xce\x2c\xb3\xa6\x90\x7d\xa4\xc2\x2d\xd8\x75\xe3\xd1\xdf\xe2\xf5\x52\x25\xef\x09\xf5\x61\xc0\x77\x7f\x4b\x8b\xa6\xd7\x8a\x3c\xaf\x7d\x72\x72\xf3\xa1\xd2\xcd\x2a\xc1\xaa\x97\xff\xc6\x55\x8a\xb2\xde\x87\x86\x7b\x87\x1a\x84\xd2\x39\x4f\xf3\x0a\x00\x00\xff\xff\x56\x36\x21\x56\x55\x01\x00\x00")

func type_certificate_graphql() ([]byte, error) {
	return bindata_read(
		_type_certificate_graphql,
		"type/certificate.graphql",
	)
}

var _type_company_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xcf\x6a\xc3\x30\x0c\xc6\xef\x7e\x0a\xf5\xb6\x41\x2f\xdb\xd1\xb7\x36\x65\x50\xd8\x5f\x46\x4f\x63\x07\x11\xab\xc6\x90\xd8\xc6\x56\x06\xa1\xf4\xdd\x87\xed\xa4\x84\x26\xdd\x71\x37\xe9\xf3\x67\xf9\x27\xc9\xdc\x7b\x82\xca\xb5\x1e\x6d\x0f\x27\x01\x80\xde\x07\xf7\x43\x4a\xc2\xd6\xb9\x86\xd0\x0a\x00\x13\x2b\x67\x39\x60\xcd\x17\x75\x25\x00\xba\xce\x28\x09\x87\xc3\x7e\x97\xb2\x3a\x10\x32\xa9\x0d\x4b\xf8\xe4\x60\xac\x16\x00\x16\x5b\x1a\xd3\xe4\x69\xd1\xa2\xa6\x10\xef\x3c\x6a\x92\xf0\x8e\x9a\xd6\x70\x34\x0d\x53\x90\xf0\x52\x0e\x9f\x72\xba\x06\x17\x14\x85\x6d\x2f\xe1\xad\x04\xf7\x17\x47\xba\x96\xaa\x29\x6a\x48\x23\xd3\x72\xb9\xdd\x70\xfa\x47\xbd\xd1\x32\x16\x44\xa5\x02\xc5\x28\x61\x53\x82\x95\x38\x0b\x91\x27\x34\x08\x65\x42\x25\x7e\x36\x96\x1e\xa6\xcd\x4d\xf4\xc7\xa9\x5e\xbb\xce\x72\x3f\x55\xbc\x8b\x5c\x39\x45\x33\x57\x98\xd8\xc6\xa7\x87\xe5\x24\x46\x30\xad\x6f\xa8\x25\xcb\x31\x35\x6b\x6c\x1a\xf8\x47\x47\xa1\xac\x8e\x94\xa6\x28\xe1\x6b\xb8\xf1\x9d\x5e\x42\x4d\x7b\x7b\x74\x65\x38\x29\x4a\x75\x8d\xf5\x1d\x43\x95\x17\x36\x98\xf7\x59\x3a\x65\x90\x2c\xbc\x5e\xad\xee\xff\xba\x2e\x74\x07\xaf\x96\xe8\xae\xfe\xdc\x9c\xf5\xc6\x17\x5e\xc2\xbf\x41\x3f\x83\x5f\x60\x9f\xa3\x8b\xb3\xf8\x0d\x00\x00\xff\xff\x1e\x8a\xae\xbc\x4c\x03\x00\x00")

func type_company_graphql() ([]byte, error) {
	return bindata_read(
		_type_company_graphql,
		"type/company.graphql",
	)
}

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4b\x6f\x23\x37\x0c\xbe\xfb\x57\x30\xf1\x21\x2d\xb0\x39\x74\xf3\x40\xeb\x5b\xd6\x71\x51\x03\x59\x6c\xea\x07\xda\x22\xc8\x81\xd1\xd0\x63\x61\x35\xd2\x54\xa2\x92\x0c\x16\xfb\xdf\x0b\x49\xe3\x19\xd9\xb1\xb7\x49\xf7\x52\x0c\xe0\x79\x88\x1f\x49\x91\x1f\x49\x99\xb4\xaf\xe0\x4a\x08\x72\x6e\xd1\xd4\x04\x5f\x06\x00\x96\x1c\x5b\x29\x98\x8a\x01\x80\xa9\x49\x0f\xbe\x0e\x06\x51\x70\xce\xd6\x0b\xf6\x96\x26\x8a\x2a\xd2\x1c\xc5\x2b\x53\x78\x45\x03\x00\x26\xc7\x03\x00\x45\xce\x99\x1e\x33\x36\xde\x3a\xea\x94\x1b\xad\xa4\x0e\xd2\x42\xa1\x73\xd6\x98\x2a\x48\x4a\xcd\x64\x57\x28\x08\xe6\x8d\x52\xf8\xe0\xdd\x94\xa9\x8a\x00\x8d\x15\x8d\x82\x65\xa9\xcb\xa3\x01\x80\xf7\xb2\x18\xc1\x72\x39\xbd\x0e\x6f\xc2\x54\xb5\x22\xa6\x11\x7c\x30\x46\x11\x46\xbb\x1c\x8c\xdd\x44\x37\x40\x86\xf5\xe0\xab\xfb\x7e\xd5\x00\x55\x7d\xb6\x9c\xdd\x6c\x30\x71\xcf\xcf\x9c\xab\x60\x2c\xdd\x08\xee\x16\x58\x1e\xdd\x77\xae\xa4\x10\x44\x93\xc1\xc0\x54\x73\x14\x6d\x6a\x1a\x65\xe1\x39\xda\xe3\xd1\x03\x8a\xcf\xa5\x35\x5e\x17\xe3\x35\x89\xcf\xb9\x2b\xd8\x25\x6d\x94\x25\x30\x80\x6a\x2b\x05\x8d\xe0\x57\x65\x90\xd3\x46\x94\xb1\x99\xcf\xf4\x2c\xc8\xd6\x9c\x7d\x91\x9a\xad\x29\xbc\x60\x69\x74\xf6\x79\x6d\x9e\x16\x66\xdc\x85\x21\xfb\xee\xad\xcb\x57\xa2\xa9\x01\xc0\xd3\x1a\xf9\x2f\xe3\x6f\x08\xad\x1e\xc1\x5d\xbb\x8f\xfb\xc8\xa9\xbf\xbd\xb4\x29\x11\xdb\x2b\xae\x26\x21\x57\x52\x2c\xc8\x56\x2e\x33\x22\x90\xa9\x34\xb6\x19\xc1\xb8\x7d\x0a\x9b\x56\xca\x3c\x51\xb1\x30\x1f\x7c\x93\x07\xc3\xb5\xa9\x0d\xaa\xb3\x2c\x47\x03\x0f\xa8\x35\xd9\x69\x85\x25\x6d\xe7\x8e\x9e\x6b\x69\xc9\x4d\xf5\x47\xa3\x79\xed\xba\xc4\xc4\xef\x18\x62\xb1\x30\x13\x5d\xc4\xd5\xce\xd8\xd1\x4e\x56\x6f\xb1\xa4\x9c\x64\xb7\x58\x4a\x8d\x4c\xc5\xef\x9e\x6c\x13\x73\x4e\x45\x49\xc1\xb3\x04\x08\x2e\xd5\x58\xd2\x54\xaf\xcc\x28\x88\xc7\xa7\x54\x04\xb5\x67\x98\xe3\x23\x8d\x37\xb5\x91\x20\xd3\xb8\x90\xd1\x67\x87\x29\x21\x5a\x64\x39\x44\x11\x99\x12\x25\x02\x8b\xb3\x28\x86\xd7\xee\xe3\xab\x09\xf0\x66\xfa\xed\xb0\x6f\x0f\xf9\xda\x02\x89\x45\xf6\xad\xf4\x67\x59\x9b\xfb\xa8\x3f\x57\xe2\xd9\xd8\x7e\x4b\x30\x4c\x37\xb3\x02\x5e\x13\x2c\xc2\x2a\xfc\xd0\x35\x18\x10\xa9\xfe\x8c\x56\xcd\x8f\xa1\x88\xf1\xf9\x16\x2d\x4b\x21\x6b\x8c\x74\x4c\x11\x75\x8c\x96\xaf\x31\xf0\x79\x21\xab\xb0\x15\xd2\xc5\xd6\xbb\x32\x02\x77\x22\xb4\x8f\x43\xff\x46\xa1\xed\x5c\x7f\x8a\x0d\xf1\xf5\x89\xde\x97\xd0\x43\xd9\x1f\x82\x23\x61\x74\x81\xb6\x19\xbf\x82\x08\x30\x84\xf9\xda\x58\x86\x82\x9c\xb0\xb2\x0e\xfe\x87\xa0\xa6\xf8\x1d\xa0\x09\x0c\xe1\xe4\xea\xc1\x78\x06\x5e\x4b\xd7\xca\x9e\x04\xc3\x41\xea\x50\x17\x09\xa8\xdf\xcc\x13\xb0\xe9\xfa\xec\x61\xfc\xbe\x6e\x03\x43\x98\x38\x96\x55\x28\xb5\x24\x72\x48\xd7\xe1\xbe\x14\x9c\xf8\x63\x8d\x0c\x8d\xf1\x27\x4a\x81\x0a\xeb\x27\xf0\xe0\x95\x22\x86\xda\x48\xcd\xee\x70\xef\x0a\xe8\x59\xb6\xf4\x12\xf8\x7f\x2d\x1f\xb7\x19\xe3\x5d\x53\xea\x9a\xe5\xf7\xf1\xb9\x57\x16\x49\x9c\x26\xdc\xee\xa1\x61\x77\xd2\xf6\x68\x4b\x81\xc0\x58\xf6\x65\xb0\x3b\x0f\xb7\xc2\xd1\xb7\xe2\x05\x96\x51\xbc\x57\xfb\x7a\xec\xa6\x30\xfe\x8b\x82\xdc\xef\x8d\x9e\xb7\x38\x9f\xf0\xb7\xde\x8a\x35\xba\xb6\x07\xb8\x5e\x41\xe2\x6f\xc8\x77\x98\x4c\xf7\x31\x70\x8e\x6c\x4f\x80\x23\x18\xc2\x8d\x74\x1c\x8a\xb4\x20\x45\x25\x32\xb9\xe8\xbe\x8b\x49\x63\x8b\x53\xfd\x68\xa4\xa0\x49\x85\x52\x65\x14\x08\xfc\xab\x99\x8a\x96\x37\xdd\x70\x7b\x41\xda\xb1\xd1\x2b\x69\xab\x97\xe7\xab\x1d\xaf\x67\xe4\x6a\xa3\xdb\x53\x0e\x5b\xd4\x0e\x63\x01\x8f\x77\x4f\x51\xc1\x67\xb6\x9e\x40\xa6\x6e\x1d\x4a\x16\x75\x03\xd2\x81\xd1\x80\x20\x8c\x66\x8b\x82\x01\x75\x01\xbc\xf6\x0e\x0a\x43\x4e\x9f\x30\x68\xa2\x22\x14\x79\x8d\x0d\xf8\x7a\x65\x4d\xdb\xb6\xad\xac\x69\xac\x24\x69\x9e\x93\xb0\x94\xb7\xb3\xe9\x6a\x9f\x2f\xc1\xd6\x0a\x95\xa3\x77\xa9\x4f\x48\x07\xa5\x7c\x24\x0d\xce\xc4\x00\x83\x40\x1d\xac\x74\x7b\x9d\x13\x5a\xb1\xde\x9c\x2c\x66\xe4\xbc\xe2\x57\x0e\xfd\x17\xc7\x91\x6f\xcd\xfe\x6b\x0a\xee\x1d\x98\x05\x2f\x29\xfb\x96\xb3\x47\x0b\x39\xec\xc1\x10\xee\x60\xf2\xe7\xd5\xc7\xdb\x9b\x09\x8c\x3f\x2d\x67\xf3\x09\xcc\x17\xb3\xe5\x78\xb1\x9c\x4d\x06\x43\x00\xf8\x92\xca\xf9\x38\x9d\xf7\x8f\xdf\xb5\xc5\x72\x7c\x76\x7e\x76\x7e\x76\x9a\xff\x9e\xc7\xdb\xf1\xd7\x77\x5b\xb8\xf4\xbf\xa0\xc7\xfd\xf4\x3e\x5e\xa7\xe9\xf6\xfe\x34\x7b\xdb\x85\x86\x7f\x16\x3d\xf0\xe7\x5f\xd2\x75\x7a\x71\x19\xaf\xd3\x8b\xcb\xee\xe1\xe2\xf2\x22\x81\xef\x07\xff\x04\x00\x00\xff\xff\x6b\xb7\x15\x5f\xdb\x0c\x00\x00")

func type_course_graphql() ([]byte, error) {
	return bindata_read(
		_type_course_graphql,
		"type/course.graphql",
	)
}

var _type_delegate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xc1\x4e\xf3\x30\x10\x84\xef\x7e\x8a\xed\xed\xff\x5f\xc1\xb7\x92\x5e\x22\x01\x2a\x90\x9c\x50\x85\x96\x7a\x6b\x0c\x8e\x6d\xd9\x0e\x55\x54\xf5\xdd\x91\x93\xa6\x49\xd3\x20\xe0\xe6\x8e\x66\x5d\x7f\xb3\x93\xd8\x38\x82\x15\x69\x92\x18\x09\x0e\x0c\xa0\xae\x95\xe0\x50\x96\xf9\x6a\xc1\x00\xb6\x9e\x30\x92\x58\x46\x0e\x4f\xd1\x2b\x23\x19\x40\x51\x64\x2f\xf9\xaa\x17\x92\x8b\x2a\x54\x7a\xe4\xd8\x29\x1f\xe2\x3d\x56\x34\x36\x69\xbc\xd6\x22\x69\x72\x6f\xd6\xd0\x68\x38\xf9\x6e\xad\x54\x66\x6c\x7c\xb7\xaf\x85\x8a\xfa\x62\x78\x6b\x2b\x87\xa6\xe1\x90\x75\x87\xa4\x39\x6f\x77\x4a\x53\x5e\xa1\xa4\xd2\x8f\xdf\x84\xdb\xa8\x3e\x55\x6c\xfe\x39\x94\xc4\x61\x8d\x92\xfe\x73\x58\x9e\xd4\xf4\x93\x01\x54\x4d\x66\x6b\x1f\x28\x70\x78\xbe\x3b\x9d\x17\x1b\x76\x64\x4c\x19\x57\x47\xc8\xda\x38\xfa\xbc\xf2\x56\x3b\x0c\x4f\x49\xa9\x75\xd9\xfd\x25\x84\x69\x7a\x73\xac\x73\x41\x49\x32\xe4\x31\xd2\x1a\x43\xd8\x5b\x2f\x38\xdc\x58\xab\x09\xcd\x34\x07\xa7\x2d\x8a\xc2\x7e\xd0\x39\xd2\x81\xa8\x74\x62\x96\x68\x52\x83\xdf\xf0\xcd\xe0\xcd\xd0\xcc\x00\xcf\xc1\xfd\x04\x00\x60\x68\x3f\x90\x0f\x58\x6d\xa3\x2f\xf7\xf4\x48\xc1\x59\x13\xba\x7e\x8b\x93\xc8\xcf\xb5\x5f\x8c\xb2\x14\xdf\x5e\xd9\xbb\x53\x53\x40\x55\x4e\x53\x45\x26\x86\x54\x24\x65\xd2\xe4\x43\x4d\xbe\x69\xff\x82\x84\x6c\x1b\xd4\x8f\x6c\x12\x0f\x4a\xca\xcd\xce\x76\xcd\x4b\xa7\x61\x07\xbd\xaf\x6d\x7d\x7b\xc3\xf5\x27\xe6\x26\xef\x4a\x9a\xb1\x89\x6c\x58\xfb\x91\x7d\x05\x00\x00\xff\xff\xaf\xa0\xf8\xb6\xcf\x03\x00\x00")

func type_delegate_graphql() ([]byte, error) {
	return bindata_read(
		_type_delegate_graphql,
		"type/delegate.graphql",
	)
}

var _type_individual_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\x3d\xca\xc2\x40\x10\x87\xf1\x7e\x4e\x31\xef\x35\xd2\xbe\x55\x1a\x0b\x3f\x0e\x30\xb2\x7f\x75\x64\xb3\xbb\xcc\xcc\x2a\x22\xb9\xbb\x90\x42\x03\x69\x7f\xf0\xc0\xa3\xa5\xf5\xe0\x7f\x83\x04\xc6\x92\xf4\xa1\xa9\x4b\x1e\x17\x7d\x13\xf3\x45\xcd\x63\x27\x13\x06\x3e\x84\x69\xb9\xfe\x11\x73\x96\xad\xdd\xeb\xf9\xa8\x91\xbf\x46\xcc\x81\x8c\x76\xab\x65\x6d\x98\x44\xf3\xba\x6b\xe2\xfe\xac\x96\x7e\x36\x13\xc5\xab\x61\x33\xb5\x87\xb7\x5a\x1c\xcb\x57\x77\xd8\xc0\x27\x87\xd1\x4c\x9f\x00\x00\x00\xff\xff\xb1\x3b\x04\x79\xc5\x00\x00\x00")

func type_individual_graphql() ([]byte, error) {
	return bindata_read(
		_type_individual_graphql,
		"type/individual.graphql",
	)
}

var _type_lesson_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\xc1\x0a\x83\x30\x10\x04\xd0\x7b\xbe\x62\xfc\x8d\x5c\xeb\x45\xe8\xc1\x52\x3c\x95\x1e\x02\x99\x86\x80\xc6\x60\x36\x50\x29\xfe\x7b\x89\x82\x05\x5b\x7a\x5b\x76\x67\xde\xca\x1c\x89\x33\x53\x1a\x43\x6b\x1c\xe1\x87\xd8\x73\x60\x90\x84\xd6\x38\x1f\x8c\xd0\x5e\x32\xa7\x19\x2f\x05\xd0\x3a\x26\x8d\xdb\x56\xb8\x2b\x20\x1a\xc7\x26\x3c\x46\x5d\xe2\xeb\xa4\x16\xa5\x7c\x88\x59\x70\x9a\x68\x84\x5b\xb6\x59\x37\xc5\x08\x66\xa0\xc6\x55\x26\x1f\x5c\xa5\x00\x31\xae\x90\x5d\xd7\xd4\x55\x11\x85\x4f\xf9\xdc\x77\xac\x8b\xf6\x07\x96\xb3\xb7\x1a\x6b\xf7\x40\x1f\xa4\xaf\x47\x3b\x5c\xb3\xe7\x7f\x78\x51\xef\x00\x00\x00\xff\xff\x87\x0c\x30\xb9\x26\x01\x00\x00")

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

var _type_module_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\xcb\x4e\x83\x40\x14\x86\xf7\xf3\x14\x87\xd7\x60\xa7\x56\x13\x12\x4d\xaa\x14\x5d\x98\x2e\xa6\xcc\x91\x4c\x9c\x0b\x99\x4b\x13\x62\x78\x77\x33\x33\x2d\x0c\xa4\x9a\xba\x72\x55\xe0\x5c\xfa\xff\xdf\x0f\xe3\x86\x1e\xe1\x49\x33\x2f\x10\xb8\xec\x05\x4a\x54\xce\x42\x3d\x08\x41\x0f\xde\x56\x0e\x25\x7c\x11\x00\xef\x39\x2b\xa1\x69\xaa\x4d\x41\x00\x14\x95\x58\x42\xed\x0c\x57\x5d\xb8\x3f\x50\xa5\xd0\x54\x92\x76\xd8\xbc\x3c\x9e\x2b\x04\x80\xa1\x6d\x0d\xef\x1d\xd7\x2a\xef\x77\x86\xaa\x54\xc8\x9f\x1e\x35\x6f\x51\x1f\xd1\x2c\x77\x1c\x39\x43\x5d\xc2\x6b\xf8\x21\x00\xf6\x24\xad\x84\xf7\x5c\x65\xb1\x27\x00\xad\x0e\x16\x1c\x96\x70\xab\xb5\x40\xaa\xc8\x48\x08\x2a\x2f\xd3\xf4\x2e\x98\x0d\x6e\xde\xaa\x7a\x57\xdd\x84\x62\xf4\x1f\x8b\xb1\x10\x6e\xcb\xb9\x39\xa8\xf2\x46\xcc\x22\x47\x42\xb8\xea\xbd\x4b\x2d\x55\xbc\xbc\x72\x2e\xca\x48\xa4\xef\x13\xe6\x34\x89\xd6\x11\x00\x81\xd6\x6a\x35\xef\x4f\x8d\x13\xfe\xb4\x7f\x31\x5d\xac\x52\x99\x46\xef\x0c\x52\x87\xa7\x05\x93\xc2\x75\x64\x8e\x76\x01\x61\x9c\xdd\xff\x35\xa9\x2c\xef\xda\xb7\x2d\x5a\xbb\xd3\x9f\xa8\xf2\xd0\xce\x59\xfe\x54\xcf\x42\x8d\x22\x97\xc9\xce\xf6\x8b\xfd\x14\x53\x6e\x6c\x4b\x07\xa1\x29\x8b\xd6\x64\x7c\x72\xc6\x33\x83\x68\x7a\x76\x01\xc4\xf2\x4d\x5e\x63\xc8\x31\x5d\xa6\x72\x11\xca\xbf\x31\xc9\x3d\x5e\xc3\x64\x83\xe1\xfb\xf8\x8d\xc9\x48\xb2\x33\x61\x4b\xbb\xc5\xb9\xb0\xa5\x1d\x57\xd4\x21\x7b\xf6\x68\x86\x38\x8b\xac\xc3\x59\x5e\x60\xd8\xd3\x0e\x2b\xf5\xa1\xcb\xd0\x1e\xaf\xd6\xef\xf5\x03\x17\x0e\xcd\xea\xaf\xaf\xa3\x3f\x92\xef\x00\x00\x00\xff\xff\x8c\x12\x92\x85\xb3\x04\x00\x00")

func type_module_graphql() ([]byte, error) {
	return bindata_read(
		_type_module_graphql,
		"type/module.graphql",
	)
}

var _type_question_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xdb\x6e\xa3\x30\x10\x86\xef\xfd\x14\x83\xf6\x2d\xb8\x4b\x08\x9b\x45\xca\xee\x26\x0d\x48\x95\x2a\x14\x8d\xc2\x04\x59\x05\x9b\x62\xa3\x16\x45\xbc\x7b\x85\x09\x84\x53\xda\x26\x77\xcc\x68\x0e\xff\x7c\xbf\xd1\x65\x46\xb0\x2b\x48\x69\x2e\x05\x9c\x19\x40\x51\xf0\xc8\x86\x20\xf0\x56\x16\x03\xd0\xf4\xa1\x6d\xd8\xeb\x9c\x8b\xb8\x8e\x73\x14\x91\x4c\xb9\xa2\x85\x50\xef\x94\x2b\x1b\x96\x52\x26\x84\x82\x01\xbc\x5d\xc6\xf8\x65\x46\x76\x37\xb4\x8e\xea\x4e\x6c\x1b\x5e\x9a\x56\x2b\x64\x15\x63\x83\xfd\x5b\x8c\x09\x78\x9a\x25\x94\x92\xd0\x0a\xb6\x18\x73\x81\x9a\xa2\x5d\x41\x79\x69\xd4\x51\x14\x53\x3d\xa3\x6d\x09\x19\x40\x86\x31\x79\xe2\x24\xed\xba\xc1\x7c\x75\x93\x9b\x55\x33\x77\x71\xe5\xc8\x3c\xa7\xa3\xee\x1f\xd0\x3f\xb6\xae\x49\x31\xa6\xe0\x69\xd3\xa5\x2a\xc6\x48\x14\xe9\xe0\x34\x33\x7b\xef\xfd\x5b\x6f\xdc\x83\xf3\xe7\xbf\xe7\xb8\x5d\x59\xb3\xbc\x2b\xf2\xdd\x67\x9f\x01\x78\x7f\x17\x6b\xf7\x12\x1e\x9a\xa0\x62\x8c\x8b\xac\xd0\xe0\xe4\x84\x9a\xda\xf1\x9e\xc9\x9d\xef\x71\xc1\xba\xc3\x86\x66\xd9\x12\x15\x3f\x36\x63\xcc\x3e\x2b\x34\xb6\x63\x5c\x97\x18\x5a\xa1\x35\x16\x38\xee\x31\x1a\xa7\x44\x2d\xf8\x05\x11\x9d\xb0\x48\xb4\x02\x2d\xe1\x84\x89\xa2\x1b\x98\x7d\xf9\x4a\xa2\x97\xc4\x8e\x9d\xdd\xe3\x68\x75\xc6\x0e\x49\x6d\xb1\x4c\x24\x46\x46\x47\x7b\xfe\xf5\xf4\xab\xfc\x20\x8b\x66\xf9\xde\x7e\xf3\x0f\x3f\xf9\x01\xea\x66\xef\x14\xf5\x98\xf4\x58\xe9\x2c\xe8\xab\xd8\x2f\x38\xcd\xf9\xf1\xb0\x1d\x2d\xf4\x21\xbe\x9f\x42\x6f\x33\xbf\x79\xa2\x27\xff\xe2\x74\xbf\x3f\x0f\x64\x45\x09\x7d\x67\x5d\xf5\x19\x00\x00\xff\xff\x8f\x47\x4b\x1b\xce\x04\x00\x00")

func type_question_graphql() ([]byte, error) {
	return bindata_read(
		_type_question_graphql,
		"type/question.graphql",
	)
}

var _type_test_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\x56\xc8\x5f\xe8\x96\x26\x08\xe0\x53\x1d\xd8\x3e\x14\x81\x0f\xb4\x39\x56\x88\x52\xa4\xca\x5d\x36\x10\x8a\xfc\x7b\x41\xea\x69\xb7\xea\x03\xed\xc9\x24\xbd\x33\x9a\x9d\x1d\x52\xba\x16\xb4\x07\x0b\x99\xa6\xb5\x68\xe0\x84\x69\xd7\x59\xab\x4e\x91\x37\x82\x86\xbe\x15\x44\x4e\x35\xa8\x68\x27\xc1\xb8\xba\x2c\x88\x62\x34\xba\xa2\xc3\x61\xf3\x98\x76\x67\x9f\xa0\x82\x8a\x3e\x78\x6f\xa1\x5c\x41\x24\xaa\xe6\x8a\x5e\xf6\xaa\x2e\x8f\x05\x91\x12\x41\xd3\x0a\xdf\x5b\xeb\xdf\xa0\x2b\xda\x38\x29\x88\x5a\xc5\xbc\x45\x38\xc3\x89\xaa\x51\xd1\x93\xf5\x2a\x9d\x7f\x89\x60\x31\xde\xf1\xde\xdf\x3b\x7e\x43\x18\x01\x41\x39\xed\x1b\xc3\xe8\x8f\x79\xf9\xc9\x09\x54\xd1\xcb\xf3\xb0\x2e\x8f\xc5\x7b\x51\x4c\x4d\x6e\x55\x8d\x65\xa3\x5b\x55\x1b\xa7\x04\xfa\x39\x22\x74\xb9\x55\xe8\x1a\x59\x39\x58\x8e\x59\x62\x8d\x8d\xbb\xf8\x2a\x15\xe7\x55\x62\x34\xae\x8d\x92\x29\x9f\x8c\x15\x84\x0c\x9d\x5d\xb9\xb1\x6c\xd2\xf0\x10\xa0\x04\xbd\x92\xce\x7a\xa5\x33\x4e\xc0\x52\x65\xb2\x99\x7a\xae\xdc\xe4\xfd\xcf\xa6\x30\x78\x9c\xc7\xb0\x66\x72\xb9\xe6\x72\xb9\x6a\x73\xf9\x2b\x9f\x4b\xba\xa3\x8f\x5f\x11\x82\xd1\x60\x6a\x11\x26\x92\x09\xa3\xd2\xee\x66\x1e\xbd\xc6\x0c\x0e\x1a\x01\x9a\xac\x61\x21\x7f\x21\x79\xc5\x5c\x49\xc6\x91\xbc\x1a\xce\x96\xcc\x66\x8c\xd3\xec\xc5\x64\x2f\x46\x48\x22\x1e\x92\x48\x77\xf9\xf7\x96\x94\x4e\x30\xae\x26\x95\xb1\xd0\xc9\xa7\xbc\x5c\x47\x32\x2c\xce\x02\x3d\x14\xce\x3a\x76\xf1\xd4\x18\xb9\x1e\xca\xd9\xc7\xc0\x48\x54\xc9\x3a\xba\xa3\x99\xa7\xff\x2b\x2f\x23\x23\x50\xea\x4b\x7d\x4e\x62\xd2\x91\xe4\x4b\xe7\x86\xf9\xaf\xab\xc9\x75\x9d\x8f\xa4\x02\x06\x45\x29\x52\x63\x1b\xcb\xb4\xf7\xfe\xf4\x3e\x7f\xf2\x31\x8c\x25\x24\x7e\xa2\x9a\xb2\x38\x37\xb3\xcc\x62\xaf\x79\x27\x4a\x22\x57\xf4\xb0\xd8\x8d\x49\x4a\xc1\x5a\x84\x61\x1f\x22\xc8\x5c\xe6\x2e\xfb\x9a\xab\xef\xf5\xee\x1d\x5a\xfd\x43\xa4\xaf\x1f\x92\xab\x3b\xf3\x67\xf9\xfe\xaf\x8f\xc8\xbf\x64\xfb\xaf\xa3\x9d\x87\x30\x7b\xf2\xbb\x07\xe1\x11\xe9\x89\x5d\x77\xef\xbd\xf8\x1e\x00\x00\xff\xff\xc8\x72\x0d\x85\xcb\x05\x00\x00")

func type_test_graphql() ([]byte, error) {
	return bindata_read(
		_type_test_graphql,
		"type/test.graphql",
	)
}

var _type_tutor_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xd0\xc1\xca\x02\x21\x10\x07\xf0\xbb\x4f\x31\xdf\x6b\xec\xf5\x8b\x60\xa1\x53\xae\x0f\x20\x3a\x2d\x43\x9b\x8a\xce\x1c\x62\xe9\xdd\xc3\x05\xb3\xa0\x82\x8e\xf3\x57\x7f\xcc\x5f\xbe\x26\x84\x49\x38\x66\x58\x15\x80\x08\xf9\x01\x8c\x19\x77\x7f\x0a\x20\xd8\x0b\x0e\xa0\x39\x53\x98\xeb\xec\x28\x3c\x8f\x85\xe6\x60\x59\x32\x9a\xe3\xa1\xe7\x37\xa5\x28\x24\x61\xf8\xcf\x68\x19\x37\x7b\xdc\x82\xf5\x07\x72\x8a\x67\x7c\x9c\x74\xd3\x24\xdf\x4c\xdd\xae\x76\xfc\x44\x0b\x6a\x71\x0e\x4b\x19\xc0\xa4\x25\x5a\xbf\xef\x51\xf5\xb9\xbe\xac\xf5\x5a\xc9\x77\x72\x07\x3f\x7f\xc7\xeb\xea\x5f\x36\xbf\x07\x00\x00\xff\xff\x8f\xe4\x13\xbb\x61\x01\x00\x00")

func type_tutor_graphql() ([]byte, error) {
	return bindata_read(
		_type_tutor_graphql,
		"type/tutor.graphql",
	)
}

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\xc1\x6e\xd4\x30\x10\xbd\xfb\x2b\xde\x2a\x17\x90\x50\x3f\x20\xb7\x55\x11\x52\x25\x40\xa0\x76\x4f\xa8\x87\xd9\xf5\x24\x3b\xd4\xb1\x83\x3d\x2e\x8a\x50\xff\x1d\xd9\xd9\xa4\x81\x45\xe2\x66\x4f\xfc\xde\xbc\x37\x6f\x14\xf6\x79\xc0\x21\x71\x7c\x98\x46\xc6\x2f\x03\x0c\xe4\xa9\xe7\x68\x00\xf1\x56\x9e\xc5\x66\x72\x06\xb0\xec\xb8\x27\x65\xf3\x62\x4c\x53\x11\x90\x04\x82\x9e\xc5\x43\x0b\x58\xcf\xa4\x38\x05\xe7\x48\x39\x41\x7c\x17\xe2\x40\x2a\xc1\x83\x8e\x21\x2b\xf4\xcc\x0b\xf9\xbb\x57\xba\x06\xe4\xed\xa6\x15\x72\xa1\x2e\x84\xe9\x06\xf7\x61\x60\x74\xc2\xce\x26\x50\x64\xf8\xa0\xe8\x83\xf8\x1e\x1a\x70\x64\xd0\x33\x89\xa3\xa3\x63\x58\x1e\xd9\x5b\xf1\xbd\x69\x10\x7c\xed\x55\x45\x85\x6e\x26\x9c\x42\xae\x0c\x91\x7f\x64\x4e\x5a\x28\x2c\x29\xa1\x0b\xf1\xc6\xd4\x97\xd5\x52\x19\x40\xb9\xb5\xeb\x4c\x76\x06\xe0\x81\xc4\xb5\xb8\xd7\x58\x1a\x00\x9d\xc4\xa4\x9f\x69\xe0\xa5\x56\x1e\x39\xba\xae\x29\x3b\x1e\xcf\xc1\xf3\x06\xfc\x3d\x1c\x1f\x44\xdd\xb6\x54\xa0\x1f\x43\x2f\x7e\x8b\x3d\x85\x61\x24\x3f\xb5\xb8\x9d\x0f\x06\x18\x63\xe8\xc4\xf1\xdd\x40\x3d\x1f\xe2\x56\x11\x9d\x54\x9e\x45\xa7\x37\x23\xf5\xdc\xe2\x0b\xf5\xfc\xb6\xc5\xfe\x52\x2d\xd7\x92\xec\x74\x1b\x72\x4c\x9c\x5a\x7c\xfb\x74\x39\xef\x1e\x0d\xd0\x80\xac\x8d\x9c\x52\x8b\xfd\x7c\x28\x29\xd7\xd5\x58\x28\xd6\xf5\xf0\xfc\x73\x46\x2e\x5d\x49\xd9\x5e\xe4\x3a\x56\xb6\xeb\xd7\x8e\xc4\xad\xd7\x17\x33\x4f\x79\x2b\x09\x52\x20\x03\x7b\x4d\x45\xb1\xf8\x42\xf5\x35\x73\x9c\x6a\x27\xb6\x7d\x95\xba\x40\x8a\xd2\x62\xef\xce\x77\x61\xb6\x58\x4e\x57\xcc\x15\x9b\xb3\xd8\x16\x87\xc3\xdd\xfb\xdd\x9a\xe8\xd6\x4a\x1d\x70\xe4\xd2\x70\xaf\x7f\x4e\xbd\xc8\x2d\x43\x5f\x64\x37\x78\xa5\xe7\xb9\xfc\x5f\xf1\x0d\xfe\x92\x7f\x01\x3e\xd6\x2f\xff\x30\xd1\xa0\x76\x6a\xb0\xbf\x28\xc0\x29\xc7\xc8\x5e\xdd\x84\x23\xd7\x85\xa7\x27\xf6\x38\x4e\xa0\xba\xd0\xb3\xe7\x25\xc5\xea\x39\x29\x69\x4e\x8b\xf4\xfb\x7a\xbb\xb6\x54\x2a\x83\xf8\xac\x9c\x1e\x22\x9d\x9e\xd8\xb6\xf8\xe0\x02\xe9\x6e\xcd\x7c\x4b\x50\x99\x47\x4a\xa9\x86\x3c\x47\x5a\x7f\x0d\x4b\xde\xe6\xc5\xfc\x0e\x00\x00\xff\xff\xca\xd8\xcd\xe9\x44\x04\x00\x00")

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
	"type/blog.graphql": type_blog_graphql,
	"type/certificate.graphql": type_certificate_graphql,
	"type/company.graphql": type_company_graphql,
	"type/course.graphql": type_course_graphql,
	"type/delegate.graphql": type_delegate_graphql,
	"type/individual.graphql": type_individual_graphql,
	"type/lesson.graphql": type_lesson_graphql,
	"type/manager.graphql": type_manager_graphql,
	"type/module.graphql": type_module_graphql,
	"type/question.graphql": type_question_graphql,
	"type/test.graphql": type_test_graphql,
	"type/tutor.graphql": type_tutor_graphql,
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
		"blog.graphql": &_bintree_t{type_blog_graphql, map[string]*_bintree_t{
		}},
		"certificate.graphql": &_bintree_t{type_certificate_graphql, map[string]*_bintree_t{
		}},
		"company.graphql": &_bintree_t{type_company_graphql, map[string]*_bintree_t{
		}},
		"course.graphql": &_bintree_t{type_course_graphql, map[string]*_bintree_t{
		}},
		"delegate.graphql": &_bintree_t{type_delegate_graphql, map[string]*_bintree_t{
		}},
		"individual.graphql": &_bintree_t{type_individual_graphql, map[string]*_bintree_t{
		}},
		"lesson.graphql": &_bintree_t{type_lesson_graphql, map[string]*_bintree_t{
		}},
		"manager.graphql": &_bintree_t{type_manager_graphql, map[string]*_bintree_t{
		}},
		"module.graphql": &_bintree_t{type_module_graphql, map[string]*_bintree_t{
		}},
		"question.graphql": &_bintree_t{type_question_graphql, map[string]*_bintree_t{
		}},
		"test.graphql": &_bintree_t{type_test_graphql, map[string]*_bintree_t{
		}},
		"tutor.graphql": &_bintree_t{type_tutor_graphql, map[string]*_bintree_t{
		}},
		"users.graphql": &_bintree_t{type_users_graphql, map[string]*_bintree_t{
		}},
	}},
}}
