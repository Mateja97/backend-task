package chainpublisher

import (
	"backend-task/backend/cors"
	"context"
	"crypto/ecdsa"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

type ChainPublisher struct {
	client   *ethclient.Client
	server   *http.Server
	wallet   common.Address
	contract string
	pk       *ecdsa.PrivateKey
}

func (cp *ChainPublisher) Init(network, privateKey, contract, port string) error {

	cp.contract = contract
	var err error
	cp.client, err = ethclient.Dial(network)
	if err != nil {
		return err
	}
	cp.pk, err = crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}

	publicKey := cp.pk.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return err
	}

	cp.wallet = crypto.PubkeyToAddress(*publicKeyECDSA)
	r := mux.NewRouter()
	r.HandleFunc("/publish", cp.PublishToChain()).Methods("POST")
	cp.server = &http.Server{
		Addr:    port,
		Handler: cors.CORSEnabled(r),
	}
	return nil
}

func (cp *ChainPublisher) Run() {
	if err := cp.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("[ERROR] Chainpublisher server failed")
	}
}
func (cp *ChainPublisher) Stop() {
	if err := cp.server.Shutdown(context.Background()); err != nil {
		log.Println("[ERROR] Chainpublisher server shutdown failed")
	}
	log.Println("[INFO] ChainPublisher stop gracefully")
}
