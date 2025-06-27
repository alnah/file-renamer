# File Renamer CLI

A command-line tool for batch renaming files using pattern matching with wildcard support.

## Features

- **Pattern Matching**: Use `*` wildcards to match and replace parts of filenames
- **Dry Run Mode**: Preview changes before applying them with `--dry-run`
- **Flexible Targeting**: Specify any folder with `--folder`
- **Zero Dependencies**: Built with Go standard library only
- **Cross-Platform**: Works on Linux, Windows, and macOS
- **Safe Operations**: Comprehensive validation and error handling

## Installation

### Option 1: Download Binary

Download the latest binary from the [releases page](https://github.com/alnah/file-renamer/releases).

### Option 2: Build from Source

```bash
git clone https://github.com/alnah/file-renamer.git
cd file-renamer
make build
```

### Option 3: Install with Go

```bash
go install github.com/alnah/file-renamer@latest
```

## Usage

### Basic Syntax

```bash
file-renamer -from 'old_pattern_*' -to 'new_pattern_*' [flags]
```

### Examples

#### Basic Renaming
```bash
# Rename files from "this_pattern_*.ext" to "new_pattern_*.ext"
file-renamer -from 'this_pattern_*.txt' -to 'new_pattern_*.txt'
```

#### Preview Changes (Dry Run)
```bash
# See what would be renamed without making changes
file-renamer -from 'img_*.jpg' -to 'photo_*.jpg' -dry-run
```

#### Target Specific Folder
```bash
# Rename files in a specific directory
file-renamer -from 'old_*.pdf' -to 'new_*.pdf' -folder ./documents
```

#### Multiple Wildcards
```bash
# Support for multiple wildcards
file-renamer -from 'backup_*_*.sql' -to 'archive_*_*.sql'
```

### Flags

| Flag | Description | Default |
|------|-------------|---------|
| `-from` | Source pattern with wildcards | Required |
| `-to` | Target pattern with wildcards | Required |
| `-folder` | Target directory path | `.` (current directory) |
| `-dry-run` | Preview changes without applying | `false` |
| `-help` | Show usage information | - |

## Pattern Matching

The tool uses `*` as wildcards that capture any sequence of characters:

- `file_*.txt` matches `file_1.txt`, `file_backup.txt`, etc.
- `*_backup.*` matches `data_backup.sql`, `config_backup.json`, etc.
- The captured parts are preserved in the new name

### Pattern Rules

1. **Same number of wildcards**: Source and target patterns must have the same number of `*` wildcards
2. **Order preservation**: Wildcards are replaced in the same order they appear
3. **Case sensitive**: Pattern matching respects filename case

## Examples in Action

### Before
```
project_files/
├── this_pattern_1.txt
├── this_pattern_2.txt
├── this_pattern_config.json
└── other_file.txt
```

### Command
```bash
file-renamer -from 'this_pattern_*' -to 'new_pattern_*' -folder project_files
```

### After
```
project_files/
├── new_pattern_1.txt
├── new_pattern_2.txt
├── new_pattern_config.json
└── other_file.txt
```

## Development

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Install locally
make install
```

### Testing

```bash
# Quick test with sample files
make test

# Manual testing
make run
```

### Formatting

```bash
make fmt
```

## Error Handling

The tool provides clear error messages for common issues:

- **Invalid patterns**: When source/target patterns don't match
- **Permission errors**: When files cannot be accessed or renamed
- **Missing files**: When no files match the specified pattern
- **Invalid directories**: When the target folder doesn't exist

## Limitations

- Only processes files in the specified directory (non-recursive)
- Requires exact wildcard count match between source and target patterns
- Does not handle special characters in filenames beyond basic wildcards

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Add tests if applicable
5. Commit your changes (`git commit -m 'feat: add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/alnah/file-renamer/blob/main/LICENCE) file for details.

## Support

- **Issues**: Report bugs on [GitHub Issues](https://github.com/alnah/file-renamer/issues)
- **Feature Requests**: Suggest new features via GitHub Issues
- **Documentation**: Check this README for usage examples

---

**Built with Go**
