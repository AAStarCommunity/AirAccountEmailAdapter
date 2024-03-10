# FAQ

- F: `Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub`
- Q: `go env -w CGO_ENABLED=1`

---

- F: `cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%`
- Q: `choco install mingw -y`