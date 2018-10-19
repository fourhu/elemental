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

	info := bindataFileInfo{name: "templates/README.md", size: 77, mode: os.FileMode(420), modTime: time.Unix(1528494836, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesIdentities_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4b\x6b\xe4\x38\x17\xdd\xfb\x57\xdc\x2f\x84\x0f\x1b\x2a\xce\x66\x98\x45\x86\x2c\x42\xd3\x0d\x05\x93\x26\x24\x30\x9b\x22\x0b\xb5\xeb\xda\x11\x23\x4b\x46\x92\xd3\x5d\x18\xfd\xf7\x41\xb2\xfc\x90\xcb\xf5\x4a\xd2\xd5\xd0\xd9\xa4\xca\xd2\x7d\x9d\xa3\x7b\x74\x5d\x15\xc9\xfe\x25\x05\x42\xd3\x40\xfa\x84\x3a\xfd\x24\x78\x4e\x8b\x5a\x12\x4d\x05\x4f\xbf\x92\x12\xc1\x98\x28\xa2\x65\x25\xa4\x86\x8b\x42\xa4\xa4\x12\x12\xb5\x48\xa9\xb8\x46\x86\x25\x72\x4d\xd8\x45\x14\xbd\x12\x09\x71\x04\x00\x40\xd7\xc8\x35\xd5\x1b\x6b\xac\xee\x49\x05\xb7\x50\x92\x6a\xa5\xb4\xa4\xbc\x78\xee\x6d\xd2\xa5\xdf\x07\x8d\x33\xb3\x7f\x4d\x73\x05\x92\xf0\x02\xdb\x64\x9e\x2a\xcc\x68\x4e\x33\x97\x8c\xb2\x89\x0c\x1b\x81\xe6\xa0\x5e\x44\xcd\xd6\x8f\x58\x50\xa5\x51\x06\xbb\x21\x85\xcb\xf4\xa1\xfe\xc6\x68\x76\x2f\xd6\x18\xda\x5e\xc1\xe5\x90\x22\xdc\xdc\x42\x6a\xf7\xb0\xf4\xf3\xf0\xf0\x6a\x64\x70\xd1\x34\x7e\xc3\x23\x2a\x6d\x97\x8d\xb9\xb8\xb1\x39\x8c\xdd\x18\xd3\x15\xb4\x08\x42\x21\x5f\x4f\xa3\x8f\x1e\x99\x28\xc0\x2c\x23\x1a\x0b\x21\xe9\x6f\x08\x9c\xa8\x65\x86\x3f\x05\x3c\xc2\x28\x51\x3f\x11\xb1\xab\x33\x42\xd6\x34\x3e\xab\x4b\xba\x80\x4b\x57\xd9\xc8\xe8\xae\xad\x14\x42\x8c\xbb\x7d\x1f\x07\xec\x9e\x83\xca\xd7\xf8\x63\x06\xeb\xd5\xf3\xea\xb9\xfd\x78\x76\x8c\x2d\x02\x93\xfe\xec\xa1\xa0\x39\x30\xe4\xdd\xf2\xb2\x4d\x1e\x8c\x99\x4d\x37\x4c\xd9\x11\x90\x89\xb2\x12\x35\x5f\x8f\x38\x18\x9c\x04\x86\x9d\xbf\x66\xec\xa2\x37\x37\xc6\x65\x69\xff\x2f\x06\x60\xc1\xec\x25\xc5\x2c\x9a\x06\x90\x29\x5b\x0e\xa7\x6c\xf1\x46\xde\x92\x28\xba\xbe\x06\x97\xfc\x3f\x28\x95\x05\x53\xa2\xae\x25\x57\xa0\x5f\x10\xb2\x5a\x4a\xe4\x1a\x5e\xfd\x9a\xc8\xdd\xe3\xd2\x15\x1b\xe5\x35\xcf\x02\xdb\x38\x81\x9c\x09\xa2\xff\xfc\x03\x1a\xef\xa7\xbf\x3c\xee\x1e\x96\x4b\x9e\x8b\xb4\x0b\x63\x2b\x8c\x22\xbd\xa9\xbc\xbb\x7b\xc2\x49\x81\x12\x94\x96\x75\xa6\x1b\x13\x39\xf7\x71\x1e\xac\x26\xd0\x9d\xd8\x2f\x52\x94\x96\xcd\x98\x5b\x4a\x5b\x78\x13\x98\xed\x6a\x57\xaa\xcf\x66\x7a\xff\xac\xac\xf9\x73\x74\x4c\xb4\x4f\xad\xfc\x6e\x62\xaf\xc3\x9b\xd3\xa3\x06\x0a\xbe\xea\xfc\x1c\x17\xde\xb5\x77\xdc\x36\xf3\xd1\x81\x07\xf5\x5b\xb9\x8f\x47\x86\xe2\x9b\x98\xf0\xa1\xbe\x98\xce\x44\x4a\xba\x50\x34\x07\x0a\xb7\x90\xa7\x5b\xd4\x10\xbe\x49\xfe\x82\xff\xd1\x74\xa9\x3e\x97\x95\xde\xc4\xc9\xa8\xa1\x3a\x68\x02\x01\x99\x73\xd5\xe3\x7e\xb2\x3b\xff\x2c\x74\xe7\x71\xe4\x9b\xe4\x00\x16\x39\x25\xdf\x18\xc6\x1d\x77\xb3\x10\x4c\x9f\xb5\x36\x1d\x32\xea\x3b\xd5\xd9\x4b\xcf\xbe\xcf\xb6\x57\xf1\x3d\xba\xf7\x66\xcd\xcb\x88\x6a\xe7\xb5\xad\x8b\x64\x10\xfb\x9b\x29\x68\x5f\xf1\xfb\x0e\x93\x38\x89\x66\x64\x63\xf2\x75\x8d\x39\xa9\x99\xde\x72\xcb\x29\xf3\x6c\xec\x02\xfa\xa9\x22\x52\xe1\x5b\xe0\xde\xb6\xfc\x85\xa0\x7b\x43\x2e\x74\x7f\x0f\xa8\x47\x21\xf4\x7b\x49\x69\x8b\x7c\x0f\x35\x1f\xc6\x94\xbf\xd8\xf6\xd3\x13\x5c\x9c\x81\xfe\xf5\x83\xc1\xaa\x73\xe0\x5e\x1f\x0e\xc9\x51\xcb\xac\xed\xda\x27\xe7\x36\x50\xa5\xfd\xbd\x37\xe9\x7d\x7f\xba\x26\x52\xd0\xea\x5c\x72\x9c\x12\x1c\x28\x7e\x3e\x1d\xf5\x61\xc7\x72\xe7\xe9\x3a\xaf\x56\xfc\x7f\xce\xe0\x81\xd5\x92\x30\x30\xe6\x6f\xaa\xec\xd5\x7d\xc6\x83\xb9\x2d\x04\x47\xf3\x34\x63\xfa\xdb\xb1\xb5\x5b\x42\x7e\x21\x67\x01\xe4\x27\x75\xb7\xda\xdb\xde\xea\xe4\xfe\x7e\x44\xd6\xd2\xf7\x42\x2b\x15\x8f\xa3\x06\x2b\x2d\x4f\x72\x3a\x5d\xc9\xb9\x3d\x36\xd6\x2b\x91\x50\xfa\x79\xf6\x36\x08\x69\xe7\x5a\x3b\x70\xfb\xc5\xf1\xac\xed\xb6\x8d\x32\xb8\x1f\x99\x75\xb3\x76\xfb\x2d\x48\x74\xbc\x6d\x98\xba\xbb\xe8\x26\x72\xf1\xee\x18\xf3\xc0\x50\x54\x7d\x54\xc2\x18\xe0\x0f\xaa\xb4\x55\x6c\xda\xaf\xfb\x60\x81\x4d\x6c\xd5\xfd\xe0\xb8\x39\xb7\xe5\xfc\x6f\xd5\x7b\x5b\xe4\xb4\x9f\x10\x8c\x47\xcf\x4d\xd1\x5f\x84\xec\xeb\x1e\x43\x68\xc9\xf3\x83\x36\xe4\x42\xba\xef\x05\x7d\xc5\x61\xee\xef\x11\x9d\xfa\x39\x74\x9f\x86\xb7\xe9\x2e\x55\x3a\x02\xd6\xf3\xca\xce\xf8\x7d\xf7\x5d\x3f\x59\xec\x7f\x05\x3e\x42\xa4\xc2\x37\x01\x2b\x4d\x26\xfa\x2f\x00\x00\xff\xff\x51\x27\x8c\xfd\xd7\x14\x00\x00")

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

	info := bindataFileInfo{name: "templates/identities_registry.gotpl", size: 5335, mode: os.FileMode(420), modTime: time.Unix(1539813059, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModelGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x3c\x5b\x6f\xdc\xb6\xd2\xef\xfb\x2b\xd8\x85\x9b\x4a\xe9\x5a\xce\xf3\xa6\x2e\x90\x2f\x97\xc2\xf8\x92\xc6\x88\x83\xf4\xc1\x30\x6a\x5a\xa2\xd6\x6c\xb4\xa4\x42\x52\x76\x7c\x16\xfa\xef\x07\xbc\x49\xa4\x2e\xbb\x94\x2f\x39\x3e\x38\xce\x43\xeb\x1d\x92\xc3\xb9\x71\x38\x33\xa4\x58\xc2\xf4\x2b\x5c\x21\xb0\xd9\x80\xe4\x04\x89\xe4\x35\x25\x39\x5e\x55\x0c\x0a\x4c\x49\xf2\x27\x5c\x23\x50\xd7\xb3\x19\x5e\x97\x94\x09\x10\xcd\x00\x00\x60\x9e\xaf\xc5\x5c\xff\xc5\x6f\x48\x3a\x9f\xe9\xbf\x57\x34\x81\x25\x65\x48\xd0\x04\xd3\x03\x54\xa0\x35\x22\x02\x16\xba\xe7\x66\x03\x18\x24\x2b\x04\x92\x93\x12\xa5\xc9\xe7\x9b\x12\x1d\x33\x7a\x85\x33\xc4\x38\xd8\xaf\x6b\x8d\x42\x52\x01\xea\xba\x19\x82\x48\xa6\x1a\x87\x50\x7c\x81\x05\xce\x14\x99\x13\x10\xc5\xb3\xd9\x66\x03\xf6\x0a\x28\x10\x17\x5f\x10\xe3\x98\x12\xb0\x3c\x34\x28\xdf\x2b\xf0\x2b\x21\x18\xbe\xa8\x04\xe2\xb6\x83\xa4\xa0\x99\x7d\x0f\x2f\xc0\x1e\x22\xd5\x5a\x8e\xbb\xa8\x70\x91\xbd\x25\xd5\x9a\x6b\x14\x5d\xd4\x75\x3d\x3b\x38\x90\x04\xa8\x11\x8a\x6d\x50\xd7\x80\xa1\x92\x21\x8e\x88\xe0\x40\x5c\x22\x50\x52\xce\xf1\x45\x81\xc0\x15\x2c\x2a\xc4\x41\x4e\x19\x80\x96\x0a\xc5\x8c\x1e\xde\x50\x66\xf4\x32\x4f\x66\x42\x62\xec\xe1\xe7\x82\x61\xb2\x9a\xcd\x52\x4a\xb8\xd5\xda\x66\xb3\x6f\x39\x20\x70\x8d\x16\x60\x4f\xcd\x26\xb9\xd0\x83\xbf\xe8\xc9\x8d\x08\x0d\xd9\x44\xcf\xd4\xa5\x58\x0f\x95\x1d\xf4\x5f\x75\x9d\x58\x51\x37\x43\x7a\x54\x1d\x6a\x56\xec\x08\x4f\x39\x4a\x37\xed\xdf\x33\x49\x2d\xce\x01\xa1\xc2\xe8\xe6\x03\xcd\x50\x91\xbc\x41\x02\xa6\x97\x28\x6b\x05\xeb\xb6\xbe\x25\x02\x8b\x1b\x23\x9c\xa3\x0c\xa9\x9f\x5d\xd2\x1b\x38\xcd\xd5\x6f\x7a\xf1\x0f\x4a\x45\x32\xbb\x82\x2c\x0c\xdf\x21\x68\x6c\x3b\x69\x80\x1b\xc5\x8c\xec\xba\x04\x8d\x05\x3a\xa8\x3e\x21\x2e\x64\x6b\x5d\xcf\x17\xaa\xeb\x6b\x28\xd0\x8a\xb2\x9b\xe5\x50\x57\x5a\xb1\xb4\x51\xb2\xee\x7f\xac\x17\xea\xb2\x8f\xda\xb4\xb4\x3d\x19\xbe\x82\x42\xf6\xec\x76\xd4\x0d\x75\xbd\x98\xd5\x13\x65\xbd\xd9\x0c\xf5\x38\xe2\x9f\x28\x15\xbb\x74\x71\x5c\x54\x0c\x16\xa0\xae\xdf\x63\x2e\x5c\x6d\x40\x50\x48\x08\xcd\x03\xc6\x36\x86\x1e\x32\xc7\xe9\xd9\xf3\xd1\x9e\x92\xe1\x83\x03\xe0\x58\x87\xa8\x18\xd1\xa6\x81\x07\x4d\x83\x03\x4c\xd4\x4f\x49\x6d\x32\xcb\x2b\x92\x82\x88\x06\xd2\x12\x37\x33\x45\xf1\xb0\xdd\x28\x9d\x69\x2a\xc6\x71\xb6\xe6\x37\xd3\xf4\xbf\xa6\x65\x4b\x3b\x04\x25\xc5\x44\x20\x06\x04\x05\x10\xa4\xb2\x4d\x12\x1c\x46\xe2\x74\x96\xe4\xe4\x03\xec\xe4\x18\x5e\x14\x88\x5b\x9e\x14\x19\xcb\x43\x00\xcb\x12\x91\x2c\x0a\x43\xbe\xa9\x17\x80\x26\x49\x12\xbb\x62\x79\x26\x51\x19\xc6\x5f\x29\x6c\x06\x29\xf7\xd4\x24\xa8\xfa\x09\x01\x41\xd7\x7a\x76\xa3\xc7\x87\x92\x83\xa6\x25\xb2\xf3\x27\x49\x32\x2c\x92\x9d\xa2\xa2\x95\xb8\xa3\xa4\xe4\x96\xf1\xf7\x42\x8a\x42\x22\xd2\x7e\xde\xd2\xa5\x7d\x93\x9d\xa7\x99\x86\x56\x42\x0d\x48\xa2\x6d\xab\x25\xd6\xf8\x6b\xcf\x4e\x69\x25\x8c\x3a\xd4\x7a\x4b\x29\xb9\x42\x4c\xb8\xda\x50\x96\x48\xc6\xf8\xbe\x9d\xb8\xe5\x7f\xc7\xcd\x4e\x51\xe2\xcb\x73\x0d\xbf\xa2\x68\x4b\xf7\x05\x28\x10\x89\x68\xdc\x8a\x10\xcb\x61\x2f\x5e\x02\x0c\x7e\x33\x6d\x2f\x01\xfe\xf5\x57\x5f\x84\xa7\xf8\x0c\x1c\x02\x7a\x8a\xcf\xb6\x8a\xe6\x0d\xca\x61\x55\x88\x8f\x2c\x43\xcc\x73\x33\x99\x6e\x00\x54\xb6\x60\xb2\x02\x39\x46\x45\xc6\xad\xb5\xa6\x94\x08\x44\x6e\x21\x1f\x77\xc2\x28\x06\xa7\x67\x3a\x0c\xe8\xf8\x18\x0b\x6e\x59\xea\x04\x56\x1f\x0d\x59\x6d\x1c\xd4\x8d\x6a\x6c\x98\xd5\x6c\x74\xfe\x56\xe5\x6f\xeb\x5a\x44\x5a\x24\x9f\xe9\x49\x09\x19\x47\x9e\x38\x02\x9d\xba\x31\x32\x94\x49\xd3\xd2\x68\x42\xd7\xf5\xc1\x01\xf8\xd8\x77\xe5\xe0\x1a\x17\x05\xa0\xa4\xb8\x51\x22\x87\xa6\x69\x85\xaf\x10\x31\x2a\x49\xc0\x9f\x54\xff\x09\xd6\x08\x12\x0e\xa4\x01\x31\x64\x40\x1c\xdd\x42\x49\x56\x04\x91\x51\x7a\x92\x24\x5a\x1f\x8f\xca\xb0\x93\x0e\x99\x8d\x9f\x19\x33\x76\x6b\x1a\xae\x62\xaf\x0c\xec\xae\x76\x6d\x70\x47\x31\xc0\x44\x0c\xed\x98\x48\x24\xaf\x8e\x8f\x8e\x48\x4e\x13\x27\xf0\xd6\x41\xbb\xb1\x42\x19\xe3\xd8\x78\x67\x47\xe0\xd8\x0d\x18\xd7\xb2\x8b\x64\x02\x76\x87\xd9\xa0\x6e\x77\x7c\x62\x82\xf2\x2a\x95\xf4\x77\xd6\xdb\xf6\x75\xb6\xd9\xa8\x54\xe0\x33\x7d\xa7\x6c\x6e\x4f\x72\x0b\x72\x58\x70\xa4\x32\x1c\x8f\x31\x29\x17\x35\xb3\x45\x20\xe5\x75\xfe\x0f\xa7\x64\x39\xdf\x9f\x83\x0b\xf5\xc7\xdf\x8a\x1f\xa3\x9b\xf9\xb9\x1e\x25\xf3\xb8\xe4\x43\x25\xd0\xf7\x5e\xff\xfd\xf9\xb9\xd1\xf1\x9f\xe8\x7a\x87\xdc\x6c\x38\x22\x37\xe0\xf1\x4d\x45\xd2\xaa\x6c\x60\x07\xc2\x28\xde\x8e\xa4\x63\x0a\xcf\xb6\xf5\x6d\xed\xdc\x15\xd0\x72\x8b\xfd\x2c\x66\x8e\x33\xdb\x77\x93\x3f\xa9\x0f\xb9\x9c\x38\x65\x4e\xb2\x08\xa2\x1d\xea\x8c\x3d\xbf\x69\x82\x6e\x7e\x49\xab\x22\xfb\x8b\x61\x81\x8e\x08\x16\x18\x16\xf8\x5f\x88\x49\x35\xab\x6c\x52\x4e\xa5\xbd\x6b\xc7\x34\xf6\x92\xe3\xea\xa2\xc0\xa9\xe4\x06\x74\xd0\xee\x61\x82\x95\x9b\xb8\x1e\x40\x8b\x84\x87\xbc\x3b\x16\xe7\x66\xb8\x07\x1f\x82\xed\xbb\x0e\x3e\x0c\x64\xd6\x64\x48\x66\x37\x18\x9f\x8f\xa5\x6e\xd6\xa1\x6c\xb5\x96\x7b\x8a\xc4\x81\x17\x8a\xbb\xab\x6f\x0a\x63\x39\xee\xc4\x04\x3a\x29\xf6\xf8\xfa\x85\x83\x8a\xe0\x6f\x95\xcd\x4b\xe4\x98\x89\xbc\xca\x21\x51\x0c\xfc\x38\x40\xe7\x72\x7a\xac\x43\x8d\x35\x4e\xeb\xda\x93\x66\x82\xb6\x53\xf2\xda\x6e\xc2\x76\x1d\x37\x5a\x96\x0e\xa9\x83\x62\x3e\x9f\x75\x8c\xe0\x36\x02\x3b\x41\xc2\xa1\x92\x23\xf1\x30\x02\xf3\xa6\x89\x70\x06\xec\xae\x1c\x26\xb5\x30\x71\x81\x43\x80\xb3\xfb\x10\xca\xd0\x76\x7b\x09\x59\x96\xd2\x0c\x65\xdd\x8d\x57\x39\xfc\x60\x41\xdc\x7a\xb7\x9d\xce\xc5\x68\x84\x6c\xeb\x02\x23\x91\x72\x10\x2f\x8a\x99\xc7\x1d\x12\x4f\x94\x98\x6b\x82\xb6\x95\xa7\x0c\x97\xa2\x2d\x35\xbe\xa1\xa9\x9f\x6b\xd0\xb4\x52\x9e\x4e\xf5\x91\x91\x60\xbb\x5e\x42\x2d\xe2\x0d\x4d\x07\x9c\x88\x22\x17\x7d\x1b\xa5\x67\x7e\x4a\x68\x46\xd3\xb3\x79\xd7\x2d\x28\xf0\x89\xae\x50\x7a\xee\xc3\xef\x76\x2e\xe5\xc5\xd3\xff\x83\xe9\x57\x81\xd3\xaf\x7c\x0b\xe3\xe7\x43\x4b\x6a\x28\xf2\x0b\x10\x72\xa8\xbb\x50\xf4\xf7\xc4\x62\x68\xcf\xd7\x22\x39\x29\x19\x26\x22\x8f\xe6\xbf\xfd\xcc\x97\x3f\xf3\xdf\xe7\x32\x5d\x6f\xb7\x20\x65\x20\x2d\x48\x7b\xea\xb8\x67\x14\x81\x91\x62\x63\x1d\x3a\x9e\xf8\x03\x11\xc4\xa0\x40\x7f\x20\x21\x10\x03\x49\x2f\x5c\x38\x38\x00\x7f\x20\x21\x59\xec\xb9\xa8\x6e\x6a\xd6\xeb\x60\xbc\x0a\x43\x29\xc2\x57\x5d\x0f\xbb\xb7\x45\x66\x23\x33\x46\xb1\x3f\x8f\x2d\x17\xfb\x22\xd5\x0e\xb6\xb7\xff\x74\x82\xfc\xbe\x08\x4e\xb6\x88\xe0\x64\x44\x04\xcd\x26\xa3\x73\xc1\xfb\x95\xc2\xc8\xa4\x91\xe3\x34\x06\xe5\xd1\xec\x44\xc3\x92\x00\x87\xc0\xc1\xd0\x59\x01\xdd\x05\x00\x49\x06\xa2\xb1\x55\x10\xf7\x9b\x74\xa1\x37\x36\x62\x1b\xcc\xe2\xb9\x06\x0d\x6f\x3d\x72\xd0\xa5\xed\x8f\x32\x5b\x23\xba\xd7\x04\x7c\xc7\x82\x0d\xca\xbb\x75\x17\x37\x9d\xb6\x42\xc7\xb9\x4a\x9b\xf5\xe0\x18\x1c\x1e\x82\x17\x4e\xee\x7c\x70\x00\x08\x2d\x30\x11\x4b\xb0\xa2\xfa\xac\x8c\x37\x8d\x36\x3d\xd9\x51\xb3\xd0\x26\x0f\x9c\x7f\x6d\xda\x11\xb0\xfc\xfd\x81\x7d\xfb\x58\x82\x67\xd2\x70\x7c\xb8\xcc\x72\xba\x33\x22\x92\x39\xe8\x6a\x37\xeb\xe7\xa5\xcc\x2a\x82\x18\xa9\xdd\xaa\x64\xde\xd6\x24\x8d\xf4\x5b\x46\xf9\x35\x16\xe9\x25\xc8\xef\x8b\xf7\x14\x72\xe4\xef\xc0\x4b\xaf\x5d\xf3\x31\x20\x09\x70\x08\x9e\x45\x43\x22\x8a\xa7\x88\xc8\x28\x9b\x97\x26\x67\x3e\x86\x92\x3b\x58\x96\x85\x2e\xc9\x13\x4a\x00\xc1\x45\x1b\xbe\x42\xf0\x3c\x40\x9e\xa6\xb4\x3d\x31\xe1\x51\x93\x47\x66\x65\x6e\xb3\xf2\xd8\xdd\xd6\x55\x5e\xc8\xdb\x53\xd1\x9d\xa2\xb7\x79\xa3\x1e\x67\x80\x38\x07\x3f\xe9\x99\xdd\x4d\xef\x88\xbf\xfd\x56\xc1\x22\x72\x77\xc2\xd8\x51\x7d\x09\x09\x4e\xa3\x79\x0a\x89\x74\x41\xa5\x12\x5e\xce\xe8\x1a\x40\xa0\xb9\xb8\xc6\xe2\x12\x64\x38\xcf\x11\x43\x44\x34\x27\x36\x73\xaf\x3a\xc5\xa9\x4a\xcf\xf5\xec\x51\x80\x7c\xe3\x59\xf7\x88\xb4\xcb\x0b\x1f\xf1\xba\x3f\x1d\x2a\x75\x3a\x65\xb4\x31\xef\xfc\x9c\x0f\x19\xd7\xac\xb1\x20\xcf\xb0\xda\x1f\xbd\xd2\x95\x39\xf7\x56\x19\x10\xce\x90\xf6\xbe\x69\xc5\xb4\x3c\x48\x4e\xd9\x5a\x07\x7b\x5c\x50\x86\x32\x19\xc5\x6b\xe3\xd1\xf5\xa7\x8a\xa1\xf0\x54\xc0\x4c\x25\x13\x66\xc6\x28\xb3\x56\xa2\x7e\x28\x0b\x69\x8d\xea\xad\x82\x6d\x6c\x08\xf7\xad\xc2\x0c\x65\x6f\xb7\x75\x1c\x3a\xd4\x0f\xb0\xb5\x26\x80\xfb\xcc\x20\xe1\x58\x72\xed\xb5\x25\x6f\xbf\x97\x94\xa3\xb6\x22\x66\xc0\x9f\x0c\x4d\x7e\x6f\x19\xc0\xaa\x0d\x76\xae\x77\x83\xb9\xa3\x73\xc4\x98\x4f\xba\x95\x87\x45\x65\x82\x40\x3f\xde\x1f\xb1\x80\xf8\xa5\xc2\xd7\xb3\x97\x8e\xa8\x9a\xd3\x19\x1f\xbe\x90\x83\xe3\xd6\x58\x9c\x64\x62\x90\x19\x81\xd7\x68\x12\x2b\x9f\xf1\x1a\x3d\x46\x46\xd0\x77\x81\x18\x81\xc5\x24\x66\xde\x9a\x41\x8f\x91\x21\x4c\x04\x5a\x21\x36\x89\x9f\x23\x22\x1e\x23\x2b\x79\x41\xa1\x98\xc4\xc8\x3b\x39\xe2\x91\xb0\xb2\x8d\x33\x86\xf2\x01\xbe\x86\x09\x4d\x5a\x2f\x39\x42\x32\xea\x90\x8a\x26\x48\x9b\x32\x10\x79\x74\xbd\xc7\x5c\xcc\xe3\x0e\xf0\x03\x2c\xe7\xb1\x25\xd8\x04\x5d\xbc\xba\x70\x8e\x82\x87\x77\xa5\x96\xc6\x96\x4b\x5e\x5d\xec\x66\x29\x88\x2d\x37\x38\x1a\x65\x30\x79\x55\x14\xf4\x1a\x65\xaf\x2f\x29\x4e\x11\x0f\xb1\x26\xed\x76\x8f\x88\x3a\x0f\xee\x18\x93\xf6\xe2\xd1\x88\x4d\x2d\xda\x52\x8f\x1c\xf7\x0f\xc5\xa4\x47\xc0\xf9\x7c\x01\xe6\xe7\x12\x5b\xbd\x50\x91\xf4\xab\x4a\xd0\x95\xc9\x26\xb3\x2d\x96\x79\x17\x35\xb7\x44\x40\x16\x24\x83\x63\x28\xd3\x5a\x12\xb6\x96\x16\xaa\x98\xd2\x9d\xe3\x5c\xb3\xe7\xec\x8b\x0f\xc1\xd9\x4a\x80\xe4\x03\xfc\xfe\x1e\x91\x95\xb8\x04\x2f\x42\x78\xfb\x00\xbf\xe3\x75\xb5\xd6\x43\x42\x39\x94\xd0\x76\x1e\x09\x51\x87\x64\x0f\xc6\x12\x26\x93\x58\xc2\xe4\x96\x2c\x35\xf3\x3c\x3c\x4b\xf0\xbb\xba\x02\x08\x5e\x24\x2f\xc6\x82\xa4\x70\x97\x6f\x94\x38\xc1\xe3\x37\x3a\xfc\x62\x2e\x08\xde\x1f\xbf\xa6\xba\x18\x4a\x74\xf0\x7e\xbb\x90\xd1\x75\xd4\x21\x3b\xbe\x6f\x3d\xed\x32\xc4\xfb\xd4\x9a\xb6\xd3\xe9\x5a\xb3\x54\x3c\x80\xd6\x02\x69\xbe\x8d\xd2\x5a\xaa\x7f\x90\xd2\x9a\xb3\xe3\xc4\x4d\x9e\x9d\xb3\x65\x75\xab\x78\xf0\x6e\x73\x4f\x12\x12\xdd\xda\x72\xab\x58\x77\x8e\x8b\x5b\xfe\x35\x30\x34\xbc\x0a\xe6\x73\x7f\x24\x9e\x0a\x90\x82\xc7\xf0\x55\xc3\xaa\xa2\xac\xa9\x3a\xe8\x84\xb4\x15\x84\x7b\x31\xf9\x75\xc5\x05\x5d\xdb\x2a\x62\x8b\x21\x69\x6b\x18\x6b\x58\x96\x98\xac\xd4\xed\x66\x75\x64\xd5\x62\xfa\xa0\x9b\x12\xf3\x7f\x30\x6f\xaf\xaa\xf7\xc8\xe9\x54\x38\x2c\xd6\x61\x5d\x18\xbc\x56\x23\x74\x4c\xc6\xa6\xe4\xa5\x98\x45\x8c\x25\x91\xb8\x29\x91\x5b\x05\x51\x35\xac\x6e\xc6\xbc\x0c\x0b\xbe\x9a\x9b\x3a\x63\x88\x02\xf1\xb4\x48\xcc\x21\xdc\xe4\x71\xf5\xb0\xbd\x0c\x99\x8f\xa9\xae\xfa\x31\x7c\x0c\x7e\x77\x8a\xac\xa6\xb0\xe6\x77\x71\x8b\x3e\x06\x07\x1a\x18\xdb\x8c\x46\xbd\x51\xf6\x4c\x0a\x17\xe1\x77\x13\x64\x0b\xce\x71\xaa\xcc\xe4\x1d\x65\x4d\xd1\xc2\x2b\x88\x37\x50\xaf\x7b\x73\x04\xa7\xab\xdc\xed\x2d\x7f\x75\x6b\xfe\x2b\xba\xb1\xc5\x99\x5d\xe7\x4f\x63\x34\x44\x0a\x51\xbf\xb4\x3d\x42\x4e\x5b\xde\xbe\x5a\x00\xfa\xd5\xd8\xf2\xe8\xc4\x6d\x7d\xe6\x03\x2c\x4f\xe5\x54\x67\x2f\xe5\xb0\x9e\xa4\xaf\x5c\x21\x1f\x1c\x80\xbf\x10\x48\x69\x55\x64\x4a\xb6\x39\x26\x19\xc0\x62\x01\x38\x05\x05\x12\xbf\x70\x90\x5e\xa2\xf4\x2b\xa0\xe6\x72\x1e\xbd\x46\x4c\xdb\x2e\x26\x19\xfa\x8e\x32\xc0\x4b\x94\x82\x35\x2c\x5d\x9d\x6d\xa3\xf3\xbd\x44\xf1\x1a\x72\x34\x40\xb0\xbd\x48\x3c\x28\x10\xee\xe9\x30\xaf\x8a\xc2\xd1\x11\xf7\x7b\xae\x61\x19\xa8\xad\x91\xb9\xa2\x58\xe2\x38\xd5\xca\x3a\x0b\xd5\x55\x00\xfb\x1e\xd7\xb3\xdd\x57\xde\xbc\xfe\xbd\xfb\x6f\xb0\x54\x85\xea\x46\x0c\xd2\x84\xb7\x61\xdb\xfd\x2d\x85\x3f\xdf\xe1\x14\x29\x6c\xba\xbb\xc8\xa4\xc2\xa1\x73\xdd\xaa\x19\x21\x49\xe8\x1c\x0f\x7a\xdf\x11\x75\xb7\xce\xf9\x72\xe7\xb2\x72\x2f\x12\xec\x8f\xe6\x77\xf2\x9f\x0b\x5f\xf6\xd3\x34\x99\xa5\x79\xb8\x3a\x17\xac\xfc\xdc\x75\x39\x92\x53\xef\xd7\xf5\xa4\xbc\xb7\x0d\xbe\x9a\x61\x75\xb3\x8b\x2f\xfa\xbc\x75\x12\xe4\x96\x3a\xb7\x61\x39\x98\x4c\x6f\xe5\xce\x4e\x30\x78\xeb\x48\xfe\xf3\x1a\x96\x23\xea\x0a\x9b\x82\x21\xa5\xb9\x8f\xa4\xb8\xf1\x66\x70\xe0\x9a\x83\x4e\xcf\x20\xec\xde\x45\x17\xa7\xdd\x85\x6b\xec\x9d\x9e\x53\xb0\xdb\x50\xb6\x8b\x5d\xc1\x15\x76\x75\x1f\xd0\xeb\xed\x5d\x08\x4c\xc2\x27\x2c\x19\x4a\xbb\xda\x6e\xa1\x96\x15\xa7\x57\x20\x5e\xef\x3a\x4b\x8b\xb8\x01\x2f\x07\xae\x85\x74\xee\x82\x04\xcd\xe4\x9c\x11\xd8\x26\x03\xd2\xb4\xb7\xed\x41\xe8\xde\xe1\x42\x20\xa6\xce\x8b\x9d\xd6\x16\xaa\x91\x7a\xbd\xc2\xf0\x52\x86\xf0\x8a\xfc\x3f\xf2\x6c\xb2\x85\x1a\xbc\x6e\xaf\x20\xbc\xe6\x5e\x88\xd3\xa2\x21\x1a\x5f\xd3\x1a\x84\xcb\xb9\x0c\xe7\xb4\xb6\x50\x8d\xd3\xeb\x15\x84\xd7\x2d\xe6\x34\x8d\x0d\x70\xd9\x2f\xf8\x04\x22\xed\x2d\x13\x0b\x5b\xf6\xea\x0f\x41\x18\x9d\x02\x4d\x8b\xd2\x02\x97\xfd\x22\x4e\x20\xd2\x3e\x99\x06\xb6\xec\x25\xdc\x21\x18\xbb\x9e\xd3\x71\x98\x93\xfc\xa4\x72\x4c\x5d\x43\x6f\x80\x9a\x36\xb7\x4f\x10\xd2\x63\x86\xd7\x90\xdd\x74\xcc\xbc\x85\x6a\xb4\x5e\xaf\x20\xbc\x9f\x10\xcc\xba\x0e\xdd\xc2\x96\xa6\xf8\xd9\xf4\x08\xc4\xe8\x1f\x22\x6a\x8c\x1a\xb6\xec\x96\x53\x83\x30\x9e\xf4\x96\xe2\x89\xb3\x14\x4f\x26\x2d\xc5\x13\x7d\xe2\xeb\xe2\x52\x10\x83\xcb\xb6\x86\xe1\xaa\x2e\xcc\xc5\xa3\x16\x99\x06\xd9\x2f\x56\x9b\x0e\x61\x96\xd3\x3b\xae\x95\xff\x1a\xa0\x26\xd1\xed\x13\x86\xb4\x43\xa2\x43\xdf\x4e\xe2\xcc\x0c\xe3\xa9\xe9\xee\xb8\x79\x38\xcd\xf8\x01\x01\xf4\xc8\xc4\x8f\x3b\x92\xd6\x79\x5d\xf2\x14\x46\x3f\x85\xd1\x4f\x61\xf4\x53\x18\xfd\x14\x46\x83\xa7\x30\xfa\x29\x8c\x7e\x0a\xa3\x2d\xc6\xa7\x30\x3a\x00\xe1\x7f\x63\x18\xbd\xe9\x7f\x82\x73\x87\x2f\x10\xf4\x11\x4c\xf8\xf7\xff\xc3\x4f\xce\x84\x62\xd0\x1f\x76\x4f\x9a\xef\xf4\x6c\xd7\x95\xe3\x7b\x7b\x84\x66\x0a\x5d\xff\xd1\xa7\x68\xa6\x3d\xd8\x70\x3b\xf6\x6e\xf5\x2c\xcd\x94\x29\x1e\xe4\x71\x9a\x1f\x21\x99\x07\x7a\xa8\xe6\xf6\xb2\xbb\xdb\x73\x35\x3b\x57\xd7\x0f\x78\xb4\x66\x9a\x02\xfe\x57\x9f\xae\x99\x26\xa5\xc7\xfd\xb5\xae\xfe\xf4\xed\xb8\x80\xd8\xff\xee\x7a\xd2\xde\xe0\xbd\x62\xf3\xb0\x8b\xde\xd0\xfa\xb8\xcc\x2e\x69\xa8\xda\x6a\x80\xb7\x79\x4e\x66\x9a\x74\x6e\xff\xa8\xcc\xee\xf0\x63\xe0\x11\x99\xfe\xe7\x91\xdb\x5e\x93\x49\x82\xa2\x0e\xef\x51\x99\x81\x75\xb0\xdb\xfe\x07\x1f\x97\x11\xac\x42\xfe\xfd\xb4\x1f\xf7\xbe\x4c\x90\x60\xdd\x57\x66\x02\x64\x94\x34\x8f\xcd\xec\xee\x1b\xc5\x41\x9f\xe3\x6d\xbc\x00\x60\xf7\x80\x8d\x35\x9c\xa0\x67\x4d\x8c\xad\xf4\x3e\xf6\x0b\xf8\x8e\xed\x21\xde\x38\x09\x7d\xb5\xc4\x23\x7b\xc7\x5b\x1c\xe1\xac\xdc\xe6\x09\x13\x9c\x07\xbf\xc7\xd1\xff\xbc\xc4\x7d\xba\xa4\x76\xc5\xf5\xfc\x81\x9e\x44\x09\x7d\xe4\xe4\xbe\xe5\xfb\xa3\x5e\x3c\x79\x36\xf4\xe4\xc9\x9d\x1f\x31\x09\xe1\xf0\x4e\x3e\x7e\x68\xab\x2f\x15\xa4\x43\x99\xd1\xcb\x6d\x08\x1c\xda\xa2\x15\x64\xe8\x13\x74\xb3\x45\xef\x7c\x34\xab\x11\xf6\xa4\x53\x1d\xb3\x66\x7a\x5f\x45\xf7\x3f\x70\xad\xc4\xe8\x27\xae\xdb\xbf\x70\x75\xbe\x6a\xed\xef\xfa\x6d\xbc\xf5\xef\x00\x00\x00\xff\xff\xd2\x6b\x23\x73\x93\x5a\x00\x00")

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

	info := bindataFileInfo{name: "templates/model.gotpl", size: 23187, mode: os.FileMode(420), modTime: time.Unix(1539981584, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRelationships_registryGotpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\xdf\x6b\xe3\x38\x10\xc7\xdf\xfd\x57\x0c\x21\x1c\xed\x91\x73\xaf\xdd\xb7\x40\x1f\x8e\xf6\x58\xfa\xd0\xa5\x74\xef\xee\x25\xe4\x41\x67\x4f\x12\xdd\x2a\x92\x57\x91\xb3\x14\xe3\xff\xfd\x90\x1d\x3b\xb6\x23\xd9\x6e\xd8\xc0\x2e\xcc\x14\x0a\xb6\x66\xe6\x3b\x1a\xfd\xc8\x07\x27\x2c\xfa\xc2\xd6\x08\x59\x06\xe1\x67\x34\xe1\x83\x92\x2b\xbe\x4e\x35\x33\x5c\xc9\xf0\x13\xdb\x22\xe4\x79\x10\xf0\x6d\xa2\xb4\x81\xc9\x5a\x85\x2c\x51\x1a\x8d\x0a\xb9\xba\x41\x81\x5b\x94\x86\x89\x49\x10\x44\x4a\xee\x0c\x48\x15\xab\xe8\xb3\xd1\x5c\xae\xe1\x1e\x26\x8b\xe2\x79\x39\x81\x9b\x1b\x90\x4a\x70\x69\xe6\xb0\x67\x3a\xda\x60\xf4\x65\x16\x23\x8b\x23\x15\x63\x10\xec\x99\x06\x8d\xa2\xd0\xdc\x6d\x78\xb2\x7b\xc5\x35\xdf\x19\xfd\x06\xb5\x42\xf8\xea\x1a\x0f\x82\x55\x2a\x23\xe0\x92\x9b\xab\x6b\xc8\x82\x00\x00\x3c\x99\xee\x87\x72\x65\x79\x19\x9e\x65\xa0\x99\x5c\x23\x4c\x51\x1a\x6e\xde\x6c\x0f\x66\x30\xad\xb2\xc2\xfc\xbe\xec\x54\x2b\x89\x6d\x52\x19\xfc\x1b\xf0\x15\xec\x36\x2a\x15\x71\x99\x19\x75\xd3\x13\xa6\x36\xb8\x99\x1b\xa6\xe1\x4b\xfa\xaf\xe0\xd1\xb3\x8a\xf1\x90\xc6\x39\x85\x45\x96\xb5\xe2\xf2\xfc\x29\x2e\x1f\x97\x70\x0f\xbf\xb8\xa7\x97\x15\xf9\x1a\xa5\xad\x0d\x5c\x09\x94\xc7\x09\x85\x0f\x1a\x99\xc1\x6b\xf8\xbd\x9a\x84\xb5\xf2\xe5\x1c\xb6\x2c\x59\xec\x8a\xf5\x5c\xfe\xea\x56\x78\x92\x2b\x05\x47\x99\x4a\xea\xd0\xc3\x84\x69\x94\x66\x06\x53\x16\x55\xdd\xeb\x2a\x37\x65\xbd\x3d\x7c\x92\xd2\xdd\xc8\x52\xa0\xd5\xc4\x6e\xc2\x89\x6d\xdc\xc1\x2f\xcf\x27\x30\xf7\x35\xcb\x4e\xa5\x3d\x93\x46\x3d\x87\x19\x84\x8f\x98\x68\x8c\x98\xc1\xb8\xab\x63\xed\x38\x3a\x07\xa3\x53\x9c\x39\xd3\xa1\x74\x06\x77\x94\x5e\x98\x66\x5b\x34\xa8\x1f\x71\x65\xb7\xb8\xed\x9f\x3f\xaa\x5e\x58\x7f\x74\xf8\x8a\x5f\x53\xae\x31\xee\x2c\x76\x65\xd5\x70\x1d\xba\x9b\x37\x4e\xcd\x27\xfc\x76\x1c\x38\xb8\xda\xa1\xab\x93\x3c\xd6\x16\x4b\xfb\x57\x6e\x9d\xd3\x9e\x36\x6b\x3f\xec\x14\x3e\x83\xa9\xb8\x2d\x36\xc8\x88\x19\xb8\xca\xf7\x35\x44\xdc\x7a\xe6\xdb\xac\x75\xa8\x52\x57\xb5\x77\x45\xb5\xe2\xb6\x2f\xb7\xb3\xa0\xbb\x81\x82\xca\xa2\xc6\x94\xe4\x2a\xeb\x43\x59\xd6\xdd\x90\x02\x54\x67\x43\x7c\xb0\xe7\xe2\x74\xab\xba\x84\x50\xc6\x03\x69\xf3\xfe\x44\xe3\x92\x0c\x7b\xf5\xc8\x0c\x07\xf7\x7b\x38\x52\x5f\x9f\x73\x90\xc7\x1c\xc9\x3f\xa5\xd1\x1c\x77\x9e\x0d\xd1\x3c\x89\x8b\xe5\xf1\x2c\x3a\x32\xb9\x2f\xae\xc6\xbe\x48\x86\xce\xd6\xa1\x14\xef\xb6\x79\xa7\x7c\x65\xf6\xc7\x6a\x7e\xb8\x84\x2b\xa2\xe8\xd9\x6c\x7f\xbd\x25\x47\x77\xfb\xd0\xef\x5e\xdd\x99\x49\xf8\x88\x2b\x96\x0a\xf3\x0f\x13\xe9\xc9\x4f\x40\xd3\x9a\x7e\xb5\x50\x27\x78\x40\x70\x78\x73\xf1\x15\xe0\xd7\x7a\x06\x13\x94\xe9\x76\xd2\x57\xd4\x1f\x42\xa8\x6f\x18\x3f\x6c\x14\x8f\xb0\x58\xec\x77\x5e\x48\xff\xcd\x60\xba\x2f\x56\x38\x09\xdb\xc9\x86\xae\x81\xa2\x03\xfb\xe1\x1b\xa0\x67\xbf\x57\x36\x7c\x26\x47\xdc\xdb\xd3\x24\x7c\x4e\x85\xe1\x89\xe8\x5d\xc6\xca\xc7\xf7\x3b\x3b\x52\xd8\x51\x72\x4f\xc4\xfb\xbc\x3d\x43\x9d\x24\x1e\x2f\xc7\xeb\x46\xa0\x63\xd4\x0b\x78\x7f\x27\xf1\x29\xe0\x95\x2f\x2f\x0c\x78\xa5\x08\x01\x9e\x47\x89\x00\x8f\x00\x8f\x00\xef\x0c\x19\x02\xbc\xba\x8a\xef\x08\x78\xc4\x77\x40\x7c\x47\x7c\x37\xde\xfb\x82\x7c\xf7\xc2\x4c\xb4\x21\x3a\x23\x3a\x23\x3a\xf3\x56\x4b\x74\xe6\x32\xa2\x33\xa2\xb3\x86\x11\x9d\x11\x9d\x11\x9d\xf9\x93\x5c\xfc\xeb\xdb\x23\x0a\x3c\xf9\xfa\x56\xbe\xbc\x30\xdf\x95\x22\xc4\x77\x1e\x25\xe2\x3b\xe2\x3b\xe2\xbb\x33\x64\x88\xef\xea\x2a\x88\xef\xba\x46\x7c\xd7\x36\xe2\xbb\xf1\x11\x3f\x25\xdf\x7d\x44\xd3\xb9\x5d\x5e\xd1\x1e\xf4\xfd\xa5\xf1\xee\x23\x1a\x62\x3b\x8f\x12\xb1\x1d\xb1\x1d\xb1\xdd\x19\x32\xc4\x76\x75\x15\xc4\x76\x5d\x23\xb6\x6b\x1b\xb1\xdd\xf8\x88\x9f\x95\xed\x9e\x99\x7c\xf3\xf0\x9d\x1d\xba\x3c\xe3\x59\x15\xe2\x3c\x8f\x12\x71\x1e\x71\x1e\x71\xde\x19\x32\xc4\x79\x75\x15\xc4\x79\x5d\x23\xce\x6b\x1b\x71\xde\xf8\x88\x1f\x86\xf3\x2c\xa9\x10\x9d\x01\xd1\x19\xd1\x99\xb7\x5a\xa2\x33\x97\x11\x9d\x11\x9d\x35\x8c\xe8\x8c\xe8\x8c\xe8\xcc\x9f\xe4\x7b\x7c\x85\x2b\xff\x77\x5e\x66\x59\xf5\x94\x07\xff\x07\x00\x00\xff\xff\xbd\x9e\xc8\xb9\xa8\x46\x00\x00")

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

	info := bindataFileInfo{name: "templates/relationships_registry.gotpl", size: 18088, mode: os.FileMode(420), modTime: time.Unix(1533220355, 0)}
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

