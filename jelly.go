package main


// import (
// 	// "bufio"
// 	"log"
// 	"os"
// 	// "path/filepath"
// 	// "strings"
// 	"golang.org/x/crypto/ssh"
// 	"time"
// )
// func main() {
// 	host := "192.168.32.133"
// 	port := "22"
// 	user := "root"
// 	pass := "123456"
// 	cmd  := "ps"
// 	// get host public key
// 	// hostKey := getHostKey(host)
// 	// ssh client config
// 	config := &ssh.ClientConfig{
// 		User: user,
// 		Auth: []ssh.AuthMethod{
// 			ssh.Password(pass),
// 		},
// 		// allow any host key to be used (non-prod)
// 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// 		// verify host public key
// 		// HostKeyCallback: ssh.FixedHostKey(hostKey),
// 		// optional host key algo list
// 		HostKeyAlgorithms: []string{
// 			ssh.KeyAlgoRSA,
// 			ssh.KeyAlgoDSA,
// 			ssh.KeyAlgoECDSA256,
// 			ssh.KeyAlgoECDSA384,
// 			ssh.KeyAlgoECDSA521,
// 			ssh.KeyAlgoED25519,
// 		},
// 		// optional tcp connect timeout
// 		Timeout:         5 * time.Second,
// 	}
// 	// connect
// 	client, err := ssh.Dial("tcp", host+":"+port, config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Close()
// 	// start session
// 	sess, err := client.NewSession()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer sess.Close()
// 	// Uncomment to store output in variable
// 	//var b bytes.Buffer
// 	//sess.Stdout = &b
// 	//sess.Stderr = &b
// 	// setup standard out and error
// 	// uses writer interface
// 	sess.Stdout = os.Stdout
// 	sess.Stderr = os.Stderr
// 	// run single command
// 	err = sess.Run(cmd)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = sess.Run("pwd")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
// import (
//     "bufio"
//     "fmt"
//     "log"
//     "os"
//     "path/filepath"
//     "regexp"
//     "strings"
//     "time"
//     "github.com/golang/glog"
//     "github.com/google/goexpect"
//     "github.com/google/goterm/term"
//     "golang.org/x/crypto/ssh"
// )
// const (
//     timeout = 10 * time.Second
// )
// func main() {
//     host := "192.168.32.133"
//     port := "22"
//     user := "root"
//     pass := "123456"
//     cmd1 := "ls -lh"
//     cmd2 := "ip a"
//     promptRE := regexp.MustCompile("\\$")
//     // get host public key
//     // hostKey := getHostKey(host)
//     sshClt, err := ssh.Dial("tcp", host+":"+port, &ssh.ClientConfig{
//         User: user,
//         Auth: []ssh.AuthMethod{ssh.Password(pass)},
//         // allow any host key to be used (non-prod)
//         HostKeyCallback: ssh.InsecureIgnoreHostKey(),
//         // verify host public key
//         // HostKeyCallback: ssh.FixedHostKey(hostKey),
//     })
//     fmt.Println("3333")
//     if err != nil {
//         glog.Exitf("ssh.Dial(%q) failed: %v", host, err)
//     }
//     defer sshClt.Close()
//     e, _, err := expect.SpawnSSH(sshClt, timeout)
//     if err != nil {
//         glog.Exit(err)
//     }
//     defer e.Close()
//     e.Expect(promptRE, timeout)
//     e.Send(cmd1 + "\n")
//     result1, _, _ := e.Expect(promptRE, timeout)
//     e.Send(cmd2 + "\n")
//     result2, _, _ := e.Expect(promptRE, timeout)
//     e.Send("exit\n")
//     fmt.Println(term.Greenf("Done!\n"))
//     fmt.Printf("%s: result:\n %s\n\n", cmd1, result1)
//     fmt.Printf("%s: result:\n %s\n\n", cmd2, result2)
// }
// func getHostKey(host string) ssh.PublicKey {
//     // parse OpenSSH known_hosts file
//     // ssh or use ssh-keyscan to pull key
//     file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
//     if err != nil {
//         log.Fatal(err)
//     }
//     defer file.Close()
//     scanner := bufio.NewScanner(file)
//     var hostKey ssh.PublicKey
//     for scanner.Scan() {
//         fields := strings.Split(scanner.Text(), " ")
//         if len(fields) != 3 {
//             continue
//         }
//         if strings.Contains(fields[0], host) {
//             var err error
//             hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
//             if err != nil {
//                 log.Fatalf("error parsing %q: %v", fields[2], err)
//             }
//             break
//         }
//     }
//     if hostKey == nil {
//         log.Fatalf("no hostkey found for %s", host)
//     }
//     return hostKey
// }

import "embed"
import "fmt"
import "log"
import "net/http"


//go:embed static
var content embed.FS
func main() {
	// fmt.Println("Hello, 世界")
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.FS(content))))
	http.Handle("/static/", http.FileServer(http.FS(content)))
    // http.HandleFunc("/d5033c97b87fec3d5fab7341a3a4c88098a1989256c52e142fe2f0ad757e25978b81cd345e8ed8a3a66d1a32409cfcbb", check_dir_handler)
    // http.HandleFunc("/upload", upload)
    // log.Fatal(http.ListenAndServe(pts, http.FileServer(http.Dir(dr))))
    // http.Handle("/", http.FileServer(http.Dir(dr)))
    http.Handle("/", http.RedirectHandler("/static/", 301))
    // srv := &http.Server{
    //     Addr:           pts,
    //     IdleTimeout:    8 * time.Second,
    //     MaxHeaderBytes: 1 << 20,
    // }
    // log.Fatal(srv.ListenAndServe())
    log.Fatal(http.ListenAndServe(":8080", nil))
}
