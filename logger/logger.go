package logger

import (
	"os"
	"path/filepath"
	"time"
)

//Data struct data
type Data struct {
	origin string
	prefix string
	text   string
}

//Log entrypoint logger
func Log(origin, prefix, text string) *Data {
	return &Data{
		origin: origin,
		prefix: prefix,
		text:   text,
	}
}

//ToFile write log into a file inside /log directory
func (d *Data) ToFile() bool {
	os.MkdirAll(filepath.Join(".", "log"), os.ModePerm)

	f, err := os.OpenFile(time.Now().Format("log/log-2006-01-02.log"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	defer f.Close()

	f.WriteString(time.Now().Format(time.RFC3339) + "-> [" + d.prefix + "] (origin: " + d.origin + ") : " + d.text + "\n")
	f.Sync()

	if err != nil {
		panic(err)
	}

	return true
}
