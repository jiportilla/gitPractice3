///hznconfig.go





import (
	"flag"
	"os"
	"strings"

	"github.com/open-horizon/anax/cli/agreement"
	"github.com/open-horizon/anax/cli/agreementbot"
	"github.com/open-horizon/anax/cli/attribute"
	"github.com/open-horizon/anax/cli/cliconfig"
	"github.com/open-horizon/anax/cli/cliutils"
	"github.com/open-horizon/anax/cli/dev"
	"github.com/open-horizon/anax/cli/eventlog"
	"github.com/open-horizon/anax/cli/exchange"
	_ "github.com/open-horizon/anax/cli/helm_deployment"
	"github.com/open-horizon/anax/cli/key"
	"github.com/open-horizon/anax/cli/metering"
	_ "github.com/open-horizon/anax/cli/native_deployment"
	"github.com/open-horizon/anax/cli/node"
	"github.com/open-horizon/anax/cli/policy"
	"github.com/open-horizon/anax/cli/register"
	"github.com/open-horizon/anax/cli/service"
	"github.com/open-horizon/anax/cli/status"
	"github.com/open-horizon/anax/cli/sync_service"
	"github.com/open-horizon/anax/cli/unregister"
	"github.com/open-horizon/anax/cli/utilcmds"
	"github.com/open-horizon/anax/cutil"
	"gopkg.in/alecthomas/kingpin.v2"
)
