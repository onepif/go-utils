package utils

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	kb "github.com/eiannone/keyboard"
)

var GetPtr TGetPtr

type TGetPtr func() *string

//func GetPtr(s string) { return &s }

func GetArch() (*string, error) {
	var uname syscall.Utsname

	pstr := new(string)
	if e := syscall.Uname(&uname); e != nil {
		return nil, e
	} else {
		for _, ix := range uname.Machine {
			*pstr += string(byte(ix))
		}
	}

	return pstr, nil
}

func GetKernelVer() (*string, error) {
	var uname syscall.Utsname

	pstr := new(string)

	if e := syscall.Uname(&uname); e != nil { return new(string), e } else {
		for _, ix := range uname.Release { if ix != '\x00' { *pstr += string(byte(ix)) } else { break } }

		return pstr, nil
	}
}

func GetSizeMem() int {
	d, e := os.ReadFile("/proc/meminfo")
	if e != nil {
//		Logg.error.Printf("reading /proc/meminfo file [%s %v %s]", BROWN, e, RESET)
		return 0
	}

	lines := strings.Split(string(d), "\n")
	for _, line := range lines {
		if strings.Contains(line, "MemTotal") {
			fields := strings.Fields(line)
			mem, _ := strconv.ParseFloat(fields[1], 8)
			return int(math.Round(mem / 1048576))
		}
	}
	return 0
}

func ResolveHostIp() (string, error) {
	netInterfaceAddresses, e := net.InterfaceAddrs()
	if e != nil { return "", e }

	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIp, ok := netInterfaceAddress.(*net.IPNet)
		if ok && !networkIp.IP.IsLoopback() && networkIp.IP.To4() != nil {
			ip := networkIp.IP.String()
			return ip, nil
		}
	}
	return "", errors.New("not fount interfaces")
}

// getCurrentFuncName will return the current function's name.
// It can be used for a better log debug system.(I'm NOT sure.)
// https://gist.github.com/HouLinwei/ => Ctrl+F => golang-get-the-function's-name.go
func GetCurrentFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

// logroll
//
// checking for the presence of the build directory and log files
//
func LogRoll(pfname string) {
	if _, e := os.Stat(pfname); e != nil { return } else {
		if files, e := filepath.Glob(pfname+".*"); e != nil {
			os.Rename(pfname, pfname+".1")
		} else {
			for ix := len(files); ix > 0; ix-- {
				os.Rename(pfname+"."+fmt.Sprintf("%d", ix), pfname+"."+fmt.Sprintf("%d", ix+1))
			}
			os.Rename(pfname, pfname+".1")
		}
	}
}

func Counter(count int) kb.KeyEvent {
	var (
		Quit	= make(chan bool)
		Channel	= make(chan kb.KeyEvent)
	)

	go func() {
		for ix:=count; ix>0; ix-- {
			fmt.Printf("%d..", ix)
			for iy:=0; iy<10; iy++ {
				select {
				case <-Quit:
					fmt.Println("0")
					return
				default:
					time.Sleep(100*time.Millisecond)
				}
			}
		}
		fmt.Println("0")
		kb.Close()
	}()

	go func() {
		KeysEvents, e := kb.GetKeys(10)
		if e != nil { panic(e) }
		event := <-KeysEvents
		Channel <-event
		if event.Rune != '\x00' || event.Key != 0 { Quit <- true; kb.Close() }
		return
	}()

	data := <-Channel

	time.Sleep(150*time.Millisecond)

	close(Quit)
	close(Channel)

	return data
}

// https://github.com/kilnfi/go-utils/blob/master/os/file.go
func CopyFile(sourcePath, destPath string) error {
	if _, e := os.Stat(sourcePath); e != nil { return e }

	source, e := os.Open(sourcePath); if e != nil { return e }
	defer source.Close()

	if e = os.MkdirAll(filepath.Dir(destPath), 0o700); e != nil { return e }

	dest, e := os.Create(destPath); if e != nil { return e }
	defer dest.Close()

	_, e = io.Copy(dest, source)
	return e
}
// go get github.com/otiai10/copy
// это копирует рекурсивно папки

func listDirByWalk(path string) {
	filepath.Walk(path, func(wPath string, info os.FileInfo, e error) error {
 
		// Обход директории без вывода
		if wPath == path { return nil }

		// Если данный путь является директорией, то останавливаем рекурсивный обход 
		// и возвращаем название папки
		if info.IsDir() {
			fmt.Printf("[%s]\n", wPath)
			return filepath.SkipDir
		}

		// Выводится название файла
		if wPath != path { fmt.Println(wPath) }

		return nil
	})
}
