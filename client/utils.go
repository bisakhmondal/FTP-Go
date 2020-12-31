package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

//func UPLOAD_HANDLER(conn net.Conn, name string){
//	inputFile, err := os.Open(name)
//	if err !=nil {
//		fmt.Println(err.Error()))
//		return
//	}else{
//		conn.Write([]byte("1"))
//	}
//	time.Sleep(100*time.Microsecond)
//
//	defer inputFile.Close()
//	fileReader := bufio.NewReader(inputFile)
//
//	clientConnWriter := bufio.NewWriter(conn)
//
//	io.Copy(clientConnWriter, fileReader)
//
//	clientConnWriter.Flush()
//}

func DOWNLOAD(conn net.Conn, name string, fileSize int64){
	outputFile, err := os.Create(ROOT + name)

	if err != nil {
		fmt.Println(err)
	}

	defer outputFile.Close()
	var fileSizeReceived int64
	for {
		if (fileSize - fileSizeReceived) < BUFFERSIZE {
			io.CopyN(outputFile, conn, (fileSize - fileSizeReceived))
			conn.Read(make([]byte, (fileSizeReceived+BUFFERSIZE)-fileSize))
			break
		}
		io.CopyN(outputFile, conn, BUFFERSIZE)
		fileSizeReceived += BUFFERSIZE
	}
	fmt.Println("File Downloaded successfully")
}

func SendFile(conn net.Conn, name string){
	inputFile, err := os.Open(ROOT + "/" + name)
	defer inputFile.Close()

	if err !=nil {
		fmt.Println("ERROR: ", err.Error())
		conn.Write([]byte("-1"))
		return
	}else{
		stats,_ := inputFile.Stat()
		//send file Size
		conn.Write([]byte(strconv.FormatInt(stats.Size(),10)))
	}
	//time.Sleep(100*time.Microsecond)
	buffer := make([]byte, BUFFERSIZE)
	for {
		_, err := inputFile.Read(buffer)
		if err == io.EOF{
			break
		}
		conn.Write(buffer)
	}
	fmt.Println("File Uploaded")

}