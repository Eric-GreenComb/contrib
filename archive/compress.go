package archive

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
)

// Compress compress string
func Compress(str string) string {
	var buff bytes.Buffer
	w := zlib.NewWriter(&buff)
	w.Write([]byte(str))
	w.Close()

	com := string(buff.Bytes())
	return com
}

// Decompress decompress string
func Decompress(comp string) string {
	buff := bytes.NewBuffer([]byte(comp))
	r, err := zlib.NewReader(buff)
	defer r.Close()

	if err != nil {
		panic(err)
	}
	undatas, _ := ioutil.ReadAll(r)
	return string(undatas)
}
