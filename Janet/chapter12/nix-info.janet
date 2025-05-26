#!/usr/bin/env janet
(use sh)
(import cmd)

(cmd/def "Print the description of Nix derivation."
 package (required ["<package>" :string])
 --name (flag) "Query by name instead of attribute")

(defn query-name [name]
  ($< nix-env -qa ,name --json --meta))

(defn query-attr [attr]
  ($< nix-env -qaA nixpkgs. ^ ,attr --json --meta))

($ <(if name (query-name package) (query-attr package))
   jq -r `.[] | .name + " " + .meta.description, "", (.meta.longDescription | rtrimstr("\n"))`)
