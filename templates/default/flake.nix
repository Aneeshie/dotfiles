{
  description = "A basic development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
  };

  outputs = { self, nixpkgs }:
    let
      supportedSystems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forEachSupportedSystem = f: nixpkgs.lib.genAttrs supportedSystems (system: f {
        pkgs = import nixpkgs { inherit system; };
      });
    in
    {
      devShells = forEachSupportedSystem ({ pkgs }: {
        default = pkgs.mkShell {
          # The packages that will be available in the dev shell
          packages = with pkgs; [
            # Examples:
            # nodejs_20
            # python3
            # go
          ];

          # Environment variables can be set here
          # FOO = "bar";

          # Shell hook to run when entering the shell
          shellHook = ''
            echo "Welcome to the dev shell!"
          '';
        };
      });
    };
}
