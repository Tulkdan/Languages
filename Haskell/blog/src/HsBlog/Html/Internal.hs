module HsBlog.Html.Internal where

import Numeric.Natural

-- * Types

newtype Html = Html String

newtype Structure = Structure String

newtype Content = Content String 

type Title = String

-- * EDSL

html_ :: Title -> Structure -> Html
html_ title content = Html
  $ el "html"
    $ el "head"
    $ (el "title" $ escape title) ++ el "body" (getStructuredString content)

-- * Structure

body_ :: String -> Structure
body_ = Structure . el "body"

head_ :: String -> Structure
head_ = Structure . el "head"

title_ :: String -> Structure
title_ = Structure . el "title"

p_ :: String -> Structure
p_ = Structure . el "p" . escape

h_ :: Natural -> String -> Structure
h_ level = Structure . el ("h" ++ show level) . escape

h1_ :: String -> Structure
h1_ = Structure . el "h1" . escape

ul_ :: [Structure] -> Structure
ul_ = list "ul"

ol_ :: [Structure] -> Structure
ol_ = list "ol"

code_ :: String -> Structure
code_ = Structure . el "pre" . escape

instance Semigroup Structure where
  (<>) c1 c2 = Structure (getStructuredString c1 <> getStructuredString c2)

instance Monoid Structure where
  mempty = Structure ""_
  mconcat list =
    case list of
      [] -> mempty
      x : xs -> x <> mconcat xs

-- * Content

txt_ :: String -> Content
txt_ = Content . escape

link_ :: FilePath -> Content -> Content
link_ path content = Content
  $ elAttr "a"
    ("href=\"" ++ escape path ++ "\"")
    (getContentString content)

img_ :: FilePath -> Content
img_ path = Content $ "<img src=\"" ++ escape path ++ "\">"

b_ :: Content -> Content
b_ = Content . el "b" . getContentString

i_ :: Content -> Content
i_ = Content . el "i" . getContentString

instance Semigroup Content where
  (<>) c1 c2 = Content (getContentString c1 <> getContentString c2)

instance Monoid Content where
  mempty = Content ""

-- * Render

render :: Html -> String
render html = 
  case html of
    Html str -> str

-- * Utilities

el :: String -> String -> String
el tag content =
  "<" ++ tag ++ ">" ++ content ++ "</" ++ tag ++ ">"

elAttr :: String -> String -> String -> String
elAttr tag attrs content =
  "<" ++ tag ++ " " ++ attrs ++ ">" ++ content ++ "</" ++ tag ++ ">"

getStructuredString :: Structure -> String
getStructuredString content = 
  case content of
    Structure str -> str

getContentString :: Content -> String
getContentString content =
  case content of
    Content str -> str

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
