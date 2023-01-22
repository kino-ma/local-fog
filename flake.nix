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
      in
      {
        devShell = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            python3
            iproute2
            protobuf
            protoc-gen-go
            protoc-gen-go-grpc

            python39Packages.matplotlib
          ];
        };
        
        # TODO: build outputs
      });
}
