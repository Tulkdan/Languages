(defn sequence [& fs]
    (fn [&]
        (each f fs
            (f))))

(defn first-word []
    (prin "hello"))

(defn space []
    (prin " "))

(defn second-word []
    (prin "world"))

(defn newline []
    (print))

(def main
    (sequence first-word space second-word newline))
