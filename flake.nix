{
  description = "Watches for new posts on Twitter using TwitterAPI.IO and show them on Discord";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
  };

  outputs =
    {
      flake-parts,
      ...
    }@inputs:
    flake-parts.lib.mkFlake { inherit inputs; } {

      systems = [ "x86_64-linux" ];

      perSystem =
        { pkgs, self', ... }:
        {

          packages.twitter-bot = pkgs.callPackage ./package.nix { };

          packages.default = self'.packages.twitter-bot;
        };

    };
}
