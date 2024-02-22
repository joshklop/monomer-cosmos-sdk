package testutil

import (
	"github.com/joshklop/monomer-cosmos-sdk/testutil/configurator"
	_ "github.com/joshklop/monomer-cosmos-sdk/x/auth"           // import as blank for app wiring
	_ "github.com/joshklop/monomer-cosmos-sdk/x/auth/tx/config" // import as blank for app wiring
	_ "github.com/joshklop/monomer-cosmos-sdk/x/auth/vesting"   // import as blank for app wiring
	_ "github.com/joshklop/monomer-cosmos-sdk/x/bank"           // import as blank for app wiring
	_ "github.com/joshklop/monomer-cosmos-sdk/x/consensus"      // import as blank for app wiring
	_ "github.com/joshklop/monomer-cosmos-sdk/x/genutil"        // import as blank for app wiring
	_ "github.com/joshklop/monomer-cosmos-sdk/x/params"         // import as blank for app wiring
	_ "github.com/joshklop/monomer-cosmos-sdk/x/staking"        // import as blank for app wiring
)

var AppConfig = configurator.NewAppConfig(
	configurator.AuthModule(),
	configurator.BankModule(),
	configurator.VestingModule(),
	configurator.StakingModule(),
	configurator.TxModule(),
	configurator.ConsensusModule(),
	configurator.ParamsModule(),
	configurator.GenutilModule(),
)
