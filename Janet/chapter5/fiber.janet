(defn print-something []
  (print "something"))

(def fiber (fiber/new print-something))
(resume fiber)

(defn range [count]
  (for i 0 count
    (yield i))
  "done")

(def fiber (fiber/new (fn [] (range 5))))

(each value fiber
  (print value))

(defn yield-twice [x]
  (yield x)
  (yield x))

(defn double-range [count]
  (for i 0 count
    (yield-twice i)))

(def fiber (fiber/new (fn [] (double-range 5))))
(each value fiber
  (print value))
