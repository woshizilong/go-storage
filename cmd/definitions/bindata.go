// Code generated by go-bindata. DO NOT EDIT.
// sources:
// tmpl/function.tmpl (258B)
// tmpl/info.tmpl (1.32kB)
// tmpl/object.tmpl (1.394kB)
// tmpl/operation.tmpl (755B)
// tmpl/pair.tmpl (2.219kB)
// tmpl/service.tmpl (7.46kB)
// ../../definitions/infos.hcl (1.052kB)
// ../../definitions/operations.hcl (4.606kB)
// ../../definitions/pairs.hcl (2.075kB)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _cmdDefinitionsTmplFunctionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8e\x41\x4a\xc4\x40\x10\x45\xf7\x39\xc5\x67\xc8\x22\x33\x68\x1f\x40\x70\x15\x74\x25\x32\xa8\x17\x28\xda\x1e\x6d\x4c\x55\x9a\x54\x05\x07\xda\xba\xbb\xf4\x44\x04\x5d\x15\xd4\xff\xbc\xff\x6a\xbd\x46\x7f\x92\x0f\xdc\xdc\xa2\x0f\xf7\xab\xc4\xf0\x48\x9c\xf0\x05\x9b\x47\xe2\x34\xc1\xbd\x3b\xad\x12\x31\x28\x0e\xb5\xa2\xbf\xe4\x5a\x28\x6e\xa5\x23\x69\xa4\xd6\xda\xa3\xd6\x46\x72\x1f\xa2\x9d\x11\x67\xb1\x74\xb6\x30\x6e\xf7\xaa\xa5\x1b\xff\x48\x0b\xb1\x86\x97\x25\xf3\x03\xa9\x85\x67\x5b\xb2\xbc\xdd\xc9\xab\x7e\x66\x7b\x1f\x67\x66\x72\xc7\x5c\x0c\x87\x42\x79\xf9\xb7\xe9\xde\x1e\xcd\xf8\xef\xfa\xf0\xcb\x7f\x4a\xba\x4e\xa6\x3f\xdc\x8b\x58\x07\x00\x85\x24\xc7\x61\x97\xb9\x4c\x89\x93\x18\xb2\xed\xf6\x9d\x7f\x07\x00\x00\xff\xff\x5b\xb1\x71\xf7\x02\x01\x00\x00")

func cmdDefinitionsTmplFunctionTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplFunctionTmpl,
		"cmd/definitions/tmpl/function.tmpl",
	)
}

func cmdDefinitionsTmplFunctionTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplFunctionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/function.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf3, 0x17, 0x45, 0xc4, 0x8, 0xb1, 0x57, 0xf3, 0x99, 0x7b, 0x1f, 0xcd, 0xe4, 0x30, 0x7e, 0x3, 0x30, 0xba, 0x28, 0xed, 0x36, 0x1c, 0xd3, 0x4a, 0x24, 0xfc, 0x10, 0x73, 0xfe, 0xf7, 0xda, 0x75}}
	return a, nil
}

var _cmdDefinitionsTmplInfoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x93\xc1\x6b\xdb\x30\x14\xc6\xef\xfa\x2b\xbe\x85\x30\x6c\x48\x1d\x0a\x63\x87\xad\x39\x35\x65\x94\xd1\xf6\xd0\xb0\xc3\xc6\x18\xb2\xf3\x6c\x44\x2c\xd9\x48\xb2\x59\xe6\xf8\x7f\x1f\x92\x9d\xd8\x49\xd3\x8c\xed\xb2\x5b\x14\x7d\x7a\xbf\xf7\x7d\xef\x79\x3e\xc7\x6d\xb1\x26\x64\xa4\x48\x73\x4b\x6b\xc4\x5b\x64\xc5\xe1\x0c\xa1\x2c\x69\xc5\xf3\x79\x22\xd7\x1f\xb1\x7c\xc2\xe3\xd3\x0a\x77\xcb\xfb\x55\xc4\x4a\x9e\x6c\x78\x46\xb0\xdb\x92\x0c\x63\x42\x96\x85\xb6\x08\x18\x00\x4c\x52\x69\x27\x2c\x64\xac\x69\xae\xa0\xb9\xca\x08\xd3\xcd\x0c\x53\xa1\xd2\xc2\xe0\xc3\x02\xd1\xbd\xfb\xf5\xc0\x4b\xb4\xad\x7f\xd1\x34\x98\x1a\xd2\xb5\x48\xe8\x91\x4b\x72\x9a\xe9\x06\x3b\xd8\xe2\x96\x4b\xca\x9d\x8c\x39\xd2\x0b\x61\xdb\xc2\x58\x5d\x25\x16\x4d\x5f\xe8\x40\xfc\x31\xc3\xb4\xf6\x95\x3a\xee\x08\x55\x47\xab\x6d\x79\x28\x30\xfc\x33\x68\xae\x40\x6a\xed\xb1\xee\x8c\x58\x58\x54\x42\xd9\xf7\xef\xfc\x59\x42\xf2\xf2\x9b\xb1\x5a\xa8\xec\xbb\x0f\x29\xe5\x09\x35\x2d\xeb\xf5\x27\xbe\x4f\xba\x70\xb7\x22\x75\xcc\xbb\x9f\x3e\xb5\xb6\x65\x69\xa5\x12\x04\xf2\x8c\xbf\x10\x9f\xc8\x76\x2d\x2e\x85\x29\x73\xbe\xed\x2f\x82\xf0\xb8\xf3\x3e\x01\x4d\xb6\xd2\x0a\x32\x7a\x61\xd4\x75\x77\x89\xf3\xfc\x0a\xa7\x3e\xe6\x84\x3d\xe8\x0c\x01\x0b\xd4\xac\x33\x48\xb9\xf1\xcc\x7f\x74\x16\x1c\x21\x67\x88\x8b\x22\xdf\x83\x45\x0a\x19\xb9\x89\xbc\x45\x70\x7d\x73\xe3\x94\x1b\x5f\x6e\xb1\xc0\x75\xaf\xb9\x9c\xc4\x0c\x56\x57\xe4\x85\xc3\xc0\xbb\x91\x7c\x25\x5d\x7c\xe1\x79\x75\xd8\x85\x51\xa5\xae\xce\x58\x31\x43\xca\x73\x43\xc3\xd2\xf4\xae\xcf\x3e\xf4\x66\x76\xf8\x75\xe9\x7d\xb7\x74\x7f\x88\xed\xa1\x32\xf6\xef\x96\xe2\xd5\xc8\xde\x1c\x47\x56\x72\x25\x92\x20\x95\x36\x7a\x2e\xb5\x50\x36\x0d\x26\xa7\x0d\xb8\x8f\xf2\x33\xc5\x3c\x1e\x3e\x9d\xfd\xf4\x85\x81\x2a\x2c\x0c\xd9\x49\x18\x8e\xd2\xfd\xef\x2b\xd9\xdd\x39\xff\xbb\x05\x46\xfe\xf7\xbb\xda\xa5\x7e\x32\x84\xe1\xe7\xef\x00\x00\x00\xff\xff\x0b\x32\xec\xf0\x28\x05\x00\x00")

func cmdDefinitionsTmplInfoTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplInfoTmpl,
		"cmd/definitions/tmpl/info.tmpl",
	)
}

func cmdDefinitionsTmplInfoTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplInfoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/info.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x2e, 0x99, 0x33, 0x56, 0x13, 0xfe, 0x70, 0x2, 0x62, 0xa2, 0x18, 0x8d, 0x3a, 0x19, 0xf6, 0x65, 0x23, 0x88, 0x7b, 0xf2, 0xf8, 0xcf, 0x78, 0xa7, 0xbd, 0xba, 0xdd, 0x70, 0x32, 0x94, 0x7, 0x34}}
	return a, nil
}

var _cmdDefinitionsTmplObjectTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\x41\x6f\xd3\x40\x10\x85\xcf\xd9\x5f\xf1\x88\x2a\x64\xa3\x60\xab\x80\x38\x40\x73\x6a\x2b\xc4\xa1\xed\xa1\x15\x07\x10\x87\x8d\x33\x4e\x97\x7a\x77\xad\xdd\x71\x68\x70\xfd\xdf\xd1\xae\xdd\x3a\x51\x43\xd5\x9e\x38\x65\x27\xfb\xe6\xed\xcc\x37\x93\xe4\x39\x8e\xed\x92\xb0\x22\x43\x4e\x32\x2d\xb1\xd8\x60\x65\x1f\x62\x28\xc3\xe4\x8c\xac\xf2\x42\x2f\x3f\xe3\xe4\x02\xe7\x17\x57\x38\x3d\xf9\x7a\x95\x89\x5a\x16\x37\x72\x45\xe0\x4d\x4d\x5e\x08\xa5\x6b\xeb\x18\x89\x00\x80\x69\xa9\x79\xda\x9f\x58\x69\x1a\x8e\x7e\x63\x8a\xa9\x48\x85\x08\x29\xb8\x58\xfc\xa2\x82\xe1\xd9\x35\x05\xa3\x15\x6d\xfb\x16\x4e\x9a\x15\xe1\xe0\x66\x86\x83\x35\x3e\xcd\x91\xf5\xa2\x33\x62\x89\xae\x8b\x2e\x41\xa6\x4a\x1c\xac\xb3\x63\xab\x35\x19\xbe\xbf\x40\x9e\xa3\x6d\xf7\x5c\x84\x0c\x32\xcb\x31\x0c\x9a\xab\x4d\x4d\xe7\x52\x13\xba\x6e\xeb\x9b\xa0\xd9\x92\x8b\x49\x9e\xa3\xa8\x54\x30\x53\x1e\x7c\x4d\x0f\x91\xc1\xef\x6b\x55\x5c\xdf\x77\xa1\x3c\x64\xa5\xd6\x94\x89\xc9\xa0\xb8\x64\xeb\xe4\x8a\x5c\xf4\xd0\xf0\x6c\x1d\xf9\xf8\x11\xa0\x39\xaa\x22\x6e\x4d\x2c\x97\x92\x65\x16\x4b\x0b\x11\xb4\xac\x7f\x78\x76\xca\xac\x7e\x46\xfa\xa5\x2c\xa8\xed\x44\x14\x2c\x14\xa3\x51\x86\x3f\x7e\x10\x93\xa5\x35\x14\x83\xf7\xef\xc4\x44\x87\xdb\xc0\x37\x3b\x6b\x98\x6e\x45\x27\xc4\x73\x80\x8e\x30\x4f\x6f\xe3\xf8\xba\x4e\x94\x8d\x29\x90\x58\xbc\xe9\xa5\x29\xbe\x10\xf7\x84\x4e\x94\xaf\x2b\xb9\x19\xb0\x25\xe9\x2e\x38\xb4\xb1\x44\x47\xdc\x38\x03\x9b\x3d\xe2\x1c\x8a\x7a\x64\x7e\xf9\x0f\xf3\xf5\xae\x79\x3a\xb8\xef\xb1\xc5\x1c\x6b\x31\x8c\xad\xf2\xf1\xa1\x97\xf4\x90\xec\xbc\x33\xc3\xc2\xda\x6a\x7c\xcd\xb3\xe4\x24\xed\xe1\xab\x12\x36\x0b\x13\x78\x8d\xe4\xf0\xe8\x28\xe4\xdd\xc4\xd2\xe6\x73\x1c\x0e\x19\x4f\x13\x98\x81\x5d\x43\x51\x38\x0c\x74\x1c\xc0\x77\x72\xf6\x9b\xac\x1a\x7a\xd8\xe7\xd1\xaa\x37\xda\x56\xcc\x50\xca\xca\xd3\xb8\xdf\x43\xe7\x7b\x13\x63\x6f\x77\xf8\xf3\x54\x7e\xbf\xf0\xfb\xd0\x9d\x35\x9e\x5f\xb6\x02\xcf\xc4\xf6\x6a\x17\x5b\x2d\x8d\x2a\x92\x52\x73\x76\x59\x3b\x65\xb8\x4c\xa6\xb6\xff\x71\xf5\x6f\xdc\x8f\x5b\x79\x18\xcb\xf0\xc4\xd3\x34\x1d\x58\xfe\xbf\xc5\xeb\xef\x42\x7b\x77\x73\x6c\xb5\x27\x76\xfe\x48\xc6\xe3\xdf\x00\x00\x00\xff\xff\x1e\xec\x76\xd6\x72\x05\x00\x00")

func cmdDefinitionsTmplObjectTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplObjectTmpl,
		"cmd/definitions/tmpl/object.tmpl",
	)
}

func cmdDefinitionsTmplObjectTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplObjectTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/object.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x1, 0x4e, 0x7f, 0x65, 0xf9, 0xc5, 0x5c, 0x93, 0x43, 0xf2, 0x64, 0x89, 0xc2, 0xb, 0x37, 0xcd, 0xff, 0xf6, 0x9d, 0xbd, 0x44, 0x4c, 0xa8, 0xfa, 0x9a, 0xe7, 0xd2, 0x27, 0xc4, 0x40, 0x55, 0x1f}}
	return a, nil
}

var _cmdDefinitionsTmplOperationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x52\x4b\x4f\xc3\x30\x0c\xbe\xe7\x57\x58\x55\x0f\xad\x34\xba\x3b\x12\x27\x1e\x12\x17\x98\xe0\xc0\x11\x99\xcc\x2b\x16\x6b\x13\x12\x83\x36\x85\xfc\x77\x94\x36\x1d\xdb\x04\x42\x42\xdc\x12\xdb\xdf\xc3\x0f\x8b\xfa\x05\x5b\x02\xd9\x5a\xf2\x4a\x71\x67\x8d\x13\xa8\x14\x00\x40\xa1\x4d\x2f\xb4\x91\x62\xfc\xb1\x29\x54\xad\x54\x08\x27\xe0\xb0\x6f\x09\xca\xc7\x19\x94\x0c\xa7\x67\xd0\x5c\xf7\x42\x6e\x85\x9a\x3c\xc4\xa8\x42\x80\x92\x9b\x0b\xf2\xda\xb1\x15\x36\x7d\x0a\x26\x05\xc8\x19\xf6\x76\x8d\xdb\x1b\xec\x08\x62\x04\x9e\xc0\x10\x06\xa5\xa4\xc0\x2b\x30\x0e\x2a\x7a\x4d\xf5\x43\x61\xe1\xc9\xbd\xb3\x26\x57\xd4\x47\x71\x31\x0e\xdb\x14\x8f\x71\xc0\xdf\x8b\xe3\xbe\xad\x6a\xf0\xc3\x63\xc7\x49\xfd\x32\x19\xd9\xfd\xf7\xba\xa0\xee\x89\x96\xa9\x93\x92\x9b\xcb\xe1\x9d\xb9\xc6\xda\x5c\x70\x64\xfc\x07\xe2\x7d\x5e\x63\x33\xe9\xad\xf5\x13\x62\x3e\x1f\x18\x8d\x1d\xfd\x7f\x80\x98\x05\x7a\x8d\xeb\x34\x8b\x9c\x39\x9a\xdd\xe4\xe2\x5b\x4c\x95\x33\x57\xc6\x75\x28\x0b\x74\xd8\x25\xad\x1a\x0e\x13\x77\xe4\xdf\xd6\xe2\x1f\x58\x9e\x17\xe3\xd2\x0f\xc6\x57\x24\xc8\xaf\xfe\x12\xfa\x7c\xbc\x8a\xbf\x58\xdd\x83\x57\x5a\x36\x90\x0f\xac\xc9\xb1\xd9\x3f\x77\x12\xc2\xb4\x9a\xa8\xbe\xf6\xf4\x19\x00\x00\xff\xff\xad\xfc\xbb\x7f\xf3\x02\x00\x00")

func cmdDefinitionsTmplOperationTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplOperationTmpl,
		"cmd/definitions/tmpl/operation.tmpl",
	)
}

func cmdDefinitionsTmplOperationTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplOperationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/operation.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd0, 0x3a, 0x99, 0xf6, 0x41, 0xfa, 0xdd, 0xc0, 0x2e, 0x3, 0xc1, 0x31, 0x43, 0x68, 0xa5, 0x42, 0xca, 0x90, 0x7f, 0xb7, 0x15, 0x72, 0x87, 0xb, 0x71, 0xa6, 0x44, 0xc8, 0x5a, 0x98, 0x33, 0xa4}}
	return a, nil
}

var _cmdDefinitionsTmplPairTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x55\x4d\x6f\xdb\x38\x10\xbd\xeb\x57\xcc\x0a\x46\x20\x2d\x64\x69\xb1\x47\x2d\x7c\x58\x6c\x7c\x58\x14\x8d\x73\x30\xda\x43\x10\x14\x0c\x35\x96\x09\x53\x24\x41\xd2\x4c\x0d\x55\xff\xbd\x20\x65\x3b\xb2\x65\xe7\xe3\x50\xa0\x3c\x89\x9c\x37\x33\x6f\xde\x70\xa8\xa2\x80\xff\x64\x85\x50\xa3\x40\x4d\x2c\x56\xf0\xb4\x83\x5a\x1e\xf7\xc0\x84\x45\x2d\x08\x2f\x68\x53\xfd\x03\xb7\x0b\xb8\x5b\x2c\x61\x7e\xfb\xff\x32\x8f\x14\xa1\x1b\x52\x23\x28\xc2\xb4\x89\x22\xd6\x28\xa9\x2d\x24\x11\x00\x40\x4c\xa5\xb0\xf8\xdd\xc6\x51\xbf\xad\x99\x5d\x6f\x9f\x72\x2a\x9b\x82\x48\x33\xad\xd0\x15\xb5\x9c\x1a\x2b\x35\xa9\xb1\x70\x7f\x17\x6a\x53\x17\x28\x2a\x25\x99\xb0\xf1\x07\x7c\xa8\xc6\x0a\x85\x65\x84\x7f\xc4\x6b\x6d\xad\xa2\x9c\xe1\xfb\x73\xd9\x9d\x42\x13\x47\x69\x14\x15\x05\xfc\xcb\x39\x10\x47\x18\x27\x4f\x7c\x5f\x7f\x1e\x51\x29\x8c\x2f\xbf\x6d\xa7\xa0\x89\xa8\x11\x26\xdf\x32\x98\x38\x28\x67\x90\xdf\x7b\x0c\x74\x5d\xc8\xe6\x11\x13\x25\x48\x83\xde\x36\x71\xf9\x9d\xff\xfc\x01\x56\xde\x13\x43\x09\x3f\xe0\x8a\x02\xda\xf6\x80\xec\x3a\x78\x66\x9c\x87\x13\x97\xdf\xa2\xa1\x9a\x29\xcb\xa4\x78\x89\x3a\x80\xce\x20\xee\x81\x77\xfd\x41\x1c\x68\xa1\xa8\x3c\x3a\x8d\xde\x20\xf9\x1e\x82\x45\x01\x5f\x99\x5d\x8f\x09\x12\xa5\xf8\x0e\x4e\xb2\x83\x23\x7c\x8b\x60\x25\x2c\x02\x65\xe3\xbd\x97\x6b\x66\x82\x76\xc0\x0c\x6c\x0d\x56\xde\x7e\xb1\xba\xd5\x56\xd0\x51\xb2\xc4\xed\xc1\xcb\x9d\xf2\xfb\x14\xfe\x0c\x3d\x0a\x55\x40\x1b\x24\xd1\x68\xb7\x5a\xc0\xcd\x8b\xa1\x3f\xf7\xeb\x13\xee\xca\x13\xcd\xb2\xa3\xe9\x8b\x67\x5b\x82\xeb\x4f\xba\xa8\x1b\x8a\xe7\xa9\xdf\x13\x6d\xb0\xaf\x56\x85\x4f\x02\x9b\xa9\x83\x86\x28\x5f\x43\xb8\x0f\x60\x38\xa3\x98\xf7\xdc\x03\x3e\x69\x3c\xe0\xc1\x58\xcd\x44\xfd\x18\xe6\x6a\x45\x28\xb6\x5d\x0a\xc9\xc3\xe3\x80\x7c\x06\xa8\xb5\xd4\xe9\xbe\x88\x3e\x5c\x39\x83\x86\x6c\xf0\x1c\xf9\x57\x06\x1c\x45\xd2\xa4\x69\x3f\x68\x8e\x68\xef\xdd\x47\xe8\x8f\x56\x52\xc3\x26\x83\xd0\xe2\xbe\xe5\x0d\xbc\xc8\xe0\x1d\x94\x83\x01\x9d\xa3\xc9\x3c\x33\x4b\xd7\xb0\x19\xa0\xdf\xbe\xdb\x43\xe4\x7b\xee\xf8\x61\x51\x62\xf0\xa4\x1d\xe5\x89\x79\xc0\x47\x87\xac\x2e\x4f\xbc\x0e\xe9\x80\xdc\x28\xd4\xf1\x6e\x8c\x63\x05\x61\x1d\xcc\x40\xbb\x91\xcd\x53\x67\x2b\xef\x1e\xfa\xa6\xcf\xa9\x1e\x73\xf4\xad\xbc\x16\x3c\xb4\x11\x66\xd0\xb6\xc7\x48\x5d\x97\x68\x97\x5e\xc4\xb3\x55\x80\xff\x31\x03\xc1\xf8\x85\xa2\x0e\x6b\x7f\xa5\x05\xe3\x19\xdc\xcc\x7d\x97\xaf\x63\x17\xaa\x84\x38\x5c\xd0\x38\xbb\x0a\x9a\x6b\x5d\xfa\xd4\xd7\x11\xaf\x4c\xca\xf9\xf2\x7a\x97\x87\xa7\x67\x2f\xfe\x2b\xb9\x4f\x07\xed\x7c\x8d\x55\x1f\x9f\x0c\x06\xf3\xdc\x54\xe1\x8a\x6c\xb9\xbd\xdc\x9e\x5f\x23\xe3\x5c\x6b\x3f\x0c\xbe\xf0\xcf\xcc\x34\xc4\xd2\xf5\xef\x2e\x6b\x77\x32\xdb\x17\xa4\xbc\x2a\xa3\xff\xc3\x33\xb1\xc5\x68\x1c\xaa\x7f\xb2\x66\xfe\x57\x80\xa2\x4a\xc2\x36\x3b\x79\x86\x43\xf1\x9b\xec\x40\x55\xb9\x2e\xdd\x3f\xb6\xc3\x87\x7b\xef\x28\x18\x8f\xba\xe8\x67\x00\x00\x00\xff\xff\xa4\x19\x45\x28\xab\x08\x00\x00")

func cmdDefinitionsTmplPairTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplPairTmpl,
		"cmd/definitions/tmpl/pair.tmpl",
	)
}

func cmdDefinitionsTmplPairTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplPairTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/pair.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xab, 0x1e, 0x79, 0x73, 0x22, 0xcc, 0xed, 0x33, 0xba, 0x50, 0x21, 0x70, 0x7a, 0x5b, 0x77, 0xe, 0xd1, 0x8, 0xd9, 0x88, 0x2c, 0x47, 0xcb, 0xd7, 0xc6, 0xa8, 0xfa, 0xe4, 0x63, 0x4c, 0x5c, 0x1}}
	return a, nil
}

var _cmdDefinitionsTmplServiceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x59\x5b\x6f\xdb\xbe\x15\x7f\xf7\xa7\x38\x7f\xc1\x2d\xa4\xc2\x91\xb6\x3d\xba\xf0\xcb\x92\x36\x2d\xba\x25\xc1\x92\xad\xc0\xda\x20\x60\x24\xca\x26\x2c\x91\x2a\x49\x2b\xf1\x5c\x7d\xf7\x81\x17\x49\x94\x2c\xb9\x76\x97\x76\x7e\x08\x6c\xf1\x5c\x7f\xe7\x4a\x25\x8a\xe0\x9c\x25\x18\x96\x98\x62\x8e\x24\x4e\xe0\x71\x0b\x4b\xd6\xfc\x86\x92\x20\x20\x54\x62\x4e\x51\x16\xc5\x79\x12\x09\xcc\x4b\x12\xe3\xb7\x70\x71\x0d\x57\xd7\x77\xf0\xee\xe2\xe3\x5d\x38\x29\x50\xbc\x46\x4b\x0c\xbb\x1d\x84\x57\x28\xc7\x50\x55\x93\x09\xc9\x0b\xc6\x25\xf8\x13\x00\x00\x2f\x66\x54\xe2\x67\xe9\x99\x5f\x84\x79\x13\xf3\x6d\x49\xe4\x6a\xf3\x18\xc6\x2c\x8f\x10\x13\x67\x09\x2e\xa3\x25\x3b\x13\x92\x71\xb4\xc4\x51\xf9\x97\xa8\x58\x2f\x23\x4c\x93\x82\x11\x5a\x73\x1f\xc5\x13\x73\x9c\x60\x2a\x09\xca\x4e\xe1\x5a\x49\x59\xc4\x19\xc1\xc7\xeb\xb2\x88\x08\x43\x1f\x1e\xc1\x21\xb7\x45\x4d\x5e\x88\x63\xec\x42\x84\x0b\x6f\x12\x4c\x26\x25\xe2\xf0\x00\xad\x67\xe1\x0d\x67\x25\x49\x30\xb7\x27\x35\x4e\xfd\xe7\xb7\x46\x58\xfd\xb3\x36\x39\xbc\x35\x5f\xde\x71\xce\xea\xb3\x16\x80\xf0\xba\x90\x84\x51\x31\x99\x44\x11\xdc\x6d\x0b\x0c\x44\x80\x5c\x61\x50\xf6\x43\xca\x78\x27\xdc\x31\xa3\x42\x1a\xb2\x05\x78\xce\x89\xa7\xf9\xad\x26\x40\x25\x22\x19\x7a\xcc\x30\x68\xaf\x42\xcb\xe7\x4f\x76\xbb\x33\xe0\x88\x2e\x31\x4c\x1f\x66\x30\x2d\x61\xbe\x80\xf0\x46\xd1\x28\xe9\x0a\x2b\x45\x41\x52\xa0\x4c\xc2\xb4\x0c\x2f\x33\xf6\x88\x32\xf7\x6c\x5a\x50\xa5\x71\xbe\x50\xc7\x5a\xf9\x77\x90\xec\x06\x89\xb8\xa5\x8b\x22\x65\xb4\xa5\xac\x2a\x78\x22\x59\xa6\x9f\x94\xe1\x05\x16\x31\x27\xda\xe5\x9a\x5a\xa9\xef\x90\x1b\xcf\xa6\xb5\x6b\x0f\x86\xb3\x71\xb4\xb6\x04\xd3\x44\x89\x70\xbe\x06\x23\x20\x10\x9a\xb2\x1f\x81\xf0\x51\xd1\xbc\x24\x08\x56\x86\xf2\x99\x88\x22\x43\xdb\x3a\x86\x60\x3f\x8e\xa0\xc5\x08\x99\xe3\x9b\xfa\xa9\x6c\x34\x60\xdc\xc6\xac\xe8\x29\x35\x07\xe7\x48\xe2\x25\xe3\xdb\xfd\xb3\x31\x78\xcf\x4e\x81\xf7\x70\xfe\x8c\xc1\x76\x0c\x64\x51\x04\x9f\x89\x5c\xed\xe7\x0d\x2a\x8a\x6c\x0b\x1d\x23\xa1\x44\xd9\x06\x83\x64\x50\x17\x8f\xaa\x9d\x15\x11\x3a\xdd\x55\x01\x6d\x04\x4e\xd4\xf9\x60\xd2\xa5\x1b\x1a\xef\x29\xf3\x4b\x4b\xac\x8b\xab\xaa\x02\x78\xa3\xfc\x82\x5d\x13\x2f\x8e\xe5\x86\x53\x78\xad\xf3\xb5\x79\xfa\x09\x6f\xe7\x7b\x29\x3c\x6b\x8e\xff\xa5\x2c\x9d\x43\x69\x9e\x54\x93\x0e\xa0\xce\xd7\x11\x68\x4f\xc8\xca\x97\x4a\xca\x13\x73\x52\x31\xd8\x4e\x77\x65\xd5\x17\x9c\x50\x99\x82\xf7\x4a\xbc\x12\x1e\xf8\x03\xe9\x1a\xe8\xa7\x03\xb9\x1a\x38\x2d\xe4\x12\xcb\xfd\x6c\x58\x62\x39\x98\x0b\x29\x67\x39\xe4\x58\xa2\x04\x49\x14\x6a\x11\x3a\xcc\x3d\x21\x7e\xbe\x67\xb0\x0a\xb5\xdf\x09\xfd\x0c\x1e\x19\xcb\x02\xb0\xb1\x2f\x67\xc0\xd6\xca\xaf\x3c\xbc\xc4\xd2\xaf\x8b\xb0\x2b\xc3\x55\x12\x68\x36\x92\xc2\x1f\x6c\x6d\x65\xb4\xb8\xff\x1b\x73\xa6\x93\xc2\x45\xdd\xa6\x96\xb1\xc2\xa5\x98\x41\x8a\x32\x81\x5b\xe0\x33\x31\xce\xa8\xcd\xff\x0e\xff\x39\xc4\xdf\x06\xce\xfc\xb5\x02\xca\xb0\x0b\x41\x30\x03\xc9\x37\xd8\x12\xd6\x11\x11\x06\xcc\x16\x7b\x1d\x12\x31\x12\x12\x42\x25\x1b\x0a\x89\x38\x22\x24\x33\xd8\x2b\xc7\x7d\x22\x8b\xad\x75\x21\x0f\x6f\x8f\x8a\xce\x0c\xca\xc0\x01\x60\xa4\x1c\x07\xab\x51\x49\x13\x05\x8a\x71\xa7\x24\x25\xce\x8b\x4c\xad\x71\x9e\xea\x3e\x1e\xe4\x68\x8d\x6f\x33\x35\x81\xfc\xa1\x4a\x0c\xf4\x43\xfc\x34\x2e\x42\xb8\x32\x6a\x09\xd3\x32\x7c\xbf\xa1\xf1\x88\x66\x05\xac\xee\x84\x3f\x62\xed\xfb\x98\xe0\x94\xd0\x56\x6f\x77\xc6\x29\xa7\x09\x4d\xf0\x33\x84\xf0\xa7\x91\x6e\x32\x55\xd1\x75\x09\xff\xac\x65\xd7\xc7\x7d\x0c\x0d\x79\x6f\x0a\x8e\x23\xa8\x8c\x98\x96\xfb\x9d\xc7\x7e\x1d\xf2\xe2\xb0\x13\xee\x59\x4a\xc7\xed\x9e\xa6\x54\xd7\xfc\x34\xa5\x43\xbd\xb4\xae\x88\xa2\x69\xfc\x36\xcb\x14\x5b\x55\xfd\x1d\x15\xb0\x62\x59\x22\x00\xa9\x09\xd6\xdd\xc6\x4c\x53\x41\xfc\x10\xf3\x02\x72\x54\x7c\x11\x92\x13\xba\xbc\x17\x92\x6f\x62\xb9\xab\xda\xa1\x13\x45\xf0\x0f\xfc\x6d\x43\x38\x4e\x1c\x99\x03\x98\xeb\x79\x68\xdd\x68\x38\xfa\xed\x1e\x11\x1e\xbe\xdf\x64\x99\x2d\x96\x39\x34\x0a\x9d\x41\xd6\xeb\x1e\xd6\x0a\x33\x7f\x51\x76\xbc\x15\x0d\xc7\x0b\x5a\x71\xd9\x5c\xac\x8e\x36\xa3\x65\x79\x09\x3b\x7e\x98\x10\xf5\x52\x5f\x20\xae\x36\x13\x23\x52\xf3\xe8\x35\x7f\x8c\xcb\xd0\x39\x3b\x88\xf6\x0f\xbe\xdc\xeb\xdd\x64\xf2\x0b\xf3\x61\x3f\xe7\x3b\xc7\xb6\x2b\xff\x96\xec\xf8\x80\xc4\x61\xb3\xd4\xb4\x3e\xd6\x83\x93\x7d\x78\xb9\xdc\x6a\x97\x37\xad\xfe\x02\xa7\x68\x93\xc9\xff\xc5\xd3\x3d\x7b\x5f\x28\x7c\x6e\x3a\x73\x81\x6f\x46\xb2\x53\x4f\x7f\x4d\x61\x77\x65\xa1\x7b\xb6\x9e\xfb\x6f\x46\x52\xba\x5d\x03\x0e\x89\xf6\x59\x21\x9b\x34\x0f\xc0\x1f\x93\x36\x03\xac\x2e\xd5\x41\xb3\x09\x08\x85\xe8\x7c\x01\xaf\x47\x18\x7a\x95\x34\x07\xa5\x68\xe6\x3a\xad\x77\x17\xa1\xd7\x3d\xb4\xc6\xbe\xd3\x86\xf5\x1b\x9a\x14\xc5\x78\x67\x37\x3c\x75\x39\x7f\x50\x9b\xca\x7c\x61\x53\x41\x9b\xbd\x1b\x8a\xb9\xbf\x42\xe2\x86\xe3\x94\x3c\x1b\x53\xbc\x2b\xfc\xe4\x05\x6e\x0c\x48\xaa\x84\x99\x4d\x73\x7c\x32\x7c\x29\xc3\x4f\x78\x7b\xff\xd6\xd9\x2d\xeb\x8f\xdd\x83\x28\xc9\x66\xed\xbb\x87\x2b\xfc\xa4\x30\xfc\x27\x15\x9b\xa2\x60\x5c\xe2\x44\xbf\x86\xf0\xed\x0e\x04\xcd\x1e\x34\x92\x51\x06\x0e\xab\x15\x16\x50\x86\x7a\xb5\x74\x11\xb3\x5e\x32\xde\xed\x2a\x9d\x72\x1e\x2c\x0b\x35\x03\x4b\x70\x60\x6d\x9e\xb2\x75\x9b\xea\xee\xc6\x62\xb3\xf2\x03\xa2\x49\x86\x81\xef\xb7\xbc\x93\xdb\x9d\x5a\x11\xd6\x78\xab\x89\x46\x2a\xa7\xc5\x42\x87\x67\x51\x63\x62\x98\x3b\xe3\xe2\xac\xaa\xee\xdd\x80\x9e\x16\xa4\xda\x42\x13\xa1\x61\xf1\x43\x61\x53\xe0\xef\xeb\x51\xa5\x10\xea\x24\x52\xfe\xe9\x7b\xbf\xdd\xf2\xdd\x16\xd0\x97\xd7\x4b\x81\x16\x6c\xb6\xdf\xcb\x4f\x9e\xf2\xbf\x16\xec\x0e\x06\xd6\xff\xba\xa3\x36\x10\x34\xb7\x9a\x5f\x02\xd2\x72\x60\x58\xfc\xcc\xa0\xf8\x6d\x30\x1d\x33\x95\x8e\x86\x73\xa0\x7d\xfc\x2c\xc2\x8e\x6d\x07\xed\x1a\xa9\xb1\x01\xa5\x8d\xc6\x11\x49\xa3\x5d\xb0\xdf\x7f\x6c\x05\x1b\x25\x33\x55\xc9\xb6\x19\x8e\x5d\xad\x9c\xbb\xd9\xff\xfb\x7a\x55\x9b\x72\xc2\x15\x6b\xd4\x9d\x97\xb8\x67\x1d\x7b\xcd\x72\x96\x0e\xf3\x75\xe8\x3d\x72\x14\xd5\xe4\xfa\x55\x60\x6d\xa7\x59\x53\x62\x8e\x15\x02\x08\xec\x7f\x48\xe0\x71\xab\x3c\x52\x99\xe0\xbc\x9a\xf0\x05\xbc\xd9\xed\xa6\x05\xad\x5f\x38\xd8\x75\x64\xb7\x53\x3a\x6f\x10\x47\xb9\x08\x6f\xf5\x3a\x60\x5f\x1b\x99\xc9\xa2\x52\xc1\x3d\x68\xf3\x31\x96\xcf\xca\x3f\xab\x35\xfc\x2b\x8a\xd7\x4b\xce\x36\x34\xf1\x83\xfe\x4b\x1c\x11\xb6\x1a\x3f\x13\xb9\x3a\x37\x3c\x7e\x2c\x9f\x67\xd0\xb1\xe0\x1c\x65\x19\xe6\x75\xd5\x0c\xc1\xe4\xf0\x1f\x40\xec\x90\xcf\x3d\x0b\x1a\x0f\xec\xb3\x9e\x45\x27\x61\x92\xe0\x14\x73\xad\xdc\x0f\xba\x95\x6b\xdb\x97\x5c\xd5\x39\x61\xc5\xdf\x20\xb9\x32\x4e\xf7\xcb\xd6\x36\x09\x44\x13\xf0\xf1\x37\xcb\xec\x79\x81\xfd\x45\xc1\xb3\xff\x0e\xe8\x2e\x5b\x2d\xbb\x65\x59\x80\x37\xfb\xea\x7d\xf5\xbc\x21\x05\xbd\xa6\xa6\x3e\x98\x73\x58\x80\x08\x53\xc6\x73\x24\xcd\xc4\x6e\xe6\xf9\x75\xd1\x5b\x51\x5b\x3d\x55\x05\x4e\xaf\x73\x92\x40\xaf\x3d\x85\x3c\xbc\x38\xab\x0f\x2b\xe4\xcc\xaa\x3f\xb8\x42\xeb\x09\x14\xb8\x9d\x5f\x31\xfd\xb1\x50\x1d\x6b\x70\x29\x71\x5a\xf0\x60\x6a\x76\x8b\xf4\x1c\xe5\x58\xd5\xe8\x40\x7a\xde\x71\x92\xff\x0d\x09\x69\xf3\xf4\x1d\x4d\xc4\x93\xce\xa6\x3c\x47\x55\xa5\x1c\x08\xf6\x5a\xe6\x7f\x03\x00\x00\xff\xff\x8b\xad\x25\xeb\x24\x1d\x00\x00")

func cmdDefinitionsTmplServiceTmplBytes() ([]byte, error) {
	return bindataRead(
		_cmdDefinitionsTmplServiceTmpl,
		"cmd/definitions/tmpl/service.tmpl",
	)
}

func cmdDefinitionsTmplServiceTmpl() (*asset, error) {
	bytes, err := cmdDefinitionsTmplServiceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "cmd/definitions/tmpl/service.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xec, 0xfa, 0x6d, 0x45, 0xc5, 0xe, 0x97, 0x8b, 0xb9, 0xb0, 0x10, 0x4d, 0x1c, 0x99, 0xc0, 0x7c, 0xce, 0xb1, 0x9c, 0xcc, 0x80, 0x98, 0xbf, 0x8e, 0x9e, 0xeb, 0xe0, 0x90, 0x81, 0x3, 0x17, 0x6d}}
	return a, nil
}

var _definitionsInfosHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x93\x3f\x53\xc3\x30\x0c\xc5\xf7\x7c\x0a\x9d\x17\x16\xda\x89\xb2\x31\x51\x06\x06\x60\xe9\x1d\x63\xa3\xc6\x6a\x2b\x1a\x5b\xc1\x56\x5a\x0a\xc7\x77\xe7\xdc\x3f\xd7\xbf\x49\x8f\x66\xca\xc5\x7a\x3f\x3d\xf9\x29\x19\xfb\xb1\x80\x91\xd1\x07\x15\x6a\xc0\x38\x52\x34\x60\x0a\xf1\x4a\x5e\x3b\xce\xf6\x0c\xfc\x64\x00\xba\xac\x08\xb6\xcf\x03\x98\xa8\x81\xfd\xc4\x64\x00\x96\x63\x55\xe2\x72\xe8\xd1\x51\x3a\x79\x5c\x4b\x5f\xfa\x3d\x93\xfd\x5e\xc0\x27\xea\x1e\x7f\x8f\xdb\xa4\x24\xc5\xc9\x7f\x1d\x3d\x0d\xb0\x85\xc8\xf6\x90\x77\x80\xa2\xaf\x4a\x82\xa6\x8f\x1a\x6a\xca\x00\x0a\x71\x8e\xbc\xa6\xaa\xe7\x3e\x70\x04\x9d\x12\xd4\x9e\x3f\x6b\x82\x19\x2d\x81\x3d\x44\x0a\x73\x2e\xa8\xdb\xdc\x32\x19\xbb\xb2\xe9\x6b\x9a\x69\xd3\x36\x50\x89\xca\x73\x82\x0a\x75\x0a\x2a\x0b\x0c\x36\x6e\xdb\xdf\x44\x78\x97\x30\xeb\x73\x68\x31\x12\xf9\xfb\xe8\xfe\xd9\xeb\xfd\x5d\xb3\xe0\x28\xb0\xb5\xf3\xb7\x55\xd1\x20\x1d\x5d\x70\x9f\x6a\x20\x4e\xa5\x2e\x2d\x8c\x08\xc4\x13\xc8\x18\xf2\x31\x97\x94\xdf\x42\x1e\x35\x10\xba\xf4\x66\x39\xe4\x20\x01\x72\xf6\x73\x2c\xd9\xe6\x2d\x43\xd4\x95\x45\x25\x3b\x44\x3d\x1c\x45\xd9\x51\x77\xc0\x8e\x76\xd2\xa8\x12\x70\x42\x3b\x6d\x29\x05\x2a\x8b\x6f\x5f\xc2\x13\xd9\x49\x80\xe7\xf2\xdb\x5c\x40\x23\x64\x21\x61\xd6\xb1\x1c\xae\x07\x45\x45\xe5\xa8\x5c\xac\x7e\xa9\xda\x6b\x6b\x96\xe7\x75\x2d\x2b\xf0\x17\x00\x00\xff\xff\x34\xac\xaf\x48\x1c\x04\x00\x00")

func definitionsInfosHclBytes() ([]byte, error) {
	return bindataRead(
		_definitionsInfosHcl,
		"definitions/infos.hcl",
	)
}

func definitionsInfosHcl() (*asset, error) {
	bytes, err := definitionsInfosHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/infos.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x31, 0x6a, 0x76, 0xf8, 0x51, 0xcf, 0x20, 0xd, 0x59, 0xb, 0xdb, 0xb6, 0x78, 0xbd, 0x86, 0x89, 0xab, 0xf, 0xc, 0x12, 0x47, 0x81, 0x53, 0x7f, 0x18, 0xfe, 0x49, 0xa6, 0x65, 0x89, 0xea, 0x70}}
	return a, nil
}

var _definitionsOperationsHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x57\x4d\x8b\xe3\x46\x10\xbd\xfb\x57\x14\xba\x2c\x2c\xc6\xa7\x90\xdb\x9c\x16\x02\x39\x2c\x09\x3b\x87\x1c\xc2\x32\xf4\x48\x25\xbb\x82\xd4\x2d\xaa\x4b\xf6\x7a\xc3\xfc\xf7\xa5\x5a\xdf\x9f\x96\x86\x99\xc3\x60\xd4\xaf\x3e\xde\xeb\xea\xa7\xd6\x81\xac\x20\xa7\x26\x46\x88\x62\x57\x10\x72\x04\xff\x1f\x00\x12\xf4\x31\x53\x21\xe4\x2c\x3c\x41\x44\x1e\xe4\x82\xd0\x81\x53\xc7\xf0\xc5\x15\xf7\x53\x74\x38\x00\xb8\x22\x44\xdf\xab\xd8\x49\xf4\x8d\xb2\x0c\x74\x1d\x8c\x85\xbf\x5e\xff\xc3\x58\xc0\x31\xe4\x65\x26\x54\x64\x08\xae\x7a\x44\x36\x14\xf1\xc8\x57\x8a\xf1\x14\x85\x54\x85\x61\x93\x7b\x08\x7f\x4f\xf0\x6f\xe4\x39\x8e\x8e\x10\x25\x5e\xa2\xef\x07\x80\xb7\xc3\x5b\x9f\x42\x42\xfc\x92\x91\x97\x65\x1a\xa5\xc7\x24\x74\x9f\x10\x63\x2c\x8e\xef\xf0\x6a\xf4\x99\x17\xc7\xe6\xdc\x96\x07\x71\xa0\x99\xea\xe6\x3c\x94\x36\x41\x06\xa3\x71\x1d\x69\x45\xbc\x24\xc4\x6b\xc4\x19\xa5\x64\x5b\x25\x33\xe0\x0b\x8c\x29\xa5\xb8\xce\x33\x47\x51\xf3\x7d\x0f\x2b\x8c\xbe\xcc\xc4\xb7\x2b\x8e\x96\x48\x7b\x3c\xe7\x68\xc5\x7f\x34\xfb\x26\xef\x88\x3e\x00\xe6\xaf\x98\x40\x6f\x5f\x2a\x20\x6a\xeb\x23\x71\xda\xe6\xd6\x54\x1a\x56\xbb\x92\x79\x97\x40\x7e\x56\x20\xb2\x09\xfe\x78\xe9\x3a\xdc\x3c\xe1\x21\xb0\x51\xa8\x0a\xdf\x4a\x9e\x2c\xc9\xcb\xa0\xf0\x1a\x7b\x45\xeb\xe1\x58\x2a\x38\xd5\xa0\x30\x72\x59\x12\x01\xcf\xb5\x0a\x55\x2b\x37\x26\xc1\xed\xbd\x04\x38\x18\xad\xa8\x67\x52\xdc\xbe\xc6\xb4\xfa\x11\x22\xd6\x7f\x21\x4a\x7f\x78\xfa\x89\x73\x3b\x93\xbb\xeb\x9e\xfd\xf8\xea\xae\xd8\x1d\x3e\x0d\x5e\x23\xa2\xeb\xda\xfb\xc7\xd9\x4b\xc1\x98\xd2\x8f\xad\x67\xac\x42\xef\xb4\x97\x2a\x68\xe4\x30\xd5\xc3\x0f\x35\x99\x3a\xe5\x1e\x9f\xa9\xd9\xef\xb5\x9a\x4d\x32\x4c\x7c\xa6\xd5\x61\xbb\xd5\x8c\xfa\xdb\xec\x36\xef\x51\x68\xde\x68\x18\x4d\x7c\xd9\x33\xd0\xdf\x34\xa0\xdb\xec\x10\xbf\xd6\x77\xc1\xee\x4a\x89\x9e\xce\x9b\xb9\x1f\xe1\x76\xa1\xf8\x02\xb1\xb1\x10\x22\x43\x81\x6a\xa0\xf6\xbb\x46\xc9\xd9\x1c\xa5\x91\x6b\x86\x15\x6b\x32\x78\x02\xe1\x12\x9b\xc6\xcd\xab\x63\xd9\x62\x2e\x01\xa8\x13\xfa\xd0\x43\xfa\x0e\x16\xbb\xbc\xc8\x50\x70\x4b\x85\x06\xdb\x15\x01\x63\x13\xc8\x91\xcf\xa8\x02\xe5\xb5\xa7\xc1\x1f\x94\x2d\xba\x40\x5b\x7f\xa4\x45\x18\xda\xf9\xfd\xd5\x5d\xc8\x0d\x59\x31\x64\xeb\xab\x8d\x1a\xe8\x68\xe0\x7d\xef\xc2\xc4\x68\x64\xd5\xc0\x2a\x04\x18\xb0\x78\x6b\x12\xe9\x4b\xc9\x8b\xb1\x8b\x0e\x66\x4d\x8e\x4b\x43\x2b\x8e\x71\x20\x6c\x82\x2a\xd5\x5a\x0f\x15\x42\xc5\xdc\x5b\xbf\xa9\x71\xc6\xd5\xfd\x3a\xa3\xce\xc3\xd5\x64\x94\x4c\x6b\x84\x43\xb2\xee\xd8\xbb\xf8\xea\xb9\x7f\xe8\x0b\x26\xcb\xa6\x9d\x34\xce\x24\x17\xf2\xa3\x8e\xa6\x65\x67\xcd\xc1\x8b\x11\xf2\x42\x31\x19\xbb\xdd\x21\x9e\x9b\x28\x93\x75\xa3\xe3\xbb\x87\xeb\x47\xa1\xb4\xd2\x34\xfb\xc9\x43\x1b\xe6\x8f\xe0\xcb\xf8\x02\xc6\xc3\x33\xfd\xc4\x23\x7c\x51\xe4\x22\x9d\x3a\x6a\x9e\x54\xa5\xd3\x76\x42\xa3\x03\xd1\x91\xda\x3e\x8b\xed\x47\x44\xca\x2e\x7f\x30\x1e\xad\xe3\x35\x23\x90\xa3\x98\xc4\x88\xd9\xf0\x36\x8d\x4b\x66\xf5\x8f\x86\xe4\x27\x0f\x4d\xf4\x82\x56\xba\x3c\xa8\xc6\x68\x92\xf5\x4a\x26\x09\x22\xa5\x94\xe9\x16\xf5\x72\xcf\x32\x39\x42\x74\x5b\x98\x76\x3b\x28\xac\xbb\xb6\x56\x58\xd7\xc3\x25\x4f\x2e\xfa\x12\xd6\x53\x48\x36\x75\xe0\xd2\xee\xca\xb4\xff\x25\xe2\xa6\x17\xcf\xc7\x57\x4d\xe5\xdc\x98\x72\xba\x6c\xca\x0d\xff\xa5\xbb\xbf\x6d\xe7\xf3\x90\x12\x66\x49\xf8\x30\xaa\xaa\xcb\xbd\x40\x2d\xea\x85\xc9\x9e\xa3\xc3\x5b\x8b\x68\xec\x60\x11\x81\x3c\xca\x81\xcc\x8e\x7b\x80\xea\x8e\x3b\x80\x90\x95\x1e\x20\xcc\xc4\x60\xfd\xb9\x9a\xa7\xaf\xba\xd0\xe1\xec\x24\xc9\xef\xbf\xf5\x97\xd5\xe7\xd6\x7b\x75\xc3\xf5\xcf\xd5\x29\xe9\x03\x68\x16\xf1\xa7\x20\x1b\x19\xb0\x2a\x0c\xb1\x1f\x82\x4f\xa7\xd3\xe7\xbf\x0d\x0d\x51\x72\x79\xd0\x53\xff\xea\xba\x08\x1a\x49\x4c\xee\xf4\x0d\x4d\x82\xfd\x52\xfa\x4a\x1e\x8a\x58\x5f\x05\x7a\x90\x31\xbb\x1a\x32\x43\x2f\x7c\x8f\xac\xca\xad\x1f\x02\xeb\x4d\x77\xbe\x38\xb7\xb9\xad\x6f\x0f\x22\x26\x1d\xd6\xce\x32\xd7\x62\x78\x79\xcd\xa5\xee\xa3\xf4\xd6\xb6\xde\xe6\x6d\xa2\xed\x3f\x7a\xea\x42\x92\x5f\x01\x00\x00\xff\xff\x74\xc4\x69\xb2\xfe\x11\x00\x00")

func definitionsOperationsHclBytes() ([]byte, error) {
	return bindataRead(
		_definitionsOperationsHcl,
		"definitions/operations.hcl",
	)
}

func definitionsOperationsHcl() (*asset, error) {
	bytes, err := definitionsOperationsHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/operations.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcd, 0x75, 0x8f, 0xce, 0x5f, 0x6e, 0x9d, 0x87, 0x51, 0x16, 0xe9, 0xe1, 0x5d, 0x4d, 0xc3, 0xbb, 0xfa, 0xd3, 0xe6, 0x67, 0xa2, 0xc7, 0x60, 0x8f, 0x83, 0x65, 0x90, 0x76, 0x19, 0x7e, 0xb4, 0x5a}}
	return a, nil
}

var _definitionsPairsHcl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x95\x4f\x6f\xd3\x40\x10\xc5\xef\xf9\x14\x4f\x39\xb5\xa8\x4a\x2f\xc0\x8d\x0b\x48\x48\x1c\x10\x48\xc0\x09\x50\xb4\xd9\x1d\xc7\x43\xec\x1d\x33\x3b\x4e\x1a\x10\xdf\x1d\xad\xff\x34\x51\xeb\xa4\x55\x7a\xab\x66\xe6\xbd\xdf\xbe\x9d\x8d\x67\x8d\x63\xc5\xdc\x97\xe4\x37\xa9\xad\xe7\xf8\x3b\x03\x6c\xdf\x10\x86\xbf\x37\x98\x27\x53\x8e\xeb\xf9\x0c\x08\x94\xbc\x72\x63\x2c\xb1\x2b\x34\xe4\xb9\xd8\x63\x9c\x46\x21\x0a\x2b\x39\x41\xe9\x77\x4b\xc9\x6e\xe0\xa5\xad\x02\x56\x84\x36\x51\x80\x4b\xf0\x12\x8d\xa2\xa1\x0e\xaf\x20\x0a\x32\xb7\x9e\xcf\xfe\x8d\x18\xb9\x78\x67\x93\x14\x43\x6d\xf1\x6e\xe8\x79\x8c\x33\x74\x80\x23\x5c\x55\x8d\x0c\x7d\x63\xe1\xda\xca\x1e\x4a\xbd\x75\x7e\xb3\x56\x69\x63\xb8\xba\x3e\x82\x50\x0a\x14\x8d\x5d\x35\xc9\xf1\xe2\x50\x5f\x7c\x56\xd9\x72\x20\x3d\x93\x4d\x29\x3b\x98\xa0\xe9\x3b\x71\x18\xee\xc2\x4a\xa4\x5b\xf6\x94\x93\x48\x26\xea\xd6\x94\xa5\x1a\xa7\x89\xf4\x70\xf2\x23\xc3\x5c\x39\xa0\x52\x0c\x8d\x70\x9c\x0e\x6c\x2c\x5e\x40\x39\x8e\x3e\x9f\xf1\x60\xf6\x80\xf0\xae\x61\xa5\x49\xbe\xcc\x7d\x9a\x68\x57\x52\x84\x95\x84\x56\xf3\x55\x5a\xab\x91\x02\x56\x7b\x28\x39\x5f\x62\xc7\x55\x85\x41\xfc\x31\x4d\xf7\xef\x87\xac\x3f\x72\x94\x66\xcd\xd2\x57\x4c\xd1\x96\xd2\x99\xa5\xe9\xdb\xcd\x8d\x7d\xdf\xe2\xd3\xd0\x37\x01\x49\x4d\x07\x99\xf9\x06\xb5\x61\xf7\x09\x59\x00\xbd\xc2\xc1\x9e\x63\xa0\xbb\x0b\x52\xc8\x82\xdd\x2c\xa4\xe8\x5f\x56\xa2\x75\x4d\xfd\xcc\x93\x87\xae\xc4\xbb\xac\x78\xd9\xab\xce\xde\xa3\xc2\xc9\x45\x18\xad\xa2\xab\xa7\x6f\xf9\x59\x36\x83\x1c\x3a\x95\x7b\x4d\x29\x8a\x44\xd3\xbb\xcd\xd1\x5e\xbf\x3c\x23\xda\x8f\x4e\xfc\x1e\x8d\x4e\xdd\xfe\x24\xa2\x4d\x5e\xfb\xae\x65\x18\x59\x51\x21\x4a\x79\xcb\xc2\x99\x8c\xb3\xf9\x88\xd9\xa8\xfc\x22\x3f\xcd\xf9\xe4\xe9\x87\xe1\xee\xe4\xb7\x1c\x0e\xc4\xe7\xc2\xce\x70\x4b\xef\xaa\x6a\xe5\xfc\x66\x59\xb4\xd1\x4f\x9a\xe7\xc2\xd5\xf7\x9f\xab\xbd\xd1\xf5\xd9\x97\xe6\x0c\x26\x41\x40\x5b\xd2\x3d\x8c\x6b\xc2\xae\x8f\x00\xc1\x99\x43\xa1\x52\x23\x49\xab\xfe\x08\x22\xf1\x9f\x93\xef\xfa\xec\xdd\xe4\xc1\xa7\x6e\x46\x62\xb5\xef\x01\x2a\xae\xd9\x28\xdc\x7f\x37\x32\xd0\x33\x2f\x66\x27\xba\x59\x06\xd6\xcb\xf7\x32\x2b\x20\xb0\x9e\x58\xff\x9b\x21\x31\x69\x48\xfb\x67\xd2\xc1\xaf\x72\x76\x95\x33\xde\xd2\xfd\x76\x05\xd6\x05\x46\x20\x7c\xfc\xf6\xe5\x2b\x92\x39\x35\xec\xd8\x4a\xdc\x76\x06\xbd\xd8\x98\xc3\x60\x97\x8e\xc6\x46\xf5\xf1\x73\x66\x82\x5b\x70\x81\x28\x86\x44\xb6\xc0\x8f\x88\xf7\xa2\x28\xd2\x43\x15\x74\x6c\x31\xc8\x2e\xa1\xa9\x9c\x15\xa2\xf5\x4d\x77\xc4\x15\x95\x6e\xcb\xa2\xe0\x84\x36\x06\x2a\x38\x52\x58\xe4\x10\xff\x07\x00\x00\xff\xff\x02\x6f\xed\xe4\x1b\x08\x00\x00")

func definitionsPairsHclBytes() ([]byte, error) {
	return bindataRead(
		_definitionsPairsHcl,
		"definitions/pairs.hcl",
	)
}

func definitionsPairsHcl() (*asset, error) {
	bytes, err := definitionsPairsHclBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "definitions/pairs.hcl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa1, 0x3f, 0x7d, 0x7e, 0xac, 0x4, 0x4f, 0x97, 0x5c, 0x8e, 0x38, 0x8f, 0xee, 0x9f, 0x8e, 0x5c, 0x60, 0x2b, 0x18, 0xf5, 0x28, 0x55, 0x14, 0xd4, 0x7c, 0xe1, 0x51, 0x4d, 0x6d, 0x41, 0x30, 0x70}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"cmd/definitions/tmpl/function.tmpl":  cmdDefinitionsTmplFunctionTmpl,
	"cmd/definitions/tmpl/info.tmpl":      cmdDefinitionsTmplInfoTmpl,
	"cmd/definitions/tmpl/object.tmpl":    cmdDefinitionsTmplObjectTmpl,
	"cmd/definitions/tmpl/operation.tmpl": cmdDefinitionsTmplOperationTmpl,
	"cmd/definitions/tmpl/pair.tmpl":      cmdDefinitionsTmplPairTmpl,
	"cmd/definitions/tmpl/service.tmpl":   cmdDefinitionsTmplServiceTmpl,
	"definitions/infos.hcl":               definitionsInfosHcl,
	"definitions/operations.hcl":          definitionsOperationsHcl,
	"definitions/pairs.hcl":               definitionsPairsHcl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"cmd": &bintree{nil, map[string]*bintree{
		"definitions": &bintree{nil, map[string]*bintree{
			"tmpl": &bintree{nil, map[string]*bintree{
				"function.tmpl":  &bintree{cmdDefinitionsTmplFunctionTmpl, map[string]*bintree{}},
				"info.tmpl":      &bintree{cmdDefinitionsTmplInfoTmpl, map[string]*bintree{}},
				"object.tmpl":    &bintree{cmdDefinitionsTmplObjectTmpl, map[string]*bintree{}},
				"operation.tmpl": &bintree{cmdDefinitionsTmplOperationTmpl, map[string]*bintree{}},
				"pair.tmpl":      &bintree{cmdDefinitionsTmplPairTmpl, map[string]*bintree{}},
				"service.tmpl":   &bintree{cmdDefinitionsTmplServiceTmpl, map[string]*bintree{}},
			}},
		}},
	}},
	"definitions": &bintree{nil, map[string]*bintree{
		"infos.hcl":      &bintree{definitionsInfosHcl, map[string]*bintree{}},
		"operations.hcl": &bintree{definitionsOperationsHcl, map[string]*bintree{}},
		"pairs.hcl":      &bintree{definitionsPairsHcl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
