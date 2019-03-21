package main

import (
	"net"
	"strings"
	"ww2prom/pkgs/warewulf"

	"gopkg.in/yaml.v2"
)

// IPMIExporterConfig represents the ipmi_exporter configuration
type IPMIExporterConfig struct {
	Modules map[string]IPMIModule `yaml:"modules"`
}

// IPMIModule is a individual module within the ipmi_exporter config
type IPMIModule struct {
	User             string   `yaml:"user"`
	Pass             string   `yaml:"pass"`
	Driver           string   `yaml:"driver,omitempty"`
	Privilege        string   `yaml:"privilege,omitempty"`
	Collectors       []string `yaml:"collectors,omitempty"`
	ExcludeSensorIDs []int64  `yaml:"exclude_sensor_ids,omitempty"`
}

//Update fetches the latest data via wwsh and repopulates the given IPMIExporterConfig
func (ipmiCfg *IPMIExporterConfig) Update(nodns bool) error {
	nodes, err := warewulf.Nodes()
	if err != nil {
		return err
	}

	for _, v := range nodes {
		if v.IpmiIpaddr != "" || v.IpmiUsername != "" || v.IpmiPassword != "" {
			mod := IPMIModule{User: v.IpmiUsername, Pass: v.IpmiPassword}
			if names, err := net.LookupAddr(v.IpmiIpaddr); nodns || err != nil {
				ipmiCfg.Modules[v.IpmiIpaddr] = mod
			} else {
				name := strings.TrimSuffix(names[0], ".")
				ipmiCfg.Modules[name] = mod
			}
		}
	}
	return nil
}

// ToYaml Marshalls IPMIExporterConfig to YAML
func (ipmiCfg *IPMIExporterConfig) ToYaml() ([]byte, error) {
	d, err := yaml.Marshal(&ipmiCfg)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// FromYAML unmarshals a IPMIExporterConfig from YAML
func FromYAML(b []byte) (*IPMIExporterConfig, error) {
	var ipmiCfg IPMIExporterConfig
	if err := yaml.Unmarshal(b, &ipmiCfg); err != nil {
		return nil, err
	}
	return &ipmiCfg, nil
}

// IPMITargetsConfig represents a Prometheus File dy
type IPMITargetsConfig []struct {
	Targets []string `yaml:"targets"`
	Labels  struct {
		Job string `yaml:"job"`
	} `yaml:"labels"`
}
