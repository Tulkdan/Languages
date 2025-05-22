(def counter-prototype
  @{:add (fn [self amount] (+= (self :_count) amount))
    :increment (fn [self] (:add self 1))
    :count (fn [self] (self :_count))})

(defn new-counter []
  (table/setproto @{:_count 0} counter-prototype))

(def counter (new-counter))

(print (:count counter))
(:increment counter)
(print (:count counter))
(:add counter 3)
(print (:count counter))

(def Counter
  (let [proto @{:add (fn [self amount] (+= (self :_count) amount))
		:increment (fn [self] (:add self 1))
		:count (fn [self] (self :_count))}]
    (fn [] (table/setproto @{:_count 0} proto))))

(def counter (Counter))
(print (:count counter))
(:increment counter)
(print (:count counter))
(:add counter 3)
(print (:count counter))

