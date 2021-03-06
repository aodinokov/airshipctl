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

package errors

import "fmt"

// ErrNotImplemented returned for features not yet implemented
type ErrNotImplemented struct {
	What string
}

func (e ErrNotImplemented) Error() string {
	if e.What != "" {
		return fmt.Sprintf("not implemented: %s", e.What)
	}
	return "not implemented"
}

// ErrInvalidPhase is returned if the phase is invalid
type ErrInvalidPhase struct {
	Reason string
}

func (e ErrInvalidPhase) Error() string {
	return fmt.Sprintf("invalid phase: %s", e.Reason)
}
