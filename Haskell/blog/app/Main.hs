module Main where

import qualified HsBlog
import OptParse

import Control.Exception (bracket)
import System.Exit (exitFailure)
import System.Directory (doesFileExist)
import System.IO

main :: IO ()
main = do
  options <- parse
  case options of
    ConvertDir input output ->
      HsBlog.convertDirectory input output

    ConvertSingle input output ->
      let
        withInputHandle :: (String -> Handle -> IO a) -> IO a
        withInputHandle action =
          case input of
            Stdin -> action "" stdin
            InputFile file ->
              bracket
                (openFile file ReadMode)
                hClose
                (action file)

        withOutputHandle :: (Handle -> IO a) -> IO a
        withOutputHandle action =
          case output of
            Stdout -> action stdout
            OutputFile file -> do
              exists <- doesFileExist file
              shouldOpenFile <-
                if exists
                  then confirm
                  else pure True
              if shouldOpenFile
                then
                  bracked
                    (openFile file WriteMode)
                    hClose
                    (action file)
                else exitFailure
      in
        withInputHandle (\title -> withOutputHandle . HsBlog.convertSingle title)


confirm :: IO Bool
confirm = do
  putStrLn "Are you sure? (y/n)"
  answer <- getLine
  case answer of
    "y" -> pure True
    "n" -> pure False
    _ ->
      putStrLn "Invalid response, use y or n" *>
      confirm
