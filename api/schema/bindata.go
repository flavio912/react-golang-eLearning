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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4d\x8f\xdb\x36\x10\xbd\xeb\x57\x8c\xe1\xcb\xe6\xd0\xf4\xee\x5b\x6a\x23\x80\x81\x5d\xd8\x8d\xa5\x1f\x30\x11\xc7\x32\x11\x8a\x54\xf9\xb1\xa9\x51\xec\x7f\x2f\x48\x49\xfc\xb2\xdb\x4b\x8e\x7c\x9a\xf7\xe6\x71\x86\x33\xf6\x16\x46\x67\xd1\x72\x25\x0d\xf4\x28\x61\x54\x8c\x5f\xef\x60\x6f\x04\xec\xfb\x67\xa0\xbf\xa9\x77\x96\x18\x18\xfa\xcb\x91\xb4\x1c\x85\xb8\x7f\x6e\x1a\x2e\x27\x67\xa1\x9b\x84\x42\xf6\x95\x0b\x7a\x23\x8b\xf0\x4f\x03\x70\xe5\x82\xda\xfb\x44\x3b\xb8\x58\xcd\xe5\xb0\x69\x00\x7a\x25\x2d\x49\xfb\x4a\x72\xb0\xb7\x1d\x1c\xa5\xdd\x34\x1f\x4d\x63\xef\x13\x65\x12\xdf\xc8\x4c\x41\xc2\x69\x91\xb3\x8d\xeb\x7b\x32\xa6\x55\x3f\x48\x26\xfc\xe3\xd1\xc3\x65\x0e\x0c\x1a\xff\x49\x0a\x49\xdf\x96\x2b\x87\x50\x64\x23\x97\xaf\x6a\xe0\xf2\x25\x28\xee\xe0\x4b\x44\x36\x9f\x76\xf0\xc5\xd9\x5b\xd0\x69\x00\x46\x94\x38\x90\x2e\xa2\xdf\x32\xac\x8e\x67\x24\x68\x40\x4b\x05\xe1\x90\x83\x25\xc3\x17\x4b\x13\x5a\x5a\x44\x57\xca\x3e\x07\x8f\x1e\xf3\xbc\xe5\xec\x4b\x36\xb1\x47\x52\x97\x83\x4f\x48\x93\x56\xbe\x5b\xc7\x11\x07\x9a\xab\xf8\xcd\x37\xd9\xd8\x24\x90\xb7\xd7\x93\xcb\x6e\xd5\x89\xcf\x99\xe0\xa3\xc6\xd2\x9e\xd2\x83\x2f\xd0\x83\xf1\x43\x0e\x46\xe3\x7f\x28\x25\x08\xe5\x26\x15\x29\xf4\xa9\x2c\x51\x80\x22\x25\x9c\xa2\xcb\x22\xbc\x4b\xd0\x43\xf8\x6c\xaa\x08\x3f\x24\xe8\x7f\x0c\xad\x9d\x2d\x3d\xad\x68\x24\x96\xb0\xaf\xa4\x92\x86\x1a\x80\xed\xe2\x74\xfd\x94\x94\xf7\x6a\x9c\x50\xde\x4b\xe1\x05\x4c\xba\xf3\x39\x5e\xb8\x22\x75\x39\xf8\x84\xb4\x5d\x2e\x9e\x90\x22\xf7\xfa\x38\x1a\x00\x3f\xd2\x01\x7b\xea\x24\x04\x2c\x93\xf2\xf4\xe9\x86\x00\x4d\x3d\x4e\xb6\xbf\x61\x3e\xeb\x79\x55\x01\x70\x9a\xb4\x7a\x8f\xd7\x70\x8e\xb3\x1d\x74\xdd\xf1\x90\xfb\xf6\x3b\xc7\x89\x2b\x17\x67\x92\x8c\xcb\xe1\xa4\x19\xe9\x97\x5e\x70\x92\xf6\x42\xbd\x26\x1b\x13\x94\xea\x06\xdf\xe9\x24\x05\x97\xb4\x57\x4e\x9b\xd8\xb4\x4b\x85\x67\x95\xf2\x47\xd8\x42\x7b\x3a\x9c\x00\x85\x50\x3f\x41\x70\xf9\x63\xd1\xda\x0b\x34\x46\x2b\x35\x3e\xca\x55\x9f\x9e\x28\x2a\xa6\x00\x19\x03\xeb\xac\xd2\x60\x15\xf4\xe1\x93\x9f\x52\xa7\xfb\x1b\x9a\xc5\x8c\x59\x65\xcf\x25\x1c\x25\x2b\x3c\x3e\xae\xd8\xcd\x16\x87\xf2\x15\xb5\x38\x44\x76\x8b\x43\x6a\x3b\x5a\x1a\x94\xae\xdf\xdc\x82\xa6\x3b\x2c\x40\xca\xf0\x4a\xc6\xa8\x6a\x2e\x67\x2c\x92\xe6\x63\x64\x1c\x25\xe3\xef\x9c\x39\x14\x25\x2b\xe1\xd5\xec\xa4\x0f\xd9\xf4\x18\xf7\x7d\xe4\xb6\xcd\x16\xd8\x25\x22\x91\x9f\xa0\x33\xde\xfd\x6a\x82\xed\x82\x19\xc0\xb8\xaf\x7f\xe7\x31\x81\x01\x94\xe6\x27\x69\xe3\xbb\x82\x60\xc9\xd8\x54\xcb\x2c\xd7\x3e\x22\x95\xd7\x2c\x57\x9c\xcc\xb6\xd8\xb2\xac\x26\x26\x68\x25\xfa\x69\x08\x36\x7e\x69\x5b\xcf\xae\xff\xf4\x2c\x5e\x77\x68\x45\x2b\xf7\x2b\x5c\xdf\xa0\x16\xe9\x0a\xb4\xba\xc9\xa3\xc8\xba\x6b\xd6\x2f\xd9\x2f\x9f\x62\x4e\x54\x1b\x74\xc6\x2a\x67\x33\x58\xfb\x2a\xe9\x5d\x86\x55\x9e\x4a\x7a\xb0\xb4\x34\xd5\x8f\x60\xb6\x8c\xe7\x73\x00\x0c\xd9\x75\x35\x67\x23\x3d\x4e\x82\xc2\xbf\x89\xdf\x7c\x80\x01\x7e\xcd\x5e\x12\x4c\x68\x0c\x31\x50\x1a\xae\xc8\x05\x31\x40\xe8\x57\xee\x32\xe5\xf0\xf2\x55\x69\x08\x79\xcc\xa7\xe6\xa3\xf9\x37\x00\x00\xff\xff\x6b\x48\x61\xf8\x94\x09\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x96\xcf\x6e\xdb\x30\x0c\xc6\xef\x7a\x8a\xaf\xc8\x25\x05\x8a\xf5\xee\x5b\xd3\xa2\x43\x81\x75\x6b\xd1\xf4\x34\x0c\x83\x62\xd3\xb6\x56\x59\xf2\x24\x79\xad\x37\xf4\xdd\x07\xfd\xb1\x63\x07\x4e\xb0\xdd\x62\xf3\x23\x7f\x24\x4d\x4a\x59\xe1\x67\x47\x46\x90\x85\xad\x75\x27\x0b\x28\xed\xd0\xe8\x42\x94\x3d\x5c\x4d\x28\x76\x1f\x40\x6f\x94\x77\x8e\x0a\x08\x85\x96\x1b\xc3\xa5\x24\xc9\x5c\xdf\x12\x1e\x78\x45\x77\xaa\xd4\xf8\xc3\x00\xa7\x1d\x97\x19\xee\x94\x3b\xc3\x0a\x9f\xbb\x66\x47\x06\xba\x44\xcb\x2b\xb2\xe0\xa5\x23\x03\x57\x0b\x0b\xad\x88\x01\xba\x2c\x2d\xb9\x51\xbf\xad\x29\xbd\xf2\x3e\x41\xe7\x1d\x51\x1a\xdd\x84\x54\xac\xe3\xc6\x31\x40\x8a\x46\xcc\xdd\x1a\xfe\xe6\xd3\x7f\x55\x68\xc9\x04\x2f\x06\x54\xe2\x17\xa9\x85\x6c\x84\xa3\xc6\x46\x2b\x7b\x67\x4c\xa8\xb6\x73\xa1\x8e\x50\xc3\x24\xa9\x19\x6a\x2f\xbd\xe7\x8a\x57\x64\x6e\x85\xf4\xf5\x78\x1f\xc5\x1b\xca\xf0\xe4\x8c\x50\x15\x03\xa8\xe1\x42\x4e\x9e\x7f\xe8\xdd\x56\x38\x39\x95\x74\x9d\x28\x32\x3c\x3f\xdf\xdd\xf8\xb6\x91\xa4\xb6\xd6\x6a\x2f\x18\x59\x37\x24\xa9\xe2\x8e\x26\xb0\xb9\xab\xcb\xbf\xfb\xc7\x31\xf0\xff\xa7\x72\x82\x7e\xad\x3b\x63\xc9\x1e\xaf\x94\xe7\x39\x59\xbb\xed\x5b\xca\x70\x35\xfe\x66\xc0\x8e\xe7\x2f\x95\xd1\x9d\x2a\xae\x6b\xca\x5f\x32\x6c\xb4\x96\xc4\x15\x03\x5a\x23\x72\xca\x70\x2b\x35\xf7\x1d\xe6\x52\xea\x57\x2a\xb6\x7a\xd3\xf5\xa3\x0c\x2b\x3c\xd5\xfa\x15\x5a\xc9\x1e\x79\xcc\x02\xae\xe6\x0e\xbd\xee\xc0\x0d\x81\x77\xae\xd6\x46\xfc\xa6\x02\x4e\x63\xa7\xf5\xcb\x34\xed\xa6\xe5\xaa\x3f\xd6\xb3\x59\x0d\xa3\xd3\x27\xb2\x56\xab\x7f\xf2\x01\x1c\xaf\x6c\x86\xaf\xde\xf8\x2d\x86\x70\x64\x4a\x9e\x87\x7d\x10\x8a\x3b\x2a\x1e\x3b\x32\x7d\x08\xd4\xa6\x15\xc9\xc6\x65\xd9\x53\xbf\x98\x82\xcc\x26\xea\xb8\xcd\x49\x15\x42\x55\xd3\x6e\x95\x82\xe4\xf8\x7d\xcf\xbc\x63\x58\xbb\x7d\x74\x11\x22\x0f\x76\x80\x17\x8d\x50\xeb\x7d\xf2\xe7\x19\xae\xfc\xab\xc1\x64\xd7\x3e\x9f\x98\xcb\x60\xf3\xbf\xbd\x6f\x13\x27\x7b\xee\x9d\xc6\x7d\x6f\x9e\x46\xb8\x40\x19\x3a\x96\xcd\xb7\xe2\x02\x3a\x16\x96\x0d\x15\xee\x03\x3d\xc4\xed\x2c\xd2\x64\x4f\x60\x67\xe7\xd9\x38\xf0\x13\xc5\x32\x6f\xbe\x19\x8b\xc0\x41\x92\x88\x79\x1c\x8b\x03\x60\x1a\x96\xd1\x2e\x8e\xf0\x66\x43\xb5\x88\x4b\x8a\x44\x93\x61\x9e\x0e\x60\x71\xc8\x46\xeb\x32\x69\x3a\x89\x8b\xa0\x28\x48\x9c\x9d\xd4\xd5\x01\x65\x23\x75\x95\x2c\x73\xc2\x42\x2c\xaf\x1d\xfb\x13\xf6\xec\x48\xf5\x93\x93\xe0\x48\xf5\x5e\x31\x0b\xb5\xf6\x49\xf9\xe3\x77\x34\xfb\xb5\xb2\x3e\xde\xb3\x25\xc3\x18\xfc\xdd\x63\x9d\x38\xd6\x89\xc7\x64\x3d\x81\x1d\x24\x09\x3c\xc4\x3b\xe8\xc8\xa0\x0a\xc7\x9d\x75\xcb\xb4\x2d\x59\x77\x82\xe4\xcd\x89\xe2\x63\x1c\x10\xbc\x95\x01\x2b\x7f\x75\x76\xf2\x48\x17\xef\x83\xed\xd4\x8a\x04\x41\xa2\xc4\x48\x07\x9c\xa8\x60\x01\x65\x89\x9b\xbc\x7e\xea\xa5\xe4\xbb\x6e\x99\x38\x18\x23\xf3\x1c\x97\x97\xb8\xf2\x07\xae\x4d\xce\x42\x55\x43\xc6\x17\xc3\x50\x82\xab\x22\xb6\x29\x50\x72\xee\xa8\xd2\xfe\x1f\xc2\x3a\xf8\x7f\x24\xe7\x0f\xed\x78\x47\x53\xfa\xd4\x13\x15\x7b\x67\x7f\x03\x00\x00\xff\xff\x62\x43\x73\xa8\x54\x08\x00\x00")

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

var _type_blog_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x4d\x6e\xc3\x20\x10\x85\xf7\x3e\x05\xb9\x06\xbb\xfc\xa8\xaa\xa5\xaa\x4a\x9b\x7a\x55\x65\x41\xcd\x84\xd0\x62\xb0\xf0\x78\x61\x55\xb9\x7b\x35\xe3\xc4\xe0\x26\x8a\xda\x9d\xe7\xc1\x3c\xbe\x79\x60\x1c\x5a\x10\x2b\x17\xcc\xb2\xc7\x63\x88\xe2\xbb\x10\xe2\x60\x63\x87\xcf\xaa\x01\x29\x76\x18\xad\x37\x8b\x42\x08\xa7\xae\xb5\x36\x86\x83\x75\xb0\xb5\x75\xf5\xfa\x74\x59\x28\x4e\x45\x31\xb9\xb2\x5f\x1d\x41\x21\xe8\x25\xe6\xbd\x7d\xab\xe7\x22\x69\xbd\xd5\x52\x54\x55\xb9\xa1\x1d\x68\xd1\xcd\x8e\xfb\x08\x7a\xc8\xeb\x5a\x21\x98\x10\x07\x29\xd6\xe7\x2f\x52\x8f\xa0\x34\xc4\xb2\x51\x06\x32\x2a\x5a\x51\x3c\xa2\xcc\xc6\x65\x57\x17\xcc\x2a\xe8\x81\x3b\x3a\x29\xde\x69\x99\x8b\xc5\x7e\x36\x0b\x6b\x3c\xd0\x67\x17\x7c\xb9\x99\x4d\x13\x5d\x2a\xf3\xa6\x2d\xf5\xd8\xa6\x75\xd0\x80\xc7\x4e\x6c\x95\xb1\x9e\x06\x7f\xe9\x21\x0e\xec\x06\x3a\x9d\xbb\xa7\x54\x95\x81\xd2\x1f\x82\xa4\xcd\xfc\x45\x8e\xd6\xb7\x3d\x26\x8e\x92\xcb\xdb\x30\x18\xbe\xc0\xcf\x70\xc6\xe6\x35\xdf\x03\x5b\x4c\xdd\x7f\x0e\x99\x6e\x25\xdd\x4d\x16\xf2\x5b\x7e\xda\x94\x72\xda\x7e\xf6\xbc\x8e\x97\x21\xc6\x8c\x47\xbe\x8a\x9f\xc4\x9c\xef\xce\x93\xf8\x05\x7b\x93\xf5\x3e\xea\xbf\xb9\x1e\x93\x57\x42\xa4\x5f\x60\xd7\xd7\xd0\x75\x52\x54\xad\x0b\x4a\x3f\x8c\x0a\x49\x97\x17\x96\xc7\x77\xfa\x09\x00\x00\xff\xff\x2e\x3c\xf8\x1b\x75\x03\x00\x00")

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

var _type_delegate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\xc1\x4e\xf3\x30\x10\x84\xef\x79\x8a\xed\xed\xff\x5f\xc1\xb7\xd2\x5e\x22\x21\x54\x20\x3d\x21\x84\x96\x7a\x6a\x0c\x8e\x6d\xd9\x0e\x28\xaa\xfa\xee\xc8\x49\xd3\xa4\x69\x2f\xdc\x92\xd1\xac\x77\xe7\xdb\x4d\xad\x07\xad\x61\xa0\x38\x81\x0e\x05\x51\xd3\x68\x29\x68\xbb\x2d\xd7\x8b\x82\x68\x17\xc0\x09\x72\x99\x04\x3d\xa7\xa0\xad\x2a\x88\xaa\x6a\xf5\x56\xae\x07\x21\xbb\x50\xb3\x36\x13\xc7\x5e\x87\x98\x1e\xb8\xc6\xd4\x64\xf8\x5a\x4b\x30\xf0\x1f\xce\x62\x52\x9c\x7d\xf7\x4e\x69\x3b\x35\x7e\xba\xf7\x4a\x27\x73\x51\xbc\x73\xb5\x67\xdb\x0a\x5a\xf5\x1f\x59\xf3\xc1\xed\xb5\x41\x59\xb3\xc2\x36\x4c\x67\xe2\x5d\xd2\xdf\x3a\xb5\xff\x3c\x2b\x08\xda\xb0\xc2\x7f\x41\xcb\x93\x9a\x7f\x8b\x63\x51\x68\xeb\x9b\x44\xab\x2e\xf5\x80\xa5\xec\xb4\xc3\xd8\x31\xc3\xe9\x11\xfd\x25\xeb\x1c\xd2\xad\x48\xb7\x78\x28\x58\x04\x4e\xd8\x70\x8c\x3f\x2e\x48\x41\x77\xce\x19\xb0\x9d\xc7\xf5\xc6\xb1\xac\xdc\x17\xce\xe4\x72\xa2\x6e\xc3\x97\x81\x9e\x10\xbd\xb3\xb1\xdf\xb7\x3c\x89\xe2\x7c\x06\x8b\x49\x53\x39\x76\x9d\x3d\x39\xb8\x33\x39\xd2\xb5\x37\xa8\x61\x53\xcc\x60\xb5\xcd\x95\x8f\x0d\x42\xdb\xb5\x80\x54\x88\x82\x5e\x86\x92\xd7\x3c\x39\x2b\x94\x76\xef\xfa\x4d\xe4\xaf\x11\xff\xe0\xeb\xae\xa0\x7b\xe1\xfa\xe4\xfc\x6c\xae\xac\x59\x97\x93\x8d\x7c\x8e\xc5\x6f\x00\x00\x00\xff\xff\x41\x95\x11\x83\xdf\x02\x00\x00")

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

var _type_module_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x92\xcb\x6a\xc3\x30\x10\x45\xf7\xfa\x8a\xf1\x6f\x78\xd7\xd7\x42\xd0\x42\xa9\xe3\x76\x51\xb2\x50\xac\x21\x88\x4a\x23\xa3\x47\x20\x94\xfc\x7b\x91\x94\xd8\x72\x70\x4b\xbb\xea\xca\x96\xe6\xc1\xbd\xe7\x2a\x1c\x47\x84\x27\x2b\xa3\x46\x50\x66\xd4\x68\x90\x82\x87\xee\xa8\xb5\xd8\x45\xcf\x03\x1a\xf8\x64\x00\x31\x2a\xd9\x42\xdf\xf3\xfb\x86\x01\x90\x30\xd8\x42\x17\x9c\xa2\x7d\x3a\xef\x04\x11\x3a\x6e\xc4\x1e\xfb\x97\xc7\x4b\x85\x01\x48\xf4\x83\x53\x63\x50\x96\xea\xfe\xe0\x04\x95\x42\x7d\x7b\xb0\x6a\x40\x7b\x40\xb7\xdc\x71\x50\x12\x6d\x0b\xaf\xe9\xc3\x00\xfc\x59\x5a\x0b\xef\xb5\xca\x66\xcb\x00\x06\x9b\x2c\x04\x6c\xe1\xd6\x5a\x8d\x82\xd8\x89\x31\xa4\x68\xca\xf4\x26\x99\x4d\x6e\xde\x78\xb7\xe1\x37\xa9\x98\xfd\xe7\x62\x2e\xa4\x63\x3b\x37\x27\x55\xd1\xe9\x59\xe4\x89\x31\x45\x63\x0c\xa5\x85\xe7\xdf\x5f\xce\x65\x19\x85\xf4\x43\xc1\x5c\x26\xd1\x07\x06\xa0\xd1\x7b\x4b\xf3\xfe\xd2\x38\xe1\x2f\xfb\x17\xd3\xcd\x55\x2a\xd3\xe8\x9d\x43\x11\xf0\xbc\x60\x52\x78\x1d\x59\x10\xfb\x84\x30\xcf\x6e\xff\x9a\x54\x95\x77\x17\x87\x01\xbd\xdf\xd8\x0f\xa4\x3a\xb4\x4b\x96\xdf\xd5\xab\x50\xb3\xc8\x65\xb2\xb3\xfd\x66\x3b\xc5\x54\x1b\x7b\x16\x47\x6d\x85\xcc\xd6\x4c\xbe\xb9\xe0\x99\x41\xf4\xa3\x5c\x01\xb1\x7c\xc9\xd7\x18\x6a\x4c\xeb\x54\x56\xa1\xfc\x1b\x93\xda\xe3\x8f\x4c\xbe\x02\x00\x00\xff\xff\xd6\x8d\x98\x78\xe9\x03\x00\x00")

func type_module_graphql() ([]byte, error) {
	return bindata_read(
		_type_module_graphql,
		"type/module.graphql",
	)
}

var _type_question_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\xdd\x6e\xa3\x40\x0c\x85\xef\xe7\x29\x8c\xf6\x2d\xb8\x4b\x58\x36\x8b\x94\xdd\x26\x0d\x48\x95\x2a\x14\x59\xc1\x41\xa3\xc2\x0c\x65\x06\xb5\x28\xca\xbb\x57\x0c\x61\xc2\x5f\xaa\x26\x77\xd8\xb2\x7d\x8e\x3f\x0f\xba\x2e\x08\xb6\x15\x29\xcd\xa5\x80\x13\x03\xa8\x2a\x9e\xb8\x10\x45\xc1\x6f\x87\x01\x68\xfa\xd4\x2e\xec\x74\xc9\x45\xda\xc4\x25\x8a\x44\xe6\x5c\xd1\x42\xa8\x0f\x2a\x95\x0b\x4b\x29\x33\x42\xc1\x00\xde\x2f\x63\xc2\xba\x20\xd7\x0e\x6d\xa2\xa6\x13\xbb\x86\xd7\xb6\xd5\x89\xd9\x99\xb1\x81\xfe\x06\x53\x02\x9e\x17\x19\xe5\x24\xb4\x82\x0d\xa6\x5c\xa0\xa6\x64\x5b\x51\x59\x1b\x77\x94\xa4\xd4\xcc\xe8\x5a\x62\x06\x50\x60\x4a\x81\x38\x4a\xb7\x69\x30\x5f\x76\x72\x2b\x35\xb3\x17\x57\x9e\x2c\x4b\x3a\xe8\xfe\x02\xfd\x65\x9b\x9a\x1c\x53\x8a\x9e\xd7\x36\x75\x66\x8c\x44\x95\x0f\x56\x33\xb3\x77\xc1\xff\xd5\xda\xdf\x7b\x7f\x9f\x02\xcf\xb7\x65\xad\xb8\x2d\x0a\xfd\x97\x90\x01\x04\xff\x16\x2b\xff\x12\xee\xdb\xe0\xcc\x18\x17\x45\xa5\xc1\x2b\x09\x35\x75\xe3\x03\x93\x3b\xdd\x73\x05\xe7\x8e\x33\xb4\x62\x4b\x54\xfc\xd0\x8e\x31\x7a\x4e\x6c\xce\x8e\x69\x53\x62\x68\xc5\xce\xd8\xe0\xb8\xc7\x78\x9c\x12\x75\xe0\x17\x24\x74\xc4\x2a\xd3\x0a\xb4\x84\x23\x66\x8a\x6e\x60\x0e\xe5\x1b\x89\x5e\x12\x2d\x3b\xb7\xc7\xd1\xb1\x87\x1d\x92\xda\x60\x9d\x49\x4c\x8c\x8f\x6e\xfd\xeb\xea\x57\xfb\x51\x91\xcc\xf2\xbd\xfd\xe6\x1f\x7e\xf2\x03\xd4\xad\xee\x14\xf5\x98\xf4\xd8\xe9\x2c\xe8\xab\xd9\x6f\x38\xcd\xdd\xe3\xe1\x73\x74\xd0\x87\xf8\x7e\x0a\xbd\xcb\xfc\xe1\x99\x9e\xfc\x8b\x53\xfd\x70\x04\xe4\x2b\x00\x00\xff\xff\x60\x64\xc9\x14\xa2\x04\x00\x00")

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
	"type/blog.graphql":       type_blog_graphql,
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
		"blog.graphql":       &_bintree_t{type_blog_graphql, map[string]*_bintree_t{}},
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
