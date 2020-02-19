// Copyright 2020 Netflix Inc
// Author: Colin McIntosh (colin@netflix.com)

package exporters

import (
	"github.com/openconfig/gnmi/ctree"
)

type Exporter interface {
	Start() error
	Export(leaf *ctree.Leaf)
}
