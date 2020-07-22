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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xcb\x6e\xe3\xb8\x12\x86\xf7\x7a\x8a\x0a\xb2\x49\x2f\x4e\x9f\xbd\x77\x1d\x0b\x8d\x31\x90\x20\x99\xb6\xf4\x00\x15\xb1\x2c\x13\x4d\x91\x1a\x92\x4a\x8f\x31\xc8\xbb\x0f\x48\x49\xbc\xc9\xe3\x4d\x7a\x67\x7e\xaa\xbf\x2e\xac\x22\xe9\x7b\x18\x26\x8b\x96\x2b\x69\xa0\x43\x09\x83\x62\xfc\x74\x01\x7b\x26\x60\x6f\x5f\x81\xfe\xa6\x6e\xb2\xc4\xc0\xd0\x5f\x13\x49\xcb\x51\x88\xcb\xd7\xaa\xe2\x72\x9c\x2c\xb4\xa3\x50\xc8\xbe\x73\x41\xcf\x64\x11\xfe\xa9\x00\x4e\x5c\x50\x73\x19\x69\x07\x47\xab\xb9\xec\xef\x2a\x80\x4e\x49\x4b\xd2\x3e\x91\xec\xed\x79\x07\x07\x69\xef\xaa\x8f\xaa\xb2\x97\x91\x12\x17\x3f\xc8\x8c\xde\xc5\xa4\x45\xaa\x36\x53\xd7\x91\x31\x8d\xfa\x49\x32\xf2\x8f\x6d\x0e\xc7\xd9\xd0\xfb\xf8\x4f\x91\x0f\xfa\xbc\x94\xec\x4d\x91\x0d\x5c\x3e\xa9\x9e\xcb\x07\xef\x71\x07\xdf\x02\xb9\xfb\xb2\x83\x6f\x93\x3d\x7b\x3f\x15\xc0\x80\x12\x7b\xd2\x99\xf5\x73\xc2\x4a\x7b\x46\x82\x7a\xb4\x94\x09\xea\x14\xe6\x0a\xb7\x59\x9a\xd0\xd2\xe2\x74\x95\xec\x53\x78\x70\xcc\xe9\x96\xb5\xdb\xb2\x91\x6d\x45\x6d\x0a\xaf\x88\x46\xad\x5c\xb7\x0e\x03\xf6\x34\xef\xe2\x0f\xd7\x64\x63\xa3\x83\xb4\xbd\x4e\x9c\x77\xab\x0c\xfc\x9a\x38\xdc\xfa\x58\xda\x93\xe7\xe0\x36\x68\x93\x78\x9d\xc2\x90\xf8\xa3\x52\x82\x50\xde\xc5\x4d\xf2\x7d\xca\xb7\xc8\xa3\x20\xf1\xab\x90\x65\x66\xde\x46\xb4\x31\x9f\x93\xca\xcc\xeb\x88\x6e\x24\xb4\x76\x36\xcf\x69\xa5\x41\x98\x63\xb7\x93\x4a\x1a\x0a\x79\x96\x5e\xda\x8c\x06\x2f\x2b\x88\xe1\xf7\x6a\x18\x51\x5e\xf2\xe8\x0b\x8c\xc1\xe7\x75\x88\x56\x88\xda\x14\x5e\x11\xdd\x2f\xbb\x13\x49\x16\x7b\x9d\xa0\x0a\xc0\x9d\x7b\xcf\xae\x66\xe2\x0d\x96\xe3\x74\x75\xbe\xbd\x81\xa6\x0e\x47\xdb\x9d\x31\xbd\x10\xd2\xad\x07\xc0\x71\xd4\xea\x3d\x94\x31\x4d\x9c\xed\xa0\x6d\x0f\x75\x9a\xb7\xbb\x98\x26\x71\xe2\xe2\x95\x24\xe3\xb2\x7f\xd1\x8c\xf4\x43\x27\x38\x49\x7b\xa4\x4e\x93\x0d\x01\x72\xef\x06\xdf\xe9\x45\x0a\x2e\x69\xaf\x26\x6d\x42\x4f\x8e\x05\x4f\x76\xca\x2d\xe1\x1e\x9a\x97\xfa\x05\x50\x08\xf5\x0b\x04\x97\x3f\x17\x5f\x7b\x81\xc6\x68\xa5\x86\xad\xbb\xe2\xd3\x15\x8f\x8a\x29\x40\xc6\xc0\x4e\x56\x69\xb0\x0a\x3a\xff\xc9\xdf\xb0\xee\xc7\x23\x4a\x49\xfa\x53\xe7\x79\x9c\x74\x77\x46\xb3\x54\x65\x56\xe5\x6b\x8e\x43\x6e\x05\x4f\x46\x79\x9d\x92\xb4\xca\x3a\x61\x37\x4e\x51\x83\x7d\x3e\xc2\x0d\xf6\xc1\xbc\xc1\x3e\xce\x1c\x5a\xea\x95\x2e\x07\x7e\xa1\x71\x03\x17\x10\x23\x3c\x91\x31\xaa\xb8\x39\x66\x16\x44\xf3\x32\x9c\x92\x5c\xd1\x26\x6c\xab\x98\x4b\xcf\x15\x75\xc2\x6e\x94\x7e\x90\x8c\xbf\x73\x36\xa1\xc8\x93\x8b\xbc\xb8\x44\xe2\x87\xcd\x35\xb2\xf5\xd5\x16\x3c\xf8\x2a\x3f\x04\x5f\xfe\x25\x7d\x1b\xb8\x6d\x92\x31\x3a\x06\x12\x1c\x44\xf4\x8a\x17\x37\x52\x70\xbf\x30\x03\x18\x5e\xc1\xff\xf3\x10\xc1\x00\x4a\xf3\x8b\xb4\x71\x63\x8c\x60\xc9\xd8\xd8\xff\x24\xd6\x3e\x90\xa2\xf0\x24\x56\xa8\xb8\xc9\x66\x9d\x95\xc2\x88\xa2\x70\xee\x55\x2a\xac\x03\xb9\xda\xa7\x39\xed\x4f\x9d\xb1\xb9\xca\x3f\x9d\x8a\x97\x53\xb8\xd2\xa2\xda\x15\x97\x15\x97\x4e\xda\x8c\x16\x95\x6f\x9d\xcc\xd5\x97\x4e\xea\x8c\xde\x98\xd6\x67\xc5\x26\x51\x3c\x76\x33\x2b\xd2\x9f\x61\x99\x7c\x2e\x6f\x13\x56\x24\x5e\xca\x97\x3f\x0d\x99\xbc\x4e\xd8\xad\xbb\xc5\x5d\x9d\xc5\x70\x39\x14\xef\x17\xb7\x8a\x13\x95\x9a\xb7\x11\x6d\xcc\xfd\x95\x7c\xe4\xbd\x44\x3b\xe9\xdf\xf1\x97\xaa\xc9\x1c\x5e\x49\x21\x06\xcb\x73\xf1\x6f\xb4\x21\xbb\xfe\x3d\x48\x9e\x95\x61\x14\xe4\xff\xf6\xfe\xcf\x19\x18\xe0\xa7\xe4\x70\xc2\x88\xc6\x10\x03\xa5\xe1\x84\x5c\x10\x03\x84\x6e\xd5\x2e\x0f\x0c\x3c\x7c\x57\x7a\x8e\x63\xbe\xfc\xbe\xd7\x22\x74\xe7\x51\xa8\xe2\xea\x77\x24\xb6\x53\xa8\x3e\x6c\x50\x6a\xdb\x06\x52\xda\xbe\x09\xd5\xff\x41\xc8\x3e\x79\x64\x63\xc8\xc4\xd9\x36\x7a\xf2\xf1\x5a\x22\x8f\x8a\x5d\x3e\x91\xc6\x47\xf5\x6f\x00\x00\x00\xff\xff\x1f\xfa\x49\x9b\xad\x0d\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x4d\x6f\xdc\x36\x10\xbd\xeb\x57\x3c\x63\x2f\x36\x60\x34\x77\xdd\x6c\x07\x01\x0c\x34\x6d\x52\xaf\x4f\x45\x51\x70\xc5\x91\xc4\x9a\x22\x15\x7e\xc4\x56\x8b\xfc\xf7\x62\xa8\x6f\x55\xbb\x48\x7a\xd3\xea\xbd\x99\x37\x33\xcb\x37\xd4\x01\x5f\x22\x39\x45\x1e\xbe\xb6\x51\x4b\x18\x1b\xd0\x58\xa9\xca\x0e\xa1\x26\xc8\xd3\x4f\xa0\x37\x2a\x62\x20\x09\x65\xd0\x0a\xe7\x84\xd6\xa4\xb3\xd0\xb5\x84\x4f\xa2\xa2\x47\x53\x5a\xfc\x93\x01\xc1\x06\xa1\x73\x3c\x9a\x70\x85\x03\x7e\x89\xcd\x89\x1c\x6c\x89\x56\x54\xe4\x21\xca\x40\x0e\xa1\x56\x1e\xd6\x50\x06\xd8\xb2\xf4\x14\x26\xfe\xb1\xa6\xe1\x15\xc7\x24\x1e\x07\xa2\x74\xb6\x49\xa5\xf8\x20\x5c\xc8\x00\xad\x1a\xb5\x0e\x6b\xc4\x1b\x97\xff\x6a\xd0\x92\x4b\x51\x19\x50\xa9\xaf\x64\x76\xaa\x51\x81\x1a\xdf\xa3\xd9\xb7\x2c\x53\xa6\x8d\x21\xf5\x91\x7a\x58\x14\xb5\x92\x9a\xa9\x1f\x85\x11\x15\xb9\x0f\x4a\x73\x3f\x1c\x63\x44\x43\x39\x9e\x82\x53\xa6\xca\x00\x6a\x84\xd2\x8b\xdf\x7f\xd9\xd3\x51\x05\xbd\xa4\xc4\xa8\x64\x8e\xe7\xe7\xc7\xf7\x3c\x36\xd2\xd4\xd6\xd6\xcc\x84\x49\xeb\x3d\x69\xaa\x44\xa0\x85\xd8\x3a\x34\x14\x7f\xf2\xcf\x29\xf1\x8f\x97\x72\x41\xfd\xc1\x46\xe7\xc9\x9f\xef\x54\x14\x05\x79\x7f\xec\x5a\xca\x71\x37\x3d\x67\xc0\x49\x14\x2f\x95\xb3\xd1\xc8\x87\x9a\x8a\x97\x1c\xf7\xd6\x6a\x12\x26\x03\x5a\xa7\x0a\xca\xf1\x41\x5b\xc1\x13\x16\x5a\xdb\x57\x92\x47\x7b\x1f\xbb\x89\x86\x03\x9e\x6a\xfb\x0a\x6b\x74\x87\xa2\xaf\x02\xa1\x16\x01\x9d\x8d\x10\x8e\x20\x62\xa8\xad\x53\x7f\x93\x44\xb0\x38\x59\xfb\xb2\x2c\xbb\x69\x85\xe9\xce\xcd\x6c\xd5\xc3\x14\xf4\x33\x79\x6f\xcd\x77\xc5\x00\x41\x54\x3e\xc7\xef\x0c\xfe\x31\xa7\x78\xea\xb4\x16\xa7\x78\x61\x5e\xf4\x56\xe8\x28\xa9\xd7\x5a\x0e\x65\x00\x3e\x5a\x19\xf9\xcf\xf9\x0f\x70\x24\x1f\xe6\xd7\x93\xe2\xa3\x91\xea\xab\x92\x51\xe8\xef\x2c\xfc\xff\x9f\x87\x40\xae\x14\x45\x32\xbc\x32\x22\x90\xfc\x1c\xc9\x75\x49\xb0\x1d\x76\x40\x3e\x6d\x83\xb9\xc2\x5f\x9d\x24\x77\xdf\xf3\x84\x2f\xc8\x48\x65\xaa\x65\x83\xa5\x22\x3d\x1d\xe0\x2b\x0e\x4c\x7b\x65\xce\xae\x52\xe6\x11\x07\x84\x6c\x94\xb9\x9e\x9b\xbc\xc9\x71\xc7\xaf\x46\xc8\x5f\x73\x3d\x7d\x2d\x23\xc6\xcf\x57\x59\x06\x34\xbd\x77\xd7\xe1\x83\xa1\x67\x78\x99\xe2\x16\x65\x1a\x6d\xbe\xf6\xfd\x2d\x6c\xdf\x59\x3e\xb6\x38\x27\xe2\x30\x16\x93\x83\x79\x17\x6a\x57\x37\xf9\xe4\xe9\x05\x63\x5f\x70\x6d\xfe\x5d\xc5\x91\x32\x4a\x16\xfd\xd1\xdf\x28\x0e\x86\x98\x70\x75\x46\x70\x65\x9c\x5d\xbd\x81\x31\xca\xe9\x74\x90\x37\x6a\xfd\xe9\x9e\xd0\x7d\xa9\xa5\xdd\x76\x95\x7a\xc2\x28\x74\xd2\xb6\xda\xc8\xdc\x6b\x5b\x0d\xc8\x5a\x62\x27\x19\x73\xe7\x11\xa5\x75\x72\x66\x00\x8b\x85\x77\x66\x00\xcc\xf8\xd4\xdf\x30\x7d\xaa\x6b\xae\x8a\x6f\x99\x09\x66\x13\x7a\xce\xf7\xec\xd3\xb1\x52\x93\x4d\x37\x3d\xcc\xfe\x5d\xb1\xf6\x8b\xdb\x9a\x7d\xb7\xbe\x99\x34\xf6\xfb\x25\x92\x0f\xea\xdc\xff\xf0\x79\x40\x2f\xa4\x1c\x29\x43\xd3\x63\xbe\x4d\x2b\x23\x2b\x6d\x10\x1f\xf6\xd5\x78\x8d\x5d\x50\x62\x78\x50\xe1\x1c\x1b\x05\x46\xd9\xa3\x69\x47\x9e\xb1\x68\xc2\x2e\x39\x34\x11\x06\x8d\x3e\xd3\x46\xa5\x67\xf0\xe0\x3c\x09\x57\xd4\xe3\x52\xdf\xd5\x5b\x6f\xfc\x9b\x1c\x4f\xab\x98\xdf\xc8\x47\x1d\x70\xc0\x1d\x5f\x72\x7e\xc8\xa8\x4c\x35\x36\x71\x3b\x7a\x04\xc2\xc8\x7e\x6e\x19\x70\x48\xf7\xcb\x75\xa0\xb7\x30\x2e\xbe\x1b\xbc\x7b\x37\x84\xf3\x97\x94\xd6\xfd\x27\x12\x25\x26\x9f\x44\x72\x41\x95\xaa\x10\x21\xed\xdf\xeb\x60\x5f\xf8\xdb\x67\xd8\x9a\x7c\x30\xd7\x04\x5e\xa4\xfc\x5c\x59\xb7\x5d\x06\x2b\xd9\x1c\x0f\x3d\xab\x1b\x66\x76\x80\xa7\xc2\x1a\x29\x5c\xf7\xf0\xe3\xf1\xdf\xb2\x7f\x03\x00\x00\xff\xff\xdc\x6e\xe5\x58\x6f\x0a\x00\x00")

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

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4b\x6f\xe3\x36\x10\xbe\xfb\x57\x4c\xe2\x43\x5a\x60\x73\xe8\xe6\x81\xd6\xb7\xac\xe3\xa2\x06\xb2\xd8\xd4\x0f\xb4\x45\x90\xc3\x84\x1a\xcb\xc4\x52\xa4\x4a\x0e\x93\x08\x8b\xfd\xef\x05\x49\x59\xa2\x1d\x7b\x9b\xb4\x28\x50\x18\xb0\x1e\xf3\xfa\x38\xf3\xcd\x90\x22\xed\x2b\xb8\x12\x82\x9c\x5b\x34\x35\xc1\x97\x01\x80\x25\xc7\x56\x0a\xa6\x62\x00\x60\x6a\xd2\x83\xaf\x83\x41\x54\x9c\xb3\xf5\x82\xbd\xa5\x89\xa2\x8a\x34\x47\xf5\xca\x14\x5e\xd1\x00\x80\xc9\xf1\x00\x40\x91\x73\xa6\xb7\x19\x1b\x6f\x1d\x75\xce\x8d\x56\x52\x07\x6d\xa1\xd0\x39\x6b\x4c\x15\x34\xa5\x66\xb2\x2b\x14\x04\xf3\x46\x29\x7c\xf0\x6e\xca\x54\x45\x03\x8d\x15\x8d\x42\x64\xa9\xcb\xa3\x01\x80\xf7\xb2\x18\xc1\x72\x39\xbd\x0e\x4f\xdc\xd4\x49\xba\x85\x2b\x48\x84\xa9\x6a\x45\x4c\x23\xf8\x60\x8c\x22\x8c\x88\x82\x3a\xdc\x44\x80\x20\x83\x3c\x68\xbb\xff\x32\x28\x40\x55\x9f\x2d\x67\x37\x1b\x6f\x31\x4f\xcf\x9c\x3b\x67\x2c\xdd\x08\xee\x16\x58\x1e\xdd\x77\x20\x53\xda\x22\x98\x10\x7a\x9a\xfc\xa7\xc8\x7d\x4a\x8f\xf6\x60\x7d\x40\xf1\xb9\xb4\xc6\xeb\x62\xbc\x26\xf1\x39\x87\x82\x5d\xa1\x47\x59\xd1\x83\x51\x6d\xa5\xa0\x11\xfc\xac\x0c\xb6\x0b\x51\xc6\x66\x98\xe9\x59\x90\xad\x39\x7b\x23\x35\x5b\x53\x78\xc1\xd2\xe8\xec\xf5\xda\x3c\x2d\xcc\xb8\x4b\x43\xf6\xde\x5b\x97\x4b\x62\xa8\x01\xc0\xd3\x1a\xf9\x0f\xe3\x6f\x08\xad\x1e\xc1\x5d\xbb\x8e\xfb\xc8\xc3\x3f\xbd\xb4\xa9\x44\xdb\x12\x57\x93\x90\x2b\x29\x16\x64\x2b\x97\x05\x11\xc8\x54\x1a\xdb\x8c\x60\xdc\xde\x85\x45\x2b\x65\x9e\xa8\x58\x98\x0f\xbe\xc9\x93\xe1\xda\xa2\x07\xd7\x59\xfd\x63\x80\x07\xd4\x9a\xec\xb4\xc2\x92\xb6\x6b\x47\xcf\xb5\xb4\xe4\xa6\xfa\xa3\xd1\xbc\x76\x5d\x61\xe2\x7b\x0c\xb9\x58\x98\x89\x2e\xa2\xb4\x0b\x76\xb4\x53\xd5\x5b\x2c\x29\xa7\xdf\x2d\x96\x52\x23\x53\xf1\xab\x27\xdb\xc4\x9a\x53\x51\x52\x40\x96\x0c\x02\xa4\x1a\x4b\x9a\xea\x95\x19\x05\xf5\x78\x97\x1a\xa7\xf6\x0c\x73\x7c\xa4\xf1\xa6\x9f\x92\xc9\x34\x0a\x32\xfa\xec\x30\x25\x64\x8b\x2c\x87\x2c\x22\x53\xa2\x44\xe0\x77\x96\xc5\xf0\xd8\xbd\x7c\x35\x01\xde\x4c\xbf\x1d\xf6\xed\x21\x5f\xdb\x20\xb1\xfd\xbe\x55\xfe\xac\x6a\x73\x1f\xfd\xe7\x4e\x3c\x1b\xdb\x2f\x09\x86\xe9\x62\x56\xc0\x6b\x82\x45\x90\xc2\x77\xdd\x50\x02\x91\xfa\xcf\x68\xd5\x7c\x1f\x9a\x18\x9f\x6f\xd1\xb2\x14\xb2\xc6\x48\xc7\x94\x51\xc7\x68\xf9\x1a\x03\x9f\x17\xb2\x0a\x4b\x21\x5d\x6c\x3d\x2b\x23\x70\x27\x43\xfb\x38\xf4\x77\x14\xda\xae\xf5\xa7\x38\x44\x5f\x5f\xe8\x7d\x05\x3d\x54\xfd\x21\x38\x12\x46\x17\x68\x9b\xf1\x2b\x88\x00\x43\x98\xaf\x8d\x65\x28\xc8\x09\x2b\xeb\x80\x3f\x24\x35\xe5\xef\x00\x4d\x60\x08\x27\x57\x0f\xc6\x33\xf0\x5a\xba\x56\xf7\x24\x04\x0e\x5a\x87\xa6\x48\xb0\xfa\xc5\x3c\x01\x9b\x6e\xce\x1e\xb6\xdf\x37\x6d\x60\x08\x13\xc7\xb2\x0a\xad\x96\x54\x0e\xf9\x3a\x3c\x97\x02\x88\xdf\xd6\xc8\xd0\x18\x7f\xa2\x14\xa8\x20\x3f\x81\x07\xaf\x14\x31\xd4\x46\x6a\x76\x87\x67\x57\xb0\x9e\x65\xa2\x97\x86\xff\xd7\xf6\x71\x9b\xdd\xae\x1b\x4a\xdd\xb0\xfc\x77\x7c\xee\x9d\x45\x12\x1f\xde\x5b\xf3\x3d\xb8\xb7\xb6\x14\x08\x8c\x65\xdf\x06\xbb\xfb\xe1\x56\x3a\xfa\x51\xbc\xc0\x32\xaa\xf7\x6e\x5f\x6f\xbb\x69\x8c\x7f\xe2\x20\xc7\xbd\xf1\xf3\x16\xf0\xc9\xfe\xd6\x5b\xb1\x46\xd7\xce\x00\xd7\x3b\x48\xfc\x0d\xf5\x0e\x3b\xd3\x7d\x4c\x9c\x23\xdb\x13\xe0\x08\x86\x70\x23\x1d\x87\x26\x2d\x48\x51\x89\x4c\x2e\xc2\x77\xb1\x68\x6c\x71\xaa\x1f\x8d\x14\x34\xa9\x50\xaa\x8c\x02\x81\x7f\x35\x53\xd1\xf2\xa6\xdb\xdc\x5e\x90\x76\x6c\xf4\x4a\xda\xea\xe5\xc9\x6b\x07\xf5\x8c\x5c\x6d\x74\x7b\xca\x61\x8b\xda\x61\x6c\xe0\xf1\xee\x29\x2a\x60\x66\xeb\x09\x64\x9a\xd6\xa1\x65\x51\x37\x20\x1d\x18\x0d\x08\xc2\x68\xb6\x28\x18\x50\x17\xc0\x6b\xef\xa0\x30\xe4\xf4\x09\x83\x26\x2a\x42\x93\xd7\xd8\x80\xaf\x57\xd6\xb4\x63\xdb\xca\x9a\xc6\x4a\x92\xe6\x39\x09\x4b\xf9\x38\x9b\xae\xf6\x61\x09\xb1\x56\xa8\x1c\xbd\x4b\x73\x42\x3a\x28\xe5\x23\x69\x70\x26\x26\x18\x04\xea\x10\xa5\x5b\xeb\x9c\xd0\x8a\xf5\xe6\x64\x31\x23\xe7\x15\xbf\x72\xd3\x7f\x71\x1c\xf9\xd6\xde\x7f\x4d\x01\xde\x81\xbd\xe0\x25\x65\xdf\x72\xf6\x68\x4d\x0e\x23\x18\xc2\x1d\x4c\x7e\xbf\xfa\x78\x7b\x33\x81\xf1\xa7\xe5\x6c\x3e\x81\xf9\x62\xb6\x1c\x2f\x96\xb3\xc9\x60\x08\x00\x5f\x52\x3b\x1f\xa7\x6f\x84\xe3\x77\x6d\xb3\x1c\x9f\x9d\x9f\x9d\x9f\x9d\xe6\xff\xe7\xf1\x72\xfc\xf5\xdd\x96\x5d\xfa\x96\xe8\xed\x7e\x78\x1f\x7f\xa7\xe9\xf2\xfe\x34\x7b\xda\x35\x0d\x5f\x23\xbd\xe1\x8f\x3f\xa5\xdf\xe9\xc5\x65\xfc\x9d\x5e\x5c\x76\x37\x17\x97\x17\xc9\xf8\x7e\xf0\x57\x00\x00\x00\xff\xff\x44\x82\xaf\xf1\x0f\x0d\x00\x00")

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

var _type_individual_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x91\xcd\x4a\x04\x31\x10\x84\xef\x79\x8a\xf6\x35\x72\x13\xbd\x0c\x88\xac\x3f\x73\x12\x0f\xd1\xd4\xc6\x96\x4c\x12\x92\x8e\xb2\x2c\xbe\xbb\x64\x44\x67\x9c\xd9\xb9\x08\xde\x92\xa2\xab\xa9\xaf\x4b\x0e\x09\xd4\x05\xcb\x6f\x6c\xab\xf1\x74\x54\x44\xb5\xb2\xd5\xd4\xf7\xdd\xe5\x99\x22\x7a\xce\x30\x02\x7b\x2e\x9a\xee\x24\x73\x70\x8a\x08\x83\x61\xff\xfd\x6f\x43\x7b\xce\x45\xae\xcd\x80\xb9\xe8\xcd\x5a\x7b\x8d\x4f\xf7\x2c\x1e\xb3\x65\x02\x8f\xf4\x12\xc3\x5c\x6b\xd6\xab\xe8\x38\x4c\xde\x0f\xa5\x16\x61\x77\xc6\x81\x78\x48\x1e\x03\x82\x14\xda\x19\xc7\xa1\x65\xbd\xa9\xc8\x87\x46\x02\xeb\x50\x34\x3d\x4c\x9e\x47\x45\x94\x8c\x43\x17\xf6\x51\x37\xc7\xf8\x6a\xcb\x39\xa4\x2a\x74\x31\xd2\x4e\xf3\xdd\xa8\x1e\xff\x01\x71\x75\xc3\x64\x4a\x79\x8f\xd9\xae\x91\x97\xa1\x6e\x51\x52\x0c\x05\x5f\x6d\x15\x64\x4d\x7d\x41\x9e\x28\xfa\x64\x37\x28\x7e\x77\xbb\x62\x3a\x81\xf4\x47\xa2\x13\x40\x3f\x3c\xcb\x78\x9b\x3c\x9f\x01\x00\x00\xff\xff\xea\x36\x9c\x17\x9e\x02\x00\x00")

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

var _type_module_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\xcb\x4e\xf3\x30\x10\x85\xf7\x7e\x8a\xc9\x6b\x64\xf7\xff\x14\xa4\x48\x20\x15\xd2\xc0\x02\x75\xe1\xc6\x43\x64\xe1\x4b\xe4\x4b\xa5\x08\xe5\xdd\x91\xed\x36\x71\x42\x41\x65\xc5\xaa\x49\x66\xce\xf4\xcc\x77\x12\xbb\xa1\x47\x78\xd0\xcc\x0b\x04\x2e\x7b\x81\x12\x95\xb3\x50\x0f\x42\xd0\x83\xb7\x95\x43\x09\x1f\x04\xc0\x7b\xce\x4a\x68\x9a\x6a\x53\x10\x00\x45\x25\x96\x50\x3b\xc3\x55\x17\xee\xc3\x94\x78\xef\x5b\xe7\x0d\xde\xa6\x31\xa1\x72\xa0\x4a\xa1\xa9\x24\xed\xb0\x79\xba\x3f\x6b\x08\x00\x43\xdb\x1a\xde\x3b\xae\xd5\x62\x92\xa1\x2a\x15\xf2\xa7\x47\xcd\x5b\xd4\x47\x34\xcb\x19\x47\xce\x50\x97\xf0\x1c\x7e\x08\x80\x3d\x99\x2e\xe1\x35\xf7\x5f\xec\x09\x40\xab\xc3\x72\x0e\x4b\xf8\xaf\xb5\x40\xaa\xc8\x48\x08\x2a\x2f\x93\x7a\x17\x30\x84\x3d\x5f\xaa\x7a\x57\xfd\x0b\xc5\x48\x26\x16\x63\x21\xad\x38\x35\x07\x57\xde\x88\xd9\xe4\x48\x08\x57\xbd\x77\xa9\xa5\x8a\x97\x57\xea\xa2\x8d\x94\xc1\x89\x5c\x52\xa2\x75\x04\x40\xa0\xb5\x5a\xcd\xf3\x53\xe3\x14\x4c\x9a\xbf\x50\x17\xab\xbc\x26\xe9\x8d\x41\xea\xf0\x34\x60\x72\xf8\x25\x4c\xda\x05\x84\x51\xbb\xff\x6d\x52\x59\xde\xb5\x6f\x5b\xb4\x76\xa7\xdf\x51\xe5\xa1\x9d\xb3\xfc\xae\x9e\x85\x1a\x4d\x2e\x93\x9d\xd7\x2f\xf6\x53\x4c\xf9\x62\x5b\x3a\x08\x4d\x59\x5c\x4d\xc6\x27\x67\x3c\x33\x88\xa6\x67\x17\x40\x2c\xdf\xf1\x35\x86\x1c\xd3\x65\x2a\x17\xa1\xfc\x19\x93\x7c\xc7\x6b\x98\x6c\x30\x7c\x1f\x3f\x31\x19\x49\x76\x5a\x6c\x69\xb7\x38\x31\xb6\xb4\xe3\x8a\x3a\x64\x8f\x1e\xcd\x10\xb5\xc8\x3a\x9c\xed\x05\x86\x3d\xed\xb0\x52\x6f\xba\x0c\xed\xf1\x6a\xfd\x5e\xdf\x71\xe1\xd0\xac\xfe\xfa\x3a\xfa\x23\xf9\x0c\x00\x00\xff\xff\x5e\x7d\xd8\xc5\xcd\x04\x00\x00")

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

var _type_test_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xcb\x6e\xdb\x30\x10\xbc\xeb\x2b\x56\xf0\x5f\xe8\x96\x26\x0d\xe0\x53\x1d\xd8\x3e\x14\x81\x0f\xb4\x39\x56\x88\x4a\xa4\x4a\x2e\x1b\x18\x45\xfe\xbd\xe0\x52\x2f\xbb\x55\x1f\x68\x4e\xe6\x63\x67\x3c\x3b\x3b\x22\x5f\x3a\xd0\x0e\x81\xc9\xb4\x5d\x83\x16\x96\x03\x6d\x2f\x4d\xa3\x8e\x31\xac\x19\x2d\x7d\x2f\x88\xac\x6a\x51\xd1\x96\xbd\xb1\x75\x59\x10\xc5\x68\x74\x45\xfb\xfd\xfa\x21\xed\x12\x87\xdc\xc6\x13\x47\x8f\x8f\x99\x26\xdd\x9c\x5c\x22\x65\x54\xf4\xc1\xb9\x06\xca\xa6\x6a\x55\x87\x8a\x9e\x77\xaa\x2e\x0f\x05\x91\x62\x46\xdb\x71\xb8\x6b\x1a\xf7\x0a\x5d\xd1\xda\x72\x41\xd4\xa9\x10\x36\xf0\x27\x58\x56\x35\x2a\x7a\x6c\x9c\x4a\xe7\x5f\x23\x02\x1b\x67\xc3\xce\xdd\xd9\xf0\x0a\x3f\x00\xbc\xb2\xda\xb5\x26\x20\x1f\x87\xf9\x5f\x8e\xa0\x8a\x9e\x9f\xfa\x75\x79\x28\xde\x8a\x62\x6c\x7f\xa3\x6a\xcc\x2d\xd8\xa8\xda\x58\xc5\xd0\x4f\x11\xfe\x22\x26\x40\xd7\x10\xe5\x08\x7c\x10\x89\x35\xd6\xf6\xec\xaa\x54\x2c\xab\xc4\x68\x6c\x17\x59\x28\x1f\x4d\xc3\xf0\x02\x9d\xfc\xba\x31\x73\xd4\x70\xef\xa1\x18\x59\xc9\xa5\x71\x4a\x0b\x8e\x11\xb8\x12\xb2\x89\x7a\xaa\x5c\xcb\xfe\x57\xf3\xe9\x3d\x96\x01\x2d\x99\x5c\x2e\xb9\x5c\x2e\xda\x5c\xfe\xce\xe7\x92\x56\xf4\xe9\x1b\xbc\x37\x1a\x81\x3a\xf8\x91\x64\xc4\xa8\xb4\xbb\x99\x47\xd6\x28\x60\xaf\xe1\xa1\xa9\x31\x81\xc9\x9d\x89\x5f\x30\x55\x92\xb1\xc4\x2f\x26\x88\x25\x93\x19\xc3\x34\xb3\x18\xf1\x62\x80\x24\xe2\x3e\xa3\xb4\x92\xdf\x5b\x52\x3a\xc2\xd8\x9a\x94\x60\xa1\x93\x4f\xb2\x5c\x46\x06\x34\x38\x31\x74\x5f\x38\xe9\xd8\xc6\x63\x6b\xf8\x7a\x28\x27\x17\x7d\x40\xa2\x4a\xd6\xd1\x8a\x26\x9e\x7c\x25\xcb\x18\xe0\x29\xf5\xa5\xbe\x24\x31\xe9\x88\xe5\x73\xb4\xfd\xfc\x97\xd5\x48\xdd\xc5\x45\x52\x1e\xbd\xa2\x14\xa9\xa1\x8d\x79\xda\xb3\x3f\xd9\xe7\xcf\x2e\xfa\xa1\x84\xd8\x8d\x54\x63\x16\xa7\x66\xe6\x59\xcc\x9a\xb7\xac\x38\x86\x8a\xee\x67\xbb\x21\x49\x29\x58\xb3\x30\xec\x7c\x04\x99\xf3\xd4\x65\xae\xb9\xfa\xbf\xec\xde\xbe\xd3\x3f\x45\xfa\xfa\x89\xb9\xfa\x66\xfe\x2e\xdf\xef\xfa\x88\xfc\x4f\xb6\xff\x39\xda\x32\x84\xc9\x93\x3f\x3d\x08\x0f\x48\x4f\xec\xb2\x7b\x6f\xc5\x8f\x00\x00\x00\xff\xff\x88\x67\x13\xe7\xe5\x05\x00\x00")

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

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x54\x4d\x8f\xd4\x30\x0c\xbd\xe7\x57\xbc\x51\x2f\xbb\x12\xda\x1f\xd0\xdb\x68\x57\xa0\x91\x00\x81\x76\xe6\x84\x38\x78\x1a\xb7\x13\x36\x1f\x25\x71\x16\x55\x68\xff\x3b\x4a\x3a\x2d\x85\xb9\xd9\x6e\xfc\xfc\x9e\x9f\x55\xf6\xd9\xe1\x94\x38\x1e\xa7\x91\xf1\x5b\x01\x8e\x3c\x0d\x1c\x15\x60\xbc\x36\xaf\x46\x67\xb2\x0a\xd0\x6c\x79\x20\x61\xf5\xa6\x54\x53\x3b\x60\x12\x08\x72\x31\x1e\x52\x9a\xe5\x42\x82\x2e\x58\x4b\xc2\x09\xc6\xf7\x21\x3a\x12\x13\x3c\xe8\x1c\xb2\x40\x2e\xbc\x80\xbf\xfb\x0b\xd7\x80\xbc\xde\x8c\x42\x2e\xd0\x05\x30\x3d\xe0\x39\x38\x46\x6f\xd8\xea\x04\x8a\x0c\x1f\x04\x43\x30\x7e\x80\x04\x9c\x19\xf4\x4a\xc6\xd2\xd9\x32\x34\x8f\xec\xb5\xf1\x83\x6a\x10\x7c\x9d\x55\x49\x85\x7e\x06\x9c\x42\xae\x08\x91\x7f\x66\x4e\x52\x20\x34\x09\xa1\x0f\xf1\x41\xd5\x97\x55\x52\x59\x40\xc9\xda\x75\x27\x3b\x05\xb0\x23\x63\x5b\x3c\x4b\x2c\x03\x80\xde\xc4\x24\x9f\xc9\xf1\x52\x2b\x8f\x2c\xdd\xd6\x84\x2d\x8f\x97\xe0\x79\xd3\xfc\x23\x9c\x8f\x46\xec\xb6\x54\x5a\x3f\x86\xc1\xf8\x6d\x6f\x17\xdc\x48\x7e\x6a\xf1\x38\x07\x0a\x18\x63\xe8\x8d\xe5\x83\xa3\x81\x4f\x71\xcb\x88\x3a\x31\xaf\x46\xa6\xbb\x91\x06\x6e\xf1\x85\x06\xbe\x6f\xb1\xbf\x56\x4b\x5a\x9c\x9d\x1e\x43\x8e\x89\x53\x8b\x6f\x9f\xae\xf1\xee\x7b\xfd\x50\x5f\xf2\x5c\xba\x33\xba\xc5\xc1\xcb\xee\xbe\xc5\xf2\x0c\x0d\x3e\xb0\xc0\x4d\xe8\xe6\xbc\x8f\xc1\x5d\xe3\xc3\x93\x02\x1a\x90\xd6\x91\x53\x6a\xb1\x9f\x83\x72\x28\xf5\xba\x16\x16\xeb\x85\x79\xfe\x35\xa3\x2e\xc4\x49\x58\x5f\x15\x5b\x16\xd6\xeb\xd7\x9e\x8c\x5d\xd3\x37\x35\x1b\xb5\x55\x05\x53\x5a\x1c\x7b\x49\x45\xb4\xf1\x05\xea\x6b\xe6\x38\xd5\x49\xac\x87\xaa\x76\x69\x29\x62\xcb\x86\x0e\xbe\x0f\xf3\x96\x4a\x74\x83\x5c\x7b\x73\x2e\x7b\x38\x9d\x0e\x4f\xbb\xf5\x28\xb6\x52\xaa\x47\x91\xcb\xc0\xbd\xfc\x6b\x5c\xa1\x5b\x7c\x5b\x68\x37\x68\xb0\x5f\x36\xd7\xe5\x18\xd9\x8b\x9d\x70\xe6\x7a\xc9\xf4\xc2\x1e\xe7\x09\x54\x2f\x75\x66\xb2\xee\xbd\x30\x49\x42\x92\xd3\x02\xf8\x5c\xb3\xdb\x41\xa5\xe2\x8c\xcf\xc2\xe9\x18\xa9\x7b\x61\xdd\xe2\xbd\x0d\x24\xf5\x82\x7d\x0c\xd6\xfe\xcf\x34\x8f\xc7\x30\x4b\x5c\xcd\xda\xce\xa8\xc3\x17\x53\x56\x37\xea\x8f\x61\xad\xbe\xa9\x3f\x01\x00\x00\xff\xff\x16\x28\x3a\xfe\x42\x04\x00\x00")

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
