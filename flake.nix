{
  inputs = {
    # nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    systems.url = "github:nix-systems/default";
  };

  outputs =
    { systems, nixpkgs, ... }@inputs:
    let
      eachSystem = f: nixpkgs.lib.genAttrs (import systems) (system: f nixpkgs.legacyPackages.${system});
    in
    {
      devShells = eachSystem (pkgs: {
        default =
          with pkgs;
          mkShell {
            packages =
              [
                go
                fyne
              ]
              ++ (
                if stdenv.isLinux then
                  [
                    libGL
                    pkg-config
                    xorg.libX11.dev
                    xorg.libXcursor
                    xorg.libXi
                    xorg.libXinerama
                    xorg.libXrandr
                    xorg.libXxf86vm
                    libxkbcommon
                    wayland
                  ]
                else
                  [ ]
              );
          };
      });
    };
}
