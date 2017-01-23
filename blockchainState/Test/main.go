package main

import (
	"fmt"
	"os"

	. "github.com/FactomProject/factomd/blockchainState"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/database/databaseOverlay"
	"github.com/FactomProject/factomd/database/hybridDB"
)

const level string = "level"
const bolt string = "bolt"

func main() {
	fmt.Println("Usage:")
	fmt.Println("Test level/bolt DBFileLocation")

	if len(os.Args) < 3 {
		fmt.Println("\nNot enough arguments passed")
		os.Exit(1)
	}
	if len(os.Args) > 3 {
		fmt.Println("\nToo many arguments passed")
		os.Exit(1)
	}

	levelBolt := os.Args[1]

	if levelBolt != level && levelBolt != bolt {
		fmt.Println("\nFirst argument should be `level` or `bolt`")
		os.Exit(1)
	}
	path := os.Args[2]

	var dbase *hybridDB.HybridDB
	var err error
	if levelBolt == bolt {
		dbase = hybridDB.NewBoltMapHybridDB(nil, path)
	} else {
		dbase, err = hybridDB.NewLevelMapHybridDB(path, false)
		if err != nil {
			panic(err)
		}
	}

	CheckDatabase(dbase)
}

func CheckDatabase(db interfaces.IDatabase) {
	if db == nil {
		return
	}

	dbo := databaseOverlay.NewOverlay(db)
	bs := new(BlockchainState)
	bs.Init()
	bl := new(BalanceLedger)
	bl.Init()

	dBlock, err := dbo.FetchDBlockHead()
	if err != nil {
		panic(err)
	}
	if dBlock == nil {
		panic("DBlock head not found")
	}

	fmt.Printf("\tStarting\n")

	for i := 0; i < int(dBlock.GetDatabaseHeight()); i++ {

		set := FetchBlockSet(dbo, i)
		if i%1000 == 0 {
			fmt.Printf("\"%v\", //%v\n", set.DBlock.DatabasePrimaryIndex(), set.DBlock.GetDatabaseHeight())
		}

		err := bs.ProcessBlockSet(set.DBlock, set.FBlock, set.ECBlock, set.EBlocks)
		if err != nil {
			panic(err)
		}

		err = bl.ProcessFBlock(set.FBlock)
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("\tFinished!\n")

	b, err := bs.MarshalBinaryData()
	if err != nil {
		panic(err)
	}
	fmt.Printf("BS size - %v\n", len(b))

	b, err = bl.MarshalBinaryData()
	if err != nil {
		panic(err)
	}
	fmt.Printf("BL size - %v\n", len(b))
}

type BlockSet struct {
	ABlock  interfaces.IAdminBlock
	ECBlock interfaces.IEntryCreditBlock
	FBlock  interfaces.IFBlock
	DBlock  interfaces.IDirectoryBlock
	EBlocks []interfaces.IEntryBlock
}

func FetchBlockSet(dbo interfaces.DBOverlay, index int) *BlockSet {
	bs := new(BlockSet)

	dBlock, err := dbo.FetchDBlockByHeight(uint32(index))
	if err != nil {
		panic(err)
	}
	bs.DBlock = dBlock

	if dBlock == nil {
		return bs
	}
	entries := dBlock.GetDBEntries()
	for _, entry := range entries {
		switch entry.GetChainID().String() {
		case "000000000000000000000000000000000000000000000000000000000000000a":
			aBlock, err := dbo.FetchABlock(entry.GetKeyMR())
			if err != nil {
				panic(err)
			}
			bs.ABlock = aBlock
			break
		case "000000000000000000000000000000000000000000000000000000000000000c":
			ecBlock, err := dbo.FetchECBlock(entry.GetKeyMR())
			if err != nil {
				panic(err)
			}
			bs.ECBlock = ecBlock
			break
		case "000000000000000000000000000000000000000000000000000000000000000f":
			fBlock, err := dbo.FetchFBlock(entry.GetKeyMR())
			if err != nil {
				panic(err)
			}
			bs.FBlock = fBlock
			break
		default:
			eBlock, err := dbo.FetchEBlock(entry.GetKeyMR())
			if err != nil {
				panic(err)
			}
			bs.EBlocks = append(bs.EBlocks, eBlock)
			break
		}
	}

	return bs
}
