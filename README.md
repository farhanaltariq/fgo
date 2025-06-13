# fgoinstall

[![Go Version](https://img.shields.io/badge/Go-1.17+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`fgoinstall` is a simple CLI tool to **find and install the latest compatible version** of a Go module for your current Go environment.

Perfect for developers working with older Go versions like `1.17` or `1.18`, where newer modules might break compatibility.

---

## 📦 Installation

### Option 1: Build from source

```bash
git clone https://github.com/farhanaltariq/fgoinstall.git
cd fgoinstall
go build -o fgoinstall
sudo mv fgoinstall /usr/local/bin/  # Optional: Add to PATH
```

### Option 2: Direct install (requires Go)
```
go install github.com/farhanaltariq/fgoinstall@latest
```

### 🚀 Usage
```
fgoinstall {go-module}
```
- #### Example
```
fgoinstall github.com/vektra/mockery/v2
```

