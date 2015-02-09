package main
import "fmt"
import "log"
import "os"
import "golang.org/x/crypto/ssh"
import "io/ioutil"
import "os/user"

func main() {
    if len(os.Args) != 4 {
        log.Fatalf("Usage: %s <user> <host:port> <command>", os.Args[0])
    }

    client, session, err := connectToHost(os.Args[1], os.Args[2])
    if err != nil {
        panic(err)
    }
    out, err := session.CombinedOutput(os.Args[3])
    if err != nil {
        panic(err)
    }
    fmt.Println(string(out))
    client.Close()
}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
    var pass string
    fmt.Print("Password: ")
    fmt.Scanf("%s\n", &pass)

    var auth []ssh.AuthMethod
    if pass != "" {
        auth = []ssh.AuthMethod{ssh.Password(pass)}
    } else {
        fmt.Println("...Using private key")
        key, err := getKeyFile()
        check(err)
        if err != nil {
            return nil, nil, err
        }
        auth = []ssh.AuthMethod{ssh.PublicKeys(key)}
    }

    sshConfig := &ssh.ClientConfig{
        User: user,
        Auth: auth,
    }

    client, err := ssh.Dial("tcp", host, sshConfig)
    check(err)
    if err != nil {
        return nil, nil, err
    }

    session, err := client.NewSession()
    check(err)
    if err != nil {
        client.Close()
        return nil, nil, err
    }

    return client, session, nil
}

func getKeyFile() (key ssh.Signer, err error) {
    usr, _ := user.Current()
    file := usr.HomeDir + "/.ssh/id_rsa"
    buf, err := ioutil.ReadFile(file)
    check(err)
    if err != nil {
        return
    }
    key, err = ssh.ParsePrivateKey(buf)
    check(err)
    if err != nil {
        return
    }
    return
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}