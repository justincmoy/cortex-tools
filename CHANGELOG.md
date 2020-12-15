# Changelog

Order should be `CHANGE`, `FEATURE`, `ENHANCEMENT`, and `BUGFIX`

## Unreleased

## v0.6.0

* [CHANGE] When using `rules` commands, cortex ruler API requests will now default to using the `/api/v1/` prefix. The `--use-legacy-routes` flag has been added to allow users to use the original `/api/prom/` routes. #99
* [FEATURE] Add support for position rule-files arguments to `rules sync` and `rules diff` #125
* [FEATURE] Add an allow-list of namespaces for `rules sync` and `rules diff` #125
* [ENHANCEMENT] Handle trailing slashes in URLs on `cortextool`. #128
* [BUGFIX] Fix inaccuracy in `e2ealerting` caused by invalid purging condition on timestamps. #117

## v0.5.0

* [FEATURE] Release `cortextool` via Homebrew for macOS #109
* [FEATURE] Add new binary `e2ealerting` for measuring end to end alerting latency. #110
* [BUGFIX] Do not panic if we're unable to contact GitHub for the `version` command. #107

## v0.4.1

* [ENHANCEMENT] Upgrade the Go version used in build images and tests to golang 1.14.9 to match upstream Cortex. #104
* [FEATURE] Add `chunktool chunk validate-index` and `chunktool chunk clean-index` commands to the chunktool. These commands are used to scan Cortex index backends for invalid index entries. #104

## v0.4.0

* [ENHANCEMENT] Loadgen: Allow users to selectively disable query or write loadgen by leaving their respective URL configs empty. #95
* [FEATURE] Add overrides-exporter to cortextool, which exports Cortex runtime configuration overrides as metrics. #91

## v0.3.2

* [BUGFIX] Supports `rules lint` with LogQL: [#92](https://github.com/grafana/cortex-tools/pull/92).

## v0.3.1

* [FEATURE] Add support for GME remote-write rule groups. #82
* [BUGFIX] Fix issue where rule comparisons would be affected by the format of the YAML file. #88

## v0.3.0

* [FEATURE] Added loki backend support for the rules commands, configurable with `--backend=loki` (defaults to cortex).
* [FEATURE] Introduces a new `version` command. The command will also let you know if there's a new version available.

## v0.2.4

* [BUGFIX] Fix alertmanager registration subcommand and path for the configuration API #72

## v0.2.3

* [FEATURE] Added `alerts verify` command, which can be used to compare `ALERTS` series in Cortex to verify if Prometheus and Cortex Ruler are firing the same alerts
* [BUGFIX] Renamed module from cortextool to cortex-tools to ensure `go get` works properly.
* [BUGFIX] When using `--disable-color` for `rules get`, it now actually prints rules instead of the bytes of the underlying string
* [ENHANCEMENT] Allow mutualTLS for Cortex API client for `rules` and `alertmanager` cmds with:
  - `--tls-ca-path` or `CORTEX_TLS_CA_PATH`
  - `--tls-cert-path` or `CORTEX_TLS_CERT_PATH`
  - `--tls-key-path` or `CORTEX_TLS_KEY_PATH`

## v0.2.2 / 2020-06-09

* [BUGFIX] Remove usage of alternate PromQL parser in `rules prepare lint`.
* [BUGFIX] `rules check` does not require an argument.

## v0.2.1 / 2020-06-08

* [FEATURE] Add `rules check` command. It runs various [best practice](https://prometheus.io/docs/practices/rules/) checks against rules.
* [ENHANCEMENT] Ensure `rules prepare` takes into account Vector Matching on PromQL Binary Expressions.
* [BUGFIX] `rules prepare` and `rules lint` do not require an argument.

## v0.2.0 / 2020-06-02

* [FEATURE] Add `rules prepare` command. It allows you add a label to PromQL aggregations and lint your expressions in rule files.
* [FEATURE] Add `rules lint` command. It lints, rules YAML and PromQL expression formatting within the rule file
* [FEATURE] Add `logtool` binary. It parses Loki/Cortex query-frontend logs and formats them for easy analysis.

## v0.1.4 / 2020-03-10

* [CHANGE] Ensure 404 deletes do not trigger an error for `rules` and `alertmanager` commands #28
* [ENHANCEMENT] Add separate `chunktool` for operating on cortex chunk backends #23 #26
* [ENHANCEMENT] Add `--disable-color` flag to `rules print`, `rules diff`, and `alertmanager get` command #25
* [BUGFIX] `alertmanager load` command will ensure provided template files are also loaded #25
* [BUGFIX] Ensure rules commands use escaped URLs when contacting cortex API #24

## v0.1.3 / 2020-02-04

* [FEATURE] Add `rules sync` command
* [FEATURE] Add `rules diff` command

## v0.1.2 / 2019-12-18

* [CHANGE] Info log when a resource does not exist instead of exiting fatally
* [FEATURE] Windows build during Makefile cross compilation
* [BUGFIX] Fix env var `CORTEX_TENTANT_ID` to `CORTEX_TENANT_ID` for rule commands
