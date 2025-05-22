(defn element-to-struct [tag attrs children]
  {:tag tag :attrs (struct ;attrs) :children children})

(def html-peg (peg/compile
	       ~{:main (* :nodes -1)
		 :nodes (any (+ :element :text))
		 :element (unref
			   {:main (/ (* :open-tag (group :nodes) :close-tag) ,element-to-struct)
			    :open-tag (* "<" (<- :w+ :tag-name) (group (? (* :s+ :attributes))) ">")
			    :attributes
			      {:main (some (* :attribute (? :s+)))
			       :attribute (* (<- :w+) "=" :quoted-string)
			       :quoted-string (* `"` (<- (any (if-not `"` 1))) `"`)}
			    :close-tag (* "</" (backmatch :tag-name) ">")})
		 :text (<- (some (if-not "<" 1)))}))

(defn main [&]
  (def input (string/trim (file/read stdin :all)))
  (pp (peg/match html-peg input)))
