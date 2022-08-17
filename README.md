# repo-metrics

## About this plugin

This plugin provides metrics regarding Artifactory ocal / Remote / Virtual / Federated / Build repositories such as :

* Number of repositories / Package Types
* Top 5 biggest repositories by size
* Top 5 biggest repositories by number of artifacts
* Only for Virtual repositories : Top 5 biggest repositories based on nb of aggregated repositories and size

## Installation with JFrog CLI

Installing the latest version:

`$ jf plugin install repo-metrics`

Installing a specific version:

`$ jf plugin install repo-metrics@version`

Uninstalling a plugin

`$ jf plugin uninstall repo-metrics`

## Usage

### Commands

* report
  * Arguments:
        - test - blabllala
    * Flags:
      * file: [Optional] path to the JSON output from this [Artifactory API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-GetStorageSummaryInfo)
    * Example:

```bash
jf repo-metrics report [--repoType=all --file="/path/to/json]"
```

### Environment variables

* HELLO_FROG_GREET_PREFIX - Adds a prefix to every greet **[Default: New greeting: ]**

## Additional info

None.

## Release Notes

The release notes are available [here](RELEASE.md).
