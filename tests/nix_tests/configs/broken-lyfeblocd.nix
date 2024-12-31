{ pkgs ? import ../../../nix { } }:
let lyfeblocd = (pkgs.callPackage ../../../. { });
in
lyfeblocd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-lyfeblocd.patch
  ];
})
