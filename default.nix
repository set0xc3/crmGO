let
  pkgs = import <nixpkgs> {};

  libs = with pkgs; [
  ];
in
  pkgs.mkShell {
    LD_LIBRARY_PATH = "${pkgs.lib.makeLibraryPath libs}";

    buildInputs = with pkgs; [
      jq
      templ
      go
      SDL2
      odin
    ];

    shellHook = ''
      export GOPATH="$HOME/go";
    '';
  }
