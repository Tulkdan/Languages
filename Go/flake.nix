{
  description = "A Nix-flake-based Elixir development environment";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        inherit (pkgs.lib) optional optionals;

        pkgs = import nixpkgs { inherit system; };
      in
        {
          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              libgccjit
              go
              golangci-lint
            ];
          };
        }
    );
}
