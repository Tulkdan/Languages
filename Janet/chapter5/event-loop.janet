(defn visualize-time []
  (var stopped false)
  (while (not stopped)
    (prin ".")
    (flush)
    (if (= (ev/sleep 0.1) :stop)
      (set stopped true))))

(def background-fiber (ev/call visualize-time))

(print "hello")
(ev/sleep 1)
(print "goodbye")

(ev/go background-fiber :stop)
