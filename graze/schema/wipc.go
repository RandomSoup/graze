package schema

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strings"

	"luminoso.dev/graze/graze/render"
)

//impements the experimental "webipc" piper-derived protocol
//in this implementation, Graze acts as a webIPC debugger. As such, the URL format is interpreted in an unusual and rather specific manner

/*
about webIPC:
webIPC is a piper-derived protocol for inter-server communication over the same network (or remotely, such as with a discord bot managing a server).
it reallocates and defines the 0x and 1x ranges, and doesn't support redirects.
webIPC adds additional header fields to a request. A webIPC request format looks like the following:
| Num Bytes | Purpose               |
| --------- | --------------------- |
| 2         | Command string len    |
| 2         | Paramstring len       |
| Remaining | Command string, following by Paramstring, in UTF8 |

the standardized delimeter for params is a space.
*/

type WIPCSchemaHandler struct {
}

func (w WIPCSchemaHandler) Name() string {
	return "WebIPC Debugger (use wipc://<target>$<command>?<paramstring>"
}

func (w WIPCSchemaHandler) Query(uri string) SchemaQueryResponse {
	grazeQueryResponse := SchemaQueryResponse{}
	grazeQueryResponse.ShouldRedir = false // webIPC doesn't do redirects.

	parts := strings.Split(uri, "$")
	command := parts[1]
	host := parts[0]
	if !strings.Contains(host, ":") {
		host += ":61"
	}
	parts = strings.Split(parts[1], "?")
	params := parts[1]
	fmt.Printf("Sending command %s to host %s with params %s", command, host, params)

	/*Netcode*/
	conn, err := net.Dial("tcp", host)
	if err != nil {
		grazeQueryResponse.Status = "Error"
		grazeQueryResponse.StatusColor = 1
		return grazeQueryResponse
	}

	reqBytes := make([]byte, 4)
	binary.LittleEndian.PutUint16(reqBytes, uint16(len(command)))
	binary.LittleEndian.PutUint16(reqBytes, uint16(len(params)))
	reqBytes = append(reqBytes, []byte(command)...)
	conn.Write(reqBytes)
	reader := bufio.NewReader(conn)
	respBytes := []byte{}
	gotHeader := false
	contentType := 0
	clen := uint64(0)
	for {
		b, _ := reader.ReadByte()
		respBytes = append(respBytes, b)
		// uint64 is 8 bytes, + contenttype is 1, so once we have 9 bytes we have ourselves a header
		if len(respBytes) >= 9 && !gotHeader {
			contentType = int(respBytes[0])
			clb := respBytes[1:]
			buf := bytes.NewBuffer(clb)
			binary.Read(buf, binary.LittleEndian, &clen)
			gotHeader = true
			fmt.Printf("Clen:  %v | CType: %v\n", clen, contentType)
			fmt.Printf("%v\n", respBytes)
		}
		if len(respBytes) >= 9+int(clen) && gotHeader {
			break
		}
	}
	grazeQueryResponse.Status = "dbgview"
	grazeQueryResponse.StatusColor = 0
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("wIPC Debug View", "", 2))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("Req Bytes:"), "", 0))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("%v", reqBytes), "", 6))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("Resp Bytes:"), "", 0))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("%v", respBytes), "", 6))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("Resp Content Type:"), "", 0))
	if contentType < 0x10 {
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("0x%02x (0X series OK)", contentType), "", 6))
	} else {
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("0x%02x (1X series Error)", contentType), "", 8))
	}
	return grazeQueryResponse
}
