{
  description = "Watches for new posts on Twitter using TwitterAPI.IO and show them on Discord";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
  };

  outputs = { self, nixpkgs }: {

    packages.x86_64-linux.twitter-bot = nixpkgs.legacyPackages.x86_64-linux.callPackage ./package.nix {} ;

    packages.x86_64-linux.default = self.packages.x86_64-linux.twitter-bot;

  };
}
