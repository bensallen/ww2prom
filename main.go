package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/imdario/mergo"
	"github.com/integrii/flaggy"
)

var version = "unknown"

var ipmiCmd *flaggy.Subcommand
var ipmiOutput string
var ipmiTargetsOutput string
var ipmiMergeOutput bool
var nodns bool

func init() {
	// Set your program's name and description.  These appear in help output.
	flaggy.SetName("Warewulf to Prometheus")
	flaggy.SetDescription("Helper app to query Warewulf and output configs into various Prometheus config files")

	// create any subcommands and set their parameters
	ipmiCmd = flaggy.NewSubcommand("ipmi")

	ipmiCmd.Description = "Prometheus ipmi_exporter configs"

	ipmiCmd.String(&ipmiOutput, "o", "output", "Path to ipmi_exporter.yml config")
	ipmiCmd.String(&ipmiTargetsOutput, "t", "targets", "Path to ipmi_targets.yml config")
	ipmiCmd.Bool(&nodns, "n", "nodns", "Do not try to do reverse lookups of IP address to host names")
	ipmiCmd.Bool(&ipmiMergeOutput, "m", "merge", "Merge output with existing configuration file")

	flaggy.AttachSubcommand(ipmiCmd, 1)

	// set the version and parse all inputs into variables
	flaggy.SetVersion(version)
	flaggy.Parse()
}

func main() {

	if ipmiCmd.Used {
		ipmiCfg := IPMIExporterConfig{Modules: make(map[string]IPMIModule)}
		if err := ipmiCfg.Update(nodns); err != nil {
			log.Fatalf("Error updating IPMI config from wwsh, %v", err)
		}

		if ipmiOutput == "" {
			if ipmiCfgYml, err := ipmiCfg.ToYaml(); err != nil {
				log.Fatalf("Error marshalling IPMI Config to YAML, %v", err)
			} else {
				fmt.Print(string(ipmiCfgYml))
			}
		} else {

			if ipmiMergeOutput {
				if _, err := os.Stat(ipmiOutput); os.IsNotExist(err) {
					log.Fatalf("Merge output is specified and output file does not exist, %v", err)
				} else {
					var f *os.File

					if f, err = os.Open(ipmiOutput); err != nil {
						log.Fatalf("Error opening existing output file, %v", err)
					}
					defer f.Close()

					currentIpmiCfgYml, err := ioutil.ReadAll(f)
					if err != nil {
						log.Fatalf("Error reading existing output file, %v", err)
					}
					currentIpmiCfg, err := FromYAML(currentIpmiCfgYml)
					if err != nil {
						log.Fatalf("Error unmarshalling existing output file, %v", err)
					}
					// Merge ipmiCfg over currentIpmiCfg
					if err := mergo.Merge(currentIpmiCfg, ipmiCfg); err != nil {
						log.Fatalf("Error merging existing output file with new configuration from wwsh, %v", err)
					}
					ipmiCfg = *currentIpmiCfg
				}
			}

			ipmiCfgYml, err := ipmiCfg.ToYaml()
			if err != nil {
				log.Fatalf("Error marshalling IPMI Config to YAML, %v", err)
			}

			tmpfile, err := ioutil.TempFile(filepath.Dir(ipmiOutput), filepath.Base(ipmiOutput))
			if err != nil {
				log.Fatalf("Error creating temp output file, %v", err)
			}

			if _, err := tmpfile.Write(ipmiCfgYml); err != nil {
				log.Fatalf("Error writing to output file, %v", err)
			}

			if err := os.Rename(tmpfile.Name(), ipmiOutput); err != nil {
				log.Fatalf("Error renaming temp file to output file, %v", err)
			}
		}
	} else {
		flaggy.ShowHelpAndExit("")
	}
}
