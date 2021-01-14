package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket/layers"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	//deviceName  string = "\\Device\\NPF_Loopback"
	deviceName  string = "eth0" //tether
	snapshotLen int32  = 1024
	promiscuous bool   = true //for all network
	err         error
	timeout     time.Duration = -1 * time.Second
	handle      *pcap.Handle
	packetCount int = 1
)

func main() {

	fmt.Println("Start")

	//Ethernet 3: "\\Device\\NPF_{8E551C5E-FCE9-42F6-8D42-AF66421424E9}"
	//Local Loopback:"\\Device\\NPF_Loopback"
	handle, err = pcap.OpenLive(deviceName, snapshotLen, promiscuous, pcap.BlockForever)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		tcpLayer := packet.Layer(layers.LayerTypeTCP)
		ipv4Layer := packet.Layer(layers.LayerTypeIPv4)

		log.Println("-------------------------------------------------------------")

		if tcpLayer != nil && ipv4Layer != nil {
			log.Printf("'%v' ---> '%v'", ipv4Layer.(*layers.IPv4).SrcIP.String()+":"+tcpLayer.(*layers.TCP).SrcPort.String(), ipv4Layer.(*layers.IPv4).DstIP.String()+":"+tcpLayer.(*layers.TCP).DstPort.String())
		}
	}
}
