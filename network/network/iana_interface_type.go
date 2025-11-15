// network/iana_interface_type.go
package network

import "fmt"

var IanaInterfaceTypeMap = map[uint32]string{
	1:   "Other",
	6:   "Ethernet",
	9:   "Token Ring",
	23:  "PPP (Point-to-Point Protocol)",
	24:  "Loopback",
	28:  "SLIP (Serial Line IP)",
	37:  "ATM",
	53:  "IP over ATM",
	71:  "Wi-Fi (IEEE 802.11)",
	131: "Tunnel (e.g., VPN, IPv6-in-IPv4)",
	144: "Virtual IP Address",
	243: "Cellular (WWAN - e.g., 4G/5G)",
	244: "Cellular (WWAN v2)",
	245: "USB Network Device",
	246: "Bluetooth PAN",
	247: "WiMAX (IEEE 802.16)",
	248: "FireWire (IEEE 1394)",
	249: "InfiniBand",
	250: "Wireless WAN",
	257: "IEEE 802.15.4 (Low-Rate Wireless PAN)",
	265: "DOCSIS Cable Modem",
	266: "Satellite",
	267: "Quantum Network",
}

func GetIanaInterfaceTypeName(t uint32) string {
	if name, ok := IanaInterfaceTypeMap[t]; ok {
		return name
	}
	return fmt.Sprintf("Unknown (%d)", t)
}

func IsWireless(t uint32) bool {
	switch t {
	case 71, 243, 244, 246, 247, 250, 257:
		return true
	default:
		return false
	}
}
