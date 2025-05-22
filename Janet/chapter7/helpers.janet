(def- upcase string/ascii-upper)

(defn shout [x]
  (printf "%s!"
	  (upcase x)))
