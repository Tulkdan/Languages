(defn print-dots []
  (while true
    (prin ".")
    (flush)
    (ev/sleep 0)))

(ev/call print-dots)

(def f (os/open "lorem-ipsum.txt" :r))
(print "About to read")
(def bytes (ev/read f 10))
(print "Done reading")
(ev/close f)
(print "Done closing the file")
(printf "read %q" bytes)
(os/exit 0)
