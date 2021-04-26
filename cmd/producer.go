package cmd

import (
	"accounting_migration/config"
	"accounting_migration/config/logger"
	"accounting_migration/queue"

	"github.com/spf13/cobra"
)

var (
	ProducerHandler = func(cmd *cobra.Command, args []string) {
		config.Start(".")
		logger.Info().Msg("Producer Start")

		for i := 0; i < 10; i++ {
			err := queue.Enqueue(queue.MessageQueue{
				Type:     1,
				UserID:   uint32(i),
				CabangID: 1,
			})

			if err != nil {
				logger.Fatal().Msg("Error publish sream!: " + err.Error())
			}
		}

	}

	ProducerCommand = &cobra.Command{
		Use:   "producer",
		Short: "Run migration producer.",
		Run:   ProducerHandler,
	}
)

func init() {
	rootCmd.AddCommand(ProducerCommand)
}
