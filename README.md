# Required Packages
- go get github.com/sandertv/gophertunnel/minecraft/protocol/packet
- go get github.com/sandertv/gophertunnel/minecraft/protocol/login
- go get github.com/sandertv/gophertunnel/minecraft/auth
- go get github.com/sandertv/gophertunnel/minecraft

# Wha doe thi do?
- It freezes stupid servers without game packets limiter
- disables everything and commands except moving & /help as it is client side
# Requirements:
- Linux Server with 2gb *
- Golang (1.23) Min *
- electric fence (optional)
# How to
- git clone https://github.com/NoSpacesFlies/MC-Go-Killer/
- cd MC-Go-Killer
- `ulimit -n 99999`
- go mod init main
- go run main.go <TARGET> <PORT> <DURATION>
