package v0

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"time"
)

type FileTransfer struct {
	io.Reader
	total    int64
	fileSize int64
}

type File struct {
	uuid string
}

func (ft *FileTransfer) Read(p []byte) (int, error) {
	n, err := ft.Reader.Read(p)
	ft.total += int64(n)
	return n, err
}

func sftpRemoteToLocal(srcRemotePath, dstLocalPath string) {
	config := &ssh.ClientConfig{
		User:            "kitak",
		HostKeyCallback: nil,
		Auth: []ssh.AuthMethod{
			ssh.Password("PASSWORD"),
		},
	}
	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", "example.com:22", config)
	if err != nil {
		panic(err)
	}
	defer sshConn.Close()

	// create new SFTP client
	client, err := sftp.NewClient(sshConn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// create destination file
	dstFile, err := os.Create(dstLocalPath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// open source file
	srcFile, err := client.Open(srcRemotePath)
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	// flush in-memory copy
	err = dstFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}

func sftpLocalToRemote(srcLocalPath, dstRemotePath string) {
	config := &ssh.ClientConfig{
		User:            "kitak",
		HostKeyCallback: nil,
		Auth: []ssh.AuthMethod{
			ssh.Password("PASSWORD"),
		},
	}
	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", "example.com:22", config)
	if err != nil {
		panic(err)
	}
	defer sshConn.Close()

	// create new SFTP client
	client, err := sftp.NewClient(sshConn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// create destination file
	dstFile, err := client.Create(dstRemotePath)
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open(srcLocalPath)
	if err != nil {
		log.Fatal(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)
}

func fileCopy(srcFile, dstFile, controller string) error {
	log.Println("---Copy disk image")
	log.Println("src: " + srcFile)
	log.Println("dst: " + dstFile)
	src, err := os.Open(srcFile)
	if err != nil {
		log.Println("Error: open error")
		return fmt.Errorf("open error")
	}
	defer src.Close()
	file, err := src.Stat()
	if err != nil {
		log.Println("Error: file gateway error")
		return err
	}

	dst, err := os.Create(dstFile)
	if err != nil {
		log.Println("Error: file create")
		return err
	}
	defer dst.Close()

	done := make(chan bool)
	tmp := FileTransfer{Reader: src, fileSize: file.Size()}
	go func() {
		select {
		case <-done:
			return
		default:
			if tmp.fileSize != tmp.total {
				<-time.NewTimer(2 * time.Second).C
				sendServer()
			} else {
				return
			}
		}
	}()

	_, err = io.Copy(dst, &tmp)
	if err != nil {
		log.Println("Error: file copy error")
		return err
	}
	<-done

	return nil
}

func sendServer() {

}
