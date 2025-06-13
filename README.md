# fgo

[![Go Version](https://img.shields.io/badge/Go-1.17+-blue.svg)](https://golang.org/dl/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`fgo` is a simple CLI tool to **find and install the latest compatible version** of a Go module for your current Go environment.

Perfect for developers working with older Go versions like `1.17` or `1.18`, where newer modules might break compatibility.

---

## 📦 Installation

### Option 1: Build from source

```bash
git clone https://github.com/farhanaltariq/fgo.git
cd fgo
go build -o fgo
sudo mv fgo /usr/local/bin/  # Optional: Add to PATH
```

### Option 2: Direct install (requires Go)
```
go install github.com/farhanaltariq/fgo@latest
```

### 🚀 Usage
```
fgo install {go-module}
```

- #### Example
```
fgo install github.com/vektra/mockery/v2
```

```
fgo get github.com/vektra/mockery/v2
```
