package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Renamer struct {
	folder     string
	oldPattern string
	newPattern string
	dryRun     bool
}

func NewRenamer(folder, oldPattern, newPattern string, dryRun bool) *Renamer {
	return &Renamer{
		folder:     folder,
		oldPattern: oldPattern,
		newPattern: newPattern,
		dryRun:     dryRun,
	}
}

func (r *Renamer) Execute() error {
	// Convert glob pattern to regex for matching
	regex, err := r.patternToRegex(r.oldPattern)
	if err != nil {
		return fmt.Errorf("invalid pattern '%s': %v", r.oldPattern, err)
	}

	// Read directory
	entries, err := os.ReadDir(r.folder)
	if err != nil {
		return fmt.Errorf("cannot read folder '%s': %v", r.folder, err)
	}

	matches := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()
		if !regex.MatchString(filename) {
			continue
		}

		newName, err := r.generateNewName(filename, regex)
		if err != nil {
			fmt.Printf("Warning: cannot generate new name for '%s': %v\n", filename, err)
			continue
		}

		oldPath := filepath.Join(r.folder, filename)
		newPath := filepath.Join(r.folder, newName)

		if r.dryRun {
			fmt.Printf("Would rename: %s -> %s\n", filename, newName)
		} else {
			if err := os.Rename(oldPath, newPath); err != nil {
				fmt.Printf("Error renaming '%s' to '%s': %v\n", filename, newName, err)
				continue
			}
			fmt.Printf("Renamed: %s -> %s\n", filename, newName)
		}
		matches++
	}

	if matches == 0 {
		fmt.Printf("No files matching pattern '%s' found in '%s'\n", r.oldPattern, r.folder)
	} else {
		action := "renamed"
		if r.dryRun {
			action = "would be renamed"
		}
		fmt.Printf("\nTotal: %d files %s\n", matches, action)
	}

	return nil
}

// Convert glob-like pattern to regex
func (r *Renamer) patternToRegex(pattern string) (*regexp.Regexp, error) {
	// Escape special regex characters except *
	escaped := regexp.QuoteMeta(pattern)

	// Replace escaped \* with (.+) to capture the variable part
	regexPattern := strings.ReplaceAll(escaped, "\\*", "(.+)")

	// Anchor to match entire filename
	regexPattern = "^" + regexPattern + "$"

	return regexp.Compile(regexPattern)
}

// Generate new filename by replacing captured groups
func (r *Renamer) generateNewName(filename string, regex *regexp.Regexp) (string, error) {
	matches := regex.FindStringSubmatch(filename)
	if len(matches) < 2 {
		return "", fmt.Errorf("no wildcards captured")
	}

	// Count asterisks in new pattern
	asteriskCount := strings.Count(r.newPattern, "*")
	capturedCount := len(matches) - 1 // matches[0] is the full match

	if asteriskCount != capturedCount {
		return "", fmt.Errorf(
			"pattern mismatch: old pattern has %d wildcards but new pattern has %d",
			capturedCount,
			asteriskCount,
		)
	}

	result := r.newPattern
	for i := 1; i < len(matches); i++ {
		result = strings.Replace(result, "*", matches[i], 1)
	}

	return result, nil
}
