package genesisblock

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
	"time"

	"github.com/zoobc/lib/address"

	"github.com/spf13/cobra"

	"github.com/zoobc/zoobc-core/common/chaintype"
	"github.com/zoobc/zoobc-core/common/constant"
	"github.com/zoobc/zoobc-core/common/crypto"
	"github.com/zoobc/zoobc-core/common/database"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
	"github.com/zoobc/zoobc-core/common/transaction"
	"github.com/zoobc/zoobc-core/common/util"
	"github.com/zoobc/zoobc-core/core/service"
	coreUtil "github.com/zoobc/zoobc-core/core/util"
)

type (
	genesisEntry struct {
		AccountAddress     string
		AccountSeed        string
		AccountBalance     int64
		NodeSeed           string
		NodeAccountAddress string
		NodePublicKey      []byte
		NodeAddress        string
		LockedBalance      int64
		ParticipationScore int64
		Smithing           bool
	}
	clusterConfigEntry struct {
		NodePublicKey       string `json:"nodePublicKey"`
		NodeSeed            string `json:"nodeSeed"`
		OwnerAccountAddress string `json:"ownerAccountAddress"`
		NodeAddress         string `json:"myAddress,omitempty"`
		Smithing            bool   `json:"smithing,omitempty"`
	}
	accountNodeEntry struct {
		NodeAddress    string
		AccountAddress string
	}
)

var (
	genesisGeneratorCmd = &cobra.Command{
		Use:   "generate",
		Short: "genesis generate command used to generate a new genesis.go and cluster_config.json file",
		Long:  `genesis generate command generate a genesis.go file from a list of accounts and/or from current database.`,
		Run: func(cmd *cobra.Command, args []string) {
			if _, ok := envTargetValue[envTarget]; !ok {
				log.Fatal("Invalid env-target flag given, only: develop,staging,alpha")
			}
			generateGenesisFiles(withDbLastState, dbPath, extraNodesCount)
		},
	}
)

func Commands() *cobra.Command {
	return genesisCmd
}

func init() {
	// TODO: this is broken since implementation of proof of participation mechanism (needs more than just last blockchain status).
	//  see snapshots implementation as reference
	// genesisGeneratorCmd.Flags().BoolVarP(&withDbLastState, "withDbLastState", "w", false,
	// 	"add to genesis all registered nodes and account balances from current database")
	genesisGeneratorCmd.Flags().StringVarP(&dbPath, "dbPath", "f", "../resource/",
		"path of blockchain's database to be used as data source in case the -w flag is used. If not set, the default resource folder is used")
	genesisGeneratorCmd.Flags().IntVarP(&extraNodesCount, "extraNodes", "n", 0,
		"number of 'extra' autogenerated nodes to be deployed using cluster_config.json")
	genesisGeneratorCmd.Flags().StringVar(&logLevels, "logLevels", "fatal error panic",
		"default log levels for all nodes (for kvConsulScript.sh). example: 'warn info fatal error panic'")
	genesisGeneratorCmd.Flags().StringVar(&wellKnownPeers, "wellKnownPeers", "127.0.0.1:8001",
		"default wellKnownPeers for all nodes (for kvConsulScript.sh). example: 'n0.alpha.proofofparticipation.network n1.alpha."+
			"proofofparticipation.network n2.alpha.proofofparticipation.network'")
	genesisGeneratorCmd.Flags().StringVar(&deploymentName, "deploymentName", "zoobc-alpha",
		"nomad task name associated to this deployment")
	genesisGeneratorCmd.Flags().StringVar(&kvFileCustomConfigFile, "kvFileCustomConfigFile", "",
		"(optional) full path (path + fileName) of a custom cluster_config.json file to use to generate consulKvInitScript."+
			"sh instead of the automatically generated in resource/generated/genesis directory")

	genesisGeneratorCmd.Flags().StringVarP(&envTarget, "env-target", "e", "alpha", "env mode indeed a.k.a develop,staging,alpha")
	genesisGeneratorCmd.Flags().StringVarP(&output, "output", "o", "resource", "output generated files target")
	genesisCmd.AddCommand(genesisGeneratorCmd)
}

// generateGenesisFiles generate genesis files starting from a source json file.
// PreRegisteredNodes contains a list of known nodes-accountOwners' public keys to be included in genesis.
// Data that can be pre-set for node registration and and account balance are:
//   AccountAddress (mandatory): node's owner
//   AccountBalance (for funded accounts only): the balance of that account at genesis block
//   NodeSeed (this should be set only for testing nodes): it will be copied into cluster_config.json to
//       automatically deploy new nodes that are already registered
//   NodeAccountAddress (mandatory): base64 encoded node public key to be registered
//   NodeAddress (optional): if known, the node address that will be registered and put in cluster_config.json too
//   LockedBalance (optional): account's locked balance
//   ParticipationScore (optional): set custom initial participation score (mainly for testing the smith process and POP algorithm).
//       if not set, defaults to constant.DefaultParticipationScore
//
// withDbLastState if set to true, we also scan a given blockchain database and extract latest state to be included in genesis
//  (account balances and registered nodes/participation scores)
func generateGenesisFiles(withDbLastState bool, dbPath string, extraNodesCount int) {
	var (
		bcState, preRegisteredNodes, preRegisteredAccountAddresses []genesisEntry
		accountNodes                                               []accountNodeEntry
		err                                                        error
	)

	if withDbLastState {
		bcState, err = getDbLastState(dbPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	file, err := ioutil.ReadFile(path.Join(getRootPath(), fmt.Sprintf("./cmd/genesisblock/templates/%s.preRegisteredNodes.json", envTarget)))
	if err == nil {
		err = json.Unmarshal(file, &preRegisteredNodes)
		if err != nil {
			log.Fatalf("preRegisteredNodes.json parsing error: %s", err)
		}

		// merge duplicates: if preRegisteredNodes contains entries that are in db too, add the parameters that are't available in db,
		// which is are NodeSeed and Smithing
		for _, prNode := range preRegisteredNodes {
			found := false
			var pubKey = make([]byte, 32)
			for i, e := range bcState {
				if prNode.AccountAddress != e.AccountAddress {
					continue
				}
				bcState[i].NodeSeed = prNode.NodeSeed
				bcState[i].Smithing = prNode.Smithing
				err := address.DecodeZbcID(prNode.NodeAccountAddress, pubKey)
				if err != nil {
					log.Fatal(err)
				}
				bcState[i].NodePublicKey = pubKey
				found = true
				break
			}
			if !found {
				err := address.DecodeZbcID(prNode.NodeAccountAddress, pubKey)
				if err != nil {
					log.Fatal(err)
				}
				prNode.NodePublicKey = pubKey
				bcState = append(bcState, prNode)
			}
		}
	}

	var idx int
	for idx = 0; idx < extraNodesCount; idx++ {
		bcState = append(bcState, generateRandomGenesisEntry(idx, ""))
	}

	// generate extra nodes from a json file containing only account addresses
	file, err = ioutil.ReadFile(path.Join(getRootPath(), fmt.Sprintf("./cmd/genesisblock/templates/%s.genesisAccountAddresses.json", envTarget)))
	if err == nil {
		// read custom addresses from file
		err = json.Unmarshal(file, &preRegisteredAccountAddresses)
		if err != nil {
			log.Fatalf("preRegisteredAccountAddresses.json parsing error: %s", err)
		}
		if idx == 0 {
			idx--
		}
		for _, preRegisteredAccountAddress := range preRegisteredAccountAddresses {
			idx++
			bcState = append(bcState, generateRandomGenesisEntry(idx, preRegisteredAccountAddress.AccountAddress))
		}
	}

	// append to preRegistered nodes/accounts previous entries from a blockchain db file
	var outPath = path.Join(getRootPath(), fmt.Sprintf("%s/generated/genesis", output))

	if err := os.MkdirAll(outPath, os.ModePerm); err != nil {
		log.Fatalf("can't create folder %s. error: %s", outPath, err)
	}
	generateGenesisFile(bcState, fmt.Sprintf("%s/genesis.go", outPath))
	clusterConfig := generateClusterConfigFile(bcState, fmt.Sprintf("%s/cluster_config.json", outPath))
	// generate a bash script to init consul key/value data store in case we automatically deploy all nodes in genesis
	generateConsulKvInitScript(clusterConfig, fmt.Sprintf("%s/consulKvInit.sh", outPath))

	// also generate a file to be shared with node owners, so they know from the wallet what node to configure as their own node

	for _, entry := range bcState {
		newEntry := accountNodeEntry{
			NodeAddress:    entry.NodeAddress,
			AccountAddress: entry.AccountAddress,
		}
		accountNodes = append(accountNodes, newEntry)
	}
	generateAccountNodesFile(accountNodes, fmt.Sprintf("%s/accountNodes.json", outPath))
	fmt.Printf("Command executed successfully\ngenesis.go.new has been generated in %s\n", outPath)
	fmt.Printf("to apply new genesis to the core-node, please overwrite common/constant/genesis."+
		"go with the new one: %s/genesis.go", outPath)
}

// generateRandomGenesisEntry randomly generates a genesis node entry
// note: the account address is mandatory for the node registration, but as there is no wallet connected to it
//       and we are not storing the relative seed, needed to sign transactions, these nodes can smith but their owners
//       can't perform any transaction.
//       This is only useful to test multiple smithing-nodes, for instence in a network stress test of tens of nodes connected together
func generateRandomGenesisEntry(nodeIdx int, accountAddress string) genesisEntry {
	var (
		ed25519Signature = crypto.NewEd25519Signature()
	)
	if accountAddress == "" {
		var (
			seed       = util.GetSecureRandomSeed()
			privateKey = ed25519Signature.GetPrivateKeyFromSeed(seed)
			publicKey  = privateKey[32:]
		)
		accountAddress, _ = ed25519Signature.GetAddressFromPublicKey(constant.PrefixZoobcNormalAccount, publicKey)
	}
	var (
		nodeSeed       = util.GetSecureRandomSeed()
		nodePrivateKey = ed25519Signature.GetPrivateKeyFromSeed(nodeSeed)
		nodePublicKey  = nodePrivateKey[32:]
	)
	nodeAccountAddress, _ := address.EncodeZbcID(constant.PrefixZoobcNodeAccount, nodePublicKey)

	return genesisEntry{
		AccountAddress:     accountAddress,
		NodePublicKey:      nodePublicKey,
		NodeAccountAddress: nodeAccountAddress,
		NodeSeed:           nodeSeed,
		ParticipationScore: constant.GenesisParticipationScore,
		Smithing:           true,
		LockedBalance:      0,
		NodeAddress:        fmt.Sprintf("n%d.alpha.proofofparticipation.network", nodeIdx),
	}
}

func getDbLastState(dbPath string) (bcEntries []genesisEntry, err error) {
	var (
		db *sql.DB
	)
	_, err = os.Stat(dbPath)
	if ok := os.IsNotExist(err); ok {
		return nil, err
	}

	dbInstance := database.NewSqliteDB()
	db, err = dbInstance.OpenDB(dbPath, "zoobc.db", 10, 10, 20*time.Minute)
	if err != nil {
		log.Fatal(err)
	}
	queryExecutor := query.NewQueryExecutor(db)
	accountBalanceQuery := query.NewAccountBalanceQuery()
	nodeRegistrationQuery := query.NewNodeRegistrationQuery()
	participationScoreQuery := query.NewParticipationScoreQuery()
	// get all account balances
	// get the participation score for this node registration
	qry := accountBalanceQuery.GetAccountBalances()
	balanceRows, err := queryExecutor.ExecuteSelect(qry, false)
	if err != nil {
		return nil, err
	}
	defer balanceRows.Close()
	accountBalances, err := accountBalanceQuery.BuildModel([]*model.AccountBalance{}, balanceRows)
	if err != nil {
		return nil, err
	}
	for _, acc := range accountBalances {
		if acc.AccountAddress == constant.MainchainGenesisAccountAddress {
			continue
		}

		var nodeRegistrations []*model.NodeRegistration

		bcEntry := new(genesisEntry)
		bcEntry.AccountAddress = acc.AccountAddress
		bcEntry.AccountBalance = acc.Balance

		err := func() error {
			// get node registration for this account, if exists
			qry, args := nodeRegistrationQuery.GetNodeRegistrationByAccountAddress(acc.AccountAddress)
			nrRows, err := queryExecutor.ExecuteSelect(qry, false, args...)
			if err != nil {
				return err
			}
			defer nrRows.Close()

			nodeRegistrations, err = nodeRegistrationQuery.BuildModel([]*model.NodeRegistration{}, nrRows)
			if err != nil {
				return err
			}
			return nil
		}()

		if err != nil {
			return nil, err
		}

		if len(nodeRegistrations) > 0 {
			nr := nodeRegistrations[0]
			bcEntry.LockedBalance = nr.LockedBalance
			if nr.NodeAddress.Port > 0 {
				bcEntry.NodeAddress = fmt.Sprintf("%s:%d", nr.NodeAddress.Address, nr.NodeAddress.Port)
			} else {
				bcEntry.NodeAddress = nr.NodeAddress.Address
			}
			bcEntry.NodePublicKey = nr.NodePublicKey
			bcEntry.NodeAccountAddress = base64.StdEncoding.EncodeToString(nr.NodePublicKey)
			err := func() error {
				// get the participation score for this node registration
				qry, args := participationScoreQuery.GetParticipationScoreByNodeID(nr.NodeID)
				psRows, err := queryExecutor.ExecuteSelect(qry, false, args...)
				if err != nil {
					return err
				}
				defer psRows.Close()

				participationScores, err := participationScoreQuery.BuildModel([]*model.ParticipationScore{}, psRows)
				if (err != nil) || len(participationScores) > 0 {
					bcEntry.ParticipationScore = participationScores[0].Score
				}
				return nil
			}()
			if err != nil {
				return nil, err
			}
		}
		bcEntries = append(bcEntries, *bcEntry)
	}

	return bcEntries, err
}

// generateGenesisFile generates a genesis file with given entries, starting from a template
func generateGenesisFile(genesisEntries []genesisEntry, newGenesisFilePath string) {
	// read and execute genesis template, outputting the genesis.go to stdout
	tmpl, err := template.ParseFiles(path.Join(getRootPath(), "./cmd/genesisblock/templates/genesis.tmpl"))
	if err != nil {
		log.Fatalf("Error while reading genesis.tmpl file: %s", err)
	}
	err = os.Remove(newGenesisFilePath)
	if err != nil {
		log.Printf("remove %s file: %s\n", newGenesisFilePath, err)
	}
	f, err := os.Create(newGenesisFilePath)
	if err != nil {
		log.Printf("create %s file: %s\n", newGenesisFilePath, err)
		return
	}
	defer f.Close()

	config := map[string]interface{}{
		"MainchainGenesisBlockID": getGenesisBlockID(genesisEntries),
		"MainchainGenesisConfig":  genesisEntries,
	}
	err = tmpl.Execute(f, config)
	if err != nil {
		log.Fatal(err)
	}
}

func getGenesisBlockID(genesisEntries []genesisEntry) int64 {
	var (
		genesisConfig []constant.GenesisConfigEntry
	)
	for _, entry := range genesisEntries {
		cfgEntry := constant.GenesisConfigEntry{
			AccountAddress:     entry.AccountAddress,
			AccountBalance:     entry.AccountBalance,
			LockedBalance:      entry.LockedBalance,
			NodeAddress:        entry.NodeAddress,
			NodePublicKey:      entry.NodePublicKey,
			ParticipationScore: entry.ParticipationScore,
		}
		genesisConfig = append(genesisConfig, cfgEntry)
	}
	bs := service.NewBlockMainService(
		&chaintype.MainChain{},
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		&transaction.TypeSwitcher{},
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		&transaction.Util{},
		&coreUtil.ReceiptUtil{},
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		query.GetPruneQuery(&chaintype.MainChain{}),
		nil,
	)
	block, err := bs.GenerateGenesisBlock(genesisConfig)
	if err != nil {
		log.Fatal(err)
	}
	return block.ID
}

func generateClusterConfigFile(genesisEntries []genesisEntry, newClusterConfigFilePath string) (clusterConfig []clusterConfigEntry) {
	for _, genesisEntry := range genesisEntries {
		// exclude entries that don't have NodeSeed set from cluster_config.json
		// (they should be nodes already registered/run by someone, thus they shouldn't be deployed automatically)
		if genesisEntry.NodeSeed != "" {
			entry := clusterConfigEntry{
				NodeAddress:         genesisEntry.NodeAddress,
				NodePublicKey:       genesisEntry.NodeAccountAddress,
				NodeSeed:            genesisEntry.NodeSeed,
				OwnerAccountAddress: genesisEntry.AccountAddress,
				Smithing:            genesisEntry.Smithing,
			}
			clusterConfig = append(clusterConfig, entry)
		}
	}
	file, err := json.MarshalIndent(clusterConfig, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling json file %s: %s\n", newClusterConfigFilePath, err)
	}
	err = ioutil.WriteFile(newClusterConfigFilePath, file, 0644)
	if err != nil {
		log.Fatalf("create %s file: %s\n", newClusterConfigFilePath, err)
	}
	return clusterConfig
}

func generateAccountNodesFile(accountNodeEntris []accountNodeEntry, configFilePath string) {
	var (
		accountNodes []accountNodeEntry
	)

	for _, entry := range accountNodeEntris {
		entry := accountNodeEntry{
			NodeAddress:    entry.NodeAddress,
			AccountAddress: entry.AccountAddress,
		}
		accountNodes = append(accountNodes, entry)
	}
	file, err := json.MarshalIndent(accountNodes, "", "  ")
	if err != nil {
		log.Fatalf("error marshaling json file %s: %s\n", configFilePath, err)
	}
	err = ioutil.WriteFile(configFilePath, file, 0644)
	if err != nil {
		log.Fatalf("create %s file: %s\n", configFilePath, err)
	}
}

// generateGenesisFile generates a genesis file with given entries, starting from a template
func generateConsulKvInitScript(clusterConfigEntries []clusterConfigEntry, consulKvInitScriptPath string) {
	// read and execute genesis template, outputting the genesis.go to stdout
	// genesisTmpl, err := helpers.ReadTemplateFile("./genesis.tmpl")
	tmpl, err := template.ParseFiles(path.Join(getRootPath(), "./cmd/genesisblock/templates/consulKvInit.tmpl"))
	if err != nil {
		log.Fatalf("Error while reading consulKvInit.tmpl file: %s", err)
	}
	err = os.Remove(consulKvInitScriptPath)
	if err != nil {
		log.Printf("remove %s file: %s\n", consulKvInitScriptPath, err)
	}
	f, err := os.Create(consulKvInitScriptPath)
	if err != nil {
		log.Fatalf("create %s file: %s\n", consulKvInitScriptPath, err)
	}
	defer f.Close()

	if kvFileCustomConfigFile != "" {
		jsonFile, err := os.Open(kvFileCustomConfigFile)
		if err != nil {
			log.Fatalf("opening file %s: error %s\n", kvFileCustomConfigFile, err)
		}
		defer jsonFile.Close()
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatalf("reading file %s: error %s\n", kvFileCustomConfigFile, err)
		}
		err = json.Unmarshal(byteValue, &clusterConfigEntries)
		if err != nil {
			log.Fatalf("parsing file %s: error %s\n", kvFileCustomConfigFile, err)
		}
	}

	config := map[string]interface{}{
		"nomadJobName":         deploymentName,
		"wellKnownPeers":       wellKnownPeers,
		"logLevels":            logLevels,
		"clusterConfigEntries": clusterConfigEntries,
	}
	err = tmpl.Execute(f, config)
	if err != nil {
		log.Fatal(err)
	}
}

func getRootPath() string {
	wd, _ := os.Getwd()
	if strings.Contains(wd, "zoobc-core/") {
		return path.Join(wd, "../")
	}
	return wd
}

func (ge *genesisEntry) FormatPubKeyByteString() string {
	if ge.NodeAccountAddress == "" {
		return ""
	}
	pubKey, err := base64.StdEncoding.DecodeString(ge.NodeAccountAddress)
	if err != nil {
		log.Fatalf("Error decoding node public key: %s", err)
	}
	return util.RenderByteArrayAsString(pubKey)
}

func (ge *genesisEntry) HasParticipationScore() bool {
	return ge.ParticipationScore > 0
}

func (ge *genesisEntry) HasLockedBalance() bool {
	return ge.LockedBalance > 0
}

func (ge *genesisEntry) HasAccountBalance() bool {
	return ge.AccountBalance > 0
}

func (ge *genesisEntry) HasNodeAddress() bool {
	return ge.NodeAddress != ""
}

func (ge *genesisEntry) HasNodePublicKey() bool {
	return ge.NodeAccountAddress != ""
}
