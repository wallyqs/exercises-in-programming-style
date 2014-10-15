import qualified Control.Monad
import qualified Data.Char
import qualified Data.List
import qualified Data.Text
import qualified Data.Text.IO
import qualified System.Directory
import qualified System.FilePath
main :: IO ()
main = do
srcRoot <- System.Directory.getCurrentDirectory
let stopWordsPath          = System.FilePath.joinPath [srcRoot, "src", "exercises-in-programming-style", "stop_words.txt"]
let prideAndPrejudicePath  = System.FilePath.joinPath [srcRoot, "src", "exercises-in-programming-style", "pride-and-prejudice.txt"]
stopWordsFile <- Data.Text.IO.readFile stopWordsPath
let stopWords = Data.Text.split (== ',') stopWordsFile


prideAndPrejudiceFile <- Data.Text.IO.readFile prideAndPrejudicePath
let prideAndPrejudice =  filter (not . (`elem` stopWords)) $
                         filter ((> 1) . Data.Text.length) $
                         Data.Text.split (not . Data.Char.isAlphaNum) $
                         Data.Text.toLower prideAndPrejudiceFile


let groups = Data.List.group $
             Data.List.sort prideAndPrejudice

let frequencies = Data.List.sortBy (\(_, a) (_, b) -> compare b a) $
                  map (\xs@(x:_) -> (x, length xs)) groups


Control.Monad.forM_ (take 25 frequencies) $ \(w, f) ->
    putStrLn $ Data.Text.unpack w ++ "  -  " ++ show f
