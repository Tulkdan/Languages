(defn choose [rng selections]
  (def index (math/rng-int rng (length selections)))
  (in selections index))

(defn verbosify [rng word]
  (choose rng
	  (case word
	    "quick" ["alacritous" "expeditious"]
	    "lazy"  ["indolent" "lackadaisical" "languorous"]
	    "jumps" ["gambols"]
	    [word])))

(defn main [&]
  (def rng (math/rng (os/time)))
  (as-> stdin $
	(file/read $ :all)
	(string/split " " $)
	(map (partial verbosify rng) $)
	(string/join $ " ")
	(prin $)))
