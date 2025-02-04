//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=models/models.cfg.yaml api/openapi.yaml

package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"strconv"

	"github.com/Taraxa-project/taraxa-indexer/api"
	"github.com/Taraxa-project/taraxa-indexer/internal/chain"
	"github.com/Taraxa-project/taraxa-indexer/internal/common"
	"github.com/Taraxa-project/taraxa-indexer/internal/indexer"
	"github.com/Taraxa-project/taraxa-indexer/internal/lara"
	"github.com/Taraxa-project/taraxa-indexer/internal/logging"
	"github.com/Taraxa-project/taraxa-indexer/internal/metrics"
	"github.com/Taraxa-project/taraxa-indexer/internal/oracle"
	"github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble"
	migration "github.com/Taraxa-project/taraxa-indexer/internal/storage/pebble/migrations"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	log "github.com/sirupsen/logrus"
)

var (
	http_port                        *int
	metrics_port                     *int
	blockchain_ws                    *string
	chain_id                         *int
	data_dir                         *string
	log_level                        *string
	yield_saving_interval            *int
	validators_yield_saving_interval *int
	sync_queue_limit                 *int
	signing_key                      *string
	oracle_address                   *string
	lara_address                     *string
	graphQLEndpoint                  *string
	lara_enabled                     *bool
	general_block_time               *int
	last_snapshot_id                 *uint64
	chain_stats_interval             *int
)

func init() {
	http_port = flag.Int("http_port", 8080, "port to listen")
	metrics_port = flag.Int("metrics_port", 2112, "metrics http port")
	blockchain_ws = flag.String("blockchain_ws", "ws://localhost:8777", "ws url to connect to blockchain")
	chain_id = flag.Int("chain_id", 841, "chain id")
	data_dir = flag.String("data_dir", "./data", "path to directory where indexer database will be saved")
	log_level = flag.String("log_level", "info", "minimum log level. could be only [trace, debug, info, warn, error, fatal]")
	yield_saving_interval = flag.Int("yield_saving_interval", 100, "interval for saving total yield")
	validators_yield_saving_interval = flag.Int("validators_yield_saving_interval", 100, "interval for saving validators yield")
	sync_queue_limit = flag.Int("sync_queue_limit", 10, "limit of blocks in the sync queue")
	oracle_address = flag.String("oracle_address", "0x8299040F890AD1c13Edd4D1381AB862Cc8b4a464", "oracles address")
	lara_address = flag.String("lara_address", "0x45225cd7B294E17d88eb0E62c935Af525d67798F", "lara address")
	signing_key = flag.String("signing_key", "", "signing key")
	graphQLEndpoint = flag.String("graphQLEndpoint", "https://indexer.community.taraxa.io/subgraphs/name/Liquid-staking/lara-subgraph", "graphql endpoint")
	lara_enabled = flag.Bool("lara_enabled", false, "enable lara")
	general_block_time = flag.Int("general_block_time", 3600, "general block time in milliseconds")
	last_snapshot_id = flag.Uint64("last_snapshot_id", 335, "last snapshot id")
	chain_stats_interval = flag.Int("chain_stats_interval", 100, "interval for saving chain stats")

	flag.Parse()

	logging.Config(filepath.Join(*data_dir, "logs"), *log_level)
	log.Print("\n\n\n")
	log.WithFields(log.Fields{
		"http_port":          *http_port,
		"blockchain_ws":      *blockchain_ws,
		"chain_id":           *chain_id,
		"signing_key":        *signing_key,
		"oracle_address":     *oracle_address,
		"lara_address":       *lara_address,
		"data_dir":           *data_dir,
		"log_level":          *log_level,
		"graphQLEndpoint":    *graphQLEndpoint,
		"lara_enabled":       *lara_enabled,
		"general_block_time": *general_block_time,
		"last_snapshot_id":   *last_snapshot_id}).
		Info("Application started")
}

func setupCloseHandler(fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	go func() {
		<-c
		fn()
		os.Exit(0)
	}()
}

func main() {
	// Lara param sanity checks
	if *signing_key == "" && *oracle_address == "" && *lara_address == "" {
		log.WithFields(log.Fields{"signing_key": *signing_key, "oracle_address": *oracle_address, "lara_address": *lara_address}).Fatal("Oracle address, Lara address and signing key should be both set but both empty")
	}

	is_lara_enabled := *lara_enabled
	rpc, err := ethclient.Dial(*blockchain_ws)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to blockchain")
	}
	log.Info("RPC initialized")

	if is_lara_enabled {
		log.Info("Starting Taraxa Indexer in Lara mode")
		lara := lara.MakeLara(rpc, *signing_key, *lara_address, *oracle_address, *graphQLEndpoint, *chain_id, *last_snapshot_id)
		log.Info("Lara initialized")

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			lara.Run(*validators_yield_saving_interval, *general_block_time)
		}()
		wg.Wait()
	} else {
		log.Info("Starting Taraxa Indexer in non-Lara mode")
		st := pebble.NewStorage(filepath.Join(*data_dir, "db"))
		setupCloseHandler(func() { st.Close() })
		fin := st.GetFinalizationData()
		swagger, err := api.GetSwagger()
		if err != nil {
			log.WithError(err).Fatal("Error loading swagger spec")
		}

		manager := migration.NewManager(st, *blockchain_ws)
		err = manager.ApplyAll()
		if err != nil {
			log.WithError(err).Fatal("Error applying migrations")
		}

		swagger.Servers = nil

		e := echo.New()

		e.Use(echomiddleware.OapiRequestValidator(swagger))
		// Add http error handler to return a proper error JSON on request error
		e.HTTPErrorHandler = func(err error, ctx echo.Context) {
			_ = ctx.JSON(http.StatusInternalServerError, map[string]any{"message": err.Error()})
		}

		c := common.DefaultConfig()
		c.TotalYieldSavingInterval = uint64(*yield_saving_interval)
		c.ValidatorsYieldSavingInterval = uint64(*validators_yield_saving_interval)
		c.SyncQueueLimit = uint64(*sync_queue_limit)
		c.ChainStatsInterval = *chain_stats_interval

		log.WithFields(log.Fields{"pbft_count": fin.PbftCount, "dag_count": fin.DagCount, "trx_count": fin.TrxCount}).Info("Loaded db with")
		chainStats := chain.MakeStats(c.ChainStatsInterval)
		apiHandler := api.NewApiHandler(st, c, chainStats)
		api.RegisterHandlers(e, apiHandler)

		o := oracle.MakeOracle(rpc, *signing_key, *oracle_address, *chain_id, *graphQLEndpoint, *st)

		go indexer.MakeAndRun(*blockchain_ws, st, c, o, chainStats)

		// start a http server for prometheus on a separate go routine
		go metrics.RunPrometheusServer(":" + strconv.FormatInt(int64(*metrics_port), 10))

		err = e.Start(":" + strconv.FormatInt(int64(*http_port), 10))
		log.WithError(err).Fatal("Can't start http server")
	}
}
