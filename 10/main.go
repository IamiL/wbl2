package telnetUtil

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)


//Для примера
//GET / HTTP/1.1
//Host: golang.org
//
//^Z
func MainTelnet() {
	//создание флагов и аргументов
	timeout:=10
	flag.IntVar(&timeout,"timeout",10,"Время на подключение к серверу")
	// Получение host и port для утилиты
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)
	address:=host+":"+port
	//address:="golang.org:80"
	fmt.Println("Server Address",address)
	//Запуск Сервера
	//StartTCPServer(address)
	//создание подключения
	conn, err := net.DialTimeout("tcp",address,time.Duration(timeout)*time.Second )
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//Обработка отправляемых сообщений
	for true {
		text:=""
		var scanner = bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			text+=scanner.Text()+"\n"
		}
		_,err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}
		buf := make([]byte,1024*1024)
		n,err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Server answer:")
		fmt.Println(string(buf[0:n]))
	}
}