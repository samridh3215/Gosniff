package main

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var inteface_choice int
var promisc_mode bool = false

var packet_channel chan gopacket.Packet = make(chan gopacket.Packet, 1000)

func PrintDevs(devs []pcap.Interface) {
	for idx, iDev := range devs {
		fmt.Print("Device ", idx, "\t")
		fmt.Println(iDev.Name)
	}
}
func Sniff() {

	infoLog.Println("Started Sniff()")
	var devs []pcap.Interface
	devs, err := pcap.FindAllDevs()

	if err != nil {
		panicLog.Panicln("No device interface found")
	} else {
		infoLog.Println("Found devices to lisen on\n", devs)
		PrintDevs(devs)
	}

	fmt.Println("Enter interface number to proceed: ")
	fmt.Scanln(&inteface_choice)
	fmt.Println("Chose ", inteface_choice)

	handler, err := pcap.OpenLive(devs[inteface_choice].Name, 2000, promisc_mode, pcap.BlockForever)
	if err != nil {
		panicLog.Panicln(err)
	} else {
		infoLog.Println("Starting sniffing on ", devs[inteface_choice].Name)
		packetSource := gopacket.NewPacketSource(handler, handler.LinkType())
		for packets := range packetSource.Packets() {
			// fmt.Println(packets)
			packet_channel <- packets
		}
	}

}
