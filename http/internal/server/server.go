package server
import(
	"fmt"
	"net"
	"sync/atomic"
	"github.com/shradiphylleia/network/internal/response"
	"github.com/shradiphylleia/network/internal/request"
	"strconv"
)

type Server struct {
	listener net.Listener
	handler  Handler
	closed   atomic.Bool
}
// type HandlerError struct {
// 	StatusCode response.StatusCode
// 	Message    string
// }

type Handler func(w *response.Writer, req *request.Request)

func Serve(port int, handler Handler) (*Server, error) {
	ln, err:=net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err!=nil {
		return nil, err
	}
	s:=&Server{
		listener: ln,
		handler: handler,
	}
	go s.listen()
	return s, nil
}

func (s *Server) Close() error {
	s.closed.Store(true)
	return s.listener.Close()
}

func (s *Server) listen(){
for {
		conn, err:=s.listener.Accept()
		if err!=nil {
			if s.closed.Load() {
				return
			}
			continue
		}
		go s.handle(conn)
	}
}

func (s *Server) handle(conn net.Conn) {
	defer conn.Close()
	
	resp:=response.NewWriter()
	req, err:=request.RequestFromReader(conn)
	if err!=nil {
	resp.WriteStatusLine(response.StatusBadRequest)
	resp.WriteBody([]byte(err.Error()))
	} else {
	s.handler(resp, req)
	}

	resp.Headers["content-length"]=strconv.Itoa(resp.Body.Len())
	fmt.Fprintf(conn,"HTTP/1.1 %d\r\n",resp.StatusCode,)

	for key, value := range resp.Headers {
		fmt.Fprintf(conn, "%s: %s\r\n", key, value)
	}

	conn.Write([]byte("\r\n"))
	conn.Write(resp.Body.Bytes())
	
}
