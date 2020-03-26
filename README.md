# KubeSphere Monitoring Dashboard

The project is inspired by [Grafana](http://grafana.com/) with significant difference in terms of data storage, multi-tenancy and dashboard sharing to fit KubeSphere's context. But it is not a replacement to Grafana. It requires KubeSphere backend and frontend to work.

This repository is aimed at KubeSphere developers who want to understand dashboard data model, storage, template sharing and how to contribute, as the custom monitoring feature is introduced since v3.0.    

# Get Started

KubeSphere custom monitoring functions work out of the box. Follow the installation guide, if you don't have KubeSphere installed, and view the demo below:

TODO(@FeynmanZhou)

The stack that makes these possible includes KubeSphere [backend](https://github.com/kubesphere/kubesphere), [console](https://github.com/kubesphere/console) and custom resources for dashboards. 

- backend: enforce multi-tenancy, allow dashboard template sharing, perform metrics query
- console: render charts by decoding dashboard data
- custome resource: store dashboard data

# Data Model

Dashboards are backed by the custom resource definition (CRD) `Dashboard`, which defines metadata and properties of a dashboard, such as time range, panels, template variables, queries in panel, etc. When a user creates a new dashboard, a new `Dashboard` custom resource is instantiated with typical content as the example presents:

```yaml
apiVersion: monitoring.kubesphere.io/v1alpha1
kind: Dashboard
metadata:
  name: mysql-overview
  namespace: demo
spec:
  title: MySQL Overview
  description: MySQL dashboard for the mysql exporter
  time:
    from: now-1h
    to: now
  datasource: prometheus
  panels:
  - type: singlestat
    title": Instance Up
    targets:
    - expr: mysql_up{release="$release"}
      instant: "true"
  - type: graph
    title: mysql  disk reads vs writes
    targets:
    - expr": irate(mysql_global_status_innodb_data_reads{release="$release"}[10m])
      legendFormat": reads
    - expr": irate(mysql_global_status_innodb_data_writes{release="$release"}[10m])
      legendFormat": write
  templatings:
  - name: release
    query: label_values(mysql_up,release)
    type: query
    sort: 0
```

> Note that the data model is heavily inspired by [Grafana JSON model](https://grafana.com/docs/grafana/latest/reference/dashboard/) to gain compatibility. However, to adapt it to KubeSphere's context, we may bring in new fields and break changes.

## Panel

Four types of panels are supported so far:   

- row: groups relevant panels
- singlestat: shows a single stat value
- table: displays ranking
- graph: shows graph for series data

Except for row, the other three have features in common, which are panel title, query expressions, etc., from base panel. The `targets` field in panel is used for fetching values for a query. See data model documentation for more information.

## Metadata

Templating defines an array of template variables, so that you don't need to hard code things like host, app, job name in your panel queries, which makes your dashboard portable.

TODO: a snapshot for templating settings

Time range defines default time span of charts for series data on display. Time range is a dashboard-level setting.

# Dashboard Sharing

Any dashboard custom resource with the label-value pair `dashboard.monitoring.kubesphere.io/template: true` should be assumed as dashboard templates which can be shared to others. The difference between dashboards and dashboard templates is that the latter removes sensitive information and can be exported into other namespaces. It is the responsibility of dashboard template owners to secure the content before sharing.

There are two types of dashboard templates:

- cluster-scoped: dashboard templates in the `kube-public` namespace
- namespace-scoped: those in user-owned namespaces

Importing a dashboard template into your namespace is as easy as duplicating the dashboard template custom resource object, removing the label `dashboard.monitoring.kubesphere.io/template`.

# Contributing

You can contribute the project in three ways:

## Backend

If you find some fields should be included in the CRD, edit [api package](https://github.com/kubesphere/monitoring-dashboard/tree/master/api/v1alpha1) and regenerate the project by `make`. Kubebuilder is the tool we are using.

If you find bugs or want to add new APIs, implement new datasources to KubeSphere, read KubeSphere [developer guides](https://github.com/kubesphere/community/tree/master/developer-guide/concepts-and-designs) for monitoring.

## Frontend

@TODO(justahole)

## Dashboard Gallery

If you want to share your dashboard templates, you can submit dashboard template yaml files to gallery folder with an elaborate readme. Outstanding templates will be selected to ship with KubeSphere future releases.