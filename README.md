# VCC - NeoVim Color Scheme Changer

A fast, lightweight CLI tool for managing Neovim colorschemes with support for multiple popular themes and variants.

## Features

- **5 Popular Themes**: Catppuccin, Tokyo Night, Gruvbox, Rose Pine, and Kanagawa
- **Multiple Variants**: Each theme comes with its authentic variants
- **Fast Performance**: Built in Go with clean architecture
- **Custom Paths**: Support for custom configuration file paths
- **Zero Dependencies**: Works out of the box with any Neovim setup
- **Cross Platform**: Works on macOS, Linux, and Windows

## Installation

### Build from Source

```bash
git clone https://github.com/atasoya/vcc.git
cd vcc
go build
```

## Available Themes

| Theme           | Variants                          | Command          |
| --------------- | --------------------------------- | ---------------- |
| **Catppuccin**  | latte, frappe, macchiato, mocha   | `vcc catppuccin` |
| **Tokyo Night** | storm, night, day, moon           | `vcc tokyonight` |
| **Gruvbox**     | dark, light, hard, soft, material | `vcc gruvbox`    |
| **Rose Pine**   | main, moon, dawn                  | `vcc rosepine`   |
| **Kanagawa**    | wave, dragon, lotus               | `vcc kanagawa`   |

## Usage

### Basic Usage

```bash
# Apply default variant of a theme
vcc catppuccin
vcc tokyonight
vcc gruvbox

# Apply specific variant
vcc catppuccin -v mocha
vcc tokyonight -v storm
vcc kanagawa -v dragon
```

### Custom Configuration Path

```bash
# Use custom config file path
vcc catppuccin -v latte -p ~/.config/nvim/colors/my-theme.lua
vcc gruvbox --variant dark --path ./custom-colorscheme.lua
```

### Command Options

All theme commands support:

| Flag        | Short | Description             | Example                              |
| ----------- | ----- | ----------------------- | ------------------------------------ |
| `--variant` | `-v`  | Specify theme variant   | `-v mocha`                           |
| `--path`    | `-p`  | Custom config file path | `-p ~/.config/nvim/colors/theme.lua` |
| `--help`    | `-h`  | Show help for command   | `vcc catppuccin -h`                  |

## Architecture

VCC follows a clean, modular architecture:

```
vcc/
├── cmd/                   # Command definitions
│   ├── catppuccin.go      # Catppuccin theme command
│   ├── tokyonight.go      # Tokyo Night theme command
│   ├── gruvbox.go         # Gruvbox theme command
│   ├── rosepine.go        # Rose Pine theme command
│   ├── kanagawa.go        # Kanagawa theme command
│   └── root.go            # Root command
├── utils/                  # Utility functions
│   ├── variants.go        # Theme variant management
│   └── config/            # Configuration file writers
│       ├── common.go      # Unified path management
│       ├── catppuccinFile.go
│       ├── tokyonightFile.go
│       ├── gruvboxFile.go
│       ├── rosepineFile.go
│       └── kanagawaFile.go
└── main.go                # Application entry point
```

## Configuration

VCC writes Neovim configuration files in Lua format compatible with only lazy.nvim:

- [lazy.nvim](https://github.com/folke/lazy.nvim)

### Default Configuration Path

```
~/.config/nvim/lua/plugins/colorscheme.lua
```

### Example Generated Configuration

```lua
return {
  {
    "catppuccin/nvim",
    name = "catppuccin",
    priority = 1000,
    config = function()
      require("catppuccin").setup({ flavour = "mocha" })
      vim.cmd.colorscheme("catppuccin")
    end,
  },
}
```

## Development

### Adding a New Theme

1. **Add variants** in `utils/variants.go`:

```go
func NewThemeVariants() *VariantSet {
    return NewVariantSet([]string{"variant1", "variant2"})
}

func GetDefaultNewThemeVariant() string {
    return "variant1"
}
```

2. **Create config writer** in `utils/config/newthemeFile.go`:

```go
func WriteNewThemeConfig(variant string, customPath string) error {
    configPath := GetConfigPath(customPath)
    // Implementation here
}
```

3. **Add command** in `cmd/newtheme.go`:

```go
var newthemeCmd = &cobra.Command{
    Use:   "newtheme",
    Short: "Apply the New Theme to Neovim",
    // Implementation here
    // You can also use cobra-cli to create a new command
}
```
