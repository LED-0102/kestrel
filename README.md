# Kestrel - Distributed Commit Log

A high-performance, distributed commit log system built in Go, inspired by Apache Kafka. Kestrel provides a scalable, fault-tolerant platform for building distributed systems that require ordered, replicated log storage with strong consistency guarantees.

## 🎯 Project Concept

Kestrel implements a distributed commit log - an ordered, append-only sequence of records that serves as the source of truth for distributed systems. Think of it as a distributed database transaction log that multiple services can read from and write to, ensuring all participants see the same sequence of events in the same order.

### Key Concepts

- **Commit Log**: An immutable, ordered sequence of records with monotonically increasing offsets
- **Segments**: Log data is partitioned into segments for efficient storage and retrieval
- **Consensus**: Uses Raft consensus algorithm to ensure all nodes agree on the log state
- **Service Discovery**: Automatic peer discovery and cluster formation using gossip protocols

## ✨ Features & Capabilities

### Core Features
- **Distributed Architecture**: Multi-node cluster with automatic leader election and failover
- **Strong Consistency**: Raft consensus ensures all nodes have identical logs
- **High Performance**: Memory-mapped files and efficient indexing for fast reads/writes
- **Horizontal Scaling**: Add/remove nodes dynamically without downtime
- **Fault Tolerance**: Survives node failures with configurable replication

### Advanced Capabilities
- **gRPC API**: High-performance RPC interface with streaming support
- **TLS Security**: Mutual TLS authentication and encrypted communication
- **Access Control**: Fine-grained authorization using Casbin ACL policies
- **Service Discovery**: Automatic cluster membership with HashiCorp Serf
- **Observability**: OpenCensus integration for metrics and distributed tracing
- **Kubernetes Ready**: Helm charts and health probes for container orchestration

### API Operations
- **Produce**: Append records to the log with automatic offset assignment
- **Consume**: Read records by offset with support for streaming consumption
- **Cluster Management**: Get cluster topology and server information
- **Health Checks**: Built-in gRPC health checking for load balancers

## 🛠 Dependencies

### Runtime Dependencies
- **Go 1.22+**: Core runtime
- **HashiCorp Raft**: Consensus algorithm implementation
- **HashiCorp Serf**: Service discovery and cluster membership
- **gRPC**: High-performance RPC framework
- **Protocol Buffers**: Efficient serialization
- **Casbin**: Access control and authorization
- **Zap**: Structured logging
- **OpenCensus**: Observability and metrics

### Development Dependencies
- **CFSSL**: TLS certificate generation
- **Docker**: Containerization
- **Kubernetes/Helm**: Container orchestration
- **Protocol Buffer Compiler**: Code generation

### External Tools Required
```bash
# Certificate generation
go install github.com/cloudflare/cfssl/cmd/cfssl@latest
go install github.com/cloudflare/cfssl/cmd/cfssljson@latest

# Protocol buffer compilation
# Install protoc from https://grpc.io/docs/protoc-installation/
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 🚀 Setup & Installation

### Prerequisites
1. Go 1.22 or later
2. CFSSL tools for certificate generation
3. Protocol Buffer compiler (protoc)

### Quick Start

#### 1. Clone and Build
```bash
git clone https://github.com/LED-0102/kestrel.git
cd kestrel

# Install dependencies
go mod download

# Generate certificates for development
make gencert

# Initialize configuration directory
make init

# Run tests to verify setup
make test
```

#### 2. Start a Single Node (Development)
```bash
# Bootstrap the first node
go run cmd/kestrel/main.go \
  --bootstrap \
  --data-dir=/tmp/kestrel-node1 \
  --bind-addr=127.0.0.1:8401 \
  --rpc-port=8400 \
  --server-tls-cert-file=$HOME/.kestrel/server.pem \
  --server-tls-key-file=$HOME/.kestrel/server-key.pem \
  --server-tls-ca-file=$HOME/.kestrel/ca.pem \
  --acl-model-file=test/model.conf \
  --acl-policy-file=test/policy.csv
```

#### 3. Start Additional Nodes
```bash
# Node 2
go run cmd/kestrel/main.go \
  --data-dir=/tmp/kestrel-node2 \
  --bind-addr=127.0.0.1:8402 \
  --rpc-port=8401 \
  --start-join-addrs=127.0.0.1:8401 \
  --server-tls-cert-file=$HOME/.kestrel/server.pem \
  --server-tls-key-file=$HOME/.kestrel/server-key.pem \
  --server-tls-ca-file=$HOME/.kestrel/ca.pem \
  --acl-model-file=test/model.conf \
  --acl-policy-file=test/policy.csv

# Node 3
go run cmd/kestrel/main.go \
  --data-dir=/tmp/kestrel-node3 \
  --bind-addr=127.0.0.1:8403 \
  --rpc-port=8402 \
  --start-join-addrs=127.0.0.1:8401 \
  --server-tls-cert-file=$HOME/.kestrel/server.pem \
  --server-tls-key-file=$HOME/.kestrel/server-key.pem \
  --server-tls-ca-file=$HOME/.kestrel/ca.pem \
  --acl-model-file=test/model.conf \
  --acl-policy-file=test/policy.csv
```

### Docker Deployment
```bash
# Build Docker image
make build-docker TAG=v1.0.0

# Run with Docker
docker run -p 8400:8400 -p 8401:8401 \
  -v $HOME/.kestrel:/etc/kestrel \
  github.com/LED-0102/kestrel:v1.0.0 \
  --bootstrap \
  --server-tls-cert-file=/etc/kestrel/server.pem \
  --server-tls-key-file=/etc/kestrel/server-key.pem \
  --server-tls-ca-file=/etc/kestrel/ca.pem \
  --acl-model-file=/etc/kestrel/model.conf \
  --acl-policy-file=/etc/kestrel/policy.csv
```

### Kubernetes Deployment
```bash
# Install using Helm
helm install kestrel deploy/kestrel/ \
  --set replicas=3 \
  --set storage=10Gi
```

## 📋 Configuration

### TLS Configuration
The system requires TLS certificates for secure communication:
- **Server certificates**: For client connections
- **Peer certificates**: For inter-node Raft communication
- **CA certificates**: For certificate validation

### Access Control
Uses Casbin for role-based access control:
- `test/model.conf`: ACL model definition
- `test/policy.csv`: Permission policies

### Environment Variables
- `GRPC_GO_LOG_VERBOSITY_LEVEL`: gRPC logging verbosity
- `GRPC_GO_LOG_SEVERITY_LEVEL`: gRPC log severity

## 🧪 Development

### Available Make Targets
```bash
make init         # Initialize configuration directory
make gencert      # Generate TLS certificates
make test         # Run tests with race detection
make compile      # Generate protobuf code
make build-docker # Build Docker image
```

### Running Tests
```bash
# Run all tests
make test

# Run specific package tests
go test -race ./internal/log/...

# Run with verbose output
go test -v -race ./...
```

### Protocol Buffer Development
When modifying `.proto` files:
```bash
make compile
```

## 🏗 Architecture

### Component Overview
- **Agent**: Main orchestrator managing all components
- **Distributed Log**: Raft-based distributed storage engine
- **gRPC Server**: API layer with authentication/authorization
- **Service Discovery**: Cluster membership management
- **Storage Engine**: Segment-based log storage with indexing

### Network Architecture
- **Port 8400**: gRPC API (default)
- **Port 8401**: Serf gossip protocol (default)
- Single port multiplexing for gRPC and Raft traffic

### Data Flow
1. Clients connect via gRPC with TLS certificates
2. Requests are authenticated and authorized
3. Log operations go through Raft consensus
4. Data is replicated across cluster nodes
5. Responses are returned to clients

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Ensure all tests pass: `make test`
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🆘 Support

- **Issues**: Report bugs and feature requests via GitHub Issues
- **Documentation**: Additional documentation in `CLAUDE.md`
- **Examples**: Sample configurations in the `test/` directory