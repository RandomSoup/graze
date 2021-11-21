package schema

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	"luminoso.dev/graze/graze/render"
)

const PIPER_DEFAULT_PORT = 60

/*
piper.go
implements the piper:// schema handler (and therefore the piper protocol)
*/

type PiperSchemaHandler struct {
}

func (p PiperSchemaHandler) Name() string {
	return "piper:// schema handler"
}

func (p PiperSchemaHandler) Query(uri string) SchemaQueryResponse {
	grazeQueryResponse := SchemaQueryResponse{}

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
	switch contentType {
	case 0x0:
		grazeQueryResponse.Status = "Ok/Txt"
		grazeQueryResponse.StatusColor = 0
		grazeQueryResponse.Contents = parsePureText(string(respBytes[9:]))
	case 0x1:
		grazeQueryResponse.Status = "Ok/Gmi"
		grazeQueryResponse.StatusColor = 0
		grazeQueryResponse.Contents = parseGemtext(string(respBytes[9:]))
	case 0x10:
		//rdl = Raw DownLoad
		grazeQueryResponse.Status = "Ok/Rdl"
		grazeQueryResponse.StatusColor = 0
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("Saving file to downloads/", "", 0))
		ioutil.WriteFile(strings.Split(path, "/")[len(strings.Split(path, "/"))-1], respBytes[9:], 0667)

	case 0x22:
		grazeQueryResponse.Status = "Error"
		grazeQueryResponse.StatusColor = 1
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("0x22 Resource Not Found", "", 0))
	case 0x23:
		grazeQueryResponse.Status = "Error"
		grazeQueryResponse.StatusColor = 1
		grazeQueryResponse.Contents = append(grazeQueryResponse.Contents, render.NewRendLine("0x23 Internal Server Error", "", 0))
	}
	return grazeQueryResponse
}
