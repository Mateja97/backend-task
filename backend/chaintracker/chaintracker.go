package chaintracker

import (
	"backend-task/backend/chain"
	"backend-task/backend/kafka"
	"backend-task/backend/storage"
	"context"
	"log"
	"math/big"
	"strings"
	"time"

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
	FilteredLogs    []types.Log
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
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []common.Address{
			ct.ContractAddress,
		},
	}
	ct.Logs = make(chan types.Log)
	ct.FilteredLogs, err = ct.EthClient.FilterLogs(context.Background(), query)
	if err != nil {
		return err
	}
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
	go func() {
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
				date := time.Unix(e[2].(*big.Int).Int64(), 0)
				layout := "02-01-2006"
				t := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC)
				entity := chain.ChainEntity{
					Symbol: e[0].(string),
					Amount: e[1].(*big.Int),
					Date:   t.Format(layout),
				}
				//Send chain entity to the kafka
				ct.producer.SendMessage(entity)

			}
		}
	}()
	//Check if there is already data on the chain
	if len(ct.FilteredLogs) > 0 {
		for _, log := range ct.FilteredLogs {
			ct.Logs <- log
		}
	}
}

func (ct *ChainTracker) Stop() {
	ct.Sub.Unsubscribe()
	ct.EthClient.Close()
	log.Println("Graceful shutdown complete.")
}
