package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pkg/pool/goroutine"
	"log"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// 不带阻塞逻辑

type echoServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (es *echoServer) React(
	frame []byte,
	c gnet.Conn,
) (
	out []byte,
	action gnet.Action,
) {
	data := append([]byte{'a'}, frame...)
	fmt.Println(string(data))

	_ = es.pool.Submit(func() {
		time.Sleep(time.Second * 1)
		// 异步写入数据，减少阻塞，并赋值给out变量并返回
		c.AsyncWrite(data)
	})
	return
}

func (es *echoServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("Echo server is listening on %s (multi-cores: %t, loops: %d)\n",
		srv.Addr.String(), srv.Multicore, srv.NumEventLoop)
	return
}

func main() {
	p := goroutine.Default()
	defer p.Release()

	echo := &echoServer{}
	echo.pool = goroutine.Default()
	log.Fatal(gnet.Serve(echo, "tcp://:9001", gnet.WithMulticore(true)))
}

var (
	res         string
	errMsg      = "Internal Server Error"
	errMsgBytes = []byte(errMsg)
)

type request struct {
	proto, method string
	path, query   string
	head, body    string
	remoteAddr    string
}

type httpServer struct {
	*gnet.EventServer
}

type httpCodec struct {
	req request
}

func (hc *httpCodec) Encode(c gnet.Conn, buf []byte) (out []byte, err error) {
	if c.Context() == nil {
		return buf, nil
	}
	return appendResp(out, "500 Error", "", errMsg+"\n"), nil
}

func (hc *httpCodec) Decode(c gnet.Conn) (out []byte, err error) {
	buf := c.Read()
	c.ResetBuffer()
	var leftover []byte
pipeline:
	leftover, err = parseReq(buf, &hc.req)
	if err != nil {
		c.SetContext(err)
		return nil, err
	} else if len(leftover) == len(buf) {
		return
	}
	out = appendHandle(out, res)
	buf = leftover
	goto pipeline
}

func appendResp(b []byte, status, head, body string) []byte {
	b = append(b, "HTTP/1.1"...)
	b = append(b, ' ')
	b = append(b, status...)
	b = append(b, '\r', '\n')
	b = append(b, "Server: gnet\r\n"...)
	b = append(b, "Date: "...)
	b = time.Now().AppendFormat(b, "Mon, 02 Jan 2006 15:04:05 GMT")
	b = append(b, '\r', '\n')
	if len(body) > 0 {
		b = append(b, "Content-Length: "...)
		b = strconv.AppendInt(b, int64(len(body)), 10)
		b = append(b, '\r', '\n')
	}
	b = append(b, head...)
	b = append(b, '\r', '\n')
	if len(body) > 0 {
		b = append(b, body...)
	}
	return b
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func parseReq(data []byte, req *request) (leftover []byte, err error) {
	sdata := b2s(data)
	var i, s int
	var head string
	var clen int
	var q = -1
	// method, path, proto line
	for ; i < len(sdata); i++ {
		if sdata[i] == ' ' {
			req.method = sdata[s:i]
			for i, s = i+1, i+1; i < len(sdata); i++ {
				if sdata[i] == '?' && q == -1 {
					q = i - s
				} else if sdata[i] == ' ' {
					if q != -1 {
						req.path = sdata[s:q]
						req.query = req.path[q+1 : i]
					} else {
						req.path = sdata[s:i]
					}
					for i, s = i+1, i+1; i < len(sdata); i++ {
						if sdata[i] == '\n' && sdata[i-1] == '\r' {
							req.proto = sdata[s:i]
							i, s = i+1, i+1
							break
						}
					}
					break
				}
			}
			break
		}
	}
	if req.proto == "" {
		return data, fmt.Errorf("malformed request")
	}
	head = sdata[:s]
	for ; i < len(sdata); i++ {
		if i > 1 && sdata[i] == '\n' && sdata[i-1] == '\r' {
			line := sdata[s : i-1]
			s = i + 1
			if line == "" {
				req.head = sdata[len(head)+2 : i+1]
				i++
				if clen > 0 {
					if len(sdata[i:]) < clen {
						break
					}
					req.body = sdata[i : i+clen]
					i += clen
				}
				return data[i:], nil
			}
			if strings.HasPrefix(line, "Content-Length:") {
				n, err := strconv.ParseInt(strings.TrimSpace(line[len("Content-Length:"):]), 10, 64)
				if err == nil {
					clen = int(n)
				}
			}
		}
	}
	// not enough data
	return data, nil
}

func appendHandle(b []byte, res string) []byte {
	return appendResp(b, "200 OK", "", res)
}
