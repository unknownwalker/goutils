package utils

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

//读取key=value类型的配置文件
func InitConfig(path string) map[string]string {
	config := make(map[string]string)

	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		//panic(err)
		fmt.Println("error file:" + path)
		return nil
	}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}
		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])
		if len(value) == 0 {
			continue
		}
		config[key] = value
	}
	return config
}

func PortInUse(port string) bool {
	checkStatement := fmt.Sprintf("lsof -i:%s ", port)
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return true
	}
	return false
}

func GetCPUID() string {
	checkStatement := fmt.Sprintf("dmidecode -t 4 | grep ID ")
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return string(output)
	}
	return ""
}

func GetSYSUUID() string {
	checkStatement := fmt.Sprintf("dmidecode -t system | grep UUID ")
	output, _ := exec.Command("sh", "-c", checkStatement).CombinedOutput()
	if len(output) > 0 {
		return string(output)
	}
	return ""
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
