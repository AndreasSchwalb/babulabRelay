package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

func getEnv(name string) string {
	env := os.Getenv(name)
	if env == "" {
		slog.Error(fmt.Sprintf("Environment variable not set: %s", name))
		os.Exit(1)
	}
	return env
}

func main() {

	serverAddress := getEnv("SERVER_ADDRESS")
	destinationAddress := getEnv("DESTINATION_ADDRESS")

	udpClient, err := net.Dial("udp", destinationAddress)
	if err != nil {
		panic(err)
	}
	defer udpClient.Close()

	udpServer, err := net.ListenPacket("udp4", serverAddress)
	if err != nil {
		panic(err)
	}
	defer udpServer.Close()

	for {

		buf := make([]byte, 512)
		_, addr, err := udpServer.ReadFrom(buf)

		if err != nil {
			panic(err)
		}

		slog.Info("got message", "source", addr)

		_, err = udpClient.Write(buf)
		if err != nil {
			slog.Error(fmt.Sprintf("Could not sent package to: %s", destinationAddress), "error", err)
			continue
		}
		slog.Info("send message", "destination", destinationAddress)
	}
}
