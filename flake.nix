{
  description = "Aneesh's darwin system";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-26.05-darwin";

    nix-darwin.url = "github:nix-darwin/nix-darwin/nix-darwin-26.05";
    nix-darwin.inputs.nixpkgs.follows = "nixpkgs";

    home-manager.url = "github:nix-community/home-manager/release-26.05";
    home-manager.inputs.nixpkgs.follows = "nixpkgs";

    nix-homebrew.url = "github:zhaofengli/nix-homebrew";
  };

  outputs = inputs@{ self, ... }: {
    templates = {
      default = {
        path = ./templates/default;
        description = "A basic development environment template";
      };
      rust = {
        path = ./templates/rust;
        description = "Rust development environment template (cargo, rustc, rustfmt, clippy, rust-analyzer)";
      };
      go = {
        path = ./templates/go;
        description = "Go development environment template (go, gopls, gotools, golangci-lint)";
      };
      python = {
        path = ./templates/python;
        description = "Python 3 development environment template (python3, pip, venv, pyright, black)";
      };
      java = {
        path = ./templates/java;
        description = "Java development environment template (jdk17, maven, gradle)";
      };
    };

    darwinConfigurations."mac" = inputs.nix-darwin.lib.darwinSystem {
      specialArgs = { inherit inputs; };
      modules = [
        ./configuration.nix

        inputs.nix-homebrew.darwinModules.nix-homebrew

        inputs.home-manager.darwinModules.home-manager
        {
          home-manager.useGlobalPkgs = true;
          home-manager.useUserPackages = true;

          home-manager.users.nara = import ./home.nix;
        }
      ];
    };
  };
}
