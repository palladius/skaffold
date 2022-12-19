/*
Copyright 2020 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v2beta16

import (
	"context"

	"github.com/GoogleContainerTools/skaffold/v2/pkg/skaffold/output/log"
	"github.com/GoogleContainerTools/skaffold/v2/pkg/skaffold/schema/util"
	next "github.com/GoogleContainerTools/skaffold/v2/pkg/skaffold/schema/v2beta17"
	pkgutil "github.com/GoogleContainerTools/skaffold/v2/pkg/skaffold/util"
)

// Upgrade upgrades a configuration to the next version.
// Config changes from v2beta16 to v2beta17
func (c *SkaffoldConfig) Upgrade() (util.VersionedConfig, error) {
	var newConfig next.SkaffoldConfig
	pkgutil.CloneThroughJSON(c, &newConfig)
	newConfig.APIVersion = next.Version

	err := util.UpgradePipelines(c, &newConfig, upgradeOnePipeline)
	return &newConfig, err
}

func upgradeOnePipeline(oldPipeline, newPipeline interface{}) error {
	for _, a := range oldPipeline.(*Pipeline).Build.Artifacts {
		if a.DockerArtifact != nil && a.DockerArtifact.Secret != nil && a.DockerArtifact.Secret.Destination != "" {
			log.Entry(context.TODO()).Warnf("Artifact %q: Docker secret destination is no longer supported: %q", a.ImageName, a.DockerArtifact.Secret.Destination)
		}
	}
	return nil
}
