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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x94\xcf\x8e\xdb\x2c\x14\xc5\xf7\x7e\x0a\xa2\x6f\xf3\x75\x33\x0f\x90\xdd\x24\x51\xd5\x48\x13\x75\x34\x89\x1f\xe0\x8e\xb9\x43\x50\x31\x50\xfe\x44\xb5\xaa\xbc\x7b\x05\x76\x30\x60\xcf\xaa\x5d\xf2\xf3\x3d\x87\xcb\x3d\xe0\xff\x48\xef\x1d\x38\xae\xa4\x25\x1d\x48\xd2\x2b\xca\x3f\x06\xe2\xae\x48\xe8\xfb\x13\xc1\x5f\xd8\x79\x87\x94\x58\xfc\xe9\x51\x3a\x0e\x42\x0c\x4f\x4d\xc3\xa5\xf6\x8e\xb4\x5a\x28\xa0\x5f\xb9\xc0\x13\x3a\x20\xbf\x1b\x42\x3e\xb8\xc0\xcb\xa0\x71\x4b\xce\xce\x70\xc9\x36\x0d\x21\x9d\x92\x0e\xa5\x7b\x41\xc9\xdc\x75\x4b\x8e\xd2\x6d\x9a\x7b\xd3\xb8\x41\x63\x66\xf1\x86\x56\x47\x0b\x6f\x44\xae\xb6\xbe\xeb\xd0\xda\x8b\xfa\x81\x72\xe6\xf7\x65\x0f\xe7\xb1\x30\x7a\x7c\x2a\x8a\x9b\x9e\xa6\x23\xc7\x52\xa0\x3d\x97\x2f\x8a\x71\xf9\x7f\x74\xdc\x92\xe7\x44\x36\x5f\xb6\xe4\xd9\xbb\x6b\xf4\x69\x08\xe9\x41\x02\x43\x53\x54\x9f\x32\x56\xd7\x53\x14\xc8\xc0\x61\x21\x38\xe4\xb0\x56\x74\x06\xc1\xe1\xe4\xf9\x50\xec\x73\x78\x0c\x2c\xc8\xa6\x75\x98\x98\xa6\x4b\x51\x9b\xc3\x15\x91\x36\x2a\x84\x75\xec\x81\xe1\x38\xc4\xb7\x90\xb1\x75\xb3\x41\x9e\x6e\x10\x97\x61\xd5\x1b\xbf\x66\x86\x4b\x8f\x29\x9d\xb2\x87\x30\x9f\x45\xe3\x87\x1c\xa6\xc6\x77\x4a\x09\x04\xb9\x49\x33\x8a\x29\x95\x13\x8a\x28\x29\xe2\x2a\x35\x59\x94\xb7\x33\x5a\x94\x8f\x3d\x15\xe5\x87\x19\x7d\xde\xcf\x23\xd6\xb2\xa5\x07\x4d\xba\x12\x87\x39\x2a\x69\x31\xb9\xec\x55\xaf\x41\x0e\xa5\xc9\x04\x67\x8f\x71\x9d\xce\x56\x89\xda\x1c\xae\x88\x8a\x9d\x1e\xa1\x37\x84\x84\x97\x1a\xd9\xea\xbe\xb1\x60\x7a\x00\xab\x57\x32\x16\x18\xec\x40\xbb\xee\x0a\xf9\x13\x2e\xc7\x05\x5a\x1b\x75\x4b\x4d\x7b\xcf\xe9\x96\xb4\xed\xf1\x50\x76\x69\xe1\x86\xdf\xa5\xe0\x12\xf7\xca\x1b\x9b\xe6\x7a\xae\x78\x76\xc0\xb0\x9c\x94\x7b\x01\xd6\x1a\xa5\xfa\xa5\xb8\xfa\xb4\xd4\x8f\xf3\xb9\x00\x2b\x53\xb8\x00\x4b\xb5\x17\x60\xf3\x20\xc1\x21\x53\xa6\xce\x6c\xa2\xb3\xfd\x04\x92\xee\x05\xad\x55\xd5\x0d\x1e\x59\xd2\x8c\xcb\x94\x73\xa9\x68\x33\xb6\x54\x8c\xd7\xb8\x54\x1c\x32\xb6\x76\x91\xb5\x37\xdd\x15\xec\x34\x58\xfb\x90\xbd\x96\x38\x29\x2b\xbe\xb8\xcb\x47\x49\xf9\x8d\x53\x0f\xa2\x3c\xe4\xcc\xab\x57\x31\x7f\x58\x78\xed\x84\xaa\xe2\x08\x64\x3e\x84\x50\x21\x90\x77\xa1\xd8\x37\x04\x8a\xe6\x1f\xfc\xd6\x76\xa5\x59\x39\xf6\xea\xe3\x5a\x23\x3b\x45\x87\xbf\x68\xe3\xde\xfc\x09\x00\x00\xff\xff\x36\xed\xd0\x32\x9c\x07\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xcd\x6a\xdc\x4c\x10\xbc\xeb\x29\xca\xf8\x62\x83\xf9\x1e\x40\x37\xaf\x8d\xc1\xf0\x25\x71\xf0\xfa\x14\x42\xe8\x95\x5a\xd2\xc4\xa3\x69\x65\x7e\xb2\x56\x82\xdf\x3d\x8c\xfe\xb5\x68\x03\xb9\x49\xea\xaa\xae\xea\xde\xea\xbd\xc4\x8f\xc0\x56\xb1\x83\xab\x24\xe8\x1c\x46\x3c\x6a\xc9\x55\xd1\xc2\x57\x8c\xfc\xf0\x1f\xf8\x8d\xb3\xe0\x39\x87\x32\x68\xc8\x5a\xd2\x9a\x75\xe2\xdb\x86\xf1\x44\x25\x3f\x9a\x42\xf0\x3b\x01\xbc\x78\xd2\x29\x1e\x8d\xbf\xc0\x25\x3e\x86\xfa\xc0\x16\x52\xa0\xa1\x92\x1d\xa8\xf0\x6c\xe1\x2b\xe5\x20\x86\x13\x40\x8a\xc2\xb1\x9f\xf0\xfb\x8a\x87\x4f\x91\xd3\xe1\x22\x11\x85\x95\xba\xb3\xe2\x3c\x59\x9f\x00\x5a\xd5\x6a\x4d\xab\xe9\x2d\xda\x3f\x1a\x34\x6c\x3b\x56\x02\x94\xea\x27\x9b\x0d\x37\xca\x73\xed\xfa\x6a\xf2\x9e\x24\xca\x34\xc1\x77\x73\x74\x33\x2c\x4c\xad\xa4\x66\xe8\x07\x32\x54\xb2\x7d\x50\x3a\xce\x13\x39\x86\x6a\x4e\xf1\xec\xad\x32\x65\x02\x70\x4d\x4a\x2f\xde\xbf\xcb\x61\xaf\xbc\x5e\x42\x42\x50\x79\x8a\x97\x97\xc7\xfb\xb8\x36\xd6\xdc\x54\x62\x66\xc0\xa4\x75\xcf\x9a\x4b\xf2\xbc\x10\x5b\x53\x7d\xf6\x2d\xbe\x4e\x8d\xff\xdd\xca\x5f\xd4\xef\x24\x58\xc7\xee\xfc\xa4\x94\x65\xec\xdc\xbe\x6d\x38\xc5\xed\xf4\x9c\x00\x07\xca\x5e\x4b\x2b\xc1\xe4\x77\x15\x67\xaf\x29\x76\x22\x9a\xc9\x24\x40\x63\x55\xc6\x29\x1e\xb4\x50\xdc\x30\x69\x2d\x47\xce\xf7\xb2\x0b\xed\x04\xc3\x25\x9e\x2b\x39\x42\x8c\x6e\x91\xf5\x2e\xe0\x2b\xf2\x68\x25\x80\x2c\x83\x82\xaf\xc4\xaa\x5f\x9c\xc3\x0b\x0e\x22\xaf\x4b\xdb\x75\x43\xa6\x3d\xb7\xb3\xd5\x0c\x13\xe9\x7f\x76\x4e\xcc\xd9\x3d\x9f\x2e\x8d\x4a\x97\xe2\x4b\xac\x7e\xed\x7b\x78\xb6\x05\x65\xdd\x41\x28\x43\x9e\xf3\xcf\x81\x6d\xdb\x75\x6a\x86\x1b\x49\xa7\x6b\x99\x65\x3f\xd9\x9c\xed\xae\xc7\x91\xcb\xd8\xe4\xca\x94\xcb\x75\x15\x8a\xf5\xf4\x03\x5f\x44\x62\x77\x77\x73\x77\xd5\x75\x1e\xeb\x00\xe5\xb5\x32\x57\xb3\xfb\x8b\xeb\x14\xb7\xf1\xdb\x58\x73\x57\xd1\x50\x6f\x66\xac\xc5\xe7\x48\xae\xfb\x6c\x2f\xe8\xd7\xe9\x18\xf8\xb9\xbc\xec\x70\x83\xa2\xdb\x59\xba\xbe\x8b\x1b\x48\x3f\x59\x3a\x8e\x38\x37\x7a\xea\xef\x33\x1f\xb2\x7d\xe2\x75\x8c\xfc\x02\xb1\xad\xb7\xbe\x8d\x4d\xc1\x11\x32\x28\x66\x7d\x30\x4e\x04\x87\xb8\x4c\x75\x75\x46\x6f\x15\xab\x4d\xb9\x01\x31\xa8\xe9\x2e\x51\x27\x62\x7d\xcc\xa6\xea\xb6\xd2\x32\x8b\x9b\x42\x3d\x60\xd0\x39\x68\x29\x4f\x54\x76\x5a\xca\xa1\xb2\x56\xd8\xe8\x15\xb1\xd3\x7e\xba\x4b\x3b\x33\xfd\xe2\xbf\xe0\xcc\xf4\x11\x31\xb4\x0a\x2e\x92\x5e\x1c\x5b\x24\xef\xc9\x9f\x00\x00\x00\xff\xff\xba\xdc\x95\x6b\x65\x06\x00\x00")

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

var _type_blog_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x92\xc1\x4e\x03\x21\x10\x86\xef\xfb\x14\xd3\xd7\xe0\x66\xdb\x18\x37\x31\xa6\x5a\xf7\x64\x7a\xc0\x32\xa5\x28\x0b\x04\x86\xc3\xc6\xf4\xdd\x0d\xb3\x6d\xa1\xa9\x07\x6f\xcc\xc0\xff\xcf\xb7\xff\x2c\x4d\x01\x61\x69\xbd\x7e\xc8\x74\xf4\x11\x7e\x3a\x80\x83\x89\x89\x5e\xe4\x88\x02\xb6\x14\x8d\xd3\x8b\x0e\xc0\xca\xfb\x5e\x88\xfe\x60\x2c\x6e\xcc\x7e\x78\x7b\xbe\x5c\x74\xa7\xae\xbb\xba\xb2\x5f\xce\x46\x09\x18\x86\x7e\x5d\x44\x64\xc8\xde\xb8\x7c\x7a\x35\xb5\xf5\x5e\x12\x6a\x1f\x27\x01\xab\xf3\xa9\x74\x8f\x28\x15\xc6\x7e\x94\x1a\x9b\x61\xe5\x46\x32\xb9\x68\xbe\x82\x5d\xad\xd7\x4b\xaf\x26\x56\x24\x01\x1f\xe5\x9a\x8b\xc5\xee\x06\x91\x7b\xcc\xf9\x95\xbc\xeb\xd7\xad\x75\x8e\xb6\x96\xad\x68\x53\x34\x66\x0c\x16\x47\x74\x94\x60\x23\xb5\x71\x92\x50\xbd\x66\x8c\x13\xbb\xa1\xaa\x73\x77\x25\x2c\xa9\xb1\x77\x07\x2f\xca\x63\x3e\x15\x47\xe3\x42\xa6\xca\xd1\x73\xf9\x37\x0c\xf9\x6f\x74\x37\x38\xb3\x78\x15\x51\x12\xb2\xc5\x55\xfd\xef\x90\xcb\x56\xea\x6e\x9a\x90\xdf\xdb\x69\xd7\x94\xeb\xf3\xb3\xe7\x7d\xbc\x0c\x31\x67\x3c\xf3\x0d\x41\x9d\xf9\x9e\xaa\x7d\x45\x2d\x7f\xd0\x36\xef\x31\x25\x01\x43\xb0\x5e\xaa\xc7\xb9\x53\x5a\x97\x4d\xb6\x98\xa7\xdf\x00\x00\x00\xff\xff\x57\x95\x72\x58\xb4\x02\x00\x00")

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

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x56\x51\x6f\xdb\x36\x10\x7e\xd7\xaf\xb8\xd8\x0f\xd9\x80\xfa\xa1\x71\x1c\x6c\x7e\x4b\x5d\x0f\x33\x90\xa2\x59\xec\x60\x1b\x82\x3c\x5c\xa8\xb3\x4c\x84\x22\x35\xf2\x98\x54\x28\xfa\xdf\x87\xa3\x64\x4b\x76\xe3\xa2\xed\xf6\xb0\x41\x86\x49\x91\x77\x1f\x3f\xde\x7d\x47\x8a\x6c\x2c\xe1\x52\x29\x0a\x61\x55\x57\x04\x1f\x33\x00\x4f\x81\xbd\x56\x4c\x79\x06\xe0\x2a\xb2\xd9\xa7\x2c\x4b\x86\x4b\xf6\x51\x71\xf4\x34\x37\x54\x92\xe5\x64\x5e\xba\x3c\x1a\xca\x00\x98\x02\x67\x00\x86\x42\x70\x9d\xcf\xbb\x34\xdd\x77\x78\xd9\x6e\xe6\xa2\x0f\xb4\x23\xe1\xac\xd1\x56\x50\x95\xc1\x10\xbc\x73\xa5\x58\xb2\x4c\x37\x96\xc9\x4a\xe7\x53\x58\x58\x3e\x11\xd4\xba\xa2\x69\x0f\x45\xc6\x2c\x96\x34\x15\xd6\xda\x16\xf2\xfe\x80\xea\xb1\xf0\x2e\xda\x7c\xb6\x21\xf5\x38\x85\x37\xce\x19\x42\x9b\x01\x54\x5e\x2b\x9a\xc2\x2f\xc6\x61\x82\x53\xce\x38\xbf\xf5\xcd\x00\xe8\x83\x22\x5f\x71\x6f\x44\x5b\xf6\x2e\x8f\x8a\xb5\xb3\xbd\xe1\x8d\x7b\x5e\xb9\x99\x2b\x2b\x43\x4c\x7b\xe3\xd1\x87\xfe\x4c\x5a\x2a\x03\x78\xde\x20\xff\xe9\xe2\x15\xa1\xb7\x53\xb8\x6b\xd9\xde\xa7\x44\xfc\x15\xb5\x4f\x81\x0b\xfb\x33\xa1\x22\xa5\xd7\x5a\xad\xc8\x97\xa1\xb7\x88\x42\xa6\xc2\xf9\x7a\x0a\xb3\xb6\x97\x01\xa0\x31\xee\x99\xf2\x95\x7b\x13\xeb\x6e\xcb\xfb\xd1\xbc\xc6\x82\x40\x0b\xb5\xb4\x1a\x5c\x63\xa1\x2d\x32\xe5\xbf\x45\xf2\x75\x8a\x35\xe5\x05\x09\x8d\xc6\x41\x58\x54\x58\xd0\xc2\xae\xdd\x54\xcc\x53\x4f\x50\xb5\xad\x22\xc3\x12\x9f\x68\xb6\x4d\x5d\xe3\xb2\x48\x13\xbd\xb4\x1d\x64\xa8\xc7\xff\xf6\x76\xf1\x76\x0a\xf2\xff\x2d\xa1\xff\x42\x7a\x71\x27\xf1\x69\x4f\xee\x07\x79\x7f\x21\xed\x8c\x85\xec\x59\x98\x7c\x31\xf0\x0f\x68\x2d\xf9\x45\x89\x05\x2d\x63\xc2\xef\x83\x44\x76\xbe\xdb\x12\x0c\x9b\xc6\xad\x81\x37\x04\x2b\x99\x85\x1f\x76\x3a\x07\xd5\xe8\xdb\x59\x53\xff\x28\xf5\x85\x1f\xae\xd1\xb3\x56\xba\xc2\x24\x84\x26\x72\x81\xd1\xf3\x5b\x14\x25\xad\x74\x29\x5b\x21\x9b\xef\xbd\x1b\xa7\x70\x2f\x42\x7b\xc9\x79\x9f\x0a\xec\x9f\x65\x66\x08\x81\x94\xb3\x39\xfa\x7a\xf6\x15\x99\x83\x21\x2c\x37\xce\x33\xe4\x14\x94\xd7\x95\x90\x93\x28\x34\x1b\x3e\x92\x57\x18\xc2\xe9\xe5\x83\x8b\x0c\xbc\xd1\xa1\xb5\x3d\x95\x85\xc5\xea\x58\xc1\x89\xd7\xaf\xee\x19\xd8\x81\x6a\xa7\x8e\xfb\xbf\x54\x98\x30\x84\x79\x60\x5d\x4a\x0d\x34\x26\xc7\xb0\x8e\x97\xb0\x90\xf8\x7d\x83\x0c\xb5\x8b\xa7\xc6\x80\x91\xf9\x53\x78\x88\xc6\x10\x43\xe5\xb4\xe5\x70\xbc\xcc\xc5\xfb\xa6\x37\xf5\xb9\xe3\x7f\x55\xef\x61\x7b\x4d\xec\x4e\x8b\x05\x53\x79\x72\x0f\x9d\x02\xbb\xe1\xe6\x4a\x48\x44\x0f\xaf\x17\x39\x87\x63\x14\x49\x26\x3e\x22\x10\x26\xa1\x71\xd7\x5c\x2a\x0d\xea\x49\x87\xda\x0d\xf7\x50\xf7\x2e\xa0\x43\xc8\x8e\x90\x27\x64\x5a\x61\xd1\xd5\xc2\xe1\xed\xb1\x17\xab\x93\xdd\x01\xba\xc2\x22\x99\x77\xb0\x5f\xef\xbb\xad\x9a\xef\x01\xe8\xf3\xde\xe2\x7c\x0b\xf9\xc6\xff\x3a\x7a\xb5\xc1\xd0\x1e\x04\xa1\x03\x68\xc4\x2d\xc1\x96\x0b\xf6\x3e\x05\x2e\x90\xef\xd4\x71\x02\x43\xb8\xd2\x81\xa5\x82\x73\x32\x54\x20\x53\x48\xf4\x43\x2a\x7d\xf6\xb8\xb0\x4f\x4e\x2b\x9a\x97\xa8\x4d\x4f\x1f\x22\xce\x8a\x29\x6f\x45\xd5\xaa\xf6\x85\x0b\x7a\xe6\xec\x5a\xfb\xf2\xf3\x4b\xeb\x80\xf5\x0d\x85\xca\xd9\xf6\x9b\x80\x3d\xda\x80\xa9\xba\xbb\x9a\xde\xae\x01\x43\x60\x1f\x09\x74\x73\xf6\x4a\x3d\xa3\xad\x41\x07\x70\x16\x10\x94\xb3\xec\x51\x31\xa0\xcd\x81\x37\x31\x40\xee\x28\xd8\x53\x06\x4b\x94\xcb\x09\x50\x61\x0d\xb1\x5a\x7b\xd7\x1e\xc2\x5e\x57\x34\x33\x9a\x2c\x2f\x49\x79\xea\x9f\x75\x8b\xf5\x4b\x5c\x64\xad\x35\x9a\x40\xaf\x9a\x43\x44\x07\x28\xf4\x13\x59\x08\x2e\x05\x18\x14\x5a\x59\x45\xf6\x3a\x84\x3b\x98\xff\x71\xf9\xee\xfa\x6a\x0e\xb3\xf7\xb7\x37\xcb\x39\x2c\x57\x37\xb7\xb3\xd5\xed\xcd\x3c\x1b\x02\xc0\xc7\x46\xe2\x83\xe6\x1b\x6c\xf0\xaa\xd5\xd0\x60\x7c\x3e\x3e\x1f\x8f\xfa\xff\xe7\xa9\x19\x7c\x7a\xd5\xf8\xc1\x81\x63\x1a\xdd\x4a\x70\xf0\xfa\x6c\x2c\xbf\x51\xdb\x9e\xb5\x9d\xd4\x8c\x77\xc6\xdb\x62\x6c\x5f\x3b\x36\xcd\x97\x5d\xc7\xe6\xf5\x59\x7a\x46\x4d\x23\x68\xbb\xb7\x2d\xa1\xbe\xbb\x7c\x1f\x76\xce\x3f\xfd\xdc\x3c\xa3\xc9\x45\x7a\x46\x93\x8b\x5d\x67\x72\x31\xe9\x00\xee\x53\xfb\xaf\x6c\x10\x86\x22\x10\x1d\x80\xa9\xac\x0c\x32\x41\x89\x8f\x94\xf4\x51\xd5\xff\x87\xdd\x1f\xa1\x32\x4e\x22\x38\x6f\x24\x21\x9d\xc9\xe4\x3c\x3d\xa3\xf4\x3f\x38\x70\xff\x0e\x2a\xf7\xd9\xdf\x01\x00\x00\xff\xff\x9e\x21\x2d\xef\x52\x0c\x00\x00")

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

var _type_users_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x53\x4d\x6b\xdc\x3c\x10\xbe\xeb\x57\x3c\xc1\x97\xf7\x85\x92\x1f\xe0\xdb\x92\x52\x58\x28\xa5\x25\xd9\x53\xc9\x61\xd6\x1a\x7b\xa7\x91\x25\x57\x1a\xa5\x98\x92\xff\x5e\x24\xef\x3a\x4e\xb3\xf4\x66\xcd\xcc\xf3\x31\x1f\x66\x9f\x47\x1c\x12\xc7\x87\x79\x62\xfc\x36\xc0\x48\x9e\x06\x8e\x06\x10\x6f\xe5\x59\x6c\x26\x67\x00\xcb\x8e\x07\x52\x36\x2f\xc6\x34\x15\x01\x49\x20\xe8\x49\x3c\xb4\x80\xf5\x44\x8a\x2e\x38\x47\xca\x09\xe2\xfb\x10\x47\x52\x09\x1e\x74\x0c\x59\xa1\x27\xbe\x90\x7f\x78\xa5\x6b\x40\xde\x6e\xa4\x90\x0b\x75\x21\x4c\xb7\xb8\x0f\x23\xa3\x17\x76\x36\x81\x22\xc3\x07\xc5\x10\xc4\x0f\xd0\x80\x23\x83\x9e\x49\x1c\x1d\x1d\xc3\xf2\xc4\xde\x8a\x1f\x4c\x83\xe0\xab\x56\x35\x15\xfa\x85\x70\x0e\xb9\x32\x44\xfe\x99\x39\x69\xa1\xb0\xa4\x84\x3e\xc4\x5b\x53\x2b\x6b\x4b\x65\x00\xe5\xd5\xae\x33\xb9\x31\x00\x8f\x24\xae\xc5\xbd\xc6\x22\x00\xf4\x12\x93\x7e\xa1\x91\x2f\xb1\x52\xe4\xe8\x7d\x4c\xd9\xf1\x74\x0a\x9e\x37\xe0\x1f\xe1\xf8\x20\xea\xb6\xa1\x02\xfd\x1c\x06\xf1\x5b\x6c\x17\xc6\x89\xfc\xdc\xe2\x6e\xf9\x30\xc0\x14\x43\x2f\x8e\xf7\x23\x0d\x7c\x88\x5b\x47\xd4\xa9\x3c\x8b\xce\xff\x4d\x34\x70\x8b\xaf\x34\xf0\xff\x2d\x76\xe7\x68\x79\x1a\xa0\x59\xca\xf8\x2e\xe4\x98\x38\x5d\xa9\x3d\xa7\xd6\xfa\x8e\xa3\x4a\x2f\x5d\xd9\xe8\x42\x60\x6d\xe4\x94\x5a\xec\x96\x8f\x72\x0d\xf5\x84\x2e\x52\xeb\x19\x79\xfe\xb5\x90\x5d\xdc\x91\xb2\x3d\xb7\xe5\x58\xd9\xae\xd9\x9e\xc4\xad\xcf\x17\xb3\x6c\x63\x6b\x1d\x52\x20\x23\x7b\x4d\xc5\xad\xf8\x42\xf5\x2d\x73\x9c\xab\x12\xdb\x81\x53\x8b\xef\x17\xc8\x63\x99\x14\x0d\xbc\xf7\x7d\x58\xda\x2b\x5f\xef\x98\x2b\x36\x67\xb1\x2d\x0e\x87\xfd\xc7\x9b\x75\xf3\xdb\x56\xea\x22\x22\x17\xc1\x9d\xbe\xdd\x4e\xb1\x5b\x96\x73\xb1\xdd\xe0\x95\x7e\x33\xc7\x7f\x9a\x6f\xf0\x97\xfd\x33\xf0\xb1\x66\xae\x34\xd1\xa0\x2a\x35\xd8\x9d\x1d\xa0\xcb\x31\xb2\x57\x37\xe3\xc8\xf5\xc7\xa0\x27\xf6\x38\xce\xa0\x7a\xf8\xd7\x6c\x9d\x85\xdf\xb6\x70\xb3\xc4\x16\xb6\x9d\x2a\x8f\x93\xb6\xd8\x7b\x5d\x12\xa3\xf8\xac\x9c\x1e\x22\x75\x4f\x6c\x5b\x7c\x72\x81\x6a\xea\xc5\xfc\x09\x00\x00\xff\xff\xec\xf6\x6e\x9e\x43\x04\x00\x00")

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
		"users.graphql": &_bintree_t{type_users_graphql, map[string]*_bintree_t{
		}},
	}},
}}
