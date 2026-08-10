#!/usr/bin/env bash
set -euo pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"

ln -sfn "$DIR" ~/.dotfiles

if command -v darwin-rebuild >/dev/null 2>&1; then
    exec sudo darwin-rebuild switch --flake ~/.dotfiles#mac
else
    echo "==> darwin-rebuild not found. Bootstrapping nix-darwin for the first time..."
    exec nix run nix-darwin#darwin-rebuild -- switch --flake ~/.dotfiles#mac
fi

