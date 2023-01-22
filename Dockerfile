FROM nixos/nix:master

WORKDIR /app

COPY nix.conf /etc/nix/nix.conf
COPY flake.nix flake.lock /app/
RUN nix develop -c true
COPY go.mod go.sum /app/
RUN nix develop -c sh -c "\
        go mod download && \
        function sudo(){ \$@; } && \
        export -f sudo && \
        cd \$(go env GOPATH)/pkg/mod/gocv.io/x/gocv@v0.32.1 && \
        make install"