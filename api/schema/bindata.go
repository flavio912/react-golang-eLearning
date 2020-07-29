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

var _mutation_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x97\xcd\x6e\xe3\x36\x10\xc7\xef\x7a\x8a\x09\x72\xc9\x02\xed\xf6\xee\x5b\x6c\x61\x51\x03\x49\xe3\xc6\xd6\x03\xd0\xe2\x58\x26\x96\x22\x55\x92\x72\x6a\x14\x79\xf7\x82\x94\xc4\x2f\xa9\x3e\x34\xbe\x99\x3f\xce\xfc\x67\x38\x1c\x52\xf4\x23\xb4\xbd\x21\x86\x49\xa1\xa1\x26\x02\x5a\x49\xd9\xe9\x0a\xe6\x8c\x40\x8f\xdf\x01\xff\xc6\xba\x37\x48\x41\xe3\x5f\x3d\x0a\xc3\x08\xe7\xd7\xef\x45\xc1\x44\xd7\x1b\xa8\x3a\x2e\x09\xfd\xc1\x38\xbe\xa2\x21\xf0\x4f\x01\x70\x62\x1c\x0f\xd7\x0e\x57\xb0\x37\x8a\x89\xe6\xa1\x00\xa8\xa5\x30\x28\xcc\x0b\x8a\xc6\x9c\x57\xb0\x15\xe6\xa1\xf8\x2c\x0a\x73\xed\x30\x92\x78\x47\xdd\x39\x89\x5e\xf1\xd8\x5b\xf7\x75\x8d\x5a\x1f\xe4\x4f\x14\x81\x7f\xce\x73\xd8\x0f\x86\x4e\xe3\x3f\x9d\x5c\xd0\xd7\x71\xc9\xce\x94\xd0\x96\x89\x17\xd9\x30\xf1\xe4\x14\x57\xf0\xec\xc9\xc3\xb7\x15\x3c\xf7\xe6\xec\x74\x0a\x80\x96\x08\xd2\xa0\x4a\xac\x5f\x23\x96\xdb\x53\xe4\xd8\x10\x83\x89\x43\x19\xc3\xd4\xc3\x16\x4b\x21\x31\x38\x8a\x4e\x2e\x9b\x18\x6e\x2d\xb3\x7e\xe3\xd8\x96\xac\xa3\x73\xa7\x2a\x86\x0b\x4e\x9d\x92\x76\xb7\xb6\x2d\x69\x70\xa8\xe2\xbb\xdd\x64\x6d\x82\x40\xbc\xbd\xd6\x39\xdd\xad\x3c\xf0\x2e\x12\x9c\x6b\x8c\xdb\x93\xe6\x60\x0b\x34\x4b\xbc\x8c\xa1\x4f\x7c\x2d\x25\x47\x22\x1e\x42\x91\xdc\x3e\xa5\x25\x72\xc8\xbb\xb8\x91\xcf\x32\x31\xaf\x02\x9a\x99\x0f\x49\x25\xe6\x65\x40\x37\x12\x9a\x76\x36\xcd\x69\xa2\xde\x31\xc5\xb6\x92\x52\x68\xf4\x79\xe6\x2a\x55\x42\xbd\xca\x04\x42\xf8\x8d\x6c\x3b\x22\xae\x69\xf4\x11\x86\xe0\xc3\xd8\x47\xcb\x9c\xaa\x18\x2e\x38\x3d\x8e\xd5\x09\x24\x89\x3d\x75\x50\x01\x60\xcf\xbd\x63\x8b\x99\x38\x83\xf1\x38\x2d\xf6\xb7\x33\x50\x58\x93\xce\xd4\x67\x12\x5f\x08\x71\xe9\x01\x48\xd7\x29\x79\xf1\xcb\xe8\x7b\x46\x57\x50\x55\xdb\x32\xcd\x5b\xa3\xd9\xc8\x5e\x69\xdc\xf5\x47\xce\xf4\x19\xe9\x53\xed\xc6\xdb\x72\xb8\x91\x7e\x81\x6e\x9a\xf1\xfa\xd9\x26\x9f\x7a\x7e\x62\x7c\x87\x82\x32\xd1\xbc\x29\x8a\xea\xa9\xe6\x0c\x85\xd9\x63\xad\xd0\xf8\x1c\xd3\x04\x35\xb9\xe0\x9b\xe0\x4c\xe0\x90\xc0\x54\xe9\x7d\xc6\xa3\x62\xdb\x21\x3c\xc2\xe1\xad\x7c\x03\xc2\xb9\xfc\x00\xce\xc4\xcf\x51\x6b\xc3\x89\xd6\x4a\xca\x76\x2e\x97\x4d\x2d\x28\x4a\x2a\x81\x50\x0a\xa6\x37\x52\x81\x91\x30\x14\xc1\x5d\xd2\xf6\xc7\x9a\x08\x81\xea\x4b\x57\x42\xd7\xab\xfa\x4c\xf4\xb8\x2a\x3d\x79\xee\x52\xec\x73\xcb\x78\x74\x1a\xa6\x46\x8b\x57\x59\x46\xec\xc6\x41\x3c\x90\x26\x3d\x05\x07\xd2\x78\xf3\x03\x69\x42\xdb\x12\x83\x8d\x54\xf9\x99\x19\x69\x28\xe0\x08\x6c\x04\x8e\x5a\x4b\x71\x87\x42\x0d\x19\xbc\x38\xb9\x34\xfe\xc0\x7c\xf4\x61\xe8\x4f\x6c\xea\x51\x45\x6c\xee\x31\xd4\x30\xf5\x28\x23\x76\xa3\x86\x5b\x41\xd9\x85\xd1\x9e\xf0\x34\xb9\xc0\xb3\x0b\x2d\x4c\xcc\xae\xb4\xb9\x56\x95\x71\xaf\x95\x4f\xcc\x1a\x62\xae\x55\x66\x7c\x71\x51\xba\x3f\xb6\xcc\x1c\xa2\x2d\xda\x7b\xe2\x1d\x02\xda\x91\xab\xdd\x2e\x78\x1c\x99\x06\xe2\xbf\xe6\xbf\x31\x1f\x4a\x03\x11\xfa\x03\x95\xb6\x67\x89\x80\x41\x6d\x42\x13\x46\xb1\x36\x9e\x64\x45\x8b\x62\xf9\x6a\x1d\x92\x3e\xa2\xb9\x63\x40\xc1\x71\x28\x4d\xec\x58\x7a\xb2\x58\x8e\x21\xed\x3b\xf4\xef\x9f\xd6\x8b\xe5\x1d\x3c\xd1\x6c\xb5\x13\xce\x57\x9c\x8b\x54\x09\xcd\x56\x3e\x17\x19\x56\x9f\x8b\x94\x09\xbd\xd1\xe9\xaf\x92\xf6\x3c\xfb\x68\x0f\x2c\x4b\x7f\x80\x21\x6e\xeb\xc6\x77\xb8\x0a\xc6\x67\x54\x92\x47\x15\xb1\xac\x02\x79\x1e\xe3\x2b\x2a\x71\x2f\x23\xb6\xb8\xf6\x8b\x64\x35\xca\x0b\xaa\xff\x97\x74\x68\x73\xfb\x29\xc9\xfa\xdc\xa2\x70\xdf\xda\x51\x68\xee\xd8\xbc\x0a\x68\x66\xee\x3e\x51\x7b\xd6\x08\x62\x7a\x75\x8f\x57\xea\x21\x11\x5c\x48\x21\x04\x4b\x73\x09\x5f\x0b\x54\x86\x9d\x58\x6d\x1d\xae\x5d\xd6\x31\xd9\x64\x68\x9d\x94\x87\x87\xd7\xb2\x58\xb5\x34\x79\x43\xac\x0e\x64\x2d\xe9\xf5\x0b\x75\x0a\xcb\x7c\x7e\xfe\xa3\x6f\x8f\xf9\xdf\x0f\x8f\x43\x36\x13\x09\x8b\xca\x5d\xab\x14\x2f\xb8\xba\x47\xa5\x46\x33\xbd\x67\xa3\x47\x4c\xdb\x71\x74\xff\xd3\x7e\xb5\x06\x1a\xd8\x29\xba\x85\xa1\x23\x5a\x23\x05\xa9\xe0\x44\x18\x47\x0a\x04\xea\xc9\x77\x7c\xce\xc0\xd3\x0f\xa9\x86\x5d\xd4\xdf\xee\xf7\x36\xf1\x95\x5a\x73\x99\x3d\x34\x2c\x09\xc7\x8d\xcb\xc6\x97\x26\xb6\xad\x3c\xc9\x6d\x8f\x5c\x36\xbf\x23\xa1\x77\xb9\x50\xd6\xa9\xd8\x3c\x7a\x34\xb9\x94\xc8\xd7\x1b\x4a\x61\x83\x02\x55\xda\xd1\x93\xc0\xfb\xd2\xe4\xfc\xae\xfa\x2c\xfe\x0d\x00\x00\xff\xff\xa1\xc2\x5a\xaf\xa5\x10\x00\x00")

func mutation_graphql() ([]byte, error) {
	return bindata_read(
		_mutation_graphql,
		"mutation.graphql",
	)
}

var _query_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x4d\x8f\xe3\x36\x0c\xbd\xfb\x57\x30\xc8\x25\x01\x06\xdd\xbb\x6f\x99\x2c\x16\x18\xa0\xbb\xdd\x6d\x32\xa7\xa2\x28\x14\x8b\xb6\xd5\x91\x25\x8f\x3e\x66\xc6\x2d\xf6\xbf\x17\x94\x2c\x7f\xd5\x09\x76\xf6\x66\xeb\x91\x7c\x24\x6d\x3e\x6a\x0b\xcf\x1e\x8d\x40\x0b\xb6\xd6\x5e\x72\x50\xda\x41\xa3\xb9\x28\x3b\x70\x35\x02\xbf\xfc\x02\xf8\x86\x85\x77\xc8\x41\x28\x68\x99\x31\x4c\x4a\x94\x99\xeb\x5a\x84\xaf\xac\xc2\x07\x55\x6a\xf8\x37\x03\x70\xda\x31\x99\xc3\x83\x72\x1b\xd8\xc2\x17\xdf\x5c\xd0\x80\x2e\xa1\x65\x15\x5a\x60\xa5\x43\x03\xae\x16\x16\xb4\xc2\x0c\x40\x97\xa5\x45\x37\xd8\x9f\x6b\xec\x8f\xc8\x27\xd8\x91\x23\x94\x46\x37\x21\x15\xeb\x98\x71\x19\x80\x14\x8d\x98\xbb\x35\xec\x8d\xd2\x7f\x55\xd0\xa2\x09\x5e\x19\x40\x25\x5e\x50\xad\x64\x23\x1c\x36\x36\xa2\xd9\xf7\x2c\x13\xaa\xf5\x2e\xd4\x11\x6a\x98\x24\x35\xa3\x1a\x4d\x3f\x33\xc5\x2a\x34\x9f\x84\xa4\x7a\xc8\x47\xb1\x06\x73\x38\x39\x23\x54\x95\x01\x60\xc3\x84\x9c\xbc\xff\xad\x2f\x67\xe1\xe4\xd4\xc4\x7b\xc1\x73\x78\x7c\x7c\xf8\x48\x6d\x43\x89\x6d\xad\xd5\x68\x30\x70\x7d\x44\x89\x15\x73\x38\x21\x9b\xbb\xba\xe2\x2f\x7a\x1d\x02\xbf\x3f\x95\x1b\xec\x47\xed\x8d\x45\x7b\xbd\x52\x56\x14\x68\xed\xb9\x6b\x31\x87\xc3\xf0\x9c\x01\x5c\x58\xf1\x54\x19\xed\x15\x3f\xd6\x58\x3c\xe5\x70\xaf\xb5\x44\xa6\x32\x80\xd6\x88\x02\x73\xf8\x24\x35\xa3\x0e\x33\x29\xf5\x2b\xf2\xb3\xbe\xf7\xdd\x60\x06\x5b\x38\xd5\xfa\x15\xb4\x92\x1d\x14\x31\x0b\x70\x35\x73\xd0\x69\x0f\xcc\x20\x30\xef\x6a\x6d\xc4\x3f\xc8\xc1\x69\xb8\x68\xfd\x94\x01\x14\xcc\x61\xa5\x4d\x47\xcd\xe9\x5b\x34\xa9\xa5\x69\x99\xea\xae\x35\x72\x59\x58\xdb\x1a\xfd\x82\x7c\x4c\x7c\x08\xf4\x2b\x5a\xab\xd5\x0f\xc6\x71\xac\xb2\x39\xfc\x41\xe0\x9f\x63\x88\x53\x27\x25\xbb\xf8\x1b\x8d\xc5\xb7\x42\x7a\x8e\x91\x6b\xda\xbd\x1e\xf8\xac\xb9\xa7\xaf\xf8\x3f\xe0\x8c\xd6\xad\x24\x7d\x44\xe3\x44\x29\xa8\x41\xf4\x85\xae\x13\x1b\xac\xbc\x64\x4e\x68\x75\xc6\x37\x37\x03\x9e\xbd\x30\x68\x8f\x87\xc3\x17\x3d\x25\xa6\xb9\x3b\x1b\x26\x94\x50\xd5\x09\x0b\x72\x5d\x4b\xe0\x70\x88\x13\x38\xa1\x16\x1c\x15\x65\x85\x66\x3a\x19\x76\xb5\xeb\x0f\x8a\x8b\x17\xc1\x3d\x93\x3f\xd8\xf9\x9f\xff\xf3\x1d\x9a\x92\x15\x41\xda\x84\x62\x0e\xf9\x37\x8f\xa6\x0b\x84\x6d\xaf\x76\xf9\xa0\x7b\x63\x86\xbf\x19\x8e\xe6\x3e\xda\x31\x5b\xa0\xe2\x42\x55\xd3\x46\x95\x02\xe5\x30\xaa\x1b\x72\x0c\x0a\x3a\x46\x17\x21\x72\xc2\x01\x18\x6f\x84\xda\x8d\x45\xee\x73\x38\xd0\x51\x82\xec\x8e\xf2\x89\xb9\x24\x8c\x9e\x37\x59\x06\xd0\x44\x95\x9a\xbb\xf7\xd2\x35\xc2\xd3\x10\x77\x50\x86\xd6\xe6\x73\x85\xbb\x03\x1d\x2b\xcb\x53\x89\x63\x20\x72\x23\x32\xde\xcb\xd4\x84\x6d\xb3\xcf\x07\xf5\x9a\x58\xac\x13\xce\x65\x6e\x95\x31\x99\x24\xca\x22\xce\xf3\x82\xb1\x9f\xf2\x01\x17\x57\x08\x67\x6a\xb0\xca\xd7\x5b\x24\x3a\x19\x26\x71\xc1\x16\xc7\x73\x40\xd7\xa9\xa6\x7a\xb1\xca\x14\x0d\x12\xd1\x45\xea\x6a\x41\x73\x2f\x75\xd5\x23\x73\x8a\x95\x60\x64\x3b\xb6\x28\x08\xe7\x95\x06\x4c\xa4\xfd\x4a\x03\xc8\xe2\x6b\xdc\xa5\x31\xd4\x8e\xb2\xa2\x7d\x3a\xc0\x71\x60\x4d\x0e\x8f\x36\xfc\x56\x62\x18\xd3\x45\x0d\xe3\xfc\xce\xac\xec\x2e\x03\x88\x73\x15\x33\x0c\xaf\x29\xc9\xe5\xd0\x07\x70\x99\x6a\x06\x30\x8b\x9f\xaa\x7f\xf6\x68\x49\x8c\xd6\xeb\xff\xd6\xa3\x37\x1a\x90\x4c\xfa\xb4\x52\xbc\x45\x61\xc9\x2a\xe8\x89\x75\xeb\x6c\xa4\xca\x37\x98\x08\xee\x59\x28\xc6\x82\x81\x50\x9a\xd8\x20\xf9\x57\x06\x36\x60\xb7\xe6\x35\x18\xf4\x1c\x31\xd2\x82\x25\x5a\x50\xe3\x9c\x77\xda\x2c\x73\xa0\xb3\x84\x5d\xa9\x91\xa0\x5b\x45\x12\x9e\x3e\x8e\x45\x66\x8a\x3a\xed\xc1\xd5\x78\xf3\x25\xb9\xcf\xe1\x34\xf3\xf9\x1d\xad\x97\x0e\xb6\x70\xa0\x0b\x84\xed\x23\x0a\x55\xa5\x46\xdd\xa5\xa9\x04\xa6\x78\xfc\x36\x19\xc0\x36\xac\xe4\x9d\x9b\x2c\xb7\x3d\x7c\xf8\xd0\xbb\xd3\x2d\x55\xca\x78\xfd\xc4\x60\x49\xff\xfe\xb8\x3b\x49\xf1\x77\x4e\x3f\xd1\xbd\xb2\xd7\x69\x1a\x85\xb9\xc1\x66\xee\x43\xfb\xf6\xf6\x6f\xbe\xba\x9c\xe3\x6f\xbd\x80\xd2\x34\xce\x4f\x97\x1a\x38\x47\x83\x12\x30\xa6\xc2\xf6\xbd\x22\x06\xf3\xe5\xbc\x9f\x9c\x0c\x5a\x12\xaf\x56\x4b\x3d\x9d\xf5\x31\x87\x63\x7f\x01\x4b\x79\xf6\xaf\xcb\x04\xfb\xe3\xf0\x39\x2c\x16\x5a\x71\x66\xba\xe3\xfb\x19\xbe\x67\xff\x05\x00\x00\xff\xff\x01\x45\x9a\xbc\xbe\x0c\x00\x00")

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

var _type_certificate_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\xcd\x6e\x9c\x40\x0c\xbe\xf3\x14\xce\x6b\x70\x23\x44\x95\x56\x8a\x56\x69\x96\x3d\x55\x3d\xb8\xe0\x4c\xac\xc0\x0c\xf5\x78\xd4\xa0\xaa\xef\x5e\x0d\x69\x80\x01\xb6\xdd\x1e\x72\x59\x2d\x98\xef\xf3\xf7\x63\x1d\x7a\x82\x92\x44\xf9\x89\x6b\x54\x3a\xd8\x27\x07\x3f\x33\x80\xda\x05\xf1\x54\xb1\xb6\x94\xc3\x49\x85\xad\xb9\xc9\x00\xe8\xb5\x67\x19\xee\x50\x93\xb7\xb5\xeb\xfa\x96\x94\x9d\xdd\x9b\xa0\x1d\x8e\xd8\x4d\xaf\x33\x00\xc5\x17\x92\x4f\x2c\x5e\x97\x83\x9b\xf7\xc9\x3d\x6e\x07\xf5\xac\xf1\xd6\x35\xc3\xf9\xf1\x7e\xc1\x27\x64\x42\x8b\x71\x7f\x45\xaf\xba\x84\x95\x45\x71\x74\xcb\xcd\x6b\x43\x6c\xbd\x4a\xa8\xd5\xc9\x7a\xe3\x3c\x29\x0f\xc7\xfd\xc1\x89\x8d\x45\x0d\x42\xa9\x9c\x85\xd6\x63\xe8\xbe\x91\xcc\xe8\x5f\x59\xc6\xb6\x0f\x0a\x8f\x64\xc8\x92\xa0\x52\x92\x7e\x1c\xc5\xf8\x9f\xd9\xab\x13\xae\xb1\x2d\xc7\x22\xce\xe7\xc3\x5d\x0e\xf1\x77\xe4\x78\xab\xad\x28\xde\xe8\x47\x44\x08\xdc\xbc\x7f\x01\x50\x0b\xa1\x52\x53\x24\x61\x70\x43\x36\xae\x5a\x0a\x02\x08\x9e\x9a\x1c\x6e\x9d\x6b\x09\xed\x0e\xfb\x03\x1a\x02\x8e\x05\x77\x64\xd5\xc3\x03\x1a\xb6\x91\xfc\x73\x20\x19\xe2\x6a\x6a\x0c\xf9\x1c\xbe\x4c\x90\xaf\x19\x40\x8f\x66\xbc\xa6\x3c\x02\xc6\x7f\xb3\xf9\x72\x54\x37\x7d\x3e\xdb\xde\x53\x38\xa1\xce\x7d\xb3\x8f\x4a\xad\x6f\x39\xd6\x26\x67\x8f\x73\xf4\x55\x7c\xde\x92\xd9\xf5\x19\xee\xe5\xba\xba\xcd\x43\x87\x86\xae\x3d\x50\xa1\xef\x81\x85\xfc\x9f\x43\x9d\x6a\x00\xf0\xcf\xee\x47\x25\xc8\x96\xad\x39\x51\x1d\xc1\x7b\x35\xa5\x16\xfe\x5e\x16\x24\x6d\xa5\xc8\x2b\x3b\x4b\x41\x73\x07\xeb\xa0\xfe\xd7\xf1\x3f\x0c\x6f\x53\xae\xdc\x0b\xd9\x29\xe2\xf5\x91\x5c\x92\x79\xb9\xdd\x8b\x9a\x3f\x4c\xf2\xef\x00\x00\x00\xff\xff\x0a\xfd\x10\x75\x7c\x05\x00\x00")

func type_certificate_graphql() ([]byte, error) {
	return bindata_read(
		_type_certificate_graphql,
		"type/certificate.graphql",
	)
}

var _type_company_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x52\xcd\x6a\xf3\x30\x10\xbc\xeb\x29\xf6\xbb\x7d\x85\x5c\xda\xa3\x6e\x89\xd3\x42\xa0\xbf\x94\x9c\x4a\x0f\x8b\xb5\x11\x02\x5b\x12\xd2\xba\x60\x42\xde\xbd\x48\x72\x52\x13\x3b\xa5\xa5\xb7\xdd\xd1\x68\xb5\x33\x23\xee\x3d\x41\xe5\x5a\x8f\xb6\x87\xbd\x00\x40\xef\x83\xfb\x20\x25\x61\xe5\x5c\x43\x68\x05\x80\x89\x95\xb3\x1c\xb0\xe6\x13\xfa\x4f\x00\x74\x9d\x51\x12\xb6\xdb\xcd\x3a\x75\x75\x20\x64\x52\x4b\x96\xf0\xca\xc1\x58\x2d\x00\x2c\xb6\x74\x6c\x13\xa7\x45\x8b\x9a\x42\xfc\xef\x51\x93\x84\x67\xd4\xb4\x80\x9d\x69\x98\x82\x84\x87\x72\x78\x97\xdb\x05\xb8\xa0\x28\xac\x7a\x09\x4f\xa5\xb8\x3a\x31\xd2\xb5\x34\x4d\x51\x43\x1a\x99\xe6\xc7\xad\x87\xd3\x6f\xe6\x1d\x29\xc7\x81\xa8\x54\xa0\x18\x25\x2c\x4b\x91\x65\x39\xcb\x58\xf3\x6d\x8b\xa6\xf9\x92\x72\x10\x22\x3b\x37\x10\x8b\x73\xa5\xbe\x37\x96\xae\xc7\xa2\x47\xf8\xcd\x18\xaf\x5d\x67\xb9\x1f\x23\xde\x45\xae\x9c\xa2\x09\x2b\xf4\xd3\xa7\x87\xd0\xd2\xee\x60\x5a\xdf\x50\x4b\x96\x63\x32\xc1\xd8\x14\xc4\x4b\x47\xa1\x44\x4a\x4a\x53\x94\xf0\x36\xdc\x78\x4f\x2f\xa1\xa6\x8d\xdd\xb9\x62\x5a\xaa\xd2\x5c\x63\x7d\xc7\x50\xe5\x20\x07\xf2\x26\x43\xfb\xbc\x48\x06\x1e\xcf\x22\xfd\xbd\xea\x39\x3f\xff\xe2\x46\xd9\x7a\xeb\xd5\xdc\xd6\x67\x7f\x74\xaa\xe1\xe7\x5f\xfe\xc2\xf2\x17\x3c\xb8\x60\xc1\x44\xe9\x8c\xd0\xa9\x4e\x71\x10\x9f\x01\x00\x00\xff\xff\x67\xd9\xca\xc5\xa9\x03\x00\x00")

func type_company_graphql() ([]byte, error) {
	return bindata_read(
		_type_company_graphql,
		"type/company.graphql",
	)
}

var _type_course_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4d\x6f\xe3\x36\x13\xbe\xeb\x57\x4c\xe2\x43\xde\x17\x48\x0e\xdd\x7c\xa0\xf5\x2d\xab\xb8\xa8\x81\x2c\x36\xf5\x07\xda\x62\x91\xc3\x98\x1a\x4b\xc4\x52\xa4\x4a\x0e\x93\x35\x16\xfb\xdf\x0b\x92\xb2\x24\x3b\xf6\x36\x69\x2f\x85\x81\x48\x22\xe7\x19\x0e\x9f\x79\x66\xc8\x90\xf6\x35\xdc\x0a\x41\xce\x2d\x36\x0d\xc1\xd7\x0c\xc0\x92\x63\x2b\x05\x53\x91\x01\x98\x86\x74\xf6\x2d\xcb\xa2\xe1\x9c\xad\x17\xec\x2d\x4d\x14\xd5\xa4\x39\x9a\xd7\xa6\xf0\x8a\x32\x00\x26\xc7\x19\x80\x22\xe7\x4c\x8f\xc9\x8d\xb7\x8e\x3a\xe7\x46\x2b\xa9\x83\xb5\x50\xe8\x9c\x35\xa6\x0e\x96\x52\x33\xd9\x35\x0a\x82\xf9\x46\x29\x5c\x79\x37\x65\xaa\x23\x40\x63\x4d\xe3\xb0\xb2\xd4\xe5\x49\x06\xe0\xbd\x2c\xc6\xb0\x5c\x4e\xef\xc2\x17\x6f\x9a\x34\xbb\x13\x57\x98\x11\xa6\x6e\x14\x31\x8d\xe1\xbd\x31\x8a\x30\x46\x14\xcc\xdb\x88\xa2\xf3\xe0\x6a\x9a\xec\x93\xa7\x3e\xda\x93\x03\x6b\xaf\x50\x7c\x2e\xad\xf1\xba\xc8\x2b\x12\x9f\x7b\xd7\x00\xd8\x71\x38\x1e\xf0\x19\x40\x8d\x95\x82\xc6\xf0\xb3\x32\xd8\x06\xa6\x8c\xdd\x7a\xcd\x00\xe8\x8b\x20\xdb\xf0\x60\x44\x6a\xb6\xa6\xf0\x82\xa5\xd1\x83\xe1\xca\x3c\x2f\x4c\xde\x6d\x6b\x30\xee\xad\x1b\xce\xc4\xa5\x32\x80\xe7\x0a\xf9\x0f\xe3\xef\x09\xad\x1e\xc3\xa7\x76\x1f\x8f\x31\xc5\x7f\x7a\x69\x23\x57\x6e\x77\xc6\x35\x24\xe4\x5a\x8a\x05\xd9\xda\x0d\x16\x11\xc8\x54\x1a\xbb\x19\x43\xde\xbe\x85\x4d\x2b\x65\x9e\xa9\x58\x98\xf7\x7e\x33\x24\xc3\xb5\x49\x0c\xae\x07\xf9\x8c\x0b\xac\x50\x6b\xb2\xd3\x1a\x4b\x5a\xce\xee\x77\x78\x68\xa4\x25\x37\xd5\x1f\x8c\xe6\xca\x75\x89\x89\xe3\x18\xb8\x58\x98\x89\x2e\xe2\x6c\xb7\x58\x24\xd8\xaf\x94\x74\x15\x15\x3b\xa3\x82\x2c\x87\x8d\x20\x53\xca\x4a\xbe\x3b\xb0\x27\x87\x07\x2c\x09\x64\x60\x30\x92\x02\x0f\x58\x4a\x8d\x4c\xc5\xaf\x9e\xec\x26\x8a\x85\x8a\x92\xc2\x96\x12\x20\xec\xa5\xc1\x92\xa6\x7a\x6d\xc6\xc1\x3c\xbe\x25\x31\x37\x9e\x61\x8e\x4f\x94\x6f\x35\x9e\x20\xd3\x38\x31\xd0\xdd\x9e\xc4\x0e\x45\x1d\x84\x3e\xa0\x3f\x7c\x76\x83\xaf\x56\xce\x9b\x75\xbb\x27\xdb\x03\xaa\x65\x2c\x03\x17\xb1\x0e\xbf\xa7\x9b\x41\xba\xe7\x3e\xfa\x1f\x3a\xf1\x6c\x6c\xbf\x25\x18\xa5\x87\x59\x03\x57\x04\x8b\x30\x0b\xff\xeb\x1a\x05\x88\x54\xb8\x46\xab\xcd\xff\x43\xe3\xc1\x2f\x0f\x68\x59\x0a\xd9\x60\xd4\x71\x62\xd4\x31\x5a\xbe\xc3\x50\x08\x0b\x59\x87\xad\x90\x2e\x76\xbe\x95\x11\xb8\xc7\xd0\x21\xf1\xfd\x9d\xf6\x76\x73\xfd\x31\x36\xb6\xd7\x27\xfa\x50\x42\x8f\x65\x7f\x04\x8e\x84\xd1\x05\xda\x4d\xfe\x0a\x21\xc0\x08\xe6\x95\xb1\x0c\x05\x39\x61\x65\x13\xe2\x0f\xa4\x26\xfe\x8e\xc8\x04\x46\x70\x76\xbb\x32\x9e\x81\x2b\xe9\x5a\xdb\xb3\xb0\x70\xb0\x3a\xd6\x7e\x02\xea\x17\xf3\x0c\x6c\xba\x86\x7b\x1c\x7f\xa8\x4d\xc1\x08\x26\x8e\x65\x1d\x4a\x2d\x99\x1c\xf3\x75\xbc\xa1\x85\x20\x7e\xab\x90\x61\x63\xfc\x99\x52\xa0\xc2\xfc\x19\xac\xbc\x52\xc4\xd0\x18\xa9\xd9\x1d\x6f\x7a\x01\x3d\x1b\x4c\xbd\x04\xfe\x57\xcb\xc7\x6d\x8f\xbd\xae\x29\x75\x5d\xf6\xdf\xe9\xb9\x77\x16\x45\x7c\xfc\x90\x1d\x1e\xc6\x3d\xda\x52\x10\x30\x96\x7d\x19\xec\x1f\xa4\x3b\x74\x9c\x74\xad\x78\x81\x65\x34\xef\xdd\xbe\x1e\xbb\x2d\x8c\x7f\xe2\x60\x18\xf7\xd6\xcf\x5b\x82\x4f\xf8\x07\x6f\x45\x85\xae\xed\x01\xae\x77\x90\xf4\x1b\xf2\x1d\x8e\xb4\xc7\x48\x9c\x23\xdb\x0b\x00\x46\x70\x2f\x1d\x87\x1a\x2d\x48\x51\x89\x4c\x2e\x46\xef\x62\xce\xd8\xe2\x54\x3f\x19\x29\x68\x52\xa3\x54\x03\x05\x04\xf9\x35\x4c\x45\x2b\x9b\xc1\xf1\xb7\xa7\xd9\xdc\xe8\xb5\xb4\xf5\xcb\xcb\xd0\x5e\xd0\x33\x72\x8d\xd1\xed\xed\x88\x2d\x6a\x87\xb1\x7e\xf3\xfd\xdb\xd4\x09\x8c\x80\xad\x27\x90\xa9\x59\x87\x8a\x45\xbd\x01\xe9\xc0\x68\x40\x10\x46\xb3\x45\xc1\x80\xba\x00\xae\xbc\x83\xc2\x90\xd3\x67\x0c\x9a\xa8\x08\x35\xde\xe0\x06\x7c\xb3\xb6\xa6\xed\xda\x56\x36\x94\x2b\x49\x9a\xe7\x24\x2c\x0d\xbb\xd9\x74\x7d\x28\x96\xb0\xd6\x1a\x95\xa3\xf3\xd4\x26\xa4\x83\x52\x3e\x91\x06\x67\x22\xbf\x20\x50\x87\x55\xba\xbd\xce\x09\xad\xa8\xb6\x37\x92\x19\x39\xaf\xf8\x95\x67\xfe\x8b\x6b\xcc\xf7\x8e\xfe\x3b\x0a\xe1\x1d\x39\x0a\x5e\x2a\xf6\x2d\x57\x8f\x16\x72\x3c\x82\x11\x7c\x82\xc9\xef\xb7\x1f\x1e\xee\x27\x90\x7f\x5c\xce\xe6\x13\x98\x2f\x66\xcb\x7c\xb1\x9c\x4d\xb2\x11\x00\x7c\x4d\xd5\x7c\x9a\xae\xed\xa7\xe7\x6d\xad\x9c\x5e\x5e\x5d\x5e\x5d\x5e\x0c\xff\x5e\xc5\xc7\xe9\xb7\xf3\x1d\x5c\xba\xde\xf7\xb8\x1f\xde\xc5\xdf\x45\x7a\xbc\xbb\x18\x7c\xed\x43\xc3\x3f\x08\x3d\xf0\xc7\x9f\xd2\xef\xe2\xfa\x26\xfe\x2e\xae\x6f\xba\x97\xeb\x9b\xeb\x04\x7e\xcc\xfe\x0a\x00\x00\xff\xff\xc5\xa8\xbd\xb0\xa2\x0c\x00\x00")

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

var _type_individual_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\xd1\xcf\x4a\xf4\x30\x14\x05\xf0\x7d\x9e\xe2\x7e\xaf\x91\xdd\x87\xb3\x29\x88\x8c\x7f\xba\x12\x17\xd1\x9c\x89\x57\xd2\x24\x24\x37\xca\x30\xcc\xbb\x4b\x2a\xda\xda\x8e\x0c\x08\xee\xda\x4b\x4e\x38\xbf\x5c\xd9\x27\x50\x17\x2c\xbf\xb2\xad\xc6\xd3\x41\x11\xd5\xca\x56\x53\xdf\x77\x9b\x7f\x8a\xe8\x29\xc3\x08\xec\x7f\xd1\x74\x2b\x99\x83\x53\x44\x18\x0c\xfb\xcf\xff\x76\x68\xc7\xb9\xc8\x95\x19\x30\x1f\x7a\xb3\x9e\xbd\xc4\xc7\x3b\x16\x8f\xd9\x65\x02\x8f\xf4\x1c\xc3\x7c\xd6\xa2\x97\xd1\x71\x98\xb2\x47\xa5\x16\x65\xb7\xc6\x81\x78\x48\x1e\x03\x82\x14\xda\x1a\xc7\xa1\x75\xbd\xae\xc8\xfb\x26\x81\x75\x28\x9a\xee\xa7\xcc\x83\x22\x4a\xc6\xa1\x0b\xbb\xa8\x5b\x62\xfc\x6a\x97\x73\x48\x55\xe8\x62\xd4\x4e\xe7\xbb\x71\x7a\xf8\x03\xe2\xea\x0d\x93\x29\xe5\x2d\x66\xbb\x26\x2f\x4b\xdd\xa0\xa4\x18\x0a\x3e\xb6\x55\x90\x35\xf5\x05\x79\x52\xf4\xc9\xfe\xa0\xf8\xbe\xdb\x95\xe9\x04\xe9\x97\xa2\x13\xa0\x2f\xcf\xb2\xde\x59\xcf\x06\x1e\xe7\x3d\xc7\xf7\x00\x00\x00\xff\xff\xe1\x40\xf9\x02\xcd\x02\x00\x00")

func type_individual_graphql() ([]byte, error) {
	return bindata_read(
		_type_individual_graphql,
		"type/individual.graphql",
	)
}

var _type_lesson_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x92\xc1\x4e\xf3\x30\x10\x84\xef\x7e\x8a\xed\x6b\xe4\xf8\xff\xe5\x10\xa9\x87\x42\x1b\x2e\xa8\x87\x6d\x32\x58\x16\xc9\x3a\x72\xd6\x95\x22\xd4\x77\x47\x76\x40\xb8\x28\xc0\x8d\x53\x92\xdd\xc9\xe8\x9b\xb1\x75\x1e\x41\x3b\x4c\x93\x17\x72\xc3\xd8\x63\x80\xe8\x44\x87\xb9\xef\xf9\x1c\xa7\x5a\x31\xd0\xab\x21\x12\x1e\x50\xd1\x41\x83\x13\xbb\x31\x44\x31\xba\xae\xa2\xa6\xa9\xb7\xe9\x2b\xb9\xe4\x6d\x6c\x35\x06\xdc\x2d\x36\x69\xd3\xfa\x64\xaa\xa8\xe8\x9f\xf7\x3d\x58\x0c\x51\x87\xa9\x0d\x6e\x54\xe7\xa5\xb4\x54\xb6\x53\x45\x4f\x47\xb6\x9b\x93\x21\x3a\xb3\x08\x42\x3d\xb0\x45\xf3\xb0\xfb\x10\x1a\xa2\x8b\x77\x2d\xfc\x05\xe1\x76\xac\x81\x65\xf1\x2d\xb5\xae\x83\xaf\xe8\x31\x3d\xcc\xd5\x98\x22\xee\x9e\x2d\xca\xc8\x7b\xb6\x4e\x58\xd1\xdd\x47\x84\x39\x87\x46\x67\x91\x90\x96\x1f\x12\xd4\xc8\x16\xb5\x3c\xfb\x2a\xc9\xf3\x5b\x72\x75\x32\x46\xa5\xff\x01\xac\x58\xb4\x75\x9e\xac\x15\xf7\x9e\x32\x37\x77\xfa\xbe\x8c\x22\xfd\xd1\xbf\x40\xd6\xf2\x7f\x5d\xfc\xda\x40\xa6\xfa\x04\x6e\xc6\x6e\x05\xf8\xf6\x64\x4b\xfc\x75\xda\x95\x4c\x7f\x01\xbf\x45\xba\x56\x3f\xc1\x5f\xcd\x5b\x00\x00\x00\xff\xff\x21\x02\xa8\x9e\xdd\x02\x00\x00")

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

var _type_module_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\xcf\x6e\x9c\x30\x10\xc6\xef\x7e\x8a\xe1\x35\xb8\xb5\xdd\x56\x42\x6a\xa5\x6d\x81\xe6\x10\x71\xf0\xe2\x89\x65\xc5\xd8\xc8\x7f\x56\x42\x11\xef\x1e\xd9\xde\x05\x43\x36\xd1\xe6\x94\xd3\x02\x33\xf3\xf1\xcd\xef\x63\xed\xa6\x11\xe1\x8f\x66\x5e\x22\x88\x61\x94\x38\xa0\x72\x16\xea\x49\x4a\x7a\xf2\xb6\x72\x38\xc0\x0b\x01\xf0\x5e\xb0\x12\xda\xb6\x3a\x14\x04\x40\xd1\x01\x4b\xa8\x9d\x11\x8a\x87\xfb\xa0\x12\xef\x7d\xef\xbc\xc1\x9f\x49\x26\x54\x4e\x54\x29\x34\xd5\x40\x39\xb6\xff\x7e\x5f\x67\x08\x00\x43\xdb\x1b\x31\x3a\xa1\xd5\x46\xc9\x50\x95\x0a\xf9\xd3\xb3\x16\x3d\xea\x33\x9a\xad\xc6\x59\x30\xd4\x25\xfc\x0f\x3f\x04\xc0\x5e\x4c\x97\xf0\x98\xfb\x2f\x3a\x02\xd0\xeb\xb0\x9c\xc3\x12\xbe\x6b\x2d\x91\xaa\xf0\x2e\xca\x43\x6f\x43\x79\xd1\x15\x64\x26\x04\x95\x1f\x92\x5a\x13\xb0\x84\xbd\x1f\xaa\xba\xa9\xbe\x85\x62\x24\x15\x8b\xb1\x90\x56\x5e\x9a\x83\x4b\x6f\xe4\x6a\x7a\x26\x44\xa8\xd1\xbb\xd4\x52\xc5\xcb\x3b\xe7\xa2\x8d\x94\xc9\x85\x64\x9a\x44\xeb\x08\x80\x44\x6b\xb5\x5a\xf5\x53\xe3\x12\x54\xd2\xdf\x4c\x17\xbb\xfc\x96\xd1\x1f\x06\xa9\xc3\x8b\xc0\xe2\xf0\x4d\xb8\x09\x53\x9c\xed\x3e\x9b\x5c\x96\x7f\xed\xfb\x1e\xad\x6d\xf4\x33\xaa\x3c\xc4\x6b\xb6\xef\xd5\xb3\x90\xa3\xc9\x6d\xd2\xeb\xfa\x45\xb7\xc4\x94\x2f\x76\xa4\x93\xd4\x94\xc5\xd5\x86\xf8\xe4\x8a\x67\x05\xd1\x8e\xec\x06\x88\xed\x37\xbf\xc7\x90\x63\xba\x4d\xe5\x26\x94\x2f\x63\x92\xef\x78\x0f\x93\x03\x86\xff\xcb\x47\x4c\x66\x92\x9d\x1e\x47\xca\x37\x27\xc8\x91\x72\xa1\xa8\x43\xf6\xd7\xa3\x99\xe2\x2c\x32\x8e\xab\xbd\xc0\x70\xa4\x1c\x2b\xf5\xa4\xcb\xd0\x1e\xaf\xf6\xdf\xf5\x2f\x21\x1d\x9a\xdd\xab\xef\xa3\x3f\x93\xd7\x00\x00\x00\xff\xff\xca\x33\x8f\xf7\xdd\x04\x00\x00")

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

var _type_tutor_graphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x90\xcf\x6a\xc4\x20\x18\xc4\xef\x3e\xc5\xd7\xd7\xf0\xda\xb2\x10\xe8\x61\xdb\xac\xa7\xd2\x83\xc4\x59\xf9\x68\x62\x44\x3f\x0f\xcb\x92\x77\x2f\x0a\xa9\xfd\x0b\xdd\x9b\x33\x8e\x3f\x87\x91\x4b\x04\x9d\x8a\xac\x89\xae\x8a\xa8\x14\x76\x9a\x8c\x19\x1e\xee\x14\x51\xb0\x0b\x34\x8d\x92\x38\xf8\xaa\x27\x0e\x9f\x65\x66\x1f\xac\x94\x04\xf3\xfc\xd8\xfd\x4d\xa9\x0e\x3d\x5a\x0f\xe2\x25\xce\x58\x10\x24\xd3\xd1\x7a\x0e\x56\xe0\x9e\x0a\xd2\xa5\x7d\x09\xe7\x91\x35\xbd\xb4\xfc\xab\x22\x8a\xd6\x63\x08\xe7\x55\xd7\x74\x3b\x55\x26\x87\x58\x84\xee\x13\xac\xa0\x45\x87\x66\x5c\x6f\xa8\x79\x5a\xdf\xf0\x71\xd3\x99\x26\xba\x9d\x39\xee\xd1\x0e\x3f\xf3\x8c\xb1\x4c\x13\x72\xd6\x64\xe2\xbc\x5a\x77\xe8\x56\xe5\x4b\x7d\x59\x27\xdb\x87\xfb\x8d\xdc\x81\x7f\x4f\xfc\xb5\xfa\x3f\x9a\x37\xf2\x81\x67\x41\xfa\x31\xc4\x37\xd8\xf6\x1e\x00\x00\xff\xff\xeb\x08\xbf\x89\xe9\x01\x00\x00")

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
