package response
import(
	"github.com/shradiphylleia/network/headers"
	"strconv"
	"bytes"
	"errors"
)

type StatusCode int

const(
	StatusOK  StatusCode=200
	StatusBadRequest StatusCode=400
	StatusInternalServerError StatusCode=500
)

type writerState int

const (
	writerStateStatusLine writerState = iota
	writerStateHeaders
	writerStateBody
)

type Writer struct {
	StatusCode StatusCode
	Headers    headers.Headers
	Body       bytes.Buffer
	state writerState
}

func NewWriter() *Writer {
	return &Writer{
		StatusCode: StatusOK,
		Headers:    GetDefaultHeaders(0),
		state:      writerStateStatusLine,
	}
}

// moving to another implementation now:
// func WriteStatusLine( w io.Writer, statusCode StatusCode) error{
// 	var reason string
// 	switch statusCode {
// 	case StatusOK:
// 		reason="OK"

// 	case StatusBadRequest:
// 		reason="Bad Request"

// 	case StatusInternalServerError:
// 		reason="Internal Server Error"

// 	default:
// 		return fmt.Errorf("unsupported status code: %d", statusCode)
// 	}

// 	_, err:=fmt.Fprintf(
// 		w,
// 		"HTTP/1.1 %d %s\r\n",
// 		statusCode,
// 		reason,
// 	)

// 	return err
// }


func GetDefaultHeaders(contentLen int) headers.Headers{
	h:=headers.NewHeaders()

	h["content-length"]=strconv.Itoa(contentLen)
	h["connection"]="close"
	h["content-type"]="text/plain"
	return h
}

// moving to another implementation now:
// func WriteHeaders (w io.Writer, headers headers.Headers) error{
// for key,value:=range headers {
// 		_, err:=fmt.Fprintf(
// 			w,
// 			"%s: %s\r\n",
// 			key,
// 			value,
// 		)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (w *Writer) WriteStatusLine(statusCode StatusCode) error {
	if w.state != writerStateStatusLine {
		return errors.New("status line already written")
	}

	w.StatusCode = statusCode
	w.state = writerStateHeaders

	return nil
}

func (w *Writer) WriteHeaders(h headers.Headers) error {
	if w.state!=writerStateHeaders {
		return errors.New("headers must be written after status line")
	}
	for k,v:=range h {
		w.Headers[k]=v
	}
	w.state=writerStateBody
	return nil
}

func (w *Writer) WriteBody(p []byte) (int, error) {
	if w.state!=writerStateBody{
		return 0, errors.New("body must be written after headers")
	}
	return w.Body.Write(p)
}