package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/gopacket"
)

type PacketInfo struct {
	LinkLayer        *LinkLayerInfo        `json:"link_layer,omitempty"`
	NetworkLayer     *NetworkLayerInfo     `json:"network_layer,omitempty"`
	TransportLayer   *TransportLayerInfo   `json:"transport_layer,omitempty"`
	ApplicationLayer *ApplicationLayerInfo `json:"application_layer,omitempty"`
}

type LinkLayerInfo struct {
	Protocol string `json:"protocol"`
	MacSrc   string `json:"mac_src"`
	MacDst   string `json:"mac_dst"`
}

type NetworkLayerInfo struct {
	Protocol string `json:"protocol"`
	IPSrc    string `json:"ip_src"`
	IPDst    string `json:"ip_dst"`
}

type TransportLayerInfo struct {
	Protocol string `json:"protocol"`
	PortSrc  string `json:"port_src"`
	PortDst  string `json:"port_dst"`
}

type ApplicationLayerInfo struct {
	Payload string `json:"payload"`
}

func ParsePacket(packet gopacket.Packet) string {
	var packet_info PacketInfo

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------")

	if link_layer := packet.LinkLayer(); link_layer != nil {

		defer fmt.Println("|\t\t\t\tProtocol: ", link_layer.LinkFlow().EndpointType().String())
		defer fmt.Println("|\t\t\t\tMAC src: ", link_layer.LinkFlow().Src().String())
		defer fmt.Println("|\t\t\t\tMAC dst: ", link_layer.LinkFlow().Dst().String())
		packet_info.LinkLayer = &LinkLayerInfo{
			Protocol: link_layer.LinkFlow().EndpointType().String(),
			MacSrc:   link_layer.LinkFlow().Src().String(),
			MacDst:   link_layer.LinkFlow().Dst().String(),
		}

		if network_layer := packet.NetworkLayer(); network_layer != nil {
			defer fmt.Println("|\t\t\tProtocol: ", network_layer.NetworkFlow().Dst().EndpointType())
			defer fmt.Println("|\t\t\tIP src: ", network_layer.NetworkFlow().Src().String())
			defer fmt.Println("|\t\t\tIP dst: ", network_layer.NetworkFlow().Dst().String())
			packet_info.NetworkLayer = &NetworkLayerInfo{
				Protocol: network_layer.NetworkFlow().Dst().EndpointType().String(),
				IPSrc:    network_layer.NetworkFlow().Src().String(),
				IPDst:    network_layer.NetworkFlow().Dst().String(),
			}

			if transport_layer := packet.TransportLayer(); transport_layer != nil {
				defer fmt.Println("|\t\tProtocol: ", transport_layer.TransportFlow().Dst().EndpointType())
				defer fmt.Println("|\t\tPort src: ", transport_layer.TransportFlow().Src().String())
				defer fmt.Println("|\t\tPort dst: ", transport_layer.TransportFlow().Dst().String())
				packet_info.TransportLayer = &TransportLayerInfo{
					Protocol: transport_layer.TransportFlow().Dst().EndpointType().String(),
					PortSrc:  transport_layer.TransportFlow().Src().String(),
					PortDst:  transport_layer.TransportFlow().Dst().String(),
				}

				if app_layer := packet.ApplicationLayer(); app_layer != nil {
					defer fmt.Println("|\tPayload: ", app_layer.Payload())
					packet_info.ApplicationLayer = &ApplicationLayerInfo{
						Payload: string(app_layer.Payload()),
					}
				}
			}
		}
	}

	fmt.Println("-------------------------------------------------------------------------------------------------------------------------")
	jsonData, err := json.MarshalIndent(packet_info, "", "  ")
	if err != nil {
		panicLog.Panicln("Error marshaling to JSON:", err)
	}
	return string(jsonData)
}
