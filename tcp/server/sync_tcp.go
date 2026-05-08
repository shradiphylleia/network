package server

import(
	"net"
	"strconv"	
	"io"
	"log"
	"github.com/shradiphylleia/network/tcp/config"

)


func readCommand( c net.Conn)(string, error){
	var buf []byte =make([]byte,512)
	n,err := c.Read(buf[:])
	if err!= nil{
		return "",err
	}
	return string(buf[:n]),nil
}

func respond(cmd string, c net.Conn) error{
	if _,err := c.Write([]byte(cmd)); err!=nil{
		return err
	}
	return nil
}



func RunSyncTCP(cfg config.Config){
	log.Println("starting a sync tcp server on", cfg.Host, cfg.Port)

	var concurrent_clients int=0

	// listen to conf host:port
	 lsnr, err := net.Listen("tcp", cfg.Host+":"+strconv.Itoa(cfg.Port))
	 if err!=nil{
		panic(err)
	 }

	 for{
		c, err := lsnr.Accept()
		if err!=nil{
			panic(err)
		}

		concurrent_clients+=1
		log.Println("client connected to address:", c.RemoteAddr(), "concurrent client", concurrent_clients)


		for{
			cmd, err := readCommand(c)
			if err!=nil{
				c.Close()
				concurrent_clients -=1
				log.Println("client is now disconnected", c.RemoteAddr(), "concurrent clients", concurrent_clients)
				if err == io.EOF{
					break
				}

				log.Println("error", err)
			}
			log.Println("command", cmd)
			if err = respond(cmd, c); err !=nil{
				log.Print("error", err)
			}

		}
	 }

}