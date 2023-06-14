module Main where

newtype Html = Html String

newtype Structure = Structure String

type Title = String

el :: String -> String -> String
el tag content =
  "<" ++ tag ++ ">" ++ content ++ "</" ++ tag ++ ">"

html_ :: Title -> Structure -> Html
html_ title content = Html
  $ el "html"
    $ el "head"
    $ (el "title" title) ++ el "body" (getStructuredString content)

body_ :: String -> Structure
body_ = Structure . el "body"

head_ :: String -> Structure
head_ = Structure . el "head"

title_ :: String -> Structure
title_ = Structure . el "title"

p_ :: String -> Structure
p_ = Structure . el "p"

h1_ :: String -> Structure
h1_ = Structure . el "h1"

append_ :: Structure -> Structure -> Structure
append_ c1 c2 = Structure (getStructuredString c1 ++ getStructuredString c2)

render :: Html -> String
render html = 
  case html of
    Html str -> str

getStructuredString :: Structure -> String
getStructuredString content = 
  case content of
    Structure str -> str

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
