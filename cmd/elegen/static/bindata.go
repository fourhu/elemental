// Code generated by go-bindata.
// sources:
// templates/README.md
// templates/identities_registry.gotpl
// templates/model.gotpl
// templates/relationships_registry.gotpl
// DO NOT EDIT!

package static

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesReadmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xcb\x11\x02\x21\x10\x45\xd1\x7d\x47\xf1\x2c\x53\x22\x81\xc6\xbe\x0a\xc5\x47\x6a\x60\x33\xd9\xcf\x79\x2b\x31\x56\xf7\xc3\x36\x4b\x85\x8d\xbe\xb5\xa3\x09\xa1\xf3\x57\x46\x8c\x4c\x04\xa1\x3a\x75\x0a\xca\x75\xfa\x75\xbf\xcc\x24\x69\x78\x43\xcb\x3f\xcd\x7f\xd8\x13\x00\x00\xff\xff\xaa\x97\xff\x85\x4d\x00\x00\x00")

func templatesReadmeMdBytes() ([]byte, error) {
	return bindataRead(
		_templatesReadmeMd,
		"templates/README.md",
	)
}

func templatesReadmeMd() (*asset, error) {
	bytes, err := templatesReadmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1515709395, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\xcf\x6b\xe3\x3a\x10\xbe\xfb\xaf\x18\x42\x79\x38\xd0\x38\x97\xc7\x3b\x14\x7a\x28\xe5\x15\x02\xdb\x52\x5a\xd8\x4b\xe9\x41\x75\xc6\xee\xb0\xb2\x14\xa4\x49\xbb\x45\xf8\x7f\x5f\xfc\x43\xb2\x93\x38\x5e\xef\xa6\xbb\x39\xd9\xf2\xcc\x37\x9f\xe6\xfb\x46\xca\x46\xa4\xdf\x44\x8e\xe0\x1c\x24\x8f\xc8\xc9\xb5\x56\x19\xe5\x5b\x23\x98\xb4\x4a\xee\x44\x81\x50\x96\x51\x44\xc5\x46\x1b\x86\x59\x4e\xfc\xba\x7d\x49\x52\x5d\x2c\xc5\x46\x1b\x64\xbd\x20\x95\x2e\x51\x62\x81\x8a\x85\x9c\x45\x51\xb6\x55\x29\x90\x22\x8e\xe7\xe0\x22\x00\xa8\xb0\x8d\x50\x39\x36\x15\x1e\x37\x98\x52\x46\x69\x5d\xc1\x56\xe8\x55\x4c\x40\x48\x1e\x30\x27\xcb\x68\x56\x6b\x54\x4c\xfc\x11\x3b\x97\xfc\x5f\x3f\x55\x6c\xca\xd2\xaf\xcf\x9d\x03\x54\xeb\x0a\xa0\x8c\xa2\xe5\x12\x6e\xf5\x1a\xe5\x57\x34\x96\xb4\x02\x83\xbc\x35\xca\x02\xbf\x22\xa4\x5b\x63\x50\x31\xbc\xb5\xdf\x74\x56\x2f\x17\x55\x7c\xd2\xf0\xed\xe7\xc6\x73\xc8\xa4\x16\xfc\xdf\xbf\xe0\x5a\x9c\xd0\x9e\xab\xfb\xd5\x4a\x65\x3a\xf1\x65\xca\x12\x9a\xe2\x0d\xab\x8c\xc4\x8b\xc4\x1b\x1d\xc8\x07\x1e\x02\x14\xbe\x03\x29\xcb\x42\xa5\xe8\x29\xf4\xb3\x20\xd3\xa6\x5e\xcc\xe9\x0d\x15\x90\x47\x50\xa2\xc0\x96\xe5\x91\x22\x71\x88\xb5\x6c\x48\xe5\xf3\x5e\x37\xfb\x29\x2e\x8a\x00\xec\x3b\x71\xfa\xda\xc1\x4f\x97\x08\x20\x15\xb6\x71\x4a\xa7\x07\x74\x82\xd4\x6e\xb9\x68\x43\xc1\x77\xee\x0e\xdf\x0f\x32\xe2\x79\x5b\x74\xe1\x15\xac\x5e\xd7\x98\x89\xad\x64\x8f\xd0\xe6\x2b\x92\x11\x80\xd7\x78\xaf\x03\xd7\x82\x31\xd7\xe6\x84\x36\xa7\x1e\xe1\x78\x9b\x7d\x91\x38\xc4\x4e\x6f\x73\x48\xf9\xc4\x36\x7b\x3e\x7f\xb8\xd5\xd7\x5a\x31\x2a\xfe\x75\x63\x8b\xa1\xd4\x09\xee\x1e\x2f\x38\x66\xf2\x81\xcc\x93\xbd\xee\x1c\x50\x06\x4a\x33\x24\x2b\xfb\xa0\x35\xc3\xe2\xa4\x39\xf8\x67\x37\xfe\x5e\x6e\x8d\x90\x50\x96\x5f\xc8\xb2\xeb\x8a\xee\xc8\xf4\x69\xaa\x4d\x98\x93\x29\xaa\x0d\x0d\xcb\x78\xc1\xb1\x99\xf9\x89\x6a\xbf\x35\x3a\x27\xa8\x76\x74\xac\xfe\x92\x72\x57\x52\xb6\x5c\x08\x6d\x27\x94\x94\x80\xdf\xc9\x32\xa9\xdc\x1b\x99\xd0\xb6\xcd\xdf\xc9\x89\xe7\xf0\xf4\xbc\x7f\x24\xd5\xae\x8f\x42\xc9\xa1\x00\x17\x18\x4f\x6b\xf0\x70\xff\xce\x0f\xf7\x5d\xef\xec\x4d\x18\x10\x92\x84\x45\x7b\x2b\x36\x70\x09\x85\xd8\x3c\x35\x66\x18\x26\x3b\x91\x4a\x15\x74\x86\x1d\x93\x8b\xcb\x1d\x62\x8b\x10\xd6\x42\x9d\xd1\x39\x9c\xd5\x44\xea\xd0\xab\x86\x92\x47\x9b\x39\xe7\xbf\x96\xe5\xec\xa2\x4a\xeb\x83\x8f\x6e\xb3\xf7\xd2\xbf\xa4\xf8\xe3\xc6\xe8\xa2\xae\xb3\xf3\x6f\x24\x6c\x55\x58\xab\x53\x12\x8c\x6b\x60\xdd\x9b\xb2\x9a\xc6\xce\x55\xd4\x83\x8a\x1b\x92\xc7\x6e\xa0\x3d\xb9\xbb\xc6\x3f\xd5\x8f\xcf\xc1\x6b\xf5\xfa\xe0\x79\x2e\x65\xcd\xa5\x4d\x3d\x72\x6e\x07\xff\xed\xe3\x74\xc7\xf4\x21\xb5\xca\x9f\x0d\x71\x18\x3d\x9e\x27\xa8\x3f\x3e\xcb\x7b\x73\xe6\xab\xba\xa9\x6e\x38\x74\xc4\x79\x58\xdf\x9b\xec\xea\x37\x30\xf2\x65\x4f\x84\x6a\xcc\xcb\xe8\x47\x00\x00\x00\xff\xff\x3a\xcb\x38\xd5\x69\x0b\x00\x00")

func templatesIdentities_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesIdentities_registryGotpl,
		"templates/identities_registry.gotpl",
	)
}

func templatesIdentities_registryGotpl() (*asset, error) {
	bytes, err := templatesIdentities_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 2921, mode: os.FileMode(420), modTime: time.Unix(1515695604, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5a\x5b\x6f\xdb\x38\xf6\x7f\xf7\xa7\xe0\x5f\xe8\xcc\xd8\x7f\x38\x4a\x9f\x3d\x9b\x01\x8a\x34\x1d\x04\xdb\x4e\x8b\xba\xdb\x7d\x28\x8a\x09\x23\x1d\x3b\x6c\x25\x52\xa5\xa8\x34\x59\x43\xdf\x7d\xc1\xab\x48\x5d\x2c\x39\x93\x4e\x8b\x85\x5f\x82\xf8\x90\xfc\x9d\x2b\x0f\xcf\x21\x55\xe0\xe4\x33\xde\x02\xda\xed\x50\xbc\x06\x11\x9f\x33\xba\x21\xdb\x8a\x63\x41\x18\x8d\xff\xc0\x39\xa0\xba\x9e\xcd\x48\x5e\x30\x2e\xd0\x7c\x86\x10\x42\xd1\x26\x17\x91\xfe\xaf\xbc\xa7\x49\x34\xd3\xff\x6f\x89\xb8\xa9\xae\xe3\x84\xe5\xa7\xb8\x60\x1c\x04\x3b\x21\x34\x39\x85\x0c\x72\xa0\x02\x67\x7a\xc9\x6e\x87\x38\xa6\x5b\x40\xf1\xba\x80\x24\x7e\x77\x5f\xc0\x1b\xce\x6e\x49\x0a\xbc\x44\x27\x75\xad\xb1\xa4\x38\xa8\xae\xdd\x12\xa0\xa9\x1a\x5c\xcc\x66\x33\x87\xf0\x84\x2c\xd1\x13\xa0\x55\x8e\x56\x67\x28\xbe\xa0\x55\x5e\x4a\x61\x4f\x4f\xe5\x0a\x35\xa0\xe0\x51\x5d\x23\x0e\x05\x87\x12\xa8\x28\x91\xb8\x01\x54\xb0\xb2\x24\xd7\x19\xa0\x5b\x9c\x55\x50\xa2\x0d\xe3\x08\x0b\xc1\xc9\x75\x25\x40\x71\xd7\xcb\x9f\x59\x9a\x31\x44\x14\xcf\x84\x44\xec\xe0\x97\x82\x13\xba\x9d\xcd\x12\x46\x4b\x6b\xa6\xdd\xee\xc4\x0a\x4a\x71\x0e\x4b\xf4\x44\x71\x93\xc2\xea\xc5\xef\x35\x73\xa3\xb3\x11\x9b\x6a\x4e\x6d\x89\xf5\x52\x39\x41\xff\x57\xd7\xb1\xb5\x8d\x5b\xd2\x91\xea\x4c\xab\x62\x57\x04\xd6\x54\xc6\x6c\xfe\x37\x56\xd3\x4e\xb9\xa0\x82\x88\x7b\xa3\xf3\x65\x0a\xea\x67\x5b\x22\x47\x67\x1b\xf5\x9b\x5d\x7f\x82\x44\xc4\xb3\x5b\xcc\xc7\x90\xce\x90\x0b\x8a\xd8\x11\x77\x4a\x3a\x39\x75\x85\x5c\x0c\x28\x90\xb7\x50\x0a\x49\xaf\xeb\x68\xa9\x26\x9d\x63\x01\x5b\xc6\xef\x57\xe1\x24\x56\xf1\xc4\x79\x6a\x39\xab\x75\xac\x90\x0d\xa2\x4c\x98\x59\x97\xe5\x5b\xc6\x44\x13\x25\x6d\x29\xdf\x64\x15\xc7\x19\xaa\xeb\x97\xa4\x14\xbe\xc6\x18\x65\x92\xc2\x36\x7b\x57\xb9\xe8\xd8\x8f\xfb\xe1\xe3\xff\xf7\xcc\x31\x5e\x38\x67\x54\x00\x15\x9e\xd9\x45\xc5\xa9\xb6\x39\xe9\xb5\x79\x89\x08\x55\x3f\xa5\x88\xf1\x6c\x53\xd1\x04\xcd\xd9\xa8\x18\x8b\x36\xab\xf9\xa2\xdf\x2f\xca\xe6\x5a\x8c\x3e\xd0\xc6\xb1\x33\xab\x41\xd1\x88\x8d\x51\xc1\x08\x15\xc0\x91\x60\x08\xa3\x44\x8e\x49\x59\xc7\xa4\x3b\x4c\x8f\x22\x14\x3e\x50\x6c\x43\xb0\xdc\xea\x46\x0f\x25\xc0\xea\x0c\xe1\xa2\x00\x9a\xce\xc7\xc0\x77\xf5\x12\xb1\x38\x8e\x17\xbe\x11\x7e\x96\x20\x46\xd9\x67\x0a\xc7\xc0\x95\x81\x57\x04\x53\x3f\x31\xa2\xf0\x55\xf3\x35\x6e\x7b\x4c\xdd\x35\xff\xb9\xe5\x19\xc7\x71\xdb\x87\x5a\xff\x89\xe6\x61\x95\x78\xb0\x75\x64\x16\xfd\x73\x29\xd5\x97\x10\x3a\xf5\x59\xb9\xf4\xee\xb6\x1c\x1c\x03\x56\x09\xb5\x20\x9e\xf7\xef\x88\x85\x46\xae\x83\x18\x64\x95\x30\xc6\x57\xbb\x29\x61\xf4\x16\xb8\xf0\x6d\xaf\x62\x8d\x76\xa2\x59\xab\x5a\x1e\x6a\x62\xf9\xb7\x67\x6f\x78\x68\x2d\xeb\xed\x99\xb9\xab\x7d\x53\x11\x01\xb9\x67\xab\xbd\x56\x92\x73\xf7\x5b\xe3\x39\x6c\x70\x95\x89\xd7\x3c\x05\x1e\xa4\x8d\x54\x0f\x20\x26\x47\x08\xdd\xa2\x0d\x81\x2c\x2d\x6d\x38\x26\x3a\x1c\x0e\x31\x89\xcf\x6a\xbe\x40\x1f\x3e\xea\x03\xb0\x95\x2c\x2c\xb9\x51\xab\x75\xfa\xbf\x36\x02\xb9\x73\xb6\x29\x01\xdc\x11\xd0\x24\x74\x0f\xc4\x9c\x5a\xda\x1a\x5a\xfb\xf7\xc0\x4b\xc2\x68\xa0\xf8\xad\xa1\x3d\x5c\x51\x83\x3a\x5f\x20\x42\x8d\x9b\xfd\x4c\x08\x22\x7e\xf6\xe6\xf2\x92\x6e\x58\x6c\xf9\xd7\x4a\x20\xef\x64\x1d\x3e\x5a\xdb\x47\x6a\xce\x52\xc8\xa4\xb0\x18\x75\x0e\xbf\x7d\xa7\x8b\xa9\x40\xaa\x44\x4a\xb8\xdb\xf9\x16\x6e\x59\x76\xb7\x43\x39\xfe\x0c\x92\xaa\x0a\xac\x99\x2c\x52\xac\xa0\xd2\x9a\xaf\xa4\x08\x56\x15\xa9\xf2\xd5\xa7\x92\xd1\x55\x74\x12\xa1\x6b\xf5\xcf\x9f\x4a\x48\x63\xd8\xe8\x4a\xaf\x92\x65\x60\xfc\xaa\x12\x70\x67\x9c\xf1\x07\x7c\x1d\x54\xd9\x9e\x0b\x32\x2b\xf6\xed\x7a\x29\x8b\x72\xd0\x20\xc8\x7c\x31\xb4\xb0\x15\x80\x3f\xf7\xcf\x6a\xe2\xd1\x57\x77\xb5\xc7\xa5\xcb\x6e\x04\xdb\xa2\x4e\x95\x7c\xd2\xf6\x8c\xab\x3a\x34\x34\xfb\x25\x25\x82\xe0\x8c\xfc\xc7\x2f\x6f\x5b\x85\x9b\x62\x1c\xe0\xb4\x18\xf6\x44\x7b\x6f\x79\x30\x54\x92\xd9\x70\x1f\x30\xda\x02\xfd\xc5\x0a\x00\x75\x4a\x00\x9b\xf7\x5a\x89\x48\x57\xa2\x81\x74\xbf\x94\xa8\xa2\xe4\x4b\x65\x8b\x1b\xb9\x66\xb2\xc4\x72\xf2\x7c\x81\xc2\xe4\xa3\x6b\x3e\x53\xef\x35\x72\x58\xe3\xdb\xb4\x19\x3b\xe8\x66\x92\x3c\x19\xe5\x59\x02\xa9\x0d\x44\x5b\xc8\x43\x56\x42\x1b\x22\x8a\x9a\x61\xed\x21\xad\xfc\x1a\x84\xc7\xb7\x04\xf1\xd8\xca\x07\x0c\xe6\x24\x35\x06\x58\x4c\xb4\xc0\x34\xd5\xd1\x19\x22\x69\xbf\x82\x7d\xa9\xf6\x06\xf3\x34\x61\x29\xa4\xed\xa4\xab\xf2\xc5\x04\xa5\x7a\x32\xed\xc4\x5c\xbb\xef\xe8\xb3\x55\xfb\xc0\x11\x38\x22\x97\x12\xec\x47\x3b\xeb\x7c\xef\x3e\x87\x32\xe1\xa4\x10\xc6\x18\xd2\x12\x2c\x09\xcf\x7e\x96\x54\x6a\x43\xab\x39\xb2\xee\x68\xc2\x6f\xdc\x29\xcf\x59\xd2\xb3\xbf\x4e\xa4\x00\xf0\xa5\x47\x86\xe8\x03\x65\x29\x4b\x3e\x46\xed\xbd\xa2\xc8\x6b\xdd\x25\x07\x7b\x2a\x9c\x76\xe5\x04\x09\x35\xbb\xea\x06\xa2\x77\xc2\x2a\x35\xf6\xef\x18\xc5\xb9\xa3\x8a\xe1\xba\xc9\x45\xbc\x2e\x38\xa1\x62\x33\x8f\xfe\xf1\x53\xb9\xfa\xa9\xfc\x2d\x92\x95\x6d\x93\x17\x95\x6b\x1a\x92\x4e\x3c\x0b\xe3\x8e\x81\x03\x57\x9f\xb7\xd2\x57\xbf\x83\x90\x1d\x90\xf6\xd0\xef\x20\xa4\x98\x9d\xfd\xe6\x7b\xad\x77\x82\xd9\x50\x1c\x12\x20\xb7\xed\x44\xf1\xa4\x57\xef\x01\x5e\xf3\x45\xc8\xc1\x5e\x18\x84\x66\xd1\x79\xa2\x93\x12\x83\xea\xc6\x06\xa3\xaf\xe0\x7a\x40\x41\x97\x09\xb7\xe4\x16\xe8\xa3\xe9\x38\xc0\x6e\xee\x6d\xa8\x5e\x6d\x5d\xba\xec\xd7\x13\x9d\x21\x0f\x21\x08\xb8\xf0\xda\xe4\x3d\xce\x48\x8a\x85\x4a\xf0\x24\x05\xad\x62\x52\x71\x0e\x54\x20\x42\x37\x8c\xe7\x7a\xf3\x95\x82\x71\x48\x65\x7a\xd3\x8d\xa1\x3e\xee\x2b\x0e\x53\xb2\xa3\x61\x22\x4f\x68\xce\x19\xb7\xb2\xab\x1f\x65\xd8\x74\x5c\x28\xda\xce\x6e\xac\x2f\x15\xe1\x90\x5e\xec\x9b\xd8\x77\x33\x17\x06\x72\x73\xac\xa8\xcb\x94\x77\x1c\xd3\x92\x48\x05\x83\xb1\xf8\xe2\xae\x60\x25\x34\xe5\xa4\x21\xbf\x35\x42\x84\xb3\x65\x06\x51\xbe\x88\xf4\xae\x8c\xec\xb0\x1c\xe3\x3c\x94\xd5\x1a\xc0\x42\x99\x1d\x1d\xa6\xcd\x01\x57\x2e\x7e\x55\x78\xff\x77\x86\x28\xc9\xbc\x26\xab\x65\x1b\xd7\x6f\x85\xf4\xa5\x5c\x6c\x3b\xaf\x30\x27\xf7\x2a\x23\x48\x0e\x07\xa9\xf2\x8e\xe4\xf0\x23\x2a\x02\x77\x02\x38\xc5\xd9\x41\xca\x5c\x98\x45\xdf\x59\xa1\x41\xf5\xe2\x67\x59\xc6\xbe\x42\x7a\x7e\xc3\x48\xd2\xc4\xf6\x3e\xd5\x74\xa8\x5d\x52\x75\x0f\xd0\x52\x4b\x47\xee\x7c\x40\xbb\x65\x53\x16\xc8\x75\x9f\x18\xa1\x1d\x01\xae\xa2\x25\x8a\xae\x24\x5a\xbd\x54\x19\xe7\x59\x25\xd8\x16\x28\x70\x2c\xd4\x8e\x19\xb2\x11\xb4\x6c\x03\x07\x38\xb9\x11\x02\xf3\x49\x36\x78\x83\x65\x82\xa7\xd3\xbc\xba\xd4\xe7\x78\x8b\xc7\x95\x56\xcf\xcb\x05\xdf\x42\xb3\xad\x40\xf1\x2b\x7c\xf7\x12\xe8\x56\xdc\xa0\xa7\x53\x74\x7b\x85\xef\x48\x5e\xe5\x7a\xc9\x54\x0d\x25\xb5\xe1\x23\x29\x1b\x9c\x95\xf0\xcd\x54\x22\xf4\x20\x95\x08\x7d\xa0\x4a\x8e\xcf\xb7\x57\x09\xdf\xa9\x17\x10\xf4\x34\x7e\x3a\x74\x30\x6c\x32\x86\xc5\xa4\xfc\x63\x9c\xf8\x42\x2e\x38\xd0\x87\xef\xcd\xfb\xc8\xe3\xe9\x6b\x0a\xdb\xa9\x42\x5f\xd2\xc9\x22\x13\x2a\xe6\x2d\xb1\x17\x8f\xed\xa7\xb1\x40\x7c\x4c\xaf\xe9\x38\x3d\xdc\x6b\x56\x8a\x6f\xe0\xb5\x89\x32\x3f\xc4\x69\x8d\xd4\x7f\x9f\xd3\x7e\xd4\xe2\xeb\xaf\xa4\x8f\xef\x5a\x6c\x3d\x8e\xe0\xdf\xb1\xb8\x7a\xa4\xd8\xea\x1d\xec\x99\x4a\x36\x28\x03\xda\x2a\xdd\x16\xe8\x37\xf4\xd4\xc9\x64\x1a\xce\x70\x8a\xff\xd6\x61\x30\xa0\x67\xad\x5b\x0d\x9d\x55\xf6\xee\x81\x64\xf6\x66\xae\x80\x84\x6c\x48\xa2\xfa\xb1\x17\x8c\xbb\x1e\x27\x68\xbe\x1d\x35\x98\xee\xee\x4e\x74\xf7\xda\x7c\x2e\xa0\x6e\x71\x3f\xc3\xbd\xed\xe2\x86\x2f\x21\x86\xb8\xcf\x15\x84\xbd\xc4\x6b\xdc\x3f\x20\x88\x6a\xfd\xc8\x06\xdd\x2e\x11\xfb\x2c\x03\xa6\x9f\x61\xd3\xc0\xbd\xc2\xc5\x07\xc9\xe2\xe3\xaf\x72\xc1\xce\xb7\xcd\xed\xcc\x58\xeb\xf4\x14\xfd\x1b\x50\xc2\xaa\x2c\x55\x2d\xde\x86\xd0\x14\x11\xb1\x44\x25\x43\x19\x88\x5f\x4a\x94\xdc\x40\xf2\x19\x31\xf3\xd4\xcc\xbe\x02\x47\x09\x2e\x01\x11\x9a\xc2\x1d\xa4\xa8\x2c\x20\x41\x39\x2e\x66\x23\x57\xc5\x2f\xe5\xd2\x73\x5c\x42\x8f\x80\xf6\x51\xb5\x57\xf1\x32\xf0\xd2\xa6\xca\x32\xcf\x0b\x65\x38\x33\xc7\xc5\xa8\x3f\x06\xb8\xcc\x17\x72\xf5\x07\xed\x8e\x8f\xd3\xbc\xb1\x57\xe1\x40\xcf\xd9\xbe\x6f\x2d\x82\x99\x9d\xd7\x21\x5c\xa8\xb7\x21\xa7\xb2\x0c\xc8\x7e\x9c\x7d\x5f\x61\x84\x3c\xce\x0e\xd1\x55\xa9\xda\x7c\xd6\x62\xef\x35\xc9\x96\x50\x9c\xad\x99\x4c\x3d\xdd\x5b\x83\xa8\x2f\x35\x45\xab\xd1\x30\xb7\xfb\xdb\x5c\x37\xf6\x76\x2c\x08\xf9\xd4\x55\xb7\xed\x90\x5d\x87\x87\xe3\x25\x29\x6f\xa9\xea\xc3\x56\x03\xfd\xe1\x49\x5d\x1f\xd4\xc3\x35\x85\x84\x5b\x56\xbb\x84\xb8\x6c\xeb\xd4\x6a\xf5\xac\x64\x3e\x79\xd5\xdb\x14\x0e\x6a\xe5\xa0\x6f\x31\xc9\xf0\x35\xc9\x88\xb8\xf7\x90\x3d\xaa\xf9\x7a\xa6\x35\x31\x1a\x45\x3e\xbf\xc1\x94\x42\xd6\x0c\x18\x82\xc1\x6b\x86\x27\x40\x71\x50\xbe\x7e\x4d\x33\x4f\x48\x9f\xaa\xb5\x6f\xcd\x1b\xc5\x0d\x9e\x03\xdc\xa8\x4f\xd5\xb8\xad\x79\x53\x71\x6d\x11\xe7\x46\x87\x8e\xf6\x93\x36\x6f\xb5\xd2\x85\x5a\x0b\x4e\x4a\xe4\xbd\x17\x6a\x8a\x09\x26\x4a\xb2\xde\x28\x52\xc3\x21\x77\x5a\xe5\x0f\xe0\xdc\x77\xb7\xda\xb7\x73\x77\x3b\x94\xe0\x82\x08\x25\x23\x9a\xeb\x73\x20\xc0\x5b\x3c\x50\x70\x5b\x87\x1e\x22\xfa\x49\x5d\x47\x9e\x23\xad\x3e\xd1\x44\xfe\x07\xb2\xea\xe1\x14\x30\x42\xfd\x9c\xda\x91\x34\x10\x5a\x05\x87\x24\xcc\x03\x0d\xcd\x86\xab\x37\x67\x42\xb0\x06\x6f\x43\x16\xd2\x11\x4d\xba\x6c\x3d\xb4\x8c\xc2\x7a\x57\xbc\x7a\xc0\x10\xb4\x88\xcd\xe8\x28\xd0\x0b\x92\x09\xe0\xea\xe3\x23\x37\xd6\xd0\x34\x5c\x30\x67\x1c\x91\x71\x20\x5b\xfa\x4f\xf0\x92\x49\x43\x33\x88\xfe\x9c\x29\x88\x39\x16\x01\x5a\x8e\x85\xc9\x74\x6e\x70\x3c\xd1\x35\x0f\x40\x9a\xae\x7f\x6b\x89\xdc\xd8\x28\x8a\xf7\x96\xeb\xc6\x1a\x9a\x46\x0b\xe6\x8c\x23\xca\xda\xcd\x03\x93\x3f\x0d\x8e\x19\x19\x85\xf0\x2f\xc1\xcc\x90\x23\xad\xba\xd7\x64\x13\xe0\x5a\xe9\xd5\x52\x56\x9d\xfb\x9a\x51\x2c\xef\x32\xcb\x82\x59\xd2\xaa\x7b\xdd\x35\x01\xae\x2d\x9a\xa1\xac\x3a\x97\x12\x63\x58\xfe\xf7\x05\xf6\xab\xd7\xde\x27\xe0\x41\x04\x75\x68\x85\xbb\xc7\x91\xb4\x3c\xfe\x8c\x51\xb8\x37\x9c\xe4\x98\xdf\x07\x7b\xa7\xa1\x69\xc0\x60\xce\x28\xe2\x5b\xc0\x69\x78\xb0\x5b\xca\xca\x5c\x04\xbb\xf1\x09\x58\xe1\x1d\x86\xc4\xd2\x94\x55\xfb\x52\x79\x14\x6b\xdd\xda\x8b\x6b\x6f\x2f\xae\x27\xef\xc5\xb5\x7e\xcc\x6b\x50\xd4\x6f\x83\x62\xc7\xc6\x51\xaa\x6b\xf3\x1a\x69\x61\x34\xc1\x7e\xd9\xec\x86\xc7\xe3\xa1\xf3\x24\x87\x90\x23\x69\xb1\xfc\x19\xe3\x70\x81\x58\x9e\x4c\x93\x05\xfa\x97\xfe\xb2\xc5\xd1\xf5\x6f\x29\x8a\x19\x1a\x94\x42\xd1\x3b\x5f\x9d\x1c\xd2\x4d\x7e\xa3\xde\x69\x80\xd9\xdf\xd3\x44\x99\x52\xeb\xd8\x41\xa1\x63\x07\x75\xec\xa0\x8e\x1d\xd4\xb1\x83\x6a\x87\xd6\xb1\x83\x3a\x76\x50\xfb\x51\x8e\x1d\xd4\xb1\x83\x3a\x76\x50\x16\xeb\xd8\x41\x05\x50\xff\xd3\x1d\xd4\x7f\x03\x00\x00\xff\xff\x8a\x05\xf9\x2c\xe9\x3d\x00\x00")

func templatesModelGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelGotpl,
		"templates/model.gotpl",
	)
}

func templatesModelGotpl() (*asset, error) {
	bytes, err := templatesModelGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.gotpl", size: 15849, mode: os.FileMode(420), modTime: time.Unix(1515715849, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x94\x41\x8b\xdb\x30\x14\x84\xef\xfa\x15\x83\x59\x4a\x02\x59\xbb\xe7\x40\x0e\x65\x0b\xcb\x1e\xb6\x94\x2c\x3d\x85\x1c\x14\xf9\xc5\x16\x2b\x4b\xae\xfc\x9c\x12\x84\xfe\x7b\x89\x9d\xec\xc6\xad\x4b\x4a\xb3\xb4\x3d\xea\x59\xef\x9b\x19\x46\xb8\x96\xea\x59\x16\x84\x10\x90\x3e\x11\xa7\x77\xce\x6e\x75\xd1\x7a\xc9\xda\xd9\xf4\x93\xac\x08\x31\x0a\xa1\xab\xda\x79\x46\x52\x68\x2e\xdb\x4d\xaa\x5c\x95\xc9\xda\x79\x62\x77\xab\xad\xca\xc8\x50\x45\x96\xa5\x49\x84\x50\xce\x36\x0c\xeb\x72\xa7\x9e\xd8\x6b\x5b\x60\x81\x64\xd5\x9d\xd7\x09\xb2\x0c\xd6\x19\x6d\x79\x8e\x9d\xf4\xaa\x24\xf5\x3c\xcb\x49\xe6\xca\xe5\x24\x84\xd8\x49\x0f\x4f\xa6\x53\x6f\x4a\x5d\x37\x4b\x2a\x74\xc3\x7e\x8f\x17\x89\x74\x39\xf6\x5d\x88\x2c\xc3\xe0\x0b\x3c\x71\xeb\x6d\x03\x2e\x09\x95\xcb\xc9\x0c\xc9\xa9\xd8\xb6\x56\x0d\x77\x26\xd3\x4b\x3a\x08\x42\xe0\x88\x1e\x77\x2a\xa2\xe8\xc9\xda\x6a\x9e\x4c\x4f\x0b\x63\xac\xc5\x25\xb5\x10\x0f\xcb\x21\xc0\x4b\x5b\x10\x6e\xc8\xb2\xe6\xfd\xa1\x95\x19\x6e\x4e\x4c\xcc\x17\x7d\x77\xc3\xf8\x31\xfe\x4a\x77\x15\xc2\x19\x29\xc6\x87\xbc\x3f\xad\xb1\xc0\xbb\x71\x47\x41\x00\x07\x23\xb7\xd0\x5b\x14\x8c\x89\x21\xfb\xea\x20\xfd\x60\x8c\xfb\xd6\xdc\x79\x92\x4c\x53\xbc\xef\xb5\x81\xf3\xf1\x1c\x95\xac\x57\x4d\xf7\x22\xd6\x1b\xe7\x0c\x7a\x66\x4f\x3d\xe6\xab\xa5\x27\xcb\x87\x40\xaf\xec\x7b\x62\x24\xaa\x63\x24\x27\x30\x90\x84\xf0\x72\x3d\xc6\x04\x73\xb0\x6f\x69\x76\x86\x24\x9b\x9f\xae\xc7\x99\x18\x99\x5e\x48\xf3\xa5\xce\xc7\xd2\xf4\xe3\xeb\xd2\xb4\x1d\xe3\x0d\xd2\xe8\x2d\xe8\xeb\x19\xfd\xd1\xe5\x84\xa4\xa2\x6a\x43\x3e\x19\xfa\xfe\x2c\x59\x95\x7f\x6c\xbb\x63\xbc\xa1\xf1\x1f\x6a\xf8\xfd\x56\x3e\x92\xa1\x91\x56\xfa\xf1\x75\xad\xe4\x1d\xe3\xaf\xbe\xb1\x7b\xe2\x9f\xa2\x2c\x89\xbd\xa6\xdd\x95\x61\x0a\xe2\x2b\x93\x0c\xdd\x3c\x4a\xbb\xff\x3f\x1c\x3d\xd8\xad\xfb\xd7\x4e\x06\xd3\xd8\xff\xa2\x8f\xe7\x28\xbe\x07\x00\x00\xff\xff\x1b\x1a\x93\x85\x53\x07\x00\x00")

func templatesRelationships_registryGotplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRelationships_registryGotpl,
		"templates/relationships_registry.gotpl",
	)
}

func templatesRelationships_registryGotpl() (*asset, error) {
	bytes, err := templatesRelationships_registryGotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 1875, mode: os.FileMode(420), modTime: time.Unix(1515701556, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
var _bindata = map[string]func() (*asset, error){
	"templates/README.md": templatesReadmeMd,
	"templates/identities_registry.gotpl": templatesIdentities_registryGotpl,
	"templates/model.gotpl": templatesModelGotpl,
	"templates/relationships_registry.gotpl": templatesRelationships_registryGotpl,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"README.md": &bintree{templatesReadmeMd, map[string]*bintree{}},
		"identities_registry.gotpl": &bintree{templatesIdentities_registryGotpl, map[string]*bintree{}},
		"model.gotpl": &bintree{templatesModelGotpl, map[string]*bintree{}},
		"relationships_registry.gotpl": &bintree{templatesRelationships_registryGotpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

