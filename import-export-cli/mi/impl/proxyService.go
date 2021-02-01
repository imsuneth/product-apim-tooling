/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package impl

import (
	"github.com/wso2/product-apim-tooling/import-export-cli/utils"
)

// ActivateProxy activates a proxy service deployed in the micro integrator in a given environment
func ActivateProxy(env, proxyName string) (interface{}, error) {
	return updateProxySerivceState(env, proxyName, "active")
}

// DeactivateProxy deactivates a proxy service deployed in the micro integrator in a given environment
func DeactivateProxy(env, proxyName string) (interface{}, error) {
	return updateProxySerivceState(env, proxyName, "inactive")
}

func updateProxySerivceState(env, proxyName, state string) (interface{}, error) {
	url := utils.GetMIManagementEndpointOfResource(utils.MiManagementProxyServiceResource, env, utils.MainConfigFilePath)
	return updateArtifactState(url, proxyName, state, env)
}
