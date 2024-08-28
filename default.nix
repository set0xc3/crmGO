let
  pkgs = import <nixpkgs> {};

  libs = with pkgs; [
    "/run/opengl-driver/lib"
    "/run/opengl-driver-32/lib"

    stdenv.cc.cc
    xorg.libX11

    libGL
  ];
in
  pkgs.mkShell {
    LD_LIBRARY_PATH = "${pkgs.lib.makeLibraryPath libs}";

    buildInputs = with pkgs; [
      jq
      go
      SDL2
      odin
    ];

    shellHook = ''
      export GOPATH="$HOME/go";
    '';
  }
