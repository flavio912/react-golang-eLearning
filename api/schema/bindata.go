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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x94\xcf\x8e\xda\x30\x10\xc6\xef\x79\x8a\x41\xbd\xb4\x97\x7d\x00\x6e\x0b\xa8\x2a\x12\xa8\xab\x85\x3c\xc0\x6c\x3c\x1b\xac\x3a\xb6\xeb\x3f\xa8\x51\xc5\xbb\x57\x76\x82\x63\x27\xec\xa1\xda\x13\xf2\x8f\xf9\xbe\x8c\xe7\xb3\xfd\x05\x3a\xef\xd0\x71\x25\x2d\x34\x28\xa1\x53\x8c\xbf\xf7\xe0\x2e\x04\xec\xed\x09\xe8\x0f\x35\xde\x11\x03\x4b\xbf\x3d\x49\xc7\x51\x88\xfe\xa9\xaa\xb8\xd4\xde\x41\xad\x85\x42\xf6\x9d\x0b\x3a\x92\x43\xf8\x5b\x01\xbc\x73\x41\xe7\x5e\xd3\x1a\x4e\xce\x70\xd9\xae\x2a\x80\x46\x49\x47\xd2\x1d\x48\xb6\xee\xb2\x86\xbd\x74\xab\xea\x56\x55\xae\xd7\x94\x59\xbc\x92\xd5\xd1\xc2\x1b\x91\xab\xad\x6f\x1a\xb2\xf6\xac\x7e\x91\x9c\xf8\x6d\xd9\xc3\x69\x28\x8c\x1e\x1f\x8a\xe2\x47\x8f\xe3\x96\x63\x29\xb2\x8e\xcb\x83\x6a\xb9\xfc\x1a\x1d\xd7\xf0\x9c\xc8\xea\xdb\x1a\x9e\xbd\xbb\x44\x9f\x0a\xa0\x43\x89\x2d\x99\xa2\xfa\x98\xb1\x79\x3d\x23\x41\x2d\x3a\x2a\x04\xbb\x1c\xce\x15\x8d\x21\x74\x34\x7a\xde\x15\xdb\x1c\xee\x03\x0b\xb2\x71\x1d\x26\xa6\xd9\x52\x54\xe7\xf0\x81\x48\x1b\x15\xc2\xda\x77\xd8\xd2\x30\xc4\xd7\x90\xb1\x75\x93\x41\x9e\x6e\x10\x97\x61\xcd\x3f\xfc\x92\x19\x2e\x3d\xc6\x74\xca\x1e\xc2\x7c\x16\x8d\xef\x72\x98\x1a\xdf\x28\x25\x08\xe5\x2a\xcd\x28\xa6\x54\x4e\x28\xa2\xa4\x88\xab\xd4\x64\x51\x5e\x4f\x68\x51\x3e\xf4\x54\x94\xef\x26\xf4\x71\x3f\xf7\x58\xcb\x96\xee\x34\xe9\x4a\x1c\xe6\xa8\xa4\xa5\xe4\xb2\x55\x9d\x46\xd9\x97\x26\x23\x9c\x3c\x86\x75\xda\xdb\x4c\x54\xe7\xf0\x81\xa8\xf8\xd2\x3d\xf4\x0a\x20\xdc\xd4\xc8\x1e\x7e\x37\x16\x8c\x17\xe0\xe1\x91\x8c\x05\x86\x1a\xd4\xae\xb9\x60\x7e\x85\xcb\x71\xa1\xd6\x46\x5d\x53\xd3\xde\x73\xb6\x86\xba\xde\xef\xca\x2e\x2d\x5e\xe9\xa7\x14\x5c\xd2\x56\x79\x63\xd3\x5c\x4f\x33\x9e\x6d\x30\x2c\x47\xe5\x56\xa0\xb5\x46\xa9\x6e\x29\x9e\xfd\xb5\xd4\x0f\xf3\x39\x63\x5b\xa6\x70\xc6\x36\xd5\x9e\xb1\x9d\x06\x89\x8e\x5a\x65\xe6\x99\x8d\x74\xb2\x1f\x41\xd2\x1d\xc8\x5a\x35\x3b\xc1\x03\x4b\x9a\x61\x99\x72\x2e\x15\x75\xc6\x96\x8a\xe1\x18\x97\x8a\x5d\xc6\x1e\x1d\x64\xed\x4d\x73\x41\x3b\x0e\xd6\xde\x65\x2f\x25\x4e\xca\x19\x5f\x9c\xe5\xbd\x64\xfc\xca\x99\x47\x51\x6e\x72\xe2\xb3\x5b\x31\xfd\xb1\xf0\xda\x08\x35\x8b\x23\x90\x69\x13\x42\x85\x40\xc2\xcf\x0f\x42\x46\x26\x7b\xd6\xc6\x97\xe7\x3f\x9e\xb5\x37\xa1\xda\x8d\x62\xfd\x27\xde\xc6\x5b\xf5\x2f\x00\x00\xff\xff\x0c\xa7\x42\x89\x5a\x07\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xcd\x6a\xdc\x4c\x10\xbc\xeb\x29\xca\xf8\x62\x83\xf9\x1e\x40\x37\xff\x60\x30\x7c\x49\x1c\xbc\x3e\x85\x10\x7a\x35\x2d\x69\xe2\xd1\xb4\x32\x3f\xb1\x95\xe0\x77\x0f\xa3\x7f\x2d\xbb\x81\xdc\x24\x75\x55\x57\x75\xab\xfa\x1c\x3f\x22\x3b\xcd\x1e\xbe\x96\x68\x14\xac\x04\x34\xa2\x74\xd9\x21\xd4\x0c\xb5\xff\x0f\xfc\xc6\x45\x0c\xac\xa0\x2d\x5a\x72\x8e\x8c\x61\x93\x85\xae\x65\x3c\x52\xc5\x0f\xb6\x14\xfc\xce\x80\x20\x81\x4c\x8e\x07\x1b\xce\x70\x8e\x8f\xb1\xd9\xb3\x83\x94\x68\xa9\x62\x0f\x2a\x03\x3b\x84\x5a\x7b\x88\xe5\x0c\x90\xb2\xf4\x1c\x66\xfc\xae\xe6\xf1\x53\xe2\xf4\xb8\x44\x44\xe9\xa4\xe9\xad\xf8\x40\x2e\x64\x80\xd1\x8d\xde\xd2\x1a\x7a\x4b\xf6\x5f\x2d\x5a\x76\x3d\x2b\x03\x2a\xfd\x93\xed\x11\x37\x3a\x70\xe3\x87\x6a\xf6\x9e\x65\xda\xb6\x31\xf4\x73\xf4\x33\xac\x4c\x6d\xa4\x16\xe8\x07\xb2\x54\xb1\xbb\xd7\x26\xcd\x93\x38\x96\x1a\xce\xf1\x14\x9c\xb6\x55\x06\x70\x43\xda\xac\xde\xbf\xcb\x7e\xa7\x83\x59\x43\x62\xd4\x2a\xc7\xf3\xf3\xc3\x5d\x5a\x1b\x1b\x6e\x6b\xb1\x0b\x60\xd6\xba\x63\xc3\x15\x05\x5e\x89\x6d\xa9\xa1\xf8\x96\x5e\xe7\xc6\xff\x6e\xe5\x2f\xea\xb7\x12\x9d\x67\x7f\x7a\x52\x2a\x0a\xf6\x7e\xd7\xb5\x9c\xe3\x7a\x7e\xce\x80\x3d\x15\x2f\x95\x93\x68\xd5\x6d\xcd\xc5\x4b\x8e\x1b\x11\xc3\x64\x33\xa0\x75\xba\xe0\x1c\xf7\x46\x28\x6d\x98\x8c\x91\x57\x56\x3b\xb9\x89\xdd\x0c\xc3\x39\x9e\x6a\x79\x85\x58\xd3\xa1\x18\x5c\x20\xd4\x14\xd0\x49\x04\x39\x06\xc5\x50\x8b\xd3\xbf\x58\x21\x08\xf6\x22\x2f\x6b\xdb\x4d\x4b\xb6\x3b\xb5\xb3\xcd\x0c\x33\xe9\x7f\xf6\x5e\xec\xc9\x3d\x1f\x2e\x8d\x2a\x9f\xe3\x4b\xaa\x7e\x1d\x7a\x04\x76\x25\x15\xfd\x41\x68\x4b\x81\xd5\xe7\xc8\xae\xeb\x3b\xb5\xe3\x8d\xe4\xf3\xb5\x2c\xb2\x9f\x9c\x62\x77\x33\xe0\xc8\x17\x6c\x95\xb6\xd5\x7a\x5d\xa5\x66\x33\xff\xe0\xb3\x44\xec\xef\x6e\xe9\xae\xfb\xce\x53\x1d\x20\xd5\x68\x7b\xb1\xb8\x3f\xbb\xcc\x71\x9d\xbe\x4d\x35\x7f\x91\x0c\x0d\x66\xa6\x5a\x7a\x4e\xe4\x66\xc8\xf6\x8a\x7e\x99\x4f\x81\x5f\xca\xeb\x0e\x57\x28\xfb\x9d\xe5\xdb\xbb\xb8\x82\x0c\x93\xe5\xd3\x88\x4b\xa3\xc7\xe1\x3e\xd5\x98\xed\x03\xaf\x53\xe4\x57\x88\xe3\x7a\xdb\xdb\x38\x2a\x38\x41\x46\xc5\x62\x08\xc6\x81\xe0\x18\x97\xb9\xae\x4f\xe8\x6d\x62\x75\x54\x6e\x44\x8c\x6a\xa6\x4f\xd4\x81\xd8\x10\xb3\xb9\x7a\x5c\x69\x9d\xc5\xa3\x42\x03\x60\x9e\xaa\xbf\x8f\x13\x9e\x57\x17\x7c\xc2\x73\x42\x8c\xad\xa2\x4f\xa4\x67\xcf\x0e\xd9\x7b\xf6\x27\x00\x00\xff\xff\x53\x0a\x15\xbb\x1b\x06\x00\x00")

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

var _type_blog_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8f\xcd\x4e\x80\x30\x10\x84\xef\x7d\x8a\xe5\x35\x7a\x53\xb8\x90\x18\x63\x54\x4e\xc6\x43\x85\xa5\x54\x4b\x4b\xca\xf6\x40\x08\xef\x6e\xba\xe1\xa7\xc4\x93\xb7\x76\x76\x67\xf6\x1b\x5a\x26\x84\x47\xeb\xf5\x43\xa4\xc1\x07\x58\x05\x40\x6f\xc2\x4c\xcf\x6a\x44\x09\x6f\x14\x8c\xd3\x85\x00\xb0\xea\xaf\x36\x05\xdf\x1b\x8b\x2f\xa6\x6d\x5e\x9f\x8e\x81\xd8\x84\x38\x53\x39\x2f\x46\xd3\x49\x68\x9a\xba\x4a\x26\x32\x64\x6f\x29\x5f\xbe\x5b\xf2\x7f\xab\x08\xb5\x0f\x8b\x84\x72\x7f\x25\x75\x40\xd5\x61\xa8\x47\xa5\x31\x3b\x96\x26\x8a\xc9\x65\xd6\xa2\x48\x0c\xc6\x4d\x91\x58\x64\x13\x93\x7c\xcf\xde\xd5\x55\x6e\x26\xff\x83\xee\x12\x4e\x5f\x19\x50\x11\xb2\x9b\xff\xeb\x7f\xc8\x53\xd5\xab\x70\x46\xfe\x7e\xbf\x76\xb0\x5f\xfb\x7b\x28\x2f\xcf\x12\x3e\x4e\xfc\xe2\x53\x6c\xbf\x01\x00\x00\xff\xff\x13\xfa\x85\x55\xac\x01\x00\x00")

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

var _type_delegate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x41\x4f\xc3\x30\x0c\x85\xef\xfd\x15\xde\x9d\x5f\xd0\x1b\x6c\x97\x4a\x08\x0d\xe8\x4e\x08\x21\xb3\xbc\x85\x40\x1a\x47\x49\x2a\x34\x4d\xfb\xef\x28\x19\xa5\x5d\xb7\x0b\xb7\xe8\xe9\x3d\xdb\x9f\x9d\xb4\xf7\xa0\x15\x2c\x34\x27\x10\x1d\x2a\xa2\xbe\x37\xaa\xa6\xcd\xa6\x59\x2d\x2a\xa2\x6d\x00\x27\xa8\xdb\x54\xd3\x73\x0a\xc6\xe9\x8a\xa8\x6d\x97\x6f\xcd\x6a\x10\xb2\x0b\x1d\x1b\x3b\x71\xec\x4c\x88\xe9\x81\x3b\x4c\x4d\x96\x2f\xb5\x04\x0b\xff\x21\x0e\x93\x70\xf6\xdd\x8b\x36\x6e\x6a\xfc\x94\xf7\xd6\x24\x7b\x16\xde\x4a\xe7\xd9\xed\x6b\x5a\x9e\x1e\x59\xf3\x41\x76\xc6\xa2\xe9\x58\x63\x13\xc6\x99\x8e\x55\x65\x9c\xef\x13\x2d\x0b\xd0\x80\xdc\x14\xed\x30\x16\xcb\xdc\x27\xfa\xff\x60\xcc\xf9\xaf\x4d\x7b\x0d\x55\xc3\x21\x70\xc2\x9a\x63\xfc\x96\xa0\x6a\xba\x13\xb1\x60\x37\x27\xf1\x56\x58\xb5\xf2\x05\x37\x25\x2a\xd7\x3b\x07\x7a\x42\xf4\xe2\x22\x0a\x93\xfa\x15\xeb\xbf\x13\x2f\x6e\x26\x5d\xd5\xd8\x76\x56\x73\xb0\xaf\x59\x83\x4c\xe7\x2d\x3a\xb8\x14\x69\xcd\xda\xb8\x9c\x7c\xec\x11\xf6\xa5\x07\x94\x46\xac\xe9\x65\x88\xbc\xe6\xd1\x59\xa3\x71\x3b\xa9\x73\xa0\xbc\xc6\xfd\x0f\xbe\x72\xe1\x52\xe1\xf2\x3b\xf9\xd9\x5c\x59\x73\x92\xd1\xc6\x05\x1d\xab\x9f\x00\x00\x00\xff\xff\x18\xe4\xd4\x61\xbc\x02\x00\x00")

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
