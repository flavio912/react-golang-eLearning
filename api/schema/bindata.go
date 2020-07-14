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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x96\xcd\x6e\xe3\x36\x10\xc7\xef\x7a\x8a\x09\x7c\xc9\x1e\xba\xbd\xfb\xb6\xb1\xb0\xa8\x81\x04\x49\xd7\xd6\x03\x4c\xc4\xb1\x4c\x2c\x45\xaa\x24\x95\xad\x51\xe4\xdd\x0b\x52\x12\xbf\xa4\xe6\x92\xde\xc2\x9f\xe7\x3f\xdf\xa4\xb2\x83\x7e\xb4\x68\xb9\x92\x06\x5a\x94\xd0\x2b\xc6\x2f\x37\xb0\x57\x02\xf6\xfa\x15\xe8\x6f\x6a\x47\x4b\x0c\x0c\xfd\x35\x92\xb4\x1c\x85\xb8\x7d\xad\x2a\x2e\x87\xd1\x42\x33\x08\x85\xec\x3b\x17\xf4\x44\x16\xe1\x9f\x0a\xe0\xc2\x05\x9d\x6f\x03\xed\xe1\x64\x35\x97\xdd\x5d\x05\xd0\x2a\x69\x49\xda\x47\x92\x9d\xbd\xee\xe1\x28\xed\x5d\xf5\x5e\x55\xf6\x36\x50\xe2\xe2\x07\x99\xc1\xbb\x18\xb5\x48\xd5\x66\x6c\x5b\x32\xe6\xac\x7e\x92\x8c\xfc\x7d\x9d\xc3\x69\x32\xf4\x3e\xfe\x53\xe4\x83\x3e\xcd\x25\x7b\x53\x64\x3d\x97\x8f\xaa\xe3\xf2\xde\x7b\xdc\xc3\xb7\x40\xee\xbe\xec\xe1\xdb\x68\xaf\xde\x4f\x05\xd0\xa3\xc4\x8e\x74\x66\xfd\x94\xb0\xd2\x9e\x91\xa0\x0e\x2d\x65\x82\x3a\x85\xb9\xc2\x35\x4b\x13\x5a\x9a\x9d\x2e\x92\x43\x0a\x8f\x8e\x39\xdd\x7c\x76\x2d\x1b\xd8\x5a\xd4\xa4\x70\x43\x34\x68\xe5\xa6\x75\xec\xb1\xa3\xa9\x8b\x3f\xdc\x90\x8d\x8d\x0e\xd2\xf1\x3a\x71\x3e\xad\x32\xf0\x4b\xe2\x70\xed\x63\x1e\x4f\x9e\x83\x6b\xd0\x2a\xf1\x3a\x85\x21\xf1\x07\xa5\x04\xa1\xbc\x8b\x4d\xf2\x73\xca\x5b\xe4\x51\x90\xf8\x53\xc8\x32\x33\x6f\x22\x5a\x99\x4f\x49\x65\xe6\x75\x44\x1f\x24\xb4\x4c\x36\xcf\x69\xa1\x41\x98\x63\xd7\x49\x25\x0d\x85\x3c\x4b\x2f\x4d\x46\x83\x97\x05\xc4\xf0\x07\xd5\x0f\x28\x6f\x79\xf4\x19\xc6\xe0\xd3\x39\x44\x2b\x44\x4d\x0a\x37\x44\xbb\xb9\x3b\x91\x64\xb1\x97\x0d\xaa\x00\xdc\xbd\xf7\x6c\x33\x13\x6f\x30\x5f\xa7\xcd\xfd\xf6\x06\x9a\x5a\x1c\x6c\x7b\xc5\xf4\x41\x48\x5b\x0f\x80\xc3\xa0\xd5\x5b\x28\x63\x1c\x39\xdb\x43\xd3\x1c\xeb\x34\x6f\xf7\x30\x8d\xe2\xc2\xc5\x0b\x49\xc6\x65\xf7\xac\x19\xe9\xfb\x56\x70\x92\xf6\x44\xad\x26\x1b\x02\xe4\xde\x0d\xbe\xd1\xb3\x14\x5c\xd2\x41\x8d\xda\x84\x99\x9c\x0a\x9e\x74\xca\x1d\x61\x07\xe7\xe7\xfa\x19\x50\x08\xf5\x0b\x04\x97\x3f\x67\x5f\x07\x81\xc6\x68\xa5\xfa\xb5\xbb\xe2\xa7\x0d\x8f\x8a\x29\x40\xc6\xc0\x8e\x56\x69\xb0\x0a\x5a\xff\x93\x7f\x61\xdd\x1f\x0f\x28\x25\xe9\x4f\xdd\xe7\x61\xd4\xed\x15\xcd\x5c\x95\x59\x94\x2f\x39\x0e\xb9\x15\x3c\x59\xe5\x65\x4b\xd2\x2a\xeb\x84\x7d\x70\x8b\xce\xd8\xe5\x2b\x7c\xc6\x2e\x98\x9f\xb1\x8b\x3b\x87\x96\x3a\xa5\xcb\x85\x9f\x69\x6c\xe0\x0c\x62\x84\x47\x32\x46\x15\x2f\xc7\xc4\x82\x68\x3a\x86\x5b\x92\x2b\x9a\x84\xad\x15\x53\xe9\xb9\xa2\x4e\xd8\x07\xa5\x1f\x25\xe3\x6f\x9c\x8d\x28\xf2\xe4\x22\x2f\x1e\x91\xf8\x43\xd2\x7b\x33\xbe\xf6\xdc\x9e\x93\xc9\x9f\x02\x09\xfa\x88\x5e\xf0\xe6\xb6\x00\x76\x33\x33\x80\xe1\xc3\xf5\x3b\x0f\x01\x0c\xa0\x34\xbf\x48\x1b\xb7\x79\x08\x96\x8c\x8d\x23\x4b\x62\x1d\x02\x29\x72\x4d\x62\x85\xbe\x9e\xb3\xf5\x64\xa5\x30\xa2\x28\x9c\xda\x9b\x0a\xeb\x40\x36\x5b\x3b\xa5\xfd\xa9\x6b\x31\x55\xf9\xa7\x53\xf1\x72\x71\x16\x5a\x54\xbb\xe0\xb2\xe2\xd2\x49\x93\xd1\xa2\xf2\xb5\x93\xa9\xfa\xd2\x49\x9d\xd1\x0f\x16\xec\x49\xb1\x51\x14\xdf\xa7\x89\x15\xe9\x4f\xb0\x4c\x3e\x97\x37\x09\x2b\x12\x2f\xe5\xf3\x77\x3e\x93\xd7\x09\xdb\x4c\x79\xb7\x6c\x97\x7b\xef\xfc\x79\x5e\x1a\x7f\xf6\xc0\x90\x5d\xbe\x83\xc9\xfb\xd9\x0f\x82\xfc\xff\x77\xbf\x39\x03\x03\xfc\x92\xac\x34\x0c\x68\x0c\x31\x50\x1a\x2e\xc8\x05\x31\x40\x68\x17\xed\xfc\x92\xc2\xfd\x77\xa5\xc1\xc7\x31\x5f\xfe\xbf\x67\x31\x8c\xe1\x41\xa8\xe2\x8d\x73\x24\x36\x41\xa8\x2e\xf4\x3c\xb5\x6d\x02\x29\x6d\x5f\x85\xea\xfe\x20\x64\x9f\x5c\xf4\x18\x32\x71\xb6\x8e\x9e\xfc\xb8\x95\xc8\x83\x62\xb7\x4f\xa4\xf1\x5e\xfd\x1b\x00\x00\xff\xff\x05\x99\x08\x17\x96\x0c\x00\x00")

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

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4f\x6f\xe3\xb6\x13\xbd\xeb\x53\x4c\xe2\x43\x7e\x3f\x60\x73\xe8\xe6\x0f\x5a\xdf\xb2\x8a\x8b\x1a\x48\xb1\xa9\x2d\xa3\x2d\x82\x1c\x26\xd4\x58\x26\x96\x22\x55\x72\x98\x44\x58\xec\x77\x2f\x48\x2a\x16\xed\x8d\x8b\x4d\xf7\x52\x18\xb0\x2d\x71\xde\xcc\x70\xde\x9b\x21\x49\xfb\x16\xae\x84\x20\xe7\xaa\xbe\x23\xf8\x5c\x00\x58\x72\x6c\xa5\x60\xaa\x0b\x00\xd3\x91\x2e\xbe\x14\x45\x34\x5c\xb2\xf5\x82\xbd\xa5\x99\xa2\x96\x34\x47\xf3\xd6\xd4\x5e\x51\x01\xc0\xe4\xb8\x00\x50\xe4\x9c\x19\x31\xa5\xf1\xd6\xd1\xd6\xb9\xd1\x4a\xea\x60\x2d\x14\x3a\x67\x8d\x69\x83\xa5\xd4\x4c\x76\x8d\x82\x60\xd9\x2b\x85\x0f\xde\xcd\x99\xda\x08\xd0\xd8\xd2\x34\x44\x96\xba\x39\x2a\x00\xbc\x97\xf5\x14\x56\xab\xf9\x75\x78\x12\xa6\xed\x14\x31\x4d\xe1\x83\x31\x8a\x30\xc6\xe5\x10\xec\x26\xa6\x01\x32\xac\x87\x5c\xdd\xf7\xbb\x06\x68\xbb\xb3\xd5\xe2\xe6\x05\x13\xf7\xfc\xcc\xb9\x0b\xc6\xc6\x4d\xe1\xae\xc2\xe6\xe8\x7e\x9b\x4a\x2a\x41\x0c\x19\x02\xcc\x35\x47\xd3\xbe\xa3\x69\x56\x9e\xa3\x57\x32\x7a\x40\xf1\xa9\xb1\xc6\xeb\xba\xdc\x90\xf8\x94\xa7\x82\x5b\xd2\xa6\x19\x81\x01\xd4\x59\x29\x68\x0a\x3f\x2b\x83\x9c\x36\xa2\x8c\xcd\x72\xa6\x67\x41\xb6\xe3\xec\x8d\xd4\x6c\x4d\xed\x05\x4b\xa3\xb3\xd7\x1b\xf3\x54\x99\x72\x5b\x86\xec\xbd\xb7\x2e\x5f\x89\xa1\x0a\x80\xa7\x0d\xf2\x9f\xc6\xdf\x10\x5a\x3d\x85\xbb\x61\x1f\xf7\x51\x53\x7f\x79\x69\x13\x11\xbb\x2b\xae\x23\x21\xd7\x52\x54\x64\x5b\x97\x05\x11\xc8\xd4\x18\xdb\x4f\xa1\x1c\xfe\x85\x4d\x2b\x65\x9e\xa8\xae\xcc\x07\xdf\xe7\xc5\x70\x03\xb5\xc1\x75\xc6\x72\x0c\xf0\x80\x5a\x93\x9d\xb7\xd8\x50\xce\xdd\x2e\x39\xb7\xd8\x50\xae\x95\x5b\x6c\xa4\x46\xa6\xfa\x37\x4f\xb6\x8f\xd4\x51\xdd\x50\x08\x90\x00\xc1\x73\x87\x0d\xcd\xf5\xda\x4c\x83\x79\xfc\x97\xb4\xdc\x79\x86\x25\x3e\x52\xf9\x22\xf1\x04\x99\xc7\x85\x4c\x05\x7b\x84\x67\x9b\x0e\x1a\x4c\x4a\x7c\x0b\x5f\x6f\x56\xcb\x9e\x58\x5e\xd1\xca\xa0\xe7\xd8\x13\xff\xc4\x56\x56\xe4\xa5\x8f\xfe\x73\x27\x9e\x8d\x1d\xb7\x04\x93\xf4\x63\xd6\xc0\x1b\x82\x2a\xac\xc2\xff\xb6\xf3\x00\x44\x6a\x17\xa3\x55\xff\xff\xd0\x73\xf8\x7c\x8b\x96\xa5\x90\x1d\x46\xf5\xa4\xca\x39\x46\xcb\xd7\x18\xe4\x57\xc9\x36\x6c\x85\x74\xbd\xf3\xac\x8c\xc0\x9d\x0a\xed\x90\xf3\x31\x0e\xa2\xef\x63\x66\x02\x8e\x84\xd1\x35\xda\xbe\xfc\x06\xe6\x60\x02\xcb\x8d\xb1\x0c\x35\x39\x61\x65\x17\x92\x0b\x55\x48\x1b\x3e\xc0\x2b\x4c\xe0\xe4\xea\xc1\x78\x06\xde\x48\x37\xd8\x9e\x84\xc0\xc1\xea\x50\x97\x06\xd4\x2f\xe6\x09\xd8\x6c\xe7\xd8\x61\xfc\x6b\xdd\x0c\x13\x98\x39\x96\x6d\xe8\x81\x64\x72\xc8\xd7\xe1\xbe\x0f\x49\xfc\xbe\x41\x86\xde\xf8\x13\xa5\x40\x85\xf5\x13\x78\xf0\x4a\x11\x43\x67\xa4\x66\x77\x78\x36\x04\xf4\x22\x5b\xfa\x1a\xf8\x5f\xd5\xbb\x7b\x39\x26\xb7\xd3\x62\x18\x46\x5b\x01\x8e\x6f\xa3\xea\xd2\x51\xb0\x7f\xba\xee\x1f\x49\x23\xda\x12\x32\x55\xd8\x8c\xba\xdd\x3f\x38\x76\xf6\x75\xb4\x1d\x76\x15\x36\xd1\x7c\x74\xfb\xed\xd8\x17\x85\xff\x1b\x07\x79\xde\x2f\x7e\xde\x92\x7c\xc2\xdf\x7a\x2b\x36\xe8\x86\xa6\x75\xa3\x83\x24\xc4\x40\x5c\x38\x5b\xef\x63\xe1\x1c\xd9\x91\xc9\x23\x98\xc0\x8d\x74\x1c\xba\xad\x26\x45\x0d\x32\xb9\x98\xbe\x8b\x6d\xca\x16\xe7\xfa\xd1\x48\x41\xb3\x16\xa5\xca\xb8\x0c\x42\xea\x98\xea\x41\x00\x83\xc2\x5e\x39\x9b\x4b\xa3\xd7\xd2\xb6\x5f\x5f\x44\xf6\xb2\x5e\x90\xeb\x8c\x1e\xae\x03\x6c\x51\x3b\x8c\x9d\x58\xee\x5f\x37\x42\xce\x6c\x3d\x81\x4c\x73\x32\xf4\x1e\xea\x1e\xa4\x03\xa3\x01\x41\x18\xcd\x16\x05\x03\xea\x1a\x78\xe3\x1d\xd4\x86\x9c\x3e\x61\xd0\x44\x75\xe8\xd6\x0e\x7b\xf0\xdd\xda\x9a\x61\x60\x5a\xd9\x51\xa9\x24\x69\x5e\x92\xb0\x94\xcf\xa5\xf9\xfa\xb5\x5c\x42\xac\x35\x2a\x47\xef\x52\xc3\x4b\x07\x8d\x7c\x24\x0d\xce\xc4\x02\x83\x40\x1d\xa2\x8c\x14\x5d\x53\xc0\x1d\x98\xaa\x91\xca\x09\xdc\xc1\xec\x8f\xab\x5f\x6f\x6f\x66\x50\x7e\x5c\x2d\x96\x33\x58\x56\x8b\x55\x59\xad\x16\xb3\x62\x02\x00\x9f\x53\x37\x1c\xa7\x7b\xe5\xf1\xbb\x41\x6b\xc7\x67\xe7\x67\xe7\x67\xa7\xf9\xf7\x79\xfc\x39\xfe\xf2\x6e\x07\x97\xee\x9f\x23\xee\x87\xf7\xf1\x73\x9a\x7e\xde\x9f\x66\x4f\xfb\xd0\x70\x83\x1d\x81\x3f\xfe\x94\x3e\xa7\x17\x97\xf1\x73\x7a\x71\xb9\xfd\x73\x71\x79\x91\xc0\xf7\xc5\xdf\x01\x00\x00\xff\xff\x24\x05\x0b\x11\x43\x0b\x00\x00")

func type_course_graphql() ([]byte, error) {
	return bindata_read(
		_type_course_graphql,
		"type/course.graphql",
	)
}

var _type_delegate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x93\x51\x4f\xc3\x20\x14\x85\xdf\xfb\x2b\xee\xde\xf4\x2f\xf0\x36\xbb\x97\x26\xc6\x4c\x6d\x9f\xcc\x62\xae\xe3\x0e\x51\x0a\x04\xa8\x4b\xb3\xec\xbf\x1b\xda\x75\xed\x3a\x8c\xfa\x06\x27\xe7\x52\xbe\xc3\x69\x68\x2d\xc1\x8a\x14\x09\x0c\x04\x87\x0c\xa0\x69\x24\x67\x50\x55\xc5\x6a\x91\x01\x6c\x1d\x61\x20\xbe\x0c\x0c\x9e\x83\x93\x5a\x64\x00\x65\x99\xbf\x16\xab\x41\x88\x2e\xaa\x51\xaa\x89\x63\x27\x9d\x0f\x0f\x58\xd3\xd4\xa4\xf0\x5a\x0b\xa4\xc8\xbe\x1b\x4d\x93\xe1\xe8\xbb\x37\x42\xea\xa9\xf1\xc3\xbc\x95\x32\xa8\x8b\xe1\xad\xa9\x2d\xea\x96\x41\xde\x2f\xa2\x66\x9d\xd9\x49\x45\x45\x8d\x82\x2a\x37\xbd\x13\x6e\x83\xfc\x92\xa1\xbd\xb1\x28\x88\xc1\x1a\x05\xdd\x32\x58\x9e\xd4\xb8\x1d\x4c\x94\x9b\xc6\x79\xf2\x0c\x5e\x96\x93\xfd\x62\x93\x1d\xb3\x4c\x6a\xdb\x04\xc8\xbb\x58\x86\xdc\x8a\x4e\x3b\x8c\x57\x8a\xe9\xf5\x19\xfe\x27\x8c\x79\x8a\x29\xe6\x54\x60\x82\x34\x39\x0c\xb4\x46\xef\xf7\xc6\x71\x06\x77\xc6\x28\x42\x3d\xcf\xc3\x2a\x83\xbc\x34\x9f\x74\x8e\x76\x24\xaa\x2c\x4f\x12\xcd\xea\xf0\x17\xbe\x04\x5e\x82\x26\x01\x9c\x82\xfb\x0d\x00\x40\xd3\x7e\x24\x1f\xb1\xba\x66\x5f\xbe\xd3\x13\x79\x6b\xb4\xef\x7b\xce\x4f\x22\x3b\xd7\x7f\x31\xc9\x92\xff\x78\xe4\xe0\x8e\x8d\x01\x59\x5b\x45\x35\xe9\xe0\x63\xa1\xa4\x8e\x93\x8f\x0d\xb9\xb6\xfb\x04\x71\xd1\xb5\x68\x18\xd9\x44\x1e\x14\x54\xe8\x9d\xe9\x1b\x18\x57\xe3\x1b\x0c\xbe\xae\xfd\xdd\x09\xd7\xbf\x9a\x9d\xdd\x2b\x6a\xda\x44\xb2\xf1\xd9\x8f\xd9\x77\x00\x00\x00\xff\xff\x0e\xc3\xfb\xcb\xd7\x03\x00\x00")

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
