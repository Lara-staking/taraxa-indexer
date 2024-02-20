package lara

import (
	"log"

	dpos_contract "github.com/Taraxa-project/taraxa-indexer/abi/dpos"
	lara_factory "github.com/Taraxa-project/taraxa-indexer/abi/factory"
	apy_oracle "github.com/Taraxa-project/taraxa-indexer/abi/oracle"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

// Start Generation Here

// LaraFactoryListener is a struct that handles listening to LaraFactory events
type LaraFactory struct {
	contract                    *lara_factory.LaraFactory // LaraFactory contract instance
	laraInstances               map[string]*Lara          // Map to store Lara instances by address
	Eth                         *ethclient.Client         // Ethereum client instance
	signer                      *bind.TransactOpts        // Transaction signer
	singing_key                 string                    // Private key
	chainID                     *int                      // Chain ID
	oracle                      *apy_oracle.ApyOracle     // ApyOracle contract instance
	dpos                        *dpos_contract.DposContract
	laraCreatedSubscription     event.Subscription
	laraDeactivatedSubscription event.Subscription
	laraActivatedSubscrpition   event.Subscription
}

func MakeLaraFactory(rpc *ethclient.Client, signing_key, deployment_address, oracle_address string, chainID int) *LaraFactory {
	f := new(LaraFactory)
	f.Eth = rpc
	f.singing_key = signing_key
	f.chainID = &chainID
	contract, err := lara_factory.NewLaraFactory(common.HexToAddress(deployment_address), f.Eth)
	if err != nil {
		log.Fatalf("Failed to create contract: %v", err)
	}
	f.oracle, err = apy_oracle.NewApyOracle(common.HexToAddress(oracle_address), f.Eth)
	if err != nil {
		log.Fatalf("Failed to create oracle: %v", err)
	}
	f.dpos, err = dpos_contract.NewDposContract(common.HexToAddress("0x00000000000000000000000000000000000000fe"), f.Eth)
	if err != nil {
		log.Fatalf("Failed to create dpos: %v", err)
	}
	f.contract = contract
	return &LaraFactory{
		contract:      contract,
		laraInstances: make(map[string]*Lara),
	}
}

func (f *LaraFactory) CreateLara(lara_address, delegator_address string) {
	laraInstance := MakeLara(f.Eth, f.signer, f.oracle, lara_address, delegator_address)
	f.laraInstances[lara_address] = laraInstance
}

func (f *LaraFactory) GetLaraInstance(lara_address string) *Lara {
	return f.laraInstances[lara_address]
}

func (f *LaraFactory) GetLaraInstances() map[string]*Lara {
	return f.laraInstances
}

func (f *LaraFactory) DeactivateLara(lara_address string) {
	// Implement logic to deactivate Lara instance
	laraInstance := f.GetLaraInstance(lara_address)
	laraInstance.Stop()
}

func (f *LaraFactory) ActivateLara(lara_address string) {
	// Implement logic to activate Lara instance
	laraInstance := f.GetLaraInstance(lara_address)
	laraInstance.Restart()
}

func (f *LaraFactory) Run() {
	// Get all past LaraCreated events and initialize Lara instances
	f.InitializePastLaraInstances()
	// once done, start listening to LaraFactory events
	f.StartListening()
}

// StartListening starts listening to LaraFactory events
func (f *LaraFactory) StartListening() {
	// Start listening to LaraCreated events
	laraCreatedCh := make(chan *lara_factory.LaraFactoryLaraCreated)
	sub, err := f.contract.WatchLaraCreated(nil, laraCreatedCh, nil, nil)
	if err != nil {
		log.Fatal("Error watching LaraCreated event: ", err)
	}
	f.laraCreatedSubscription = sub

	// Start listening to LaraDeactivated events
	laraDeactivatedCh := make(chan *lara_factory.LaraFactoryLaraDeactivated)
	subDeactivated, err := f.contract.WatchLaraDeactivated(nil, laraDeactivatedCh, nil, nil)
	if err != nil {
		log.Fatal("Error watching LaraDeactivated event: ", err)
	}
	f.laraDeactivatedSubscription = subDeactivated

	// Start listening to LaraActivated events
	laraActivatedCh := make(chan *lara_factory.LaraFactoryLaraActivated)
	subActivated, err := f.contract.WatchLaraActivated(nil, laraActivatedCh, nil, nil)
	if err != nil {
		log.Fatal("Error watching LaraActivated event: ", err)
	}
	f.laraActivatedSubscrpition = subActivated

	go func() {
		for {
			select {
			case event := <-laraCreatedCh:
				// Handle LaraCreated event
				f.handleLaraCreated(event)
			case event := <-laraDeactivatedCh:
				// Handle LaraDeactivated event
				f.handleLaraDeactivated(event)
			case event := <-laraActivatedCh:
				// Handle LaraActivated event
				f.handleLaraActivated(event)
			}
		}
	}()
}

func (f *LaraFactory) InitializePastLaraInstances() {
	// Filter for all past LaraCreated events
	laraCreatedIterator, err := f.contract.FilterLaraCreated(nil, nil, nil)
	if err != nil {
		log.Fatal("Error filtering LaraCreated events: ", err)
	}
	for laraCreatedIterator.Next() {
		event := laraCreatedIterator.Event
		laraAddress := event.LaraAddress.Hex()
		delegatorAddress := event.Creator.Hex()
		// Initialize Lara instance for each past event
		f.CreateLara(laraAddress, delegatorAddress)
	}

	laraDeactivatedIterator, err := f.contract.FilterLaraDeactivated(nil, nil, nil)
	if err != nil {
		log.Fatal("Error filtering LaraDeactivated events: ", err)
	}
	for laraDeactivatedIterator.Next() {
		event := laraDeactivatedIterator.Event
		laraAddress := event.LaraAddress.Hex()
		// Deactivate Lara instance for each past event
		f.DeactivateLara(laraAddress)
	}

	laraActivatedIterator, err := f.contract.FilterLaraActivated(nil, nil, nil)
	if err != nil {
		log.Fatal("Error filtering LaraActivated events: ", err)
	}
	for laraActivatedIterator.Next() {
		event := laraActivatedIterator.Event
		laraAddress := event.LaraAddress.Hex()
		// Activate Lara instance for each past event
		f.ActivateLara(laraAddress)
	}
}

func (f *LaraFactory) StopListening() {
	f.laraCreatedSubscription.Unsubscribe()
	f.laraDeactivatedSubscription.Unsubscribe()
	f.laraActivatedSubscrpition.Unsubscribe()
}

// handleLaraCreated is a function to handle LaraCreated event
func (f *LaraFactory) handleLaraCreated(event *lara_factory.LaraFactoryLaraCreated) {
	f.CreateLara(event.LaraAddress.Hex(), event.Creator.Hex())
}

// handleLaraDeactivated is a function to handle LaraDeactivated event
func (f *LaraFactory) handleLaraDeactivated(event *lara_factory.LaraFactoryLaraDeactivated) {
	f.DeactivateLara(event.LaraAddress.Hex())
}

// handleLaraActivated is a function to handle LaraActivated event
func (l *LaraFactory) handleLaraActivated(event *lara_factory.LaraFactoryLaraActivated) {
	l.ActivateLara(event.LaraAddress.Hex())
}
