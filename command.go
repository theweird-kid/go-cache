package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Command string

const (
	CMDSet Command = "SET"
	CMDGet Command = "GET"
)

type Message struct {
	Cmd   Command
	Key   []byte
	Value []byte
	TTL   time.Duration
}

func parseMessage(rawCmd []byte) (*Message, error) {
	rawStr := string(rawCmd)
	parts := strings.Split(rawStr, " ")

	if len(parts) == 0 {
		return nil, fmt.Errorf("invalid protocol format")
	}

	msg := &Message{
		Cmd: Command(parts[0]),
		Key: []byte(parts[1]),
	}

	if msg.Cmd == CMDSet {
		if len(parts) != 4 {
			return nil, errors.New("invalid SET commnad")
		}

		ttl, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, errors.New("invalid SET commnad")
		}

		msg.Value = []byte(parts[2])
		msg.TTL = time.Duration(ttl)

		return msg, nil
	}

	if msg.Cmd == CMDGet {
		if len(parts) != 2 {
			return nil, errors.New("invalid GET commnad")
		}

		return msg, nil
	}

	return nil, errors.New("invalid command")
}
