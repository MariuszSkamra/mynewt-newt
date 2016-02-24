package cli

import (
	"fmt"

	"git-wip-us.apache.org/repos/asf/incubator-mynewt-newt/newtmgr/protocol"
	"github.com/spf13/cobra"
)

func mempoolStatsRunCmd(cmd *cobra.Command, args []string) {
	runner, err := getTargetCmdRunner()
	if err != nil {
		nmUsage(cmd, err)
	}

	srr, err := protocol.NewMempoolStatsReadReq()
	if err != nil {
		nmUsage(cmd, err)
	}

	nmr, err := srr.EncodeWriteRequest()
	if err != nil {
		nmUsage(cmd, err)
	}

	if err := runner.WriteReq(nmr); err != nil {
		nmUsage(cmd, err)
	}

	rsp, err := runner.ReadResp()
	if err != nil {
		nmUsage(cmd, err)
	}

	msrsp, err := protocol.DecodeMempoolStatsReadResponse(rsp.Data)
	if err != nil {
		nmUsage(cmd, err)
	}

	fmt.Printf("Return Code = %d\n", msrsp.ReturnCode)
	if msrsp.ReturnCode == 0 {
		for k, info := range msrsp.MPools {
			fmt.Printf("  %s ", k)
			fmt.Printf("(blksize=%d nblocks=%d nfree=%d)",
				int(info["blksiz"].(float64)),
				int(info["nblks"].(float64)),
				int(info["nfree"].(float64)))
			fmt.Printf("\n")
		}
	}
}

func mempoolStatsCmd() *cobra.Command {
	mempoolStatsCmd := &cobra.Command{
		Use:   "mpstats",
		Short: "Read statistics from a remote endpoint",
		Run:   mempoolStatsRunCmd,
	}

	return mempoolStatsCmd
}
