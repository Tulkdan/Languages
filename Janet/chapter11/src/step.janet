(defn enemy-of-enemy [name people]
  (debug)
  (def subject (people name))
  (def nemesis (people (subject :nemesis)))
  (people (nemesis :nemesis)))

(defn main[&]
  (def people {"ian" {:age "young at heart"
		      :nemesis "jeffrey"}
	       "jeffrey" {:age 7.5
			  :nemesis "sarah"}})
  (print (enemy-of-enemy "ian" people)))
