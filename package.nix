{
  pkgs ? import <nixpkgs> { },
  ...
}:

pkgs.buildGoModule {
  pname = "twitter-bot";
  version = "1.1.2";

  src = ./.;

  doCheck = false;

  vendorHash = "sha256-mXJv30sdkG6bGGm2MzHPShVogJe2TmYYCkyNKw87vN4=";
}
