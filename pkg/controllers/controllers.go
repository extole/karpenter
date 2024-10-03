/*
Copyright The Kubernetes Authors.

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

package controllers

import (
	"k8s.io/utils/clock"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/extole/karpenter/pkg/operator/controller"

	"github.com/extole/karpenter/pkg/cloudprovider"
	"github.com/extole/karpenter/pkg/controllers/disruption"
	"github.com/extole/karpenter/pkg/controllers/disruption/orchestration"
	"github.com/extole/karpenter/pkg/controllers/leasegarbagecollection"
	metricsnode "github.com/extole/karpenter/pkg/controllers/metrics/node"
	metricsnodepool "github.com/extole/karpenter/pkg/controllers/metrics/nodepool"
	metricspod "github.com/extole/karpenter/pkg/controllers/metrics/pod"
	"github.com/extole/karpenter/pkg/controllers/node/termination"
	"github.com/extole/karpenter/pkg/controllers/node/termination/terminator"
	nodeclaimconsistency "github.com/extole/karpenter/pkg/controllers/nodeclaim/consistency"
	nodeclaimdisruption "github.com/extole/karpenter/pkg/controllers/nodeclaim/disruption"
	nodeclaimgarbagecollection "github.com/extole/karpenter/pkg/controllers/nodeclaim/garbagecollection"
	nodeclaimlifecycle "github.com/extole/karpenter/pkg/controllers/nodeclaim/lifecycle"
	nodeclaimtermination "github.com/extole/karpenter/pkg/controllers/nodeclaim/termination"
	nodepoolcounter "github.com/extole/karpenter/pkg/controllers/nodepool/counter"
	nodepoolhash "github.com/extole/karpenter/pkg/controllers/nodepool/hash"
	"github.com/extole/karpenter/pkg/controllers/provisioning"
	"github.com/extole/karpenter/pkg/controllers/state"
	"github.com/extole/karpenter/pkg/controllers/state/informer"
	"github.com/extole/karpenter/pkg/events"
)

func NewControllers(
	clock clock.Clock,
	kubeClient client.Client,
	cluster *state.Cluster,
	recorder events.Recorder,
	cloudProvider cloudprovider.CloudProvider,
) []controller.Controller {

	p := provisioning.NewProvisioner(kubeClient, recorder, cloudProvider, cluster)
	evictionQueue := terminator.NewQueue(kubeClient, recorder)
	disruptionQueue := orchestration.NewQueue(kubeClient, recorder, cluster, clock, p)

	return []controller.Controller{
		p, evictionQueue, disruptionQueue,
		disruption.NewController(clock, kubeClient, p, cloudProvider, recorder, cluster, disruptionQueue),
		provisioning.NewPodController(kubeClient, p, recorder),
		provisioning.NewNodeController(kubeClient, p, recorder),
		nodepoolhash.NewController(kubeClient),
		informer.NewDaemonSetController(kubeClient, cluster),
		informer.NewNodeController(kubeClient, cluster),
		informer.NewPodController(kubeClient, cluster),
		informer.NewNodePoolController(kubeClient, cluster),
		informer.NewNodeClaimController(kubeClient, cluster),
		termination.NewController(kubeClient, cloudProvider, terminator.NewTerminator(clock, kubeClient, evictionQueue), recorder),
		metricspod.NewController(kubeClient),
		metricsnodepool.NewController(kubeClient),
		metricsnode.NewController(cluster),
		nodepoolcounter.NewController(kubeClient, cluster),
		nodeclaimconsistency.NewController(clock, kubeClient, recorder),
		nodeclaimlifecycle.NewController(clock, kubeClient, cloudProvider, recorder),
		nodeclaimgarbagecollection.NewController(clock, kubeClient, cloudProvider),
		nodeclaimtermination.NewController(kubeClient, cloudProvider),
		nodeclaimdisruption.NewController(clock, kubeClient, cluster, cloudProvider),
		leasegarbagecollection.NewController(kubeClient),
	}
}
