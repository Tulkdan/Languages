(def file (file/open "output.txt" :w))
(print "everything is normal")
(with-dyns [*out* file]
  (print "but this writes to a file"))
(print "back to normal")
