{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    # systems.url = "github:nix-systems/default";
    utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      utils,
    }:
    utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs {
          inherit system;

          config.allowUnfree = true;
          config.android_sdk.accept_license = true;
        };
      in
      {
        devShells =
          let
            android-sdk =
              let
                androidComposition = pkgs.androidenv.composeAndroidPackages {
                  platformVersions = [ "latest" ];
                  abiVersions = [
                    "armeabi-v7a"
                    "arm64-v8a"
                  ];
                  toolsVersion = null;
                  includeNDK = true;
                };
              in
              androidComposition.androidsdk;
          in
          {
            default =
              with pkgs;
              mkShell {
                ANDROID_HOME = "${android-sdk}/libexec/android-sdk";
                packages =
                  [
                    android-sdk
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
          };
      }
    );
}
