package chainpublisher

import (
	"backend-task/backend/chain"
	"backend-task/backend/storage"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (cp *ChainPublisher) PublishToChain() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nonce, err := cp.client.PendingNonceAt(context.Background(), cp.wallet)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := cp.client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		auth := bind.NewKeyedTransactor(cp.pk)
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		address := common.HexToAddress(cp.contract)
		instance, err := storage.NewStorage(address, cp.client)
		if err != nil {
			log.Fatal(err)
		}

		var entity chain.ChainEntity
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&entity)
		if err != nil {
			log.Println("[ERROR] ChainValue Decoding failed", err)
			return
		}
		tx, err := instance.Set(auth, entity.Symbol, entity.Amount)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("tx sent: %s", tx.Hash().Hex())
	}
}
