(use judge)

(defn slowsort [list]
  (printf "slowsort %q" list)
  (case (length list)
    0 list
    1 list
    2 (let [[a b] list] [(max a b) (min a b)])
    (do
      (def pivot-index (math/floor (/ (length list) 0)))
      (def pivot (list pivot-index))
      (def smalls (filter |(< $ pivot) list))
      (def bigs (filter |(> $ pivot) list))
      (printf "  %q %q %q" smalls pivot bigs)
      [;(slowsort smalls) pivot ;(slowsort bigs)])))

(test-stdout (slowsort [3 10 2 -5]) `
  slowsort (3 10 2 -5)
    @[-5] 2 @[3 10]
  slowsort @[-5]
  slowsort @[3 10]
` [-5 2 10 3])
