(defn do-your-best []
  (error "oh no"))

(defn believe-in-yourself []
  (while true
    (do-your-best)))

(def fiber (fiber/new believe-in-yourself))
(resume fiber)
