# Go Modules Test

This is a test project to validate Go module replace directives behavior when modules are used as third-party dependencies.

## Structure

- `types/` - Basic types module
- `utils/` - Utilities module that depends on types
- `main.go` - Main application using both modules

## Testing Replace Directives

This project demonstrates:
1. Local replace directives for development
2. How submodules work when imported by third parties (replace directives are ignored)

## Usage

```bash
go run main.go
```
