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

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x96\x4d\x6f\xe3\x46\x0c\x86\xef\xfa\x15\x6f\xe0\x8b\x03\x04\xdd\xbb\x6e\x49\x16\x0b\x04\xe8\xb6\xbb\x8d\x73\x2a\x8a\x62\x2c\x51\xd2\x34\xa3\x19\xed\x7c\x34\x56\x8b\xfd\xef\x05\x67\xf4\xe9\xca\xc6\xe6\x26\x8b\x2f\xf9\x90\x34\x49\x68\x87\x6f\x81\xac\x24\x07\xd7\x98\xa0\x4a\x68\xe3\xd1\x9a\x52\x56\x3d\x7c\x43\x28\x8f\x3f\x81\x4e\x54\x04\x4f\x25\xa4\x46\x27\xac\x15\x4a\x91\xca\x7c\xdf\x11\xbe\x88\x9a\x9e\x74\x65\xf0\x6f\x06\x78\xe3\x85\xca\xf1\xa4\xfd\x0d\x76\xf8\x25\xb4\x47\xb2\x30\x15\x3a\x51\x93\x83\xa8\x3c\x59\xf8\x46\x3a\x18\x4d\x19\x60\xaa\xca\x91\x9f\xf4\x87\x86\x86\x57\xec\x13\x75\xec\x88\xca\x9a\x36\xa6\xe2\xbc\xb0\x3e\x03\x94\x6c\xe5\xda\xad\x15\x27\x4e\xff\x4d\xa3\x23\x1b\xbd\x32\xa0\x96\x7f\x93\xde\xc8\x46\x7a\x6a\x5d\xb2\x66\xdf\xb3\x4c\xea\x2e\xf8\x58\x47\xac\x61\x91\xd4\x0a\x35\x4b\x3f\x0b\x2d\x6a\xb2\x9f\xa4\xe2\x7a\xd8\x47\x8b\x96\x72\x3c\x7b\x2b\x75\x9d\x01\xd4\x0a\xa9\x16\xbf\xff\x32\xc7\x83\xf4\x6a\x29\x09\x41\x96\x39\x5e\x5e\x9e\x3e\x72\xdb\x48\x51\xd7\x18\x3d\x0b\x26\xd6\x47\x52\x54\x0b\x4f\x0b\xd8\xda\xd5\x17\x7f\xf2\xcf\x29\xf0\xfb\x53\xb9\x42\x7f\x34\xc1\x3a\x72\x97\x2b\x15\x45\x41\xce\x1d\xfa\x8e\x72\xdc\x4f\xcf\x19\x70\x14\xc5\x6b\x6d\x4d\xd0\xe5\x63\x43\xc5\x6b\x8e\x07\x63\x14\x09\x9d\x01\x9d\x95\x05\xe5\xf8\xa4\x8c\xe0\x0e\x0b\xa5\xcc\x1b\x95\x07\xf3\x10\xfa\x49\x86\x1d\x9e\x1b\xf3\x06\xa3\x55\x8f\x22\x65\x01\xdf\x08\x8f\xde\x04\x08\x4b\x10\xc1\x37\xc6\xca\x7f\xa8\x84\x37\x38\x1a\xf3\xba\x4c\xbb\xed\x84\xee\x2f\xf5\x6c\x55\xc3\xe4\xf4\x33\x39\x67\xf4\x0f\xf9\x00\x5e\xd4\x2e\xc7\xef\x6c\xfc\x63\x0e\xf1\xdc\x2b\x25\x8e\xe1\x4a\xbf\xe8\x54\xa8\x50\x52\x62\x2d\x9b\x32\x18\x3e\x9b\x32\xf0\x9f\xf3\x3f\xc3\x81\x9c\x9f\x5f\x47\xa2\x27\x5b\x89\x22\x6e\xa0\xd4\xc2\x53\xf9\x35\x90\xed\x23\xb5\x1b\x96\x32\x9f\xd6\x73\x4e\xf2\x57\x5b\x92\x7d\x48\x3a\xe1\x0a\xd2\xa5\xd4\xf5\x92\x58\x49\x52\xd3\x44\xdd\xb0\x63\x5c\xf4\x39\xba\x8c\x91\x47\x3b\x20\xca\x56\xea\xfd\xdc\xae\xdb\x1c\xf7\xfc\x6a\x34\xb9\x3d\xe7\x93\x72\x19\x6d\xfc\x7c\x93\x65\x40\x9b\x96\x69\xed\x3e\x6c\xd8\x6c\x5e\x86\xb8\x43\x15\xfb\x9b\xaf\x17\xf1\x0e\x26\x55\x96\x8f\x25\xce\x81\xd8\x8d\x61\xe5\xb0\x4d\x0b\xda\xcd\x6d\x3e\x2d\xd9\x42\xb1\x0d\x5c\x6f\xe3\x26\x71\x94\x8c\xc8\x22\xcd\xe2\x19\x71\x98\xd0\xc9\x2e\x2f\x00\x57\x93\xbc\xc9\x1b\x14\x23\x4e\xc5\xc9\x3a\xa3\xa5\x71\x9b\xac\xdb\xa8\xe5\xfc\x6f\x92\x92\x60\x04\x1d\x95\xa9\xcf\x30\x0f\xca\xd4\x83\x65\x8d\xd8\x08\xc6\xda\xb9\x45\x71\xbf\x2f\x34\x60\x71\x81\x2e\x34\x80\x15\x5f\xd2\xc9\x4f\xa1\xf6\x9c\x15\x9f\xfd\xc9\xcc\xeb\xec\x38\xde\x8b\x23\xcb\xcc\x6f\x81\x9c\x97\x97\x7a\xf1\x75\xb0\x5e\xc1\x8e\x92\x01\x3c\xc6\x3b\x6b\xc9\xa8\x8a\x67\xd6\xf9\x6d\x1a\xef\xf6\x15\x12\x9b\x07\x0a\xc7\x38\x23\xb0\x95\xf7\x24\x1e\x8e\x0b\x6b\x12\x6d\xd7\xb6\x24\x0a\x06\x46\x8a\x74\x46\x49\x0a\x6e\x9c\x23\x61\x8b\x66\xbc\x74\x9b\xbc\xf5\x19\xbc\xcd\xf1\xbc\xf2\xf9\x8d\x5c\x50\x1e\x3b\xdc\xf3\xe5\x77\x43\x44\xa9\xeb\xb1\x88\xbb\x71\x4e\x21\x74\x99\xfa\x96\x01\xbb\x78\x74\xf7\x9e\x4e\x7e\x3c\x3e\xb7\xf8\xf0\x61\x70\xe7\xcf\x0b\xa5\xd2\x77\x03\x45\x25\x4f\x03\x59\x2f\x2b\x59\x08\x1f\x6f\xe0\xde\x9b\x57\xfe\x20\x18\x2e\x17\x0f\xc7\x5a\xc0\xc7\x8c\x9f\x6b\x63\xcf\x17\x72\x85\xcd\xf1\x98\x54\xfd\xd0\xb3\x1d\x1c\x15\x46\x97\xc2\xf6\x8f\xef\xf7\xff\x9e\xfd\x17\x00\x00\xff\xff\x55\x93\x79\xc3\x84\x09\x00\x00")

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

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x54\xc1\x6e\xdb\x30\x0c\xbd\xeb\x2b\x5e\xe0\x4b\x0b\x0c\xfd\x00\xdf\x82\x16\x1b\x02\x6c\xc3\x86\x36\xa7\x61\x07\x26\xa2\x1d\xad\xb2\xe4\x49\x54\x07\x63\xe8\xbf\x0f\x94\x63\xcf\x45\x6f\x24\x2d\x3e\x3e\xbe\x47\x98\x43\x19\x70\xcc\x9c\x9e\xa6\x91\xf1\xd7\x00\x03\x05\xea\x39\x19\xc0\x05\xeb\x5e\x9c\x2d\xe4\x0d\x60\xd9\x73\x4f\xc2\xe6\xd5\x98\xa6\x76\xc0\x65\x10\xe4\xe2\x02\x44\x9b\xe5\x42\x82\x73\xf4\x9e\x84\x33\x5c\xe8\x62\x1a\x48\x5c\x0c\xa0\x53\x2c\x02\xb9\xf0\x02\xfe\xe1\x3f\x5c\x03\x0a\x76\x33\x0a\x45\xa1\x15\x30\xdf\xe1\x31\x0e\x8c\xce\xb1\xb7\x19\x94\x18\x21\x0a\xfa\xe8\x42\x0f\x89\x38\x31\xe8\x85\x9c\xa7\x93\x67\x58\x1e\x39\x58\x17\x7a\xd3\x20\x86\x3a\xab\x92\x8a\xdd\x0c\x38\xc5\x52\x11\x12\xff\x2e\x9c\x45\x21\x2c\x09\xa1\x8b\xe9\xce\xd4\x97\x75\x25\x15\x40\xb3\x76\xd5\x64\x67\x00\x1e\xc8\xf9\x16\x8f\x92\x74\x00\xd0\xb9\x94\xe5\x2b\x0d\xbc\xd4\xf4\x91\xa7\xf7\x35\x61\xcf\xe3\x25\x06\xde\x34\xff\x8a\xa7\x27\x27\x7e\x5b\xd2\xd6\xcf\xb1\x77\x61\xdb\x7b\x8e\xc3\x48\x61\x6a\x71\x3f\x07\x06\x18\x53\xec\x9c\xe7\xc3\x40\x3d\x1f\xd3\x96\x11\x9d\xc5\xbd\x38\x99\x6e\x46\xea\xb9\xc5\x37\xea\xf9\xb6\xc5\xfe\x5a\xd5\x54\x9d\x9d\xee\x63\x49\x99\x73\x8b\x1f\x5f\xae\xf1\xee\x67\xfd\x50\x5f\xf2\x5c\xba\x71\xb6\xc5\x21\xc8\xee\xb6\xc5\xf2\x0c\x0d\x3e\xb1\x60\x98\x70\x9e\xf3\x2e\xc5\xe1\x1a\x1f\x1e\x0c\xd0\x80\xac\x4d\x9c\x73\x8b\xfd\x1c\xe8\xa1\xd4\xeb\x5a\x58\xac\x17\x16\xf8\xcf\x8c\xba\x10\x27\x61\x7b\xdd\xd8\xb3\xb0\x5d\xbf\x76\xe4\xfc\x9a\xbe\x9a\xd9\xa8\xed\x56\x70\xda\x32\x70\x90\xac\x4b\xbb\xa0\x50\xdf\x0b\xa7\xa9\x4e\x62\xdb\xd7\x6d\x97\x16\x5d\x56\x15\x3a\x84\x2e\xce\x2a\x69\xf4\x0e\xb9\xf6\x96\xa2\x3a\x1c\x8f\x87\x87\xdd\x7a\x14\xdb\x55\xaa\x47\x89\x75\xe0\x5e\xde\x1a\xa7\x74\xd5\xb7\x85\x76\x83\x06\xfb\x45\xb9\x73\x49\x89\x83\xf8\x09\x27\xae\x97\x4c\xcf\x1c\x70\x9a\x40\xf5\x52\x67\x26\xab\xee\xca\x24\x0b\x49\xc9\x0b\xe0\x63\xcd\xde\x0f\xd2\xca\xe0\x42\x11\xce\x4f\x89\xce\xcf\x6c\x5b\x7c\xf4\x91\xa4\x5e\x70\x48\xd1\xfb\xb7\x4c\x17\x7f\xb6\xb0\x75\xde\x48\x39\x57\x43\x66\xf9\xeb\x9f\x60\xf1\xc6\xbc\x9a\x7f\x01\x00\x00\xff\xff\xf3\x3c\x6d\x80\x33\x04\x00\x00")

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
