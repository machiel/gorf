// Package gorf provides a simple interface in Go to communicate with running
// frog instance
package gorf

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Client communicates with the frog server
type Client struct {
	Conn net.Conn
}

// NewClient constructs a Client, by setting up the connection with the frog
// server
func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)

	if err != nil {
		return nil, fmt.Errorf("gorf: Could not connect with address '%s'", addr)
	}

	client := &Client{Conn: conn}
	return client, nil
}

// Parse sends a text to be parsed by the frog server to the frog server and
// parses the response given by the frog server.
func (c Client) Parse(text string) ([]Sentence, error) {
	return parseText(c.Conn, text)
}

func parseText(conn net.Conn, text string) ([]Sentence, error) {
	fmt.Fprintf(conn, text+"\n")
	fmt.Fprintf(conn, "EOT\n")

	reader := bufio.NewReader(conn)
	var sentences []Sentence
	var sentence Sentence

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			return sentences, fmt.Errorf("gorf: Could not read from frog server")
		}

		l := len(line)

		if l == 1 {
			sentences = append(sentences, sentence)
			sentence = Sentence{}
		} else if line == "READY\n" {
			break
		} else {
			token, err := parseLine(line)

			if err != nil {
				return sentences, err
			}

			sentence = append(sentence, token)
		}
	}

	return sentences, nil
}

func parseLine(line string) (Token, error) {
	var t Token
	var err error

	size := len(line)
	line = line[:size-1]
	columns := strings.Split(line, "\t")

	t.Position, err = strconv.Atoi(columns[0])

	if err != nil {
		return t, fmt.Errorf("gorf: Could not parse position")
	}

	t.POSConfidence, err = strconv.ParseFloat(columns[5], 64)

	if err != nil {
		return t, fmt.Errorf("gorf: Could not parse POS Confidence")
	}

	t.Token = columns[1]
	t.Lemma = columns[2]
	t.POSTag = columns[4]
	t.NamedEntityType = columns[6]

	return t, nil
}

// Sentence is a collection of tokens
type Sentence []Token

// Token holds the data as provided by Frog
type Token struct {
	Position        int
	Token           string
	Lemma           string
	POSTag          string
	POSConfidence   float64
	NamedEntityType string
}
