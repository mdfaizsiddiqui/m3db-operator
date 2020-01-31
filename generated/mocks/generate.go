// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// mockgen rules for generating mocks using file mode
//go:generate sh -c "mockgen -package=placement -destination=$GOPATH/src/$PACKAGE/pkg/m3admin/placement/client_mock.go -source=$GOPATH/src/$PACKAGE/pkg/m3admin/placement/types.go"
//go:generate sh -c "mockgen -package=namespace -destination=$GOPATH/src/$PACKAGE/pkg/m3admin/namespace/client_mock.go -source=$GOPATH/src/$PACKAGE/pkg/m3admin/namespace/types.go"
//go:generate sh -c "mockgen -package=topic -destination=$GOPATH/src/$PACKAGE/pkg/m3admin/topic/client_mock.go -source=$GOPATH/src/$PACKAGE/pkg/m3admin/topic/types.go"
//go:generate sh -c "mockgen -package=m3admin -destination=$GOPATH/src/$PACKAGE/pkg/m3admin/client_mock.go -source=$GOPATH/src/$PACKAGE/pkg/m3admin/client.go"
//go:generate sh -c "mockgen -package=m3db -destination=$GOPATH/src/$PACKAGE/pkg/k8sops/m3db/k8sops_mock.go -source=$GOPATH/src/$PACKAGE/pkg/k8sops/m3db/types.go"
//go:generate sh -c "mockgen -package=podidentity -destination=$GOPATH/src/$PACKAGE/pkg/k8sops/podidentity/provider_mock.go -source=$GOPATH/src/$PACKAGE/pkg/k8sops/podidentity/provider.go"

package mocks
