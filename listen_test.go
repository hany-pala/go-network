package ch03

import(
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	
	defer func () { _ = listener.Close() }()
	
	t.Logf("bound to %q", listener.Addr())
	
	// 하나 이상의 수신 연결을 처리하기 위해 for loop
	for {
		//wait to connect
		//Accept() - 클라-서버간의 tcp 핸드셰이크 절차가 완료될때까지 블로킹됨. 
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		
		go func (c net.Conn) {
			// 서버로 FIN 패킷 보냄
			defer c.Close()
		}(conn)
	}
}