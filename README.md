contract: 0x3af33feDF748f5439CD04130A15356b96d3Ad3c6

abigen --bin=Storage.bin --abi=Storage.abi --pkg=storage --out=Storage.go

$ go run ./backend/cmd/dbt -kafka.brokers=localhost:9092
go run ./backend/cmd/chaintrackerService -contract=0x3af33feDF748f5439CD04130A15356b96d3Ad3c6 -kafka.brokers=localhost:9092