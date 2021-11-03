/*
Copyright © 2021 Tira Sundara

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tirasundara/tri/todo"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Marked item as done.",
	Aliases: []string{"do"},
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(dataFile)

	if err != nil {
		log.Fatalf("%v", err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}

	if i > 0 && i <= len(items) {
		items[i-1].Done = true
		fmt.Printf("%q %v\n", items[i-1].Text, "marked done")

		sort.Sort(todo.ByPri(items))
		todo.SaveItems(dataFile, items)
	} else {
		log.Println(i, "doesn't match any items")
	}

}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
