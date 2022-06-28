package chaintracker

import (
	"backend-task/internal/chain"
	"backend-task/internal/storage"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
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
	ServerURL       string
}

func (ct *ChainTracker) Init(port, contract, network, server string) error {
	ct.ServerURL = server
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
			ct.SendToServer(e)
			//fmt.Println(reflect.TypeOf(e[0]), "   ", reflect.TypeOf(e[1]))
			/*entity := chain.ChainEntity{
				Symbol: e[0].(string),
				Amount: e[1].(*big.Int),
			}*/

		}
	}
}

func (ct *ChainTracker) Stop() {
	ct.Sub.Unsubscribe()
	ct.EthClient.Close()
	log.Println("Graceful shutdown complete.")
}

func (ct *ChainTracker) SendToServer(e []interface{}) {
	requestBody, err := json.Marshal(chain.ChainEntity{
		Symbol: e[0].(string),
		Amount: e[1].(*big.Int),
	})
	if err != nil {
		log.Println("[ERROR] SendToServer. Failed to Marshall data", err.Error())
	}
	_, err = http.Post(ct.ServerURL+"/price/chain", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println("[ERROR] SendToServer. Failed to post data for :", string(requestBody), err.Error())
	}
	log.Println("[INFO] Data sent to: ", ct.ServerURL)
}
