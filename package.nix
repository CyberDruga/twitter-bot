{
  pkgs ? import <nixpkgs> { },
  ...
}:

pkgs.buildGoModule {
  pname = "twitter-bot";
  version = "1.1.2";

  src = ./.;

  doCheck = false;

  vendorHash = "sha256-ObK67jv0DdEbITIUfHqr2hZxuVkZvs8rgJJUWwgHVdc=";
}
