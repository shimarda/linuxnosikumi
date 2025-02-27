package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func main() {
	pid := os.Getpid()
	fmt.Println("*** testfileのメモリマップ前のプロセスの仮想アドレス空間 ***")
	cmd := exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/maps")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal("catの実行に失敗")
	}
	
	file, err := os.OpenFile("testfile", os.O_RDWR, 0)
	if err != nil {
		log.Fatal("testfileを開けない")
	}
	defer file.Close()

	data, err := syscall.Mmap(int(file.Fd()), 0, 5, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		log.Fatal("mmap()に失敗")
	}
	fmt.Println("")
	fmt.Println("testfileをマップしたアドレス: %p\n", &data[0])
	fmt.Println("")

	fmt.Println("*** testfileのメモリマップ後のプロセスの仮想アドレス空間 ***")
	cmd = exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/maps")
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		log.Fatal("catの実行に失敗")
	}

	replaceBytes := []byte("HELLO")
	for i, _ := range data {
		data[i] = replaceBytes[i]
	}
}