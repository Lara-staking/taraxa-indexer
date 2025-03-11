package lara

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/machinebox/graphql"
)

type StakerResponse struct {
	Stakers []Staker `json:"stakers"`
}

type Staker struct {
	ID string `json:"id"`
}

type SnapshotResponse struct {
	Snapshots []Snapshot `json:"snapshots"`
}

type Snapshot struct {
	ID           string `json:"id"`
	TotalRewards string `json:"totalRewards"`
	Block        string `json:"block"`
}

type SnapshotWithBlock struct {
	ID           uint64   `json:"id"`
	TotalRewards *big.Int `json:"totalRewards"`
	Block        uint64   `json:"block"`
}

func GetStakedTaraHolders(endpoint string, blockNumber uint64) []string {
	if endpoint == "" {
		log.Fatal("GraphQL endpoint is not set")
	}
	client := graphql.NewClient(endpoint)

	totalStakers := 0
	stakers := []string{}

	first := 100
	skip := 0

	for {
		req := graphql.NewRequest(`
			query($blockNumber: Int!, $first: Int!, $skip: Int!) {
				stakers(
				block: {number: $blockNumber},
				where: { stTaraBalance_gt: 0 },
				first: $first,
				skip: $skip
			) {
				id
			}
		}
	`)

		req.Var("blockNumber", int(blockNumber))
		req.Var("first", first)
		req.Var("skip", skip)

		var resp StakerResponse

		if err := client.Run(context.Background(), req, &resp); err != nil {
			log.Printf("GraphQL query error: %v", err)
			return nil
		}

		if len(resp.Stakers) == 0 {
			fmt.Println("No stakers found in the response")
			return nil
		}

		totalStakers += len(resp.Stakers)
		for _, staker := range resp.Stakers {
			stakers = append(stakers, staker.ID)
		}

		skip += first

		if len(resp.Stakers) < first {
			break
		}
	}
	log.Infof("Total stakers: %d", totalStakers)
	return stakers
}

func GetSnapshotsWithBlocks(endpoint string) []SnapshotWithBlock {
	if endpoint == "" {
		log.Fatal("GraphQL endpoint is not set")
	}
	client := graphql.NewClient(endpoint)

	snapshots := []SnapshotWithBlock{}

	first := 100
	skip := 0

	for {
		log.Infof("Fetching the first %d snapshots, skipping %d", first, skip)
		req := graphql.NewRequest(`
			query snapshotTaken($first: Int!, $skip: Int!){
			snapshots(first: $first, skip: $skip){
				id
				totalRewards
				block
			}
		}
	`)

		req.Var("first", first)
		req.Var("skip", skip)
		var resp SnapshotResponse

		if err := client.Run(context.Background(), req, &resp); err != nil {
			log.Printf("GraphQL query error: %v", err)
			return nil
		}

		if len(resp.Snapshots) == 0 {
			break
		}

		for _, snapshot := range resp.Snapshots {
			totalRewards, ok := big.NewInt(0).SetString(snapshot.TotalRewards, 10)
			if !ok {
				log.Printf("Failed to convert totalRewards to big.Int: %v", snapshot.TotalRewards)
				continue
			}
			if totalRewards.Cmp(big.NewInt(0)) > 0 {
				snapshotId, err := strconv.ParseUint(snapshot.ID, 10, 64)
				if err != nil {
					log.Printf("Failed to convert snapshot.ID to uint64: %v", snapshot.ID)
					continue
				}
				block, err := strconv.ParseUint(snapshot.Block, 10, 64)
				if err != nil {
					log.Printf("Failed to convert snapshot.Block to uint64: %v", snapshot.Block)
					continue
				}
				snapshots = append(snapshots, SnapshotWithBlock{ID: snapshotId, TotalRewards: totalRewards, Block: block})
			}
		}

		skip += first

		if len(resp.Snapshots) < first {
			break
		}
	}
	log.Infof("Total snapshots: %d", len(snapshots))

	return snapshots
}
