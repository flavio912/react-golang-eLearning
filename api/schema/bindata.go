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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xcd\x6e\xe3\x36\x10\xc7\xef\x7a\x8a\x09\x72\xc9\x1e\xba\xbd\xfb\xb6\xb1\xb1\xa8\x81\x04\x71\xd7\xd6\x03\x4c\xc4\xb1\x4c\x2c\x45\xaa\xfc\xc8\xd6\x28\xf2\xee\x05\x29\x89\x5f\x72\x73\x49\x8f\xfc\x69\xfe\x33\x7f\xce\x90\xb4\xef\x61\x70\x16\x2d\x57\xd2\x40\x87\x12\x06\xc5\xf8\xf9\x0a\xf6\x42\xc0\x5e\xbf\x02\xfd\x4d\x9d\xb3\xc4\xc0\xd0\x5f\x8e\xa4\xe5\x28\xc4\xf5\x6b\xd3\x70\x39\x3a\x0b\xed\x28\x14\xb2\xef\x5c\xd0\x33\x59\x84\x7f\x1a\x80\x33\x17\x74\xba\x8e\xb4\x81\xa3\xd5\x5c\xf6\x77\x0d\x40\xa7\xa4\x25\x69\x9f\x48\xf6\xf6\xb2\x81\xbd\xb4\x77\xcd\x7b\xd3\xd8\xeb\x48\x59\x8a\x1f\x64\xc6\x90\xc2\x69\x91\xab\x8d\xeb\x3a\x32\xe6\xa4\x7e\x92\x4c\xfc\x7d\xed\xe1\x38\x05\x86\x1c\xff\x29\x0a\x45\x9f\xe7\x2d\x87\x50\x64\x03\x97\x4f\xaa\xe7\xf2\x21\x64\xdc\xc0\xb7\x48\xee\xbe\x6c\xe0\x9b\xb3\x97\x90\xa7\x01\x18\x50\x62\x4f\xba\x88\x7e\xce\x58\x1d\xcf\x48\x50\x8f\x96\x0a\xc1\x2e\x87\xa5\xc2\x37\x4b\x13\x5a\x9a\x93\x2e\x92\x6d\x0e\xf7\x9e\x79\xdd\xbc\xf6\x2d\x1b\xd9\x5a\xd4\xe6\xf0\x86\x68\xd4\xca\x4f\x6b\x3f\x60\x4f\x53\x17\x7f\xf8\x21\x1b\x9b\x12\xe4\xe3\xf5\xe2\x72\x5a\x75\xe1\x43\x96\x70\x9d\x63\x1e\x4f\xe9\xc1\x37\x68\x65\x7c\x97\xc3\x68\xfc\x51\x29\x41\x28\xef\x52\x93\xc2\x9c\xca\x16\x05\x14\x25\x61\x15\x5d\x16\xe1\x6d\x42\xab\xf0\xc9\x54\x11\xbe\x4b\xe8\x03\x43\xcb\x64\x4b\x4f\x0b\x8d\xc2\x12\xfb\x4e\x2a\x69\x28\xfa\xac\xb3\xb4\x05\x8d\x59\x16\x90\xca\x6f\xd5\x30\xa2\xbc\x96\xd5\x67\x98\x8a\x4f\xeb\x58\xad\x12\xb5\x39\xbc\x21\xba\x9f\xbb\x93\x48\x51\x7b\x39\x41\x0d\x80\xbf\xf7\x81\xdd\x74\x12\x02\xe6\xeb\x74\xf3\x7c\x87\x00\x4d\x1d\x8e\xb6\xbb\x60\xfe\x20\xe4\xad\x07\xc0\x71\xd4\xea\x2d\x6e\xc3\x39\xce\x36\xd0\xb6\xfb\x5d\xee\xdb\x3f\x4c\x4e\x9c\xb9\x38\x90\x64\x5c\xf6\x2f\x9a\x91\x7e\xe8\x04\x27\x69\x8f\xd4\x69\xb2\xb1\x40\x99\xdd\xe0\x1b\xbd\x48\xc1\x25\x6d\x95\xd3\x26\xce\xe4\x58\xf1\xac\x53\x7e\x09\xf7\x70\x7a\xd9\xbd\x00\x0a\xa1\x7e\x81\xe0\xf2\xe7\x9c\x6b\x2b\xd0\x18\xad\xd4\xb0\x4e\x57\x7d\xba\x91\x51\x31\x05\xc8\x18\x58\x67\x95\x06\xab\xa0\x0b\x9f\xfc\x55\x76\xba\xbb\xa0\x99\xcd\x98\x25\xed\xa1\xc4\x31\x65\xc5\xe3\x09\x8c\xd3\x3c\x61\x5f\x9e\xa2\x13\xf6\x51\x7d\xc2\x3e\x8d\x1d\x2d\xf5\x4a\xd7\x67\x6e\xa6\x69\x0f\x33\x48\x15\x9e\xc8\x18\x55\x5d\xde\x89\x45\xd1\xb4\x8c\x07\xb5\x54\xb4\x19\x5b\x2b\xa6\x33\x5a\x2a\x76\x19\xfb\xe0\x0e\xef\x25\xe3\x6f\x9c\x39\x14\xa5\xb9\xc4\xab\x7b\x9c\x3e\x64\x37\xd9\xb8\xd7\x81\xdb\x53\xf6\x98\x1e\x23\x89\xfa\x84\x0e\x78\xf5\xcf\x24\xdc\xcf\xcc\x00\xc6\xdf\x8e\xdf\x79\x2c\x60\x00\xa5\xf9\x45\xda\xf8\xe1\x23\x58\x32\x36\x8d\x2c\xab\xb5\x8d\xa4\xf2\x9a\xd5\x8a\x7d\x3d\x15\x2f\x3e\xab\x85\x09\x2d\x42\x7f\xe9\x82\x8d\x4f\xfd\x72\x4c\xae\xff\xf4\x2a\x5e\x1f\x84\x85\x56\xee\x17\x5c\xef\xa0\x4e\xd2\x16\xb4\xda\xc9\x3a\xc9\xf2\xa4\x2d\x5f\xb2\x5f\x61\xc5\x9c\xa8\x5e\xf3\x89\x55\xce\x26\x58\xfb\x2a\xe5\x6d\xc6\x2a\x4f\xa5\x3c\x58\x9a\x87\xea\x6f\x7a\x58\xcf\xb3\x0a\xeb\x00\x0c\xd9\xe5\x17\x20\x7b\x39\x86\x51\x50\xf8\x67\xf3\x9b\x0f\x30\xc0\xcf\xd9\x49\x82\x11\x8d\x21\x06\x4a\xc3\x19\xb9\x20\x06\x08\xdd\xa2\x9d\x1f\x13\x78\xf8\xae\x34\x84\x3a\xe6\xcb\xff\xf7\xb2\x00\xc4\xa6\x3e\x0a\x55\x3d\x2e\x9e\xa4\x1b\x29\x54\x1f\x3b\x98\xc7\xb6\x91\xd4\xb1\xaf\x42\xf5\x7f\x10\xb2\x4f\x9e\xc8\x54\x32\x4b\xb6\xae\x9e\x7d\xbc\x65\xe4\x51\xb1\xeb\x27\x6c\xbc\x37\xff\x06\x00\x00\xff\xff\xd0\xe8\xf9\x88\x92\x0b\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x96\xc1\x6e\xe3\x46\x0c\x86\xef\x7a\x8a\x3f\xc8\xc5\x01\x82\xee\x5d\xb7\x64\x17\x5b\x04\xe8\xb6\x59\xc4\x39\x15\x45\x41\x4b\x94\x34\xcd\x68\x46\x9d\x19\x6d\xa2\x76\xf3\xee\x05\x67\x24\x59\x72\x65\xa3\xbd\xd9\xe6\x4f\x7e\x24\x4d\x52\xba\xc6\x9f\x3d\x3b\xc5\x1e\xbe\xb1\xbd\x2e\x61\x6c\x40\x6b\x4b\x55\x0d\x08\x0d\xa3\x3c\xfc\x00\x7e\xe3\xa2\x0f\x5c\x42\x19\x74\xe4\x1c\x69\xcd\x3a\x0b\x43\xc7\x78\xa4\x9a\x1f\x4c\x65\xf1\x77\x06\x04\x1b\x48\xe7\x78\x30\xe1\x0a\xd7\xf8\xb9\x6f\x0f\xec\x60\x2b\x74\x54\xb3\x07\x55\x81\x1d\x42\xa3\x3c\xac\xe1\x0c\xb0\x55\xe5\x39\xcc\xfa\x7d\xc3\xe3\x4f\xe2\x13\x75\xe2\x88\xca\xd9\x36\xa6\xe2\x03\xb9\x90\x01\x5a\xb5\x6a\xed\xd6\xd2\x9b\xa4\xff\x6a\xd0\xb1\x8b\x5e\x19\x50\xab\x6f\x6c\x36\xb2\x51\x81\x5b\x9f\xac\xd9\x7b\x96\x29\xd3\xf5\x21\xd6\x11\x6b\x58\x24\xb5\x42\x1d\xa5\x5f\xc8\x50\xcd\xee\xb3\xd2\x52\x8f\xf8\x18\x6a\x39\xc7\x53\x70\xca\xd4\x19\xc0\x2d\x29\xbd\xf8\xfe\x87\x3d\xec\x55\xd0\x4b\x49\xdf\xab\x32\xc7\xf3\xf3\xc3\x27\x69\x1b\x6b\xee\x1a\x6b\x8e\x82\x99\xf5\x89\x35\xd7\x14\x78\x01\x5b\xbb\x86\xe2\x77\xf9\x3a\x07\xfe\xff\xa9\x5c\xa0\x7f\xb4\xbd\xf3\xec\xcf\x57\x4a\x45\xc1\xde\xef\x87\x8e\x73\xdc\xcd\x9f\x33\xe0\x40\xc5\x4b\xed\x6c\x6f\xca\x8f\x0d\x17\x2f\x39\xee\xad\xd5\x4c\x26\x03\x3a\xa7\x0a\xce\xf1\x59\x5b\x92\x0e\x93\xd6\xf6\x95\xcb\xbd\xbd\xef\x87\x59\x86\x6b\x3c\x35\xf6\x15\xd6\xe8\x01\x45\xca\x02\xa1\xa1\x80\xc1\xf6\x20\xc7\xa0\x3e\x34\xd6\xa9\xbf\xb8\x44\xb0\x38\x58\xfb\xb2\x4c\xbb\xed\xc8\x0c\xe7\x7a\xb6\xaa\x61\x76\xfa\x89\xbd\xb7\xe6\x3f\xf9\x00\x81\x6a\x9f\xe3\x57\x31\xfe\x96\x42\x04\x76\x15\x15\x71\x1f\x94\xa1\xc0\xe5\xd7\x9e\xdd\x10\x03\x75\xe3\x8a\xe4\xf3\xb2\x1c\xa9\xbf\xb8\x92\xdd\x7d\xd2\x91\x2f\xd8\x94\xca\xd4\xcb\x6e\x55\x8a\xf5\xfc\xff\x5e\x89\x63\x5c\xbb\x63\x74\x15\x23\x4f\x76\x80\xca\x56\x99\xdd\x31\xf9\x9b\x1c\x77\xf2\xd3\x64\xf2\x3b\xc9\x27\xe5\x32\xd9\xe4\xb3\xf8\xb6\x69\xb2\xd7\xde\xe3\xb8\x1f\xcd\xcb\x08\xb7\xa8\x62\xc7\xf2\xf5\x56\xdc\xc2\xa6\xc2\xf2\xa9\xc2\x63\xa0\xc7\xb4\x9d\xe5\x38\xd9\x0b\xd8\xd5\x4d\x3e\x0f\xfc\x42\xb1\xcd\x5b\x6f\xc6\x26\x70\x92\x8c\xc4\x22\x8d\xc5\x09\x70\x1c\x96\xd9\xae\xce\xf0\x56\x43\xb5\x89\x1b\x15\x23\x4d\xc7\x79\x3a\x81\xa5\x21\x9b\xad\xdb\xa4\xe5\x24\x6e\x82\x92\x60\xe4\x1c\xb4\xad\x4f\x28\xf7\xda\xd6\xa3\x65\x4d\xd8\x88\x25\xda\xb9\x3f\x71\xcf\xce\x54\xbf\xb8\x04\x67\xaa\x17\xc5\x2a\xd4\x4e\x92\x92\xf3\x3b\x9b\x65\xad\xbc\xc4\x7b\xf6\xec\xb2\x0c\xf2\xec\xf1\x41\x9d\xeb\xc4\xd7\xd1\x7a\x01\x3b\x49\x46\xf0\x14\xef\xa4\x23\x93\x2a\x9e\x3b\x1f\xb6\x69\x7b\xf6\xe1\x02\x49\xcc\x23\x45\x62\x9c\x10\xc4\x2a\x2b\x62\xcb\x5e\x9f\xe9\xe1\x97\x68\xbb\xb4\x20\x51\x30\x32\x52\xa4\x13\x4a\x52\x48\xe3\xae\xe1\x99\x5c\xd1\x3c\x0d\x5a\xd3\xa1\xdf\x26\x4e\xc6\xc4\xbc\xc1\x87\x0f\xb8\x93\x73\xeb\x47\x67\x65\xea\x29\xe3\xdb\x69\x24\x41\xa6\x4c\x4d\x8a\x94\x82\x02\xd7\x56\xde\x0f\x76\x9a\xbf\xb1\xce\xf1\xe8\x54\x4b\x6e\xc0\xf7\xef\x78\xe2\xc2\x9a\x92\xdc\x70\x8b\xc0\x6f\x61\x3a\x43\x11\xf4\x23\x07\xb9\xed\xe9\x51\xce\xe3\x44\x2c\xc2\xc5\xe8\x72\x47\x77\xff\x72\x4d\xc9\xc9\x1b\xc3\xd1\x5f\x94\xd9\x7b\xf6\x4f\x00\x00\x00\xff\xff\x25\x62\xa8\xdf\xae\x08\x00\x00")

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

var _type_company_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xcf\x6a\xc3\x30\x0c\xc6\xef\x7e\x0a\xf5\xb6\x41\x2f\xdb\xd1\xb7\x36\x65\x50\xd8\x5f\x46\x4f\x63\x07\x11\xab\xc6\x90\xd8\xc6\x56\x06\xa1\xf4\xdd\x87\xed\xa4\x84\x26\xdd\x71\x37\xe9\xf3\x67\xf9\x27\xc9\xdc\x7b\x82\xca\xb5\x1e\x6d\x0f\x27\x01\x80\xde\x07\xf7\x43\x4a\xc2\xd6\xb9\x86\xd0\x0a\x00\x13\x2b\x67\x39\x60\xcd\x17\x75\x25\x00\xba\xce\x28\x09\x87\xc3\x7e\x97\xb2\x3a\x10\x32\xa9\x0d\x4b\xf8\xe4\x60\xac\x16\x00\x16\x5b\x1a\xd3\xe4\x69\xd1\xa2\xa6\x10\xef\x3c\x6a\x92\xf0\x8e\x9a\xd6\x70\x34\x0d\x53\x90\xf0\x52\x0e\x9f\x72\xba\x06\x17\x14\x85\x6d\x2f\xe1\xad\x04\xf7\x17\x47\xba\x96\xaa\x29\x6a\x48\x23\xd3\x72\xb9\xdd\x70\xfa\x47\xbd\xd1\x32\x16\x44\xa5\x02\xc5\x28\x61\x53\x82\x95\x38\x0b\x91\x27\x34\x08\x65\x42\x25\x7e\x36\x96\x1e\xa6\xcd\x4d\xf4\xc7\xa9\x5e\xbb\xce\x72\x3f\x55\xbc\x8b\x5c\x39\x45\x33\x57\x98\xd8\xc6\xa7\x87\xe5\x24\x46\x30\xad\x6f\xa8\x25\xcb\x31\x35\x6b\x6c\x1a\xf8\x47\x47\xa1\xac\x8e\x94\xa6\x28\xe1\x6b\xb8\xf1\x9d\x5e\x42\x4d\x7b\x7b\x74\x65\x38\x29\x4a\x75\x8d\xf5\x1d\x43\x95\x17\x36\x98\xf7\x59\x3a\x65\x90\x2c\xbc\x5e\xad\xee\xff\xba\x2e\x74\x07\xaf\x96\xe8\xae\xfe\xdc\x9c\xf5\xc6\x17\x5e\xc2\xbf\x41\x3f\x83\x5f\x60\x9f\xa3\x8b\xb3\xf8\x0d\x00\x00\xff\xff\x1e\x8a\xae\xbc\x4c\x03\x00\x00")

func type_company_graphql() ([]byte, error) {
	return bindata_read(
		_type_company_graphql,
		"type/company.graphql",
	)
}

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4f\x6f\xe3\xb6\x13\xbd\xeb\x53\x4c\xe2\x43\x7e\x3f\x60\x73\xe8\x66\x13\xb4\xbe\x65\x15\x17\x35\x90\x62\x53\x5b\x46\x5b\x04\x39\x4c\xa8\xb1\x4c\x2c\x45\xaa\xe4\x30\x89\xb0\xd8\xef\x5e\x90\x54\x2c\xda\x89\x8b\x4d\xf7\x52\x18\xb0\xfe\x70\xde\xcc\x70\xde\x9b\xa1\x48\xfb\x16\x2e\x85\x20\xe7\xaa\xbe\x23\xf8\x52\x00\x58\x72\x6c\xa5\x60\xaa\x0b\x00\xd3\x91\x2e\xbe\x16\x45\x34\x5c\xb2\xf5\x82\xbd\xa5\x99\xa2\x96\x34\x47\xf3\xd6\xd4\x5e\x51\x01\xc0\xe4\xb8\x00\x50\xe4\x9c\x19\x31\xa5\xf1\xd6\xd1\xd6\xb9\xd1\x4a\xea\x60\x2d\x14\x3a\x67\x8d\x69\x83\xa5\xd4\x4c\x76\x8d\x82\x60\xd9\x2b\x85\xf7\xde\xcd\x99\xda\x08\xd0\xd8\xd2\x34\x44\x96\xba\x39\x2a\x00\xbc\x97\xf5\x14\x56\xab\xf9\x55\x78\x12\xa6\xed\x14\x31\x4d\xe1\xa3\x31\x8a\x30\xc6\xe5\x10\xec\x3a\xa6\x01\x32\xac\x87\x5c\xdd\xf7\xbb\x06\x68\xbb\xb3\xd5\xe2\xfa\x19\x13\xf7\xfc\xc4\xb9\x0b\xc6\xc6\x4d\xe1\xb6\xc2\xe6\xe8\x6e\x9b\x4a\x2a\x41\x0c\x19\x02\xcc\x35\x47\xd3\xbe\xa3\x69\x56\x9e\xa3\x57\x32\xba\x47\xf1\xb9\xb1\xc6\xeb\xba\xdc\x90\xf8\x9c\xa7\x82\x5b\xd2\xa6\x19\x81\x01\xd4\x59\x29\x68\x0a\x3f\x2b\x83\x9c\x36\xa2\x8c\xcd\x72\xa6\x27\x41\xb6\xe3\xec\x8d\xd4\x6c\x4d\xed\x05\x4b\xa3\xb3\xd7\x1b\xf3\x58\x99\x72\x5b\x86\xec\xbd\xb7\x2e\x5f\x89\xa1\x0a\x80\xc7\x0d\xf2\x9f\xc6\x5f\x13\x5a\x3d\x85\xdb\x61\x1f\x77\x51\x53\x7f\x79\x69\x13\x11\xbb\x2b\xae\x23\x21\xd7\x52\x54\x64\x5b\x97\x05\x11\xc8\xd4\x18\xdb\x4f\xa1\x1c\xee\xc2\xa6\x95\x32\x8f\x54\x57\xe6\xa3\xef\xf3\x62\xb8\x81\xda\xe0\x3a\x63\xf9\x05\x07\x37\xd8\x50\x2e\x89\x1b\x6c\xa4\x46\xa6\xfa\x37\x4f\xb6\x8f\x0c\x51\xdd\x50\xf0\x93\x00\x21\xc3\x0e\x1b\x9a\xeb\xb5\x99\x06\xf3\x78\x97\x24\xdb\x79\x86\x25\x3e\x50\xf9\xac\xe4\x04\x99\xc7\x85\x8c\xec\x3d\x5e\xb3\xbd\x05\xa9\x25\xc1\xbd\x85\x96\x37\x8b\x62\x4f\x13\xaf\x48\x62\x90\x6d\x94\xfe\x3f\x91\x72\x8f\x5a\x93\x9d\xb7\xd8\xd0\xd2\x47\xff\xb9\x13\xcf\xc6\x8e\x5b\x82\x49\xba\x98\x35\xf0\x86\xa0\x0a\xab\xf0\xbf\x6d\xdb\x83\x48\x5d\x61\xb4\xea\xff\x1f\x5a\x0b\x9f\x6e\xd0\xb2\x14\xb2\xc3\x28\x92\x54\x39\xc7\x68\xf9\x0a\x83\xca\x2a\xd9\x86\xad\x90\xae\x77\x9e\x95\x11\xb8\x53\xa1\x1d\x72\x3e\xc5\x79\xf3\x7d\xcc\x4c\xc0\x91\x30\xba\x46\xdb\x97\xdf\xc0\x1c\x4c\x60\xb9\x31\x96\xa1\x26\x27\xac\xec\x42\x72\xa1\x0a\x69\xc3\x07\x78\x85\x09\x9c\x5c\xde\x1b\xcf\xc0\x1b\xe9\x06\xdb\x93\x10\x38\x58\x1d\x6a\xc6\x80\xfa\xc5\x3c\x02\x9b\xed\xb8\x3a\x8c\x7f\xad\x69\x61\x02\x33\xc7\xb2\x0d\x3d\x90\x4c\x0e\xf9\x3a\xdc\xde\x21\x89\xdf\x37\xc8\xd0\x1b\x7f\xa2\x14\xa8\xb0\x7e\x02\xf7\x5e\x29\x62\xe8\x8c\xd4\xec\x0e\x8f\x80\x80\x5e\x64\x4b\x2f\x81\xff\x55\xbd\xbb\xe7\xd3\x70\x3b\x2d\xc6\x99\x93\x04\x38\xbe\x8d\xaa\x4b\x13\x7f\xff\x10\xdd\x3f\x79\x46\xb4\x25\x64\xaa\xb0\x19\x75\xbb\x7f\x3e\xec\xec\xeb\x68\x3b\xec\x2a\x6c\xa2\xf9\xe8\xf6\xdb\xb1\xcf\x0a\xff\x37\x0e\xf2\xbc\x9f\xfd\xbc\x25\xf9\x84\xbf\xf1\x56\x6c\xd0\x0d\x4d\xeb\x46\x07\x49\x88\x81\xb8\x70\x84\xde\xc5\xc2\x39\xb2\x23\x93\x47\x30\x81\x6b\xe9\x38\x74\x5b\x4d\x8a\x1a\x64\x72\x31\x7d\x17\xdb\x94\x2d\xce\xf5\x83\x91\x82\x66\x2d\x4a\x95\x71\x19\x84\xd4\x31\xd5\x83\x00\x06\x85\xbd\x72\x04\x97\x46\xaf\xa5\x6d\x5f\x7e\x6f\xec\x65\xbd\x20\xd7\x19\x3d\x9c\xfa\x6c\x51\x3b\x8c\x9d\x58\xee\x7f\x55\x84\x9c\xd9\x7a\x02\x99\xe6\x64\xe8\x3d\xd4\x3d\x48\x07\x46\x03\x82\x30\x9a\x2d\x0a\x06\xd4\x35\xf0\xc6\x3b\xa8\x0d\x39\x7d\xc2\xa0\x89\xea\xd0\xad\x1d\xf6\xe0\xbb\xb5\x35\xc3\xc0\xb4\xb2\xa3\x52\x49\xd2\xbc\x24\x61\x29\x9f\x4b\xf3\xf5\x6b\xb9\x84\x58\x6b\x54\x8e\xde\xa5\x86\x97\x0e\x1a\xf9\x40\x1a\x9c\x89\x05\x06\x81\x3a\x44\x19\x29\xba\xa2\x80\x3b\x30\x55\x23\x95\x13\xb8\x85\xd9\x1f\x97\xbf\xde\x5c\xcf\xa0\xfc\xb4\x5a\x2c\x67\xb0\xac\x16\xab\xb2\x5a\x2d\x66\xc5\x04\x00\xbe\xa4\x6e\x38\x4e\x9f\x8f\xc7\xef\x06\xad\x1d\x9f\x7d\x38\xfb\x70\x76\x9a\xff\x7f\x88\x97\xe3\xaf\xef\x76\x70\xe9\x33\x73\xc4\xfd\xf0\x3e\xfe\x4e\xd3\xe5\xfd\x69\xf6\xb4\x0f\x0d\x1f\xaa\x23\xf0\xc7\x9f\xd2\xef\xf4\xfc\x22\xfe\x4e\xcf\x2f\xb6\x37\xe7\x17\xe7\x09\x7c\x57\xfc\x1d\x00\x00\xff\xff\x37\x20\x35\x52\x2a\x0b\x00\x00")

func type_course_graphql() ([]byte, error) {
	return bindata_read(
		_type_course_graphql,
		"type/course.graphql",
	)
}

var _type_delegate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\xc1\x4e\xf3\x30\x10\x84\xef\x7e\x8a\xed\xed\xff\x5f\xc1\xb7\xd2\x5e\x22\x21\x54\x20\x39\x21\x84\x96\x7a\x6b\x0c\x8e\x6d\xd9\x0e\x55\x54\xf5\xdd\x91\x93\xa6\x49\x53\x23\xe0\x96\x8e\x66\x6d\x7f\xb3\xd3\xd8\x3a\x82\x35\x69\x92\x18\x09\x0e\x0c\xa0\x69\x94\xe0\x50\x55\xc5\x7a\xc1\x00\xb6\x9e\x30\x92\x58\x46\x0e\x8f\xd1\x2b\x23\x19\x40\x59\xae\x5e\x8a\xf5\x20\x24\x17\xd5\xa8\xf4\xc4\xb1\x53\x3e\xc4\x3b\xac\x69\x6a\xd2\x78\xad\x45\xd2\xe4\xde\xac\xa1\xc9\x70\xf2\xdd\x5a\xa9\xcc\xd4\xf8\x6e\x5f\x4b\x15\xf5\xc5\xf0\xd6\xd6\x0e\x4d\xcb\x61\xd5\x7f\x24\xcd\x79\xbb\x53\x9a\x8a\x1a\x25\x55\x7e\xfa\x26\xdc\x46\xf5\xa9\x62\xfb\xcf\xa1\x24\x0e\x1b\x94\xf4\x9f\xc3\xf2\xa4\xa6\x9f\xec\xc8\x98\x32\xae\x89\xb0\xea\xa8\x87\x58\x8a\x4e\x3b\x8c\x37\xa6\x70\xfa\x88\xfe\xc2\x3a\x0f\x29\x87\x94\xcb\x43\x92\x21\x8f\x91\x36\x18\xc2\xde\x7a\xc1\xe1\xc6\x5a\x4d\x68\xe6\xb8\x4e\x5b\x14\xa5\xfd\xa0\x73\x72\x23\x51\xe5\x44\x96\x68\xb6\xed\xdf\xf0\x65\xf0\x32\x34\x19\xe0\x1c\xdc\x4f\x00\x00\x86\xf6\x23\xf9\x88\xd5\x15\xf7\x72\x4f\x0f\x14\x9c\x35\xa1\xaf\xb1\x38\x89\xfc\xdc\xee\xc5\x24\x4b\xf1\xed\x91\x83\x3b\x15\x02\x54\xed\x34\xd5\x64\x62\x48\x7d\x51\x26\x4d\xde\x37\xe4\xdb\xee\x0a\x12\x92\x02\x87\xa7\x61\xe4\x39\xf1\xa0\xa4\xc2\xec\x6c\x5f\xb0\xf4\x35\xee\x60\xf0\x75\xe5\xee\x4e\xb8\xfe\x27\xb9\xd9\xbb\x92\x66\x6c\x22\x1b\xd7\x7e\x64\x5f\x01\x00\x00\xff\xff\x97\x3b\xd0\xa3\xb6\x03\x00\x00")

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

var _type_test_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\x56\xf0\x5f\xe8\x96\x36\x08\xe0\x53\x1d\xd8\x3e\x14\x81\x0f\xb4\x38\x51\x88\x52\xa4\xca\x5d\x36\x30\x8a\xfc\x7b\x41\x4a\x16\x6d\xb7\xea\x03\xcd\xc9\x24\xbd\x33\x9a\x9d\x1d\x52\x4e\x03\x68\x07\x16\x32\xfd\x60\xd1\xc3\x09\xd3\xf6\x64\xad\x3a\x46\x5e\x0b\x7a\xfa\x5e\x11\x39\xd5\xa3\xa1\xad\x04\xe3\xba\xba\x22\x8a\xd1\xe8\x86\xf6\xfb\xf5\x7d\xda\xb5\x3e\x41\x05\x0d\x7d\xf0\xde\x42\xb9\x8a\x48\x54\xc7\x0d\x3d\xed\x54\x57\x1f\x2a\x22\x25\x82\x7e\x10\xbe\xb3\xd6\xbf\x42\x37\xb4\x76\x52\x11\x0d\x8a\x79\x83\xd0\xc2\x89\xea\xd0\xd0\x83\xf5\x2a\x9d\x7f\x8d\x60\x31\xde\xf1\xce\xdf\x39\x7e\x45\x38\x03\x82\x72\xda\xf7\x86\x31\x1e\xf3\xe5\x27\x67\x50\x43\x4f\x8f\xd3\xba\x3e\x54\x6f\x55\x35\x37\xb9\x51\x1d\x2e\x1b\xdd\xa8\xce\x38\x25\xd0\x8f\x11\xe1\x94\x5b\x85\xee\x90\x95\x83\xe5\x90\x25\x76\x58\xbb\x67\xdf\xa4\xe2\xbc\x4a\x8c\xc6\x0d\x51\x32\xe5\x83\xb1\x82\x90\xa1\xc5\x95\x1b\xcb\x66\x0d\x1f\x03\x94\x60\x54\x72\xb2\x5e\xe9\x8c\x13\xb0\x34\x99\xac\x50\x97\xca\x75\xde\xff\x6a\x0a\x93\xc7\x79\x0c\x4b\x26\xd7\x4b\x2e\xd7\x8b\x36\xd7\xbf\xf3\xb9\xa6\x15\x7d\xfa\x86\x10\x8c\x06\xd3\x80\x30\x93\xcc\x18\x95\x76\x37\xf3\x18\x35\x66\x70\xd0\x08\xd0\x64\x0d\x0b\xf9\x67\x92\x17\x94\x4a\x32\x8e\xe4\xc5\x70\xb6\xa4\x98\x71\x9e\xe6\x28\x26\x7b\x71\x86\x24\xe2\x29\x89\xb4\xca\xbf\xb7\xa4\x74\x84\x71\x1d\xa9\x8c\x85\x4e\x3e\xe5\xe5\x32\x92\x61\xd1\x0a\xf4\x54\x58\x74\x6c\xe3\xb1\x37\x72\x3d\x94\xd6\xc7\xc0\x48\x54\xc9\x3a\x5a\x51\xe1\x19\xff\xca\xcb\xc8\x08\x94\xfa\x52\x5f\x92\x98\x74\x24\xf9\xd2\xb9\x69\xfe\xcb\x6a\x72\xdd\xc9\x47\x52\x01\x93\xa2\x14\xa9\x73\x1b\x97\x69\x1f\xfd\x19\x7d\xfe\xec\x63\x38\x97\x90\xf8\x99\x6a\xce\x62\x69\xe6\x32\x8b\x1c\xdb\x16\x7c\x31\xef\xd2\xfd\x7e\xd0\x3f\x45\xf2\xfa\x21\xb8\xca\xfc\xdf\xe5\xf3\x5d\x1f\x81\xff\xc9\xe6\x3f\x47\x33\x9b\x58\x3c\xf9\xd3\x85\xbe\x47\x7a\x22\x97\xdd\x7b\xfb\x11\x00\x00\xff\xff\x4d\xa0\x76\x8a\x8a\x05\x00\x00")

func type_test_graphql() ([]byte, error) {
	return bindata_read(
		_type_test_graphql,
		"type/test.graphql",
	)
}

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x53\xcd\x6a\xdc\x30\x10\xbe\xeb\x29\xbe\xc5\x97\x16\x4a\x1e\xc0\xb7\x25\xa5\xb0\x50\x4a\x4b\xb2\xa7\x92\xc3\xac\x35\xf6\x4e\x23\x4b\xae\x34\x4a\x31\x25\xef\x5e\x24\xaf\x37\x4e\x53\xe8\x4d\x33\xd6\xf7\x37\x23\xb3\xcf\x23\x8e\x89\xe3\xfd\x3c\x31\x7e\x1b\x60\x24\x4f\x03\x47\x03\x88\xb7\xf2\x24\x36\x93\x33\x80\x65\xc7\x03\x29\x9b\x67\x63\x9a\x8a\x80\x24\x10\xf4\x2c\x1e\x5a\xc0\x7a\x26\x45\x17\x9c\x23\xe5\x04\xf1\x7d\x88\x23\xa9\x04\x0f\x3a\x85\xac\xd0\x33\xaf\xe4\x1f\x5e\xe8\x1a\x90\xb7\x1b\x29\xe4\x42\x5d\x08\xd3\x0d\xee\xc2\xc8\xe8\x85\x9d\x4d\xa0\xc8\xf0\x41\x31\x04\xf1\x03\x34\xe0\xc4\xa0\x27\x12\x47\x27\xc7\xb0\x3c\xb1\xb7\xe2\x07\xd3\x20\xf8\xaa\x55\x4d\x85\x7e\x21\x9c\x43\xae\x0c\x91\x7f\x66\x4e\x5a\x28\x2c\x29\xa1\x0f\xf1\xc6\xd4\x9b\x35\x52\x19\x40\xa9\xda\xeb\x4c\x76\x06\xe0\x91\xc4\xb5\xb8\xd3\x58\x04\x80\x5e\x62\xd2\x2f\x34\xf2\xda\x2b\x97\x1c\xbd\xed\x29\x3b\x9e\xce\xc1\xf3\x06\xfc\x23\x9c\xee\x45\xdd\xb6\x55\xa0\x9f\xc3\x20\x7e\x8b\xed\xc2\x38\x91\x9f\x5b\xdc\x2e\x07\x03\x4c\x31\xf4\xe2\xf8\x30\xd2\xc0\xc7\xb8\x75\x44\x9d\xca\x93\xe8\xfc\x6e\xa2\x81\x5b\x7c\xa5\x81\xdf\xb7\xd8\x5f\xba\xa5\x5c\x2f\xf1\x6d\xc8\x31\x71\x6a\xf1\x7d\xbf\xa9\x77\x0f\x06\x68\xd0\x71\x54\xe9\xa5\x2b\x2b\xac\x0d\xb2\x36\x72\x4a\x2d\xf6\xcb\xa1\xac\xbf\xbe\x99\x95\xfb\xfa\x6e\x3c\xff\x5a\xa8\x56\x25\x52\xb6\x97\x1c\x8e\x95\xed\xf5\x6b\x4f\xe2\xae\xe5\xb3\x59\xc6\xbf\xf5\x0a\x29\x90\x91\xbd\xa6\x12\x45\x7c\xa1\xfa\x96\x39\xce\x55\x89\xed\xf0\xe2\x5f\x74\x2e\xd6\x4b\xee\x83\xef\xc3\x92\xbd\x9c\xde\x30\x57\x6c\xce\x62\x5b\x1c\x8f\x87\x8f\xbb\xeb\xaa\xb7\x51\xea\xe4\x23\x17\xc1\xbd\xbe\x5e\x47\xb1\x5b\xb6\xb1\xda\x6e\xf0\x42\x7f\x99\xe2\x7f\xcd\x37\xf8\xcb\xfe\x05\xf8\x50\xbf\xfc\x23\x44\x83\xaa\xd4\x60\x7f\x71\x80\x2e\xc7\xc8\x5e\xdd\x8c\x13\xd7\x3f\x81\x1e\xd9\xe3\x34\x83\xea\x4b\x37\x6f\x4c\xd5\xdc\xaf\xed\xd7\x40\x0b\xcf\x5e\x95\xc7\x49\x5b\x1c\xbc\x96\xf6\x28\x3e\x2b\xa7\xfb\x48\xdd\x23\xdb\x16\x9f\x5c\x20\xdd\x99\x67\xf3\x27\x00\x00\xff\xff\x27\xaa\xe4\x2d\x2a\x04\x00\x00")

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
	"type/company.graphql": type_company_graphql,
	"type/course.graphql": type_course_graphql,
	"type/delegate.graphql": type_delegate_graphql,
	"type/individual.graphql": type_individual_graphql,
	"type/lesson.graphql": type_lesson_graphql,
	"type/manager.graphql": type_manager_graphql,
	"type/module.graphql": type_module_graphql,
	"type/question.graphql": type_question_graphql,
	"type/test.graphql": type_test_graphql,
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
		"users.graphql": &_bintree_t{type_users_graphql, map[string]*_bintree_t{
		}},
	}},
}}
