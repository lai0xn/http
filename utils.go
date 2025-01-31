package http

import (
	"errors"
	"fmt"
	"strconv"
)

func constructResponse(res *Response) (string, error) {
	res.Headers["Content-Length"] = strconv.Itoa(len(res.Body))

	// Construct the resonse

	phrase, ok := StatusPhrases[res.Status]

	if !ok {
		return "", errors.New("Invalid status code")
	}

	response := fmt.Sprintf("%s %d %s\r\n", res.Proto, res.Status, phrase)
	for key, value := range res.Headers {
		response += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	response += "\r\n" + res.Body
	return response, nil

}
