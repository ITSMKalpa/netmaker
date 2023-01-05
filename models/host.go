package models

import (
	"net"

	"github.com/google/uuid"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

// WIREGUARD_INTERFACE name of wireguard interface
const WIREGUARD_INTERFACE = "netmaker"

// Host - represents a host on the network
type Host struct {
	ID               uuid.UUID        `json:"id" yaml:"id"`
	Verbosity        int              `json:"verbosity" yaml:"verbosity"`
	FirewallInUse    string           `json:"firewallinuse" yaml:"firewallinuse"`
	Version          string           `json:"version" yaml:"version"`
	IPForwarding     bool             `json:"ipforwarding" yaml:"ipforwarding"`
	DaemonInstalled  bool             `json:"daemoninstalled" yaml:"daemoninstalled"`
	HostPass         string           `json:"hostpass" yaml:"hostpass"`
	Name             string           `json:"name" yaml:"name"`
	OS               string           `json:"os" yaml:"os"`
	Interface        string           `json:"interface" yaml:"interface"`
	Debug            bool             `json:"debug" yaml:"debug"`
	ListenPort       int              `json:"listenport" yaml:"listenport"`
	LocalAddress     net.IPNet        `json:"localaddress" yaml:"localaddress"`
	LocalRange       net.IPNet        `json:"localrange" yaml:"localrange"`
	LocalListenPort  int              `json:"locallistenport" yaml:"locallistenport"`
	ProxyListenPort  int              `json:"proxy_listen_port" yaml:"proxy_listen_port"`
	MTU              int              `json:"mtu" yaml:"mtu"`
	PublicKey        wgtypes.Key      `json:"publickey" yaml:"publickey"`
	MacAddress       net.HardwareAddr `json:"macaddress" yaml:"macaddress"`
	TrafficKeyPublic []byte           `json:"traffickeypublic" yaml:"trafficekeypublic"`
	InternetGateway  net.UDPAddr      `json:"internetgateway" yaml:"internetgateway"`
	Nodes            []string         `json:"nodes" yaml:"nodes"`
	Interfaces       []Iface          `json:"interfaces" yaml:"interfaces"`
	EndpointIP       net.IP           `json:"endpointip" yaml:"endpointip"`
	ProxyEnabled     bool             `json:"proxy_enabled" yaml:"proxy_enabled"`
	IsDocker         bool             `json:"isdocker" yaml:"isdocker"`
	IsK8S            bool             `json:"isk8s" yaml:"isk8s"`
	IsStatic         bool             `json:"isstatic" yaml:"isstatic"`
	IsDefault        bool             `json:"isdefault" yaml:"isdefault"`
}

// FormatBool converts a boolean to a [yes|no] string
func FormatBool(b bool) string {
	s := "no"
	if b {
		s = "yes"
	}
	return s
}

// ParseBool parses a [yes|no] string to boolean value
func ParseBool(s string) bool {
	b := false
	if s == "yes" {
		b = true
	}
	return b
}