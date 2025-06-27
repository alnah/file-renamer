package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		folder     = flag.String("folder", ".", "Target folder path")
		oldPattern = flag.String("from", "", "Old pattern (e.g., 'this_pattern_*.ext')")
		newPattern = flag.String("to", "", "New pattern (e.g., 'new_pattern_*.ext')")
		dryRun     = flag.Bool("dry-run", false, "Show what would be renamed without doing it")
		help       = flag.Bool("help", false, "Show usage")
	)
	flag.Parse()

	if *help || *oldPattern == "" || *newPattern == "" {
		showUsage()
		return
	}

	renamer := NewRenamer(*folder, *oldPattern, *newPattern, *dryRun)

	if err := renamer.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func showUsage() {
	fmt.Println("File Renamer CLI")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println(
		"  file-renamer -from 'this_pattern_*.ext' -to 'new_pattern_*.ext' [-folder path] [-dry-run]",
	)
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  file-renamer -from 'old_*' -to 'new_*'")
	fmt.Println(
		"  file-renamer -from 'this_pattern_*.txt' -to 'new_pattern_*.txt' -folder ./documents",
	)
	fmt.Println("  file-renamer -from 'img_*.jpg' -to 'photo_*.jpg' -dry-run")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
