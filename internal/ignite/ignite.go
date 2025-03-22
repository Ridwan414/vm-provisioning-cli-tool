package ignite

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v2"
)

// VMConfig represents the configuration for a VM
type VMConfig struct {
	Name      string
	UID       string
	CPUs      int
	Memory    string
	DiskSize  string
	ImageOCI  string
	EnableSSH bool
}

// Manifest represents the Ignite VM manifest structure
type Manifest struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name string `yaml:"name"`
		UID  string `yaml:"uid"`
	} `yaml:"metadata"`
	Spec struct {
		Image     map[string]string `yaml:"image"`
		CPUs      int               `yaml:"cpus"`
		Memory    string            `yaml:"memory"`
		DiskSize  string            `yaml:"diskSize"`
		SSH       bool              `yaml:"ssh"`
		CopyFiles []struct {
			HostPath string `yaml:"hostPath"`
			VMPath   string `yaml:"vmPath"`
		} `yaml:"copyFiles,omitempty"`
	} `yaml:"spec"`
}

// LaunchVM creates and starts a new VM using Ignite
func LaunchVM(config *VMConfig) (string, error) {
	// Create manifest
	manifest := createManifest(config)

	// Create temporary manifest file
	manifestFile, err := createTempManifestFile(manifest)
	if err != nil {
		return "", fmt.Errorf("failed to create manifest file: %v", err)
	}
	defer os.Remove(manifestFile)

	// Run ignite command
	if err := runIgniteCommand("run", "--config", manifestFile); err != nil {
		return "", fmt.Errorf("failed to run ignite: %v", err)
	}

	// Get VM IP
	ip, err := getVMIP(config.Name)
	if err != nil {
		return "", fmt.Errorf("failed to get VM IP: %v", err)
	}

	return ip, nil
}

// createManifest creates an Ignite manifest from the VM configuration
func createManifest(config *VMConfig) *Manifest {
	manifest := &Manifest{
		APIVersion: "ignite.weave.works/v1alpha4",
		Kind:       "VM",
	}
	manifest.Metadata.Name = config.Name
	manifest.Metadata.UID = config.UID
	manifest.Spec.Image = map[string]string{"oci": config.ImageOCI}
	manifest.Spec.CPUs = config.CPUs
	manifest.Spec.Memory = config.Memory
	manifest.Spec.DiskSize = config.DiskSize
	manifest.Spec.SSH = config.EnableSSH

	return manifest
}

// createTempManifestFile creates a temporary YAML file containing the manifest
func createTempManifestFile(manifest *Manifest) (string, error) {
	yamlData, err := yaml.Marshal(manifest)
	if err != nil {
		return "", fmt.Errorf("failed to marshal manifest: %v", err)
	}

	tmpFile, err := os.CreateTemp("", "ignite-manifest-*.yaml")
	if err != nil {
		return "", fmt.Errorf("failed to create temp file: %v", err)
	}

	if _, err := tmpFile.Write(yamlData); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to write manifest: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		os.Remove(tmpFile.Name())
		return "", fmt.Errorf("failed to close manifest file: %v", err)
	}

	return tmpFile.Name(), nil
}

// runIgniteCommand executes an ignite command with the given arguments
func runIgniteCommand(args ...string) error {
	fmt.Println("Running ignite command:", args)
	cmd := exec.Command("sudo", append([]string{"ignite"}, args...)...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ignite command failed: %v\nStdout: %s\nStderr: %s",
			err, stdout.String(), stderr.String())
	}

	return nil
}

// getVMIP retrieves the IP address of a VM
func getVMIP(vmName string) (string, error) {
	cmd := exec.Command("sudo", "ignite", "ps")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to get VM status: %v", err)
	}

	// Parse output to find IP
	for _, line := range strings.Split(stdout.String(), "\n") {
		if strings.Contains(line, vmName) {
			fields := strings.Fields(line)
			if len(fields) >= 13 {
				return fields[12], nil
			}
		}
	}

	return "", fmt.Errorf("IP address not found for VM: %s", vmName)
}
