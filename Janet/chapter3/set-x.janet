(defn set-x [& expressions]
    ~(upscope
        ,;(mapcat (fn [expression]
            [~(print ,(string/format "about to execute %q" expression))
                expression])
        expressions)))
