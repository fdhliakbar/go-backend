Here’s a concise overview of commonly used Go commands to help you navigate the Go programming environment effectively:

General Commands
go run: Compiles and runs Go programs.
go build: Compiles source code into an executable binary.
go install: Installs the binary into the $GOPATH/bin directory.
go clean: Removes object files and cached files.
go test: Runs tests in the current package.
go fmt: Formats Go source code according to Go standards.
go vet: Reports potential issues in the code.
go mod: Manages Go modules (e.g., go mod init, go mod tidy).
go get: Adds, updates, or downloads dependencies.
go list: Lists packages or modules.
go doc: Displays documentation for a package or symbol.
go version: Displays the installed Go version.
go env: Prints Go environment variables.
Module-Specific Commands
go mod init: Initializes a new module.
go mod tidy: Cleans up unused dependencies.
go mod download: Downloads module dependencies.
go mod verify: Verifies dependencies have not been tampered with.
Debugging and Profiling
go tool: Runs various Go tools (e.g., pprof, cover).
go trace: Generates and views execution traces.

These commands are typically used as subcommands of the go tool. For more details, you can explore the official Go documentation or use go help <command> for specific guidance.