// Copyright (c) 2020 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

package context

import (
	"os"

	"github.com/ProtonMail/proton-bridge/internal/bridge"
	"github.com/ProtonMail/proton-bridge/pkg/pmapi"
	"github.com/ProtonMail/proton-bridge/test/fakeapi"
	"github.com/ProtonMail/proton-bridge/test/liveapi"
)

type PMAPIController interface {
	GetClient(userID string) bridge.PMAPIProvider
	TurnInternetConnectionOff()
	TurnInternetConnectionOn()
	AddUser(user *pmapi.User, addresses *pmapi.AddressList, password string, twoFAEnabled bool) error
	AddUserLabel(username string, label *pmapi.Label) error
	GetLabelIDs(username string, labelNames []string) ([]string, error)
	AddUserMessage(username string, message *pmapi.Message) error
	GetMessageID(username, messageIndex string) string
	PrintCalls()
	WasCalled(method, path string, expectedRequest []byte) bool
	GetCalls(method, path string) [][]byte
}

func newPMAPIController() PMAPIController {
	switch os.Getenv(EnvName) {
	case EnvFake:
		return newFakePMAPIController()
	case EnvLive:
		return newLivePMAPIController()
	default:
		panic("unknown env")
	}
}

func newFakePMAPIController() PMAPIController {
	return newFakePMAPIControllerWrap(fakeapi.NewController())
}

type fakePMAPIControllerWrap struct {
	*fakeapi.Controller
}

func newFakePMAPIControllerWrap(controller *fakeapi.Controller) PMAPIController {
	return &fakePMAPIControllerWrap{Controller: controller}
}

func (s *fakePMAPIControllerWrap) GetClient(userID string) bridge.PMAPIProvider {
	return s.Controller.GetClient(userID)
}

func newLivePMAPIController() PMAPIController {
	return newLiveAPIControllerWrap(liveapi.NewController())
}

type liveAPIControllerWrap struct {
	*liveapi.Controller
}

func newLiveAPIControllerWrap(controller *liveapi.Controller) PMAPIController {
	return &liveAPIControllerWrap{Controller: controller}
}

func (s *liveAPIControllerWrap) GetClient(userID string) bridge.PMAPIProvider {
	return s.Controller.GetClient(userID)
}
