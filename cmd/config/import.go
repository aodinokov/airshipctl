/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"opendev.org/airship/airshipctl/pkg/config"
)

const (
	useImportLong = `
Merge the clusters, contexts, and users from an existing kubeConfig file into the airshipctl config file.
`

	useImportExample = `
# Import from a kubeConfig file"
airshipctl config import $HOME/.kube/config
`
)

// NewImportCommand creates a command that merges clusters, contexts, and
// users from a kubeConfig file into the airshipctl config file.
func NewImportCommand(cfgFactory config.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "import <kubeConfig>",
		Short:   "Merge information from a kubernetes config file",
		Long:    useImportLong[1:],
		Example: useImportExample,
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := cfgFactory()
			if err != nil {
				return err
			}
			kubeConfigPath := args[0]
			err = cfg.ImportFromKubeConfig(kubeConfigPath)
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Updated airship config with content imported from %q.\n", kubeConfigPath)
			return nil
		},
	}

	return cmd
}
