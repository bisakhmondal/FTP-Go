package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func getDIR(conn net.Conn){
	files, err := ioutil.ReadDir(ROOT)
	if err!=nil{
		conn.Write([]byte(err.Error()))
		fmt.Println(err)
		return
	}
	fileINFO := ""
	for _, file := range files{
		if file.IsDir() {
			fileINFO += "D" + file.Name() + "|"
		}else{
			fileINFO += "F" + file.Name()+"|"
		}

	}
	conn.Write([]byte(fileINFO))
}


func changeDIR(conn net.Conn, dir string){
	if dir==".."{
		splits := strings.Split(ROOT,"/")
		ROOT = strings.Join(splits[:len(splits)-1],"/" )
	}else{
		tempROOT := ROOT  + "/"+ dir
		_, err := ioutil.ReadDir(tempROOT)
		if err!=nil{
			conn.Write([]byte(err.Error()))
			return
		}
		ROOT = tempROOT
	}
	conn.Write([]byte(ROOT))
}

func SendFile(conn net.Conn, name string){
	inputFile, err := os.Open(ROOT + "/" + name)
	defer inputFile.Close()

	if err !=nil {
		conn.Write([]byte(err.Error()))
		return
	}else{
		stats,_ := inputFile.Stat()
		//send file Size
		conn.Write([]byte(strconv.FormatInt(stats.Size(),10)))
	}
	buffer := make([]byte, BUFFERSIZE)
	for {
		_, err := inputFile.Read(buffer)
		if err == io.EOF{
			break
		}
		conn.Write(buffer)
	}
	fmt.Println("File Sent")

}

func GetFile(conn net.Conn, name string, fileSize int64){
	outputFile, err := os.Create(ROOT + "/" + name)

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
	fmt.Println("File Received successfully")
}

func deleteDIR(conn net.Conn, name string){
	fmt.Println(ROOT + "/" +name)
	err := os.Remove(ROOT + "/" +name)
	if err != nil{
		conn.Write([]byte(err.Error()))
		log.Println(err)
		return
	}
	conn.Write([]byte("File Successfully Deleted"))
}