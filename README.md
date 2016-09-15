# prombench
Benchmark and integration testing tool for Prometheus

Prombench generates load with synthetic exporters and then queries Prometheus
to compare what we think we sent with what we hope was stored.

The eventual goal is to have Prombench monitor Prometheus internal metrics and
increase the load until Prometheus can't keep up, then throttle back until it's
just below the tipping point.  For now it just generates a static load.

# Usage

If the prometheus binary is in your PATH, you can run it without any arguments.
It won't run if the 'data' dir already exists, so either remove it yourself
between invocations or use the -rmdata option.

Each exporter generates 100 metrics with 100 labels, or 10000 real metrics.  I
can run about 30 such exporters on my system (8GB RAM, AMD FX(tm)-6300) at 1s
scrape interval for a total load of 300K samples/second.  Going above 300K
samples/second requires more RAM and will probably require tuning Prometheus,
see the tips in the [storage](https://prometheus.io/docs/operating/storage/)
part of the docs.  You can provide Prometheus command-line arguments at the end
of the prombench command line, after --, e.g.

    ./prombench -num-exporters 60 -- -storage.local.memory-chunks 2097152 -storage.local.max-chunks-to-persist 1048576 

# Exporters

The `-exporter` flag selects which load exporter to use.

The `inc` exporter increments the value of each metric on each scrape.

The `static` exporter exports unchanging metrics.

The `randcyclic` exporter exports semi-random values.

The `oscillate` exporter toggles between two sets of values on each cycle.
Unlike the others it doesn't actually go through the standard Prometheus client
library except during initialization, so it has much lower CPU needs.

# Dashboards

I've put up a [rudimentary dashboard](https://grafana.net/dashboards/445) at
Grafana.net which is helpful for seeing what's going on under the hood.
