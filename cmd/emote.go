package cmd

import (
	"github.com/spf13/cobra"

	"github.com/orginux/my-cli-tool/pkg/emoticons"
)

func init() {
	buildEmoteCommand()
}

func buildEmoteCommand() *cobra.Command {

	app, err := emoticons.New()
	cobra.CheckErr(err)

	var dest string

	emote := &cobra.Command{
		Use: "emote",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return app.Config.Load(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			emoticonName := args[0]
			app.Emote(emoticonName, dest)
		},
		Args: cobra.ExactArgs(1),
	}
	rootCmd.AddCommand(emote)

	emote.Flags().StringVar(&dest, "dest", "clipboard", "where to send your emoticon")

	return emote
}
