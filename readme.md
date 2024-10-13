# arbfeedtowasm

This project, written in Go, compiles into a WebAssembly (.wasm) module that processes messages from the Arbitrum Sequencer Feed, decoding them into human-readable transactions. It’s designed for easy integration across different environments where parsing Arbitrum sequencer data is required, offering an efficient and lightweight solution for interpreting Arbitrum messages.

## Development Setup

### Building with Makefile

To automate tasks like compiling contracts and setting up dependencies, you can use the Makefile provided in the Nitro submodule. The Makefile defines a set of tasks that streamline the build process.

```bash
$ cd nitro
$ make contracts
$ make build-node-deps
```

### Configuring `go.mod`

```mod
module arbfeedtowasm

replace github.com/offchainlabs/nitro => ./nitro
replace github.com/ethereum/go-ethereum => ./nitro/go-ethereum
replace github.com/VictoriaMetrics/fastcache => ./nitro/fastcache

go 1.23.2

```

### Setting Up Dependencies

```shell
$ go mod tidy
```

## Initial Setup

<b>For reference only, no longer needed after initial run.</b>

### Module Initialization

To initialize the Go module for this project, you can run the following command:

```bash
$ git mod init arbfeedtowasm
```

This command creates a new Go module named `arbfeedtowasm`, setting up the necessary `go.mod` file, which tracks dependencies and module versions for your project.

### Adding Submodules

Since this project relies on code from the Arbitrum Nitro repository, you need to add it as a Git submodule. This allows you to integrate external code while keeping it in sync with your project.

```bash
$ git submodule add https://github.com/OffchainLabs/nitro
```

This command adds the Nitro repository as a submodule, ensuring that you can track changes from the upstream Nitro project directly in your repository.

### Initializing Submodules Recursively

After adding the submodule, you’ll need to initialize it, along with any nested submodules that the Nitro project might depend on. This ensures all required dependencies are set up.

```bash
$ git submodule update --init --recursive
```

This command initializes the submodules recursively, pulling down all the necessary files for the project to work properly.