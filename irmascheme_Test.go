package main

import (
	"fmt"
	"log"
	"irmaethereumscheme/contracts"
	"irmaethereumscheme/proto"
	"github.com/golang/protobuf/proto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
)

func main() {

	// Generate keys
	key, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(key)



	// Create simulator with prefunded account for generated keys
	alloc := core.GenesisAlloc{auth.From: {Balance: big.NewInt(10000000000)}}
	alloc[crypto.PubkeyToAddress(key.PublicKey)] = core.GenesisAccount{Balance: big.NewInt(10000000000)}
	sim := backends.NewSimulatedBackend(alloc)
	// WARNING: for all this to work, gas Limit needs to be set to a high value in protocol_params.go



	// Deploy contract including metadata
	metadata := &irmaproto.IRMASchemeMetadata{
		Version: 7,
		Id: "pbdf",
		Url: "https://privacybydesign.foundation/schememanager/pbdf",
		Name: []*irmaproto.IRMASchemeMetadata_LocalizedName{
			{ Lang: "en", Name: "Privacy by Design Foundation" },
			{ Lang: "nl", Name: "Stichting Privacy by Design" },
		},
		Description: []*irmaproto.IRMASchemeMetadata_LocalizedDescription{
			{
				Lang: "en",
				Name: "The Privacy by Design Foundation develops the IRMA app and the IRMA " +
					"infrastructure, and issues basic attributes for free.",
			},{
				Lang: "nl",
				Name: "De stichting Privacy by Design ontwikkelt de IRMA app en de IRMA" +
					"infrastructuur, en geeft gratis een set basisattributen uit.",
			},
		},
		Keyshareserver: "https://privacybydesign.foundation/tomcat/irma_keyshare_server/api/v1",
		Keysharewebsite: "https://privacybydesign.foundation/mijnirma/",
		Keyshareattribute: "pbdf.pbdf.mijnirma.email",
		Contact: "https://privacybydesign.foundation/",
	}
	metadataBuffer, err := proto.Marshal(metadata);
	fmt.Printf("Size of metadata: %d, %d\n", len(metadataBuffer), cap(metadataBuffer));

	addr, createTransaction, createdContract, err := contracts.DeployIRMAScheme(auth, sim, "pbdf", metadataBuffer)
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}
	fmt.Printf("Contract deployed to %s\n", addr.String())
	fmt.Printf("Transaction cost %d gas\n", createTransaction.Gas())
	sim.Commit()
	fmt.Printf("Created contract: %+v\n\n", createdContract.IRMASchemeCaller)


	// Opening the new contract
	contractAddr := addr.String();
	schemeContractAddress := common.HexToAddress(contractAddr);
	fmt.Printf("Opening contract at: %s\n", contractAddr)
	contract, err := contracts.NewIRMAScheme(schemeContractAddress, sim)
	if err != nil {
		log.Fatalf("could not load contract: %v", err)
	}
	fmt.Printf("Contract: %+v\n", contract)



	// Add Issuer
	issuerId := "pbdf"
	issuerLogoUrl := "logo.png"
	issuerMetadata := &irmaproto.IRMAIssuerMetadata{
		Version: 4,
		Id: "pbdf",
		Shortname: []*irmaproto.IRMAIssuerMetadata_LocalizedName{
			{ Lang: "en", Name: "PbD.f" },
			{ Lang: "nl", Name: "PbD.f" },
		},
		Name: []*irmaproto.IRMAIssuerMetadata_LocalizedName{
			{ Lang: "en", Name: "Privacy by Design Foundation" },
			{ Lang: "nl", Name: "Stichting Privacy by Design" },
		},
		Schememanager: "pbdf",
		Contactaddress: "",
		Contactemail: "info@privacybydesign.foundation",
		Baseurl: "",
	}
	issuerMetadataBuffer, err := proto.Marshal(issuerMetadata);
	fmt.Printf("Size of metadata: %d\n", len(issuerMetadataBuffer));

	transaction, err := contract.AddIssuer(auth, issuerId, issuerLogoUrl, issuerMetadataBuffer);
	if err != nil {
		log.Fatalf("could not add issuer to contract: %v", err)
	}
	fmt.Printf("Transaction cost %d gas\n\n", transaction.Gas())
	sim.Commit();





	// Add key for Issuer
	fmt.Printf("Adding a public key for issuer: %s\n", issuerId)
	issuerPk := &irmaproto.IRMAIssuerPublicKey{
		Counter: 0,
		ExpiryDate: 1509100474,
		N: big.NewInt(0).SetBytes([]byte("153083626846953041148798984867206431967476314548248599675213756172644063474530260495216333346051157056954585429878833407616458713470258614557644549055442336816765036994343083706077637150817233586215957823887604466160968116348696406527285353127675492765345080512715060583469637176051691364962519805621590733473")).Bytes(),
		Z: big.NewInt(0).SetBytes([]byte("151145812782966949004562123511348360915116849203600768286954722250271879925704384516978099406159776962580833602916965512771761193134735415560803760227980183523828836195681657377097451091691150194016329773950347273245413663949426820549440343819565194141128790683089713015200119110876669077666902685958565588052")).Bytes(),
		S: big.NewInt(0).SetBytes([]byte("12312959352926251135494460490921161934091634286217473593489364830603639867013234990280140537834447349617690724678772590211285489304226992425280976827684960652990650901939858111822201353449384669401038093259267422292728115066715560921099291332761199984328277806112630508352069021431703076854924493715632084017")).Bytes(),
		Bases: [][]byte{
			big.NewInt(0).SetBytes([]byte("33782263087307562864480900198045833374750619733798476098212578035590943035435908463276779959724126424200275797822368649153968236010770095399382254627624062049873397397148222688098561316951905533280118904360834591505517638401560113752147680276096553203012894130910046851906811597361183525806497169535320267797")).Bytes(),
			big.NewInt(0).SetBytes([]byte("106650570690598786746122399462689251329692794802696146056700015003878191154514962185126379529043804828101603851946686685047883979478767043459817003452809280670215303313729694075682202573717900691699940400992806388900776826125304705168190092591934276598534161653172184630278229357028689651051120991109783239418")).Bytes(),
			big.NewInt(0).SetBytes([]byte("15298239839860390193358254564804457906266515465341533356190062050419116312979329542279219089657703762356805992539904795731928568189593093326930999088049658794260334276361398481273113503679467911624784327724732821470644564910487906724603870356007744252411582449304505800960397177753083969705554755531491013892")).Bytes(),
			big.NewInt(0).SetBytes([]byte("119066226692287819271875205901586639650489573717572307887184191313582736075762107770891140910408012576128835436512037336555603529640574391487403617673171733123091116559826742583990565080161828170801549308103209653228377405385430038775730835604527224770360388572595118747448777665560549662969777140478002986750")).Bytes(),
			big.NewInt(0).SetBytes([]byte("129913086049966305403869561126313443451444121083223960693716417021389622144817079534833329207100281587838531613358041436912233357693211443475354476448528015632789105900387849851565355518078877678322456516657971212996179096284449116132018202522949146292861568614254890901658677663877866512283758799371230765394")).Bytes(),
			big.NewInt(0).SetBytes([]byte("145104541700986954834811916153770119511376699407541153940614568447297271862831143600503082220257879882981771742535556735723327824112871311965286618770837804132102249980186789885524646714064585936345725068676894598443507527866568410542387237466055090563274803398866766396883938181477987695941045696131164827282")).Bytes(),
			big.NewInt(0).SetBytes([]byte("121975771562025657137263368460544407537529264362427723555966904337127659036310143462972467344879087445927625640211546904243882795825921815721853618824168655249204623261906941336225556086235702807507089763177646769305371479073357562582106501871986042802579595901287313798460386958206868829738947885749875921464")).Bytes(),
			big.NewInt(0).SetBytes([]byte("134784291069778492417742687439629834108955047750467893071736863260687496258877053992997159773051476418352088057591360417734175516138917875691716294275993852050319208245562733306322722638394106478184704081911353238849080952807440667050664824921945489615404090751538209271833254635314840706620820939777866125317")).Bytes(),
			big.NewInt(0).SetBytes([]byte("59499567765381098013124037668849147083157608921772381241458621155624658671533915671214317795983068740288855650984075028960441717590316201738555075102795818639726244025406484796446393394636105164830228823564733237533767366856954721008880593330400537775586339734161930485177024190484281910272040211499205917997")).Bytes(),
			big.NewInt(0).SetBytes([]byte("45709009769068698106911189786724026693231735357757758257888447057268030046254337180385065006493500496853381439056622150821947274365444229945489137295200823356348255710147195907904262125637201840831057103466272405360437928661971384780420653449902556914988088473959670196098580608097946869544142109124596262051")).Bytes(),
			big.NewInt(0).SetBytes([]byte("82575382879741981612611049007775177332995326153980371603891354995029361207639675020634601898860694930668036971810598779935819455772155346095437375649062471786991444053949822773167687023015719052451751987269448961102241651942875358339789253846356716234660409238405055894042743094095369433988999229172460295469")).Bytes(),
			big.NewInt(0).SetBytes([]byte("150704157305178319999507727279818285132827483776851999442370479760729281625182204110395890197656083368949330459083714205511232927095262193943570698751715088094839474069403443330183153814131068326014968334716704512874193257712084970456192873304110290595898119357653854902255101163441669702566413074944456399386")).Bytes(),
			big.NewInt(0).SetBytes([]byte("10941699380645120706679647780500747083571706790657666333024692848721448169749538350895654157898679917773636629107971785863462903979551818673597201646785894723557096272876357316920028249274200158725594126811039874640806723765906083070793783762210468404617213642175589262193372063761067859163962912351167160763")).Bytes(),
			big.NewInt(0).SetBytes([]byte("114671364810417843235910756615359826154834137981143621047594194795242915278320344022944706704225995502078654012489373545743775002727718571334921636200244575898837153440434013902474500297501059943743101661599702996015361091949119373203834767238098786275719533625169668553474005956256414041472797651621709675694")).Bytes(),
			big.NewInt(0).SetBytes([]byte("82400369738799392781091605890773851749725409715183985847657479661133102811405483638231928557731498682803518872532782026273943182664937397623681921899741169214970651791139513416191760344850417642828797799523594553710861568280402733497593145901208551087787425609097055317670233821656540828127921697917876865514")).Bytes(),
			big.NewInt(0).SetBytes([]byte("62535227893948466606240086845781045436355537938190463264822158055859133233386280473861935758869030248352014797229400574590981440379141058555846754179259898038691456407620848886892292836474960137133674247894620482532312498143064642185967807152733234227411732245128937945414161243984996366526857416252024061070")).Bytes(),
			big.NewInt(0).SetBytes([]byte("119079091514930448917472632561691384279264570107053459247870729346378967437592848391519021518568438711940498396324501491420303333764083724858609876363098604601580438378583912026645682000575795305613010802782541761200581180252052663489786300291289834185809329920870795705516479970424332298317148952718770980215")).Bytes(),
			big.NewInt(0).SetBytes([]byte("87537134759956964561538108070954803212893162601646310968168723758920558716776051807072678729762134192547668818632990630804909030011492097221235005735519667046453023235293082881896791254332909900793738355670056635643023443596854148936384004920095117635071260064624946722396177299358041779950358487467437868577")).Bytes(),
			big.NewInt(0).SetBytes([]byte("26787466773161010745967555596804733307074259511568966244396266613360421420681107356830528347555290395701857395924421662407131316549991314475995132804529738765871467058492142698509620952622145089481283016066610171477597432327295707457104026821453040817024682545367417563834615698280345821733691065583620897054")).Bytes(),
			big.NewInt(0).SetBytes([]byte("142752723192160729177054222850182334657393213102478174952496741081452852114761225277226714647828162676249426705371769046270852001959158619006346344363962289199372364378247846927835689830842537857783322312409206961524939589066894613967030785165837254277180604351770947416191287141715231296015191823879323586349")).Bytes(),
		},
		EpochLength: 432000,
	}
	issuerPkBuffer, err := proto.Marshal(issuerPk);
	fmt.Printf("Size of public key: %d\n", len(issuerPkBuffer));

	keyTransaction, err := contract.AddIssuerPublicKey(auth, issuerId, issuerPkBuffer)
	if err != nil {
		log.Fatalf("could not pk to issuer : %v", err)
	}
	fmt.Printf("Transaction cost %d gas\n\n", keyTransaction.Gas())
	sim.Commit();




	// Add issuespecification/credential for Issuer
	credentialId := "ageLimits"
	fmt.Printf("Adding credential %s for issuer: %s\n", credentialId, issuerId)
	credential := &irmaproto.IRMAIssueSpecification{
		Version: 4,
		Name: []*irmaproto.IRMAIssueSpecification_LocalizedName{
			{ Lang: "en", Name: "Age Limits" },
			{ Lang: "nl", Name: "Leeftijdslimieten" },
		},
		ShortName: []*irmaproto.IRMAIssueSpecification_LocalizedName{
			{ Lang: "en", Name: "Age" },
			{ Lang: "nl", Name: "Leeftijd" },
		},
		SchemeManager: "pbdf",
		IssuerId: "pbdf",
		CredentialId: "pbdf",
		Description: []*irmaproto.IRMAIssueSpecification_LocalizedDescription{
			{ Lang: "en", Name: "Your age attributes, issued after revealing your birthdate, e.g. from the iDIN bank credentials." },
			{ Lang: "nl", Name: "Uw leeftijdsgrenzen zoals berekend uit uw geboortedatum, bijvoorbeeld uit een iDIN bank credential." },
		},
		ShouldBeSingleton: true,
		Attributes: []*irmaproto.IRMAIssueSpecification_Attribute{
			{
				Id: "over12",
				Name: []*irmaproto.IRMAIssueSpecification_LocalizedName{
					{ Lang: "en", Name: "Over 12" },
					{ Lang: "nl", Name: "Ouder dan 12" },
				},
				Description: []*irmaproto.IRMAIssueSpecification_LocalizedDescription{
					{ Lang: "en", Name: "If you are over 12" },
					{ Lang: "nl", Name: "Of u ouder dan 12 bent" },
				},
			},{
				Id: "over16",
				Name: []*irmaproto.IRMAIssueSpecification_LocalizedName{
					{ Lang: "en", Name: "Over 16" },
					{ Lang: "nl", Name: "Ouder dan 16" },
				},
				Description: []*irmaproto.IRMAIssueSpecification_LocalizedDescription{
					{ Lang: "en", Name: "If you are over 16" },
					{ Lang: "nl", Name: "Of u ouder dan 16 bent" },
				},
			},{
				Id: "over18",
				Name: []*irmaproto.IRMAIssueSpecification_LocalizedName{
					{ Lang: "en", Name: "Over 18" },
					{ Lang: "nl", Name: "Ouder dan 18" },
				},
				Description: []*irmaproto.IRMAIssueSpecification_LocalizedDescription{
					{ Lang: "en", Name: "If you are over 18" },
					{ Lang: "nl", Name: "Of u ouder dan 18 bent" },
				},
			},{
				Id: "over21",
				Name: []*irmaproto.IRMAIssueSpecification_LocalizedName{
					{ Lang: "en", Name: "Over 21" },
					{ Lang: "nl", Name: "Ouder dan 21" },
				},
				Description: []*irmaproto.IRMAIssueSpecification_LocalizedDescription{
					{ Lang: "en", Name: "If you are over 21" },
					{ Lang: "nl", Name: "Of u ouder dan 21 bent" },
				},
			},{
				Id: "over65",
				Name: []*irmaproto.IRMAIssueSpecification_LocalizedName{
					{ Lang: "en", Name: "Over 65" },
					{ Lang: "nl", Name: "Ouder dan 65" },
				},
				Description: []*irmaproto.IRMAIssueSpecification_LocalizedDescription{
					{ Lang: "en", Name: "If you are over 65" },
					{ Lang: "nl", Name: "Of u ouder dan 65 bent" },
				},
			},
		},
	}
	credentialBuffer, err := proto.Marshal(credential);
	fmt.Printf("Size of credential: %d\n", len(credentialBuffer));

	credentialTransaction, err := contract.AddIssuerCredential(auth, issuerId, credentialId, "credentiallogo.jpg", credentialBuffer)
	if err != nil {
		log.Fatalf("could not add credential to contract: %v", err)
	}
	fmt.Printf("Transaction cost %d gas\n\n", credentialTransaction.Gas())
	sim.Commit();



	// Reading scheme info from blockchain
	fmt.Printf("SCHEME INFO:\n\n")

	schemeMetadataBuffer, err := contract.Metadata(nil);
	if err != nil {
		log.Fatalf("could not get metadata: %v", err);
	}
	schemeMetadata := &irmaproto.IRMASchemeMetadata{}
	proto.Unmarshal(schemeMetadataBuffer, schemeMetadata)
	fmt.Printf("Metadata: %+v\n", schemeMetadata)

	numIssuers, err := contract.GetNumberOfIssuers(nil)
	fmt.Printf("Number of issuers: %+v\n", numIssuers.Sign())
	for i := 0; i < numIssuers.Sign(); i++ {
		issuerId, err := contract.IssuerIds(nil, big.NewInt(int64(i)));
		fmt.Printf(">> ISSUER %+v\n", i)
		if err != nil {
			log.Fatalf("could not get issuer with id: %v", i);
		}
		_, logoUrl, ownerAddress, issuerMetadataBuffer, numKeys, numCredentials, err := contract.GetIssuerById(nil, issuerId)
		if err != nil {
			log.Fatalf("could not get issuer: %v", err);
		}
		fmt.Printf("Issuer id: %+v\n", issuerId)
		fmt.Printf("Issuer logoURl: %+v\n", logoUrl)
		fmt.Printf("Issuer address: %+v\n", ownerAddress.String())
		fmt.Printf("Issuer numKeys: %+v\n", numKeys)
		fmt.Printf("Issuer numCredentials: %+v\n", numCredentials)

		issuerMetadata := &irmaproto.IRMAIssuerMetadata{}
		proto.Unmarshal(issuerMetadataBuffer, issuerMetadata)
		fmt.Printf("Issuer metadata: %+v\n", issuerMetadata)

		fmt.Printf("Issuer keys:\n")
		for i := 0; i < numKeys.Sign(); i++ {
			_, keyBuffer, err := contract.GetIssuerPublicKeyById(nil, issuerId, big.NewInt(int64(i)))
			if err != nil {
				log.Fatalf("could not get issuer key: %v", i);
			}
			key := &irmaproto.IRMAIssuerPublicKey{}
			proto.Unmarshal(keyBuffer, key)
			fmt.Printf("-- %+v\n", key)
		}

		fmt.Printf("Issuer credentials:\n")
		for i := 0; i < numKeys.Sign(); i++ {
			credentialId, credentialLogo, issueSpecBuffer, err := contract.GetIssuerCredentialIdByCredentialIndex(nil, issuerId, big.NewInt(int64(i)))
			if err != nil {
				log.Fatalf("could not get issuer credential: %v", i);
			}
			issueSpec := &irmaproto.IRMAIssueSpecification{}
			proto.Unmarshal(issueSpecBuffer, issueSpec)
			fmt.Printf("-- [%+v] Logo: %+v Spec: %+v\n", credentialId, credentialLogo, issueSpec)
		}
	}

}