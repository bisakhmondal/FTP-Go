package main

import (
	"FTP/server/credentials"
	"bufio"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)
const (
	PORT = "9090"
	BUFFERSIZE = 4096
)

var  ROOT = "/filestore"
//dynamic root dir
func init(){
	cdir, _ := os.Getwd()
	splits := strings.Split(cdir, "/")
	ROOT = strings.Join(splits[:len(splits)-1],"/" )+ ROOT
}


func main(){
	server, err := net.Listen("tcp", ":"+PORT)

	defer server.Close()

	if err != nil{
		log.Fatal(err)
	}
	log.Println("TCP server is UP @ localhost:", PORT)

	for{
		connection, err := server.Accept()
		if err != nil{
			log.Println("Client Connection failed")
		}
		go HandleServer(connection)
	}
}


func GetCred() *credentials.CredArr{

	f, _ := os.Open("credential.json")
	var creds credentials.CredArr
	creds.FromJSON(f)
	return &creds
}

func AuthenticateClient(conn net.Conn){
	reader := bufio.NewScanner(conn)
	validated := false
	conn.Write( []byte("Connection Established"))
	CREDS := GetCred()

	//validate user
	for !validated {
		reader.Scan()
		uname := reader.Text()
		reader.Scan()
		passwd := reader.Text()

		for _, cred := range *CREDS{

			if cred.Username == uname && cred.Password == passwd{
				validated=true
				log.Println("Client Validated")
				break
			}
		}
		if validated{
			conn.Write([]byte("1"))
			break
		}
		conn.Write([]byte("0"))
	}
}

func HandleServer(conn net.Conn){
	defer conn.Close()
	AuthenticateClient(conn)


	buffer := make([]byte, BUFFERSIZE)
	for {
		n,_ := conn.Read(buffer)
		command := string(buffer[:n])
		commandArr := strings.Split(command," ")


		switch strings.ToLower(commandArr[0]) {

		case "upload":
			log.Println("UPLOAD Request")
			n, _ := conn.Read(buffer)
			fileSize,err := strconv.ParseInt(string(buffer[:n]),10, 64)
			log.Println(fileSize)
			if err!=nil || fileSize ==-1{
				log.Println("FILE ERROR")
				continue
			}
			GetFile(conn,commandArr[1], fileSize)


		case "download":
			log.Println("Download")
			SendFile(conn, commandArr[1])


		case "ls":
			log.Println("ls")
			getDIR(conn)


		case "pwd":
			log.Println("pwd")
			conn.Write([]byte(ROOT))


		case "cd":
			changeDIR(conn, commandArr[1])

		case "delete":
			deleteDIR(conn, commandArr[1])


		case "close":
			log.Println("CLOSE")
			return
		}
	}

}