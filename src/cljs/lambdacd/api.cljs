(ns lambdacd.api
  (:require [lambdacd.ajax :as ajax]
            [lambdacd.route :as route])
  (:import goog.History))

(defn get-build-history []
  (let [result (ajax/GET "api/builds/")]
    result))

(defn get-build-state [build-number]
  (let [result (ajax/GET (str "api/builds/" build-number "/"))]
    result))

(defn- nop [response])

(defn- confirm-triggered [response]
  (js/alert "triggered"))

(defn- after-retriggered [response]
  (let [build-number (get response "build-number")]
    (route/set-build-number build-number)))


(defn trigger [trigger-id data]
  (ajax/POST (str "api/dynamic/" trigger-id) data confirm-triggered))

(defn retrigger [build-number step-id]
  (ajax/POST (str "api/builds/" build-number "/" step-id "/retrigger") {} after-retriggered))

(defn kill [build-number step-id]
  (ajax/POST (str "api/builds/" build-number "/" step-id "/kill") {} nop))