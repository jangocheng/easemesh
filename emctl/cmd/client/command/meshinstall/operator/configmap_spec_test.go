/*
 * Copyright (c) 2017, MegaEase
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package operator

import (
	"testing"

	"github.com/megaease/easemeshctl/cmd/client/command/meshinstall/base/fake"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestDeployOperatorConfigMap(t *testing.T) {

	client := testclient.NewSimpleClientset()
	stageContext := fake.NewStageContextForApply(client, nil)

	err := configMapSpec(stageContext).Deploy(stageContext)
	if err != nil {
		t.Fatalf("deployment operator configmap err %s", err)
	}

}
