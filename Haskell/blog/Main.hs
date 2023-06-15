module Main where

import Html

main :: IO ()
main = putStrLn
  $ render
  $ html_
    "My page title"
    ( append_
      (h1_ "Hello World!")
      ( append_
        (p_ "Paragraph #1")
        (p_ "Paragraph #2")
      )
    )
