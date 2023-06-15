module Main where

import Html

main :: IO ()
main = putStrLn
  $ render
  $ html_
    "My page title"
    (
      (h1_ "Hello World!") <>
      (
        (p_ "Paragraph #1") <>
        (p_ "Paragraph #2")
      )
    )
