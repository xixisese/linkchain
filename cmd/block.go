package cmd

import (
	"github.com/linkchain/common/math"
	"github.com/linkchain/common/util/log"
	meta_block "github.com/linkchain/meta/block"
	"github.com/linkchain/poa/manage"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	RootCmd.AddCommand(chainInfoCmd)
	RootCmd.AddCommand(blockCmd)
	blockCmd.AddCommand(heightCmd, mineCmd, hashCmd)
}

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "block command",
}

var mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "generate a new block",
	Run: func(cmd *cobra.Command, args []string) {
		block, err := manage.GetManager().BlockManager.CreateBlock()
		if err != nil {
			log.Error("mine", "New Block error", err)
			return
		}
		txs := manage.GetManager().TransactionManager.GetAllTransaction()
		block.SetTx(txs)

		block, err = manage.GetManager().BlockManager.RebuildBlock(block)
		if err != nil {
			log.Error("mine", "Rebuild Block error", err)
			return
		}
		manage.GetManager().BlockManager.ProcessBlock(block)
		manage.GetManager().NewBlockEvent.Post(meta_block.NewMinedBlockEvent{Block: block})
	},
}

var heightCmd = &cobra.Command{
	Use:   "height",
	Short: "get a block by height in chainmanage",
	Run: func(cmd *cobra.Command, args []string) {
		example := []string{"example", "block height 0"}
		if len(args) != 1 {
			log.Error("getblockbyheight", "error", "please input height", example[0], example[1])
			return
		}

		height, error := strconv.Atoi(args[0])
		if error != nil {
			log.Error("getblockbyheight ", "error", error, example[0], example[1])
			return
		}

		if uint32(height) > manage.GetManager().ChainManager.GetBestBlock().GetHeight() || height < 0 {
			log.Error("getblockbyheight ", "error", "height is out of range", example[0], example[1])
			return
		}
		block, err := manage.GetManager().ChainManager.GetBlockByHeight(uint32(height))
		if err != nil {
			log.Error("getblockbyheight ", "error", err)
		} else {
			log.Info("block", "data", block)
		}
	},
}

var hashCmd = &cobra.Command{
	Use:   "hash",
	Short: "get a block by hash in chainmanage",
	Run: func(cmd *cobra.Command, args []string) {
		example := []string{"example", "block hash 98acd27a58c79eaab05ea4abd0daa8e63021df3bf2e65fcb38e2474fb706c3fe"}
		if len(args) != 1 {
			log.Error("getblockbyhash", "error", "please input hash", example[0], example[1])
			return
		}

		hash, err := math.NewHashFromStr(args[0])
		if err != nil {
			log.Error("getblockbyhash", "error", err, example[0], example[1])
		}
		block, err := manage.GetManager().ChainManager.GetBlockByHash(*hash)
		if err != nil {
			log.Error("getblockbyhash ", "error", err)
		} else {
			log.Info("block", "data", block)
		}
	},
}