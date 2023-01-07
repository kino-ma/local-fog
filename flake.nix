{
  description = "LocalFog";

  inputs = {
    nixpkgs = { url = "github:NixOS/nixpkgs/nixpkgs-unstable"; };
    flake-utils = { url = "github:numtide/flake-utils"; };
  };

  outputs = { self, nixpkgs, flake-utils }: 
    flake-utils.lib.eachDefaultSystem (system:
      let
        inherit (nixpkgs.lib) optional;
        pkgs = import nixpkgs { inherit system; };

        go = pkgs.go;
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = [
            go
          ];
        };

        apps.default = {
          type = "app";
          program = "go run .";
        };
      });
}
