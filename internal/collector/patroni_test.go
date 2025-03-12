// Copyright 2024 - 2025 Crunchy Data Solutions, Inc.
//
// SPDX-License-Identifier: Apache-2.0

package collector

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/crunchydata/postgres-operator/internal/feature"
	"github.com/crunchydata/postgres-operator/internal/testing/require"
	"github.com/crunchydata/postgres-operator/pkg/apis/postgres-operator.crunchydata.com/v1beta1"
)

func TestEnablePatroniLogging(t *testing.T) {
	t.Run("NilInstrumentationSpec", func(t *testing.T) {
		gate := feature.NewGate()
		assert.NilError(t, gate.SetFromMap(map[string]bool{
			feature.OpenTelemetryLogs: true,
		}))
		ctx := feature.NewContext(context.Background(), gate)

		config := NewConfig(nil)
		cluster := new(v1beta1.PostgresCluster)
		require.UnmarshalInto(t, &cluster.Spec, `{
			instrumentation: {
				logs: { retentionPeriod: 5h },
			},
		}`)

		EnablePatroniLogging(ctx, cluster, config)

		result, err := config.ToYAML()
		assert.NilError(t, err)
		assert.DeepEqual(t, result, `# Generated by postgres-operator. DO NOT EDIT.
# Your changes will not be saved.
exporters:
  debug:
    verbosity: detailed
extensions:
  file_storage/patroni_logs:
    create_directory: true
    directory: /pgdata/patroni/log/receiver
    fsync: true
processors:
  batch/1s:
    timeout: 1s
  batch/200ms:
    timeout: 200ms
  batch/logs:
    send_batch_size: 8192
    timeout: 200ms
  groupbyattrs/compact: {}
  resource/patroni:
    attributes:
    - action: insert
      key: k8s.container.name
      value: database
    - action: insert
      key: k8s.namespace.name
      value: ${env:K8S_POD_NAMESPACE}
    - action: insert
      key: k8s.pod.name
      value: ${env:K8S_POD_NAME}
  resourcedetection:
    detectors: []
    override: false
    timeout: 30s
  transform/patroni_logs:
    log_statements:
    - context: log
      statements:
      - set(instrumentation_scope.name, "patroni")
      - set(cache, ParseJSON(body["original"]))
      - set(severity_text, cache["levelname"])
      - set(severity_number, SEVERITY_NUMBER_DEBUG)  where severity_text == "DEBUG"
      - set(severity_number, SEVERITY_NUMBER_INFO)   where severity_text == "INFO"
      - set(severity_number, SEVERITY_NUMBER_WARN)   where severity_text == "WARNING"
      - set(severity_number, SEVERITY_NUMBER_ERROR)  where severity_text == "ERROR"
      - set(severity_number, SEVERITY_NUMBER_FATAL)  where severity_text == "CRITICAL"
      - set(time, Time(cache["asctime"], "%F %T,%L"))
      - set(attributes["log.record.original"], body["original"])
      - set(body, cache["message"])
receivers:
  filelog/patroni_jsonlog:
    include:
    - /pgdata/patroni/log/*.log
    operators:
    - from: body
      to: body.original
      type: move
    storage: file_storage/patroni_logs
service:
  extensions:
  - file_storage/patroni_logs
  pipelines:
    logs/patroni:
      exporters:
      - debug
      processors:
      - resource/patroni
      - transform/patroni_logs
      - resourcedetection
      - batch/logs
      - groupbyattrs/compact
      receivers:
      - filelog/patroni_jsonlog
`)
	})

	t.Run("InstrumentationSpecDefined", func(t *testing.T) {
		gate := feature.NewGate()
		assert.NilError(t, gate.SetFromMap(map[string]bool{
			feature.OpenTelemetryLogs: true,
		}))
		ctx := feature.NewContext(context.Background(), gate)

		cluster := new(v1beta1.PostgresCluster)
		cluster.Spec.Instrumentation = testInstrumentationSpec()
		config := NewConfig(cluster.Spec.Instrumentation)

		EnablePatroniLogging(ctx, cluster, config)

		result, err := config.ToYAML()
		assert.NilError(t, err)
		assert.DeepEqual(t, result, `# Generated by postgres-operator. DO NOT EDIT.
# Your changes will not be saved.
exporters:
  debug:
    verbosity: detailed
  googlecloud:
    log:
      default_log_name: opentelemetry.io/collector-exported-log
    project: google-project-name
extensions:
  file_storage/patroni_logs:
    create_directory: true
    directory: /pgdata/patroni/log/receiver
    fsync: true
processors:
  batch/1s:
    timeout: 1s
  batch/200ms:
    timeout: 200ms
  batch/logs:
    send_batch_size: 8192
    timeout: 200ms
  groupbyattrs/compact: {}
  resource/patroni:
    attributes:
    - action: insert
      key: k8s.container.name
      value: database
    - action: insert
      key: k8s.namespace.name
      value: ${env:K8S_POD_NAMESPACE}
    - action: insert
      key: k8s.pod.name
      value: ${env:K8S_POD_NAME}
  resourcedetection:
    detectors: []
    override: false
    timeout: 30s
  transform/patroni_logs:
    log_statements:
    - context: log
      statements:
      - set(instrumentation_scope.name, "patroni")
      - set(cache, ParseJSON(body["original"]))
      - set(severity_text, cache["levelname"])
      - set(severity_number, SEVERITY_NUMBER_DEBUG)  where severity_text == "DEBUG"
      - set(severity_number, SEVERITY_NUMBER_INFO)   where severity_text == "INFO"
      - set(severity_number, SEVERITY_NUMBER_WARN)   where severity_text == "WARNING"
      - set(severity_number, SEVERITY_NUMBER_ERROR)  where severity_text == "ERROR"
      - set(severity_number, SEVERITY_NUMBER_FATAL)  where severity_text == "CRITICAL"
      - set(time, Time(cache["asctime"], "%F %T,%L"))
      - set(attributes["log.record.original"], body["original"])
      - set(body, cache["message"])
receivers:
  filelog/patroni_jsonlog:
    include:
    - /pgdata/patroni/log/*.log
    operators:
    - from: body
      to: body.original
      type: move
    storage: file_storage/patroni_logs
service:
  extensions:
  - file_storage/patroni_logs
  pipelines:
    logs/patroni:
      exporters:
      - googlecloud
      processors:
      - resource/patroni
      - transform/patroni_logs
      - resourcedetection
      - batch/logs
      - groupbyattrs/compact
      receivers:
      - filelog/patroni_jsonlog
`)
	})
}
