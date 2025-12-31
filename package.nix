{pkgs ? import <nixpkgs> {},...}:

pkgs.buildGoModule {
	pname = "twitter-bot";
	version = "1.0.0";

	src = ./. ;

	doCheck = false;

	vendorHash = "sha256-Upjt0Q2G6x5vGf0bG0TS9uWrHBow8/cQsZexhMgVb2I=";
}
