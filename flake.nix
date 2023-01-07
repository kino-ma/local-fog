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
        protoc = pkgs.protoc-gen-go;
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = [
            go
            protoc
          ];
        };

        apps.default = {
          type = "app";
          program = "go run .";
        };
      });
}
