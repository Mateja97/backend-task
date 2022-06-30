package chaintracker

import (
	"backend-task/backend/chain"
	"backend-task/backend/kafka"
	"backend-task/backend/storage"
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ChainTracker struct {
	ContractAddress common.Address
	EthClient       *ethclient.Client
	Logs            chan types.Log
	ABI             abi.ABI
	Sub             ethereum.Subscription
	producer        kafka.KafkaProducer
}

func (ct *ChainTracker) Init(port, contract, network, topic string, brokers []string) error {
	ct.ContractAddress = common.HexToAddress(contract)
	var err error
	ct.EthClient, err = ethclient.Dial(network)
	if err != nil {
		return nil
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{
			ct.ContractAddress,
		},
	}
	ct.Logs = make(chan types.Log)

	ct.Sub, err = ct.EthClient.SubscribeFilterLogs(context.Background(), query, ct.Logs)
	if err != nil {
		return err
	}

	ct.ABI, err = abi.JSON(strings.NewReader(string(storage.StorageABI)))
	if err != nil {
		return err
	}
	err = ct.producer.Init(brokers, topic)
	if err != nil {
		log.Println("[ERROR] Producer init failed")
		return err
	}

	return nil
}

func (ct *ChainTracker) Run() {

	for {
		select {
		case err := <-ct.Sub.Err():
			if err != nil {
				log.Println("[ERROR]:Sub err:", err)
			}
		case vLog := <-ct.Logs:
			e, err := ct.ABI.Unpack("Sent", vLog.Data)
			if err != nil {
				log.Println("Cannot unpack log", err.Error())
				continue
			}
			fmt.Println(e[0], "   ", e[1])
			entity := chain.ChainEntity{
				Symbol: e[0].(string),
				Amount: e[1].(*big.Int),
			}
			ct.producer.SendMessage(entity)

		}
	}
}

func (ct *ChainTracker) Stop() {
	ct.Sub.Unsubscribe()
	ct.EthClient.Close()
	log.Println("Graceful shutdown complete.")
}
