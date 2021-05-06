package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/jaakali/ipfiter"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedIpFilterServer
}

func (s *server) Rewrite(ctx context.Context, in *pb.IpReq) (*pb.IpRep, error) {
	ip := net.ParseIP(in.Ip4).To4()
	if ip == nil {
		return &pb.IpRep{Ret: true}, nil
	}
	ip2 := binary.BigEndian.Uint32(ip)
	log.Printf("Received: %v %d", in.Ip4, ip2)
	return &pb.IpRep{Ret: bSearch(ipp, ip2)}, nil
}

var ipp []Ipblock

func main() {

	ipp = InitIpp(true)
	//随机端口获取
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println(os.Getpid(), "grpc using", ln.Addr().String())

	s := grpc.NewServer()
	pb.RegisterIpFilterServer(s, &server{})
	go func() {
		if err := s.Serve(ln); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	for {
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.Signal(10))
		switch <-quit {
		case syscall.SIGINT: //Ctrl+c
			goto L
		case syscall.SIGTERM:
			goto L
		case syscall.Signal(10):	
			if ipb := InitIpp(false); ipb == nil {
				log.Println("[SIG]", "reinit nil")
			} else {
				ipp = ipb
				log.Println("[SIG]", "reinit ok")
			}
		} 
	}
L:	
	s.Stop()
	log.Println("server down...")
}

type Ipblock struct {
	S uint32 
	E uint32
}
func InitIpp(b bool) []Ipblock {
	var scan *bufio.Scanner
	var s, e uint32
	var ib Ipblock
	ipp := make([]Ipblock, 0)

	f, err := os.Open("../ipblocks.txt")
	if err != nil {
		goto F
	}
	defer f.Close()
	scan = bufio.NewScanner(f)
	if scan.Scan() {
		if _, err = fmt.Sscanf(scan.Text(), "%d %d", &ib.S, &ib.E); err != nil {
			goto F
		}
	}
	for scan.Scan() {
		if _, err = fmt.Sscanf(scan.Text(), "%d %d", &s, &e); err != nil {
			goto F
		}
		if ib.E + 1 == s {
			ib.E = e
		} else {
			ipp = append(ipp, ib)
			ib = Ipblock{s, e}
		}
	}
	return append(ipp, ib)
F:
	if b {
		log.Fatalln(err)
	}
	log.Println("[InitIpp]", err)
	return nil	
}
func bSearch(ipbs []Ipblock, ip uint32) bool {
	l, r := 0, len(ipbs)
	m := l/2
	for r - l > 1 {	
		if ipbs[m].S > ip { //向左侧查找
			r = m
		} else {
			l = m
		}
		m = (l + r) /2
	}
	if ipbs[m].S <= ip && ipbs[m].E >= ip {
		return true
	}
	return false
}