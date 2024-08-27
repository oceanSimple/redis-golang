package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"net"
	"server-v5-refactor-server/insMap"
	"server-v5-refactor-server/persistence"
	"server-v5-refactor-server/persistence/aof"
	"server-v5-refactor-server/static/output"
	"server-v5-refactor-server/static/structure"
	"strconv"
	"strings"
	"time"
)

var (
	port = strconv.Itoa(viper.Get("server.port").(int))
)

func main() {
	// create a listener for a port
	listener := listenPort()

	// load aof file
	aofLoad()
	// according to the aof tactic, execute aof persistence method
	aofTactic()

	// accept connections and handle them
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(output.Error(), "Accept error: "+err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

// create a listener for a port
func listenPort() net.Listener {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println(output.Error(), "Listen error: "+err.Error())
		panic(err)
	}
	fmt.Println(output.Info(), "Listening on port:", port)
	return listener
}

// handle a connection
func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(output.Error(), "Error closing connection: "+err.Error())
		}
	}(conn)

	fmt.Println(output.Default()+"Connected to:", conn.RemoteAddr())

	// read incoming message
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(output.Error(), "Read error: "+err.Error())
			break
		}
		// fmt.Printf("Message from client: %s", message)

		// Handle instruction
		// 1. convert message to instruction
		instruction, err := convertMessageToInstruction(message)
		if err != nil {
			_, err = conn.Write([]byte("error: " + err.Error() + "\n"))
			continue
		}
		// 2. execute instruction
		responses := instruction.Execute(instruction)

		// 3. aof log
		aof.WriteToAof(instruction)

		// send response
		_, err = conn.Write([]byte(strings.Join(responses, " ") + "\n"))
	}
}

// convert message to instruction
func convertMessageToInstruction(message string) (*structure.Instruction, error) {
	// 1. split message
	message = strings.TrimSpace(message) // remove leading and trailing white spaces
	splits := strings.Split(message, " ")
	// 2. according to the first word, get the instruction class from map
	if value, ok := insMap.InstructionMap[splits[0]]; !ok { // if not found
		return nil, fmt.Errorf("instruction not found: %s", splits[0])
	} else {
		// 3. fill key and value into the instruction
		value.Values = splits[1:]
		return value, nil
	}
}

func aofTactic() {
	fmt.Println(output.Info(), "Aof tactic:", aof.GetTactic())
	// create a routine to flush buffer every second
	if aof.GetTactic() == "everySecond" {
		go func() {
			for {
				aof.FlushBuffer()
				<-time.After(time.Second)
			}
		}()
	}
}

func aofLoad() {
	aofFileConnector, _ := persistence.GetFileConnector("aof.txt")
	scanner := bufio.NewScanner(aofFileConnector)
	for scanner.Scan() {
		line := scanner.Text()
		instruction, err := convertMessageToInstruction(line)
		if err != nil {
			fmt.Println(output.Error() + "error aof record: " + line)
			continue
		}
		instruction.Execute(instruction)
	}
	fmt.Println(output.Info() + "aof load finish")
}
