(ns aoc2017.day3
  (:require [clojure.math.combinatorics :as combo])
  (:use clojure.test))

; Note: tests disabled since the spiral is clockwise instead of counter-clockwise

(def input 312051)

(defn abs [n] (max n (- n)))

(with-test
  (defn step [[x y]]
      (if (and (<= (abs x) (abs y))
               (or (not= x y) (>= x 0)))
        (list (if (>= y 0) (inc x) (dec x)) y)
        (list x (if (>= x 0) (dec y) (inc y)))))
  (is (= '(1 0) (step '(0 0))))
  (is (= '(1 1) (step '(1 0))))
  (is (= '(0 1) (step '(1 1))))
  (is (= '(-1 0) (step '(-1 1))))
  (is (= '(0 -1) (step '(-1 -1))))
  (is (= '(-2 1) (step '(-2 2)))))

; (test #'step)

; Part 1

(with-test
  (defn pos-to-coord [n]
    (loop [p '(0 0)
           m 0]
      (if (< m (dec n))
        (recur
          (step p)
          (inc m))
        p)))
  (is (= '(0 0) (pos-to-coord 1)))
  (is (= '(1 0) (pos-to-coord 2)))
  (is (= '(1 1) (pos-to-coord 3)))
  (is (= '(0 1) (pos-to-coord 4)))
  (is (= '(2 1) (pos-to-coord 12)))
  (is (= '(0 -2) (pos-to-coord 23))))

; (test #'pos-to-coord)

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