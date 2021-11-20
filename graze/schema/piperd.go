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

/*
 */

type PiperDSchemaHandler struct {
}

func (w PiperDSchemaHandler) Name() string {
	return "Piper Debugger (piperd://)"
}

func (w PiperDSchemaHandler) Query(uri string) SchemaQueryResponse {
	grazeQueryResponse := SchemaQueryResponse{}
	grazeQueryResponse.ShouldRedir = false

	parts := strings.Split(uri, "/")
	hostname := parts[0] + ":60"
	if strings.Contains(parts[0], ":") {
		hostname = parts[0]
	}
	path := "/" + strings.Join(parts[1:], "/")
	conn, err := net.Dial("tcp", hostname)
	if err != nil {
		grazeQueryResponse.Status = "Error"
		grazeQueryResponse.StatusColor = 1
		return grazeQueryResponse
	}
	reqBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(reqBytes, uint16(len(path)))
	reqBytes = append(reqBytes, []byte(path)...)
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
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Piper Debug View", "", 2))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("Req Bytes:"), "", 0))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("%v", reqBytes), "", 6))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("Resp Bytes:"), "", 0))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("%v", respBytes), "", 6))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("Resp Content Type:"), "", 0))
	grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine(fmt.Sprintf("0x%02x", contentType), "", 6))

	return grazeQueryResponse
}
