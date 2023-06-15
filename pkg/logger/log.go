package logger

import (
	"io"
	"log"
	"os"
	"time"
	"void-project/pkg"
)

func Log(info string) {
	file, err := os.OpenFile(pkg.GetRootPath()+"/runtime/log/"+time.Now().Format("2006-01-02")+".txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	multiWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(info)
}
