package v2

import (
	"fmt"
	"github.com/pkg/sftp"
	controllerInt "github.com/vmmgr/imacon/pkg/api/core/controller"
	controller "github.com/vmmgr/imacon/pkg/api/core/controller/v2"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type Progress struct {
	total int64
	size  int64
}

func (p *Progress) Write(data []byte) (int, error) {
	n := len(data)
	p.size += int64(n)

	return n, nil
}

func Copy(uuid, url, srcPath, dstPath, addr, user, pk string) error {
	//config := &ssh.ClientConfig{User: auth.User, HostKeyCallback: nil, Auth: []ssh.AuthMethod{ssh.Password(auth.Pass)}}
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{PublicKeyFile(pk)},
	}
	config.SetDefaults()
	sshConn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(err)
	}
	defer sshConn.Close()

	// SFTP Client
	client, err := sftp.NewClient(sshConn)
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	dstPathArray := strings.Split(dstPath, "/")
	err = client.MkdirAll(dstPath[:len(dstPath)-(len(dstPathArray))])
	if err != nil {
		log.Println(err)
	}

	// dstFileの作成
	dstFile, err := client.Create(dstPath)
	if err != nil {
		log.Println(err)
	}
	defer dstFile.Close()

	srcFile, err := os.Open(srcPath)
	if err != nil {
		log.Println(err)
	}

	file, err := srcFile.Stat()
	if err != nil {
		return err
	}

	p := Progress{total: file.Size()}
	log.Println(file.Size())

	go func() {
		for {
			if p.size != p.total {
				<-time.NewTimer(2 * time.Second).C
				log.Println(p.size)
				controller.SendController(url, controllerInt.Controller{
					UUID:     uuid,
					Progress: uint(float64(p.size) / float64(p.total) * 100),
					Finish:   false,
				})
				//bar.Set(int(float64(p.size) / float64(p.total) * 100))
			} else {
				return
			}
		}
	}()

	bytes, err := io.Copy(dstFile, io.TeeReader(srcFile, &p))
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%d bytes copied\n", bytes)

	go func() {
		for {
			<-time.NewTimer(200 * time.Microsecond).C
			err = controller.SendController(url, controllerInt.Controller{
				UUID:     uuid,
				Progress: 100,
				Finish:   true,
			})
			if err == nil {
				break
			}
		}
	}()

	return nil
}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}
