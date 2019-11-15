(ns aoc2017.day3
  (:use clojure.test))

(def input 312051)

(defn abs [n] (max n (- n)))

; Part 1

(with-test
  (defn pos-to-coord [n]
    (loop [x 0 y 0
           dx 0 dy -1
           s 0]
      (if (< s (dec n))
        (if (or (= x y)
                (and (< x 0)
                     (= x (- y)))
                (and (> x 0)
                     (= x (- 1 y))))
          (recur (+ x (- dy)) (+ y dx) (- dy) dx (inc s))
          (recur (+ x dx) (+ y dy) dx dy (inc s)))
        (list x y))))
  (is (= '(0 0) (pos-to-coord 1)))
  (is (= '(1 0) (pos-to-coord 2)))
  (is (= '(1 1) (pos-to-coord 3)))
  (is (= '(0 1) (pos-to-coord 4)))
  (is (= '(2 1) (pos-to-coord 12)))
  (is (= '(0 -2) (pos-to-coord 23))))

(test #'pos-to-coord)

(with-test
  (defn distance-to-origin [p]
    (apply + (map abs p)))
  (is (= 0 (distance-to-origin '(0 0))))
  (is (= 3 (distance-to-origin '(2 1))))
  (is (= 2 (distance-to-origin '(0 -2))))
  (is (= 31 (distance-to-origin (pos-to-coord 1024)))))

(test #'distance-to-origin)

; Main

(defn -main []
  (println "Part 1:" (distance-to-origin (pos-to-coord input))))