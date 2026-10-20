# test-repo

A small sample repository used as an analysis target while developing a
Crossplane provider for SonarQube Cloud. The provider manages the SonarQube
Cloud side (organization, GitHub binding, project, permissions, quality gates);
this repo just gives that project something real to analyze.

## Contents

| Path | What |
|------|------|
| `internal/textstats` | Line/word/char counts and top-N word frequency, with table-driven tests |
| `cmd/wordcount` | CLI that reads stdin and prints the stats |
| `sonar-project.properties` | Scanner settings for CI-based analysis (organization + project key) |
| `.github/workflows/ci.yaml` | `go vet`, tests with coverage, and a SonarQube Cloud scan when `SONAR_TOKEN` is set |

## Usage

```sh
go test ./...
echo "the cat and the hat" | go run ./cmd/wordcount -top 3
```

## Hooking up SonarQube Cloud

1. Create the organization and bind it to this GitHub organization, then
   create/import the project (via the provider, or manually).
2. Make sure `sonar.organization` and `sonar.projectKey` in
   `sonar-project.properties` match what was created.
3. Either keep **Automatic Analysis** on (no CI changes needed), or turn it off
   and add a `SONAR_TOKEN` repository secret so the CI scan step runs.
   The two modes are mutually exclusive.
