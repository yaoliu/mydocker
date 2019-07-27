//__author__ = "YaoYao"
//Date: 2019-07-26
package network

import (
	"github.com/vishvananda/netlink"
	"net"
)

type NetWork struct {
	Name    string
	IpRange *net.IPNet
	Driver  string
}

type Endpoint struct {
	ID          string           `json:"id"`
	Device      netlink.Veth     `json:"dev"`
	IpAddress   net.IP           `json:"ip"`
	MacAddress  net.HardwareAddr `json:"mac"`
	PortMapping []string         `json:"portmapping"`
	NetWork     *NetWork
}

type Driver interface {
	Name() string
	Create(subent string, name string) (*NetWork, error)
	Delete(network NetWork) error
	Connect(network *NetWork, endpoint *Endpoint) error
	Disconnect(nework *NetWork, endpoint *Endpoint) error
}

func CreateNetWork(driver, subnet, name string) error {
	_, cidr, _ := net.ParseCIDR(subnet)
	gatewayIp, err := ipAllocator

}
