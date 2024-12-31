// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)
package upgrade

import (
    "fmt"
    "os"
    "os/exec"
    "regexp"
    "sort"
    "github.com/hashicorp/go-version"
)

// LyfeblocVersions is a custom comparator for sorting semver version strings.
type LyfeblocVersions []string

// Len is the number of stored versions.
func (v LyfeblocVersions) Len() int { return len(v) }

// Swap swaps the elements with indexes i and j. It is needed to sort the slice.
func (v LyfeblocVersions) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

// Less compares semver versions strings properly
func (v LyfeblocVersions) Less(i, j int) bool {
    v1, err := version.NewVersion(v[i])
    if err != nil {
        panic(fmt.Sprintf("couldn't interpret version as SemVer string: %s: %s", v[i], err.Error()))
    }
    v2, err := version.NewVersion(v[j])
    if err != nil {
        panic(fmt.Sprintf("couldn't interpret version as SemVer string: %s: %s", v[j], err.Error()))
    }
    return v1.LessThan(v2)
}

// CheckUpgradeProposalVersion checks the upgrade proposal version for the given version string.
// Since this is a new blockchain, it always returns DefaultProposal.
func CheckUpgradeProposalVersion(version string) ProposalVersion {
    return DefaultProposal
}

// RetrieveUpgradesList parses the app/upgrades folder and returns a slice of semver upgrade versions
// in ascending order, e.g ["v1.0.0", "v1.0.1", "v1.1.0", ... , "v10.0.0"]
func RetrieveUpgradesList(upgradesPath string) ([]string, error) {
    dirs, err := os.ReadDir(upgradesPath)
    if err != nil {
        return nil, err
    }
    // Preallocate slice to store versions
    versions := make([]string, 0, len(dirs))
    // Pattern to find quoted string (upgrade version) in a file e.g. "v10.0.0"
    pattern := regexp.MustCompile(`"(.*?)"`)
    for _, d := range dirs {
        if !d.IsDir() {
            continue
        }
        // Creating path to upgrade dir file with constant upgrade version
        constantsPath := fmt.Sprintf("%s/%s/constants.go", upgradesPath, d.Name())
        if _, err = os.Stat(constantsPath); os.IsNotExist(err) {
            continue
        }
        f, err := os.ReadFile(constantsPath)
        if err != nil {
            return nil, err
        }
        v := pattern.FindString(string(f))
        // v[1 : len(v)-1] subslice used to remove quotes from version string
        versions = append(versions, v[1:len(v)-1])
    }
    sort.Sort(LyfeblocVersions(versions))
    return versions, nil
}

// ExportState executes the 'docker cp' command to copy container .lyfeblocd dir
// to the specified target dir (local)
//
// See https://docs.docker.com/engine/reference/commandline/cp/
func (m *Manager) ExportState(targetDir string) error {
    /* #nosec G204 */
    cmd := exec.Command(
        "docker",
        "cp",
        fmt.Sprintf("%s:/root/.lyfeblocd", m.ContainerID()),
        targetDir,
    )
    return cmd.Run()
}