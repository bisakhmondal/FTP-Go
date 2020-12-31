## FTP CLIENT-SERVER simulation in GoLang
A File Transfer Protocol(FTP) **Multi Client** Simulation in go environment.

### Supported Command
- **pwd** [print working directory]
- **ls** [list directory]
- **cd** [change directory]
- **upload** [upload file from current "filestore/clientDir"]
- **download** [download file from $(pwd) of the ftp client terminal]
-  **delete** [delete specific file from the current pwd]
-  **close/exit** [close the connection]


### Directory Structure
- [x] client : Contains implementation of the client.
- [x] server : server side code in the FTP

```bash
cd server 
go run server.go utils.go
#Multiple clients can be attached 
cd ../client
go run client.go utils.go
```

#### Add credential in read time in server/credential.json

easy- peasy right!

~Bisakh.