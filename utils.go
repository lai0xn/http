package http

import (
	"fmt"
	"strconv"
)

func constructResponse(res *Response) string {
	res.Headers["Content-Length"] = strconv.Itoa(len(res.Body))

	// Construct the resonse
	response := fmt.Sprintf("%s %d OK\r\n", res.Proto, res.Status)
	for key, value := range res.Headers {
		response += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	response += "\r\n" + res.Body
	return response

}
