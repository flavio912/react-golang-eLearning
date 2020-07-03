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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x56\x4d\x8f\xe3\x36\x0c\xbd\xfb\x57\x30\xc8\x65\xf6\xd0\xed\x3d\xb7\x6d\x82\x05\x02\xcc\x20\xe9\xc6\xfe\x01\x5c\x8b\x71\x84\x95\x25\x57\x1f\x33\x0d\x8a\xf9\xef\x85\x64\x5b\x5f\x49\x7b\xd4\x33\xdf\xe3\x13\x29\x32\xd9\xc2\xe8\x2c\x5a\xae\xa4\x81\x1e\x25\x8c\x8a\xf1\xeb\x1d\xec\x8d\x80\xfd\xfc\x0a\xf4\x37\xf5\xce\x12\x03\x43\x7f\x39\x92\x96\xa3\x10\xf7\xaf\x4d\xc3\xe5\xe4\x2c\x74\x93\x50\xc8\xbe\x73\x41\x6f\x64\x11\xfe\x69\x00\xae\x5c\x50\x7b\x9f\x68\x07\x17\xab\xb9\x1c\x36\x0d\x40\xaf\xa4\x25\x69\x5f\x49\x0e\xf6\xb6\x83\xa3\xb4\x9b\xe6\xb3\x69\xec\x7d\xa2\x4c\xe2\x07\x99\x29\x48\x38\x2d\x72\xb6\x71\x7d\x4f\xc6\xb4\xea\x17\xc9\x84\x7f\x3e\x7a\xb8\xcc\x81\x41\xe3\x3f\x49\x21\xe9\xdb\x72\xe5\x10\x8a\x6c\xe4\xf2\x55\x0d\x5c\xbe\x04\xc5\x1d\x7c\x8b\xc8\xe6\xcb\x0e\xbe\x39\x7b\x0b\x3a\x0d\xc0\x88\x12\x07\xd2\x45\xf4\x5b\x86\xd5\xf1\x8c\x04\x0d\x68\xa9\x20\x1c\x72\xb0\x64\xf8\x62\x69\x42\x4b\x8b\xe8\x4a\xd9\xe7\xe0\xd1\x63\x9e\xb7\x9c\x7d\xc9\x26\xf6\x48\xea\x72\xf0\x09\x69\xd2\xca\x77\xeb\x38\xe2\x40\x73\x15\x7f\xf8\x26\x1b\x9b\x04\xf2\xf6\x7a\x72\xd9\xad\x3a\xf1\x39\x13\x7c\xd4\x58\xda\x53\x7a\xf0\x05\x7a\x30\x7e\xc8\xc1\x68\xfc\x0f\xa5\x04\xa1\xdc\xa4\x22\x85\x3e\x95\x25\x0a\x50\xa4\x84\x53\x74\x59\x84\x77\x09\x7a\x08\x9f\x4d\x15\xe1\x87\x04\xfd\x8f\xa1\xb5\xb3\xa5\xa7\x15\x8d\xc4\x12\xf6\x95\x54\xd2\x50\x03\xb0\x5d\x9c\xae\x9f\x92\xf2\x5e\x8d\x13\xca\x7b\x29\xbc\x80\x49\x77\x3e\xc7\x0b\x57\xa4\x2e\x07\x9f\x90\xb6\xcb\xc5\x13\x52\xe4\x5e\x1f\x47\x03\xe0\x47\x3a\x60\x4f\x9d\x84\x80\x65\x52\x9e\x3e\xdd\x10\xa0\xa9\xc7\xc9\xf6\x37\xcc\x67\x3d\xaf\x2a\x00\x4e\x93\x56\xef\xf1\x1a\xce\x71\xb6\x83\xae\x3b\x1e\x72\xdf\x7e\xe7\x38\x71\xe5\xe2\x4c\x92\x71\x39\x9c\x34\x23\xfd\xd2\x0b\x4e\xd2\x5e\xa8\xd7\x64\x63\x82\x52\xdd\xe0\x3b\x9d\xa4\xe0\x92\xf6\xca\x69\x13\x9b\x76\xa9\xf0\xac\x52\xfe\x08\x5b\x68\x4f\x87\x13\xa0\x10\xea\x03\x04\x97\xbf\x16\xad\xbd\x40\x63\xb4\x52\xe3\xa3\x5c\xf5\xe9\x89\xa2\x62\x0a\x90\x31\xb0\xce\x2a\x0d\x56\x41\x1f\x3e\xf9\x29\x75\xba\xbf\xa1\x59\xcc\x98\x55\xf6\x5c\xc2\x51\xb2\xc2\xe3\xe3\x8a\xdd\x6c\x71\x28\x5f\x51\x8b\x43\x64\xb7\x38\xa4\xb6\xa3\xa5\x41\xe9\xfa\xcd\x2d\x68\xba\xc3\x02\xa4\x0c\xaf\x64\x8c\xaa\xe6\x72\xc6\x22\x69\x3e\x46\xc6\x51\x32\xfe\xce\x99\x43\x51\xb2\x12\x5e\xcd\x4e\xfa\x90\x4d\x8f\x71\x3f\x47\x6e\xdb\x6c\x81\x5d\x22\x12\xf9\x09\x3a\xe3\xdd\xaf\x26\xd8\x2e\x98\x01\x8c\xfb\xfa\x77\x1e\x13\x18\x40\x69\x3e\x48\x1b\xdf\x15\x04\x4b\xc6\xa6\x5a\x66\xb9\xf6\x11\xa9\xbc\x66\xb9\xe2\x64\xb6\xc5\x96\x65\x35\x31\x41\x2b\x31\xa6\xfc\xd3\xcf\x20\xaf\xcb\xbb\xa2\x55\xea\x15\xae\xd3\xd7\x22\x5d\x81\x56\x36\x1e\x45\xd6\x45\xb1\x7e\xc9\x7e\xb6\x14\x73\xa2\x5a\x7f\x33\x56\x39\x9b\xc1\xda\x57\x49\xef\x32\xac\xf2\x54\xd2\x83\xa5\xa5\x23\x7e\x7e\xb2\x4d\x3a\x9f\x03\x60\xc8\xae\x7b\x35\x9b\xc7\x71\x12\x14\xfe\x0a\xfc\xe6\x03\x0c\xf0\x6b\xf6\x0c\x60\x42\x63\x88\x81\xd2\x70\x45\x2e\x88\x01\x42\xbf\x72\x97\x11\x85\x97\xef\x4a\x43\xc8\x63\xbe\x34\x9f\xcd\xbf\x01\x00\x00\xff\xff\x43\x64\xf5\xd8\x51\x09\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\x5d\x6f\xac\x36\x10\x7d\xf7\xaf\x38\xd1\xbe\x6c\xa4\xa8\x79\xe7\x2d\x1f\x4a\x15\xa9\x69\x13\x65\xf3\x54\x55\xd5\x2c\x0c\xe0\xc6\xd8\xd4\x36\x4d\xe8\x55\xfe\xfb\x95\x6d\x60\x61\xc5\xae\xee\x7d\x03\xce\x99\x39\x67\x86\x19\x7b\x83\x7f\x3b\xb6\x92\x1d\x5c\x6d\x3a\x55\x40\x1b\x8f\xc6\x14\xb2\xec\xe1\x6b\x46\xb1\xff\x05\xfc\xc9\x79\xe7\xb9\x80\xd4\x68\xc9\x5a\x52\x8a\x95\xf0\x7d\xcb\x78\xa6\x8a\x1f\x75\x69\xf0\x4d\x00\xde\x78\x52\x19\x1e\xb5\xbf\xc0\x06\xbf\x77\xcd\x9e\x2d\x4c\x89\x96\x2a\x76\xa0\xd2\xb3\x85\xaf\xa5\x83\xd1\x2c\x00\x53\x96\x8e\xfd\xc4\xdf\xd5\x3c\x7c\x0a\x31\x91\x17\x02\x51\x5a\xd3\x44\x2b\xce\x93\xf5\x02\x50\xb2\x91\xcb\xb0\x86\x3e\x83\xfd\x0f\x8d\x96\x6d\x8c\x12\x40\x25\xff\x63\xbd\xe2\x46\x7a\x6e\x5c\x42\xc5\x97\x10\x52\xb7\x9d\x8f\x75\xc4\x1a\x66\xa6\x16\x52\x07\xea\x13\x69\xaa\xd8\x3e\x48\x15\xea\x09\x31\x9a\x1a\xce\xf0\xea\xad\xd4\x95\x00\xb8\x21\xa9\x66\xef\xff\x98\xfd\x4e\x7a\x35\xa7\x74\x9d\x2c\x32\xbc\xbd\x3d\xde\x87\xb6\xb1\xe2\xb6\x36\xfa\x40\x98\xb4\xee\x59\x71\x45\x9e\x67\x62\xcb\x50\x9f\xff\x1d\x5e\xa7\xc4\x3f\x6f\xe5\x8c\xfa\x9d\xe9\xac\x63\x77\xba\x52\xca\x73\x76\x6e\xd7\xb7\x9c\xe1\x66\x7a\x16\xc0\x9e\xf2\xf7\xca\x9a\x4e\x17\x77\x35\xe7\xef\x19\x6e\x8d\x51\x4c\x5a\x00\xad\x95\x39\x67\x78\x50\x86\x42\x87\x49\x29\xf3\xc1\xc5\xce\xdc\x76\xfd\x44\xc3\x06\xaf\xb5\xf9\x80\xd1\xaa\x47\x9e\x5c\xc0\xd7\xe4\xd1\x9b\x0e\x64\x19\xd4\xf9\xda\x58\xf9\x3f\x17\xf0\x06\x7b\x63\xde\xe7\xb6\x9b\x96\x74\x7f\xaa\x67\x8b\x1a\xa6\xa0\xdf\xd8\x39\xa3\x7f\x28\x06\xf0\x54\xb9\x0c\x7f\x06\xf0\xaf\x94\xc2\xb3\x2d\x29\x8f\xfb\x20\x35\x79\x2e\x5e\x3a\xb6\x7d\x4c\xd4\x0e\x2b\x92\x4d\xcb\x72\x50\xfd\xc3\x16\x6c\x6f\x13\x8f\x5c\xce\xba\x90\xba\x9a\x77\xab\x94\xac\xa6\xff\x7b\x11\x02\xe3\xda\x1d\xb2\xcb\x98\x79\xc4\x01\x2a\x1a\xa9\xb7\x07\xf3\x97\x19\x6e\xc2\xa7\x11\x72\xdb\xe0\x27\x79\x19\xb1\xf0\x1c\x62\x9b\x34\xd9\xcb\xe8\x61\xdc\x0f\xf0\x3c\xc3\x15\xca\xd8\xb1\x6c\xb9\x15\x57\x30\xa9\xb0\x6c\xac\xf0\x90\xe8\x39\x6d\x67\x31\x4c\xf6\x4c\xec\xe2\x32\x9b\x06\x7e\xc6\x58\xd7\x5b\x6e\xc6\xaa\xe0\x48\x19\x14\xf3\x34\x16\x47\x82\xc3\xb0\x4c\xb8\x3c\xa1\xb7\x18\xaa\x55\xb9\x81\x31\xa8\xa9\x38\x4f\x47\x62\x69\xc8\x26\x74\x5d\x69\x3e\x89\xab\x42\x89\x30\x55\x15\xb7\xe3\x84\xe7\xd9\xfe\x9e\xf0\x1c\x18\x8b\x54\xdb\x60\x38\x1c\x9a\x13\x1c\x96\xc1\x85\x7c\x6f\x8e\xad\x10\x08\x37\x86\xf3\xf2\x94\xff\x97\x01\x3d\x23\x3b\x52\x06\xe1\x31\xdf\x51\xb7\x46\x56\x3c\xa4\x9c\x5f\x57\xdb\xb1\xf3\x67\x94\x02\x3c\xa8\x84\x1c\x47\x0a\x01\x15\xc0\x26\x5c\x78\x9d\x3a\xd1\xc5\xa7\x88\x9d\x1b\xec\x48\x18\x54\x52\xa6\x23\x9d\xc4\x10\x51\xca\x31\xd9\xbc\x7e\xed\x95\xa2\x7d\xb7\xae\x38\x82\x49\xf3\x12\xd7\xd7\xb8\x09\xc7\xa4\x1b\x82\xa5\xae\x46\xc7\x57\xe3\x28\x81\x74\x91\xda\x14\x55\x72\xf2\x5c\x99\x70\xaf\x6f\x63\xfc\xaf\xec\xc3\x51\x9b\x6e\x56\x1e\x7e\xf5\x8c\x25\xbe\xc4\xf7\x00\x00\x00\xff\xff\x3f\x8b\xc4\x87\x0a\x08\x00\x00")

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

var _type_blog_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x41\x4f\xc3\x30\x0c\x85\xef\xf9\x15\xde\xdf\xc8\x8d\x6d\x42\x54\x42\x08\x31\x7a\x42\x1c\x42\xe3\x75\x81\xac\xae\x52\xe7\x50\xa1\xfe\x77\x14\xb7\x6b\x32\x6d\x07\x6e\xf5\x73\xdf\xf3\xeb\x57\x1e\x7b\x84\xad\xa7\xf6\x21\xf2\x89\x02\xfc\x2a\x80\xa3\x0b\x03\xbf\x98\x33\x6a\x38\x70\x70\x5d\xbb\x51\x00\xde\xdc\x6a\x7d\xa0\xa3\xf3\xf8\xea\x9a\xfa\xed\xf9\xb2\x50\x93\x52\x6b\xaa\xe4\xc5\xe8\xac\x86\xba\xae\xf6\xc9\xc4\x8e\xfd\x55\xca\x17\xd9\xb1\x9c\x1b\xc3\xd8\x52\x18\x35\xec\x96\xa7\xa4\x9e\xd0\x58\x0c\xd5\xd9\xb4\x58\x1c\x4b\x1b\x23\xcd\x75\xf1\x15\x92\xea\xa9\xdd\x92\x1d\xc5\x31\x68\xf8\x48\x6b\x19\x36\x9f\x57\x15\x45\x93\x9e\xdf\x03\x75\xd5\xbe\x8c\x8e\xc1\xe7\x71\x52\xca\x75\x7d\xe4\xec\xaa\x64\xbc\x6f\x65\xfa\xc1\xee\x8e\x79\x17\xd0\x30\x4a\xc4\xea\xfe\x37\x92\xc4\x30\x93\x2c\x90\xbc\x97\xd7\x56\x26\xf9\xf5\x25\xf3\x16\x86\x94\x98\x89\xcc\xfd\xea\xde\x2e\xfd\x9e\x72\x7c\xae\x9a\xfe\xf7\x21\x36\x38\x0c\x1a\xea\xde\x93\xb1\x8f\xb3\x92\xa4\x0b\xf7\xb2\xe6\xf4\x17\x00\x00\xff\xff\xa4\xcd\x6f\xe3\x62\x02\x00\x00")

func type_blog_graphql() ([]byte, error) {
	return bindata_read(
		_type_blog_graphql,
		"type/blog.graphql",
	)
}

var _type_company_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\x4f\x6b\x33\x21\x18\xc4\xef\x7e\x8a\xc9\xed\x7d\x21\x97\xf6\xe8\x2d\x49\x29\x04\xfa\x97\xb2\xa7\xd2\x83\xc4\x27\x22\xec\x3e\x8a\xba\x85\x25\xe4\xbb\x17\x75\x53\x96\xec\xe6\xda\xdb\x38\x8e\xfa\x73\x78\xd2\xe0\x09\x3b\xd7\x79\xc5\x03\x4e\x02\x50\xde\x07\xf7\x4d\x5a\x62\xeb\x5c\x4b\x8a\x05\xd0\xf7\x56\x4b\x34\xcd\xfe\x61\x25\x80\x43\x20\x95\x48\x6f\x92\xc4\x47\x0a\x96\x8d\x00\x58\x75\x74\x59\xe6\x4c\xa7\x58\x19\x0a\xf1\x9f\x57\x86\x24\xde\x94\xa1\x35\x8e\xb6\x4d\x14\x24\x9e\xeb\xe6\x63\x59\xae\xe1\x82\xa6\xb0\x1d\x24\x5e\xab\xf8\xff\x9b\xc8\xc7\xf2\x6d\x4a\xeb\x40\x31\x4a\x6c\xaa\x58\x89\xb3\x10\x85\x7c\x34\x2a\x79\xd5\x4f\x96\xe9\x6e\xca\x32\xf1\xef\xa7\xfe\xc1\xf5\x9c\x86\xa9\xe3\x5d\x4c\x3b\xa7\x69\x96\x0a\x93\xd8\xe5\xe9\xb1\xb4\xcc\x08\xdb\xf9\x96\x3a\xe2\x14\xf3\x57\x2d\xe7\x7e\xde\x7b\x0a\xb5\x52\xd2\x86\xa2\xc4\xe7\x78\xe2\x2b\xbf\xa4\x0c\xed\xf9\xe8\x6a\x35\x59\xe5\x7b\x2d\xfb\x3e\x61\x57\xfa\x1d\xc3\xfb\x62\x9d\x0a\x48\x31\x5e\xae\x9a\xfe\xbb\x5f\x57\xba\xc6\xeb\x25\xba\xab\x11\x99\xb3\xde\x18\xad\x25\xfc\x1b\xf4\x33\xf8\x05\xf6\x39\xba\x38\x8b\x9f\x00\x00\x00\xff\xff\xc7\xdd\x28\x06\xe4\x02\x00\x00")

func type_company_graphql() ([]byte, error) {
	return bindata_read(
		_type_company_graphql,
		"type/company.graphql",
	)
}

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4d\x6f\xe3\x36\x13\xbe\xeb\x57\x4c\xe2\x43\xde\x17\xd8\x1c\xba\xf9\x40\xeb\x5b\x56\xeb\xa2\x06\x52\x6c\x6a\xcb\x68\x8b\x20\x87\x09\x35\x96\x89\xa5\x48\x95\x1c\x26\x11\x16\xfb\xdf\x0b\x92\x8a\x45\x3b\x71\xb1\xe9\x5e\x0a\x03\xd6\x07\xe7\x99\x0f\x3e\xcf\x0c\x45\xda\xb7\x70\x25\x04\x39\x57\xf5\x1d\xc1\x97\x02\xc0\x92\x63\x2b\x05\x53\x5d\x00\x98\x8e\x74\xf1\xb5\x28\xa2\xe1\x92\xad\x17\xec\x2d\xcd\x14\xb5\xa4\x39\x9a\xb7\xa6\xf6\x8a\x0a\x00\x26\xc7\x05\x80\x22\xe7\xcc\x88\x29\x8d\xb7\x8e\xb6\xce\x8d\x56\x52\x07\x6b\xa1\xd0\x39\x6b\x4c\x1b\x2c\xa5\x66\xb2\x6b\x14\x04\xcb\x5e\x29\xbc\xf7\x6e\xce\xd4\x46\x80\xc6\x96\xa6\x21\xb2\xd4\xcd\x51\x01\xe0\xbd\xac\xa7\xb0\x5a\xcd\x3f\x86\x27\x61\xda\x4e\x11\xd3\x14\x3e\x18\xa3\x08\x63\x5c\x0e\xc1\xae\x63\x1a\x20\xc3\x7a\xc8\xd5\x7d\xbf\x6b\x80\xb6\x3b\x5b\x2d\xae\x9f\x31\xb1\xe6\x27\xce\x5d\x30\x36\x6e\x0a\xb7\x15\x36\x47\x77\xdb\x54\xd2\x16\xc4\x90\x21\xc0\x5c\x73\x34\xed\x3b\x9a\x66\xdb\x73\xf4\x4a\x46\xf7\x28\x3e\x37\xd6\x78\x5d\x97\x1b\x12\x9f\xf3\x54\x70\x4b\xda\x34\x23\x30\x80\x3a\x2b\x05\x4d\xe1\x67\x65\x90\x53\x21\xca\xd8\x2c\x67\x7a\x12\x64\x3b\xce\xde\x48\xcd\xd6\xd4\x5e\xb0\x34\x3a\x7b\xbd\x31\x8f\x95\x29\xb7\xdb\x90\xbd\xf7\xd6\xe5\x2b\x31\x54\x01\xf0\xb8\x41\xfe\xd3\xf8\x6b\x42\xab\xa7\x70\x3b\xd4\x71\x17\x35\xf5\x97\x97\x36\x11\xb1\xbb\xe2\x3a\x12\x72\x2d\x45\x45\xb6\x75\x59\x10\x81\x4c\x8d\xb1\xfd\x14\xca\xe1\x2e\x14\xad\x94\x79\xa4\xba\x32\x1f\x7c\x9f\x6f\x86\x1b\xa8\x0d\xae\x33\x96\x5f\x70\x70\x83\x0d\xe5\x92\xb8\xc1\x46\x6a\x64\xaa\x7f\xf3\x64\xfb\xc8\x10\xd5\x0d\x05\x3f\x09\x10\x32\xec\xb0\xa1\xb9\x5e\x9b\x69\x30\x8f\x77\x49\xb2\x9d\x67\x58\xe2\x03\x95\xcf\x4a\x4e\x90\x79\x5c\xc8\xc8\xde\xe3\x35\xab\x2d\x48\x2d\x09\xee\x2d\xb4\xbc\x59\x14\x7b\x9a\x78\x45\x12\x83\x6c\xa3\xf4\xff\x89\x94\x7b\xd4\x9a\xec\xbc\xc5\x86\x96\x3e\xfa\xcf\x9d\x78\x36\x76\x2c\x09\x26\xe9\x62\xd6\xc0\x1b\x82\x2a\xac\xc2\xff\xb6\x6d\x0f\x22\x75\x85\xd1\xaa\xff\x7f\x68\x2d\x7c\xba\x41\xcb\x52\xc8\x0e\xa3\x48\xd2\xce\x39\x46\xcb\x1f\x31\xa8\xac\x92\x6d\x28\x85\x74\xbd\xf3\xac\x8c\xc0\x9d\x1d\xda\x21\xe7\x53\x9c\x37\xdf\xc7\xcc\x04\x1c\x09\xa3\x6b\xb4\x7d\xf9\x0d\xcc\xc1\x04\x96\x1b\x63\x19\x6a\x72\xc2\xca\x2e\x24\x17\x76\x21\x15\x7c\x80\x57\x98\xc0\xc9\xd5\xbd\xf1\x0c\xbc\x91\x6e\xb0\x3d\x09\x81\x83\xd5\xa1\x66\x0c\xa8\x5f\xcc\x23\xb0\xd9\x8e\xab\xc3\xf8\xd7\x9a\x16\x26\x30\x73\x2c\xdb\xd0\x03\xc9\xe4\x90\xaf\xc3\xed\x1d\x92\xf8\x7d\x83\x0c\xbd\xf1\x27\x4a\x81\x0a\xeb\x27\x70\xef\x95\x22\x86\xce\x48\xcd\xee\xf0\x08\x08\xe8\x45\xb6\xf4\x12\xf8\x5f\xd5\xbb\x7b\x3e\x0d\xb7\xd3\x62\x9c\x39\x49\x80\xe3\xdb\xa8\xba\x34\xf1\xf7\x0f\xd1\xfd\x93\x67\x44\x5b\x42\xa6\x0a\x9b\x51\xb7\xfb\xe7\xc3\x4e\x5d\x47\xdb\x61\x57\x61\x13\xcd\x47\xb7\xdf\x8e\x7d\x56\xf8\xbf\x71\x90\xe7\xfd\xec\xe7\x2d\xc9\x27\xfc\x8d\xb7\x62\x83\x6e\x68\x5a\x37\x3a\x48\x42\x0c\xc4\x85\x23\xf4\x2e\x6e\x9c\x23\x3b\x32\x79\x04\x13\xb8\x96\x8e\x43\xb7\xd5\xa4\xa8\x41\x26\x17\xd3\x77\xb1\x4d\xd9\xe2\x5c\x3f\x18\x29\x68\xd6\xa2\x54\x19\x97\x41\x48\x1d\x53\x3d\x08\x60\x50\xd8\x2b\x47\x70\x69\xf4\x5a\xda\xf6\xe5\xf7\xc6\x5e\xd6\x0b\x72\x9d\xd1\xc3\xa9\xcf\x16\xb5\xc3\xd8\x89\xe5\xfe\x57\x45\xc8\x99\xad\x27\x90\x69\x4e\x86\xde\x43\xdd\x83\x74\x60\x34\x20\x08\xa3\xd9\xa2\x60\x40\x5d\x03\x6f\xbc\x83\xda\x90\xd3\x27\x0c\x9a\xa8\x0e\xdd\xda\x61\x0f\xbe\x5b\x5b\x33\x0c\x4c\x2b\x3b\x2a\x95\x24\xcd\x4b\x12\x96\xf2\xb9\x34\x5f\xbf\x96\x4b\x88\xb5\x46\xe5\xe8\x5d\x6a\x78\xe9\xa0\x91\x0f\xa4\xc1\x99\xb8\xc1\x20\x50\x87\x28\xa1\xd6\x09\xdc\xc2\xec\x8f\xab\x5f\x6f\xae\x67\x50\x7e\x5a\x2d\x96\x33\x58\x56\x8b\x55\x59\xad\x16\xb3\x62\x02\x00\x5f\x92\xca\x8f\xd3\x67\xe1\xf1\xbb\x41\x43\xc7\x67\xe7\x67\xe7\x67\xa7\xf9\xff\x79\xbc\x1c\x7f\x7d\xb7\x83\x4b\x9f\x8f\x23\xee\x87\xf7\xf1\x77\x9a\x2e\xef\x4f\xb3\xa7\x7d\x68\xf8\x00\x1d\x81\x3f\xfe\x94\x7e\xa7\x17\x97\xf1\x77\x7a\x71\xb9\xbd\xb9\xb8\xbc\x48\xe0\xbb\xe2\xef\x00\x00\x00\xff\xff\x49\x7c\x37\x5a\x02\x0b\x00\x00")

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

var _type_lesson_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcf\xb1\x0a\x83\x30\x10\x06\xe0\x3d\x4f\x71\xbe\x46\xd6\xba\x08\x1d\x2c\xad\x53\x71\x08\xe4\xef\x11\xd0\x18\xcc\x05\x2a\xe2\xbb\x17\x23\xb4\xa5\x15\xb7\xe4\x72\xf9\xee\x7e\x99\x02\xe8\x8c\x18\x07\x4f\xb3\x22\x4a\xc9\x59\x4d\x4d\x53\x95\x85\x22\x12\x27\x1d\x34\x5d\x65\x74\x9e\x73\xc1\x70\xd4\x74\xbf\x19\x2e\xda\x7c\xc7\x53\x3e\xef\x8b\x52\x5f\x5e\x6d\x18\xe4\xfa\xd0\xa1\x87\x97\x48\xb5\x61\xe7\x8d\xc0\x5e\x12\xc6\x29\x4f\x83\x65\xac\xde\xf6\xa1\x55\x44\xc1\x30\x2a\xff\x18\xf4\xda\x9e\x4f\xab\xea\x7c\x48\x42\xa7\x11\x46\xb0\xf5\x56\xb9\x32\x1f\xec\x98\x33\xb4\x7b\x3b\x6e\x5a\x13\xec\x8e\x76\x90\xff\x87\xfa\x9b\xf4\x96\x4b\x74\x38\x96\x97\x57\x00\x00\x00\xff\xff\x9d\x86\x38\x67\x77\x01\x00\x00")

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

var _type_module_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\xcb\x6a\xc3\x30\x10\x45\xf7\xfa\x8a\xf1\x6f\x78\xd7\xd7\x42\xd0\x42\xa9\xe3\x76\x51\xb2\x50\xac\x21\x88\x4a\x23\xa3\x47\x20\x94\xfc\x7b\x91\x94\xd8\x72\x70\x4b\xbb\xea\xca\x96\xe6\xc1\xbd\xe7\x2a\x1c\x47\x84\x27\x2b\xa3\x46\x50\x66\xd4\x68\x90\x82\x87\xee\xa8\xb5\xd8\x45\xcf\x03\x1a\xf8\x64\x00\x31\x2a\xd9\x42\xdf\xf3\xfb\x86\x01\x90\x30\xd8\x42\x17\x9c\xa2\x7d\x3a\xef\x04\x11\x3a\x6e\xc4\x1e\xfb\x97\xc7\x4b\x85\x01\x48\xf4\x83\x53\x63\x50\x96\xea\xfe\xe0\x04\x95\x42\x7d\x7b\xb0\x6a\x40\x7b\x40\xb7\xdc\x71\x50\x12\x6d\x0b\xaf\xe9\xc3\x00\xfc\x59\x5a\x0b\xef\xb5\xca\x66\xcb\x00\x06\x9b\x2c\x04\x6c\xe1\xd6\x5a\x8d\x82\xd8\x89\x31\xa4\x68\xca\xf4\x26\x99\x4d\x6e\xde\x78\xb7\xe1\x37\xa9\x98\xfd\xe7\x62\x2e\xa4\x63\x3b\x37\x27\x55\xd1\xe9\x59\xe4\x89\x31\x45\x63\x0c\xa5\x85\xe7\xdf\x5f\xce\x65\x19\x85\xf4\x43\xc1\x5c\x26\xd1\x07\x06\xa0\xd1\x7b\x4b\xf3\xfe\xd2\x38\xe1\x2f\xfb\x17\xd3\xcd\x55\x2a\xd3\xe8\x9d\x43\x11\xf0\xbc\x60\x52\x78\x1d\x59\x10\xfb\x84\x30\xcf\x6e\xff\x9a\x54\x95\x77\x17\x87\x01\xbd\xdf\xd8\x0f\xa4\x3a\xb4\x4b\x96\xdf\xd5\xab\x50\xb3\xc8\x65\xb2\xb3\xfd\x66\x3b\xc5\x54\x1b\x7b\x16\x47\x6d\x85\xcc\xd6\x4c\xbe\xb9\xe0\x99\x41\xf4\xa3\x5c\x01\xb1\x7c\xc9\xd7\x18\x6a\x4c\xeb\x54\x56\xa1\xfc\x1b\x93\xda\xe3\x8f\x4c\xbe\x02\x00\x00\xff\xff\xd6\x8d\x98\x78\xe9\x03\x00\x00")

func type_module_graphql() ([]byte, error) {
	return bindata_read(
		_type_module_graphql,
		"type/module.graphql",
	)
}

var _type_question_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\xdd\x4e\xc3\x30\x0c\x85\xef\xf3\x14\xae\x78\x8b\xde\xb1\x32\xa0\xd2\x04\x1b\x5b\xaf\x50\x85\xac\xd5\x8b\x22\xda\xa4\xe4\x47\x50\x4d\x7b\x77\x94\x94\x6e\xeb\xcf\x24\x06\x77\x8d\x65\xfb\x9c\xf3\x25\xb5\x4d\x4d\xb0\x72\x64\xac\x50\x12\xf6\x0c\xc0\x39\x51\xc4\x90\x65\xe9\x5d\xc4\x00\x2c\x7d\xd9\x18\xd6\x56\x0b\xc9\xfd\x59\xa3\x2c\x54\x25\x0c\xdd\x4a\xf3\x49\xda\xc4\x30\x53\xaa\x24\x94\x0c\xe0\xe3\x67\xcd\xa6\xa9\x29\x3e\x2e\xf5\x27\x3f\x89\xdd\xc0\x6b\x3b\x1a\xe5\xec\xc0\x58\x4f\x7f\x89\x9c\x40\x54\x75\x49\x15\x49\x6b\x60\x89\x5c\x48\xb4\x54\xac\x1c\xe9\x26\xb8\xa3\x82\x93\xdf\xd1\x8d\xe4\x0c\xa0\x46\x4e\xa9\xdc\xa9\xd8\x0f\x84\xaf\xe3\xe6\x56\x6a\x22\x97\x30\x89\xd2\x9a\xb6\xf6\x3c\xc0\x79\x58\xdf\x53\x21\xa7\xec\x65\x71\x2c\x1d\x18\x23\xe9\xaa\x5e\xb4\xb0\x7b\x9d\x3e\x3d\x2c\xe6\x6f\xc9\xe3\x73\x9a\xcc\x7d\x9b\x90\xb5\xb3\x90\x68\x42\x4b\x5d\x77\x1a\x6a\xfb\x6b\xa0\x46\x57\x50\x6d\xc5\x66\x68\xc4\xb6\x5d\x13\xf4\xa2\x3c\xdc\x22\x72\xdf\x12\xc2\xe7\xd1\xd0\xe0\x70\x26\x78\x1c\x03\x8a\xe0\x06\x0a\xda\xa1\x2b\xad\x01\xab\x60\x87\xa5\xa1\x0b\xd4\x36\xea\x9d\xe4\x39\xb7\x70\x1b\x7d\x1e\x4b\x6c\x4a\x85\x45\x50\xeb\x42\x9e\x02\x9e\x4c\x66\x75\x31\x49\xf1\xf2\x43\xfd\xf3\x3b\xed\x01\x6d\x75\xc7\x40\x87\x3c\x87\x4e\x27\x71\x9e\xcc\x4e\xb2\xfd\x37\xda\x3e\xa4\xdf\xa2\xed\x2a\xf7\xa2\xb4\xa3\xdf\x64\xac\xbf\x19\xc4\xfe\x0e\x00\x00\xff\xff\xb9\x45\xf1\xb2\x3d\x04\x00\x00")

func type_question_graphql() ([]byte, error) {
	return bindata_read(
		_type_question_graphql,
		"type/question.graphql",
	)
}

var _type_test_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xc1\x8e\xd3\x30\x10\xbd\xe7\x2b\x26\xea\x5f\xe4\xb6\x80\x56\xea\x89\xae\xda\x1e\xd0\xaa\x07\x37\x7e\x9b\xb5\x70\xec\xe0\x19\xb3\xaa\x10\xff\x8e\xec\xa4\x71\x5a\x08\x02\xc1\xa9\x63\x77\xde\xcb\x9b\x37\x2f\x91\xcb\x00\x3a\x80\x85\x4c\x3f\x58\xf4\x70\xc2\xb4\xbf\x58\xab\xce\x91\xb7\x82\x9e\xbe\x55\x44\x4e\xf5\x68\x68\x2f\xc1\xb8\xae\xae\x88\x62\x34\xba\xa1\xe3\x71\xfb\x21\x9d\x5a\x9f\xa0\x82\x86\xde\x79\x6f\xa1\x5c\x45\x24\xaa\xe3\x86\x9e\x0f\xaa\xab\x4f\x15\x91\x12\x41\x3f\x08\x3f\x58\xeb\xdf\xa0\x1b\xda\x3a\xa9\x88\x06\xc5\xbc\x43\x68\xe1\x44\x75\x68\xe8\xd1\x7a\x95\xee\xbf\x44\xb0\x18\xef\xf8\xe0\x1f\x1c\xbf\x21\x5c\x01\x41\x39\xed\x7b\xc3\x18\xaf\x79\xf9\xc8\x19\xd4\xd0\xf3\xd3\x54\xd7\xa7\xea\x7b\x55\xcd\x43\xee\x54\x87\xe5\xa0\x3b\xd5\x19\xa7\x04\xfa\x29\x22\x5c\xf2\xa8\xd0\x1d\xb2\x72\xb0\x9c\xb2\xc4\x0e\x5b\xf7\xe2\x9b\xd4\x9c\xab\xc4\x68\xdc\x10\x25\x53\x3e\x1a\x2b\x08\x19\x5a\x5c\xb9\xb3\x6c\xd6\xf0\x3e\x40\x09\x46\x25\x17\xeb\x95\xce\x38\x01\x4b\x93\xc9\x0a\x75\xe9\xdc\xe6\xf3\xaf\xb6\x30\x79\x9c\xd7\xb0\x66\x72\xbd\xe6\x72\xbd\x6a\x73\xfd\x3b\x9f\x6b\xda\xd0\xc7\xaf\x08\xc1\x68\x30\x0d\x08\x33\xc9\x8c\x51\xe9\x74\xb7\x8f\x51\x63\x06\x07\x8d\x00\x4d\xd6\xb0\x90\x7f\x21\x79\x45\xe9\x24\xe3\x48\x5e\x0d\x67\x4b\x8a\x19\xd7\x6d\x8e\x62\xb2\x17\x57\x48\x22\x9e\x92\x48\x9b\xfc\x7b\x4f\x4a\x67\x18\xd7\x91\xca\x58\xe8\xe4\x53\x2e\xd7\x91\x0c\x8b\x56\xa0\xa7\xc6\xa2\x63\x1f\xcf\xbd\x91\xdb\xa5\xb4\x3e\x06\x46\xa2\x4a\xd6\xd1\x86\x0a\xcf\xf8\x57\x2e\x23\x23\x50\x9a\x4b\x7d\x4e\x62\xd2\x95\xe4\x97\xce\x4d\xfb\x5f\x57\x93\xfb\x2e\x3e\x92\x0a\x98\x14\xa5\x48\x5d\xc7\x58\xa6\x7d\xf4\x67\xf4\xf9\x93\x8f\xe1\xda\x42\xe2\x67\xaa\x39\x8b\x65\x98\x65\x16\x39\xb6\x2d\x78\xb1\xef\x32\xfd\x71\xd0\x3f\x45\xf2\xf6\x43\x70\x93\xf9\x3f\xcb\xe7\x7f\xfd\x08\xfc\x4b\x36\xff\x3a\x9a\xd9\xc4\xe2\xc9\xea\x0b\xfd\x23\x00\x00\xff\xff\xac\x4d\xf0\xb4\x62\x05\x00\x00")

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
	"mutation.graphql":        mutation_graphql,
	"query.graphql":           query_graphql,
	"schema.graphql":          schema_graphql,
	"type/admin.graphql":      type_admin_graphql,
	"type/auth.graphql":       type_auth_graphql,
	"type/company.graphql":    type_company_graphql,
	"type/course.graphql":     type_course_graphql,
	"type/delegate.graphql":   type_delegate_graphql,
	"type/individual.graphql": type_individual_graphql,
	"type/lesson.graphql":     type_lesson_graphql,
	"type/manager.graphql":    type_manager_graphql,
	"type/module.graphql":     type_module_graphql,
	"type/question.graphql":   type_question_graphql,
	"type/test.graphql":       type_test_graphql,
	"type/users.graphql":      type_users_graphql,
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
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"mutation.graphql": &_bintree_t{mutation_graphql, map[string]*_bintree_t{}},
	"query.graphql":    &_bintree_t{query_graphql, map[string]*_bintree_t{}},
	"schema.graphql":   &_bintree_t{schema_graphql, map[string]*_bintree_t{}},
	"type": &_bintree_t{nil, map[string]*_bintree_t{
		"admin.graphql":      &_bintree_t{type_admin_graphql, map[string]*_bintree_t{}},
		"auth.graphql":       &_bintree_t{type_auth_graphql, map[string]*_bintree_t{}},
		"company.graphql":    &_bintree_t{type_company_graphql, map[string]*_bintree_t{}},
		"course.graphql":     &_bintree_t{type_course_graphql, map[string]*_bintree_t{}},
		"delegate.graphql":   &_bintree_t{type_delegate_graphql, map[string]*_bintree_t{}},
		"individual.graphql": &_bintree_t{type_individual_graphql, map[string]*_bintree_t{}},
		"lesson.graphql":     &_bintree_t{type_lesson_graphql, map[string]*_bintree_t{}},
		"manager.graphql":    &_bintree_t{type_manager_graphql, map[string]*_bintree_t{}},
		"module.graphql":     &_bintree_t{type_module_graphql, map[string]*_bintree_t{}},
		"question.graphql":   &_bintree_t{type_question_graphql, map[string]*_bintree_t{}},
		"test.graphql":       &_bintree_t{type_test_graphql, map[string]*_bintree_t{}},
		"users.graphql":      &_bintree_t{type_users_graphql, map[string]*_bintree_t{}},
	}},
}}
