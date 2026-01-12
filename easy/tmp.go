{
	"annotations": {
	  "list": [
		{
		  "builtIn": 1,
		  "datasource": {
			"type": "grafana",
			"uid": "-- Grafana --"
		  },
		  "enable": true,
		  "hide": true,
		  "iconColor": "rgba(0, 211, 255, 1)",
		  "name": "Annotations & Alerts",
		  "target": {
			"limit": 100,
			"matchAny": false,
			"tags": [],
			"type": "dashboard"
		  },
		  "type": "dashboard"
		}
	  ]
	},
	"editable": true,
	"fiscalYearStartMonth": 0,
	"graphTooltip": 0,
	"id": 742,
	"links": [],
	"liveNow": false,
	"panels": [
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"gridPos": {
		  "h": 3,
		  "w": 24,
		  "x": 0,
		  "y": 0
		},
		"id": 16,
		"options": {
		  "includeVars": false,
		  "keepTime": false,
		  "maxItems": 10,
		  "query": "",
		  "showFolderNames": true,
		  "showHeadings": false,
		  "showRecentlyViewed": false,
		  "showSearch": true,
		  "showStarred": false,
		  "tags": [
			"RecomMerger"
		  ]
		},
		"pluginVersion": "11.1.0",
		"title": "Dasboards",
		"type": "dashlist"
	  },
	  {
		"alert": {
		  "alertRuleTags": {},
		  "conditions": [
			{
			  "evaluator": {
				"params": [
				  1
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "and"
			  },
			  "query": {
				"params": [
				  "B",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "max"
			  },
			  "type": "query"
			}
		  ],
		  "executionErrorState": "keep_state",
		  "for": "1m",
		  "frequency": "1m",
		  "handler": 1,
		  "name": "Response 0.95 percentile alert",
		  "noDataState": "keep_state",
		  "notifications": [
			{
			  "uid": "Q9eA0f84z"
			}
		  ]
		},
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 7,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				}
			  ]
			},
			"unit": "s"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 24,
		  "x": 0,
		  "y": 3
		},
		"id": 10,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last",
			  "max",
			  "mean"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "histogram_quantile(0.99, rate(wbx_catalog_proxy_latency_bucket{cluster=~\"${cluster}\",kubernetes_namespace=~\"${namespace}\", path=~\"/personal.*|/recom.*|/interests.*|/visual.*\", kubernetes_pod_name=~\"${podName}\"}[$__rate_interval]))",
			"hide": true,
			"instant": false,
			"legendFormat": "{{kubernetes_pod_name}}-{{cluster}} {{path}}",
			"range": true,
			"refId": "A"
		  },
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "histogram_quantile(0.95, rate(wbx_catalog_proxy_latency_bucket{cluster=~\"k8s.elrec|k8s.xcrec\",kubernetes_namespace=~\"rec-pers-online\", app=\"recom-merger\"}[$__rate_interval]))",
			"hide": true,
			"instant": false,
			"legendFormat": "{{kubernetes_pod_name}} {{path}}",
			"range": true,
			"refId": "B"
		  },
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "max by (kubernetes_pod_name, cluster) (histogram_quantile(0.99, rate(wbx_catalog_proxy_latency_bucket{cluster=~\"${cluster}\",kubernetes_namespace=~\"${namespace}\", path=~\"/personal.*|/recom.*|/interests.*|/visual.*\", kubernetes_pod_name=~\"${podName}\"}[$__rate_interval])))",
			"hide": false,
			"instant": false,
			"legendFormat": "{{kubernetes_pod_name}}-{{cluster}}",
			"range": true,
			"refId": "C"
		  }
		],
		"thresholds": [
		  {
			"colorMode": "critical",
			"op": "gt",
			"value": 1,
			"visible": true
		  }
		],
		"title": "Response 0.99 percentile",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"alert": {
		  "alertRuleTags": {},
		  "conditions": [
			{
			  "evaluator": {
				"params": [
				  20
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "and"
			  },
			  "query": {
				"params": [
				  "B",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "last"
			  },
			  "type": "query"
			}
		  ],
		  "executionErrorState": "keep_state",
		  "for": "1m",
		  "frequency": "1m",
		  "handler": 1,
		  "name": "Status code alert",
		  "noDataState": "keep_state",
		  "notifications": [
			{
			  "uid": "Q9eA0f84z"
			}
		  ]
		},
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 7,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			}
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 24,
		  "x": 0,
		  "y": 10
		},
		"id": 2,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last",
			  "max",
			  "mean"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum by (code) (rate(wbx_catalog_proxy_responses_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\",path=~\"/personal.*|/recom.*|/interests.*|/visual.*\", kubernetes_pod_name=~\"${podName}\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "{{code}}",
			"range": true,
			"refId": "A"
		  }
		],
		"thresholds": [
		  {
			"colorMode": "critical",
			"op": "gt",
			"value": 20,
			"visible": true
		  }
		],
		"title": "Status code",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 7,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			}
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 24,
		  "x": 0,
		  "y": 17
		},
		"id": 12,
		"interval": "30",
		"options": {
		  "legend": {
			"calcs": [
			  "last",
			  "max",
			  "mean"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum by (app, exp, group) (rate(wbx_catalog_proxy_ab_responses_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\", kubernetes_pod_name=~\"${podName}\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "{{exp}}_{{group}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "AB RPS",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"alert": {
		  "alertRuleTags": {},
		  "conditions": [
			{
			  "evaluator": {
				"params": [
				  0
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "and"
			  },
			  "query": {
				"params": [
				  "A",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "last"
			  },
			  "type": "query"
			}
		  ],
		  "executionErrorState": "keep_state",
		  "for": "1m",
		  "frequency": "1m",
		  "handler": 1,
		  "name": "Restarts alert",
		  "noDataState": "keep_state",
		  "notifications": [
			{
			  "uid": "Q9eA0f84z"
			}
		  ]
		},
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 7,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			}
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 24,
		  "x": 0,
		  "y": 24
		},
		"id": 11,
		"interval": "30",
		"options": {
		  "legend": {
			"calcs": [
			  "last",
			  "max",
			  "mean"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "increase(kube_pod_container_status_restarts_total{cluster=~\"${cluster}\", namespace=~\"${namespace}\", pod=~\"${podName}\"}[5m])",
			"hide": false,
			"instant": false,
			"legendFormat": "{{pod}}-{{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"thresholds": [
		  {
			"colorMode": "critical",
			"op": "gt",
			"value": 0,
			"visible": true
		  }
		],
		"title": "Restarts",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 7,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "reqps"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 24,
		  "x": 0,
		  "y": 31
		},
		"id": 9,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last",
			  "max",
			  "mean"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "fdiinprxonta8a"
			},
			"editorMode": "code",
			"expr": "sum(rate(wbx_catalog_proxy_responses_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\", path=~\"/personal.*|/recom.*\", kubernetes_pod_name=~\"${podName}\"}[1m])) by (kubernetes_pod_name, cluster)",
			"hide": false,
			"legendFormat": "{{kubernetes_pod_name}}.{{cluster}}",
			"queryType": "randomWalk",
			"range": true,
			"refId": "A"
		  },
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum(rate(wbx_catalog_proxy_responses_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\", path=~\"/personal.*|/recom.*|/interests.*\", kubernetes_pod_name=~\"${podName}\"}[1m])) by (kubernetes_pod_name, cluster)",
			"hide": true,
			"instant": false,
			"legendFormat": "{{kubernetes_pod_name}}.{{cluster}}",
			"range": true,
			"refId": "B"
		  }
		],
		"title": "RPS",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"alert": {
		  "alertRuleTags": {},
		  "conditions": [
			{
			  "evaluator": {
				"params": [
				  1677721600
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "and"
			  },
			  "query": {
				"params": [
				  "B",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "last"
			  },
			  "type": "query"
			}
		  ],
		  "executionErrorState": "keep_state",
		  "for": "2m",
		  "frequency": "1m",
		  "handler": 1,
		  "name": "Memory Usage GiB alert",
		  "noDataState": "keep_state",
		  "notifications": [
			{
			  "uid": "Q9eA0f84z"
			}
		  ]
		},
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "line"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				}
			  ]
			},
			"unit": "bytes"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 8,
		  "w": 13,
		  "x": 0,
		  "y": 38
		},
		"id": 6,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "container_memory_working_set_bytes{cluster=~\"${cluster}\", namespace=~\"${namespace}\", pod=~\"${podName}\"}",
			"hide": false,
			"instant": false,
			"legendFormat": "{{pod}} {{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"thresholds": [
		  {
			"colorMode": "critical",
			"op": "gt",
			"value": 1677721600,
			"visible": true
		  }
		],
		"title": "Memory Usage GiB",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"alert": {
		  "alertRuleTags": {},
		  "conditions": [
			{
			  "evaluator": {
				"params": [
				  2.4
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "and"
			  },
			  "query": {
				"params": [
				  "B",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "last"
			  },
			  "type": "query"
			}
		  ],
		  "executionErrorState": "keep_state",
		  "for": "2m",
		  "frequency": "1m",
		  "handler": 1,
		  "name": "CPU Usage alert",
		  "noDataState": "keep_state",
		  "notifications": [
			{
			  "uid": "Q9eA0f84z"
			}
		  ]
		},
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "line"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 8,
		  "w": 11,
		  "x": 13,
		  "y": 38
		},
		"id": 4,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "rate(container_cpu_usage_seconds_total{cluster=~\"${cluster}\", namespace=~\"${namespace}\", pod=~\"${podName}\"}[5m])",
			"hide": false,
			"instant": false,
			"legendFormat": "{{pod}} {{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"thresholds": [
		  {
			"colorMode": "critical",
			"op": "gt",
			"value": 2.4,
			"visible": true
		  }
		],
		"title": "CPU Usage",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"alert": {
		  "alertRuleTags": {},
		  "conditions": [
			{
			  "evaluator": {
				"params": [
				  10
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "and"
			  },
			  "query": {
				"params": [
				  "C",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "last"
			  },
			  "type": "query"
			},
			{
			  "evaluator": {
				"params": [
				  0
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "or"
			  },
			  "query": {
				"params": [
				  "D",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "last"
			  },
			  "type": "query"
			}
		  ],
		  "executionErrorState": "keep_state",
		  "for": "1m",
		  "frequency": "1m",
		  "handler": 1,
		  "name": "Errors alert",
		  "noDataState": "keep_state",
		  "notifications": [
			{
			  "uid": "Q9eA0f84z"
			}
		  ]
		},
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 13,
		  "x": 0,
		  "y": 46
		},
		"id": 7,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"exemplar": true,
			"expr": "sum by (error, cluster) (rate(wbx_catalog_proxy_errors_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\", kubernetes_pod_name=~\"${podName}\", error!=\"user-info: status code is not 200: 204\",error!=\"user-info: status code is not 200: 403\",error!=\"user-info: status code is not 200: 500\", error!=\"request exact: similar, err: error while parsing binary response\", error!=\"request exact: related, err: error while parsing binary response\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "{{error}} {{cluster}}",
			"range": true,
			"refId": "A"
		  },
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum(rate(wbx_catalog_proxy_panic_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "panic",
			"range": true,
			"refId": "B"
		  }
		],
		"thresholds": [
		  {
			"colorMode": "critical",
			"op": "gt",
			"value": 10,
			"visible": true
		  }
		],
		"title": "Errors",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"alert": {
		  "alertRuleTags": {},
		  "conditions": [
			{
			  "evaluator": {
				"params": [
				  10
				],
				"type": "gt"
			  },
			  "operator": {
				"type": "and"
			  },
			  "query": {
				"params": [
				  "A",
				  "1m",
				  "now"
				]
			  },
			  "reducer": {
				"params": [],
				"type": "last"
			  },
			  "type": "query"
			}
		  ],
		  "executionErrorState": "keep_state",
		  "for": "1m",
		  "frequency": "1m",
		  "handler": 1,
		  "name": "Empty responses alert",
		  "noDataState": "keep_state",
		  "notifications": [
			{
			  "uid": "Q9eA0f84z"
			}
		  ]
		},
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 8,
		  "w": 11,
		  "x": 13,
		  "y": 46
		},
		"id": 8,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum(rate(wbx_catalog_proxy_empty_responses_total{cluster=~\"${cluster}\", kubernetes_pod_name=~\"${podName}\"}[1m])) by (type, cluster)",
			"hide": false,
			"instant": false,
			"legendFormat": "empty responses {{type}} {{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"thresholds": [
		  {
			"colorMode": "critical",
			"op": "gt",
			"value": 10,
			"visible": true
		  }
		],
		"title": "Empty responses",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 13,
		  "x": 0,
		  "y": 53
		},
		"id": 18,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"exemplar": true,
			"expr": "sum by (cluster, kubernetes_pod_name) (rate(wbx_catalog_proxy_errors_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\", kubernetes_pod_name=~\"${podName}\", error!=\"user-info: status code is not 200: 204\",error!=\"user-info: status code is not 200: 403\",error!=\"user-info: status code is not 200: 500\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "{{kubernetes_pod_name}}.{{cluster}} ",
			"range": true,
			"refId": "A"
		  },
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum(rate(wbx_catalog_proxy_panic_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "panic",
			"range": true,
			"refId": "B"
		  }
		],
		"title": "Errors",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 8,
		  "w": 11,
		  "x": 13,
		  "y": 54
		},
		"id": 19,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last",
			  "mean",
			  "max"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum(rate(wbx_catalog_proxy_action_default_requests_counter{cluster=~\"${cluster}\"})) by  (cluster)",
			"hide": false,
			"instant": false,
			"legendFormat": "action default requests {{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "Action default requests",
		"transformations": [
		  {
			"id": "calculateField",
			"options": {}
		  }
		],
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 13,
		  "x": 0,
		  "y": 60
		},
		"id": 17,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"exemplar": true,
			"expr": "sum by (error, cluster) (rate(wbx_catalog_proxy_errors_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\", kubernetes_pod_name=~\"${podName}\", error=~\"user-info.*\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "{{error}} {{cluster}}",
			"range": true,
			"refId": "A"
		  },
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum(rate(wbx_catalog_proxy_panic_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "panic",
			"range": true,
			"refId": "B"
		  }
		],
		"title": "Errors user info",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 9,
		  "w": 11,
		  "x": 13,
		  "y": 62
		},
		"id": 14,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "wbx_catalog_proxy_retry_pool{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\",kubernetes_pod_name=~\"${podName}\"}",
			"hide": false,
			"instant": false,
			"legendFormat": "{{kubernetes_pod_name}}-{{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "Retry pool by pod",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 7,
		  "w": 13,
		  "x": 0,
		  "y": 67
		},
		"id": 13,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"exemplar": true,
			"expr": "sum by (app, cluster) (rate(wbx_catalog_proxy_retry_counter{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\"}[1m]))",
			"hide": false,
			"instant": false,
			"legendFormat": "Retry {{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "RPS Retries",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			}
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 9,
		  "w": 11,
		  "x": 13,
		  "y": 71
		},
		"id": 22,
		"interval": "30",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"exemplar": true,
			"expr": "sum(rate(wbx_catalog_proxy_partition_types_gauge{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\",kubernetes_pod_name=~\"${podName}\"}[5m])) by (run, platform, type)",
			"instant": false,
			"legendFormat": "{{platform}}-{{type}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "RPS Types",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "smooth",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			},
			"unit": "none"
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 6,
		  "w": 13,
		  "x": 0,
		  "y": 74
		},
		"id": 15,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Last",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"exemplar": true,
			"expr": "100 * sum(increase(container_cpu_cfs_throttled_periods_total{cluster=~\"${cluster}\",pod=~\"${podName}\",namespace=~\"${namespace}\"}[5m])) by (pod, namespace, cluster) / sum(increase(container_cpu_cfs_periods_total{cluster=~\"${cluster}\",pod=~\"${podName}\",namespace=~\"${namespace}\"}[5m])) by (pod, namespace, cluster)",
			"hide": false,
			"instant": false,
			"legendFormat": "{{pod}}-{{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "CPU Throttling",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 0,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			}
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 6,
		  "w": 24,
		  "x": 0,
		  "y": 80
		},
		"id": 21,
		"interval": "30",
		"options": {
		  "legend": {
			"calcs": [
			  "last"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true,
			"sortBy": "Name",
			"sortDesc": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"exemplar": true,
			"expr": "sum by (error, cluster) (rate(wbx_catalog_proxy_errors_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\", kubernetes_pod_name=~\"${podName}\", error=~\".*binary.*\"}[1m]))",
			"instant": false,
			"legendFormat": "{{error}} {{cluster}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "Errors binary",
		"transparent": true,
		"type": "timeseries"
	  },
	  {
		"datasource": {
		  "type": "prometheus",
		  "uid": "PDCCC65849FA2FB5B"
		},
		"fieldConfig": {
		  "defaults": {
			"color": {
			  "mode": "palette-classic"
			},
			"custom": {
			  "axisBorderShow": false,
			  "axisCenteredZero": false,
			  "axisColorMode": "text",
			  "axisLabel": "",
			  "axisPlacement": "auto",
			  "barAlignment": 0,
			  "drawStyle": "line",
			  "fillOpacity": 7,
			  "gradientMode": "none",
			  "hideFrom": {
				"legend": false,
				"tooltip": false,
				"viz": false
			  },
			  "insertNulls": false,
			  "lineInterpolation": "linear",
			  "lineWidth": 2,
			  "pointSize": 5,
			  "scaleDistribution": {
				"type": "linear"
			  },
			  "showPoints": "auto",
			  "spanNulls": false,
			  "stacking": {
				"group": "A",
				"mode": "none"
			  },
			  "thresholdsStyle": {
				"mode": "off"
			  }
			},
			"mappings": [],
			"thresholds": {
			  "mode": "absolute",
			  "steps": [
				{
				  "color": "green"
				},
				{
				  "color": "red",
				  "value": 80
				}
			  ]
			}
		  },
		  "overrides": []
		},
		"gridPos": {
		  "h": 8,
		  "w": 13,
		  "x": 0,
		  "y": 86
		},
		"id": 20,
		"interval": "30s",
		"options": {
		  "legend": {
			"calcs": [
			  "last",
			  "max",
			  "mean"
			],
			"displayMode": "table",
			"placement": "right",
			"showLegend": true
		  },
		  "tooltip": {
			"mode": "single",
			"sort": "none"
		  }
		},
		"targets": [
		  {
			"datasource": {
			  "type": "prometheus",
			  "uid": "PDCCC65849FA2FB5B"
			},
			"editorMode": "code",
			"expr": "sum by (code) (rate(wbx_catalog_proxy_responses_total{cluster=~\"${cluster}\", kubernetes_namespace=~\"${namespace}\",path=~\".*/preview.*\", kubernetes_pod_name=~\"${podName}\"}[1m]))",
			"instant": false,
			"legendFormat": "{{code}}",
			"range": true,
			"refId": "A"
		  }
		],
		"title": "Status code",
		"transparent": true,
		"type": "timeseries"
	  }
	],
	"refresh": "1m",
	"schemaVersion": 39,
	"tags": [],
	"templating": {
	  "list": [
		{
		  "current": {
			"selected": false,
			"text": [
			  "k8s.elrec",
			  "k8s.nbrec",
			  "k8s.dmrec"
			],
			"value": [
			  "k8s.elrec",
			  "k8s.nbrec",
			  "k8s.dmrec"
			]
		  },
		  "hide": 0,
		  "includeAll": true,
		  "label": "Cluster",
		  "multi": true,
		  "name": "cluster",
		  "options": [
			{
			  "selected": false,
			  "text": "All",
			  "value": "$__all"
			},
			{
			  "selected": false,
			  "text": "k8s.nbstage",
			  "value": "k8s.nbstage"
			},
			{
			  "selected": true,
			  "text": "k8s.elrec",
			  "value": "k8s.elrec"
			},
			{
			  "selected": false,
			  "text": "k8s.xcrec",
			  "value": "k8s.xcrec"
			},
			{
			  "selected": true,
			  "text": "k8s.nbrec",
			  "value": "k8s.nbrec"
			},
			{
			  "selected": true,
			  "text": "k8s.dmrec",
			  "value": "k8s.dmrec"
			},
			{
			  "selected": false,
			  "text": "k8s.elcat",
			  "value": "k8s.elcat"
			},
			{
			  "selected": false,
			  "text": "k8s.xccat",
			  "value": "k8s.xccat"
			},
			{
			  "selected": false,
			  "text": "k8s.dmcat",
			  "value": "k8s.dmcat"
			},
			{
			  "selected": false,
			  "text": "k8s.xscat",
			  "value": "k8s.xscat"
			}
		  ],
		  "query": "k8s.nbstage, k8s.elrec, k8s.xcrec, k8s.nbrec, k8s.dmrec, k8s.elcat, k8s.xccat, k8s.dmcat, k8s.xscat",
		  "queryValue": "",
		  "skipUrlSync": false,
		  "type": "custom"
		},
		{
		  "current": {
			"selected": false,
			"text": [
			  "rec-pers-online"
			],
			"value": [
			  "rec-pers-online"
			]
		  },
		  "hide": 0,
		  "includeAll": true,
		  "label": "Namespace",
		  "multi": true,
		  "name": "namespace",
		  "options": [
			{
			  "selected": false,
			  "text": "All",
			  "value": "$__all"
			},
			{
			  "selected": true,
			  "text": "rec-pers-online",
			  "value": "rec-pers-online"
			},
			{
			  "selected": false,
			  "text": "rec-online",
			  "value": "rec-online"
			},
			{
			  "selected": false,
			  "text": "rec-visual-online",
			  "value": "rec-visual-online"
			},
			{
			  "selected": false,
			  "text": "wbx-search",
			  "value": "wbx-search"
			}
		  ],
		  "query": "rec-pers-online, rec-online, rec-visual-online,wbx-search",
		  "queryValue": "",
		  "skipUrlSync": false,
		  "type": "custom"
		},
		{
		  "current": {
			"selected": false,
			"text": "recom-merger.*",
			"value": "recom-merger.*"
		  },
		  "hide": 0,
		  "includeAll": false,
		  "label": "podName",
		  "multi": false,
		  "name": "podName",
		  "options": [
			{
			  "selected": true,
			  "text": "recom-merger.*",
			  "value": "recom-merger.*"
			},
			{
			  "selected": false,
			  "text": "wbx-recom-similar-merger.*",
			  "value": "wbx-recom-similar-merger.*"
			},
			{
			  "selected": false,
			  "text": "recom-interests-merger.*",
			  "value": "recom-interests-merger.*"
			},
			{
			  "selected": false,
			  "text": "recom-visual-merger.*",
			  "value": "recom-visual-merger.*"
			}
		  ],
		  "query": "recom-merger.*, wbx-recom-similar-merger.*, recom-interests-merger.*, recom-visual-merger.*",
		  "queryValue": "",
		  "skipUrlSync": false,
		  "type": "custom"
		}
	  ]
	},
	"time": {
	  "from": "now-1h",
	  "to": "now"
	},
	"timepicker": {
	  "refresh_intervals": [
		"1m",
		"5m",
		"15m",
		"30m",
		"1h",
		"2h",
		"1d"
	  ]
	},
	"timezone": "",
	"title": "Recom merger",
	"uid": "_3CFhetIz",
	"version": 35,
	"weekStart": ""
  }