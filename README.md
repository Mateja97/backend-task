#### DBT(Decentar backend task)<br>
    contract: 0x3af33feDF748f5439CD04130A15356b96d3Ad3c6 
## Abigen setup<br>
abigen --bin=Storage.bin --abi=Storage.abi --pkg=storage --out=Storage.go<br>
### Services<br>

## Dbt<br>
Service that is providing:
1. rest api for token values on chain from smart contract events and coin gecko
2. rest api for token values on chain from smart contract events and coin gecko for a specified date
3. every 60 seconds track coin gecko values for configuration token ids and publish to publisher if price differs more then 2% from on chain value
4. read chain data from kafka and store to the chaincache
5. provide websocket for current onchain values
    go run ./backend/cmd/dbt -kafka.brokers=localhost:9092 <br>

## Chain Tracker <br>
Service that tracks on chain events and publish them to the kafka
    go run ./backend/cmd/chaintrackerService -contract=0x3af33feDF748f5439CD04130A15356b96d3Ad3c6 -kafka.brokers=localhost:9092 <br>

## Chain Publisher <br>
On post request call contract Set function
    go run ./backend/cmd/chainpublisherService -private.key=<private.key> -contract=0x3af33feDF748f5439CD04130A15356b96d3Ad3c6

## Diagram
![alt.text](https://github.com/Mateja97/backend-task/blob/master/diagram.png?raw=true)