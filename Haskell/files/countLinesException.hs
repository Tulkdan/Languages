import System.Environment
import System.IO.Error
import Control.Exception

main = toTry `catch` handler

toTry :: IO ()
toTry = do (fileName:_) <- getArgs
           contents <- readFile fileName
           putStrLn $ "The file has " ++ (show $ length $ lines contents) ++ " lines!"

handler :: IOError -> IO ()
handler e
  | isDoesNotExistError e =
      case ioeGetFileName e of Just name -> putStrLn ("Whoops! File does not exist at: " ++ name)
                               Nothing -> putStrLn "Whoops! File does not exist at unknown location!"
  | otherwise = ioError e

