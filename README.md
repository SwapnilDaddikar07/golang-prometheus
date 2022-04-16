# golang-prometheus

This is a simple integration with official prometheus golang client library to expose metrics.
Four endpoints are configred.

1. /all-success-simulator  (needs monitoring)
2. /internal-server-error-simulator (needs monitoring)
3. /slow-simulator (needs monitoring)
4. /healthz (needs to be excluded from monitoring)

A middleware for capturing metrics sits at the top of all routes.
The middleware accepts urls which need to be excluded from monitoring.
This is because , every application has some urls created for readiness/liveness probes which we dont intend to monitor.

The middleware captures metrics like
1. Number of requests received to the service.
2. Duration for each request (divided into custom buckets.)
3. Number of requests with status code 200.
4. Number of requests with status code 500.

There are 4 types of metrics.
1. Counter  (Used for naturally increasing values only)
2. Gauge   (Used for increasing and decreasing values)
3. Summary
4. Histogram  (Used to group data into buckets and primarily used for data like response time and response size)

The library provided wrappers which can be used to define metrics , register them automatically to default registerer.