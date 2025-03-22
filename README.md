# VM Provisioning CLI Tool

A command-line interface for provisioning virtual machines using Weave Ignite.

## Installation

```bash
go build -o poridhictl 
```

## Usage

To launch a new VM:

```bash
./poridhictl launch vm --name my-vm-name [flags]
```

To list all VMs:

```bash
./poridhictl list
```

Example output:
```
┌────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
| VM ID            | NAME       | IMAGE                              | KERNEL                           | SIZE   | CPUS | MEMORY    | CREATED | STATUS | IP          |
| 61f98da33b021682 | my-vm      | shajalahamedcse/only-k3-go:v1.0.10 | weaveworks/ignite-kernel:5.10.51 | 3.0 GB | 2    | 1024.0 MB | 10s ago | Up     | 10.62.0.200 |
└────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```

### Available Flags

- `--name`, `-n`: Name of the VM (required)
- `--cpus`, `-c`: Number of CPUs (default: 2)
- `--memory`, `-m`: Memory size (default: "1GB")
- `--disk-size`, `-d`: Disk size (default: "3GB")
- `--image`, `-i`: OCI image (default: "shajalahamedcse/only-k3-go:v1.0.10")
- `--ssh`, `-s`: Enable SSH (default: true)

### Example

```bash
./poridhictl launch vm --name master-1 --cpus 4 --memory 2GB --disk-size 5GB
``` 