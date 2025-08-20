(lang dune 3.19)

(name they)

(generate_opam_files true)

(source
 (github clfs/they))

(authors "Calvin Figuereo-Supraner <mail@calvin.page>")

(maintainers "Calvin Figuereo-Supraner <mail@calvin.page>")

(license LICENSE)

(documentation https://github.com/clfs/they)

(package
 (name they)
 (synopsis "UCI chess engine")
 (description "A UCI chess engine written in OCaml.")
 (depends ocaml)
 (tags
  (chess)))
