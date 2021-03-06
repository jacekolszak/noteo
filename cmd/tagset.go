package cmd

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
)

func tagSet() *cobra.Command {
	var (
		stdin bool
		name  string
	)
	tagSet := &cobra.Command{
		Use:     "set",
		Short:   "Set a tag on notes",
		Aliases: []string{"add", "update"},
		RunE: func(cmd *cobra.Command, args []string) error {
			if name == "" {
				return fmt.Errorf("no name given using -n flag")
			}
			repo, err := workingDirRepository()
			if err != nil {
				return err
			}
			tagFileWith := func(file, name string) {
				ok, err := repo.TagFileWith(file, name)
				if err != nil {
					_, _ = fmt.Fprintln(cmd.ErrOrStderr(), "skipping:", err)
				} else if ok {
					printer := NewPrinter()
					printer.PrintFile(file)
					printer.Println(" updated")
				}
			}
			if stdin {
				scanner := bufio.NewScanner(cmd.InOrStdin())
				scanner.Split(bufio.ScanLines)
				for scanner.Scan() {
					file := scanner.Text()
					tagFileWith(file, name)
				}
				return nil
			}
			for _, file := range args {
				tagFileWith(file, name)
			}
			return nil
		},
	}
	tagSet.Flags().BoolVarP(&stdin, "stdin", "", false, "read file names from standard input")
	tagSet.Flags().StringVarP(&name, "name", "n", "", "short name without space. Can have form of name:number-or-date")
	return tagSet
}
