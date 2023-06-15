module Html.Internal where

-- * Types

newtype Html = Html String

newtype Structure = Structure String

type Title = String

-- * EDSL

html_ :: Title -> Structure -> Html
html_ title content = Html
  $ el "html"
    $ el "head"
    $ (el "title" $ escape title) ++ el "body" (getStructuredString content)

body_ :: String -> Structure
body_ = Structure . el "body"

head_ :: String -> Structure
head_ = Structure . el "head"

title_ :: String -> Structure
title_ = Structure . el "title"

p_ :: String -> Structure
p_ = Structure . el "p" . escape

h1_ :: String -> Structure
h1_ = Structure . el "h1" . escape

append_ :: Structure -> Structure -> Structure
append_ c1 c2 = Structure (getStructuredString c1 ++ getStructuredString c2)


ul_ :: [Structure] -> Structure
ul_ = list "ul"

ol_ :: [Structure] -> Structure
ol_ = list "ol"

code_ :: String -> Structure
code_ = Structure . el "pre" . escape

-- * Render

render :: Html -> String
render html = 
  case html of
    Html str -> str

-- * Utilities

el :: String -> String -> String
el tag content =
  "<" ++ tag ++ ">" ++ content ++ "</" ++ tag ++ ">"

getStructuredString :: Structure -> String
getStructuredString content = 
  case content of
    Structure str -> str

escape :: String -> String
escape =
  let
    escapeChar c =
      case c of
        '<' -> "&lt;"
        '>' -> "&gt;"
        '&' -> "&amp;"
        '"' -> "&quot;"
        '\'' -> "&#39;"
        _ -> [c]
  in
    concat . map escapeChar

list :: String -> [Structure] -> Structure
list listType = 
  Structure
  . el listType
  . concat
  . map (el "li" . getStructuredString)