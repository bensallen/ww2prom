package warewulf

type Node struct {
	Arch                string      `json:"ARCH"`
	Bootloader          string      `json:"BOOTLOADER"`
	Bootlocal           int64       `json:"BOOTLOCAL"`
	Bootstrapid         string      `json:"BOOTSTRAPID"`
	Console             string      `json:"CONSOLE"`
	Domain              string      `json:"DOMAIN"`
	Enabled             string      `json:"ENABLED"`
	Fileids             interface{} `json:"FILEIDS"`
	Fs                  []string    `json:"FS"`
	Groups              []string    `json:"GROUPS"`
	IpmiAutoconfig      string      `json:"IPMI_AUTOCONFIG"`
	IpmiIpaddr          string      `json:"IPMI_IPADDR"`
	IpmiLanchannel      string      `json:"IPMI_LANCHANNEL"`
	IpmiNetmask         string      `json:"IPMI_NETMASK"`
	IpmiPassword        string      `json:"IPMI_PASSWORD"`
	IpmiUID             string      `json:"IPMI_UID"`
	IpmiUsername        string      `json:"IPMI_USERNAME"`
	IpmiVlanid          string      `json:"IPMI_VLANID"`
	Ipxeurl             string      `json:"IPXEURL"`
	Kargs               interface{} `json:"KARGS"`
	Name                []string    `json:"NAME"`
	Netdevs             NetDevs     `json:"NETDEVS"`
	Nodename            string      `json:"NODENAME"`
	Nokmodcopy          string      `json:"NOKMODCOPY"`
	Noruntimeservices   string      `json:"NORUNTIMESERVICES"`
	Postnetdown         int64       `json:"POSTNETDOWN"`
	Postscript          interface{} `json:"POSTSCRIPT"`
	Postshell           int64       `json:"POSTSHELL"`
	Prescript           interface{} `json:"PRESCRIPT"`
	Preshell            int64       `json:"PRESHELL"`
	Selinux             int64       `json:"SELINUX"`
	ValidateVnfs        int64       `json:"VALIDATE_VNFS"`
	Vnfsid              string      `json:"VNFSID"`
	WwipmiLanchannel    string      `json:"WWIPMI_LANCHANNEL"`
	Wwnoruntimeservices string      `json:"WWNORUNTIMESERVICES"`
	Hwaddr              []string    `json:"_HWADDR"`
	Hwprefix            []string    `json:"_HWPREFIX"`
	ID                  string      `json:"_ID"`
	Ipaddr              []string    `json:"_IPADDR"`
	Timestamp           string      `json:"_TIMESTAMP"`
	Type                string      `json:"_TYPE"`
}

type NetDevs struct {
	Array []struct {
		Fqdn     string `json:"FQDN"`
		Hwaddr   string `json:"HWADDR"`
		Hwprefix string `json:"HWPREFIX"`
		Ipaddr   string `json:"IPADDR"`
		Mtu      string `json:"MTU"`
		Name     string `json:"NAME"`
		Netmask  string `json:"NETMASK"`
		Network  string `json:"NETWORK"`
	} `json:"ARRAY"`
}

type VNFS struct {
	Arch      string `json:"ARCH"`
	Checksum  string `json:"CHECKSUM"`
	Chroot    string `json:"CHROOT"`
	Name      string `json:"NAME"`
	Size      string `json:"SIZE"`
	ID        string `json:"_ID"`
	Timestamp string `json:"_TIMESTAMP"`
	Type      string `json:"_TYPE"`
}

type Bootstrap struct {
	Arch      string `json:"ARCH"`
	Checksum  string `json:"CHECKSUM"`
	Name      string `json:"NAME"`
	Size      string `json:"SIZE"`
	ID        string `json:"_ID"`
	Timestamp string `json:"_TIMESTAMP"`
	Type      string `json:"_TYPE"`
}

type File struct {
	Checksum  string `json:"CHECKSUM"`
	Filetype  string `json:"FILETYPE"`
	Format    string `json:"FORMAT"`
	Gid       string `json:"GID"`
	Mode      string `json:"MODE"`
	Name      string `json:"NAME"`
	Origin    string `json:"ORIGIN"`
	Path      string `json:"PATH"`
	Size      string `json:"SIZE"`
	UID       string `json:"UID"`
	ID        string `json:"_ID"`
	Timestamp string `json:"_TIMESTAMP"`
	Type      string `json:"_TYPE"`
}
